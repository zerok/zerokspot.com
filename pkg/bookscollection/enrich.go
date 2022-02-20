package bookscollection

import (
	"context"
	"fmt"
	"os"
	"regexp"

	"gitlab.com/zerok/zerokspot.com/pkg/openlibrary"
)

func EnrichData(ctx context.Context, path string) error {
	fp, err := os.Open(path)
	if err != nil {
		return err
	}
	book, err := ParseBook(fp)
	fp.Close()
	if err != nil {
		return err
	}

	enricher := NewOpenLibraryEnricher()

	if book.ISBN == "" && book.OpenLibraryID == "" {
		return fmt.Errorf("no ID present, cannot fetch additional data")
	}

	if book.ISBN != "" && book.OpenLibraryID == "" {
		if err := enricher.LoadByISBN(ctx, book, book.ISBN); err != nil {
			return err
		}
	}
	if err := WriteBook(ctx, path, book); err != nil {
		return err
	}

	return nil
}

func NewOpenLibraryEnricher() *OpenLibraryEnricher {
	return &OpenLibraryEnricher{
		client: openlibrary.NewClient(),
	}
}

type OpenLibraryEnricher struct {
	client *openlibrary.Client
}

func (e *OpenLibraryEnricher) LoadByISBN(ctx context.Context, book *Book, isbn string) error {
	isbn = normalizeISBN(isbn)
	obook, err := e.client.GetBookByISBN(ctx, isbn)
	if err != nil {
		return err
	}
	book.OpenLibraryID = obook.Key
	return nil
}

func normalizeISBN(in string) string {
	re := regexp.MustCompile("[^0-9]")
	return re.ReplaceAllString(in, "")
}
