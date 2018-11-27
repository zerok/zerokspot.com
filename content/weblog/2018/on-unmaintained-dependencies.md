---
title: "On unmaintained dependencies"
date: 2018-11-27T18:50:22+01:00
tags:
- development
- ecosystem
---

In case you missed it, over the course of the last two days quite a few news sites reported on a malware attack involving the event-stream npm package (among others). [The Register](https://www.theregister.co.uk/2018/11/26/npm_repo_bitcoin_stealer/) has a nice, compact write-up if you're curious. What follows here is just a little collection of my thoughts on this incident and the environment that made it possible.

Looking at the [original issue on Github](https://github.com/dominictarr/event-stream/issues/116) about all of this, you'll notice that sh\*t hit the fan long before the world took notice: 20 Nov 2018. By reading [the timeline](https://medium.com/@cnorthwood/todays-javascript-trash-fire-and-pile-on-f3efcf8ac8c7) composed by Chris Northwood and the changelog (if you can call a log with a single commit a log) of the [flatmap-stream](https://github.com/hugeglass/flatmap-stream) project it looks like the first pieces had been put into place in September or even earlier.

Sadly, everything about this on Github including the original tickets has become nearly unreadable as people moved on from fixing the issue to laying blame left and right but mostly on the original author (Dominic Tarr) who handed the project over quite a while ago. It's easy, it's cheap, and it completely ignores that nothing he did there was uncommon: Passing maintainership of a project on without thorough vetting.

It also demonstrates an often ignored problem inside the OSS community: What happens to a project when the original author is no longer interested in a project? Historically, there have been two approaches:

1. Don't do anything and just mark the project as archived. The community will fork it if necessary under a new name.
2. Hand the project over to a new maintainer.

This all assumes that a fork will not eventually end up being able to make new releases under the name of the original project. Such an ecosystem would basically merge approaches 1 and 2 as even a fork could start to look like the original project with a new maintainer.

Combine that with people not pinning their dependencies (and n-th level dependencies) to exact versions which leads to "patch" releases making it into the release processes of unsuspecting projects. Honestly, it's quite surprising that incidents like the one involving event-stream don't happen more often.

I don't intend to paint handing commit-bits over to someone else as "the wrong way to do it", though. Some projects have built up complete teams around them so trust *is* possible. Sometimes projects end up at corporate sponsors or foundations like the Apache Foundation. There has to be a process, though. Just randomly picking the first person who raises a hand when asked if they want to take over is definitely the wrong way. There absolutely *is* some responsibility that comes with releasing your code esp. once its getting popular.

Every community seems to be dealing with EOL'ing projects in its own way: The Django community, for instance, has created the [jazzband](https://github.com/jazzband/) project. This summer, some members of the Go community started doing something similar with [Gof.rs](https://gof.rs/). Perhaps the recent events will birth something similar for the JavaScript community if there isn't one already.

I just hope, that all of this reminds people that (1) keeping dependencies maintained is not only the job of the maintainer but also the community around it and (2) that you as a user of a library are responsible for keeping an eye on your dependencies. Or, [as Scott Hanselman put it on Twitter](https://twitter.com/shanselman/status/1067363631354925056):

> Who is responsible for your massive dependency graph? You. #minergate



