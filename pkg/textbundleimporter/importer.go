package textbundleimporter

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/zerok/textbundle-go"
	"gopkg.in/yaml.v2"
)

var imageLine = regexp.MustCompile(`!\[(.*)\]\(([^)]+)\)`)

type Importer struct {
	RepoPath     string
	TimeLocation *time.Location
	Now          time.Time
}

func New(repopath string) *Importer {
	return &Importer{
		RepoPath: repopath,
	}
}

type FrontMatter struct {
	Title string   `yaml:"title"`
	Date  string   `yaml:"date"`
	Tags  []string `yaml:"tags"`
}

func (i *Importer) processBody(text string, now time.Time) string {
	fm := FrontMatter{
		Date: now.Format(time.RFC3339),
	}
	lines := []string{}
	inputLines := strings.Split(text, "\n")
	for idx, line := range inputLines {
		if idx == 0 && strings.HasPrefix(line, "# ") {
			fm.Title = strings.TrimPrefix(line, "# ")
		} else {
			if len(inputLines)-1 == idx && (strings.HasPrefix(line, "#") || strings.HasPrefix(line, "%% #")) {
				line = strings.TrimPrefix(line, "%% ")
				for _, t := range strings.Split(line, " ") {
					tag := strings.TrimPrefix(t, "#")
					if len(tag) > 0 {
						fm.Tags = append(fm.Tags, tag)
					}
				}
				continue
			} else if strings.HasPrefix(line, "![") {
				if imageLine.MatchString(line) {
					elems := imageLine.FindStringSubmatch(line)
					img := elems[2]
					if strings.HasPrefix(img, "assets/") {
						img = fmt.Sprintf("/media/%s/%s", now.Format("2006"), strings.TrimPrefix(img, "assets/"))
					}
					lines = append(lines, fmt.Sprintf("<figure><img src=\"%s\"><figcaption>%s</figcaption></figure>", img, elems[1]))
					continue
				}
			}
			lines = append(lines, line)
		}
	}
	out := bytes.Buffer{}
	out.WriteString("---\n")
	yaml.NewEncoder(&out).Encode(fm)
	out.WriteString("---\n")
	out.WriteString(strings.Join(lines, "\n"))
	return out.String()
}

func (i *Importer) Import(ctx context.Context, path string, slug string) error {
	logger := zerolog.Ctx(ctx)
	var now time.Time
	defaultTime := time.Time{}
	if i.Now == defaultTime {
		now = time.Now()
	} else {
		now = i.Now
	}
	if i.TimeLocation != nil {
		now = now.In(i.TimeLocation)
	}
	r, err := textbundle.OpenReader(path)
	if err != nil {
		return err
	}
	// Create markdown file
	logger.Info().Msgf("Setting date of new post to %s", now.Format(time.RFC3339))
	folder := filepath.Join(i.RepoPath, fmt.Sprintf("content/weblog/%s", now.Format("2006")))
	if slug == "" {
		elems := strings.SplitN(filepath.Base(path), ".", 2)
		if len(elems) < 1 {
			return fmt.Errorf("invalid file name")
		}
		slug = elems[0]
	}
	filename := fmt.Sprintf("%s.md", slug)
	fpath := filepath.Join(folder, filename)
	logger.Info().Msgf("Creating %s\n", fpath)
	if err := os.MkdirAll(folder, 0700); err != nil {
		return fmt.Errorf("could not create %s: %w", folder, err)
	}
	fp, err := os.OpenFile(fpath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Errorf("could not create %s: %w", fpath, err)
	}
	if _, err = fp.Write([]byte(i.processBody(string(r.Text), now))); err != nil {
		fp.Close()
		return err
	}
	fp.Close()

	// Create asset files
	for _, a := range r.Assets {
		fpath = filepath.Join(i.RepoPath, fmt.Sprintf("static/media/%s/%s", now.Format("2006"), a.Name))
		logger.Info().Msgf("Creating %s\n", fpath)
		if err := os.MkdirAll(filepath.Dir(fpath), 0755); err != nil {
			return fmt.Errorf("Failed to create media folder: %w", err)
		}
		ain, err := a.File.Open()
		if err != nil {
			return err
		}
		fp, err := os.OpenFile(fpath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
		if err != nil {
			ain.Close()
			return err
		}
		if _, err := io.Copy(fp, ain); err != nil {
			fp.Close()
			ain.Close()
			return err
		}
		fp.Close()
		ain.Close()
	}
	return err
}
