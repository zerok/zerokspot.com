package main

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

	"github.com/spf13/cobra"
	"github.com/zerok/textbundle-go"
	"gopkg.in/yaml.v2"
)

var imageLine = regexp.MustCompile(`!\[(.*)\]\(([^)]+)\)`)

var importTextBundleCmd = &cobra.Command{
	Use: "import-textbundle FILE",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := logger.WithContext(cmd.Context())
		ctx = findParentTrace(ctx)
		ctx, span := tracer.Start(ctx, "import-textbundle")
		defer span.End()
		if len(args) < 1 {
			return fmt.Errorf("please specify at least one import file")
		}
		for _, f := range args {
			if err := importTextBundle(ctx, f); err != nil {
				return fmt.Errorf("failed to import %s: %w", f, err)
			}
		}
		return nil
	},
}

type FrontMatter struct {
	Title string   `yaml:"title"`
	Date  string   `yaml:"date"`
	Tags  []string `yaml:"tags"`
}

func processBody(text string, now time.Time) string {
	fm := FrontMatter{
		Date: now.Format(time.RFC3339),
	}
	lines := []string{}
	inputLines := strings.Split(text, "\n")
	for idx, line := range inputLines {
		if idx == 0 && strings.HasPrefix(line, "# ") {
			fm.Title = strings.TrimPrefix(line, "# ")
		} else {
			if len(inputLines)-1 == idx && strings.HasPrefix(line, "#") {
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

func importTextBundle(ctx context.Context, path string) error {
	now := time.Now()
	r, err := textbundle.OpenReader(path)
	if err != nil {
		return err
	}
	// Create markdown file
	folder := fmt.Sprintf("content/weblog/%s", now.Format("2006"))
	elems := strings.SplitN(filepath.Base(path), ".", 2)
	if len(elems) < 1 {
		return fmt.Errorf("invalid file name")
	}
	filename := fmt.Sprintf("%s.md", elems[0])
	fpath := filepath.Join(folder, filename)
	fmt.Printf("Creating %s\n", fpath)
	fp, err := os.OpenFile(fpath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Errorf("could not create %s: %w", fpath, err)
	}
	if _, err = fp.Write([]byte(processBody(string(r.Text), now))); err != nil {
		fp.Close()
		return err
	}
	fp.Close()

	// Create asset files
	for _, a := range r.Assets {
		fpath = fmt.Sprintf("static/media/%s/%s", now.Format("2006"), a.Name)
		fmt.Printf("Creating %s\n", fpath)
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

func init() {
	rootCmd.AddCommand(importTextBundleCmd)
}
