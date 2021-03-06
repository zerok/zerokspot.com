---
date: '2010-05-24T12:00:00-00:00'
language: en
tags:
- djangocon
- django
- berlin
- travel
title: Djangocon.eu - Day 1
---


<a class="left" href="http://www.flickr.com/photos/zerok/4640333866/" title="DSC_0101.jpg von zerok bei Flickr"><img src="http://farm5.static.flickr.com/4070/4640333866_fc95746c42_m.jpg" width="240" height="159" alt="DSC_0101.jpg" /></a>

So today was day 1 of this year's [Djangocon.eu](http://djangocon.eu) in
Berlin. After surviving last nights pre-show [beer-up](http://beerup.org/) at
the [Schleusenkrug](http://www.schleusenkrug.de/) getting a seat in the local
metro was easier than expected. I guess, on a holiday early in the morning not
all that many people are ready to leave their comfy beds.  The downside is
that you can't really ask anyone for the way, but luckily I ran into
[jezdez](http://jannisleidel.com) and [BartTC](http://mahner.org).

The venue is just fantastic with two large video walls, comfortable chairs and
power plugs for everyone. It's kind of well hidden, though, so without Jannis
and Martin I probably wouldn't have found it within 20 minutes (thank you,
telcos, for your stupid roaming contracts btw.). Part of being a great venue
is also that the beamer walls are high enough that if you have some &gt; 1.6m
sitting in front of you, you can still see the slides.

----------------------------------


## Keynotes

[Jacob Kaplan-Moss](http://jacobian.org/) talked about the state of the pony
and the community. He gave some really nice numbers regarding the user-base
with around 17.000 members currently registered on
[django-users](http://groups.google.com/group/django-users). But the main part
was not about the past but about the future: Where should Django go from here?

Similar to Rails, Django is currently or will be facing a couple of issues:

* Python is currently in flux with the transition to 3.x

* The rise of microframeworks like [Flask](http://flask.pocoo.org/) and friends

* A fracture of the community (less so in Django but still relevant)
  because of the size of the user-base out-pacing the size of the
  developer community and also because of the existence of alternatives.

* And the relationship between Django and the language community

First of all the interaction with the whole Python community should probably
be improved and also to go about standards more pro-actively by not only
implementing them but working standardization processes. Django should also
become more accessible to new users. Regarding this, naturally, the topic of
DVCS came up, but according to Jacob this would be the wrong way of addressing
the community issue and it's probably still too hot a topic right now anyway
;-).

A more detailed summary can be found on [Reinout's blog](http://reinout.vanrees.org/weblog/2010/05/24/jacob-keynote.html).

## WSGI and Python3

It's weird: Every time I meet [Armin](http://lucumr.pocoo.org/) at a Python
conference he talks about WSGI :-) This time it was all about the brokeness,
Python3's unicode handling (or string handling in general compared to how it
was done with Python 2.x) introduced with regards to WSGI and the web-related
parts of the stdlib. In general it seems like that Unicode was introduced into
parts where it doesn't really belong, like sys.stdin and urllib.

Because of that and in order to eventually upgrade to Python 3.x (since 2.7
will be the last release of the 2.x series of Python) Armin made a couple of
predictions. For example:

* The stdlib will probably be less used by WSGI apps

* Frameworks will have to re-implement certain functionalities that were
  usually provided by the stdlib.

So if we care for web development with and in Python, all of us should get
involved in the process of improving Python 3.x.

## Django and NoSQL

NoSQL is currently the big thing and so it was unavoidable to also see some
talks related to that here. [Alex Gaynor](http://alexgaynor.net/) is currently
working as part of his GSoC project on getting some kind of integration
between the model layer and various NoSQL implementation. The motivation for
that is pretty clear: Django offers a hug set of functionality regarding the
data layer and so far working with, e.g. MongoDB, requires that you have to
re-invent the wheel or jump through more hoops than necessary to use those
with alternatives data backends.

As part of his talk Alex also gave a really nice overview of the current
architecture of the query layer within the ORM, which he proposes to change a
little bit in order to move more platform dependent logic out of the Query
object and into the query compiler. A rough time frame was also presented: ~
Django 1.4.

## Running a User Group

For quite some time now I've been thinking about organizing some kind of
webdevelopment-centric user group in Graz. Webmontage and Barcamps are nice
but usually are less technically than I'd like them to be. [Sean
O'Connor](http://www.seanoc.com/) organizes the [Django NYC
usergroup](http://www.djangonyc.org/) and had some principles to share
regarding how to get a usergroup going and how to keep it running. I think it
kind of boils down to offering participants a scheduled and reliable place to
share ideas. This requires a venue that allows socializing and it has to be
promoted well. Sponsors, if they exist, shouldn't provide money but e.g. the
venue or other infrastructure.

The Django project itself could help here, in my opinion, with offering some
centralized place for people to find usergroups.

## Testing

[Honza Král](http://www.honzakral.com/) gave an overview about testing in
general just in case there were still people in the audience that have never
used the testing framework integrated in Django.

He went about some tips on how to write your tests in order to keep them
usable and also how to make your code in general more testable. For instance,
he suggests using class-based views since you then can easily fragment your
view into parts that you can easily test.

Quote of the talk:

<blockquote>
<p>Don't use doctests</p>
<cite>-- Armin Ronacher</cite>
</blockquote>

## Django and the Enterprise

Next, Jirka Schäfer of tschitschereengreen presented some tips on how you
might be able to introduce Django into a customers IT infrastructure. This is
rather problematic thanks to larger companies tried to keep their
infrastructure standardized and naturally larger companies being highly
political systems.

Basically the idea here is to use [Jython](http://jython.org/) to integrate
with the mostly Java-based infrastructure, acting kind of as a trojan horse
with something sweet in it ;-)

## Free the developers

[Will Hardy](http://willhardy.com.au/) presented a different take on reusable
apps he calls "domain specific frameworks" which are from what I understood
more loosely bound to the rest of the project using functionality like generic
relations and signals or provide for instance really generic base classes for
the model layer. Other means for such frameworks are template method that work
depending on the methods the models provide.

## Lightning talks

* What are [OpenSpaces](http://en.wikipedia.org/wiki/Open_Space_Technology)?

* Alex Gaynor about how to make template tags suck less:
  [django-templatetag-sugar](http://github.com/alex/django-templatetag-sugar)

* [LFC](http://bitbucket.org/diefenbach/django-lfc/), CMS, Release 1.0

* [RVirtualEnv](http://github.com/kvbik/rvirtualenv)

* Getting around Berlin

OK, that was the first day. Well, not really, but I think I will just write
about the [beerbucket](http://bit.ly/beerbucket) tomorrow. Thanks to
[bitbucket](http://bitbucket.org/) for sponsoring the beer :-)
