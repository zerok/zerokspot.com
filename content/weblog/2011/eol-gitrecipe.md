---
date: '2011-07-10T12:00:00-00:00'
language: en
tags:
- python
- buildout
title: 'PSA: End of Life of zerokspot.recipe.git'
---


Just a short notice to everyone who is using my git recipe for
[buildout](http://pypi.python.org/pypi/zc.buildout),
[zerokspot.recipe.git](http://pypi.python.org/pypi/zerokspot.recipe.git/):
Version 0.6 will be the last feature release and, if nothing really
significant come up, the last release of this package in general. The
[repository on github](https://github.com/zerok/zerokspot.gitrecipe) will stay
online for the time being but don't expect any real work being done on it
anymore. If you like what you see there, by all means, fork it :-)

As for the reasons for this: For virtually everything I use a combination of
virtualenv, pip and fabric now so I haven't used buildout for quite some time.
Because of that it makes little sense to keep maintaining something I don't
use myself anymore and luckily there are quite a few alternatives out for
handling software packages available only via git or other VCSs. There are for
instance [mr.developer](http://pypi.python.org/pypi/mr.developer) and
[gp.recipe.pip](http://pypi.python.org/pypi/gp.recipe.pip) just to name two of
them.

I also want to thank all the contributors this project had over the years.
Without you I'd probably have written this post probably a long time ago :-)
