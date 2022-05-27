package main

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"gitlab.com/zerok/zerokspot.com/cmd/blogsearch/cmd"
)

var logger zerolog.Logger
var localZoneName string
var localZone *time.Location

var rootCmd = &cobra.Command{
	Use: "blog",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr})
		localZone, err = time.LoadLocation(localZoneName)
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.SilenceUsage = true
	rootCmd.PersistentFlags().StringVar(&localZoneName, "tz", "Europe/Vienna", "Timezone to be used for data-relevant processing")
	rootCmd.AddCommand(cmd.RootCmd)
	rootCmd.AddCommand(generateServeCmd())
	rootCmd.AddCommand(generateResizePhotosCmd())
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		logger.Fatal().Msg(err.Error())
	}
}
