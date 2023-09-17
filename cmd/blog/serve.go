package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"gitlab.com/zerok/zerokspot.com/pkg/resizer"
	"gitlab.com/zerok/zerokspot.com/pkg/server"
)

func generateServeCmd() *cobra.Command {
	var addr string
	var publicBaseURL string
	var dataFolder string
	cmd := &cobra.Command{
		Use: "serve",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			apiKey := os.Getenv("ZS_API_KEY")
			profiles := generatePhotoProfiles()
			r := resizer.NewMagickResizer(dataFolder, profiles)
			srv, err := server.NewServer(ctx,
				server.WithPublicBaseURL(publicBaseURL),
				server.WithDataFolder(dataFolder),
				server.WithAPIKey(apiKey),
				server.WithResizer(r),
			)
			if err != nil {
				return err
			}

			hs := http.Server{}
			hs.Addr = addr
			hs.Handler = srv
			slog.InfoContext(ctx, fmt.Sprintf("Starting server on %s", addr))
			return hs.ListenAndServe()
		},
	}
	cmd.Flags().StringVar(&addr, "addr", "localhost:8080", "Address to listen on")
	cmd.Flags().StringVar(&publicBaseURL, "base-url", "https://zerokspot.com", "Base URL of this installation")
	cmd.Flags().StringVar(&dataFolder, "data-folder", "./data", "Folder where files etc. are stored")
	return cmd
}
