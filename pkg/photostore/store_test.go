package photostore

import (
	"bytes"
	"context"
	"testing"
	"time"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
)

func TestPhotoStoreWrite(t *testing.T) {
	ctx := context.Background()
	fs := afero.NewMemMapFs()
	s, err := New(fs, "/data")
	ts := time.Date(2022, 5, 7, 10, 00, 00, 00, time.UTC)
	require.NoError(t, err)
	data := bytes.Buffer{}
	require.NoError(t, s.Write(ctx, ts, "test.jpg", &data))

	exists, err := afero.Exists(fs, "/data/2022/05/07/test.jpg")
	require.NoError(t, err)
	require.True(t, exists)
}
