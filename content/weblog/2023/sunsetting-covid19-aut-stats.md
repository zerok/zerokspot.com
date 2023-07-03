---
title: Sunsetting covid19-aut-stats
date: "2023-07-03T15:18:29+02:00"
tags:
- covid19
- opendata
---

Right when the COVID19 pandemic started I wanted to have a little service that sent me every day the current number of infections etc. in Austria. For this I started scraping various sources and stored their numbers into a simple CSV file onto GitHub: [covid19-aut-stats](https://github.com/zerok/covid19-aut-stats/). You can read all about it [here](https://zerokspot.com/weblog/2020/03/28/covid19-aut-stats/).

Well, it’s now been more than three years and the numbers aren’t going anywhere anymore. So I’ve decided to stop the scraper and also the associated datasette service running on `covid19-aut-stats.h10n.me`. That being said, I intend to keep the repository in place but have marked it as archived. 
