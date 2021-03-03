---
date: '2009-03-17T12:00:00-00:00'
language: en
tags:
- python
title: Whoosh looks interesting
---


This morning I read a great post by [Arne][] where he introduced me to [Whoosh][], a
pure-Python search back-end. Traditionally when you're working with
site-specific search-technologies, you sooner or later stumble upon the
[Lucene][] ecosystem (since there is a whole forest of applications surrounding
it, I guess you could really call it an ecosystem by now). Lucene might be the
best thing since the invention of sliced bread but especially for smaller
sites it might just be too much configuration overhead. Surely, some of the
applications and libraries around Lucene have made that whole process much
easier -- [Solango][] comes to mind, there -- but that might still not be enough.

[Whoosh][] on the other hand appears to be quite Python-targeted (while Lucene
goes after *any* environment) with a very simple configuration- and
operation-workflow: Create a storage, create an index above it, define what
should be indexed, hand the index a document you want to index according to
your schema, done. No running some Java application in the background, no
XML-schema declarations.

I'm not yet sure, if I will go with Whoosh or Solango for the next iteration
of this site but Arne definitely provided me with yet another option to think
about, thanks :-)

There are a couple of aspects I'm not yet really sure about with Whoosh,
though. For example, the searcher (the thing that you pass your query to
receive the actual hits) doesn't seem to provide any offset mechanism. This
way, running any kind of pagination over the result-set would end up being a
bit of a problem. But perhaps there is already a solution for this out there
:-)

[Arne]: http://www.arnebrodowski.de/blog/add-full-text-search-to-your-django-project-with-whoosh.html
[Lucene]: http://lucene.apache.org/
[Whoosh]: http://whoosh.ca/
[Solango]: http://code.google.com/p/django-solr-search/
