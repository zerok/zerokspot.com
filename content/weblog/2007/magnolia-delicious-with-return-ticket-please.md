---
date: '2007-02-13T12:00:00-00:00'
language: en
tags:
- api
- del-icio-us
- ma-gnolia
- wishlist
title: Ma.gnolia del.icio.us with return ticket, please
---


Well, yesterday I (once again?) walked again from [del.icio.us](http://del.icio.us) and imported all my bookmarks into [Ma.gnolia](http://ma.gnolia.com) after all the [trouble](http://zerokspot.com/node/832) I had with the the del.icio.us API. The exact reasons should be quite obvious from my last post so I won't really elaborate on them here anymore. Just one thing: If you have rules for your service, really only temp. ban people from it when they obviously act against these rules.

-------------------------------

 There might have been a hidden abuse-trigger for preventing mass-deletions, but even adding update-queries in between the delete queries didn't help. 100 queries per hour was the maximum. (Btw.: Don't call it "throttling" if it's actually "banning". Throttling is for me like giving me less speed.)

Anyway, on to the world of Ma.gnolia (once again?). I haven't  tried their API yet, but at least I've found no exact rules that I could act against ;-). Ma.gnolia has changed quite a lot since I moved back to del.icio.us last May:

* Posting to groups can now be done right when you add a bookmark.
* There is now a discussion feature for each group.
* The servers got faster :)
* The auto-completion for tags in the bookmarklet feels much better to me now.

There are still some limitations though which I would like to formulate as feature requests and also explain why I would like to see them (but wouldn't be all that pissed if ... it takes longer ;-) ):

1. __Group support for the API:__ The API for selecting posts offers a neat    way to create a "daily posts" feature from the client side similar to what del.icio.us does on the server side. 
	* But what happens if you have actually a group you want to get those bookmarks from? In this sense it would be really nice to have an additional parameter for the bookmarks\_get function that lets you specify the name or id of a group you're member of.
	* As kind of a pre-requirement for this probably a function similar to the tag\_find function would be needed, that gives you all the groups (or group ids) you're a member of.
	* Perhaps also some functions to manipulate your group'd bookmarks.
2. "__How many people have bookmarked a link__" directly in the bookmark listing (IMO the del.icio.us implementation for this is quite good).
3. Is there perhaps a way to __kill the animation in the Roots__ bookmarklet? It's nearly killing my browser :-)

Esp. (1) would be really nice since I would like to integrate some daily links by the editors on [gamerslog.com](http://gamerslog.com) and using the group for this would be the ideal solution for me :-)

Once again: Thank you for this great service :-)