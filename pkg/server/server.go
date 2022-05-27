package server

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"gitlab.com/zerok/zerokspot.com/pkg/photostore"
	"gitlab.com/zerok/zerokspot.com/pkg/resizer"
)

type Server struct {
	router        chi.Router
	publicBaseURL string
	dataFolder    string
	photoStore    *photostore.Store
	apiKey        string
	resizer       resizer.Resizer
}

type ServerOption func(*Server)

func WithPublicBaseURL(u string) ServerOption {
	return func(srv *Server) {
		srv.publicBaseURL = u
	}
}

func WithDataFolder(u string) ServerOption {
	return func(srv *Server) {
		srv.dataFolder = u
	}
}

func WithAPIKey(key string) ServerOption {
	return func(srv *Server) {
		srv.apiKey = key
	}
}

func WithResizer(r resizer.Resizer) ServerOption {
	return func(srv *Server) {
		srv.resizer = r
	}
}

func NewServer(ctx context.Context, options ...ServerOption) (*Server, error) {
	srv := &Server{}
	for _, o := range options {
		o(srv)
	}
	if err := srv.setup(ctx); err != nil {
		return nil, err
	}
	return srv, nil
}

func (srv *Server) requireAPIKey() func(http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			key := r.Header.Get("Authorization")
			if key != srv.apiKey {
				http.Error(w, "Invalid authorization key", http.StatusUnauthorized)
				return
			}
			handler.ServeHTTP(w, r)
		})
	}
}

func (srv *Server) setup(ctx context.Context) error {
	r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		r.With(srv.requireAPIKey()).Post("/photos/", srv.handleUploadPhoto)
		r.Get("/photos/{year:\\d\\d\\d\\d}/{month:\\d\\d}/{day:\\d\\d}/{filename}", srv.handleGetPhoto)
	})
	srv.router = r
	store, err := photostore.New(srv.dataFolder + "/photos")
	if err != nil {
		return err
	}
	srv.photoStore = store
	return nil
}

func (srv *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	srv.router.ServeHTTP(w, r)
}
