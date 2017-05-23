---
date: 2017-05-23T11:20:13-07:00
title: A search-feature for PyVideo
tags:
- pyvideo
- pycon2017
- golang
---

Up until now, [PyVideo](http://pyvideo.org) had only a rudimentary search
feature using Google Site Search. This is not ideal and downright bad on mobile
as it's using iframes which don't work well in that environment.

As the site is statically generated
using [Pelican](https://github.com/getpelican/pelican), we have to integrate any
kind of dynamic search-feature using JavaScript. A couple of months ago I
implemented exactly that for my own blog here so I thought I should be able to
solve this also for PyVideo in a couple of days. An ideal task for the PyCon
sprints 😀

This new search feature should be as light weighted as possible as we don't want
to have to run something like an ElasticSearch cluster for that. Since the site
is about Python, an ideal implementation would be using asyncio in combination
with a simple search index like whoosh. Sadly, that seems not to support asyncio
and so I opted to go with an implementation in Go
using [bleve](http://blevesearch.com). You can find the
implementation [on Github](https://github.com/zerok/pyvideosearch) if you want
to play around with it 😊

<figure>
<img src="/media/2017/pyvideo-search.png" alt="" />
<figcaption><p>The search-feature running on my local machine.</p></figcaption>
</figure>

To host this service and stay as flexible as possible, we have created a new
sub-domain called `api.pyvideo.org` where we might offer other API-ish services
in the future.

The integration into Pelican required a search content-page with a custom
template where I could add the necessary JavaScript. Also straight forward but
something I haven't done in a while as I'm now using React in combination with
Webpack etc. for most of my projects. You can find the PR for this change
again [on Github](https://github.com/pyvideo/pyvideo/pull/77) 😊
