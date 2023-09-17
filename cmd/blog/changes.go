package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"gitlab.com/zerok/zerokspot.com/pkg/contentmapping"
	"gitlab.com/zerok/zerokspot.com/pkg/repostate"
)

var sinceRev string
var showAll bool
var asURL bool
var baseURL string

var changesCmd = &cobra.Command{
	Use: "changes",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := findParentTrace(cmd.Context())
		ctx, span := tracer.Start(ctx, "cmd:changes")
		defer span.End()
		var out bytes.Buffer
		currentMapping, err := contentmapping.LoadFromFile("public/.mapping.json.xz")
		if err != nil {
			return err
		}
		previousMapping, err := contentmapping.LoadFromURL("https://zerokspot.com/.mapping.json.xz")
		if err != nil {
			return err
		}
		s := repostate.RepoState{}
		changes, err := s.ChangedFilesSince(ctx, sinceRev)
		if err != nil {
			return err
		}
		for _, change := range changes {
			if showAll || (change.IsArticle() || change.IsNote()) {
				if asURL {
					objectID, ok := currentMapping[strings.TrimPrefix(change.Name, "content/")]
					if !ok {
						objectID, ok = previousMapping[strings.TrimPrefix(change.Name, "content/")]
						if !ok {
							return fmt.Errorf("%s not found in current mapping", change.Name)
						}
					}
					fmt.Fprintf(&out, "%s%s\n", baseURL, objectID)
				} else {
					fmt.Fprintln(&out, change.Name)
				}
			}
		}

		if outputFile == "-" {
			io.Copy(os.Stdout, &out)
		} else {
			fp, err := os.OpenFile(outputFile, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
			if err != nil {
				return err
			}
			if _, err := io.Copy(fp, &out); err != nil {
				return err
			}
			return fp.Close()
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(changesCmd)
	changesCmd.Flags().StringVar(&sinceRev, "since-rev", "", "Git rev of the previous state")
	changesCmd.Flags().StringVar(&baseURL, "base-url", "https://zerokspot.com", "Base URL")
	changesCmd.Flags().StringVar(&outputFile, "output", "-", "Output file")
	changesCmd.Flags().BoolVar(&showAll, "all", false, "List all changes")
	changesCmd.Flags().BoolVar(&asURL, "url", false, "List all changes as URLs")
}
