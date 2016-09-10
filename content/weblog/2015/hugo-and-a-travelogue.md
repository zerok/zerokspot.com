---
date: '2015-04-06T16:48:48+02:00'
language: en
tags:
- zerokspot
- travelogue
title: Hugo and a Travelogue
---


For nearly 5 years now I've been writing blog posts exclusively in a text editor
and publishing them through Git and post-receive hooks. Since early 2010 I've
for that purpose used a completely custom system but over the years I've removed
more and more features, making that system more minimal with every change. Up
until about a year ago this site still featured complete full-text
searching, for instance. Did I ever use it? About once every other month. Did
anybody else use it? Pretty sure the answer is "no" here.

Right now this site is still mostly powered by Git and ElasticSearch with a
small Go-layer to merge everything but I've started to look at some of the
static blog generators out there (mostly [Jekyll][] and [Hugo][]) to see if I
could get enough power with either of them to replace my custom solution (and
some some megs of RAM on my VPS ;-) ).

---------

I read that [Jekyll][] might have problems with larger blogs so I've opted to
give [Hugo][] a try first for a travelogue I've wanted to make for quite some
time now. It definitely helps that PyCon is about to start in a couple of days
so I'll have some material to post there right away!

Because this whole thing is bit of an experiment some things might break here
and there over the next weeks but if you enjoy me exploring the world, you will
be able to find some pictures, stories and reviews over at
[travelogue.h10n.me](http://travelogue.h10n.me) ([source][src]).

I've also written a small post over there about
[why I wanted to create a travelogue](http://travelogue.h10n.me/post/2015/welcome/)
in the first place.

[hugo]: http://gohugo.io
[jekyll]: http://jekyllrb.com/
[src]: https://github.com/zerok/travelogue
