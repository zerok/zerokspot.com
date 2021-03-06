---
date: '2008-07-19T12:00:00-00:00'
language: en
tags:
- django
- forms
- newforms
- newformsadmin
title: newforms-admin in Django's trunk and {new,}forms
---


After nearly [more than a year of development](http://code.djangoproject.com/changeset/4314) Brian Rosner [merged](http://code.djangoproject.com/changeset/7967) the newforms-admin branch into trunk last night. This marks the last missing branch getting merged into trunk before the 1.0 release. Big kudos and thanks to Brian and everybody else involved in the development on this branch.

Another [big change](http://code.djangoproject.com/changeset/7971) happened last night: newforms is now finally the "official" forms library for Django. At least it now resides directly within django.forms.

Both of these changes are backwards-incompatible, yet fully documented ([1](http://code.djangoproject.com/wiki/BackwardsIncompatibleChanges#Mergednewforms-adminintotrunk), [2](http://code.djangoproject.com/wiki/BackwardsIncompatibleChanges#Movednewformstoforms)). For everyone running on trunk (and who hasn't been preparing for this merge) this means some work esp. if you've messed around quite a lot with oldadmin. Also some of the styling in NFA looks a little bit different than before thanks to the move away from the old form manipulator modules which added some special classes to each form-element. If it annoys you, check out [#5609](http://code.djangoproject.com/ticket/5609). 