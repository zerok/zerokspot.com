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
	RunE: func(c *cobra.Command, args []string) error {
		var obj algoliasearch.Object
		fp, err := os.Open(args[0])
		if err != nil {
			return fmt.Errorf("failed to open %s: %w", args[0], err)
		}
		defer fp.Close()
		if err := json.NewDecoder(fp).Decode(&obj); err != nil {
			return fmt.Errorf("failed to decode %s: %w", args[0], err)
		}
		if err := validateSearchObject(obj); err != nil {
			return fmt.Errorf("JSON not valid: %w", err)
		}
		if _, err := index.UpdateObjects([]algoliasearch.Object{obj}); err != nil {
			return fmt.Errorf("failed to update object: %w", err)
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(setItemCmd)
}
