---
date: '2016-06-26T12:02:12+02:00'
language: en
tags:
- redis
- gis
title: Fun with Redis' Geo Features
---

A couple of days ago I started implementing a little project where I needed
quick (not necessarily precise) GEO queries. Nothing more fancy than "Give me
the closest POI to my current location". The number of data-sets was extremely
limited so going with something like MongoDB, PostGIS, or ElasticSearch smelt
like overkill. I also had heard that [Redis][] had gotten basic GEO indexing
support in [3.2][320] and so I thought I should give it a try.

Turns out, this is the easiest way to handle simple geo indexing that I've ever
seen (at least so far). Let's work on a small example here: What's the closest
Austrian state capital from where you're currently at? First, we need to index
the positions of every state capital:

```
DEL capitals
GEOADD capitals 15.3717499 47.0735683 graz
GEOADD capitals 14.242715 46.6413036 klagenfurt
GEOADD capitals 16.4643351 47.8388337 eisenstadt
GEOADD capitals 16.2399758 48.2205998 vienna
GEOADD capitals 15.5769753 48.1937587 stpoelten
GEOADD capitals 14.1873214 48.2949799 linz
GEOADD capitals 12.9863902 47.8027886 salzburg
GEOADD capitals 11.3087501 47.2854337 innsbruck
GEOADD capitals 9.7169709 47.5070747 bregenz
```

Using the [GEOADD][] command, you don't add coordinates to global keys but
instead have one global repository (using sorted sets) where each contained key
has coordinates attached to it.

Now that we have all these positions stored, let's assume we are skiing on the
[Planai][], which is located more or less in the centre of Austria near the
border between Styria and Salzburg. How far would each capital city be from
here?

```
127.0.0.1:6379> GEORADIUS capitals 13.7172856 47.3690183 1000 km WITHDIST ASC
1) 1) "salzburg"
   2) "73.0357"
2) 1) "klagenfurt"
   2) "90.2198"
3) 1) "linz"
   2) "108.8062"
4) 1) "graz"
   2) "129.2256"
5) 1) "stpoelten"
   2) "166.5275"
6) 1) "innsbruck"
   2) "181.8110"
7) 1) "vienna"
   2) "210.9414"
8) 1) "eisenstadt"
   2) "212.5248"
9) 1) "bregenz"
   2) "301.3099"
```

That's where the [GEORADIUS][] command comes in. We once again operate on that
single store ("capitals") and ask it for all the keys within a 1000km radius
(just to make sure that we really get all the capitals back) around the
Planai. We also want to see how far it would be to each city (hence `WITHDIST`)
and the result should be shown in ASCending order. Turns out that if you had a
plane, you'd first reach 3 other capital cities before reaching the Styrian one
(Graz, the Planai is located in Styria).

Having to specify a radius is the only aspect I don't like about
this. Especially for small sets like ours it would really be nice if that option
weren't mandatory.

Now, let's say we are at a place that is already stored in that "capitals"
index. How far would it be to the other cities?

```
127.0.0.1:6379> GEORADIUSBYMEMBER capitals graz 1000 km WITHDIST ASC
1) 1) "graz"
   2) "0.0000"
2) 1) "klagenfurt"
   2) "98.4143"
3) 1) "eisenstadt"
   2) "118.3056"
4) 1) "stpoelten"
   2) "125.5406"
5) 1) "vienna"
   2) "143.2082"
6) 1) "linz"
   2) "162.2352"
7) 1) "salzburg"
   2) "196.9217"
8) 1) "innsbruck"
   2) "308.0339"
9) 1) "bregenz"
   2) "429.2321"
```

If you're interested in the distance between just two entries of that mapping,
Redis has you covered, again:

```
127.0.0.1:6379> GEODIST capitals graz klagenfurt km
"98.4143"
```

I have to say, this is a really nice feature that Redis got here. Esp. if you're
operating on smaller data-sets where you don't necessarily need something like
MongoDB, PostGIS, or ElasticSearch, this might be a killer-feature. Personally,
I definitely see myself using it quite a bit in the future 😊 Big thanks to
everyone who has worked on it!

[planai]: https://en.wikipedia.org/wiki/Planai
[redis]: https://redis.io
[320]: http://antirez.com/news/104
[geoadd]: http://redis.io/commands/geoadd
[georadius]: http://redis.io/commands/georadius
