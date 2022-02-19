package bookscollection

import (
	"context"
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
)

type OPML struct {
	XMLName struct{} `xml:"opml"`
	Version string   `xml:"version,attr"`
	Head    OPMLHead `xml:"head"`
	Body    OPMLBody `xml:"body"`
}

type OPMLHead struct {
	Title string `xml:"title,omitempty"`
}

type OPMLBody struct {
	Outlines []OPMLOutline `xml:"outline"`
}

type OPMLOutline struct {
	Type          string        `xml:"type,attr,omitempty"`
	ISBN          string        `xml:"isbn,attr,omitempty"`
	OpenLibraryID string        `xml:"openLibraryId,attr,omitempty"`
	Text          string        `xml:"text,attr,omitempty"`
	Name          string        `xml:"name,attr,omitempty"`
	Author        string        `xml:"author,attr,omitempty"`
	Outlines      []OPMLOutline `xml:"outline"`
}

func GenerateOPML(ctx context.Context, outputFilePath string, title string, books []*Book) error {
	output := OPML{
		Version: "2.0",
		Head: OPMLHead{
			Title: title,
		},
		Body: OPMLBody{
			Outlines: make([]OPMLOutline, 0, 10),
		},
	}
	for _, b := range books {
		name := b.Title
		if b.Author != "" {
			name = fmt.Sprintf("%s by %s", b.Title, b.Author)
		}
		outline := OPMLOutline{
			Type:          "book",
			Name:          b.Title,
			Author:        b.Author,
			ISBN:          b.ISBN,
			OpenLibraryID: b.OpenLibraryID,
			Text:          name,
		}
		output.Body.Outlines = append(output.Body.Outlines, outline)
	}
	outputFolder := filepath.Dir(outputFilePath)
	if err := os.MkdirAll(outputFolder, 0755); err != nil {
		return err
	}
	fp, err := os.OpenFile(outputFilePath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer fp.Close()
	fp.WriteString(xml.Header)
	enc := xml.NewEncoder(fp)
	enc.Indent("", "  ")
	return enc.Encode(output)
}
