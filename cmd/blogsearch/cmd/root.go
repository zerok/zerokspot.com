package cmd

import (
	"context"
	"log/slog"
	"os"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/spf13/cobra"
)

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
	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))
	slog.SetDefault(logger)
	RootCmd.PersistentFlags().StringVar(&baseIndexName, "base-index", os.Getenv("BLOGSEARCH_BASE_INDEX"), "Name of the base index")
	RootCmd.PersistentFlags().StringVar(&appID, "app-id", os.Getenv("BLOGSEARCH_APP_ID"), "Algolia app ID")
	RootCmd.PersistentFlags().StringVar(&appKey, "app-key", os.Getenv("BLOGSEARCH_APP_KEY"), "Algolia app key")
}

func Execute() {
	ctx := context.Background()
	if err := RootCmd.ExecuteContext(ctx); err != nil {
		slog.ErrorContext(ctx, "blogsearch failed", slog.Any("err", err))
		os.Exit(1)
	}
}
