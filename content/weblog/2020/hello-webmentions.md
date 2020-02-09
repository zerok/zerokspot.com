---
title: "Hello, Webmentions"
date: "2020-02-09T17:55:00+01:00"
tags:
- webmentions
- zerokspot
---

If you've been online and had a blog in the heydays of blogging and the
blogosphere, you probably somehow stumbled across [Pingback][p] and
[Trackback][t]. Both allow owners of blog A to notify blog B if one of A's
articles links to an article on B.

A couple of months ago I stumbled across [Webmention][w] which is basically a
re-implementation of Pingback without the XML-RPC part (yes, XML-RPC, ...) so,
in theory, it should be less of a pain to implement. I also saw what could be
done with Webmentions thanks to (among others) [Aaron Parecki's blog][a] where
he exposes a simple form so that people can add mentions even if their websites
don't support Webmentions natively.

Since I wanted to get to know the protocol itself, I started implementing my
own little [Webmention receiver and server][o] and during FOSDEM last weekend I
finally switched it on for zerokspot.com. So, if you want to send me Webmentions, you
can do so now! All mentions are moderated through a simple web-interface, so
there *will* be a delay between you sending a mention and me approving it being
rendered on zerokspot.com 😉

And if you don't and nobody else does, then at least I had some fun playing
with the protocol! There are still lots of things I want to improve there (like
a better rendering here AND a simple form to send Webmentions to specific
posts), so I'm pretty sure I'm going to continue having fun 😄


[p]: https://en.wikipedia.org/wiki/Pingback
[t]: https://en.wikipedia.org/wiki/Trackback
[w]: https://webmention.net/
[a]: https://aaronparecki.com/
[o]: https://github.com/zerok/webmentiond
