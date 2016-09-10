---
date: '2008-04-27T12:00:00-00:00'
language: en
tags:
- django
- qsrf
title: Django's ORM-layer revamped
---


[Malcolm Tredinnick](http://www.pointy-stick.com/blog/) today merged his [queryset-refactor branch](http://code.djangoproject.com/wiki/QuerysetRefactorBranch) into Django's trunk, which means everyone using trunk will finally be able to for example using a more sane syntax for ordering resultsets (no table names anymore), update multiple objects with just one query, tell select_related() which fields it should follow (instead of every foreign key it comes across) and much much more.

-------------------------------

Part of this branch is also an implementation for [model inheritance](http://www.djangoproject.com/documentation/model-api/#model-inheritance) which finally adds this long awaited feature to Django.

While the public interface for the ORM-layer should stay the same, the internals have undergone massive change. So if you're relying on some private part of the API, you will have to check if it still works. Affected of this if for example the popular [django-tagging]() app (guess what it does ;-)) which in its current state ([r132](http://code.google.com/p/django-tagging/source/detail?r=132)) uses the no longer available django.db.models.query.parse_lookup function. A patch by Joes Watts is available [here](http://code.google.com/p/django-tagging/issues/detail?id=106#c5) and seems to solve this issue. For more details check out [this ticket](http://code.google.com/p/django-tagging/issues/detail?id=106) on django-tagging's google code project page.

Anyway, congratulations and big thanks to Malcolm Tredinnick for his work on this. I can't wait for the next branch to get merged with trunk ;-)