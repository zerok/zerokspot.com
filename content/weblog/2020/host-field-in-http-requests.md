---
title: "Setting Host header in http.Requests"
tags:
- golang
- programming
date: "2020-02-16T17:06:40+0100"
---

A couple of days ago I wanted to test some HTTP endpoint with Go's [net/http](https://golang.org/pkg/net/http/) library. These endpoints were reachable through a virtual host configuration so the test requests had to explicitly set the Host header. 

Turns out, you cannot simply set that like any other request header. Instead, you have to set the Host property of the request object itself:

```
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHostHeader(t *testing.T) {
	hostHeaderValue := ""
	hostFieldValue := ""
	handler := func(w http.ResponseWriter, r *http.Request) {
		hostHeaderValue = r.Header.Get("Host")
		hostFieldValue = r.Host
	}
	srv := httptest.NewServer(http.HandlerFunc(handler))
	defer srv.Close()

	client := srv.Client()

	// Setting the Host header the normal way, won't do anything:
	req, _ := http.NewRequest(http.MethodGet, srv.URL+"/", nil)
	req.Header.Set("Host", "testhost")
	// The header field *has* been set!
	require.Equal(t, req.Header.Get("Host"), "testhost")
	client.Do(req)
	require.NotEqual(t, "testhost", hostFieldValue)
	// In fact, the Host header will be the host component of the
	// srv.URL:
	require.Contains(t, srv.URL, hostFieldValue)
	require.Empty(t, hostHeaderValue)

	// If you use the Host field, it will reach the server:
	req, _ = http.NewRequest(http.MethodGet, srv.URL+"/", nil)
	req.Host = "testhost"
	client.Do(req)
	require.Equal(t, "testhost", hostFieldValue)
	require.Empty(t, hostHeaderValue)
}
```

In this example I've also depicted another special handling of the Host header: Inside a http.Handler the value of the Host-header is promoted to the http.Request.Host field and then removed from the http.Request.Header map.

Note: All of this [**is** documented](https://golang.org/pkg/net/http/#Request). It just surprised me a bit 🙂
