package cmd

import (
	"encoding/json"
	"os"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/spf13/cobra"
)

var setItemCmd = &cobra.Command{
	Use:   "set-item",
	Short: "Update a single object based on a JSON file",
	Run: func(c *cobra.Command, args []string) {

		var obj algoliasearch.Object
		fp, err := os.Open(args[0])
		if err != nil {
			logger.Fatal().Err(err).Msgf("Failed to open %s", args[0])
		}
		defer fp.Close()
		if err := json.NewDecoder(fp).Decode(&obj); err != nil {
			logger.Fatal().Err(err).Msgf("Failed to decode %s", args[0])
		}
		if _, err := index.UpdateObjects([]algoliasearch.Object{obj}); err != nil {
			logger.Fatal().Err(err).Msg("Failed to update object")
		}
	},
}

func init() {
	rootCmd.AddCommand(setItemCmd)
}
