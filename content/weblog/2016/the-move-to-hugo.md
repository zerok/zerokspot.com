---
date: 2016-09-11T10:34:23+02:00
title: "The move to Hugo"
tags:
- zerokspot
- hugo
---

Up until yesterday, this site was powered by a home-grown system based around
Go, Git, and ElasticSearch. While I absolutely love all three, having a custom
made engine and maintain it over the years more often than not got in the way of
producing content for which that engine was originally written.

I'm also running all this on a small VPS on [Digital Ocean][] so having a
component like ElasticSearch in the stack ended up eating a large portion of the
available resources there. As such, I decided many months ago that I'd be
migrating the site over to a static site generator. While my blog had full-text
search, I rarely used it, so having all content served as simple HTML pages was
feasible.

Pretty much immediately I considered going with [Hugo][] for a couple of
reasons:

- Hugo is extremely easy to setup. Just place your content in the `content`
  folder and run `hugo`.
- It performs well for large amounts of blog posts. I currently have around 1400
  blog posts and I plan to add at least 40-50 every year. Anything that needs
  more than 5 seconds to build all that is immediately out of the race for me
  since it would keep me from creating content.
- Maintained by someone else 😉
- I've already used it for [GoGraz](https://gograz.org) and my
  [Travelogue](https://travelogue.h10n.me) for blog-like content.

Sadly also nearly immediately I ran into some issues:

- The URLs for things like feeds and taxonomies are either hard to constomize or
  not customizable at all. I really liked my old URL structure as I've built and
  fine-tuned it over many years. Parting ways here would end up requiring quite
  a lot of additional rewrite-rules in nginx and (let's be honest) I already
  have a huge mess in there from previous migrations 😉
- The folder structure doesn't lend itself for having post-specific assets. You
  basically have a single `static` folder which is fair game but on my own
  system I had given each post its own folder where I could add images,
  downloads, slides, ...

So, I (obviously...) started writing my own little generator and actually got
pretty far. I could generate all the pages I had in the dynamic version and
everything went according to plan. Until time happened. I simply couldn't find
any of it to add the finishing touches to the generator.

For this reason I decided a couple of weeks ago that I'd drop all of that, try
to recreate a more lightweighted version of the blog in Hugo, and migrate the
content over in at most a week. For this reason I've spent most of last week's
evening tinkering with a simple stylesheet and templates as well as a small
converter script that would move everything into a Hugo-compatible folder
structure.

I've tried to keep all the article URLs intact, but opted for getting rid of
most of the archive pages as Hugo doesn't yet support
[yearly archives][]. It's not a huge loss but still suboptimal as a single page
listing more than 1000 entries has it's downsides 😉

Yesterday evening I finally made the switch and rsync'd everything up. I'm
pretty sure that there are still a couple of things that are broken, but I
decided to fix them after the roll-out instead of delaying things again and
again. One thing I'm not yet sure how to handle is comments: I've used Disqus
for many years and I like it. The problem is, though, that the number of
comments per year is in the single digit(s). ~~For that reason I've decided to
leave comments out for now.~~ These days I get most of the feedback via Twitter,
e-mail, or Facebook, but adding Disqus - on the other hand - doesn't really hurt
and each page gets a quick feedback mechanism. So, for now I've re-added Disqus
to the site. Will it stay? I don't really know yet.

Anyway, I hope this move will make me create more content here again that I hope
you enjoy 😊

**Update:** I've re-added Disqus for now.

[yearly archives]: https://github.com/spf13/hugo/issues/448
[digital ocean]: https://www.digitalocean.com/
[hugo]: http://gohugo.io/
