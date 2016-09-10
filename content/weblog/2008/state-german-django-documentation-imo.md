---
date: '2008-11-20T12:00:00-00:00'
language: en
tags:
- django
- documentation
- german
- state
- translation
title: State of the German Django documentation (IMO)
---


OK, first of all: The [German translation of the Django documentation](http://www.django-de.org/participate/) is not dead. While the start late last year saw some great results, the whole movement got a little slowed down over the time and right now it's still just proceeding on a more slow way right now. When the [last major refactor](http://code.djangoproject.com/changeset/8506) in upstream took place (some time before [PyConUK](http://pyconuk.org/) 2008) I started looking through it and it became quite clear, that it would be quite hard to save some of our work from before. The structure and most of the content was just too different to save even a single document completely. So I started building a little bit of the foundation (style-port, main menu, etc.) during my train-trip back home.

-------------------------------

The plan was and still is, to first of all look at the documents and try to save parts from before, while not really shying away from just re-translating it altogether. Since we couldn't really reuse the old translations right away I started a new repository for it on [github](http://github.com/zerok/django-docs-de/tree/master) last month and I'm trying to work my way through the documentation one file at a time right now. The idea is (based on what I gathered [Jannis](http://jannisleidel.com/) wants (not that I'm of a different opinion in this regard)), to have a translation for each "real" version of Django. So one for Django 1.0, one for Django 1.1 and so on... if we could get the 1.0 version ready some time **before** the release of Django 1.1 this would be a small miracle though.

Naturally this process takes time and thanks to some time-constraints on my master thesis right now (something that is probably not really worth getting into due to its weirdness). Currently I think I set up 4 people from [#django-de](irc://freenode.net/django-de) on freenode as contributors who can commit to that repo right away. The true benefit of github is, though, that *anyone* can easily contribute to projects and get attributed. So if you want to help translating the docs to German and all you want is getting your name into the history of the repository for contributing back something useful, just fork the repo and have fun with a file :-) Currently there is no real infrastructure in place to keep people from working on the same files in parallel, so perhaps if you just drop me a line or something like that, it would be helpful... or we could agree on some hash-tag on twitter. Something like `#djangodocsde` or whatnot. For all this, please mind the big "IMO" in the title ;-)

Something that I absolutely want to get done this weekend is a restructuring of the index page which happened a couple of days ago [upstream](http://code.djangoproject.com/changeset/9490) and makes it way more useable. Anyway, there will definitely be at least some progress on the docs-front by the end of this weekend.

These are just some ideas and my impression of what the current state of the translation is, so if you have some better ideas, please let me know or just join #django-de so that more people can join the discussion :-)