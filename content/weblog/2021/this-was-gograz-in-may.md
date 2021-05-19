---
title: This was GoGraz in May
date: "2021-05-19T18:22:37+02:00"
tags:
- gograz
- golang
- 100daystooffload
---

Sadly, I only again had only very little energy left for the [GoGraz meetup this month](https://gograz.org/meetup/2021-05-17/). Long workdays with lots of video calls, learning for a re-certification, and lots of other tasks simply left too little to stay in the call for longer than 45min this time around. That being said, we still had some awesome topics, albeit sadly no talk:

- [modernc.org/sqlite](https://pkg.go.dev/modernc.org/sqlite?utm_source=godoc)is an automated port for SQLite to Go to get rid of the CGo dependency in your projects that want to use SQLite.
- [Litestream](https://litestream.io/) is a replication tool for SQLite databases that basically streams your write-ahead log into a S3-compatible storage.
- [Use the Index, Luke](https://use-the-index-luke.com/)is a developer-friendly guide to database performance that mostly tries to stay vendor-agnostic.
- Matthias and Chris have talked a bit about their experience with [Vitess](https://vitess.io/). I somehow always associated large setups with this and therefore never gave it a try, but perhaps I finally should 😅 It’s basically orchestrating multiple MariaDB/MySQL server to get around some of the limitations of the single-instance setup, for example when it comes to renaming tables and doing migrations in general.
-  If you’re looking for time series databases and like PostgreSQL, you might want to take a look at [TimescaleDB](https://www.timescale.com/) and [Citus Data](https://www.citusdata.com/).

That’s it! Sadly, I was simply too exhausted to stay online for a minute longer but hopefully I’ll have more energy next time around!

Btw.: If you want to give a talk at GoGraz (we’re remote right now 😮) just contact me on `talks@gograz.org` and also take a look at [gograz.org](https://gograz.org/) 🙂
