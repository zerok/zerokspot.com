---
date: '2016-04-25T18:44:15+02:00'
language: en
tags:
- golang
title: GOPATH, Vendoring, oh my...
---

If you're using `vendor` directories (which were introduced in Go 1.5 for
vendoring project-specific dependencies) be aware that your project still has to
be in your `$GOPATH` for that vendor-directory to be seen by `go build`.

Especially with smaller projects I often just start them inside my generic
projects folder where I put everything code-related no matter the
language. Today a co-worker wanted to try out some changes I had my and noticed
that `glide install` wasn't installing everything that was needed by the
project. Turns out I had forgotten to `glide get` one of the dependencies but
everything kept building for me simply because I also had this library installed
globally within the `$GOPATH`.

But as it happened, just adding that dependency didn't solve his building issues
so we started to play around a bit. I removed the project's dependencies also
from my system and right away ran into the same difficulties my co-worker was
reporting after fixing the glide.yaml file:

```
[INFO] Running go build github.com/boltdb/bolt
[WARN] Failed to run 'go install' for github.com/boltdb/bolt: go install: no install location for directory /Users/zerok/work/code/project/subproject/vendor/github.com/boltdb/bolt outside GOPATH
    For more details see: go help gopath
```

After reading the error message and already suspecting something like that, I
found the confirmation on the
[official mailing list](https://groups.google.com/forum/?fromgroups#!topic/golang-nuts/4K9ZVVtHJSM). Turns
out that the vendor directory only works if the parent project is already living
in the `$GOPATH`. Bummer. This might make working in a mixed-language mono-repo
a bit more complicated. At least I can still simply symlink that project into my
`$GOPATH` and everything will be fine, though 😉

```
cd $GOPATH/src/company-reposerver.com/project/
ln -s ~/work/code/project/subproject $PWD/subproject
cd subproject && glide install
```

Definitely not ideal but good enough for now 😊 If you know or/and are using
other approaches that work nicely in a similar environment (monorepos with mixed
languages), please let me know!
