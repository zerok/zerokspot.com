---
title: Structured logging accepted to the Go stdlib
date: "2023-03-17T20:56:27+01:00"
tags:
- golang
---
Looks like the [log/slog proposal](https://go.googlesource.com/proposal/+/master/design/56345-structured-logging.md) for adding structured logging to Go’s standard library [has been accepted](https://github.com/golang/go/issues/56345#issuecomment-1470506816)! This also includes support for levels and more complex backend options by allowing developers to implement their own log-record handlers. Something like this has really been missing in the current `log` package where all you get is an `io.Writer` and `Print` method.

From what I’ve seen of the proposal so far it’s a really nice one and might finally make discussions around “which logging library to use” shorter. There was also a long discussion in [the GitHub issue](https://github.com/golang/go/issues/56345) about if something like `zerolog.Ctx(context.Context) *zerolog.Logger` should also be provided which would offer a standardised way to propagate a logger through a context. At this point, [this won’t happen](https://github.com/golang/go/issues/56345#issuecomment-1469945135). Personally, I would have liked to see something like this but perhaps it might come in the future 🙂

I think I will use one of my pet-projects to give this a try to get a proper opinion about this new addition to the standard library and if I can really replace zero log with it in the long run.
