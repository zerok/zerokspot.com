package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/spf13/cobra"
	"gitlab.com/zerok/zerokspot.com/pkg/receiver"
)

var receive = &cobra.Command{
	Use: "receive",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		var err error
		tz := time.UTC
		tz, err = time.LoadLocation(tzName)
		if err != nil {
			return fmt.Errorf("failed to load timezone %s: %w", tzName, err)
		}
		if err := requireStringFlags(cmd, "repo-path", "token", "github-token"); err != nil {
			return err
		}
		srv := http.Server{}
		srv.Addr = addr
		recv := receiver.New(ctx, func(cfg *receiver.Configuration) {
			cfg.RepoPath = repoPath
			cfg.TimeLocation = tz
			cfg.GitHubToken = githubToken
			cfg.AccessToken = token
		})
		srv.Handler = recv
		slog.InfoContext(ctx, fmt.Sprintf("Listening on %s", srv.Addr))
		go func() {
			<-ctx.Done()
			srv.Shutdown(context.Background())
		}()
		return srv.ListenAndServe()
	},
}

func init() {
	rootCmd.AddCommand(receive)
	receive.Flags().StringVar(&repoPath, "repo-path", "", "Path to the repository that should be updated")
	receive.Flags().StringVar(&tzName, "timezone", "Europe/Vienna", "Timezone to be used with in generated markdown files")
	receive.Flags().StringVar(&token, "token", "", "Token required for requests")
	receive.Flags().StringVar(&addr, "addr", "localhost:37080", "Address to listen on")
	receive.Flags().StringVar(&githubUser, "github-user", "", "Username to be used for creating a pull-request")
	receive.Flags().StringVar(&githubToken, "github-token", "", "Token to be used for creating a pull-request")
}
