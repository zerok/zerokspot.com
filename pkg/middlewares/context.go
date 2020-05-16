package middlewares

import (
	"net/http"

	"github.com/rs/zerolog"
)

func InjectLogger(next http.Handler, logger zerolog.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r.WithContext(logger.WithContext(r.Context())))
	})
}
