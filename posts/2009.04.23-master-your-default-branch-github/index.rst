Master your default-branch on Github
####################################

Sometimes small features are just as important as big ones. For nearly 6
months I've been waiting for such a small feature and tonight github finally
implemented it: The option to `choose the default branch
<http://github.com/blog/421-pick-your-default-branch>`_ of a repository. 

Why is this so important for me? I have a `small repository
<http://github.com/zerok/django-docs-de/tree>`_ where we translate the
official `Django documentation <http://docs.djangoproject.com/>`_ into German. Since we want to stick as close to
the structure used upstream as possible, we have the documentation for 1.1 in
the master-branch and the one for 1.0 in the releases/1.0.X branch. The
problem was until now that most of our efforts right now happen in the
1.0-branch which we'll then simply merge into 1.1 once it is mostly complete.
Previously we had to put a huge WARNING in the description of our repository
to tell people what branch to use, now we can make that just a small warning
;-)

So now, when you visit the translation's project on Github, you see the
1.0-branch right away :D
