---
date: '2010-05-26T12:00:00-00:00'
language: en
tags:
- djangocon
- berlin
- django
- travel
title: Djangocon.eu - Day 3
---


To get the usual recap from last night out of the way: Last night was no
organized social event, so social drinking games and no way too motivated
people dancing. Instead it was an evening full of small groups exploring the
city on a quest to find some beer and some food. Armin, Arne, Arthur, some guy
who isn't on Twitter and me went down Warschauer Straße and found some nice (I
guess) Spanish restaurant called
[Cayetano](http://www.qype.com/place/37617-CAYETANO-Berlin) with some comfy
tables outdoors. I'm always amazed how much a mess Nachos with cheese and
chicken can become when I'm near them.

So, now that this is out of the way, let's get to what happened at
Djangocon.eu on day 3.

-------------------------

## Relax with CouchDB

And on to more NoSQL with Benoît Chesneau giving an introduction to CouchDB,
what its all about and how you can use it easily within your Django project.
He also mentioned some of the "more recent" development around Couch like
[couchdb-lounge](http://tilgovi.github.com/couchdb-lounge/) for clustering and
the new geo indexer [GeoCouch](http://github.com/vmx/couchdb).

I also wasn't aware that couchdbkit was that tightly integrated with Django,
offering for instance a [counterpart to
ModelForms](http://github.com/benoitc/couchdbkit/blob/master/couchdbkit/ext/django/forms.py).
Plans for the next version of couchdbkit go even further with support for the
admin app and MultiDB.

## Release management with capistrano and supervisord

[Maciej Pasternacki](http://www.pasternacki.net/) of
[SetJam.com](http://www.setjam.com/) described how they do release management
using Capistrano and Supervisord. Capistrano was chosen because of some nice
support for transactional deployment and simply because it is more mature than
the competition. Dependency management with upstream libraries is handled not
with the usual suspects like pip, buildout etc. but with custom Makefiles
[which I personally thought to be a bit weird]. For scaling he recommended
[django-sqs](http://github.com/mpasternacki/django-sqs) and
[capistrano-ec2group](http://github.com/logandk/capistrano-ec2group) to
deployment to EC2.

## Design for Developers

Next was Idan Gazit's attempt on teaching software developers some design tricks to make our websites suck less visually. The goal here were minimalistic designs since they provide less opportunity to suck. Just to sum up a few tips:

* Make it clear for people what the actual content of your page is. You can do that by testing it yourself by squinting. If you can still detect the content area right away, you've succeeded.
* Use enough white space
* Provide enough contrast
* Design with a grid in mind
* Text content should not be wider than 2 alphabets
* Don't use more than 2 or 3 typefaces.
* Regarding colours you should know your audience since some colours can have unexpected meanings depending on the user's culture

## Efficient Django Hosting for the Masses

Michael P. Jung of [PyRox](http://pyrox.eu/webhosting/) gave an overview on
how they try to provide a shared-hosting infrastructure for people who don't
want or need a VPS or root server. From what I understood there each user gets
her own Apache server instance with mod_wsgi and nginx as server in front of
all those Apache instances. They also did some filesystem as well as database
replication.

## NoSQL panel

Do you want to support special NoSQL feature within the ORM even if they cannot be mimicked in SQL?

> No

How many NoSQL databases should Django support if any?

> Ideally all of them but there should perhaps first happen some kind of 
> standardization with regards to queries.

How much abstraction (should it emulate joins)?

> There are too few commonalities to really go beyond something like 
> get_by_pk or querying simple lists.

What are the approaches by other high-level frameworks?

> AFAWK it is not on their agenda right now.

This should not discourage people from using NoSQL with Django *right now* .


## Red Square

Red Square is an internal "social networking" application at
[BMW](http://www.bmw.de/) being in development since 2007. [Jörg
Kress](http://jjkress.tumblr.com/) presented the reasons for it and how
basically the upper management was involved in getting this project going.

## Testing

Eric Holscher had the honor of giving the last talk at Djangocon.eu 2010 where
he went into detail on testing infrastructures with great stuff like
[Hudson](http://hudson-ci.org/) and [devmason](http://devmason.com/) and some
tools in general like
[django-kong](http://github.com/ericholscher/django-kong) and
[patu](http://github.com/akrito/patu/commits/master). 

He also has a rather big Pony: There should be only one way on how to run
tests on newly downloaded Django apps. He proposed `python setup.py test`
since this has a good chance to end up in Python's stdlib when distutils2 and
related efforts land.

## Lightning talks

* P2P Web Applications with [CouchDB](http://couchdb.org) by Jan Lehnardt 
* Logic-less templates with {{ [mustaches](http://mustache.github.com/) }} by Jan Lehnardt
* Ray's Widget Exchange by Russell Keith-Magee on his vision about seeing a widget repository implemented
* Patches Welcome! by Ville Säävuori
* Demoserver by Remco Wendt 
* Localized documentation by me
* Continuous Performance Testing by Lukasz Dobrzanski
* Hidden Hires — Jacob Kaplan-Moss, Eric Holscher, Idan Gazit on a new job platform for Django companies and developers
* Mockity mock mock -- a little love for the mock library by Konrad Delong
* Easy package releasing with zest.releaser by Reinout van Rees
* <s>A Quick Look at the django-cms by Jonas Obrist</s>
* Surprise act (acapella of Code Monkey) by Remco Wendt and friends

