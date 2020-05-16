package main

import (
	"net/http"

	"github.com/spf13/cobra"
	"gitlab.com/zerok/zerokspot.com/pkg/middlewares"
	"gitlab.com/zerok/zerokspot.com/pkg/textbundleimporter"
	"gitlab.com/zerok/zerokspot.com/pkg/textbundlereceiver"
)

var repoPath string
var token string
var addr string
var githubUser string
var githubToken string

var receiveTextBundle = &cobra.Command{
	Use: "receive-textbundle",
	RunE: func(cmd *cobra.Command, args []string) error {
		srv := http.Server{}
		srv.Addr = addr
		imp := textbundleimporter.New(repoPath)
		recv := textbundlereceiver.New(func(r *textbundlereceiver.Receiver) {
			r.RepoPath = repoPath
			r.Importer = imp
			r.AccessToken = token
			r.GitHubUser = githubUser
			r.GitHubToken = githubToken
		})
		srv.Handler = middlewares.InjectLogger(recv, logger)
		logger.Info().Msgf("Listening on %s", srv.Addr)
		return srv.ListenAndServe()
	},
}

func init() {
	rootCmd.AddCommand(receiveTextBundle)
	receiveTextBundle.Flags().StringVar(&repoPath, "repo-path", "", "Path to the repository that should be updated")
	receiveTextBundle.Flags().StringVar(&token, "token", "", "Token required for requests")
	receiveTextBundle.Flags().StringVar(&addr, "addr", "localhost:37080", "Address to listen on")
	receiveTextBundle.Flags().StringVar(&githubUser, "github-user", "", "Username to be used for creating a pull-request")
	receiveTextBundle.Flags().StringVar(&githubToken, "github-token", "", "Token to be used for creating a pull-request")
}
