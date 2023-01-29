package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gosimple/slug"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

type goodReadsEntry struct {
	Title     string
	Author    string
	DateRead  *time.Time
	DateAdded *time.Time
}

func generateImportGoodReadsCSV() *cobra.Command {
	var outputFolder string
	var year int
	var importGoodReadsCSVCmd = &cobra.Command{
		Use: "import-goodreads-csv",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := logger.WithContext(cmd.Context())
			ctx = findParentTrace(ctx)
			ctx, span := tracer.Start(ctx, "import-goodreads-csv")
			defer span.End()
			if len(args) < 1 {
				return fmt.Errorf("no file provided")
			}
			fp, err := os.Open(args[0])
			if err != nil {
				return fmt.Errorf("failed to open file: %w", err)
			}
			defer fp.Close()
			r := csv.NewReader(fp)
			i := 0
			for {
				rec, err := r.Read()
				if err == io.EOF {
					break
				}
				if err != nil {
					return fmt.Errorf("failed to read record: %w", err)
				}
				if i == 0 {
					i++
					continue
				}
				i++
				e := goodReadsEntry{}
				e.Title = rec[1]
				e.Author = rec[2]
				if rec[14] != "" {
					dr, err := time.Parse("2006/01/02", rec[14])
					if err != nil {
						return fmt.Errorf("failed to parse read-date: %w", err)
					}
					e.DateRead = &dr
				}
				if rec[15] != "" {
					da, err := time.Parse("2006/01/02", rec[15])
					if err != nil {
						return fmt.Errorf("failed to parse read-added: %w", err)
					}
					e.DateAdded = &da
				}

				if e.DateRead != nil {
					if year == 0 || e.DateRead.Year() == year {
						filename := filepath.Join(outputFolder, generateBookFilename(e)+".md")
						_, err := os.Stat(filename)
						if err == nil {
							fmt.Println(filename + " already exists.")
							continue
						}
						if os.IsNotExist(err) {
							if err := ioutil.WriteFile(filename, []byte(generateBookContent(e)), 0600); err != nil {
								return err
							}
						}
					}
				}
			}
			return nil
		},
	}
	importGoodReadsCSVCmd.Flags().StringVar(&outputFolder, "output-folder", "content/reading", "Folder where markdown files are written to")
	importGoodReadsCSVCmd.Flags().IntVar(&year, "year", 0, "Year to import")
	return importGoodReadsCSVCmd
}

func init() {
	rootCmd.AddCommand(generateImportGoodReadsCSV())
}

func generateBookFilename(e goodReadsEntry) string {
	t := strings.Split(e.Title, " (")[0]
	return slug.Make(t)
}

func generateBookContent(e goodReadsEntry) string {
	header := make(map[string]interface{})
	header["title"] = e.Title
	header["type"] = "book"
	if e.Author != "" {
		header["author"] = e.Author
	}
	if e.DateAdded != nil {
		header["date"] = e.DateAdded.Format(time.RFC3339)
	}
	if e.DateRead != nil {
		header["finished"] = e.DateRead.Format(time.RFC3339)
	}
	hout := bytes.Buffer{}
	yaml.NewEncoder(&hout).Encode(header)
	return "---\n" + hout.String() + "---\n"
}
