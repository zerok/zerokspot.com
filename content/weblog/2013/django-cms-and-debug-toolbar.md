---
date: '2013-12-25T13:01:57+01:00'
language: en
tags:
- django
- django-debug-toolbar
- django-cms
title: Django-CMS and Debug Toolbar and syncdb
---


Recently, I ran across an issue executing syncdb on an empty database with
both [Django-CMS 3][2] and the [debug toolbar][3] listed in the installed
apps:

```
$ python manage.py syncdb --all
DatabaseError: relation "cms_title" does not exist
LINE 1: ..."django_site"."domain", "django_site"."name" FROM "cms_title...
                                                             ^
```

After some digging around with `--traceback` I discovered that this was 
caused by the debug toolbar (using the quick setup) patching the URL config 
when its `models.py` is loaded, which in turn triggers the plugin discovery 
to misbehave.

To get around this, follow the [explicit setup][1] as described in the 
toolbar's documentation which requires that you disable the automatic 
patching, add a URL pattern manually to your project's urlpatterns and the 
debug toolbar's middleware.

[1]: http://django-debug-toolbar.readthedocs.org/en/1.0/installation.html#explicit-setup
[2]: https://github.com/divio/django-cms
[3]: https://github.com/django-debug-toolbar/django-debug-toolbar