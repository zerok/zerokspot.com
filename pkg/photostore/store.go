package photostore

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"time"
)

type Store struct {
	rootPath string
}

func New(rootPath string) (*Store, error) {
	if err := os.MkdirAll(rootPath, 0755); err != nil {
		return nil, err
	}
	return &Store{
		rootPath: rootPath,
	}, nil
}

func (s *Store) GetResizedPath(ctx context.Context, path string, profileFileName string) (string, error) {
	targetFilePath := filepath.Join(s.rootPath, "resized", path, profileFileName)
	_, err := os.Stat(targetFilePath)
	if err != nil {
		return "", err
	}
	return targetFilePath, nil
}

func (s *Store) Exists(ctx context.Context, path string) bool {
	targetFilePath := filepath.Join(s.rootPath, path)
	_, err := os.Stat(targetFilePath)
	if err != nil {
		return false
	}
	return true
}

func (s *Store) Write(ctx context.Context, ts time.Time, filename string, data io.Reader) (string, error) {
	localPath := filepath.Join(ts.Format("2006/01/02"), filename)
	targetFolder := s.rootPath + "/" + ts.Format("2006/01/02")
	if err := os.MkdirAll(targetFolder, 0755); err != nil {
		return "", err
	}
	fp, err := os.OpenFile(targetFolder+"/"+filename, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return "", err
	}
	defer fp.Close()
	if _, err := io.Copy(fp, data); err != nil {
		return "", err
	}
	return localPath, nil
}
