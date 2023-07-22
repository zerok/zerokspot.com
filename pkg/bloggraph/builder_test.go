package bloggraph_test

import (
	"context"
	"os"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
	"gitlab.com/zerok/zerokspot.com/pkg/bloggraph"
)

func TestBuildMapping(t *testing.T) {
	ctx := context.Background()
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr})
	ctx = logger.WithContext(ctx)

	t.Run("Simple mapping", func(t *testing.T) {
		result, err := bloggraph.BuildMapping(ctx, "testdata/simplemapping")
		require.NoError(t, err)
		require.NotNil(t, result)

		// "new-post" should be inspired by "old-post" and "old-post" should be inspiring "new-post"
		paths, found := result["/weblog/2001/01/01/old-post/"]
		require.True(t, found)
		require.Len(t, paths.Up, 0)
		require.Len(t, paths.Down, 1)
		require.Equal(t, "/weblog/2001/01/02/new-post/", paths.Down[0].ContentID)

		paths, found = result["/weblog/2001/01/02/new-post/"]
		require.True(t, found)
		require.Len(t, paths.Up, 1)
		require.Len(t, paths.Down, 0)
		require.Equal(t, "/weblog/2001/01/01/old-post/", paths.Up[0].ContentID)
	})
}
