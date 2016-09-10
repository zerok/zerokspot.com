---
date: '2016-04-04T21:37:59+02:00'
language: en
tags:
- django-channels
- django
- djangocon
title: Channels in Docker example
---

As part of the DjangoCon Europe 2016 sprints in Budapest last weekend I wanted
to get know Andrew Godwin's [Channels][] framework for Django a bit better
especially since it might find its way into Django 1.10 and I'm in need of a
Websocket-solution anyway right now 😊 Given that channels are supposed to be run
using at least two distinct servers I thought it might be nice to visualise that
setup using docker-compose.

[channels]: https://github.com/andrewgodwin/channels

You can find the result
[here](https://github.com/zerok/channels-in-docker). Right now it doesn't do
anything beyond a little hello-world, but I want to extend it with at least a
tiny Websocket example in the near future. Nothing fancy, just the bare minimum
😉
