---
date: '2010-11-21T12:00:00-00:00'
language: en
rating: 3
tags:
- books
- review
title: 'HTML5: Up and Running by Mark Pilgrim'
---


As usual first a short **disclaimer**: I've received a review copy of this
book from [O'Reilly][].

For the impatient among you a short summary: [HTML5: Up and Running][] is a book
presenting you some of the more prominent features of [HTML5][] including [canvas][],
[localStorage][] and the [applicationCache][]. It sadly leaves out a whole range of
other features like [WebSockets][] and the [data-attributes][]. Also the writing style
reminds you often enough that this book was created out of the author's [Dive Into HTML5][]
project with each chapter and sub-chapter being able to stand alone
which (thanks to the repetition of explanations et al.) makes reading certain
chapters not all that pleasant.

In general, if you want a book that teaches you some of the new core features
of HTML5 and you really want a *book*, this might be something for you.

-------------------------------------------------------------------------------

[HTML5: Up and Running][] by Mark Pilgrim is described on the publisher's
website as a "guide [...] through the important changes" that come with HTML5.
The focus here lies on features like the new form-elements, canvas, video and
audio integration, localStorage and the applicationCache as well as the
[microdata][] component of HTML5. All these are presented with a couple of
examples so you should know how to use them for basic operations after reading
the respective chapters.

HTML5 is put into perspective with the first chapter that tells some of the
history of HTML and how changes to the language have been proposed in the past
and how HTML5 came to be. Even if you have known most of this before, it is a
nicely written summary.

But back to the present. Sadly the described range of features is far from
covering all that is important about HTML5. The author forgot about things
like the sessionStorage_, WebWorkers_, WebSockets or the data-attributes.
Instead the books contains a couple chapter telling you on how to use ffmpeg
and Handbrake to create videos in the codecs and containers supported by most
modern web browsers. Space that could have been better used for the new
network tools, in my opinion.

Surprisingly the appendix, which contains a short summary on how to detect the
new features in a browser (which the author himself summarized with "Try
Modernizr"), also contains some of the features not mentioned in the "main
chapters". That said, the appendix in itself was in my opinion completely
useless since it is basically just a collection of links to each feature and
the detection script. The "missing features" received no explanation what so
ever.

Another big issue (at least for me) is the writing style as mentioned above.
Explaining the first example line-by-line is a good thing. Repeating that with
every following example that only contains a minor variation of the solution
might end up annoying the reader (as it did with me).

-------------------------------------------------------------------------------

Don't get me wrong, though: This a good book if you're new to HTML5 and want a
*book* that teaches you some of the new features without going into
controversial decisions like microdata vs. rdfa vs. microformats. If you want
more or even a complete overview, you will have to look somewhere else,
though.

[o'reilly]: http://oreilly.com
[HTML5: Up and Running]: http://oreilly.com/catalog/9780596806033
[websockets]: http://dev.w3.org/html5/websockets/
[applicationCache]: http://www.w3.org/TR/offline-webapps/
[data-attributes]: http://dev.w3.org/html5/spec/Overview.html#embedding-custom-non-visible-data-with-the-data-attributes
[localstorage]: http://dev.w3.org/html5/webstorage/
[sessionstorage]: http://dev.w3.org/html5/webstorage/
[webworkers]: http://www.whatwg.org/specs/web-workers/current-work/
[html5]: http://dev.w3.org/html5/spec/Overview.html
[canvas]: http://dev.w3.org/html5/canvas-api/canvas-2d-api.html
[dive into html5]: http://diveintohtml5.org/
[microdata]: http://dev.w3.org/html5/md/
