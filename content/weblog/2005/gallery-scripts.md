---
date: '2005-03-25T12:00:00-00:00'
language: en
tags:
- development
- zerokspot
title: Gallery scripts
---


As you can see in the <a href="http://weblog.zerokspot.com/posts/289/">latest daily dump of del.icio.us bookmarks</a> I'm currently testing looking for a PHP gallery script which I want to use here. I've already tried quite a few but haven't really found the right one for me yet. <a href="http://gallery.menalto.com/">Gallery(2)</a> would be really close, but the problem is, that the server I want to install the script on is running in <em>safe_mode</em> :-( 

-------------------------------



So the next script I wanted to give a try was <a href="http://coppermine.sf.net">coppermine</a>, but it somehow felt a little bit to "feature-rich" for what I need.



What features do I want for my gallery?



* Load friendly (only generate pages if necessary and also already create the thumbnails when uploading the image or when requested by the admin

* Descriptions for images and albums

* It should work if safe_mode is enabled on the server

* RSS/RSS2/Atom feeds with thumbnails (only generated when an image is added or removed)

* Images should be .htaccess protected and only accessible through the script itself



Since I couldn't find a script that offers all that I started to write once again my own gallery. I hope that I will have something to upload here within a few days, right in time for holiday photos ;-)