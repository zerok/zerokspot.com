package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var fontAwesomeRoot string
var configFile string

type iconData struct {
	ViewBox []int  `json:"viewBox"`
	Data    string `json:"data"`
}

type iconsOutput struct {
	Brands map[string]map[string]iconData
}

type iconConfig struct {
	Group string `json:"group"`
	Icon  string `json:"icon"`
}

type iconsConfig struct {
	Icons []iconConfig `json:"icons"`
}

var buildIconsCmd = &cobra.Command{
	Use:   "build-icons",
	Short: "Build an icons.js file",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := logger.WithContext(cmd.Context())
		ctx = findParentTrace(ctx)
		ctx, span := tracer.Start(ctx, "build-icons")
		defer span.End()
		var config iconsConfig
		fp, err := os.Open(configFile)
		if err != nil {
			return err
		}
		if err := yaml.NewDecoder(fp).Decode(&config); err != nil {
			fp.Close()
			return err
		}
		fp.Close()
		output := make(map[string]map[string]iconData)
		for _, cfg := range config.Icons {
			group, ok := output[cfg.Group]
			if !ok {
				group = make(map[string]iconData)
			}
			d, err := buildIconData(filepath.Join(fontAwesomeRoot, "svgs", cfg.Group, cfg.Icon+".svg"))
			if err != nil {
				return err
			}
			group[cfg.Icon] = *d
			output[cfg.Group] = group
		}
		var out bytes.Buffer
		if err := json.NewEncoder(&out).Encode(&output); err != nil {
			return err
		}
		outFP, err := os.OpenFile(outputPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
		if err != nil {
			return err
		}
		defer outFP.Close()
		fmt.Fprintf(outFP, "window.icons = %s;\n", out.String())
		return nil
	},
}

func buildIconData(path string) (*iconData, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	output := iconData{}
	viewBoxRe := regexp.MustCompile("viewBox=\"([^\"]+)\"")
	dRe := regexp.MustCompile("d=\"([^\"]+)\"")
	output.Data = string(dRe.FindSubmatch(data)[1])
	viewBox := string(viewBoxRe.FindSubmatch(data)[1])
	output.ViewBox = make([]int, 0, 4)
	for _, e := range strings.Split(viewBox, " ") {
		i, err := strconv.Atoi(e)
		if err != nil {
			return nil, err
		}
		output.ViewBox = append(output.ViewBox, i)
	}
	return &output, nil
}

func init() {
	buildIconsCmd.Flags().StringVar(&fontAwesomeRoot, "fontawesome-root", "static/fontawesome-pro-5.13.0-web", "Output path")
	buildIconsCmd.Flags().StringVar(&outputPath, "output", "static/js/icons.js", "Output path")
	buildIconsCmd.Flags().StringVar(&configFile, "config", "icons.yaml", "Path to an icons.yaml")
	rootCmd.AddCommand(buildIconsCmd)
}
