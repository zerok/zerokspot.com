package main

import (
	"fmt"
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
		currentMapping, err := contentmapping.LoadFromFile("public/.mapping.json.xz")
		if err != nil {
			return err
		}
		previousMapping, err := contentmapping.LoadFromURL("https://zerokspot.com/.mapping.json.xz")
		if err != nil {
			return err
		}
		s := repostate.RepoState{}
		changes, err := s.ChangedFilesSince(sinceRev)
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
					fmt.Printf("%s%s\n", baseURL, objectID)
				} else {
					fmt.Println(change.Name)
				}
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(changesCmd)
	changesCmd.Flags().StringVar(&sinceRev, "since-rev", "", "Git rev of the previous state")
	changesCmd.Flags().StringVar(&baseURL, "base-url", "https://zerokspot.com", "Base URL")
	changesCmd.Flags().BoolVar(&showAll, "all", false, "List all changes")
	changesCmd.Flags().BoolVar(&asURL, "url", false, "List all changes as URLs")
}
