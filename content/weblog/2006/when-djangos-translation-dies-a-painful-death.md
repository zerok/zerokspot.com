---
date: '2006-12-06T12:00:00-00:00'
language: en
tags:
- django
title: When Django's translation() dies a painful death
---


Finally something about coding here ;-) I've just started working on a little project for a friend and I guess in the days to come I will probably write a little bit about things I learned about [Django](http://www.djangoproject.com/) and also describe fixes for problems I faced.

So on to the first problem: If you ever get something like this with Django:

<pre class="output">
&quot;&lt;PREFIX&gt;/django/utils/translation/trans_real.py&quot;, line 167, in _fetch
    app = getattr(__import__(appname[:p], {}, {}, [appname[p+1:]]), appname[p+1:])
AttributeError: &apos;module&apos; object has no attribute &apos;&lt;APP_IN_PROJ&gt;&apos;

-------------------------------


</pre>

this might be caused by circular imports within your INSTALLED_APPS. To make this more clear here a small example:

I have app proj.base and another app proj.menus. Now I got the great idea of writing some utility functions to make my life much easier. And to make it even more easier I put everything right into the \_\_init\_\_.py of proj.base. 

And the problem is now what exactly? Well, on of my utility functions uses a model from proj.menus. So I imported this app's models right in proj.base's \_\_init\_\_.py. For some reason now, this (or at least something in this direction) causes severe trouble with Django's translation subsystem. Sorry for that incomplete example, but I somehow couldn't completely reproduce this problem with a blank project.

So now I simply put everything I previously had in base's \_\_init\_\_.py into a utils.py in the same module, and finally everything works :)