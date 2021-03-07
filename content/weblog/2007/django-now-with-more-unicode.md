---
date: '2007-07-05T12:00:00-00:00'
language: en
tags:
- django
title: Django now with more Unicode
---


Yesterday Malcolm Tredinnick merged the so called unicode-branch of Django into the main branch and with this made all Unicode goodness available to those people sticking to the primary development tree of Django.



-------------------------------



What does this mean in detail?

<blockquote><p>Django natively supports Unicode data everywhere. Providing your database can somehow store the data, you can safely pass around Unicode strings to templates, models and the database.</p><cite><a href="http://code.djangoproject.com/browser/django/branches/unicode/docs/unicode.txt?rev=5597">Django Unicode documentation</a></cite></blockquote>

Besides that Django now also has (at least some) support for IRIs (the international relatives of URIs ;-) )

Just to name a few changes. This is a huuuuuge merge and it comes with tons of changes that are all documented in the [docs/unicode.txt](http://www.djangoproject.com/documentation/unicode/) of your Django checkout so better read on there :-) And if you want to migrate an existing app to support Unicode, [this checklist](http://code.djangoproject.com/wiki/UnicodeBranch#PortingApplicationsTheQuickChecklist) should help.

For me personally this also fixed quite a few problems I had thanks to stupid charset clashes that happened while I was jumping from blogging software to blogging software and now wanted to write an importer for this broken data into my Django-based CMS/blogging tool. I can't wait until I can put zerokspot v5 online ;-)
