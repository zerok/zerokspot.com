package cmd

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/algolia/algoliasearch-client-go/algoliasearch"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/spf13/cobra"
)

var httpAddr string

var serveCmd = &cobra.Command{
	Use: "serve",
	PreRun: func(cmd *cobra.Command, args []string) {

	},
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		router := chi.NewRouter()

		corsMiddleware := cors.New(cors.Options{
			AllowedOrigins: []string{"http://localhost:1313", "https://zerokspot.com"},
			AllowedMethods: []string{http.MethodGet, http.MethodOptions},
		})
		router.Use(corsMiddleware.Handler)

		router.Get("/search/{query}", func(w http.ResponseWriter, r *http.Request) {
			queryString := chi.URLParam(r, "query")
			index := client.InitIndex(baseIndexName)
			res, err := index.Search(queryString, algoliasearch.Map{
				"hitsPerPage":           20,
				"attributesToHighlight": []string{},
				"attributesToRetrieve": []string{
					"title", "date", "tags", "url",
				},
			})
			if err != nil {
				slog.ErrorContext(ctx, "Search query failed", slog.Any("err", err))
				http.Error(w, "Search query failed", http.StatusInternalServerError)
			}
			w.Header().Set("Content-type", "application/json")
			json.NewEncoder(w).Encode(res)
		})
		router.Get("/year/{year}", func(w http.ResponseWriter, r *http.Request) {
			year := chi.URLParam(r, "year")
			index := client.InitIndex(fmt.Sprintf("%s-datesorted", baseIndexName))
			res, err := index.Search("", algoliasearch.Map{
				"filters":               fmt.Sprintf("date_year:%s", year),
				"facets":                []string{"tags"},
				"hitsPerPage":           1000,
				"attributesToHighlight": []string{},
				"attributesToRetrieve": []string{
					"title", "date", "tags", "url",
				},
			})
			if err != nil {
				slog.ErrorContext(ctx, "Year query failed", slog.Any("err", err))
				http.Error(w, "Year query failed", http.StatusInternalServerError)
			}
			w.Header().Set("Content-type", "application/json")
			json.NewEncoder(w).Encode(res)
		})
		server := http.Server{}
		server.Handler = router
		server.Addr = httpAddr
		slog.InfoContext(ctx, "Listening", slog.String("httpAddr", httpAddr))
		if err := server.ListenAndServe(); err != nil {
			slog.ErrorContext(ctx, "Failed to start server", slog.Any("err", err))
			os.Exit(1)
		}
	},
}

func init() {
	serveCmd.Flags().StringVar(&httpAddr, "http-addr", "0.0.0.0:8000", "Address to listen for requests")
	RootCmd.AddCommand(serveCmd)
}
