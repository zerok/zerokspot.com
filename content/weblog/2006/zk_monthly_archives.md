---
date: '2006-08-01T12:00:00-00:00'
language: en
tags:
- drupal
- modules
- zerokspot
- zk_monthly_archives
title: zk_monthly_archives
---


A few days ago in #drupal c4se noticed a the archive modules I'm using here. Well, and since it's not yet released I was wondering, if it's really worth putting it out into the light of the world. Let's see ...

zk\_monthly\_archives is basically a replacement for the archives modules bundled with Drupal. It's actually a much simplified version of it that does 2 things:

* Provide a block holding a list of all the months published nodes exist in ...
* and offers for each month a listing off all these nodes.

Nothing really special but I liked it that way when I first saw something like this quite some time ago on [Jon Hick's weblog](http://www.hicksdesign.co.uk). Sure, the implementation is a little bit different, simply because I normally have more than 10 posts per month and the page would just get too long if I put a listing of _all_ posts on one single page :)

This was also my first module I've ever written for Drupal, so please don't kill me for its simplicity and probably also for the sloppy code in it. It gets the job done and that's all I care about. So there's no fancy admin/settings form for it ;)

If you find any bugs, please comment to this post for now. I think I will sooner or later install the projects module here, but first I want to test it to see if it really does what I want :)



-------------------------------



So here the obligatory installation guide:

1. Download the file \[zk\_monthly\_archives-4.7.0.tar.bz2\]
2. Extract it which should give you following directory: "zk\_monthly\_archives"
3. Put this folder into your modules directory (or where ever you normally put your contributed modules)
4. Activate the modules in admin/modules
5. Add the "Archives: List of Months" block to your sidebar or something similar.
6. Enjoy!