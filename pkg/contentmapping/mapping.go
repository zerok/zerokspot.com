package contentmapping

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/ulikunitz/xz"
)

type PageData struct {
	ObjectID string `json:"objectID"`
	File     string `json:"file"`
	DateYear int    `json:"date_year"`
}

type ContentMapping map[string]string

func Load(rd io.Reader) (ContentMapping, error) {
	r, err := xz.NewReader(rd)
	if err != nil {
		return nil, fmt.Errorf("failed to create new xz reader: %w", err)
	}
	mapping := ContentMapping{}
	if err := json.NewDecoder(r).Decode(&mapping); err != nil {
		return nil, fmt.Errorf("failed to decode mapping file: %w", err)
	}
	return mapping, nil
}

func LoadFromFile(path string) (ContentMapping, error) {
	fp, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open %s for writing: %w", path, err)
	}
	defer fp.Close()
	return Load(fp)
}

func LoadFromURL(u string) (ContentMapping, error) {
	resp, err := http.Get(u)
	if err != nil {
		return nil, fmt.Errorf("failed to execute GET request for %s: %w", u, err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET request for %s returned %d: %w", u, resp.StatusCode, err)
	}
	defer resp.Body.Close()
	return Load(resp.Body)
}

func SaveToFile(path string, mapping ContentMapping) error {
	fp, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	w, err := xz.NewWriter(fp)
	if err := json.NewEncoder(w).Encode(mapping); err != nil {
		w.Close()
		fp.Close()
		return err
	}
	if err := w.Close(); err != nil {
		return err
	}
	return fp.Close()
}

func BuildMapping(root string) (ContentMapping, error) {
	result := make(map[string]string)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.Name() == "index.json" {
			var page PageData
			fp, err := os.Open(path)
			if err != nil {
				return err
			}
			defer fp.Close()
			if err := json.NewDecoder(fp).Decode(&page); err != nil {
				return err
			}
			if strings.HasPrefix(page.ObjectID, "/weblog") && page.DateYear <= 1 {
				return nil
			}
			result[page.File] = page.ObjectID
		}
		return nil
	})
	return result, err
}
