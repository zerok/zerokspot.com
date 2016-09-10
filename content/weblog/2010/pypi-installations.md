---
date: '2010-11-01T12:00:00-00:00'
language: en
tags:
- python
- pypi
- mirror
- proxy
- pygraz
title: Proxying/Mirroring PyPI and private installations
---


If you're working with external dependencies in Python there is probably
no way around [PyPI][], the Python Package Index. On some occasions it would
be kind of nice, though, to have a way around it, say, when PyPI is down,
your own connection to the world is down or you're working on some private
packages that shouldn't really end up on PyPI in the first place.

So we end up with two problems: (1) having a mirror/proxy setup that handled
PyPI fail-overs for us and (2) having a private package index.

A short disclaimer: I have tried these solutions so far only for a couple of
minutes but plan to use one of the described tools for my home server as well
as my netbook.  This post here is more or less a summary/short version of a
presentation I'm going to give at the next [PyGRAZ][] meetup. So if you plan to
attend: **Spoilers** :-)

---------------------------

First to the mirrors and proxies: Depending on your environment one might be
more appropriate than the other. If you don't want to waste all that hard-disk
space for packages you will most likely never need, a proxy will probably be
more for you, but if you can spare the space, you gain quite a bit of
flexibility (travelling by train, anyone?) with having everything available
locally.

Luckily there are already some solutions for either going mirror or proxy:

* [z3c.pypimirror][] creates a mirror of the PyPI in a given directory which can
  be updated. Right now it doesn't really remove files no longer existing on
  upstream, though.
* [pep381client][] is the "official" implementation of [PEP381][] which lays the
  groundwork for a mirror infrastructure for the PyPI. Compared to pypimirror,
  it also removes no longer present packages and projects.
* [collective.eggproxy][] provides a simple WSGI application where requests for
  packages are first checked in a local cache and then are forwarded to the
  primary PyPI with the result being cached again. So you end up with only
  those packages stored locally that you've requested through that proxy.

If you also need a private installation to host packages that should not end
up on the public index, you have a couple of options, too:

* [mypypi][]
* [haufe.eggserver][]
* [Plone Software Center][]

All of them are somehow based on Zope/Plone and provide their own server. I
have not tried any of them but just installing haufe.eggserver displayed
enough errors and warnings that I kind of fear that they won't really work
with >=Python 2.6 out of the box. Luckily there is also a lightweight
solution: The package index isn't all that complicated so you can easily
mirror its URLs using URL-rewriting in Apache or other servers as demonstrated
by [Reinout van Rees](http://reinout.vanrees.org/weblog/2009/11/09/eggproxy-plus-private-packages.html)
which - thanks to Apache's quite powerful rewrite engine - can also forward
requests directly to a proxy or mirror.

And if you're not a friend of rewrite rules and automatically generated
indexes, there is still [basketweaver][] ;-)

For me personally, going with an Apache/nginx solution that handles all my
private packages and them falling back to pep381client for my netbook sounds
like the best approach. I use my netbook mostly when going by bus or trains
somewhere and usually I don't have any internet connectivity worth mentioning
then (thanks to all these mountains here and Orange not really having that
great a coverage in these areas). For my workstation at home, replacing
pep381client with eggproxy sounds like a good idea, though :-)

And if everything falls apart, [pip](http://pypi.python.org/pypi/pip#mirror-support)
and [buildout](http://pypi.python.org/pypi/zc.buildout#finding-distributions) can
configured to use two indexes ;-)


[Plone Software Center]: http://plone.org/products/plonesoftwarecenter/
[z3c.pypimirror]: http://pypi.python.org/pypi/z3c.pypimirror
[haufe.eggserver]: http://pypi.python.org/pypi/haufe.eggserver
[collective.eggproxy]: http://pypi.python.org/pypi/collective.eggproxy
[pep381client]: http://pypi.python.org/pypi/pep381client
[mypypi]: http://pypi.python.org/pypi/mypypi
[basketweaver]: http://pypi.python.org/pypi/basketweaver
[PEP381]: http://www.python.org/dev/peps/pep-0381/
[PyGRAZ]: http://pygraz.org
[PyPI]: http://pypi.python.org
