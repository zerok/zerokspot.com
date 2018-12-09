package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
		if !rebuild {
			// Fetch a reference commit ID from a designed URL
			gitrev, err := fetchGitRev(gitrevURL)
			if err != nil {
				logger.Fatal().Err(err).Msg("Failed to retrieve gitrev")
			}
			newMapping, err := buildMapping("public")
			if err != nil {
				logger.Fatal().Err(err).Msg("Failed to generate new mapping")
			}
			mapping, err := loadMapping(mappingURL)
			if err != nil {
				logger.Fatal().Err(err).Msg("Failed to retrieve mapping")
			}
			changedFiles, err := changedFilesSince(gitrev)
			if err != nil {
				logger.Fatal().Err(err).Msg("Failed to retrieve changes")
			}
			logger.Info().Msgf("Changes since %s", gitrev)
			for _, file := range changedFiles {
				logger.Info().Msgf(" - %s (%s, is_content: %v, is_layout: %v)", file.Name, file.Status, file.IsContent(), file.IsLayout())
				if file.IsContent() {
					objectID := mapping[strings.TrimPrefix(file.Name, "content/")]
					switch file.Status {
					case "D":
						logger.Info().Msgf("Deleting %s", objectID)
						if _, err := index.DeleteObject(objectID); err != nil {
							logger.Fatal().Err(err).Msgf("Failed to delete %s", objectID)
						}
					case "M":
						fallthrough
					case "A":
						if objectID == "" {
							objectID = newMapping[strings.TrimPrefix(file.Name, "content/")]
						}
						logger.Info().Msgf("Updating %s", objectID)
						dataPath := filepath.Join("public", objectID, "index.json")
						raw, err := ioutil.ReadFile(dataPath)
						if err != nil {
							logger.Fatal().Err(err).Msgf("Failed to load %s", dataPath)
						}
						var obj algoliasearch.Object
						if err := json.Unmarshal(raw, &obj); err != nil {
							logger.Fatal().Err(err).Msgf("Failed to decode JSON from %s", dataPath)
						}
						if _, err := index.UpdateObjects([]algoliasearch.Object{obj}); err != nil {
							logger.Fatal().Err(err).Msgf("Failed to update %s", objectID)
						}
					}
				}
			}
		} else {
			// Determine the file changes between that previous commit ID and now
			var objects []algoliasearch.Object
			err := filepath.Walk("public", func(path string, info os.FileInfo, err error) error {
				fmt.Println(path)
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
					objects = append(objects, obj)
				}
				return nil
			})
			if err != nil {
				logger.Fatal().Err(err).Msg("Failed to load JSON data")
			}
			_, err = index.UpdateObjects(objects)
			if err != nil {
				logger.Fatal().Err(err).Msg("Failed to update index")
			}
		}
	},
}

func init() {
	updateIndexCmd.Flags().StringVar(&gitrevURL, "gitrev-url", "https://zerokspot.com/.gitrev", "URL to retrieve the reference Git ref from")
	updateIndexCmd.Flags().StringVar(&mappingURL, "mapping-url", "https://zerokspot.com/.mapping.json.xz", "URL to retrieve the objectID mapping from")
	updateIndexCmd.Flags().BoolVar(&rebuild, "rebuild", false, "Rebuild the whole index")
	rootCmd.AddCommand(updateIndexCmd)
}
