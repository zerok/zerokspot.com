---
title: Self-hosted comments with remark42
date: "2022-01-28T20:47:38+01:00"
tags:
- commenting
- selfhosted
- golang
incoming:
- url: https://chaos.social/@zerok/107701844591013470
---

Every now and then I think about re-adding classic comments on my blog. Some time ago I stumbled upon [remark42](https://github.com/umputun/remark42), a system that would allow me to self-host the commenting infrastructure and integrate it into a statically generated site. 

<figure><img src="/media/2022/Screenshot%202022-01-28%20at%2020.27.24.png"><figcaption>Commenting screen of remark42 in dark-mode with e-mail and anonymous comments enable</figcaption></figure>

Let me start with saying that I think that remark42 looks like a really promising and also actively maintained application. It is written in Go for the backend and React for the frontend. Thanks to an existing Docker image and no dependencies to an external database server it’s easy to set up and allows you to integrate with various OAuth providers like GitHub and Twitter for handling logins. It even supports my favourite auth-provider: e-mail 😍😉

A feature I at least think I need, though, is some kind of moderation system. If something new should appear on my personal website, I want to see it first to make sure it’s not spam or hate etc.. At the same time I want to make it easy for people to comment and therefore would enable anonymous and e-mail comments. So, basically, what I’m looking for is some kind of pre-moderation system. Sadly, this is something that remark42 currently doesn’t have but there are at least already some [discussions](https://github.com/umputun/remark42/discussions/1236) going on about it.

Note that this is mostly just an issue due to e-mail and anonymous comments. I would not expect there to be that much spam with OAuth-based providers but at the same time, I’m not willing to put that to the test right now especially since the supported providers are completely on the corporate side of the industry so far. Coming to think of it, perhaps I will just add a GitLab provider or something like that… Oh god, not another side-project for which I won’t have time…

To summarise: I still don’t have any plans to add classic comments to this site. That being said, I will definitely keep an eye on remark42 and might use it in future projects! It at least looks extremely promising!
