package feedbin

type Subscription struct {
	ID      int64  `json:"id"`
	FeedID  int64  `json:"feed_id"`
	Title   string `json:"title"`
	FeedURL string `json:"feed_url"`
	SiteURL string `json:"site_url"`
}
