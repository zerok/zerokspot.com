---
date: '2008-01-12T12:00:00-00:00'
language: en
tags:
- django
- software-release
title: django-tagging 0.2
---


[Jonathan Buchanan](http://insin.webfactional.com/) just released version 0.2 of the django-tagging app which comes with quite some significant changes in how tags are detected. It's now also possible to have multi-word tags and as such tags no longer have to be separated by spaces but instead can now also be separated by commas.

Besides these and other new features as well as some bugfixes this update also comes with 2 backwards-incompatible changes:

1.  The database tables are now named `tagging_tag` instead of `tag` and `tagging_taggeditem` instead of `tagged_item`.
2.  The `tagging.utils.get_tag_name_list` function was removed. From what I can tell, `tagging.utils.parse_tag_input` is basically its replacement.
    
For more details check out [Jonathan's announcement](http://insin.webfactional.com/weblog/2008/jan/12/django-tagging-02-released/) and django-tagging's [project page](http://code.google.com/p/django-tagging/).
