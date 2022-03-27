---
title: Twitter’s URL shortener and HTTP/2
date: "2022-03-27T17:51:26+02:00"
tags:
- golang
---

Thanks to a [unittest failing on webmentiond](https://github.com/zerok/webmentiond/runs/5709325508?check_suite_focus=true) today I stumbled upon a weird issue: The default HTTP client in Go is running into a 404 error when trying to resolve a t.co URL. Weirdly enough the same request went through with cURL even using the same HTTP headers. Eventually I tried it with enabling/disable HTTP/2 support and lo-and-behold, all of a sudden I got 200 status codes again.

For some reason and out of the blue t.co changed its behaviour regarding HTTP/2 requests:

```go
package main

import (
	"crypto/tls"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTwitterH2(t *testing.T) {
	h1client := http.Client{}
	// Using Transport.TLSNextProto you can disable HTTP/2 support:
	tr := http.Transport{
		TLSNextProto: make(map[string]func(string, *tls.Conn) http.RoundTripper),
	}
	h1client.Transport = &tr
	h2client := http.Client{}

	req, _ := http.NewRequest(http.MethodGet, "https://t.co/JqumM1uaVE", nil)

	t.Run("h1", func(t *testing.T) {
		resp, err := h1client.Do(req)
		require.NoError(t, err)
		require.Equal(t, 200, resp.StatusCode)
	})

	t.Run("h2", func(t *testing.T) {
		resp, err := h2client.Do(req)
		require.NoError(t, err)
		require.Equal(t, 404, resp.StatusCode)
	})
}
```

In the h2 case I get a 404 error while in the h1 one I get 200. I haven't yet found out what's the difference here compared to running `curl --http2` , though.

If anyone has some details about what’s going on there, please let me know 🙂
