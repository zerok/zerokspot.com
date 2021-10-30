---
title: "go-logr: A logger abstraction interface"
date: "2021-10-30T09:18:00+02:00"
tags:
- golang
- libraries
---

During [GoGraz 2021.10](https://gograz.org/meetup/2021-10-18/), Stephan mentioned as nice logging library for Go that I had previously missed: [go-logr/logr](https://github.com/go-logr/logr). This is basically an abstraction library that provides a simple interface for logging under which you can then put an actual logger that then writes the statements to whatever system you like to have. logr tries to solve the use-case that libraries often also want to generate log output but cannot really anticipate what logging system the application around them is using. 

An example for such a scenario is [hashicorp/go-retryablehttp](https://github.com/hashicorp/go-retryablehttp) which wants to log every retry and HTTP request. Here you'd either implement a [`retryable.Logger`](https://pkg.go.dev/github.com/hashicorp/go-retryablehttp?utm_source=godoc#Logger) or [`retryable.LevelLogger`](https://pkg.go.dev/github.com/hashicorp/go-retryablehttp?utm_source=godoc#LeveledLogger) interface to work with your favorite logging framework and then attach it to the `retryablehttp.Client`.

logr now offers a general abstraction that can be used everywhere so that library authors don't have to come up with their own interfaces for loggers anymore.
Inside a library you'd basically just either take a `logr.Logger` instance directly or extract one from a given context:

```go
package main

import (
	"context"

	"github.com/go-logr/logr"
)

func helper(ctx context.Context) {
	logger := logr.FromContextOrDiscard(ctx)
	logger.Info("hello from helper")
}
```

How would this work, for instance, with zerolog underneath? This would be using go-logr/logr in combination with [go-logr/zerologr](https://github.com/go-logr/zerologr):

```go
package main

import (
	"context"
	"os"

	"github.com/go-logr/logr"
	"github.com/go-logr/zerologr"
	"github.com/rs/zerolog"
)

func main() {
	ctx := context.Background()
	zl := zerolog.New(os.Stderr).Level(zerolog.DebugLevel)
	logger := zerologr.New(&zl)
	ctx = logr.NewContext(ctx, logger.WithName("helper"))
	helper(ctx)
}
```

## Basic logging

Once you have that, you use just the methods that are available through `logr.Logger`. Want to log something with some additional fields?

```go
logger.Info("This is the message", "field1", 123, "field2", false)
```

As you can see here, additional fields are simply added to the Info call as key-value pairs. Underneath this will cause some reflection to take place depending on what underlying library you're using. Additionally, there's also an `Error` function that has basically the same signature as `Info` but takes as first argument an `error` object.

## Verbosity levels

As already hinted at, the API of a `logr.Logger` only exposes two methods for actually creating log messages:

- `Info(Info(msg [string](https://pkg.go.dev/builtin#string), keysAndValues ...interface{})`
- `Error(err [error](https://pkg.go.dev/builtin#error), msg [string](https://pkg.go.dev/builtin#string), keysAndValues ...interface{})`

Compared to zerolog, these two have slightly different semantics. Both of them somehow combine the additional fields of the log entry with its level. In zerolog, for instance,  you can have an info-entry that still has an error attached to it. Here, that's not directly possible. That being said, this whole mapping is done within the `zerologr` library, so you could create your own adapter that does away with this behaviour if you don't want it there.

In general, log levels are handled in a more customizable way within logr. Those "verbosity levels" are just integers. The higher the number the more detailed the message gets. In a sense, you end up with something like this:

| Verbosity | Meaning |
|-|-|
| 0 | Messages that you basically want to always see outside of errors |
| 1| Debug level messages |
| 2 | Trace level messages |
| ... | The higher the less you want to see them |

The verbosity level is set per logger using the `V(int)` method:

```
logger.V(1).Info("hello") // lvl 1
logger.V(1).V(1).Info("hello") // lvl 2
```

It's then the job of the underlying sink to only process those log entries that have a level equal or lower to the one they were set up to process. For zerolog that has been configured to do debug logging, the lvl 2 statement above would no longer be rendered.

```go
package main

import (
	"fmt"
	"os"

	"github.com/go-logr/zerologr"
	"github.com/rs/zerolog"
)

func main() {
	zl := zerolog.New(os.Stderr).Level(zerolog.InfoLevel)
	logger := zerologr.New(&zl)

	logger.Info("info from logr") // Logs to level 0 by default

	// Error logs are always visible
	logger.Error(fmt.Errorf("new error"), "error from logr")

	logger.V(0).Info("level 0") // Visible on InfoLevel
	logger.V(1).Info("level 1") // Visible on DebugLevel
	logger.V(2).Info("level 2") // Visible on TraceLevel
	logger.V(3).Info("level 3") // With zerolog never visible
	
	// Output:
	// {"v":0,"message":"info from logr"}
	// {"error":"new error","message":"error from logr"}
	// {"v":0,"message":"level 0"}
}
```

## Prefill values

As you've already seen in the first example when putting a logger into a context, you can also attach fields directly to a logger so that they are part of every log statement afterwards. This is done using the `WithName` and `WithValues` methods:

```go
package main

import (
	"os"

	"github.com/go-logr/zerologr"
	"github.com/rs/zerolog"
)

func main() {
	zl := zerolog.New(os.Stderr).Level(zerolog.DebugLevel)
	logger := zerologr.New(&zl).WithName("somename").WithValues("key1", "val1")
	logger.Info("hello")
	
	// Output:
	// {"key1":"val1","v":0,"logger":"somename","message":"hello"}
}
```

## Conclusion

After only playing around with it for a very limited amount of time, I really like the approach that logr is taking. There are already adapters avaible for every framework I might be using (i.e. zerolog and zap) and it looks to be well-maintained and stable. One thing I'm not yet sure about is how I'll work with the verbosity levels here. zerolog only has a very limited number of available levels so moving beyond V(2) won't make all that much sense there for now. 

The way `V(int)` call add up (e.g. `V(1).V(1)` will result in a verbosity of 2) makes it also very simple to silence chatty libraries. I'm not yet sure if the reverse is also possible. If a library author, of instance, uses completely different log levels I don't know if I could map those from the outside to the 0, 1, and 2 of zerolog. I guess I'd have to provide a custom LogSink implementation (similar to `zerlogr`) that does the mapping for me there.
