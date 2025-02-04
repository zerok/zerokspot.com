---
title: "Getting into the flow with Pomodoro"
date: 2019-01-08T20:30:47+01:00
tags:
- productivity
- pomodoro
- gtd
- workflow
---

In mid-November my mental workload finally reached a dangerous level, so I started to experiment with the [Pomodoro technique](https://francescocirillo.com/pages/pomodoro-technique) ([archive.org link](https://web.archive.org/web/20230306054358/https://francescocirillo.com/products/the-pomodoro-technique). For those who don't know it, Pomodoro is a time management method developed by Francesco Cirillo which focuses on creating highly focused work time intervals interleaved with breaks. Usually, every work iteration has 25 minutes followed by a pause of five minutes. After four such iterations there is a longer break of around 20-30 minutes.

With Pomodoro you have a timer somewhere lying around that enforces these time limits. A quick search through the AppStore resulted in far too many options but luckily also included [BFT, the Bear Focus Timer](https://itunes.apple.com/us/app/bft-bear-focus-timer/id1328806990). BFT has the additional benefit, that it uses the phone audio output to play some "rain music" which makes staying focused easier. The only disadvantage I've noticed there yet is that when used throughout the day it needs quite a lot of battery power.

That being said, these early experiments had some very encouraging results: I got tons of work done, far more than without the strict time intervals required by Pomodoro. Stress also went back to a more or less acceptable level. I only had one issue with this workflow: At the end of each day I no longer knew what I had worked on. I knew that I had achieved a lot but I wanted to have some kind of summary over all my completed tasks. There are just some things that don't fit all that well into my GTD system. Some things are just too short-lived to even make it into my inbox yet are worth mentioning inside some kind of log.

For this use-case I've now created a pair of small shell scripts:

* `it-log` opens `$HOME/Documents/iterations/$DATE/$ITERATION.md` inside VIM and I can document what I've been working on using Markdown
* `it-next` increments the `$ITERATION` counter for the current day

In the next couple of days I will probably also create a `it-summary` script for generating a simple report for every day. But even without that I'm extremely happy with my setup. I've felt so productive today in a way I haven't felt like for years. While working this way for a whole day is very exhausting, I think the added motivation is well worth it!
