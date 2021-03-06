---
date: '2008-06-29T12:00:00-00:00'
language: en
tags:
- api
- development
- ratelimit
- twitter
title: Know your rate-limit on Twitter
---


When you write an application that uses any kind of web-API out there, you eventually get to a point where you hit some kind of rate-limit. It's especially a problem when the API has a frequently changing rate-limit as [Twitter](http://twitter.com)'s has. Well, at least on this front, there is now an easy way around :-)


-------------------------------

So far Twitter had with "account/rate_limit_status.${format}" a simple call to get the number of API calls you can do during the current period.

@@ javascript @@
{
    "remaining_hits":8
}
@@

The problem was, that this information was only useful to some degree because you didn't really know, when the next reset would happen. There was also no easy way to find out the "official" rate limit (say "20 hits per hour"). That led to a quite lengthy [discussion](http://groups.google.com/group/twitter-development-talk/browse_thread/thread/12b81ff1c0f92cc6/0e0d7457216505c4) on the Twitter developer mailing-list about what kind of information would be relevant to developers. Basically what everyone wanted was a way to find out, when that next reset was going to happen and in general how many calls you could make during the whole period. 

During the course of the discussion, Alex Payne [suggested a format](http://groups.google.com/group/twitter-development-talk/msg/f2a911737ba2ef69) that solves most of that and put it live on the 27th:

@@ javascript @@
{
    "remaining_hits":8,
    "hourly_limit":20,
    "reset_time":"Sun Jun 29 10:09:14 +0000 2008",
    "reset_time_in_seconds":1214734154
}
@@

That was the first step and a huge one. Big kudos to Alex for being so fast. This was really important since in the end I guess no one really knows anymore, what rate-limit currently applies.

Today, [another suggesting](http://groups.google.com/group/twitter-development-talk/msg/6f3691c9c0bbb820) came up, to just include that information in every response by the server to save some extra calls. In my opinion this would be a good candidate to become an option. I simply see enough apps that don't need this information since their rate limit is negligible, so the extra session lookup would be unnecessary. 