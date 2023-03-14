---
title: Mage files in a directory
date: "2023-03-14T20:31:15+01:00"
tags:
- golang
- development
- til
---

Over the last couple of weeks and months I’ve grown really fond of [Mage](https://magefile.org), a build-tool with some aspects of GNU Make but written in Go. The `Magefile.go` where you define your build-targets is also just Go so there’s no extra language to learn. Just create such a file, write some public Go functions into it and the `mage` binary will pick them up as build targets.

In order to prevent any package clashes between your `Magefile.go`’s main-package and others you usually put your Magefile behind a build-tag:

	//go:build mage
	package main
	
	func Hello() error {
	  return nil
	}

Unfortunately, using build-tags can easily confuse tools like the auto-completion system of editors (e.g. gopls) which makes them sometimes annoying to work with esp. for larger files.

Luckily, since [last March (v1.13)](https://magefile.org/blog/2022/03/release-v1.13.0/) Mage now also supports targets defined in a `magefiles` directory, basically preventing any clashes with other `main` packages! Having a folder folder available has also made it more attractive to split targets by topic/namespace! With some of the projects I’m currently working on, I’ve now mostly settled on the following structure:

	/mage.go # Simple bootstrap file
	/magefiles
	    /docs.go  # For build targets around building the documentation
	    /ci.go    # Usually includes some Dagger ;-)
	    /build.go # Classic build targets

The `mage.go`[boostrap file](https://raw.githubusercontent.com/gograz/gograz-meetup/master/bootstrap.go) is mostly there because I simply don’t like global binaries and so I can just run `go run mage.go TARGET` 🙂
