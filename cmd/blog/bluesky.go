package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/bluesky-social/indigo/api/bsky"
	"github.com/bluesky-social/indigo/util"
	"github.com/bluesky-social/indigo/xrpc"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var importBlueskyIncomingCmd = &cobra.Command{
	Use: "import-bluesky-links",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		pattern := regexp.MustCompile(`https://zerokspot.com/weblog/([^"]+)`)
		client := xrpc.Client{Host: "https://public.api.bsky.app"}
		cursor := ""
		for range 2 {
			feed, err := bsky.FeedGetAuthorFeed(ctx, &client, "zerokspot.com", cursor, "", false, 100)
			if err != nil {
				return err
			}
			cursor = *feed.Cursor

			// Find relevant status messages
			for _, bskyPost := range feed.Feed {
				isRelevant := false
				post := bskyPost.Post.Record.Val.(*bsky.FeedPost)
				for _, facet := range post.Facets {
					for _, feature := range facet.Features {
						if feature.RichtextFacet_Tag != nil && feature.RichtextFacet_Tag.Tag == "blogged" {
							isRelevant = true
						}
					}
				}
				if isRelevant {
					// Extract the zerokspot URL that should be updated:
					if post.Embed == nil || post.Embed.EmbedExternal == nil || post.Embed.EmbedExternal.External == nil {
						continue
					}
					blogURL := post.Embed.EmbedExternal.External.Uri
					match := pattern.FindStringSubmatch(blogURL)
					if len(match) < 2 {
						continue
					}
					parsedURI, err := util.ParseAtUri(bskyPost.Post.Uri)
					if err != nil {
						return err
					}
					elems := strings.Split(match[1], "/")
					postFile := "content/weblog/" + elems[0] + "/" + elems[3] + ".md"
					if err := addBlueskyLink(ctx, postFile, fmt.Sprintf("https://bsky.app/profile/%s/post/%s", "zerokspot.com", parsedURI.Rkey)); err != nil {
						return err
					}
				}
			}
		}

		return nil
	},
}

func addBlueskyLink(_ context.Context, postFile string, postURL string) error {
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
		if incoming["url"] == postURL {
			log.Printf("%s present in %s", postURL, postFile)
			return nil
		}
	}

	incomings = append(incomings, map[any]any{
		"url": postURL,
	})
	frontmatter["incoming"] = incomings

	// Inject the new frontmatter:
	var fmData bytes.Buffer
	if err := yaml.NewEncoder(&fmData).Encode(frontmatter); err != nil {
		return err
	}

	log.Printf("%s not linked to from %s\n", postURL, postFile)
	return injectFrontMatter(postFile, frontmatter, frontMatterStart, frontMatterEnd)
}

func init() {
	rootCmd.AddCommand(importBlueskyIncomingCmd)
}
