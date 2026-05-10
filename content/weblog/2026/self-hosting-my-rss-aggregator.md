---
title: "Self hosting my RSS aggregator"
date: "2026-05-10T20:52:00+02:00"
tags:
- rss
- selfhosting
---

I've been using [Feedbin](https://feedbin.com/) since 2008 after moving there from Feedly. Back then, Feedly had become somehow annoying with lots of integrations that I actually didn't care about. This is not the reason why I decided to look at alternative RSS aggregators this time. I really like Feedbin and haven't had any real problems with it ever. At the same time, RSS aggregation just feels like something that I could solve for myself with a little service running on one of the servers I already have for other things.

So I gathered a couple of options:

- [TinyTiny RSS](https://tt-rss.org): Written in PHP
- [FreshRSS](https://www.freshrss.org): Also written in PHP...
- [Miniflux](https://miniflux.app): Finally something not written in PHP :D
- [commafeed](https://github.com/Athou/commafeed): And a Java implementation

Even with Docker I don't want to have to deal with PHP when I can avoid it. Same with Java plus I'm actually looking for something really lightweight. So I went with Miniflux for now which is as minimal as it gets.

<figure>
<img src="https://zerokspot.com/api/photos/2026/05/10/miniflux.jpeg?profile=1024" alt="">
<figcaption>Miniflux's front-page is quite simplistic.</figcaption>
</figure>

Still, the interface offers exactly what I want: A list of unread items from all my feeds and the option to only see those of a specific feed. Combined with good keyboard shortcuts   I've so far not even opened Reeder on my phone since starting to use Miniflux.

It also offers tons of things that I haven't used yet like ... 

- [Scraper rules](https://miniflux.app/docs/rules.html#scraper-rules) for extracting the actual post content from a website using CSS selectors
- [Entry filtering rules](https://miniflux.app/docs/rules.html#feed-filtering-rules) which would allow me to skip for instance those articles on Wired.com that are just discount codes/advertisements.
- Webhook support for specific feeds so that I could trigger other services from them.

The first week with Miniflux has been quite pleasant. So far it looks like it does exactly what I want! As a side-effect of the migration, I also found tons of feeds that were not working or that I wasn't interested in anymore 😅
