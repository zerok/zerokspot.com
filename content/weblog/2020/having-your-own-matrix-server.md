---
title: "Having your own Matrix server"
date: "2020-04-29T08:06:46+0200"
tags:
- matrix
- communication
- 100daystooffload
- decentralized
---

For many years now I've been looking for a place for private family discussions that was for the most part under our control. Sure, [Signal](https://www.signal.org/) is private but I wanted to have a system that I could more directly interact with yet that allowed for e2e encryption where needed. At first I thought about just setting up an [XMPP](https://xmpp.org/) server but since there are now competing systems for e2e encryption there and no decent iOS or macOS client, I quickly gave up.

Chaos congresses and FOSDEMs came and went and I still hadn't found a solution I was happy with but [Matrix](https://matrix.org/) became more and more interesting. It's rather new compared to XMPP but I really like the idea of treating chats like state containers that are modified by events. Having those events never actually leave a home server if all participants are located there is even better for my use-case. Oh, and I like that [Riot.im](https://about.riot.im/) is a decent client that I can use pretty much whatever platform I'm working on.

When COVID-19 happened and all of a sudden we were constrained to our own four walls I decided to just re-use one of my DigitalOcean instances, set up some disk encryption and deploy Synapse, the reference implementation for a Matrix home server, onto it.

On March 28 I was finally happy enough with the setup that I invited my partner onto it and we've been using it ever-since without any relevant downtime or issues so far. At this point I'm happy enough with it that I will probably also start to invite other family members. Let's see how it goes 😅

As for the setup itself:

- A 2GB droplet on [DigitalOcean](https://m.do.co/c/6a69f8676d65) (referral link) with Ubuntu
- The latest [Synapse](https://github.com/matrix-org/synapse) release
- [PostgreSQL](https://www.postgresql.org/) 11 as data store
- [Nginx](https://nginx.org/) as ingress proxy with [Certbot](https://certbot.eff.org/) for fetching TLS certificates from Let's Encrypt

That's pretty much it. Nothing fancy and not a lot of customisation yet but it has worked quite well so far 🙂 There was also surprisingly little configuration needed to get this working. Thankfully, there is even a script that converts a SQLite database (default configuration for Synapse) into a PostgreSQL one. If you have a spare VM somewhere (or one with enough capacity for another application running on) [give it a try](https://github.com/matrix-org/synapse/blob/master/INSTALL.md) 🙂
