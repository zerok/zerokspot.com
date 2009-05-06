MercurialRecipe
###############

A couple of months ago I wrote a small `zc.buildout`_-recipe to allow the
integration of data from some git repository into your workspace. Now I needed
mostly the same functionality but for a Mercurial repository. Luckily, `Tim Molendijk`_
created a small recipe just for that available in the `PyPI`_. Since the
README isn't really included there, you probably also want to take a look at
the project's repository_ on Bitbucket.

To use it, simply configure a part like this::
    
    [django.piston]
    recipe = mercurialrecipe
    repository = http://bitbucket.org/jespern/django-piston/

As with my git-recipe it supports the ``newest`` option (globally and locally)
to prevent the repository from getting pulled every time you update the
environment and it sets the ``location``-part-variable so that you can access
the data you just pulled.

Really nice stuff. Thanks Tim.


.. _Tim Molendijk: http://timmolendijk.nl/
.. _repository: http://bitbucket.org/tawm/mercurial-recipe/
.. _pypi: http://pypi.python.org/pypi/MercurialRecipe/
.. _zc.buildout: http://pypi.python.org/pypi/zc.buildout