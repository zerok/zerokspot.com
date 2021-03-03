---
title: "Golang UK 2017"
date: 2017-08-20T19:02:09+02:00
tags:
- conferences
- travel
- london
- golang
---

As I've spent more and more time writing Go both at work and at home for
projects like [StrangeNewPlaces](https://www.strangenewplaces.org/), I've
decided to shift my focus towards Go-related conferences this year. This means
that besides [dotGo](https://www.dotgo.eu/) later this year I've also
attended [GolangUK](https://www.golanguk.com/) in London during the last couple
of days.

## This time with a workshop

The conference took place between August 15 and August 17
in [The Brewery](https://www.thebrewery.co.uk/), a very central and beautiful
event venue near the Moorgate station. This time around I also attended one of
the workshops as one was originally planned to be *The Ultimate Go* workshop by
William Kennedy. Sadly, he couldn't make it and the workshop ended up more
targeted at intermediate users. Still, a very good workshop held by Mark Bates
which also included a short introduction to [gRPC](https://grpc.io/)
and [Protobuf](https://github.com/google/protobuf) 🙂

## The main event

The main event was split into two tracks with tons of interesting talks. You can
find [the complete schedule](https://www.golanguk.com/schedule/) on the
GolangUK.com website but I just want to list all the things I've personally
learnt thanks to these sessions:

* If your package does HTTP requests, you should allow your users to pass in
  their own net/http.Client instance so that they can control aspects like
  timeouts. (*Writing beautiful packages* by Mat Ryer)
* As the channel reader can easily determine whether a channel was closed, the
  writer should always be the one closing it. I had known about this for quite
  some time but I somehow failed to see that the existence of the ok-return
  value when reading from a channel made this rule obvious. (*Concurrency
  Patterns in Go* by Arne Claus)
* If you have an endless loop without `time.Sleep`, at least add
  `runtime.GoSched()` to yield. (*Concurrency Patterns in Go* by Arne Claus)
* The Spinning Compare and Swap (Spinning CAS) pattern: Basically use
  `atomic.CompareAndSwap` in a tight-loop to handle locks without having to
  require things like mutexes. (Concurrency Patterns in Go by Arne Claus)
* Pretty much everything
  from
  [Sean Kelly's awesome talk about embedding](https://github.com/stabbycutyou/embeddingtalk). This
  is probably the one language feature I know the least about so I really
  appreciated this long-form introduction.
* The [errgroup package](https://godoc.org/golang.org/x/sync/errgroup) for
  handling synchronisation of multiple go-routines. (*How to correctly use
  package context* by Jack Lindamood)
* [go-i18n](https://github.com/nicksnyder/go-i18n) is used for instance
  by [Buffalo](http://gobuffalo.io/) for internationalisation. (*Rapid web
  development in go* by Mark Bates)
* Pretty much everything in Buffalo can be replaced except for gorilla/mux as
  router. (*Rapid web development in go* by Mark Bates)
* Buffalo provides its own database migration framework
  with [Pop](https://github.com/markbates/pop). (*Rapid web development* in go
  by Mark Bates)
* [tfortools](https://github.com/intel/tfortools) by Mark Ryan looks like the
  perfect library for providing templating support for your command-line
  tool. It not only contains a handful of extremely useful template functions
  but also automatically generates the documentation for what is available
  within the templates. (Command-line scripting with templates by Mark Ryan)
* Pretty much all of *Production ready go* by Ian Kent. Most of the things in
  there are well-known by now. It's still great to have all the timeout-hints
  collected in one talk 🙂
* [golang.org/x/net/websocket](http://golang.org/x/net/websocket) only
  implements a subset of the RFC6455 websocket spec, hence you should probably
  prefer gorilla/websocket
  or [github.com/gobwas/ws](https://github.com/gobwas/ws).

Sadly, I missed Brian Ketelsen's final keynote about *Go for the Enterprise* but
I hope I will be able to watch it online in a couple of days 🙂

## The social event

On the evening of the first talk-day there was also a nice social event with BBQ
and drinks. The weather was still nice enough, so we could enjoy the tiki bar
without having to run between roofs 😉 While the idea of offering a completely
customisable BBQ with dishes like pulled pork, kimchi, raw salmon and the likes
was awesome, it sadly didn't scale that well for more than 400 attendees. After
40 minutes I still saw some people in the queue.

## The venue and catering

That being said, the food there and during the conference as a whole was
excellent. Everything was well prepared and the whole organisation was just
flawless.

Same goes for the venue. The seating was comfortable, thanks to some additional
TV-screens I could see the slides even from the very back of the room, and there
was water readily available everywhere. Perhaps the one thing I'd improve about
the catering next year would be fruits instead of bonbons on the tables. But
perhaps that's a British thing 🙂

## Conclusion

Over the course of these three days I've learnt so much that this post can only
offer a very brief summary. I really, really liked the conference and cannot
wait for next year's GolangUK 🙂 Sadly, London is probably not the best city for
me and my health right now, but the trip was still well worth it! Big thanks to
the organisers, speakers, sponsors, and everyone else that made this such a
great conference!

Perhaps next year I will even dare to speak to other people and not hide behind
my laptop or iPad 😢
