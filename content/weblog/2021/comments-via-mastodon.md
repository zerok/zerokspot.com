---
title: Comments via Mastodon
date: "2021-01-07T10:48:27+01:00"
tags:
- zerokspot
- comments
- mastodon
- retoots
---

In [May 2019](https://zerokspot.com/weblog/2019/05/01/removing-comments-form/) I removed Disqus from this site simply because it had turned into an ad-network. Not wanting to “own” other people’s comments also didn’t replace it with something custom made/self-hosted. About a year later I had stumbled upon webmentions and started implementing it here using the [webmentiond](https://github.com/zerok/webmentiond) project. Nearly exactly a year after removing Disqus [webmentiond supported comments](https://zerokspot.com/weblog/2020/05/05/webmentiond-supports-comments/) and I embedded them below my blogposts.

When I read [Carl Schwan’s post about embedding Mastodon comments](https://carlschwan.eu/2020/12/29/adding-comments-to-your-static-blog-with-mastodon/) directly using the Mastodon API back in December I knew I want to have that too! He solved this by fetching comments client-side via JavaScript but I wanted to have a little more control over what is shown on my blog.

Frustrated with the world I set out to solve that problem for me yesterday. A couple of hours later I had a rough implementation deployed on my website:

Similar to webmentiond there is now another service called “retoots” running on my server. When I have a post that should allow commenting via Mastodon, there is a small JavaScript embedded in that post that fetches replies to a status linked to this blog post from retoots.

retoots, in turn, then contacts the Mastodon server that status is stored at and fetch all descendants through its API, normalises the data a little bit, and then returns that to the blog for rendering. This has the advantage that I can, in the future, add some _content filtering_ (e.g. not rendering posts by known spammers etc.) and also some _caching_ so that comments are also shown if the Mastodon server is down.

From a privacy point of view this also means that readers of my blog won’t contact a third-party Mastodon server directly and therefore their IPs and other metadata won’t show up in the access logs there but just on my own server.

You can find the current state of [retoots on GitHub](https://github.com/zerok/retoots). At this point it’s mostly a proof of concept with very limited documentation and not even proper releases. If everything works out that will change, though, in the near future 😀
