package receiver_test

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"gitlab.com/zerok/zerokspot.com/pkg/receiver"
)

func TestReceiverMarkdownHandling(t *testing.T) {
	t.Run("markdown-file-without-path", func(t *testing.T) {
		ctx := context.Background()
		tmpDir := t.TempDir()
		now := time.Date(2023, 8, 8, 12, 0, 0, 0, time.UTC)
		loc, _ := time.LoadLocation("Europe/Vienna")
		recv := receiver.New(ctx, func(cfg *receiver.Configuration) {
			cfg.SkipCommit = true
			cfg.SkipPull = true
			cfg.RepoPath = tmpDir
			cfg.ForceNow = now
			cfg.TimeLocation = loc
		})
		testSrv := httptest.NewServer(recv)
		content := `---
title: "something"
slug: "something"
---

Some content`
		expectedContent := `---
date: "2023-08-08T14:00:00+02:00"
title: something
---

Some content`
		expectedResponseBody := `{"path":"content/weblog/2023/something.md"}
`
		resp, err := http.Post(testSrv.URL, "text/markdown", bytes.NewBufferString(content))
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		// Also ensure that the path is returned from the server
		respBody, err := io.ReadAll(resp.Body)
		require.NoError(t, err)
		require.Equal(t, expectedResponseBody, string(respBody))
		require.FileExists(t, filepath.Join(tmpDir, "content/weblog/2023/something.md"))
		rawFileContent, err := ioutil.ReadFile(filepath.Join(tmpDir, "content/weblog/2023/something.md"))
		require.NoError(t, err)
		fileContent := string(rawFileContent)
		require.Equal(t, expectedContent, fileContent)
	})
}
