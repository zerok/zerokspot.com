---
title: Know when to abort a project
date: "2021-03-09T18:40:20+01:00"
tags:
- side-projects
- 100daystooffload
incoming:
- url: https://chaos.social/@zerok/105861107110361910
---

For the last couple of months I’ve had a side-project on the back burner that should allow me to consume some Instagram accounts via RSS. I’ve previously tried some third party services like [rss.app](https://rss.app/) but simply wasn’t willing to pay for yet another subscription just for getting updates to 2-3 Instagram accounts…

So I started looking at the public data available when going to, for instance, [theunipiper’s profile](https://www.instagram.com/theunipiper/)and querying the endpoints directly. This project should also have been my first playground project for Rust but I eventually could no longer find time for it. The approach was just hacky and so I wanted to do two things:

1. Move the data retrieval over to using the proper Instagram API
2. Move the implementation over to Go so that I’d eventually get done with this project and be able to move on.

While (2) shouldn’t be a problem, (1) turned out to be a show-stopper: In order to get proper API access you need not only an Instagram account but have it also linked to a Facebook account. All new Instagram accounts seem to require a phone number and all of a sudden I was running from one privacy nightmare to the next just to get access to 10 photos a week…

Yesterday I’ve now done something that I should do far more often with my hobby projects: Re-evaluate the costs and benefits and, if it looks quite bleak, drop it. In the case of “Instafeed” the benefits were just too small and since I had already decided to *not* code it in Rust there was also no real learning effect left.

Perhaps I will eventually give the whole thing another try but for the foreseeable future I’ve removed it from my project list 🙂
