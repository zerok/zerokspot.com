---
title: "Comments using Webmentions (some planning)"
date: "2020-05-04T10:21:48+0200"
tags:
- webmention
- indieweb
- blogging
- commenting
---

With my current [webmention](https://webmention.net/) implementation being rather stable now I think I'll set a new goal for the project: I want it to also support [comments as defined on the IndieWeb wiki](https://indieweb.org/comments). With this extension to the protocol the target of a mention checks the content of the source for additional information about the content in order to display it like a classic "comment" at the target URL.

This has implications for the receiver *and* the sender of a mention. 

## The sender/source

The actual webmention itself looks exactly the same *but* the document that is available at the source has to have some specific characteristics in order to be recognised as comment. A slightly slimmed down version of these requirements for a source document includes these:

1. It has to have a h-entry.
2. Ideally, that h-entry should hold an [in-reply-to URL](https://indieweb.org/in-reply-to) so that the target knows if it was just mentioned or if the source is meant as a comment to the target.
3. It should contain either an e-content or p-summary element that is "not too long" and that should then be displayed by the target.

These are all things I will have to include in my Hugo setup. At this point, I think, I'm going to just add another parameter to posts that are meant as comments which points to the page I want to comment and render that link according to the comment-rules described above. If time permits, I want to get this done over the course of the next couple of days.

## The receiver/target

Since I also want to render comments on zerokspot.com, I will have to extend the content parser inside [webmentiond](https://github.com/zerok/webmentiond/) and add more microformat2 support. Or (more likely) I will finally give [willnorris/microformats](https://github.com/willnorris/microformats) a try 😅 In any case, the Mention struct will probably gain a couple of new fields:

- `Content`
- `AuthorName` 
- `AuthorURL`
- `AuthorImage`
- `Type`

In the case of a comment, the `Type` will probably simply be `comment` to distinguish it from a "normal" mention.

I have no idea yet when I will get around implementing this. I hope, though, it won't take me too long to get this into webmentiond. I really want it now that I've seen comments in the wild 😅
