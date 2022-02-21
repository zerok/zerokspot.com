package bookscollection

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/gohugoio/hugo/parser/pageparser"
	"gopkg.in/yaml.v2"
)

type Book struct {
	Title         string     `yaml:"title,omitempty"`
	Author        string     `yaml:"author,omitempty"`
	Date          *time.Time `yaml:"date,omitempty"`
	StartedDate   *time.Time `yaml:"started,omitempty"`
	FinishedDate  *time.Time `yaml:"finished,omitempty"`
	ISBN          string     `yaml:"isbn,omitempty"`
	GoodreadsID   string     `yaml:"goodreadsID,omitempty"`
	OpenLibraryID string     `yaml:"openlibraryID,omitempty"`
	Genre         string     `yaml:"genre,omitempty"`
	Pages         int        `yaml:"pages,omitempty"`
	Rating        int        `yaml:"rating,omitempty"`
	Body          string     `yaml:"-"`
}

func (b *Book) ReadingDur() time.Duration {
	if b.FinishedDate == nil || b.StartedDate == nil {
		return -1
	}
	return b.FinishedDate.Sub(*b.StartedDate)
}

func LoadBooks(ctx context.Context) ([]*Book, error) {
	var books []*Book
	paths, err := filepath.Glob("content/reading/*.md")
	if err != nil {
		return nil, err
	}
	for _, path := range paths {
		fp, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		book, err := ParseBook(fp)
		fp.Close()
		books = append(books, book)
	}
	return books, nil

}

func ParseBook(r io.Reader) (*Book, error) {
	res, err := pageparser.ParseFrontMatterAndContent(r)
	if err != nil {
		return nil, err
	}
	book := &Book{}
	fm := res.FrontMatter
	s, err := getFieldAsString(fm, "title")
	if err != nil {
		return nil, err
	}
	book.Title = s
	s, err = getFieldAsString(fm, "author")
	if err != nil {
		return nil, err
	}
	book.Author = s
	s, err = getFieldAsString(fm, "genre")
	if err != nil {
		return nil, err
	}
	book.Genre = s
	s, err = getFieldAsString(fm, "isbn")
	if err != nil {
		return nil, err
	}
	book.ISBN = s
	s, err = getFieldAsString(fm, "openlibraryID")
	if err != nil {
		return nil, err
	}
	book.OpenLibraryID = s
	i, err := getFieldAsInt(fm, "pages")
	if err != nil {
		return nil, err
	}
	book.Pages = i
	t, err := getFieldAsTime(fm, "date")
	if err != nil {
		return nil, err
	}
	book.Date = t
	t, err = getFieldAsTime(fm, "started")
	if err != nil {
		return nil, err
	}
	book.StartedDate = t
	t, err = getFieldAsTime(fm, "finished")
	if err != nil {
		return nil, err
	}
	book.FinishedDate = t
	i, err = getFieldAsInt(fm, "rating")
	if err != nil {
		return nil, err
	}
	book.Rating = i
	book.Body = string(res.Content)
	return book, nil
}

func WriteBook(ctx context.Context, outpath string, book *Book) error {
	out := bytes.Buffer{}
	out.WriteString("---\n")
	enc := yaml.NewEncoder(&out)
	if err := enc.Encode(book); err != nil {
		return err
	}
	out.WriteString("---\n")
	out.WriteString(book.Body)
	fp, err := os.OpenFile(outpath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer fp.Close()
	_, err = io.Copy(fp, &out)
	return err
}

func getFieldAsString(m map[string]interface{}, k string) (string, error) {
	v, ok := m[k]
	if !ok {
		return "", nil
	}
	value, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("%s is not a string", k)
	}
	return value, nil
}

func getFieldAsTime(m map[string]interface{}, k string) (*time.Time, error) {
	v, err := getFieldAsString(m, k)
	if err != nil {
		return nil, err
	}
	if v == "" {
		return nil, nil
	}
	t, err := time.Parse(time.RFC3339, v)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func getFieldAsInt(m map[string]interface{}, k string) (int, error) {
	v, ok := m[k]
	if !ok {
		return 0, nil
	}
	value, ok := v.(int)
	if !ok {
		return 0, fmt.Errorf("%s is not a int", k)
	}
	return value, nil
}
