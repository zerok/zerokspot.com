---
date: '2007-02-11T12:00:00-00:00'
language: en
tags:
- api
- del-icio-us
title: Temporarily banned from del.icio.us API?
---


Ok, yesterday I [wrote](http://zerokspot.com/node/831) about a story I found about strange Yahoo! customer care treatment. Today I wanted to delete a couple of posts from my del.icio.us account. Because it was quite a bunch of posts (622 or something like that) I wrote a small script to get this job done.


-------------------------------


For the first about 100 posts or so everything was fine, but then I only got a  503 error whenever I accessed the API. I checked the rules for the API usage again and again, but my script ran with a 2 seconds interval (+ some latency thanks to the slow connection ;) ). I guess I'm within the rules then:

<blockquote><p>Please wait AT LEAST ONE SECOND between queries, or you are likely to get automatically throttled. If you are releasing a library to access the API, you MUST do this.</p><cite><a href="http://del.icio.us/help/api/">del.icio.us</a></cite></blockquote>

I also noticed this:

<blockquote><p>Please watch for 503 errors and back-off appropriately. It means that you have been throttled.</p><cite><a href="http://del.icio.us/help/api/">del.icio.us</a></cite></blockquote>

Why do I get throttled when I'm clearly within the rules? I also changed my user agent since the FAQ mentioned something about default API user agents getting throttled regularly (although this is probably not the case here, since the throttling is clearly limited to my IP) after I got the first 50x errors.

It's been about one hour now since the last request went through without an error. Now I'm just curious how long I will get throttled for something that was - as far as I can tell - within the rules...

Is deleting more than 100 posts in a row perhaps also against the rules? Or does it perhaps trigger some abuse-alert?

__Update (1):__ Ok, it's been now 2 hours and the throttling has been lifted. Thanks :) But perhaps if someone knows why I was throttled in the first place, please let me know :)

__Update (2):__ Throttled again now with `sleep(random.randrange(4,8 ))`. This is really getting annoying.