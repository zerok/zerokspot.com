package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

var logger zerolog.Logger

var rootCmd = &cobra.Command{
	Use: "blog",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr})
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		logger.Fatal().Msg(err.Error())
	}
}
