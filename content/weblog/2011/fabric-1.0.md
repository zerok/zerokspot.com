---
date: '2011-03-06T12:00:00-00:00'
language: en
tags:
- software-release
- python
title: Fabric 1.0 released
---


It was just a couple of days ago when I started wondering when [Fabric][fab]
(deployment tool, SSH, awesome) would finally hit [1.0][1.0]. It turns out [Jeffrey
E.  Forcier][jef] *can* actually read thoughts and released it on Saturday with a
ton of small and medium sized changes. Just to name a few personal highlights
out of the [changelog][log]:

* You can now manipulate the target system shell's PATH using the [path
  context-manager][path]
* Another new context manager is [prefix][prefix], which allows you to prepend commands
  to ever run-call. Something that is (as the examples in the documentation
  already indicate) really useful if you're working for instance with
  [virtualenv][venv].

Big kudos to Jeffrey and everyone who contributed to Fabric. It is one of
those tools I use for virtually every project and it time and again saves me
from the mess deployment scripts can become :D

[log]: http://docs.fabfile.org/en/1.0.0/changes/1.0.html
[fab]: http://fabfile.org
[path]: http://docs.fabfile.org/en/1.0.0/api/core/context_managers.html#fabric.context_managers.path
[prefix]: http://docs.fabfile.org/en/1.0.0/api/core/context_managers.html#fabric.context_managers.prefix
[venv]: http://pypi.python.org/pypi/virtualenv
[jef]: http://bitprophet.org/
[1.0]: http://pypi.python.org/pypi/Fabric/1.0.0
