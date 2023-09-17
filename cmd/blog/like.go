package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var title string
var tags []string

type noteData struct {
	Title     string   `yaml:"title,omitempty"`
	LikeOf    string   `yaml:"likeOf,omitempty"`
	Timestamp string   `yaml:"date"`
	Tags      []string `yaml:"tags"`
}

var likeCmd = &cobra.Command{
	Use: "like SLUG URL",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := findParentTrace(cmd.Context())
		ctx, span := tracer.Start(ctx, "like")
		defer span.End()
		if len(args) < 2 {
			return fmt.Errorf("specify at least a URL to like")
		}
		u := args[1]
		slug := args[0]
		now := time.Now()
		ts := now.Format(time.RFC3339)
		folder := fmt.Sprintf("content/notes/%s", now.Format("2006/01/02"))
		if err := os.MkdirAll(folder, 0755); err != nil {
			return err
		}
		buf := bytes.Buffer{}
		if err := yaml.NewEncoder(&buf).Encode(noteData{
			LikeOf:    u,
			Timestamp: ts,
			Title:     title,
			Tags:      tags,
		}); err != nil {
			return err
		}
		fp, err := os.OpenFile(filepath.Join(folder, slug+".md"), os.O_CREATE|os.O_WRONLY, 0600)
		if err != nil {
			return err
		}
		defer fp.Close()
		fmt.Fprintf(fp, "---\n"+buf.String()+"\n---")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(likeCmd)
	likeCmd.Flags().StringVar(&title, "title", "", "Title of the like")
	likeCmd.Flags().StringSliceVar(&tags, "tag", []string{}, "Tag(s) of the like")
}
