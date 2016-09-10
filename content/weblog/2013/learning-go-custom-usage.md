---
date: '2013-11-18T21:52:03+01:00'
language: en
tags:
- golang
- flag
- cli
title: 'Learning Go: Custom Usage'
---


I’ve wanted to learn [Go][golang] for the better part of the last two years now and last weekend I finally took the time to actually write my very first lines in that language. Nothing special, mind you, but enough to get the feel for it and work on something that actually solves a real problem of mine. As I learn more I want to share some little tricks that I picked up and that were not obvious to me by just looking at the docs.

The first one is how to change the usage information provided by the [flag][] package. For everything in there there seems to exist some function to manipulate the internal state, but not for the usage.

To change that, you have to write your own Usage method:

```
import (
    "flag"
    "fmt"
)

flag.Usage = func() {
    fmt.Printf("Usage: yourtool [options] param>\n\n")
    flag.PrintDefaults()
}
```

When you now launch your program with `--help` you will see your usage message instead of the “ugly” default one :-)

As a little side-note, the first thing I did to that little CLI tool was to replace the flag package and replace it with [pflag][] for one simple feature: it lets me set optional flags *after* the positional arguments if I want to.

[flag]: http://golang.org/pkg/flag/
[pflag]: https://github.com/ogier/pflag
[golang]: http://golang.org/
