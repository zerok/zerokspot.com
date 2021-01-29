---
title: "Go 1.16: signal.NotifyContext" 
date: "2021-01-29T13:25:00+01:00"
tags:
- golang
- "100daystooffload"
---

Yesterday I somehow stumbled (probably via r/golang) onto another nice feature that’s incoming with [Go 1.16](https://tip.golang.org/doc/go1.16) which I had previously missed while skimming through the release notes: [`signal.NotifyContext`](https://tip.golang.org/pkg/os/signal/#NotifyContext).

Whenever you have some kind of long-running process there might be a need to do some state-clean-up, closing connections, or just doing some general house-keeping before exiting. Such processes usually exit when they receive, for instance, an interrupt from the operating system (e.g. someone hitting Ctrl+C in the shell the process is running in).

Previously, you had to basically setup up a signal handler including channel and separate Go routine yourself if you wanted to just cancel some context, but now there is a standard library function for that:

	package main
	
	import (
		"context"
		"fmt"
		"os"
		"os/signal"
	)
	
	func main() {
		ctx, done := signal.NotifyContext(context.Background(), os.Interrupt)
		defer done()
	
		fmt.Println("Waiting for SIGINT.")
		<-ctx.Done()
		fmt.Println("We are done here.")
	}

One more thing to look forward to with 1.16 ☺️
