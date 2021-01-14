---
title: Testable HTTP API libraries with http.RoundTripper
date: "2021-01-14T20:12:58+01:00"
tags:
- golang
- development
- testing
- http
- 100daystooffload
incoming:
- url: https://chaos.social/@zerok/105555682864766893
---

Over the course of the last couple of days I’ve had a lot to do with client libraries for various online services. For this post I’m just going to use the [official Go library for Slack](https://github.com/slack-go/slack) but there are many others where the presented pattern should help integrating them into test setups.

## Case 1: What data is sent?

In one situation I wanted to see what kind of data would be sent to the Slack API (for the `chat.postMessage` endpoint) without actually sending data out. We had wrapped the Slack API a little bit in order to make things like building message blocks easier but I wanted to have a way to evaluate the final HTTP request body in order to compare it with what was expected according to the specs.

Sadly, I couldn’t just wrap the [slack.Client struct](https://pkg.go.dev/github.com/slack-go/slack#Client) into an interface and fake a [Client.PostMessage](https://pkg.go.dev/github.com/slack-go/slack#Client.PostMessage) method because you feed it slack.MsgOption objects that act as configurators for a private object. So there’s no easy way to just fake the message creation with just the input of PostMessage.

With the easy way off the list, there was one last idea I had: Can I somehow inject a custom HTTP Client into the Slack client and just grab the raw HTTP request data? Yes, that’s possible and this is also the point where the `http.RoundTripper` interface comes in:

	package main
	
	import (
		"net/http"
		"testing"
	
		"github.com/slack-go/slack"
		"github.com/stretchr/testify/require"
	)
	
	type LoggingRoundTripper struct {
		Requests []*http.Request
	}
	
	func (rt *LoggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
		rt.Requests = append(rt.Requests, r)
		return nil, nil
	}
	
	func TestSlackAPIPostMessage(t *testing.T) {
		lrt := &LoggingRoundTripper{
			Requests: make([]*http.Request, 0, 10),
		}
		hc := &http.Client{
			Transport: lrt,
		}
		c := slack.New("no-token", slack.OptionHTTPClient(hc))
		c.PostMessage("somechannel", slack.MsgOptionText("hello world", false))
		require.Len(t, lrt.Requests, 1)
	}
	

In this example I’ve implemented a simple RoundTripper that just adds the incoming requests to a slice which can then validated again.

# Case 2: Fake API responses

In another situation I wanted to fake the response I got from the server. What would postMessage do, for instance, if it got a 500 status code from the server? If you want to have different behaviour for different URLs, you can, for instance, combine RoundTripper with a Handler (also provided by the net/http package). This makes it possible to just use a request muxer and keep you sane 😅

	package main
	
	import (
		"fmt"
		"net/http"
		"net/http/httptest"
		"testing"
	
		"github.com/go-chi/chi"
		"github.com/slack-go/slack"
		"github.com/stretchr/testify/require"
	)
	
	type FakeServerRoundTripper struct {
		Handler http.Handler
	}
	
	func (rt *FakeServerRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
		if rt.Handler == nil {
			return nil, fmt.Errorf("no handler set")
		}
		w := httptest.NewRecorder()
		rt.Handler.ServeHTTP(w, r)
		return w.Result(), nil
	}
	
	func TestSlackAPIPostMessage(t *testing.T) {
		mux := chi.NewRouter()
		mux.Post("/api/chat.postMessage", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		})
		fsrt := &FakeServerRoundTripper{Handler: mux}
		hc := &http.Client{
			Transport: fsrt,
		}
		c := slack.New("no-token", slack.OptionHTTPClient(hc))
		_, _, err := c.PostMessage("somechannel", slack.MsgOptionText("hello world", false))
		require.Error(t, err)
		require.Contains(t, err.Error(), "500")
	}
	


## Conclusion

While I would have preferred to use fake interface implementations to test some of the interactions with the Slack API and other libraries, it’s sometimes simply not possible. Luckily, most HTTP-based libraries allow you to at least inject your own `http.Client` and so you can at least fake your way around with a custom RoundTripper. It may not be nice, but it’s definitely powerful!

Now I just have to restrain myself so that I don’t abuse this pattern too much 😆
