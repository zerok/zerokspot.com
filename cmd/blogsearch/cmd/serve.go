package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"

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
				logger.Error().Err(err).Msg("Search query failed")
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
				logger.Error().Err(err).Msg("Year query failed")
				http.Error(w, "Year query failed", http.StatusInternalServerError)
			}
			w.Header().Set("Content-type", "application/json")
			json.NewEncoder(w).Encode(res)
		})
		server := http.Server{}
		server.Handler = router
		server.Addr = httpAddr
		logger.Info().Msgf("Listening on %s", httpAddr)
		if err := server.ListenAndServe(); err != nil {
			logger.Fatal().Err(err).Msg("Failed to start server")
		}
	},
}

func init() {
	serveCmd.Flags().StringVar(&httpAddr, "http-addr", "0.0.0.0:8000", "Address to listen for requests")
	RootCmd.AddCommand(serveCmd)
}
