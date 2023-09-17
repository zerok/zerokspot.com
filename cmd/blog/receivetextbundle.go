package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/spf13/cobra"
	"gitlab.com/zerok/zerokspot.com/pkg/textbundleimporter"
	"gitlab.com/zerok/zerokspot.com/pkg/textbundlereceiver"
)

var repoPath string
var token string
var addr string
var githubUser string
var githubToken string
var tzName string

var receiveTextBundle = &cobra.Command{
	Use: "receive-textbundle",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		var err error
		tz := time.UTC
		tz, err = time.LoadLocation(tzName)
		if err != nil {
			return fmt.Errorf("failed to load timezone %s: %w", tzName, err)
		}
		srv := http.Server{}
		srv.Addr = addr
		imp := textbundleimporter.New(repoPath)
		imp.TimeLocation = tz
		recv := textbundlereceiver.New(func(r *textbundlereceiver.Receiver) {
			r.RepoPath = repoPath
			r.Importer = imp
			r.AccessToken = token
			r.GitHubUser = githubUser
			r.GitHubToken = githubToken
		})
		srv.Handler = recv
		slog.InfoContext(ctx, fmt.Sprintf("Listening on %s", srv.Addr))
		return srv.ListenAndServe()
	},
}

func init() {
	rootCmd.AddCommand(receiveTextBundle)
	receiveTextBundle.Flags().StringVar(&repoPath, "repo-path", "", "Path to the repository that should be updated")
	receiveTextBundle.Flags().StringVar(&tzName, "timezone", "Europe/Vienna", "Timezone to be used with in generated markdown files")
	receiveTextBundle.Flags().StringVar(&token, "token", "", "Token required for requests")
	receiveTextBundle.Flags().StringVar(&addr, "addr", "localhost:37080", "Address to listen on")
	receiveTextBundle.Flags().StringVar(&githubUser, "github-user", "", "Username to be used for creating a pull-request")
	receiveTextBundle.Flags().StringVar(&githubToken, "github-token", "", "Token to be used for creating a pull-request")
}
