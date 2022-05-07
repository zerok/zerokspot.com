package photostore

import (
	"context"
	"io"
	"os"
	"time"

	"github.com/spf13/afero"
)

type Store struct {
	fs       afero.Fs
	rootPath string
}

func New(fs afero.Fs, rootPath string) (*Store, error) {
	if err := fs.MkdirAll(rootPath, 0755); err != nil {
		return nil, err
	}
	return &Store{
		fs:       fs,
		rootPath: rootPath,
	}, nil
}

func (s *Store) Write(ctx context.Context, ts time.Time, filename string, data io.Reader) error {
	targetFolder := s.rootPath + "/" + ts.Format("2006/01/02")
	if err := s.fs.MkdirAll(targetFolder, 0755); err != nil {
		return err
	}
	fp, err := s.fs.OpenFile(targetFolder+"/"+filename, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer fp.Close()
	if _, err := io.Copy(fp, data); err != nil {
		return err
	}
	return nil
}
