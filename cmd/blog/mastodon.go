package main

import (
	"bufio"
	"bytes"
	"context"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/mattn/go-mastodon"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var importMastodonIncomingCmd = &cobra.Command{
	Use: "import-mastodon-links",
	RunE: func(cmd *cobra.Command, args []string) error {
		pattern := regexp.MustCompile(`href="https://zerokspot.com/weblog/([^"]+)"`)
		client := mastodon.NewClient(&mastodon.Config{
			Server: "https://chaos.social",
		})
		ctx := cmd.Context()
		statuses, err := client.GetAccountStatuses(ctx, mastodon.ID("55379"), nil)
		if err != nil {
			return err
		}

		// Find relevant status messages
		for _, status := range statuses {
			isRelevant := false
			for _, tag := range status.Tags {
				if tag.Name == "blogged" {
					isRelevant = true
					break
				}
			}
			if isRelevant {
				// Extract the zerokspot URL that should be updated:
				matches := pattern.FindAllStringSubmatch(status.Content, -1)
				for _, match := range matches {
					elems := strings.Split(match[1], "/")
					postFile := "content/weblog/" + elems[0] + "/" + elems[3] + ".md"
					if err := addMastodonLink(ctx, postFile, status.URL); err != nil {
						return err
					}
				}

			}
		}

		return nil
	},
}

func addMastodonLink(ctx context.Context, postFile string, mastodonURL string) error {
	// Check front matter if it contains already a matching link:
	postContent, err := os.ReadFile(postFile)
	if err != nil {
		return err
	}
	frontmatter, frontMatterStart, frontMatterEnd, err := parseFrontmatter(string(postContent))
	if err != nil {
		return err
	}

	incomings := make([]map[any]any, 0, 5)
	if data, ok := frontmatter["incoming"]; ok {
		for _, d := range data.([]any) {
			incomings = append(incomings, d.(map[any]any))
		}

	}

	for _, incoming := range incomings {
		if incoming["url"] == mastodonURL {
			log.Printf("%s present in %s", mastodonURL, postFile)
			return nil
		}
	}

	incomings = append(incomings, map[any]any{
		"url": mastodonURL,
	})
	frontmatter["incoming"] = incomings

	// Inject the new frontmatter:
	var fmData bytes.Buffer
	if err := yaml.NewEncoder(&fmData).Encode(frontmatter); err != nil {
		return err
	}

	log.Printf("%s not linked to from %s\n", mastodonURL, postFile)
	return injectFrontMatter(postFile, frontmatter, frontMatterStart, frontMatterEnd)
}

type postFrontMatter map[string]interface{}

type postFrontMatterIncoming struct {
	URL string `yaml:"url"`
}

func injectFrontMatter(postFile string, frontmatter postFrontMatter, frontMatterStart int, frontMatterEnd int) error {
	fp, err := os.Open(postFile)
	if err != nil {
		return err
	}
	rd := bufio.NewScanner(fp)
	var output bytes.Buffer
	lineNum := 0
	for rd.Scan() {
		lineNum++
		line := rd.Text()
		if lineNum == frontMatterStart {
			output.WriteString("---\n")
			if err := yaml.NewEncoder(&output).Encode(frontmatter); err != nil {
				return err
			}
			output.WriteString("---\n")
		}
		if lineNum < frontMatterStart || lineNum > frontMatterEnd {
			output.WriteString(line)
			output.WriteRune('\n')
		}
	}
	return os.WriteFile(postFile, output.Bytes(), 0664)
}

func parseFrontmatter(raw string) (postFrontMatter, int, int, error) {
	rd := bufio.NewScanner(bytes.NewBufferString(raw))
	rawFrontmatter := bytes.Buffer{}
	inBody := false
	startLine := -1
	endLine := -1
	lineNum := 0
	for rd.Scan() {
		lineNum++
		line := rd.Text()
		if strings.HasPrefix(line, "---") {
			if inBody {
				endLine = lineNum
				break
			}
			startLine = lineNum
			inBody = true
			continue
		}
		if inBody {
			rawFrontmatter.WriteString(line)
			rawFrontmatter.WriteByte('\n')
		}
	}
	result := postFrontMatter{}
	dec := yaml.NewDecoder(&rawFrontmatter)
	err := dec.Decode(&result)
	return result, startLine, endLine, err
}

func init() {
	rootCmd.AddCommand(importMastodonIncomingCmd)
}
