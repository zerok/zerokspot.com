---
title: "JAMstack"
date: 2017-06-30T10:08:52+02:00
tags:
- development
---

In a recent [Changelog episode](https://changelog.com/podcast/251) Matt Biilman
and Chris Bach talked about a technology pattern
called [JAMstack](https://jamstack.org) and how they used it to implement the
new version of the Smashing Magazine website.

The idea here is that you generate most of your **m**arkup statically using
something like Hugo or Jekyll and add dynamic components through **J**avaScript
and the use of simple **A**PIs. This somehow reminded me of what I've been doing
for a while here and more recently also on PyVideo.org. There we generate all
the primary content pages using Pelican out of a bunch of JSON files. The one
dynamic component here is the search engine, which we've integrated with
JavaScript backed by a search API written in Go.

But that's just the beginning, as Matt and Chris have shown. They and their team
at [Netlify](https://www.netlify.com/) have created a whole toolbox of services
ready for this kind of web application. Among them is
a [CMS](https://github.com/netlify/netlify-cms) and
an [e-commerce service](https://github.com/netlify/gocommerce). With these you
can basically store all your content inside services and have a build-step
compile it out into static pages, basically replacing what a Git repository has
been for us with a custom content service API-layer.

But these are not the only services available out there. There are, for
instance, some commercial ones like [Contentful](https://contentful.com), which
provides a CMS similar to what netlify-cms does. If you prefer on-premise
solutions, there are things like [Hexo](https://hexo.io/) (a blogging system)
and [Ponzu](https://github.com/ponzu-cms/ponzu) (another simple CMS implemented
in Go).

Especially for content-heavy sites, JAMstack looks like a great approach. You
basically have cronjobs or events that are triggered when the content in a
repository or CMS changes and that statically rebuild the affected parts of the
website. Everything else is loaded dynamically from APIs and displayed through
JavaScript. Sure, this doesn't work for every kind of website, but for news
sites, online shops etc. I can definitely see its appeal 😀

I already have a couple of ideas for GoGraz where that might come in handy. Same
goes for the [Meet-The-Meetups.org](https://meet-the-meetups.org) where we are
already fetching some of the content from the meetup.com API. More on that,
though, when I actually come around implementing it 😉
