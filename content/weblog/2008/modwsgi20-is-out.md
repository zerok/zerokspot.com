---
date: '2008-03-21T12:00:00-00:00'
language: en
tags:
- python
title: mod_wsgi 2.0 is out
---


Over the last couple of months [mod_wsgi](http://code.google.com/p/modwsgi/) has become a very good alternative to mod\_python and FastCGI if you want to deploy a Python-based web application on Apache. For the last 3 months, this site has been running on it and it's been a great experience. Today, Graham Dumpleton [released](http://blog.dscpl.com.au/2008/03/version-20-of-modwsgi-is-now-available.html) version 2.0 which comes with tons of nice new features.

-------------------------------

For me personally, there are 2 new features that I've been really looking forward to (and I nearly installed some of the release candidates to get them):

You can now set a timeout on daemon processes to shut down when having been idling for a certain number of seconds. Since I want to save some memory and I have a couple of low-traffic sites that I want to host on the same slice as this one, this should help quite a lot.

You can now also name daemon processes according to the [docs](http://code.google.com/p/modwsgi/wiki/ChangesInVersion0200):

> Added 'display-name' option for WSGIDaemonProcess. On operating systems where it works, this should allow displayed name of daemon process shown by 'ps' to be changed. Note that name will be truncated to whatever the existing length of 'argv0' was for the process.

... which should come in pretty handy when trying to blame excessive memory usage on a certain application ;-)

In general, there are tons of new options for optimizing mod\_wsgi daemon processes which should help esp. those of us who are on a memory-limit and don't like getting angry calls by the Linux kernel :-)
