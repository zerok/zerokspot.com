---
date: '2010-05-25T12:00:00-00:00'
language: en
tags:
- djangocon
- django
- berlin
- travel
title: Djangocon.eu - Day 2
---


Day two ... that means we all survived the evening of day one at [Clärchens
Ballhaus](http://ballhaus.de/). As the name indicates, this is not your
typical bar or restaurant but actually a dance hall with room for more than
100 people at tables and enough room for those who also want to dance. A big
thank you, again, to Jesper and [Bitbucket](http://bitbucket.org/) for
sponsoring the drinks.

--------------

## Keynotes

As every good day does, day two started with Keynotes. This time by CouchDB's
Jan Lehnardt who reminded us about one of the major corner stones of good
design - no matter if its graphical or technical: Simplicity

If you keep solutions simple, yet effective, you usually also produce code
with less bugs since the number of errors per Klines of codes is rather
constant (~ 10-25). And here it doesn't really matter what programming
language you use or what framework.

As an example for simplicity Jan mentioned the
[mustache](http://mustache.github.com/) template engine (for which he
maintains a [JavaScript port](http://github.com/janl/mustache.js) where it
showed that keeping out features actually produced a much cleaner solution.

## Gunicorn, Django and WSGI

Next was Benoît Chesneau who gave a presentation of WSGI, how it plays with Django and
how to deploy them using Gunicorn. The first part of the talk consisted of a
short summary on WSGI. Next he introduced us to two ways on how to mix Django
with "native" WSGI applications:
[twod.wsgi](http://packages.python.org/twod.wsgi/) and
[django-wsgi](http://github.com/alex/django-wsgi).

Having WSGI is not worth a thing if you can't deploy WSGI apps. So Benoît gave
an overview of some of the servers out there with support for Python's "web
protocol": uwsgi, which is an "extension" for NGINX; mod_wsgi for Apache HTTPD
and various "native" projects like cherrypy, spawning, paster and, naturally,
Gunicorn.

## Django CMS

[Patrick Lauber](http://ch.linkedin.com/pub/patrick-lauber/5/268/734) was next
with a presentation of [Django CMS](http://www.django-cms.org/) complete with
tons of feature presentation and demos.

I haven't used it yet but it definitely looks interesting esp. since there are
even people out there writing
[extensions](http://www.django-cms.org/en/extensions/) for it.

## No! Bad pony

[Dr Russell Keith-Magee](http://cecinestpasun.com/) gave an updated version of
his famous "No! Bad Pony!" talk. And here is the original:

<embed src="http://blip.tv/play/AYG6_AgC" type="application/x-shockwave-flash" width="480" height="350" allowscriptaccess="always" allowfullscreen="true"></embed>

## MongoDB

The first talk after lunch was by Peter Bengtsson and about how to use MongoDB
with Django. After a short crash course to MongoDB he went right into all the
options you have if you want to get more NoSQL into your django project with
solutions like [Ming](http://merciless.sourceforge.net/tour.html),
[mongoengine](http://hmarr.com/mongoengine/) or
[django-mongokit](http://github.com/peterbe/django-mongokit). django-mongokit
even offers integration with signals and test-databases , the latter I'm
really missing sometimes in mongoengine.

And I'm pleased that I seem to have taken the right way with mongoengine for
this site ;-)


## Django South

Next, Andrew Godwin presented what has changed within Django South over the
last two years and his plans for the future. Sadly, during his talk I was
mostly on the hunt for some free powerplug I will simply link to one of
Reinout's [excellent
summaries](http://reinout.vanrees.org/weblog/2010/05/25/south-new-and-old.html).

(Perhaps I should change to this format, too ;-))

## Django Technical Design Panel

The last official talk for today was a technical design panel with Jacob,
Russel and Jannis and Alex acting as moderator and presenting [community
question](http://www.google.com/moderator/#15/e=751d&t=751d.41&f=751d.2c4e9).
A short summary of what I could hear from the back of the room:

* Jacob would like the middleware API to be completely rewritten and renamed and URLconf to be improved.
* Russell wants some async support inspired by how node.js does it.
* The way project templates are handled in Rails is something Jacob would like to see in Django.
* The first goal regarding Django on Python 3.x is to get a working distribution out there that allows the community to provide usable bug reports.
* The way, Django apps are found by users could be improved and the ultimate goal should be to come up with a system that can be used by Django as well as by Python as a whole since the problems are mostly the same.
* The DVCS mirrors should get listed on djangoproject.com

## Lightning talks

* Continuous integration with [Hudson](http://hudson-ci.org/) by Erik Romijn
* [Lightning Fast Shop](http://www.getlfs.com)
* No love for generic views and flatpages from Stijn Debrouwere
* [Django-ROA](http://code.welldev.org/django-roa) by David Larlet
* I18n content by using the sites framework by Stefan Wehrmeyer
* An introduction to Flask by Armin Ronacher
* Deployment tips by Thilo Fromm
* An introduction to software licensing by  Jacob Kaplan-Moss
* A call for action for a project related to Tibetan Buddism by Tomas Juriga

OK, enough for today and on to hopefully more photos tomorrow :-)
