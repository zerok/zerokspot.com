---
date: '2015-04-29T20:21:50+02:00'
language: en
tags:
- python
- pyformat
title: PyFormat.info
---


<figure>
<img alt="" src="/media/2015/pyformat.png"/>
</figure>

During PyCon 2015 just two weeks ago [Ulrich][ulope] and I started working on
something he has had the idea for quite some time ago but somehow we didn't get
around implementing it last year: A page that helps folks understanding what all
can be done with Python's awesome string formatters. And [PyFormat.info][pfi] is
the result.

---------------------

Perhaps we didn't start working on PyFormat.info last year because our goals
were too high and our time too little. So this time we opted for making a
classic *Minimal Viable Product*: An info page; nothing more and nothing less.

What you can find on the site at the time of me writing this is a set of
examples for rather common use-cases you might want to use Python's string
formatting for. These examples are presented for both, the old and the new API,
where possible. We ourselves always have to look the more complicated features
up in the specs so we thought: Why not collect them somewhere?!

The implementation is pretty straight forward. The site itself is nothing more
than a static website generated out of a set of PyTest test-cases. This way we
can always be sure that our code snippets actually work the way they are printed
on the final website. All "content tests" are run against Python 2.7, 3.2, 3.3,
and 3.4 right now.

If you want to contribute or see more details about the implementation, you can
find all that in our [Github repo][gh]!

We also plan to keep working on and improving it whenever time permits or we
find a new example. One of last year's high goals was to add an interactive
component to the site so that you can try the formatters right there. We don't
yet have any concrete plans there but at least [a ticket][t11] ;-)

In the meantime, I hope you enjoy what we have there right now!

[ulope]: https://twitter.com/ulope
[pfi]: http://pyformat.info/
[gh]: https://github.com/ulope/pyformat.info
[t11]: https://github.com/ulope/pyformat.info/issues/11
