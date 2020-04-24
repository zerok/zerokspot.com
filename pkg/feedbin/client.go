package feedbin

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const defaultBaseURL string = "https://api.feedbin.com"

type Client struct {
	BaseURL      string
	HTTPClient   http.Client
	AuthUser     string
	AuthPassword string
}

type ClientConfigurator func(*Client)

func New(configs ...ClientConfigurator) *Client {
	c := &Client{}
	c.BaseURL = defaultBaseURL
	c.HTTPClient = http.Client{}
	for _, cfg := range configs {
		cfg(c)
	}
	return c
}

func (c *Client) buildURL(path string) string {
	return fmt.Sprintf("%s%s", c.BaseURL, path)
}

func (c *Client) GetSubscriptions(ctx context.Context) ([]Subscription, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.buildURL("/v2/subscriptions.json"), nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.AuthUser, c.AuthPassword)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	subscriptions := make([]Subscription, 0, 10)
	if err := json.NewDecoder(resp.Body).Decode(&subscriptions); err != nil {
		return nil, err
	}
	return subscriptions, nil
}

func (c *Client) GetTaggings(ctx context.Context) ([]Tagging, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.buildURL("/v2/taggings.json"), nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.AuthUser, c.AuthPassword)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	taggings := make([]Tagging, 0, 10)
	if err := json.NewDecoder(resp.Body).Decode(&taggings); err != nil {
		return nil, err
	}
	return taggings, nil
}
