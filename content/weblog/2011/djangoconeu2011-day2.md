---
date: '2011-06-12T12:00:00-00:00'
language: en
tags:
- django
- djangocon
- djangoconeu2011
- travel
title: 'DjangoCon Europe 2011: Day 2'
---


Note: Sorry for the long delay in DjangoCon Europe 2011 posts, but now I've finally
found some time again to put my notes and impressions to posts :-)

After some cocktails, more burgers and even more beer and far less sleep, the
second day had as much variety to offer as the first one with talks about
Celery, multilingual content, fighting the n+1-query-beast and much more.

---------------------

## Who cares about Zope?

<img
src="http://photos.h10n.me/Conferences/DjangoCon-Europe-2011/i-cfJrkzf/1/S/DSC0279-S.jpg"
alt="" class="left" />First thing in the morning [Martijn
Faassen](http://twitter.com/faassen) gave a keynote about what Django and
its community could learn from the problems and issues the Zope community had
over the years. The points that probably stuck the most with me was the huge
amount of circular dependencies Django's core has right now and also that
Django still has to work harder on integrating with the whole Python
community. It took quite some time for Django to once again embrace the whole
package and component mentality and right now some parts of it seem to be
stuck in 2006 when it comes to best-practices.

## Bitbucket - Lessons learnt

<figure>
    <img src="http://photos.h10n.me/Conferences/DjangoCon-Europe-2011/i-496Z2v6/1/M/DSC0284-M.jpg" alt="" />
</figure>

Next was [Jesper Noehr](http://twitter.com/jespern) with a keynote'ish kind of
talk about what [Bitbucket](http://bitbucket.org) learnt over the years. His
main lessons were:

* Stay idiomatic. Don't work against the core components and architecture of
  Django or you will most likely build in some roadblocks for when the next
  update comes around. It also makes introducing new co-workers to the project
  much easier since you can just give them the Django documentation to read.

* Things change and you should have a way to deal with these changes. Keep
  only a very small part of your architecture fixed (like using Django and
  building on top of Linux) and leave the rest as flexible as possible so that
  you can, for instance, move from one storage provider to another easily.

* Work with the opensource community by (1) using as many reusable components
  as possible and (2) also giving code back.

* Be transparent and humble. Bitbucket has a great history of communicating
  the reason behind outages and what is done to resolve them.

## Continuous integration and continuous deployment

<figure>
    <img src="http://photos.h10n.me/Conferences/DjangoCon-Europe-2011/i-kSDsTJh/1/M/DSC0293-M.jpg" alt="" />
</figure>

After a short break, [Szilveszter Farkas](http://szilveszterfarkas.com/) gave
a talk on how continuous integration, testing and deployment is done at
[Prezi](http://prezi.com). The core of their current development style is that
there are no longer feature branches but just one master branch where all
development happens. The usual counter argument here would be that this makes
implementing larger features problematic. The way Prezi is doing it, though,
is with feature flags using [Gargoyle](https://github.com/disqus/gargoyle).
This also makes performance testing simpler since you can emulate A and B by
just flipping some switches.

For internal stage deployment and continuous integration
[Jenkins](http://jenkins-ci.org/) is used.

## unjoinify

<figure>
<img alt="" src="http://photos.h10n.me/Conferences/DjangoCon-Europe-2011/i-f28h6Xk/1/M/DSC0296-M.jpg" />
</figure>

Next came a very low-level talk by [Matt Westcott](http://matt.west.co.tt/)
which dove into some of the problems you might face if you're working with
more complicated model structures, and especially the n+1 query problem for
nested data sets. In order to save you from going sown the raw SQL route he
wrote [unjoinify](https://github.com/gasman/django-unjoinify).

He also invited everyone to compare unjoinify with Simon Willison's
[django-queryset-transform](https://github.com/simonw/django-queryset-transform/).

## Celery

<figure>
<img src="http://photos.h10n.me/Conferences/DjangoCon-Europe-2011/i-PRSWwfp/0/M/DSC0315-M.jpg" alt="" />
</figure>

Right after lunch it was [Markus Zapke-Gründemann](http://www.keimlink.de/)'s
turn to give an introductory talk about [Celery](http://celeryproject.org/)
and how to handle asynchronous operations with it. He also gave some examples
on how Celery can be integrated easily with Django by using
[django-celery](http://packages.python.org/django-celery/).

## Reusable apps using "Eight Spaces"

<figure>
<img src="http://photos.h10n.me/Conferences/DjangoCon-Europe-2011/i-XQqCvMC/0/M/DSC0320-M.jpg" alt="" />
</figure>

Next in line was [Klaas van Schelven](http://twitter.com/vanschelven) and he introduced a set of patterns and
tools that are used at [Legalsense](http://www.legalsense.nl/) to make
customization of reusable applications easier. The core methodology here is
that views and models are dynamically generated within an application class so
that they can be replace without breaking other parts.

Some aspects of this kind of try to solve the same kind of problem as the
app-loading GSoC project by [Arthur Koziel](http://www.arthurkoziel.com/)
which is currently in its final stage of development.

## Core developers panel

<figure>
<img src="http://photos.h10n.me/Conferences/DjangoCon-Europe-2011/i-44d8xKF/0/M/DSC0338-M.jpg" alt="" />
</figure>

... with [Andrew Godwin](http://www.aeracode.org/), [Alex
Gaynor](http://alexgaynor.net/), [Jannis Leidel](http://twitter.com/jezdez),
[Idan Gazit](http://idan.gazit.me/) and [Russell
Keith-Magee](http://twitter.com/freakboy3742) talking about the past, present
and future of Django and answering tons of community questions. For me,
personally, the most interesting part was about what should perhaps get
incorporated into core. One of these ideas was multilingual content (which was
also the topic of the next talk).

Another big one was that there are now at least temptative plans for a
redesign of the admin with even the thought going around that it might not
necessarily be part of core (at least at first).

## The Django ORM and multilingual database content

<figure>
<img src="http://photos.h10n.me/Conferences/DjangoCon-Europe-2011/i-7kLK8Vt/0/M/DSC0190-M.jpg" alt="" />
</figure>

Basically continuing part of the discussion in the core developer panel was
[Jonas Obrist](http://twitter.com/ojiidotch)'s presentation about multiple
approaches to translate database fields and
[django-nani](https://github.com/ojii/django-nani) in particular.  He also
went into quite some details regarding the differences and pros and cons
compared to other approaches like having a single table for all translations
or using language-specific columns right next to the original columns into the
same table.

django-nani's data model basically places an additional table next to the
orignal model model which then holds the translatable fields and a
reference to the language as well as the "original" model that contains the
non-translatable field. So, in a sense, you have one database row per
translation. Personally, I also see this approach as the most promising one
and I know that it works since I've already worked with it in other systems.
Now, I just want to see it work with Django :D

## Integration the Enterprise using Django

<figure>
<img src="http://photos.h10n.me/Conferences/DjangoCon-Europe-2011/i-VRdDX49/0/M/DSC0362-M.jpg" alt="" />
</figure>

Sadly I completely missed everything of [Ed Crewe](http://www.edcrewe.com/)'s
talk on how to integrate different enterprise frameworks and systems using
Django.

## Lightning talks

* [Erik Romijn](http://blog.solidlinks.nl/) on the [Promodoro technique](http://www.pomodorotechnique.com/),
  an alternative time-/workload-management technique
* [Zachary Voase](http://zacharyvoase.com/) on
  [django-qmixin](https://github.com/zacharyvoase/django-qmixin), a
  mixin-toolkit for querysets and manager.
* [Angelo Dini](https://github.com/FinalAngel) presenting some general JavaScript best-practices
* [Jorge Bastida](http://jorgebastida.com/) on the [Dajaxproject](http://dajaxproject.com/), a framework
  for facilitating AJAX in Django
* [Ben Firshman](http://benfirshman.com/) on [Needle](https://github.com/bfirsh/needle), a tool for
  Selenium to integrate visual test cases.
* [Daniele Procida](http://twitter.com/evildmp) presenting [Arkestra](http://medicine.cf.ac.uk/arkestra/), a
  project he realised for Cardiff University out of frustration with their
  previous CMS.
* Idan Gazit giving some first hints at his new company
  [Skills](http://skillsapp.com/), which focuses on creating a new way for
  finding new employees for your company.
* and last but not least [Steve Holden](http://twitter.com/holdenweb) showing that even the smallest things
  can help make a difference.
