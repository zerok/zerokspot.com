---
date: '2005-03-12T12:00:00-00:00'
language: en
tags:
- web
title: Trying Feedmarker
---


As I've written last night, I will try <a href="http://www.feedmarker.com">Feedmarker</a> today, which looks like a combination of bloglines and del.icio.us.

-------------------------------



<h4>Registration</h4>
The registration is completely easy. First go to the login area where you should find the "register" link. All you have to fill in are your e-mail address, username and password (2x). The account is directly activated, so no e-mail activation or something like that.

<h4>First steps</h4>
Ok, services like this work because people are adding links, right? So let's find a link to add ;-) So I went to the <a href="http://blog.feedmarker.com/">Feedmarker blog</a> and tried to add it via the bookmarklet for feeds (which you can find in the menu of your admin section). As you can see there is a switch for making a feed public (or private ;-) ) similiar to bloglines and as in del.icio.us you can add a description and tags to a feed (ok, in del.icio.us it's a bookmark, but anyway). 

<img src="http://www.zerokspot.com/uploads/feedmarker-add.jpg" alt="Add feed interface"/>

Since I haven't specified the URL of an RSS or Atom file Feedmarker now tried to find such a file linked in the original URL. If there isn't a feed linked in the original URL you get an error message and a hint that you should perhaps try to add this URL as a bookmark instead of a feed.

<h4>Management</h4>
You can easily manage your feeds and bookmarks in an interface that resembles del.icio.us quite a lot (which is IMO not a bad thing since the del.icio.us interface gets the job done very well :-) ).  The cool thing here is, that you can really easy mark entries as private edit the feed or whatever without a single reload of the site. All is done java scripts. All the options you see in the screenshot below become visible when you hover the feed:

<img src="http://www.zerokspot.com/uploads/feedmarker-edit.jpg" alt="Edit options of a feed"/>

Every feed also has a collapse/de-collapse button which you can use to get a list of all the posts in a feed (or hide this list) and you can simply hit the "copy" button to copy a post to your bookmarks. Again: All this is done via JavaScript, so again no reloading.

<h4>Syndication</h4>
After taking a quick look at some subsites it seems like you can get a feed for every tag as well as for your whole repository of feeds and bookmarks. The feed ... about the feeds you've subscribed to also includes all the posts from every single feed. Really cool :-) (I don't know how many feeds and posts you get when you sync one of these feeds since I don't have enough items in my repositories yet ;-) )

<h4>Social aspects</h4>
IMO one of the strengths of del.icio.us is the option to see who else has a specific bookmark in his/her repository... and this is the only thing that I can really critizise about Feedmarker. There is no option to see who else has this bookmark/feed. All you can do is go to the index, get a list of recent (public) additions (which also includes your own - also private - additons) and copy some if you want. You also see who has added these entries but nothing more. Sure, you can view the repositories of every user (http://www.feedmarker.com/user/${username}) but nothing more. There are also only 10 entries per page on the index (yes, per page since you can browser to "older" pages) which should perhaps save some traffic...

<h4>The site itself</h4>
... about the traffic.... Just two points:

* All CSS definitons are inline (no &lt;link/&gt;s here...)
* At least parts of the JavaScript are also directly in the output

These are nearly 1200 lines of code that could be saved with every hit.

The site also doesn't validate against HTML 4.01 mostly because of undefined characters (I'd guess from the tags) which look a little bit hebrew to me. But there are also some validation problems with other things.

<h4>Anything else?</h4>
Yes. You can also import bookmarks (and I guess also feeds) via OPML which I haven't tried yet, and perhaps never will ;-)

<h4>Conclusion</h4>
The site looks really nice and is IMO quite promising. There are still some things that could be improved (like the inline CSS, inline JavaScript and parts of the social aspects) but apart from that you get a very clean (I'm talking about the interface ;-) ) system for managing feeds and bookmarks for free.