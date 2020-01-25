---
title: "Some TDD tools for Go developers"
date: "2020-01-25T21:15:00+01:00"
tags:
- golang
- tdd
---

So, week two of doing TDD and I'm still absolutely loving it. I've
also rediscovered two tools/libraries I had been using before but that
make much more sense to me now that I iterate on tests more
frequently.

## GoConvey

[GoConvey][g] was one of the first tools I used to run tests in Go
back when I learnt the language all these years ago. The project has
two main parts:

1. A testing framework which is supposed to give your tests a slightly
   nicer structure using a DSL
2. A test runner that continuously watches your source files, runs the
   associated tests, and finally present the output to you in a web
   interface
   
For my purposes, the second aspect is highly relevant. Whenever I now
want to change something inside a package, I open a goconvey instance
in there. It will start watching the files in that package for changes
, run the relevant tests, and let me know if I'm done with my change
or not.

Let's say, I'm working on a new helper for [tpl][] inside the
`internal/world` package. Then I change into the folder of that
package and run `goconvey`:

```
$ goconvey
...
2020/01/25 20:34:57 goconvey.go:105: Launching browser on 127.0.0.1:8080
2020/01/25 20:34:57 goconvey.go:178: Serving HTTP at: http://127.0.0.1:8080
```

This will open a browser window and show me all the tests that have
just been run:

<figure>
<img src="/media/2020/goconvey.png">
<figcaption>Main status page of GoConvey listing all the run tests</figcaption>
</figure>

GoConvey also supports sending notifications after every testrun. This
means, I don't have to keep the browser open on a secondary display
while coding on the primary one. Instead, I can just leave the browser
somewhere in the background and have all my screen-estate for the
tests and code to iterate on. This is especially handy while coding on
my 13" laptop while at home or on the go.

Sadly, notifications are disabled by default but can easily be enabled
through the "bell" sign in GoConvey's UI as shown below:

<figure>
<img src="/media/2020/goconvey-notifications.png">
<figcaption>Notification preferences</figcaption>
</figure>

## dockertest

[Dockertest][dt] is a little library I had already written about last
year: It allow you to easily start up a Docker container, for instance
of the database system you're using in your application, and access
its metadata like exposed ports etc. This little library has made my
integration tests so much more readable, it's hard to describe :-)

But as I said, I've already written about it before. If you want my 2c
about it, go [here][dtz].

## Go's crypto package test keys

The last one is just a little shout-out to Go's awesome standard
library: I recently also improved some JWT code but ran into the
situation where I didn't felt like checking in some previously valid
tokens (as they had already expired) as test objects. What I did
instead, was to create my own little RSA keypair using Go's
[crypto/rsa][r] package and generate tokens inside my test cases using
that:

```
import "crypto/rsa"
import "crypto/rand"

privateKey, _ := rsa.GenerateKey(rand.Reader)
publicKey := privateKey.Public()
```

These two lines of code prevented me from having to check in some
pre-generated RSA keys which made the whole test suite much more
pleasant to read. It also made testing the underlying code much more
flexible because I could generate tokens with whatever properties I
needed instead of having to rely on pre-generated ones.

[tpl]: https://github.com/zerok/tpl/
[dtz]: https://zerokspot.com/weblog/2019/05/11/integrating-docker-into-go-tests/
[dt]: https://github.com/ory/dockertest
[g]: http://goconvey.co/
[r]: https://golang.org/pkg/crypto/rsa/
