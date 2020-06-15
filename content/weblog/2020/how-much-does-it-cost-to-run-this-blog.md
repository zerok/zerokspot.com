---
title: How much does it cost to run this blog?
date: "2020-06-15T08:33:12+02:00"
tags:
- blogging
- zerokspot
- infrastructure
- 100daystooffload
---

Kev posted yesterday a [nice run-down of what it costs to run his blog](https://kevq.uk/how-much-does-it-cost-to-run-this-blog/). I’ve wanted to write about this for quite a while now and so I thought this might be a good time to finally actually do it 😅

My setup is quite simple on the surface with the biggest fraction running into [DigitalOcean for a small VPS (1GB)](https://www.digitalocean.com/pricing/) and some extra object storage for larger files (e.g. photos, software packages, …). Previously, I was on a larger VPS but after removing some services altogether and moving others to dedicated instances zerokspot.com can now comfortably live on the smallest VPS offering with resources to spare! Going with Hugo and just adding some tiny APIs on top of that has definitely paid off here!

| Item | Monthly costs |
|-|-|
| Domain name | $15.17/12 = $1.26 |
| VPS on DigitalOcean (1GB)| $6 |
| Space on DigitalOcean for photos | $6 |
| [Algolia](https://www.digitalocean.com/pricing/) | $0 (free tier) |
| **Total** | **$13.26** |

There are also some shared costs that I don’t include in this table like my GitHub Pro account, Fastmail, etc.. While they are somehow related to the website, they are also used for tons of other projects.

Another item that’s not in this calculation is the time I spend every month to maintain the blog and add or update content. If I somehow convert that time into financial costs they’d most likely dwarf those $13.26 I’ve listed above 🙂

Please note that you can absolutely start your own blog for free but if you want to run a self-hosted setup where pretty much everything including the operating system is under your control, you should probably get a virtual-private-server somewhere and they cost usually around $5 per month. But even if you go with a fully managed solution like a hosted Wordpress instance, you should definitely get your own domain (around $9 - $16 per *year*) which will allow you to migrate your blog from one host to another.
