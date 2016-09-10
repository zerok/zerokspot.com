---
date: '2006-09-26T12:00:00-00:00'
language: en
tags:
- development
- macosx
- py2app
- python
title: App'ing with py2app
---


Today I wanted to write a small little tool that I could stick into my MacOSX dock and simply drag'n drop stuff onto it. Since the whole processing in the background is already done using some Python libs I wanted to see how I could get a simply Python script into that dock. The problem here is, that MacOSX has two different areas in the dock:

* One for handling applications
* and one for files and folder (like the trash bin)

... and MacOSX recognizes scripts as the later one. Everything following in the first category seems to be a really .app.



-------------------------------



And to convert a simple script into an App, there is a nice tool called [py2app](http://undefined.org/python/py2app.html). After installing it simply run `py2applet Script.py` and you will get a nice little Script.app.

You can also tell py2applet to simply put a symlink to your script into the resulting App (-A) which might make developing just a little bit easier. 

Another cool thing about it: If you drag for example a file from Finder onto that App in the dock, it will show up in sys.argv :D