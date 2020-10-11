---
title: 'Rust: More modular than expected'
date: "2020-10-11T22:10:19+02:00"
tags:
- rustlang
- learning
---

I’m still trying to learn [Rust](https://www.rust-lang.org) but slowly but steadily I’m getting used to it. Alongside completing various exercises on Exercism I’ve now also started coding on a little project, something that I’d normally write either in Go, NodeJS, or Python. It consists mostly of...

1. Fetching some data via HTTP
2. Parsing the response as JSON
3. Creating some basic XML
4. And finally serving that XML with a simple HTTP server

The last part is optional but the others should give me a good overview of Rust’s eco-system. Finishing steps 1 and 2 today already demonstrated quite a few differences to the Go eco-system that will take time getting used to. In Go you have a quite powerful standard library that already includes good tooling around HTTP incl. TLS and JSON encoding and decoding. With Rust I have to look for external libraries for these two. First stop: [hyper](https://crates.io/crates/hyper), for acting as HTTP client. 

That dumped me into a bit of a rabbit hole as it depends on [tokio](https://crates.io/crates/tokio) for handling server responses asynchronously. I hadn’t tried futures and async in Rust before as I’m just getting started and so only did the bare minimum to get hyper running (by declaring the main function as tokio’s entrypoint and marking it as `async` ).

I was surprised when I made the HTTP request to a https:// URL and received an error. I had somehow expected hyper to already come with TLS support. Turns out, this has been externalized into separate connectors (e.g. [hyper-rustls](https://crates.io/crates/hyper-rustls) and [hyper-tls](https://crates.io/crates/hyper-tls)).

In the end, I decided that async/futures/tokio was just too much of a diversion and so I moved on to [ureq](https://crates.io/crates/ureq). Compared to hyper it is much simpler but also blocking. For now that’s enough for me. It also integrates quite nicely with [serde\_json](https://crates.io/crates/serde_json), which I’ve picked for decoding the HTTP responses. It also allows for decoding data directly into native structs and also for lots of customization which I plan to play around with in the near future 🙂

So far, I’m having lots of fun here! Sure, at my current level I’d be much faster implementing that little project in Go, but there’s no time-pressure behind what I want to have implemented and so I’ll take my time learning as much as possible here 😉
