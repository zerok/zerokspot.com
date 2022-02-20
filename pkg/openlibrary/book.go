package openlibrary

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Book struct {
	Key   string `json:"key"`
	Title string `json:"title"`
	Type  Type   `json:"type"`
}

type Type struct {
	Key string `json:"key"`
}

func (c *Client) GetBookByISBN(ctx context.Context, isbn string) (*Book, error) {
	var book *Book
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("https://openlibrary.org/isbn/%s.json", isbn), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code returned: %d", resp.StatusCode)
	}
	err = json.NewDecoder(resp.Body).Decode(&book)
	if err == nil {
		book.Key = strings.TrimPrefix(book.Key, "/books/")
		book.Key = strings.TrimPrefix(book.Key, "/works/")
	}
	return book, err
}
