package cmd

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

type pageData struct {
	ObjectID string `json:"objectID"`
	File     string `json:"file"`
	DateYear int    `json:"date_year"`
}

func buildMapping(root string) (map[string]string, error) {
	result := make(map[string]string)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.Name() == "index.json" {
			var page pageData
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
