---
date: '2008-10-01T12:00:00-00:00'
language: en
tags:
- python
- django
title: Django-oopviews now in its own library
---


As you might have guessed based on the recent activity in django-zsutils, this project is more or less a dumping ground for some ideas I have and snippets I use on multiple sites (or at least see myself using on multiple sites). The new step for any of these is, whether I can see them as standalone library. If I do, I try to clean them up even more and split them of the main package. 

This happened yesterday to my little [object-oriented views-implementation for Django](http://github.com/zerok/django-oopviews/). If you don't know this library yet, please take a look at its [README](http://github.com/zerok/django-oopviews/tree/master/README.rst) or [this blog post](http://zerokspot.com/weblog/e/1037/). Basically, the idea is to be able to share some common functionality between views by making views out of Python classes.

Today I did some final cleanup, moved it over to setuptools and finally made an actual [0.2.0](http://pypi.python.org/pypi/django-oopviews/0.2.0) release including [PyPI-registration](http://pypi.python.org/pypi/django-oopviews/), files etc. So if you want to use it, just run `easy_install django-oopviews` :-) 

I also registered the project on [launchpad.net](https://launchpad.net/django-oopviews), so if you find any bugs or have some feature request or other questions, please ask there :-)
