package cmd

import (
	"github.com/spf13/cobra"
)

var deleteObjectCmd = &cobra.Command{
	Use:   "delete-object",
	Short: "Removes a single object by ID",
	Run: func(c *cobra.Command, args []string) {
		if len(args) < 1 {
			logger.Fatal().Msg("You have to specify an ObjectID")
		}
		if _, err := index.DeleteObject(args[0]); err != nil {
			logger.Fatal().Err(err).Msg("Failed to delete object")
		}
	},
}

func init() {
	RootCmd.AddCommand(deleteObjectCmd)
}
