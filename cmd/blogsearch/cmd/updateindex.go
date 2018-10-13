package cmd

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/spf13/cobra"
)

var updateIndexCmd = &cobra.Command{
	Use: "update-index",
	Run: func(cmd *cobra.Command, args []string) {
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
	},
}

func init() {
	rootCmd.AddCommand(updateIndexCmd)
}
