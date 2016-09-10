---
date: '2009-01-30T12:00:00-00:00'
language: en
tags:
- django
- flatblocks
title: World, meet django-flatblocks 0.1.0
---


A couple of months ago I found out about Clint Ecker's excellent django app [django-chunks](http://code.google.com/p/django-chunks/), which basically does one thing and that very well: It takes the idea of a django.contrib.flatpages and uses it for small chunks or blocks on a page, like a help section a small "about" part you want to have on every page, yet still keep it editable. All you have to do, is create a model instance, give it a specific name/key/slug and then use a bundled templatetag to include that object into your view.

<pre>{% load chunks %}
...
{% chunk 'my_help_section' %}</pre>

Many people started forking it to add some additional fields. Kevin Fricovsky added [an active-flag](http://github.com/howiworkdaily/django-freetext/tree/master), Peter Baumgartner added [a header field](http://github.com/lincolnloop/django-freetext/tree/master), ... which is something I personally needed for one of my projects. But I also need a couple of other things, like an inclusion-tag instead of a plain-old-django-templatetag, so that I could easily add, for example, an edit button right next to each such block. That was the time when I knew, I had to make it a real fork. 

The first step actually was to also allow the name of the chunk/flatblock being passed via a template variable, so you could all of a sudden do something like this:

<pre>{% load flatblock_tags %}
...
{% flatblock blockInAVaribale %}
</pre>

... which makes things a little bit easier if you're, for instance, operating in a multi-lingual environment. But first I have some other things I want to see in django-flatblocks with a simple view for editing being quite on top of that list ... so that I don't have to write it again and again in each and every project I'm working on ;-)

But now enough of that. Enjoy [django-flatblocks 0.1.0](http://pypi.python.org/pypi/django-flatblocks/0.1.0) :-)

**Update (2009-02-25):** [django-flatblocks 0.2.0](http://zerokspot.com/weblog/2009/02/24/django-flatblocks-020/) is now available :-)