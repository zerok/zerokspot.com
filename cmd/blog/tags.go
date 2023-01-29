package main

import (
	"fmt"
	"os"

	"github.com/gohugoio/hugo/config"
	hugodeps "github.com/gohugoio/hugo/deps"
	hugofs "github.com/gohugoio/hugo/hugofs"
	"github.com/gohugoio/hugo/hugolib"
	page "github.com/gohugoio/hugo/resources/page"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

func buildSites() (*hugolib.HugoSites, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	fs := afero.NewOsFs()
	c := config.New()
	c.Set("publishDir", "./output")
	c.Set("workingDir", wd)
	hfs := hugofs.NewFrom(fs, c)
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
		return nil, err
	}
	sites, err := hugolib.NewHugoSites(*dcfg)
	if err != nil {
		return nil, err
	}
	if err := sites.Build(hugolib.BuildCfg{
		ResetState: true,
		SkipRender: true,
	}); err != nil {
		return nil, err
	}
	return sites, nil
}

func isContentPage(p page.Page) bool {
	return (p.Section() == "weblog" || p.Section() == "notes") && p.IsPage() && !p.Draft()
}

func generateSingletonTagsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "singletons",
		RunE: func(cmd *cobra.Command, args []string) error {
			allTags := make(map[string][]string)
			sites, err := buildSites()
			if err != nil {
				return err
			}
			for _, p := range sites.Pages() {
				if isContentPage(p) {
					rawTags, err := p.Param("tags")
					if err != nil {
						continue
					}

					if tags, ok := rawTags.([]string); ok {
						for _, t := range tags {
							prev, ok := allTags[t]
							if !ok {
								prev = make([]string, 0, 5)
							}
							prev = append(prev, p.File().Path())
							allTags[t] = prev
						}
					}
				}
			}
			for t, paths := range allTags {
				if len(paths) == 1 {
					fmt.Println(t)
					for _, path := range paths {
						fmt.Printf("  - content/%s\n", path)
					}
				}
			}

			fmt.Printf("%d singleton(s) found\n", len(allTags))
			return nil
		},
	}
	return cmd
}

func generateTagsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "tags",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
	cmd.AddCommand(generateSingletonTagsCmd())
	return cmd
}

func init() {
	rootCmd.AddCommand(generateTagsCmd())
}
