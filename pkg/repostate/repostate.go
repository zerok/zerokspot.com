package repostate

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

type RepoState struct {
	PublishedStateURL string
}

func (s *RepoState) FetchPublishedGitRev(ctx context.Context) (string, error) {
	c := http.Client{}
	req, err := http.NewRequest(http.MethodGet, s.PublishedStateURL, nil)
	if err != nil {
		return "", err
	}
	resp, err := c.Do(req.WithContext(ctx))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}

type FileChangeStatus struct {
	Name   string
	Status string
}

func (f *FileChangeStatus) IsArticle() bool {
	return strings.HasPrefix(f.Name, "content/weblog/") && strings.HasSuffix(f.Name, ".md")
}

func (f *FileChangeStatus) IsNote() bool {
	return strings.HasPrefix(f.Name, "content/notes/") && strings.HasSuffix(f.Name, ".md")
}

func (f *FileChangeStatus) IsContent() bool {
	return f.IsArticle()
}

func (f *FileChangeStatus) IsLayout() bool {
	return strings.HasPrefix(f.Name, "layout") && strings.HasSuffix(f.Name, ".json")
}

func (s *RepoState) ChangedFilesSince(ctx context.Context, ref string) ([]FileChangeStatus, error) {
	var stdout bytes.Buffer
	result := make([]FileChangeStatus, 0, 5)
	cmd := exec.CommandContext(ctx, "git", "diff", "--name-status", ref)
	cmd.Stdout = &stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	for _, line := range strings.Split(stdout.String(), "\n") {
		if line != "" {
			elems := strings.SplitN(line, "\t", 2)
			if len(elems) != 2 {
				continue
			}
			result = append(result, FileChangeStatus{
				Name:   elems[1],
				Status: elems[0],
			})
		}
	}
	return result, nil
}
