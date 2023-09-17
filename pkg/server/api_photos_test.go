package server

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"gitlab.com/zerok/zerokspot.com/pkg/resizer"
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
		rootPath := t.TempDir()
		ctx := context.Background()
		srv, err := NewServer(ctx, WithDataFolder(rootPath), WithAPIKey(apiKey))
		require.NoError(t, err)
		require.NotNil(t, srv)

		req := httptest.NewRequest(http.MethodPost, "/api/photos/", nil)
		resp := httptest.NewRecorder()
		srv.ServeHTTP(resp, req)
		require.Equal(t, http.StatusUnauthorized, resp.Result().StatusCode)
	})
	t.Run("no-data", func(t *testing.T) {
		rootPath := t.TempDir()
		ctx := context.Background()
		srv, err := NewServer(ctx, WithDataFolder(rootPath), WithAPIKey(apiKey))
		require.NoError(t, err)
		require.NotNil(t, srv)

		req := httptest.NewRequest(http.MethodPost, "/api/photos/", nil)
		req.Header.Add("Authorization", apiKey)
		resp := httptest.NewRecorder()
		srv.ServeHTTP(resp, req)
		require.Equal(t, http.StatusBadRequest, resp.Result().StatusCode)
	})

	t.Run("invalid-name", func(t *testing.T) {
		rootPath := t.TempDir()
		ctx := context.Background()
		srv, err := NewServer(ctx, WithDataFolder(rootPath), WithAPIKey(apiKey))
		require.NoError(t, err)
		require.NotNil(t, srv)

		data := bytes.Buffer{}
		form := multipart.NewWriter(&data)
		photoFile, err := form.CreateFormFile("photo", "photo..jpg")
		require.NoError(t, err)
		require.NoError(t, copyFile(photoFile, "../../static/images/me.jpg"))
		require.NoError(t, form.Close())

		req := httptest.NewRequest(http.MethodPost, "/api/photos/", &data).WithContext(ctx)
		req.Header.Add("content-type", "multipart/form-data;boundary="+form.Boundary())
		req.Header.Add("Authorization", apiKey)
		resp := httptest.NewRecorder()
		srv.ServeHTTP(resp, req)
		require.Equal(t, http.StatusBadRequest, resp.Result().StatusCode, "Invalid filenames should be rejected")
	})

	t.Run("happy", func(t *testing.T) {
		rootPath := t.TempDir()
		ctx := context.Background()
		profiles := resizer.NewProfiles()
		profiles.Add("800", resizer.Profile{
			Width:  800,
			Format: resizer.FormatJPEG,
		})
		profiles.Add("400", resizer.Profile{
			Width:  400,
			Format: resizer.FormatJPEG,
		})
		r := resizer.NewMagickResizer(rootPath, profiles)
		srv, err := NewServer(ctx, WithDataFolder(rootPath), WithAPIKey(apiKey), WithResizer(r), WithPublicBaseURL("https://example.org"))
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

		photoPath := time.Now().Format("2006/01/02") + "/photo.jpg"
		loc := resp.Header().Get("Location")
		require.Equal(t, "https://example.org/api/photos/"+photoPath, loc)

		// Now check if the file also exists in the FS
		require.FileExists(t, rootPath+"/photos/"+photoPath)

		require.FileExists(t, rootPath+"/photos/resized/"+photoPath+"/800.jpg", "800px variant should exist")
		require.FileExists(t, rootPath+"/photos/resized/"+photoPath+"/400.jpg", "400px variant should exist")

		// Now use the retrieval API for a full cycle
		req = httptest.NewRequest(http.MethodGet, "/api/photos/"+photoPath+"?profile=800", nil).WithContext(ctx)
		resp = httptest.NewRecorder()
		srv.ServeHTTP(resp, req)
		require.Equal(t, http.StatusOK, resp.Result().StatusCode)
		require.Equal(t, "image/jpeg", resp.Header().Get("Content-Type"))
	})
}

func TestAPIGetPhoto(t *testing.T) {
	rootPath := t.TempDir()
	ctx := context.Background()
	profiles := resizer.NewProfiles()
	profiles.Add("800", resizer.Profile{
		Width:  800,
		Format: resizer.FormatJPEG,
	})
	profiles.Add("400", resizer.Profile{
		Width:  400,
		Format: resizer.FormatJPEG,
	})
	r := resizer.NewMagickResizer(rootPath, profiles)
	srv, err := NewServer(ctx, WithDataFolder(rootPath), WithAPIKey(apiKey), WithResizer(r), WithPublicBaseURL("https://example.org"))
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

	photoPath := time.Now().Format("2006/01/02") + "/photo.jpg"

	t.Run("existing-profile", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/photos/"+photoPath+"?profile=800", nil).WithContext(ctx)
		resp := httptest.NewRecorder()
		srv.ServeHTTP(resp, req)
		require.Equal(t, http.StatusOK, resp.Result().StatusCode)
	})

	t.Run("missing-profile", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/photos/"+photoPath+"?profile=100", nil).WithContext(ctx)
		resp := httptest.NewRecorder()
		srv.ServeHTTP(resp, req)
		require.Equal(t, http.StatusNotFound, resp.Result().StatusCode)
	})

	t.Run("no-profile-selected", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/photos/"+photoPath, nil).WithContext(ctx)
		resp := httptest.NewRecorder()
		srv.ServeHTTP(resp, req)
		require.Equal(t, http.StatusOK, resp.Result().StatusCode)

		require.Equal(t, "text/html", resp.Header().Get("Content-Type"))

		// Check that the rendered HTML page contains links to all the profiles:
		body, err := ioutil.ReadAll(resp.Result().Body)
		require.NoError(t, err)
		sbody := string(body)
		require.Contains(t, sbody, "<a href=\"?profile=800\">")
		require.Contains(t, sbody, "<a href=\"?profile=400\">")
	})
}

func TestValidatePhotoFilename(t *testing.T) {
	tests := []struct {
		Input       string
		ExpectError bool
	}{
		{
			Input:       "hello.jpg",
			ExpectError: false,
		},
		{
			Input:       "hello..jpg",
			ExpectError: true,
		},
		{
			Input:       "h😅llo.jpg",
			ExpectError: true,
		},
		{
			Input:       "h/llo.jpg",
			ExpectError: true,
		},
	}
	for _, test := range tests {
		err := ValidatePhotoFilename(test.Input)
		if test.ExpectError && err != nil {
			continue
		}
		if !test.ExpectError && err == nil {
			continue
		}
		if err != nil {
			t.Errorf("Unexpected error for input `%s`: %s", test.Input, err.Error())
		} else {
			t.Errorf("Missing error for input `%s`", test.Input)
		}
	}
}
