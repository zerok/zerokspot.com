---
date: '2007-01-16T12:00:00-00:00'
language: en
tags:
- drupal
- pathauto
title: 'Hints: Disabling pathauto for nodes'
---


As you might have noticed, I'm using pathauto here for the tag URLs. The problem is, though, that I really want to use pathauto _only_ for the taxonomy and not also for the nodes :)

<a href="http://zerokspot.com/uploads/pathauto-nonodes.big.png" title="Empty all the pattern fields" class="figure thickbox"><img alt="Empty all the pattern fields" src="http://zerokspot.com/uploads/pathauto-nonodes.png"/></a>


-------------------------------


After some messing around with it, there is a simple solution for this: Simply keep all the pattern fields for nodes empty. I also tried it with simply setting the default pattern to "node/[nid]", but this only produced aliases from node/123 to ... node/123 :-)