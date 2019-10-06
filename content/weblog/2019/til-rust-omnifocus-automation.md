---
date: 2019-10-06T21:34:20+02:00
title: "This weekend's learnings: Basic Rust and automating OmniFocus"
tags:
- til
- rust
- applescript
- omnifocus
- automation
---

Being slightly sick can also have its advantages. For instance, this
weekend I had a bit of time to listen to some podcasts and learn some
new stuff! 

First, I ported some code I had previously written to manage my
personal knowledge base from [Go][] to [Rust][] in order to get a feel
for the language. I've wanted to write something in Rust for quite
some time but eventually opted to just take a project I had barely
started and port it over.

The [code that I have right now][ds] is rough, bad, and probably as
far away from idiomatic as possible, but I want to improve upon it in
the future with things like logging and proper error handling. The
project itself is small enough that iterating ideas and concepts here
should be easily doable.

On Saturday, I also listened to [episode 5 of the NestedFolders
podcast][nf] in which the hosts discuss how they are doing GTD
reviews. Scotty Jackson, for instance, does the review always in two
phases on Friday and then on Monday. I totally forgot that one can use
AppleScript to manipulate items in OmniFocus en-masse. The hosts
mentioned various applications of that in order to set review-dates
and so I looked a bit around finding [this script][jb] by Joe Buhling.

Taking some inspiration from that I created custom version of it, more
tailored to my project structure. If that and the weekly review
(Sunday for personal stuff, Monday for work projects) works out, I
will probably write a follow-up in the near future 🙂

[rust]: https://rust-lang.org
[go]: https://golang.org/
[ds]: https://github.com/zerok/datasphere
[jb]: https://github.com/joebuhlig/OFScripts/tree/master/Update%20Reviews
[nf]: https://nestedfolderspodcast.com/podcast/episode-5-how-we-weekly-review/
