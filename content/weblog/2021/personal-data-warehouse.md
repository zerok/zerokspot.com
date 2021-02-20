---
title: Personal data warehouse
date: "2021-02-20T20:27:50+01:00"
tags:
- 100daystooffload
- personal
- datawarehouse
- pim
---

A couple of months ago, [Simon Willison posted a talk](https://simonwillison.net/2020/Nov/14/personal-data-warehouses/) he gave about building a personal data warehouse based on all the data that can be exported from the various services we are using or that are simply out there. Sadly, I only read it today but it was very inspirational!

So far I’ve only dabbled a little bit into actually using and combining that data with my [geotrace](https://github.com/zerok/geotrace) project. That talk gave me some new ideas and so I might, in the future, try to also import such exports continuously into either SQLite or PostgreSQL data stores to make them reusable.

SQLite would have the advantage that I could easily share that data and use tools of Simon’s [Dogsheep project](https://github.com/dogsheep) (which is a great pun on Wolfram Alpha 😅).  On the other hand, PostgreSQL would offer me more powerful features but would, at the same time, require some custom setup for having a local PostgreSQL server running and not interfering with my other systems. In the end, it might be a hybrid solution but perhaps I should just get started and worry about the details later 😉

I really like this idea of building your own personal data warehouse! I think I will simply start with the dumps I’ve received from Goodreads, Swarm et al. and see where it leads me. No matter what, I will put the data in close proximity to [my personal information management system](https://zerokspot.com/weblog/2021/01/19/creating-my-own-pim-system/) 😉
