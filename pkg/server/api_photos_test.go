package server

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/rs/zerolog"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
)

const apiKey = "test"

func copyFile(out io.Writer, inPath string) error {
	fp, err := os.Open(inPath)
	if err != nil {
		return err
	}
	defer fp.Close()
	if _, err := io.Copy(out, fp); err != nil {
		return err
	}
	return nil
}

func TestAPIUploadPhoto(t *testing.T) {
	t.Run("no-api-key", func(t *testing.T) {
		ctx := context.Background()
		fs := afero.NewMemMapFs()
		srv, err := NewServer(ctx, WithDataFolder("/data"), WithFS(fs), WithAPIKey(apiKey))
		require.NoError(t, err)
		require.NotNil(t, srv)

		req := httptest.NewRequest(http.MethodPost, "/api/photos/", nil)
		resp := httptest.NewRecorder()
		srv.ServeHTTP(resp, req)
		require.Equal(t, http.StatusUnauthorized, resp.Result().StatusCode)
	})
	t.Run("no-data", func(t *testing.T) {
		ctx := context.Background()
		fs := afero.NewMemMapFs()
		srv, err := NewServer(ctx, WithDataFolder("/data"), WithFS(fs), WithAPIKey(apiKey))
		require.NoError(t, err)
		require.NotNil(t, srv)

		req := httptest.NewRequest(http.MethodPost, "/api/photos/", nil)
		req.Header.Add("Authorization", apiKey)
		resp := httptest.NewRecorder()
		srv.ServeHTTP(resp, req)
		require.Equal(t, http.StatusBadRequest, resp.Result().StatusCode)
	})

	t.Run("happy", func(t *testing.T) {
		ctx := context.Background()
		logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr})
		ctx = logger.WithContext(ctx)
		fs := afero.NewMemMapFs()
		srv, err := NewServer(ctx, WithDataFolder("/data"), WithFS(fs), WithAPIKey(apiKey))
		require.NoError(t, err)
		require.NotNil(t, srv)

		// Let's create a multipart form and put a file into it:
		data := bytes.Buffer{}
		form := multipart.NewWriter(&data)
		photoFile, err := form.CreateFormFile("photo", "photo.jpg")
		require.NoError(t, err)
		require.NoError(t, copyFile(photoFile, "../../static/images/me.jpg"))
		require.NoError(t, form.Close())

		req := httptest.NewRequest(http.MethodPost, "/api/photos/", &data).WithContext(ctx)
		req.Header.Add("content-type", "multipart/form-data;boundary="+form.Boundary())
		req.Header.Add("Authorization", apiKey)
		resp := httptest.NewRecorder()
		srv.ServeHTTP(resp, req)
		require.Equal(t, http.StatusCreated, resp.Result().StatusCode)

		// Now check if the file also exists in the FS
		exists, err := afero.Exists(fs, "/data/photos/"+time.Now().Format("2006/01/02")+"/photo.jpg")
		require.NoError(t, err)
		require.True(t, exists)
	})
}
