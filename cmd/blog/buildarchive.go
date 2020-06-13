package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	hugodeps "github.com/gohugoio/hugo/deps"
	hugofs "github.com/gohugoio/hugo/hugofs"
	"github.com/gohugoio/hugo/hugolib"
	page "github.com/gohugoio/hugo/resources/page"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

type yearCtx struct {
	Title     string   `yaml:"title"`
	Year      string   `yaml:"year"`
	Date      string   `yaml:"date"`
	Paths     []string `yaml:"paths"`
	URL       string   `yaml:"url"`
	NumPhotos int      `yaml:"numPhotos"`
	NumNotes  int      `yaml:"numNotes"`
	NumPosts  int      `yaml:"numPosts"`
}

type monthCtx struct {
	Title     string   `yaml:"title"`
	Year      string   `yaml:"year"`
	Month     string   `yaml:"month"`
	Date      string   `yaml:"date"`
	Paths     []string `yaml:"paths"`
	URL       string   `yaml:"url"`
	NumPhotos int      `yaml:"numPhotos"`
	NumNotes  int      `yaml:"numNotes"`
	NumPosts  int      `yaml:"numPosts"`
}

var buildArchiveCmd = &cobra.Command{
	Use: "build-archive",
	RunE: func(cmd *cobra.Command, args []string) error {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		fs := afero.NewOsFs()
		hfs := hugofs.NewFrom(fs, &viper.Viper{})
		dcfg := &hugodeps.DepsCfg{
			Fs: hfs,
		}
		dcfg.Cfg, _, err = hugolib.LoadConfig(hugolib.ConfigSourceDescriptor{
			Fs:         fs,
			Filename:   "config.toml",
			Path:       wd,
			WorkingDir: wd,
		})
		if err != nil {
			return err
		}
		sites, err := hugolib.NewHugoSites(*dcfg)
		if err != nil {
			return err
		}
		if err := sites.Build(hugolib.BuildCfg{
			ResetState: true,
			SkipRender: true,
		}); err != nil {
			return err
		}

		years := make(map[string][]page.Page)
		months := make(map[string][]page.Page)
		pages := sites.Pages().Reverse()
		for _, p := range pages {
			if p.IsNode() {
				continue
			}
			if sec := p.Section(); sec == "weblog" || sec == "notes" || sec == "photos" {
				t := p.Date().In(localZone)
				y := t.Format("2006")
				m := t.Format("2006-01")
				posts := years[y]
				posts = append(posts, p)
				years[y] = posts

				posts = months[m]
				posts = append(posts, p)
				months[m] = posts
			}
		}
		rootFolder := filepath.Join(wd, "content", "archive")
		if err := os.RemoveAll(rootFolder); err != nil {
			return err
		}
		for year, pages := range years {
			f := filepath.Join(rootFolder, year, "_index.md")
			if err := os.MkdirAll(filepath.Join(rootFolder, year), 0700); err != nil {
				return err
			}
			paths := make([]string, 0, len(pages))
			numPhotos := 0
			numNotes := 0
			numPosts := 0
			for _, p := range pages {
				paths = append(paths, p.Path())
				switch p.Section() {
				case "photos":
					numPhotos += 1
				case "weblog":
					numPosts += 1
				case "notes":
					numNotes += 1
				}
			}
			if err := writeFile(f, yearCtx{
				Title:     year,
				Year:      year,
				Date:      year,
				Paths:     paths,
				NumPosts:  numPosts,
				NumPhotos: numPhotos,
				NumNotes:  numNotes,
				URL:       fmt.Sprintf("archive/%s/", year),
			}); err != nil {
				return err
			}
		}

		for m, pages := range months {
			elems := strings.Split(m, "-")
			year := elems[0]
			month := elems[1]
			f := filepath.Join(rootFolder, year, month+".md")
			paths := make([]string, 0, len(pages))
			numPhotos := 0
			numNotes := 0
			numPosts := 0
			for _, p := range pages {
				paths = append(paths, p.Path())
				switch p.Section() {
				case "photos":
					numPhotos += 1
				case "weblog":
					numPosts += 1
				case "notes":
					numNotes += 1
				}
			}
			if err := writeFile(f, monthCtx{
				Title:     fmt.Sprintf("%s // %s", year, month),
				Year:      year,
				Date:      year,
				Month:     month,
				Paths:     paths,
				NumPosts:  numPosts,
				NumPhotos: numPhotos,
				NumNotes:  numNotes,
				URL:       fmt.Sprintf("archive/%s/%s/", year, month),
			}); err != nil {
				return err
			}
		}
		return nil
	},
}

func writeFile(path string, data interface{}) error {
	fp, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer fp.Close()
	if _, err := fp.WriteString("---\n"); err != nil {
		return err
	}
	if err := yaml.NewEncoder(fp).Encode(data); err != nil {
		return err
	}
	if _, err := fp.WriteString("\n---\n"); err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(buildArchiveCmd)
}
