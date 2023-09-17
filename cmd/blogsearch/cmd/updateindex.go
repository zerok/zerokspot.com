package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log/slog"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/ulikunitz/xz"
)

var gitrevURL string
var mappingURL string
var rebuild bool
var dryRun bool
var updatedObjectsPath string

func fetchGitRev(u string) (string, error) {
	c := http.Client{}
	resp, err := c.Get(u)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}

type fileChangeStatus struct {
	Name   string
	Status string
}

func (f *fileChangeStatus) IsContent() bool {
	return strings.HasPrefix(f.Name, "content/weblog/") && strings.HasSuffix(f.Name, ".md")
}

func (f *fileChangeStatus) IsLayout() bool {
	return strings.HasPrefix(f.Name, "layout") && strings.HasSuffix(f.Name, ".json")
}

func changedFilesSince(ref string) ([]fileChangeStatus, error) {
	var stdout bytes.Buffer
	result := make([]fileChangeStatus, 0, 5)
	cmd := exec.Command("git", "diff", "--name-status", ref)
	cmd.Stdout = &stdout
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	for _, line := range strings.Split(stdout.String(), "\n") {
		if line != "" {
			elems := strings.SplitN(line, "\t", 2)
			if len(elems) != 2 {
				continue
			}
			result = append(result, fileChangeStatus{
				Name:   elems[1],
				Status: elems[0],
			})
		}
	}
	return result, nil
}

func loadMapping(u string) (map[string]string, error) {
	c := http.Client{}
	resp, err := c.Get(u)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to fetch %s", u)
	}
	defer resp.Body.Close()
	r, err := xz.NewReader(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to generate xz reader")
	}
	mapping := make(map[string]string)
	if err := json.NewDecoder(r).Decode(&mapping); err != nil {
		return nil, errors.Wrap(err, "failed to decode JSON")
	}
	return mapping, nil
}

