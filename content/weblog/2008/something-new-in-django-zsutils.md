---
date: '2008-08-30T12:00:00-00:00'
language: en
tags:
- contentnegotiation
- django
- django-zsutils
- http
title: Something new in django_zsutils
---


The first big change in my ["little collection of Django-related stuff"](http://github.com/zerok/django-zsutils) is a new folder structure. This was necessary because for some modules I just want to have some sort of unit-testing going on, and with the old structure this would just have been a total mess.

The actual new part of the package is the [django_zsutils.utils.oopviews.ctn](http://github.com/zerok/django-zsutils/tree/master/django_zsutils/utils/oopviews/ctn.py) module which includes a sample implementation of [content-type negotiation](http://www.w3.org/Protocols/rfc2616/rfc2616-sec12.html). Stuff like that is always interesting if you want to provide rendering for different types of clients while still using the same URL. 

-------------------------------

Just a small example:

@@ python @@
class MyView(ctn.AbstractCTNView):
    ctn_accept_binding = {
        'text/html': 'html_handler',
        'text/*': 'plain_handler',
        '*/*': 'plain_handler',
    }

    def html_handler(self, request, *args, **kwargs):
        return HttpResponse('Hello World', mimetype='text/html')
    def plain_handler(self, request, *args, **kwargs):
        return HttpResponse('Hello World', mimetype='text/plain')
myview = oopviews.create_view(MyView)
@@

This view would accept basically every request for text and give only for text/html provides a different output. Please note, though, that this is by no means a complete implementation of content-type negotiation. The specs also mention additional type-parameters, that are no further described, so I couldn't really add support for them. But for most cases, this implementation should suffice :-)