---
date: "2025-01-26T11:47:46+01:00"
incoming:
- url: https://bsky.app/profile/zerokspot.com/post/3lgnbt3gahc2j
tags:
- bluesky
- meta
- zerokspot
- samara
title: Integrating Bluesky into static websites using Samara
---

Ever since I started using Bluesky, I wanted to also integrate replies to my post into this website. I have done the same with Mastodon using [retoots](https://github.com/zerok/retoots) and so I started working on something similar for Bluesky: [Samara](https://github.com/zerok/samara). 

One of the great things about the Bluesky API is that posts etc. are accessible without any kind of authentication. Many folks have gone out and just used that for integrating posts into their static websites. While that’s definitely a viable approach, I wanted to have a layer between my website and Bluesky that is under my control. 

In the current implementation this offers the following advantages:

- If you go to zerokspot.com and read a post, your IP is not leaked to Bluesky. The actual API requests are done through my server. This also includes avatars as they are also proxied.
- There is server-side caching happening. Right now this is only there to prevent the server from hammering the Bluesky servers. In the future I also want to use that as fallback if upstream is down so that you can still read the comments on my site even if Bluesky itself has issues.

To make integrations easier, Samara also supports HTMX so that you don’t have to write your own JavaScript frontend. You can find a complete example setup inside the [examples/simple-htmx](https://github.com/zerok/samara/tree/main/examples/simple-htmx) folder 🙂

Eventually, I want to remove the caching and telemetry features of Samara into retoots and also expose things like “favorites” and “reposts” within Samara.

For now I hope that this is also useful for someone else 🙂