var updateIndexCmd = &cobra.Command{
	Use: "update-index",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		updatedObjectIDs := make([]string, 0, 5)
		if !rebuild {
			// Fetch a reference commit ID from a designed URL
			gitrev, err := fetchGitRev(gitrevURL)
			if err != nil {
				slog.ErrorContext(ctx, "Failed to retrieve gitrev", slog.Any("err", err))
				os.Exit(1)
			}
			newMapping, err := buildMapping("public")
			if err != nil {
				slog.ErrorContext(ctx, "Failed to generate new mapping", slog.Any("err", err))
				os.Exit(1)
			}
			mapping, err := loadMapping(mappingURL)
			if err != nil {
				slog.ErrorContext(ctx, "Failed to retrieve mapping", slog.Any("err", err))
				os.Exit(1)
			}
			changedFiles, err := changedFilesSince(gitrev)
			if err != nil {
				slog.ErrorContext(ctx, "Failed to retrieve changes", slog.Any("err", err))
				os.Exit(1)
			}
			for _, file := range changedFiles {
				logger := slog.With(slog.String("file_name", file.Name), slog.String("file_status", file.Status))
				if file.IsContent() {
					objectID := mapping[strings.TrimPrefix(file.Name, "content/")]
					logger := logger.With(slog.String("object_id", objectID))
					switch file.Status {
					case "D":
						if objectID == "" {
							continue
						}
						logger.InfoContext(ctx, "Deleting object")
						if dryRun {
							continue
						}
						updatedObjectIDs = append(updatedObjectIDs, objectID)
						if _, err := index.DeleteObject(objectID); err != nil {
							logger.ErrorContext(ctx, "Failed to delete object", slog.Any("err", err))
							os.Exit(1)
						}
					case "M":
						fallthrough
					case "A":
						if objectID == "" {
							objectID = newMapping[strings.TrimPrefix(file.Name, "content/")]
						}
						dataPath := filepath.Join("public", objectID, "index.json")
						raw, err := os.ReadFile(dataPath)
						if err != nil {
							if os.IsNotExist(err) {
								logger.InfoContext(ctx, "Deleting object")
								if _, err := index.DeleteObject(objectID); err != nil {
									logger.ErrorContext(ctx, "Failed to delete object", slog.Any("err", err))
								}
								objectID = newMapping[strings.TrimPrefix(file.Name, "content/")]
								dataPath = filepath.Join("public", objectID, "index.json")
								raw, err = os.ReadFile(dataPath)
							}
							if err != nil {
								logger.ErrorContext(ctx, "Failed to load file", slog.String("dataPath", dataPath), slog.Any("err", err))
								os.Exit(1)
							}
						}
						var obj algoliasearch.Object
						if err := json.Unmarshal(raw, &obj); err != nil {
							logger.ErrorContext(ctx, "Failed to decode JSON", slog.String("dataPath", dataPath), slog.Any("err", err))
							os.Exit(1)
						}
						if err := validateSearchObject(obj); err != nil {
							logger.WarnContext(ctx, "Document not valid.", slog.Any("err", err))
							continue
						}
						if dryRun {
							continue
						}
						logger.InfoContext(ctx, "Updating object")
						if _, err := index.UpdateObjects([]algoliasearch.Object{obj}); err != nil {
							slog.ErrorContext(ctx, "Failed to update object", slog.Any("err", err))
							os.Exit(1)
						}
						updatedObjectIDs = append(updatedObjectIDs, objectID)
					}
				}
			}
		} else {
			// Determine the file changes between that
			// previous commit ID and now
			var objects []algoliasearch.Object
			err := filepath.Walk("public", func(path string, info os.FileInfo, err error) error {
				if !info.IsDir() && info.Name() == "index.json" {
					var obj algoliasearch.Object
					fp, err := os.Open(path)
					if err != nil {
						return err
					}
					defer fp.Close()
					if err := json.NewDecoder(fp).Decode(&obj); err != nil {
						return err
					}
					// Let's ignore notes for now
					// until I'm sure I actually
					// want them in the index.
					if obj["type"] != "weblog" {
						return nil
					}
					slog.InfoContext(ctx, "Adding file", slog.String("dataPath", path))
					objects = append(objects, obj)
				}
				return nil
			})
			if err != nil {
				slog.ErrorContext(ctx, "Failed to load JSON data", slog.Any("err", err))
				os.Exit(1)
			}
			if dryRun {
				return
			}
			res, err := index.UpdateObjects(objects)
			if err != nil {
				slog.ErrorContext(ctx, "Failed to update index", slog.Any("err", err))
				os.Exit(1)
			}
			for _, oid := range res.ObjectIDs {
				updatedObjectIDs = append(updatedObjectIDs, oid)
			}
		}
		slog.InfoContext(ctx, "Creating updated-objects file", slog.String("outputFile", updatedObjectsPath))
		fp, err := os.OpenFile(updatedObjectsPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
		if err != nil {
			slog.ErrorContext(ctx, "Failed to create updated-objects file", slog.String("outputFile", updatedObjectsPath), slog.Any("err", err))
			os.Exit(1)
		}
		defer fp.Close()
		for _, oid := range updatedObjectIDs {
			fmt.Fprintf(fp, "%s\n", oid)
		}
	},
}

func init() {
	updateIndexCmd.Flags().StringVar(&gitrevURL, "gitrev-url", "https://zerokspot.com/.gitrev", "URL to retrieve the reference Git ref from")
	updateIndexCmd.Flags().StringVar(&mappingURL, "mapping-url", "https://zerokspot.com/.mapping.json.xz", "URL to retrieve the objectID mapping from")
	updateIndexCmd.Flags().BoolVar(&rebuild, "rebuild", false, "Rebuild the whole index")
	updateIndexCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Do everything BUT changing the index")
	updateIndexCmd.Flags().StringVar(&updatedObjectsPath, "updated-objects-path", "updated-objects.txt", "File listing updated objects")
	RootCmd.AddCommand(updateIndexCmd)
}
