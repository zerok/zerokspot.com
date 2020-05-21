package textbundleimporter

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/zerok/textbundle-go"
)

func TestImportBundle(t *testing.T) {
	ctx := context.Background()
	t.Run("repo-missing", func(t *testing.T) {
		i := New("does/not/exist")
		err := i.Import(ctx, "missing.textpack", "")
		require.Error(t, err)
	})

	t.Run("file-missing", func(t *testing.T) {
		path := createTestRepo(t, "something")
		i := New(path)
		err := i.Import(ctx, "missing.textpack", "")
		require.Error(t, err)
	})

	t.Run("timezone", func(t *testing.T) {
		now := time.Now()
		path := createTestRepo(t, "something")
		fp, err := os.OpenFile(filepath.Join(path, "test.textpack"), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
		require.NoError(t, err)
		w := textbundle.NewWriter(fp)
		w.SetText("md", "Hello world")
		w.Close()
		fp.Close()

		i := New(path)
		// First with UTC
		i.Now = now.In(time.UTC)
		err = i.Import(ctx, filepath.Join(path, "test.textpack"), "")
		require.NoError(t, err)
		expectedTS := i.Now.Format(time.RFC3339)
		expectedFile := filepath.Join(path, "content", "weblog", i.Now.Format("2006"), "test.md")
		fileContent, err := ioutil.ReadFile(expectedFile)
		require.NoError(t, err)
		require.Contains(t, string(fileContent), expectedTS)

		// Now with Europe/Vienna
		tz, err := time.LoadLocation("Europe/Vienna")
		require.NoError(t, err, "Europe/Vienna not found")
		i.Now = now.In(tz)
		err = i.Import(ctx, filepath.Join(path, "test.textpack"), "")
		require.NoError(t, err)
		expectedTS = i.Now.Format(time.RFC3339)
		expectedFile = filepath.Join(path, "content", "weblog", i.Now.Format("2006"), "test.md")
		fileContent, err = ioutil.ReadFile(expectedFile)
		require.NoError(t, err)
		require.Contains(t, string(fileContent), expectedTS)

		// Now with Europe/Vienna and then overriden with UTC again
		i.Now = now.In(tz)
		i.TimeLocation = time.UTC
		err = i.Import(ctx, filepath.Join(path, "test.textpack"), "")
		require.NoError(t, err)
		expectedTS = i.Now.In(time.UTC).Format(time.RFC3339)
		t.Logf("Expected TS: %s", expectedTS)
		expectedFile = filepath.Join(path, "content", "weblog", i.Now.Format("2006"), "test.md")
		fileContent, err = ioutil.ReadFile(expectedFile)
		require.NoError(t, err)
		require.Contains(t, string(fileContent), expectedTS)
	})
}

func createTestRepo(t *testing.T, name string) string {
	t.Helper()
	path := filepath.Join("testdata", name)
	require.NoError(t, os.RemoveAll(path))
	require.NoError(t, os.MkdirAll(path, 0700))
	return path
}
