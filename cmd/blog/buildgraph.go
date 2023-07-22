package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"gitlab.com/zerok/zerokspot.com/pkg/bloggraph"
)

type doc struct {
	ObjectID string `json:"objectID"`
	Content  string `json:"content"`
	File     string `json:"file"`
}

var buildGraphCmd = &cobra.Command{
	Use: "build-graph",
	RunE: func(command *cobra.Command, args []string) error {
		ctx := logger.WithContext(command.Context())
		ctx = findParentTrace(ctx)
		ctx, span := tracer.Start(ctx, "cmd:build-graph")
		defer span.End()
		if err := os.MkdirAll("data", 0700); err != nil {
			return err
		}
		mapping, err := bloggraph.BuildMapping(ctx, ".")
		data, err := json.Marshal(mapping)
		if err != nil {
			return err
		}
		return ioutil.WriteFile(filepath.Join("data", "postpaths.json"), data, 0600)
	},
}

func init() {
	rootCmd.AddCommand(buildGraphCmd)
}
