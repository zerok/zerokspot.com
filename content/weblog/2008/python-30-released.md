---
date: '2008-12-04T12:00:00-00:00'
language: en
tags:
- python
title: Python 3.0 released
---


When I read yesterday night that [Python 3.0 was imminent](http://mail.python.org/pipermail/python-committers/2008-December/000277.html), I looked in awe at my calendar and thought "Is it already that late in the year"? This morning I then woke up, checked the newsfeeds and realized: [Yes](http://python.org/download/releases/3.0/). Naturally this doesn't mean that everyone will just leave 2.x behind and move over to 3.x right away but this public release hopefully makes it way more attractive to people to finally look into it then was the case with all these preview-builds ;-) And there's much to look into, indeed: 
    
* A whole new way to handle strings (no longer do you distinguish between
  unicode strings and normal strings)

* ``print`` is now a function and no longer a statement (which ended up
  in quite an ugly construct if you used it for printing to a specific
  IO-object)

* There is now only one integer type anymore. So ``long`` got dropped (and
  ``int`` is the new ``long``)

* sets and dict now also have their own \*-comprehension shortcuts
    
and [much much more](http://docs.python.org/dev/3.0/whatsnew/3.0.html). There are also some syntax changes, but from what I've seen so far Python stays Python :D 

I'm not going to dive into it right away, but Graham Dumpleton also [just wrote](http://blog.dscpl.com.au/2008/12/python-30-and-modwsgi.html) that mod_wsgi should already work with Python 3.0 with some tweaks if you're using trunk.