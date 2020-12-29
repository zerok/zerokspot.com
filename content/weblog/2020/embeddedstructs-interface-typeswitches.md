---
title: Embedded structs, interfaces, type-switches, oh my
date: "2020-12-29T16:23:30+01:00"
tags:
- golang
- retryablehttp
- development
- til
---

Some time ago I ran into a weird issue related to interfaces In Go.  [hashicorp/go-retryablehttp](https://github.com/hashicorp/go-retryablehttp) is in use in various of our projects for more resilient HTTP API calls by handling retries, backoffs etc.. It also allows you to inject a logger so you can see what is happening below the surface. For this, you can implement one of two interfaces:

	// Logger interface allows to use other loggers than
	// standard log.Logger.
	type Logger interface {
		Printf(string, ...interface{})
	}
	
	// LeveledLogger is an interface that can be implemented by any logger or a
	// logger wrapper to provide leveled logging. The methods accept a message
	// string and a variadic number of key-value pairs. For log.Printf style
	// formatting where message string contains a format specifier, use Logger
	// interface.
	type LeveledLogger interface {
		Error(msg string, keysAndValues ...interface{})
		Info(msg string, keysAndValues ...interface{})
		Debug(msg string, keysAndValues ...interface{})
		Warn(msg string, keysAndValues ...interface{})
	}

Since we wanted to have support for different log levels, we implemented the LeveledLogger interface in a custom DefaultLeveledLogger type wrapping a [logrus.Logger](https://github.com/sirupsen/logrus):

	type DefaultLeveledLogger struct {
	  *logrus.Logger
	}
	
	func (l *DefaultLeveledLogger) Error(msg string, kvs ...interface{}) {
	}
	
	// ...

When we then injected such a DefaultLeveledLogger instance into the library, we were quite surprised as retryablehttp did only treat it as Logger instead of LevelLogger, producing simple output strings instead of taking full advantage of the log levels available within Logrus.

Turns out, we had run into two issues with our implementation here: One related to how we implemented the interface and one related to how retryablehttp checks what interface is implemented.

## Accidentally implementing an interface

The first issue was that we accidentally implemented the Logger interface too by just embedding the `*logrus.Logger` instead of wrapping and hiding it. If you’ve never worked with embedding structs before, please take a look at the [Embedding chapter of Effective Go](https://golang.org/doc/effective_go.html#embedding).

For us, embedding had one unexpected side-effect: `*logrus.Logger` also provides a `Printf` method. By just embedding that into our own DefaultLeveledLogger struct we exposed that method  and all of a sudden we also implemented that other interface:

	p := logrus.New()
	l := &DefaultLeveledLogger{p}
	l.Printf("hello")
	

This means that the logger will still be used but Printf always uses the info-loglevel no matter what kind of data should be logged.


## Type-switches and order

That alone wouldn’t have been an issue if retryablehttp would prioritise LeveledLogger over Logger. We had previously done the same in another project but didn’t run into that issue. Embedding was still happening but it didn’t bite us.

So what changed? In the project were everything worked we were at retryablehttp v0.6.8 while in the project were it failed we were still using v0.6.4. 

In v0.6.4, retryablehttp does the following check to see what kind of logging API should be used:

	switch v := logger.(type) {
	case Logger:
		v.Printf("[DEBUG] %s %s", req.Method, req.URL)
	case LeveledLogger:
		v.Debug("performing request", "method", req.Method, "url", req.URL)
	}

Then, with [cf855b1d14d561f66108aa47dc4b92493d8fcb37](https://github.com/hashicorp/go-retryablehttp/commit/cf855b1d14d561f66108aa47dc4b92493d8fcb37), the order was inverted:

	switch v := logger.(type) {
	case LeveledLogger:
		v.Debug("performing request", "method", req.Method, "url", req.URL)
	case Logger:
		v.Printf("[DEBUG] %s %s", req.Method, req.URL)
	}

This had the effect, that if something implements both interfaces, LeveledLogger is now detected first and therefore used. `switch` will just use the first branch that works. No magic there.

## Conclusion

So what did we learn from this? Don’t embed because you want to implement an interface for a struct. Wrap it completely and only expose the methods that you absolutely need for the task at hand!
