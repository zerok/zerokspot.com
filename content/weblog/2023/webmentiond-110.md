---
date: "2023-08-26T13:57:40+02:00"
tags:
- webmentiond
- releases
- projects
title: Webmentiond 1.1.0 released
incoming:
- url: "https://chaos.social/@zerok/110956132451395833"
---

After two years I thought it was time to finally put some of those changes that had been made to [Webmentiond](https://github.com/zerok/webmentiond) actually also into a numbered release again. So after some cleaning up and tooling improvements last night, I've now pushed [version 1.1.0](https://github.com/zerok/webmentiond/releases/tag/v1.1.0) of Webmentiond out this morning 🥳

Changes include:

- Support for STARTTLS using `EMAIL_USE_STARTTLS` environment variable
- Add `SERVER_AUTH_JWT_SECRET` environment variable
- Improve detection of likes and comments (in nested `h-like`s)
- JWT secret can now be configured via an environment variable
- Expose metrics only if a `--metrics-addr` is set
- Adding version data to binary (and `--version` flag)

You can get the new version via DockerHub: `zerok/webmentiond:v1.1.0`.

Big thanks to everyone who has used the project over the years and especially those who've provided feedback and improvements! ❤️

With the new tooling changes in place, creating proper releases should be much easier going forward so you will probably see more of these in the future!
