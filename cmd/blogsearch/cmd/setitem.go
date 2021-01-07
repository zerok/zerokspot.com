package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/spf13/cobra"
)

func validateSearchObject(obj algoliasearch.Object) error {
	dyraw, ok := obj["date_year"]
	if !ok {
		return fmt.Errorf("document has no date_year")
	}
	if dy, ok := dyraw.(float64); !ok || dy <= 1 {
		return fmt.Errorf("document's date_year is not an integer")
	}
	return nil
}

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
		if err := validateSearchObject(obj); err != nil {
			logger.Fatal().Err(err).Msg("JSON not valid.")
		}
		if _, err := index.UpdateObjects([]algoliasearch.Object{obj}); err != nil {
			logger.Fatal().Err(err).Msg("Failed to update object")
		}
	},
}

func init() {
	RootCmd.AddCommand(setItemCmd)
}
