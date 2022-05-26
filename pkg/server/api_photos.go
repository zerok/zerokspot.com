package server

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/go-chi/chi"
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
	if !ok || len(photos) == 0 {
		http.Error(w, "No photos included", http.StatusBadRequest)
		return
	}
	fp, err := photos[0].Open()
	if err != nil {
		logger.Error().Err(err).Msg("Failed to read photo")
		http.Error(w, "Could not read photo", http.StatusBadRequest)
		return
	}

	fname := photos[0].Filename
	if err := ValidatePhotoFilename(fname); err != nil {
		http.Error(w, "Invalid filename", http.StatusBadRequest)
		return
	}
	targetPath, err := srv.photoStore.Write(ctx, time.Now(), photos[0].Filename, fp)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to store photo")
		http.Error(w, "Could not save photo", http.StatusInternalServerError)
		return
	}
	if err := srv.resize(ctx, targetPath); err != nil {
		logger.Error().Err(err).Msg("Failed to resize photo")
		http.Error(w, "Could not resize photo", http.StatusInternalServerError)
	}
	w.Header().Set("Location", srv.publicBaseURL+"/api/photos/"+targetPath)
	w.WriteHeader(http.StatusCreated)
}

var photoFilenamePattern = regexp.MustCompile("^[a-z0-9A-Z._-]+$")

func ValidatePhotoFilename(fname string) error {
	if strings.Contains(fname, "..") {
		return fmt.Errorf("filename contains multiple successive dots")
	}
	if photoFilenamePattern.MatchString(fname) {
		return nil
	}
	return fmt.Errorf("filename contains invalid characters")
}

func (srv *Server) handleGetPhoto(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := zerolog.Ctx(ctx).With().Str("handler", "api/get-photo").Logger()
	ctx = logger.WithContext(ctx)
	year := chi.URLParam(r, "year")
	month := chi.URLParam(r, "month")
	day := chi.URLParam(r, "day")
	filename := chi.URLParam(r, "filename")
	profileName := r.URL.Query().Get("profile")

	path := fmt.Sprintf("%s/%s/%s/%s", year, month, day, filename)

	if !srv.photoStore.Exists(ctx, path) {
		http.NotFound(w, r)
		return
	}

	// If a profile was requested, then we want to serve the actual image if possible:
	if profileName != "" {
		prof := srv.resizer.GetProfile(profileName)
		if prof == nil {
			http.NotFound(w, r)
			return
		}
		path, err := srv.photoStore.GetResizedPath(ctx, path, prof.Filename())
		if err != nil {
			http.NotFound(w, r)
			return
		}
		http.ServeFile(w, r, path)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	for _, pname := range srv.resizer.GetProfiles().Names() {
		fmt.Fprintf(w, "<a href=\"?profile=%s\">%s</a><br>", pname, pname)
	}
	w.WriteHeader(http.StatusOK)
}

func (srv *Server) resize(ctx context.Context, photoPath string) error {
	return srv.resizer.Resize(ctx, photoPath)
}
