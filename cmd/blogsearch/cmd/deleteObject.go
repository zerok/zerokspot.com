package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteObjectCmd = &cobra.Command{
	Use:   "delete-object",
	Short: "Removes a single object by ID",
	RunE: func(c *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("you have to specify an ObjectID")
		}
		if _, err := index.DeleteObject(args[0]); err != nil {
			return fmt.Errorf("failed to delete object: %w", err)
		}
		return nil
	},
}

func init() {
	RootCmd.AddCommand(deleteObjectCmd)
}
