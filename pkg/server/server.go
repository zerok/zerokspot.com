package server

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
)

type Server struct {
	router        chi.Router
	publicBaseURL string
}

type ServerOption func(*Server)

func WithPublicBaseURL(u string) ServerOption {
	return func(srv *Server) {
		srv.publicBaseURL = u
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

func (srv *Server) setup(ctx context.Context) error {
	r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		r.Post("/photos", srv.handleUploadPhoto)
		r.Get("/photos/{year}/{slug}", srv.handleGetPhoto)
	})
	srv.router = r
	return nil
}

func (srv *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	srv.router.ServeHTTP(w, r)
}
