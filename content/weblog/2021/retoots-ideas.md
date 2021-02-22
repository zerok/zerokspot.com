---
title: "Some ideas for retoots"
date: "2021-02-22T20:30:00+01:00"
tags:
- mastodon
- retoots
- 100daystooffload
incoming:
- url: https://chaos.social/@zerok/105776615039607886
---

Today, I finally found the couple of minutes to integrate [retoots][] into
[GoGraz.org](https://gograz.org). While doing that I came up with a couple of
things that I'd like to improve about it in the near future:

- Create an ansible-role so that you can easily install it on your server using
  [Ansible][].
- Create a simple JavaScript that you can integrate into your website. Perhaps
  I will even use Go 1.16's embedding feature for that.
- Include boosts in the `/api/v1/interactions` result.

It's not yet clear when I'll find time for all of these but at least they're
now on my roadmap 😉

[retoots]: https://github.com/zerok/retoots
[ansible]: https://www.ansible.com/
