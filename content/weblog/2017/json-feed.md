---
date: 2017-06-05T13:03:16+02:00
title: JSON Feed
tags:
- blogging
- json
---

It was only a matter of time until someone would go out and try to replace Atom
and RSS, the two main formats when it comes to subscribing to blogs, with
something based on JSON instead of XML. Brent Simmons and Manton Reece have done
exactly that now with [JSON Feed](https://jsonfeed.org).

This new format tries to encompass all the things the community has learnt about
publishing content online over the many years that RSS and Atom have been around
(like micro-blogging, Twitter-feeds, …). It's also supposed to be easier to work
with, as JSON is usually easier to work with than XML.

One of the huge advantages of XML over pretty much anything else is that it is
clear how it can be extended. In practice, this is rarely done, though, as it is
rather complicated. In JSON Feed, extensions are simply JSON properties that
start with an underscore:

```json
{
  "version": "https://jsonfeed.org/version/1",
  "title": "My blog",
  ...
  "items": [
    {
      "id": "123",
      "content_text": "some text",
      "url": "https://domain.com/123",
      "_location": {
        "lon": "...",
        "lat": "..."
      }
    }
  ]
}
```

I'm pretty sure, there will eventually be some standardisation happening around
that, but making the format easy to extend is a great start.

There are some downsides (at least in version 1), though:

- You can no longer specify which language the content is in. XML has a core
  attribute for that (`xml:lang`) which hasn't been ported over yet. There
  is [an issue for that](https://github.com/brentsimmons/JSONFeed/issues/40) in
  the official tracker, though.
- XML also allows to set something like
  a [base-URL](https://www.w3.org/TR/xmlbase/), which is used inside the content
  for relative references. Like if you have an image with the URL
  `/images/test.png` you can set a base-URL so that the feed reader knows to
  fetch the image from `https://domain.com/images/test.png` instead of
  `https://feedly.com/images/test.png`, for instance.

Personally, I think that having something like JSON Feed around is a good
thing. While it is basically [XKCD 927](https://xkcd.com/927/) and Atom and RSS
are already decent formats, it was time for something to take the things we've
learnt from using them over the years and merge that into something new. The
community around this new format also appears to be quite active so even if
you're not adopting JSON Feed yet it is worth keeping an eye on it 🙂 So while
points like a custom media type and link-rel info are still not finalised, I'm
looking forward to what the community will do around this format!

Turns out, Cathal Garvey has already opened
a [feature request for Hugo](https://github.com/spf13/hugo/issues/3487) (the
blog system I'm using here) so I might get JSON Feed here sooner than expected 😀

P.S.: YAML Feed, anyone? 😉
