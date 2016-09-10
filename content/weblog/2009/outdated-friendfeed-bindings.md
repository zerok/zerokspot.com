---
date: '2009-06-14T12:00:00-00:00'
language: en
tags:
- bindings
- friendfeed
title: Are all FriendFeed bindings outdated?
---


Is it just me or are most language bindings for the [FriendFeed API][] mostly
outdated? I've now looked at the [official ones][] for Python, PHP and C# as
well as two for Java and none of them offers the extended search facilities
provided by the API , e.g. the domain search or the URL fetcher. The
["real-time"-methods][] are nowhere to be seen, either. And if you want to at
the missing functionality to some of the Java bindings, you end up running
into the private-method-wall -_-.

So far I've only found one binding for Python (albeit not the official one) by
[Chris Lasher][] that supports at least the domain-search. (A [new version][]
of it is out now, btw. :D) 

On a side-note: If you write a binding for a language, (big, big, big) please
make sure that it's easy to use. It doesn't really help if a PHP binding is
not in PEAR, a Java binding is in no Maven2 repository or a Python binding is
not on PyPI.

[chris lasher]: https://launchpad.net/friendfeed-pyapi
[official ones]: http://code.google.com/p/friendfeed-api/
[friendfeed api]: http://code.google.com/p/friendfeed-api/wiki/ApiDocumentation
["real-time"-methods]: http://code.google.com/p/friendfeed-api/wiki/ApiDocumentation#Real-time
[new version]: https://launchpad.net/friendfeed-pyapi/0.2/0.2.0
