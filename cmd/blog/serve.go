package main

import (
	"context"
	"net/http"

	"github.com/spf13/cobra"
	"gitlab.com/zerok/zerokspot.com/pkg/server"
)

func generateServeCmd() *cobra.Command {
	var addr string
	var publicBaseURL string
	cmd := &cobra.Command{
		Use: "serve",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := logger.WithContext(context.Background())
			srv, err := server.NewServer(ctx, server.WithPublicBaseURL(publicBaseURL)
			if err != nil {
				return err
			}

			hs := http.Server{}
			hs.Addr = addr
			hs.Handler = srv
			logger.Info().Msgf("Starting server on %s", addr)
			return hs.ListenAndServe()
		},
	}
	cmd.Flags().StringVar(&addr, "addr", "localhost:8080", "Address to listen on")
	cmd.Flags().StringVar(&publicBaseURL, "base-url", "https://zerokspot.com", "Base URL of this installation")
	return cmd
}
