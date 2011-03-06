Fabric 1.0 released
===================

It was just a couple of days ago when I started wondering when Fabric_
(deployment tool, SSH, awesome) would finally hit 1.0_. It turns out `Jeffrey
E.  Forcier`_ *can* actually read thoughts and released it on Saturday with a
ton of small and medium sized changes. Just to name a few personal highlights
out of the changelog_:

* You can now manipulate the target system shell's PATH using the `path
  context-manager`_
* Another new context manager is prefix_, which allows you to prepend commands
  to ever run-call. Something that is (as the examples in the documentation
  already indicate) really useful if you're working for instance with
  virtualenv_.

Big kudos to Jeffrey and everyone who contributed to Fabric. It is one of
those tools I use for virtually every project and it time and again saves me
from the mess deployment scripts can become :D

.. _changelog: http://docs.fabfile.org/en/1.0.0/changes/1.0.html
.. _fabric: http://fabfile.org
.. _path context-manager: http://docs.fabfile.org/en/1.0.0/api/core/context_managers.html#fabric.context_managers.path
.. _prefix: http://docs.fabfile.org/en/1.0.0/api/core/context_managers.html#fabric.context_managers.prefix
.. _virtualenv: http://pypi.python.org/pypi/virtualenv
.. _Jeffrey E. Forcier: http://bitprophet.org/
.. _1.0: http://pypi.python.org/pypi/Fabric/1.0.0
