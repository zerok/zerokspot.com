package server

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

	"log/slog"

	"github.com/go-chi/chi/v5"
)

func (srv *Server) handleUploadPhoto(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := slog.With(slog.String("handler", "api/upload-photo"))
	if err := r.ParseMultipartForm(10 * 1024 * 1024); err != nil {
		logger.ErrorContext(ctx, "Failed to parse form data", slog.String("err", err.Error()))
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
		logger.ErrorContext(ctx, "Failed to read photo", slog.String("err", err.Error()))
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
		slog.ErrorContext(ctx, "Failed to store photo", slog.String("err", err.Error()))
		http.Error(w, "Could not save photo", http.StatusInternalServerError)
		return
	}
	if err := srv.resize(ctx, targetPath); err != nil {
		slog.ErrorContext(ctx, "Failed to resize photo", slog.String("err", err.Error()))
		http.Error(w, "Could not resize photo", http.StatusInternalServerError)
	}
	absoluteURL := srv.publicBaseURL + "/api/photos/" + targetPath
	w.Header().Set("Location", absoluteURL)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, absoluteURL)
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
	w.WriteHeader(http.StatusOK)
	for _, pname := range srv.resizer.GetProfiles().Names() {
		fmt.Fprintf(w, "<a href=\"?profile=%s\">%s</a><br>", pname, pname)
	}
}

func (srv *Server) resize(ctx context.Context, photoPath string) error {
	return srv.resizer.Resize(ctx, photoPath)
}
