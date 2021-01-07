package cmd

import (
	"os"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

var logger zerolog.Logger
var client algoliasearch.Client
var index algoliasearch.Index
var baseIndexName string
var appID string
var appKey string

var RootCmd = &cobra.Command{
	Use: "search",
	Run: func(cmd *cobra.Command, args []string) {

	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		client = algoliasearch.NewClient(appID, appKey)
		index = client.InitIndex(baseIndexName)
	},
}

func init() {
	logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr})
	RootCmd.PersistentFlags().StringVar(&baseIndexName, "base-index", os.Getenv("BLOGSEARCH_BASE_INDEX"), "Name of the base index")
	RootCmd.PersistentFlags().StringVar(&appID, "app-id", os.Getenv("BLOGSEARCH_APP_ID"), "Algolia app ID")
	RootCmd.PersistentFlags().StringVar(&appKey, "app-key", os.Getenv("BLOGSEARCH_APP_KEY"), "Algolia app key")
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		logger.Fatal().Err(err).Msg("blogsearch failed")
	}
}
