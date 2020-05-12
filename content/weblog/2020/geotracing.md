---
title: "Geo-tracing my life"
date: "2020-05-12T08:24:31+0200"
tags:
- overland
- geotrace
- 100daystooffload
---

On Sunday I stumbled across a post by Aaron Parecki in which he describes how [he's been tracking his movement](https://aaronparecki.com/gps/) for the last couple of years. For many years I've been using Swarm/Foursquare to at least keep track of places I visit when I'm a tourist. Quite often, though, I wanted a way to visualise also my daily paths *without* having to rely on third-party services and in the process potentially expose my location to someone else.

While our use-cases for wanting this data may differ, luckily Aaron also wrote a little app for iOS that collects location data and then submits it to a server that can be defined by the user. About an hour after discovering [Overland][o] I had created a simple HTTP server in Go that writes the received coordinates into a CSV file. Another 30 minutes later I had that service deployed on one of my servers and storing the data onto an encrypted volume. I just love [Digital Ocean Volumes][v] 😍

Just in case this sounds interesting to you, check out [zerok/geotrace][g] on GitHub. Right now, geotrace only supports appending data to a CSV file but in the next couple of days I also want to add SQLite support to it!

As for actually using the collected data: Once SQLite support is in there I want to write a little exporter that creates GPX files which I can then use with OpenStreetMap et al. or just use it to attach proper coordinates to the pictures I take with my Sony camera. I probably have far too many things I want to do with that data to actually start doing any of them... 🤯

[o]: https://overland.p3k.app/ 
[v]: https://www.digitalocean.com/products/block-storage/
[g]: https://github.com/zerok/geotrace


