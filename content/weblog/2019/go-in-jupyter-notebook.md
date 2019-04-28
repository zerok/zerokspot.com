---
title: "Using Go in Jupyter Notebook"
date: "2019-04-28T11:17:14+02:00"
tags:
- golang
- jupyter
---

From time to time I want to use Jupyter with Go code elements instead
of Python. Luckily, this is possible thanks to the Go kernel provided
by the [gophernotes][] project. Under the hood this uses the plugin
system [introduced with Go 1.8][p18] and then later extended in Go
1.10. This has the consequence that some functionality is only
available on Linux and macOS but not on Windows.

In order to use gophernotes as kernel for Jupyter-Notebook you have to
run the following steps (asuming that you're on Linux or macOS):

```
# We want to install the binary so we need to deactivate Go Modules
# for now:
$ export GO111MODULE=off
$ go get -u github.com/gopherdata/gophernotes

# For macOS
$ mkdir -p $HOME/Library/Jupyter/kernels/gophernotes
$ cp $GOPATH/src/github.com/gopherdata/gophernotes/kernel/* \
    $HOME/Library/Jupyter/kernels/gophernotes

# For Linux
$ mkdir -p $HOME/.local/share/jupyter/kernels/gophernotes
$ cp $GOPATH/src/github.com/gopherdata/gophernotes/kernel/* \
    $HOME/.local/share/jupyter/kernels/gophernotes
```

Once that's done, start a new jupyter-notebook instance. In the "New"
menu you should now see a "Go" entry which will use the gophernotes
kernel.


[jupyter]: https://jupyter.org/
[gophernotes]: https://github.com/gopherdata/gophernotes
[p18]: https://golang.org/doc/go1.8#plugin
