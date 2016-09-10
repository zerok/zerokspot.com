---
date: '2007-07-31T12:00:00-00:00'
language: en
tags:
- django
- django-lifestream
title: django-lifestream
---


Without the work of the whole Django community who provided pre-built applications for example for [tagging](http://django-tagging.googlecode.com) this whole site would have taken much longer to make than it actually did. So now it's my turn to give something back. Since I'm still quite new to Django I want to start with something small: The [Lifestream](/lifestream/) component of this page :-)

-------------------------------

As already said, this is probably the smallest component of this site, but it was quite easy to make it generally applicable. It basically consists of 2 components:

1.	A commandline script that can easily be integrated with a cron-system in order to periodically regenerate the lifestream (backend)
2.	and a frontend that consists of a view (and hopefully in the future also a feed) for actually rendering the built-stream.

The setup is pretty straight-forward if you have ever worked with cron so I hope the provided documentation is sufficient (at least for now). Otherwise just file a bug report :-)

For now the whole project is hosted on [Google Code](http://code.google.com) but I also plan to at least mirror it on [Launchpad](http://launchpad.net) since I prefer working with distributed version control systems.

So, by now you can find django-lifestream on **[http://code.google.com/p/django-lifestream/](http://code.google.com/p/django-lifestream/)**. I hope you like it and it's useful :-)

I once again want to thank [Jeff Croft](http://jeffcroft.com) and [Manuela Hoffmann](http://www.manuela-hoffmann.info/) for their inspiration and the whole Django community (and esp. all the folks in #django) for their help  :-)