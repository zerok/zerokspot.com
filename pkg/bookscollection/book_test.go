package bookscollection

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParser(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		b := bytes.NewBufferString(`---
title: "title"
genre: "scifi"
author: "Author Name"
pages: 123
isbn: "9780380789030"
date: "2021-01-01T12:34:00Z"
---`)
		book, err := ParseBook(b)
		require.NoError(t, err)
		require.NotNil(t, book)
		require.Equal(t, "title", book.Title)
		require.Equal(t, int64(123), book.Pages)
		require.Equal(t, "scifi", book.Genre)
		require.Equal(t, "Author Name", book.Author)
		require.Equal(t, "9780380789030", book.ISBN)
		require.NotNil(t, book.Date)
	})
}
