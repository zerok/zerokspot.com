---
date: '2009-05-06T12:00:00-00:00'
language: en
link: http://pypi.python.org/pypi/MercurialRecipe/
tags:
- buildout
- mercurial
title: MercurialRecipe
url_title: 'Python Package Index : MercurialRecipe 0.1.2'
---


A couple of months ago I wrote a small [zc.buildout][]-recipe to allow the
integration of data from some git repository into your workspace. Now I needed
mostly the same functionality but for a Mercurial repository. Luckily, [Tim Molendijk][]
created a small recipe just for that available in the [PyPI][]. Since the
README isn't really included there, you probably also want to take a look at
the project's [repository][] on Bitbucket.

To use it, simply configure a part like this::
    
    [django.piston]
    recipe = mercurialrecipe
    repository = http://bitbucket.org/jespern/django-piston/

As with my git-recipe it supports the ``newest`` option (globally and locally)
to prevent the repository from getting pulled every time you update the
environment and it sets the ``location``-part-variable so that you can access
the data you just pulled.

Really nice stuff. Thanks Tim.


[Tim Molendijk]: http://timmolendijk.nl/
[repository]: http://bitbucket.org/tawm/mercurial-recipe/
[pypi]: http://pypi.python.org/pypi/MercurialRecipe/
[zc.buildout]: http://pypi.python.org/pypi/zc.buildout