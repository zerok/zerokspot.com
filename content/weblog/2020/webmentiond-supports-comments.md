---
title: "Webmentiond now supports comments"
date: "2020-05-05T11:21:17+0200"
tags:
- opensource
- webmention
---

After [writing about comments yesterday](https://zerokspot.com/weblog/2020/05/04/comments-using-webmentions/
) I've now implemented experimental support for comment-type mentions in [webmentiond][w]. Using Will Norris's [microformats parser for Go](https://github.com/willnorris/microformats), the current implementation will look for an h-entry at the source-URL of the mention and use its p-name for the `title`, e-content for the `content` of the mention, and u-author for its author. All these properties will be stored inside the `webmentions` table of the database. If the entry has a `u-is-reply-to` link, then the whole mention will receive the `type` `comment`.

At this point, the content is truncated after 500 characters (or actually after 497 in order to always have a maximum of 500 characters) before being stored inside the database.

These new fields are also exposed in the relevant HTTP-API endpoints for each mention. You can see an example for this by looking at the response body of [this request](https://zerokspot.com/webmentions/get?target=https%3A%2F%2Fzerokspot.com%2Fweblog%2F2020%2F04%2F09%2Fsome-people%2F):

```
[
  {
    "id":"bq7jiguu9c2cnem5k8eg",
    "source":"https://jlelse.blog/micro/2020/04/2020-04-09-eccgd/",
    "target":"",
    "created_at":"2020-04-09T15:03:31Z",
    "status":"approved",
    "title":"jlelse's Blog",
    "content":"Some people don’t wear their mask over mouth and nose.That’s something I don’t understand at all. If some people wear masks, why don’t they wear them right? I see that sometimes when I go for a walk alone. People pass by me who have a mask over their mouth, but their nose is not covered.In order to generally avoid the narrowness in the supermarket more, I now only go shopping once a week at most and drive to a supermarket that is much bigger and where one can keep more distance than in the supermarket around the corner. Another advantage: The offer is also much better and there are still more of the things I want to buy. (But they didn’t have toilet paper either.)",
    "author_name":"Jan-Lukas Else",
    "type":"comment"
  }
]
```

To use this new implementation, simply pull [zerok/webmentiond](https://hub.docker.com/r/zerok/webmentiond):latest and restart your server 🙂

I've also updated the rendering of mentions on my blog, which makes the mention listed above to appear like this:

<figure>
<img src="/media/2020/webmentiond-comment-zerokspot.png">
<figcaption>Rendering of the comment on <a href="https://zerokspot.com/weblog/2020/04/09/some-people/">zerokspot.com</a> with special indicator that it is indeed a comment</figcaption>
</figure>

[w]: https://github.com/zerok/webmentiond
