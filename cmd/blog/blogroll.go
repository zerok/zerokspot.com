package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/spf13/cobra"
	"gitlab.com/zerok/zerokspot.com/pkg/feedbin"
)

var outputPath string
var tagName string

type BlogrollItem struct {
	Title   string `json:"title"`
	FeedURL string `json:"feed_url"`
	SiteURL string `json:"site_url"`
}

type BlogrollItemsByTitle []BlogrollItem

func (items BlogrollItemsByTitle) Len() int {
	return len(items)
}

func (items BlogrollItemsByTitle) Swap(i, j int) {
	oldi := items[i]
	items[i] = items[j]
	items[j] = oldi
}

func (items BlogrollItemsByTitle) Less(i, j int) bool {
	return items[i].Title < items[j].Title
}

var blogrollCmd = &cobra.Command{
	Use:   "blogroll",
	Short: "Generate blogroll.json",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		user := os.Getenv("FEEDBIN_USER")
		password := os.Getenv("FEEDBIN_PASSWORD")

		if user == "" || password == "" {
			return fmt.Errorf("please specify FEEDBIN_USER and FEEDBIN_PASSWORD")
		}

		fb := feedbin.New(func(c *feedbin.Client) {
			c.AuthUser = user
			c.AuthPassword = password
		})

		subscriptions, err := fb.GetSubscriptions(ctx)
		if err != nil {
			return err
		}
		taggings, err := fb.GetTaggings(ctx)
		if err != nil {
			return err
		}
		relevantFeeds := make(map[int64]struct{})
		for _, t := range taggings {
			if t.Name == tagName {
				relevantFeeds[t.FeedID] = struct{}{}
			}
		}

		result := make([]BlogrollItem, 0, 10)
		for _, s := range subscriptions {
			if _, ok := relevantFeeds[s.FeedID]; !ok {
				continue
			}
			result = append(result, BlogrollItem{
				Title:   s.Title,
				FeedURL: s.FeedURL,
				SiteURL: s.SiteURL,
			})
		}

		sort.Sort(BlogrollItemsByTitle(result))

		fp, err := os.OpenFile(outputPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
		if err != nil {
			return err
		}
		defer fp.Close()
		if err := json.NewEncoder(fp).Encode(result); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(blogrollCmd)
	blogrollCmd.Flags().StringVar(&outputPath, "output", "blogroll.json", "Output path")
	blogrollCmd.Flags().StringVar(&tagName, "tag", "Personal", "Tag to use")
}
