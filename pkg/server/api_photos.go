package server

import (
	"net/http"
	"time"

	"github.com/rs/zerolog"
)

func (srv *Server) handleUploadPhoto(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := zerolog.Ctx(ctx).With().Str("handler", "api/upload-photo").Logger()
	ctx = logger.WithContext(ctx)
	if err := r.ParseMultipartForm(10 * 1024 * 1024); err != nil {
		logger.Error().Err(err).Msg("Failed to parse form data")
		http.Error(w, "Invalid multipart data", http.StatusBadRequest)
		return
	}
	photos, ok := r.MultipartForm.File["photo"]
	if !ok {
		http.Error(w, "No photos included", http.StatusBadRequest)
		return
	}
	fp, err := photos[0].Open()
	if err != nil {
		logger.Error().Err(err).Msg("Failed to read photo")
		http.Error(w, "Could not read photo", http.StatusBadRequest)
		return
	}
	if err := srv.photoStore.Write(ctx, time.Now(), photos[0].Filename, fp); err != nil {
		logger.Error().Err(err).Msg("Failed to store photo")
		http.Error(w, "Could not save photo", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (srv *Server) handleGetPhoto(w http.ResponseWriter, r *http.Request) {
}
