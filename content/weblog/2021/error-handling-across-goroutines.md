---
title: Error handling across goroutines
date: "2021-05-13T11:29:45+02:00"
tags:
- til
- golang
- 100daystooffload
- concurrency
incoming:
- url: https://chaos.social/@zerok/106227216140026354
---

When having whole groups of goroutines it gets quite tedious to always write the logic around error handling. Luckily, there is the [x/sync/errgroup](https://pkg.go.dev/golang.org/x/sync/errgroup) package which allows you things like cancelling the whole group if a single one returns with an error:

	package main
	
	import (
		"context"
		"fmt"
		"log"
	
		"golang.org/x/sync/errgroup"
	)
	
	func main() {
		g, ctx := errgroup.WithContext(context.Background())
		for idx := 0; idx < 5; idx++ {
			i := idx
			g.Go(func() error {
				defer log.Printf("%d finished", i)
				if i == 3 {
					return fmt.Errorf("goroutine %d raised an error", i)
				}
				<-ctx.Done()
				return nil
			})
		}
		if err := g.Wait(); err != nil {
			log.Printf("Err: %s\n", err.Error())
		}
	}

In the example above I create 5 goroutines attached to an errgroup. The 4th one will immediately return with an error which will cancel the context and make all the other goroutines return. `Wait()` waits for all of that and then return the error that triggered it all.

	2021/05/13 11:22:07 3 finished
	2021/05/13 11:22:07 0 finished
	2021/05/13 11:22:07 4 finished
	2021/05/13 11:22:07 2 finished
	2021/05/13 11:22:07 1 finished
	2021/05/13 11:22:07 Err: goroutine 3 raised an error

Note that only the first error being returned will be returned by Wait. If there is another error happening between the context being cancelled and all the routines returning, then the won’t be accessible through the Wait output.

One thing to keep in mind, though, is that the context is cancelled when Wait returns, so don’t use it as a parent for anything that should run after that 😉

Personally, I’ve just recently started using this package. I’ve been aware of it for a really long time but somehow also confused it with other libraries dealing with nested errors instead of making the connection in my mind that this would be actually something related more to goroutines. Once I made the mental connection, though, I added it to more and more of my code and am loving it so far! 🙂
