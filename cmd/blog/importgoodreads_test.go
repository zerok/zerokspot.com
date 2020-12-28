package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGoodReadsFileNaming(t *testing.T) {
	t.Run("simple-title", func(t *testing.T) {
		e := goodReadsEntry{}
		e.Title = "Hello World"
		require.Equal(t, "hello-world", generateBookFilename(e))
	})

	t.Run("title-with-series", func(t *testing.T) {
		e := goodReadsEntry{}
		e.Title = "Hello World (Strange new worlds #123)"
		require.Equal(t, "hello-world", generateBookFilename(e))
	})
}

func TestGoodReadsContent(t *testing.T) {
	t.Run("simple", func(t *testing.T) {
		e := goodReadsEntry{}
		e.Title = "Hello World (Strange new worlds #123)"
		require.Equal(t, "---\ntitle: 'Hello World (Strange new worlds #123)'\n---\n", generateBookContent(e))
	})
}
