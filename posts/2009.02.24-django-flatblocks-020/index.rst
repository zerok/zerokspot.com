django-flatblocks 0.2.0
#######################

While working on a small site for my family where I'm using flatblocks (and
which originally actually led me to create that project to begin with) I also
needed a simple view for editing a flatblock. Since this is something I can
absolutely see my self needing on every single project I use flatblocks, I've
now also included that view in django-flatblocks.

The other change included in version 0.2.0 is that you can finally localize
the fieldnames of models. So far the only translation included is the German
one but I definitely welcome patches with more translations :-) 

So that's basically it for 0.2.0. I plan to make such small releases more or
less right after adding a new feature so that you can easily specify a real
version number in your dependencies. 

As always you can get this release on `PyPI
<http://pypi.python.org/pypi/django-flatblocks/0.2.0/>`_. There you can also
find detailed descriptions of the new view (or take a look at the docstring of
``flatblocks.views.edit``).

Another issue I've been thinking about for the last couple of days is if this
project might benefit from an issue tracker. For me personally the only real
option here would be to move over to `bitbucket.org <http://bitbucket.org/>`_
since I actually like hg quite a bit and that service has a very simple and
elegant issue tracker. Let me think about this a little longer for now ;-)