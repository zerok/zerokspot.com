package photostore

import (
	"bytes"
	"context"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPhotoStoreWrite(t *testing.T) {
	ctx := context.Background()
	dataDir := t.TempDir()
	s, err := New(dataDir)
	ts := time.Date(2022, 5, 7, 10, 00, 00, 00, time.UTC)
	require.NoError(t, err)
	data := bytes.Buffer{}
	require.NoError(t, s.Write(ctx, ts, "test.jpg", &data))

	require.FileExists(t, filepath.Join(dataDir, "/2022/05/07/test.jpg"))
}
