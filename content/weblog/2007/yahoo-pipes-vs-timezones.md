---
date: '2007-07-19T12:00:00-00:00'
language: en
tags:
- pipes
- timezones
- yahoo
title: Yahoo! Pipes vs. timezones
---


If you have taken a closer look at my lifestream, you might have noticed that the order of the items seems not to be completely correct. I'm currently trying to trace this problem down, but it looks like the culprit might be Yahoo! pipes and not the date-conversion I'm doing on my end.

-------------------------------

I'm currently quite sure that this is the case, because I actually don't change the order on my end but let Yahoo! pipes do the whole feed merging. I only then try to convert all the dates to CEST.

So far I think, that Yahoo! Pipes might have a problem with feeds using different timezones. To demonstrate this I built a small example pipe that only tries to combine my weblog feed (UTC) with my bookmarks feed provided by ma.gnolia which sends the with a UTC-7 timezone. Both are using Atom as feed format.

A screenshot of my configuration for this merge can be seen here:

<img class="figure" src="/media/2007/yahoopipestz.png" alt="Merging two feeds using Yahoo! Pipes"/>

The problem is now, that the generated order is completely wrong:

* 2007-07-19T15:32:01Z
* 2007-07-19T11:28:25-07:00
* 2007-07-18T21:18:21Z
* and so on

The last time I checked, 11:28:25-07:00 should be **after** 15:32:01Z ...

Yahoo! Pipes seems to have [some history](http://discuss.pipes.yahoo.com/Message_Boards_for_Pipes/threadview?m=tm&bn=pip-DeveloperHelp&tid=1711&mid=1711&tof=25&frt=2) with timezone issues which led some [brave people](http://discuss.pipes.yahoo.com/Message_Boards_for_Pipes/threadview?bn=pip-DeveloperHelp&tid=1800&mid=1803) to try to solve the problem on YP's side using [quite complex pipes](http://pipes.yahoo.com/pipes/pipe.edit?_id=wIO7C0E13BGwIqYsdbq02Q). I honestly don't want to go this way, so I will have to look around for alternatives in the near future :-(
