---
date: '2007-10-12T12:00:00-00:00'
language: en
tags:
- development
- pylons
- python
- web
title: 'First taste of Pylons: Rocky'
---


For the last week or so I've been messing a little bit around with [Pylons](http://pylonshq.com/) during my limited free time in order to be of at least of some help for [Martin with his new project](http://mgratzer.topmind.at/2007/10/06/my-next-coding-project-photowalking-metadata/). I'm not really sure how I should feel about this framework, though.

-------------------------------

A few points on this:

* It's probably the same with everything: It takes a lot of time getting used to something different, so please bear with me here. I don't know all that much yet about Pylons, so these are just my first impressions.
* I like the idea of re-using components where possible. Being able to use SQLAlchemy or any other DB-layer or template engine wherever I feel like it, is definitely nice.
* ... but there should be some templates for paster bundled with Pylons that already do the basic configuration stuff for often-used components like SQLAlchemy.
* The debugging output is absolutely great. From what I've seen so far, Pylons creates a separate debugging page (with its own URL) whenever an exception stays unhandled, which should make debugging broken POST or AJAX calls way easier.
* The documentation is completely fragmented. You have (1) the [docs section](http://wiki.pylonshq.com/display/pylonsdocs/Home) on pylonshq.com, then you have the [WIKI](http://wiki.pylonshq.com/dashboard.action) (which is actually a superset of the documentation) and then you have the documentation for every component you want to use on its respective project site.
* Pylons has quite a low level feeling to it. While Django seems to be at the 6th floor above WSGI, Pylons seems to be on the first half-floor. Definitely one of the strongest points of Pylons in my opinion (but as I demonstrated in Vienna: I like explicit ;-)).
* I'm not yet sure, whether I like the deployment of projects or not. I'm simply not really a fan of setuptools.

Pylons so far looks really interesting, but given the lack of a more complete set of paster templates, it takes quite some time to get a project under way if you're using for example SQLAlchemy and other non-core components. 