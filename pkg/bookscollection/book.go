package bookscollection

import (
	"fmt"
	"io"
	"time"

	"github.com/gohugoio/hugo/parser/pageparser"
)

type Book struct {
	Title        string     `yaml:"title"`
	Author       string     `yaml:"author"`
	Pages        int        `yaml:"pages"`
	ISBN         string     `yaml:"isbn"`
	Genre        string     `yaml:"genre"`
	GoodreadsID  string     `yaml:"goodreadsID"`
	FinishedDate *time.Time `yaml:"finished"`
	StartedDate  *time.Time `yaml:"started"`
	Date         *time.Time `yaml:"date"`
}

func (b *Book) ReadingDur() time.Duration {
	if b.FinishedDate == nil || b.StartedDate == nil {
		return -1
	}
	return b.FinishedDate.Sub(*b.StartedDate)
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
	return book, nil
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
