---
date: '2008-09-29T12:00:00-00:00'
language: en
tags:
- django
title: Django 1.0 now with its own branch
---


Right after 1.0 got released, some people started to wonder how to easily stay up to date on bugfixes for this major release. Back then I think it was James Bennett who told me to stay tuned for something on this front. Today, [Jacob Kaplan-Moss announced](http://groups.google.com/group/django-developers/t/3767a05601a68448?hl=en) that Django 1.0 finally has [its own maintenance branch](http://code.djangoproject.com/svn/django/branches/releases/1.0.X) in the repository. 

-------------------------------

This means mostly one thing for you according the [release process documentation](http://docs.djangoproject.com/en/dev/internals/release-process/): If you want to stay up to date with the latest fixes for the Django 1.0 release without getting some unexpected and perhaps not really welcome changes to the core, all you have to do is::
    
    $ svn co http://code.djangoproject.com/svn/django/branches/releases/1.0.X/ django-1.0
    $ cd django-1.0
    $ svn up
    # The latest bugfix-release just went live
    $ svn up

This way it should be easy to keep using your previous update process (if you were running on trunk that is) while having a solid base for your sites. 
