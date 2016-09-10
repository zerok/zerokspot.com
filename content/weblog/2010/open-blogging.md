---
date: '2010-03-20T12:00:00-00:00'
language: en
tags:
- blogging
- dvcs
- jekyll
- nodeblog
- openblogging
title: Open Blogging
---


52 podcasts in the backlog ... I really should find some time to listen to
more of them again. Anyway, I just finished listening to [episode 0.1.7](http://changelogshow.com/105/3197-episode-0-1-7-open-source-publishing-with-geoffrey-grosenbach-brandon-mathis-and-tim-caswell) of the
[Changelog Show](http://changelogshow.com/) and was "Open Blogging/Publishing".

The term "Open Blogging" as used among others by [Brandon Mathis](http://brandonmathis.com/blog/2010/02/09/edgerails.info-and-open-blogging/) means a blogging system based on a publicly available repository (e.g. on Github or BitBucket) which contributors fork and write articles as patches. If a DVCS is used, the single point of failure - the current server - mostly disappears since everyone can easily create a mirror.

-------------------------------

The technical part of that idea isn't really new. I'm definitely not the only one out there who has been thinking about building a blog based on a DVCS on and off again. It's just nice to have a real versioning system in the background to help you track (and transparently present your users) changes made to each post over the time.

Contributing to such blogs via pull requests is perhaps also not new, but services like [Github](http://github.com) and [Bitbucket](http://bitbucket.org) definitely facilitate that process.

## What's out there?

There are already quite a few collaborative blogs out there that use this approach. During the podcast sites like [edgerails.org](http://edgerails.info/) or [howtonode.org](http://howtonode.org) were mentioned with one of them probably using [Jekyll](http://github.com/mojombo/jekyll) (which also powers Github Pages) and the other being built upon a custom build-system called [node-blog](http://github.com/creationix/node-blog/).

## For everyone?

As these two examples show, this approach is mostly intended for technical writers. If you are not comfortable with a text editor and versioning systems using them for your blog won't make you happy (at least not until you've learnt them). 

But for technical content, this is a really nice approach, especially if you can integrate it somehow into the documentation process. Think about an opensource community around project X. As most OSS projects out there, the documentation of project X is a bit lacking but already uses e.g. ReST as their documentation system. Now you could easily build a tips & tricks blog using the same format and the open blogging approach and people already writing documentation for project X would feel right at home with your blog.

But what about a single author?

## For singles

Every couple of months I'm thinking about rewriting my blog and just dump the
whole database thing in order to pull all the content out of simple
plain old text files. Every time I start to write a new blog post, like this
one, the first thing I do is open a text editor ... and about a minute into 
writing it I remember that my blog actually has a web interface for that ... and 
not a bad one either since it's 100% crafted to what I usually want out of it.

Going to that web-interface kind of pulls me out of my normal process: Treating
a text editor at the one-size-fits-all solution for 90% of the problems I want
to solve. That's my primary reason to want to use a blog based on static files
on and off again. 

Naturally, the whole shared aspect of open blogging is not totally relevant here, but people could still contribute patches for typos and things like that. Also the tools used for a collaborative open blog would also apply here.

### The structural part

Mostly every time I dump the idea again, though, because of a couple of problems I never can decide on a solution for ...

1.  How should I handle metadata? 
    
2.  How should I support multiple input formats? Up until 2 months ago I mostly
    wrote these posts using ReST, before that in Markdown and now mostly in
    HTML using CKEditor. And I really don't want to stick to a single format ;-)
    The format itself actually isn't the problem, more the interaction with
    the metadata.
    
3.  How should searching (and tagging for that matter) be implemented?

4.  How should the directory structure look like to stay manageable after more
    than 1000 posts (or 100 to begin with)?

My main problem is the directory structure. Should I use the directory name as
some kind of slug, should it contain some kind of separate ID so that I can
change that slug later on ...

### Static vs. dynamic

Another question is how dynamic the frontend should be. Both Jekyll and node-blog basically just transform input data into a ready-to-host website consisting of static files. If you want to use your own commenting system or search engine, this is hardly an option. 

Right at this point the whole system might become quite complicated. How should the content be indexed, how should comments be associated with a post, etc. Especially if a site contains not only these articles but also other kind of content (e.g. a community forum, some project related database ...) some part of the post *should* end up in the database to allow cross-referencing.

In the end I'd really love to play around with this technique, be it as part of a community blog or with me being the only author. Perhaps we could use something like that for the [German Django community](http://django-de.org) or for the tech community in Graz ;-)
