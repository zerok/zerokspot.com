---
date: "2025-01-14T07:51:11+01:00"
incoming:
- url: https://chaos.social/@zerok/113825539795114710
tags:
- challenges
- meta
title: Blog Questions Challenge 2025
---

I haven’t done a blog challenge in a while so thanks to [Hyde](https://lazybea.rs/bqc-25/) for giving me a jump-start for the new year! This time all the questions are about this blog and my process:

## Why did you start blogging in the first place?

Back in the days I was very much into online forums from your specialized Star Trek board to just the software behind them. I used to shared things like book reviews and the latest software news I had come across. Eventually, I wanted a bit more control over my own experience there and so set up a little weblog with pretty much that kind of link-heavy content.

## What platform are you using to manage your blog, and why did you choose it?

Right now I’m using Hugo with lots of tooling and automation around it. I wanted to keep the deployment target environment for the core blog as simple as possible and nothing is simpler than a bunch of static files served by a random HTTPD (in this case Caddy).

Any dynamic functionality like the comments shown below posts is implemented using tiny standalone services. Thanks to [HTMX](https://htmx.org/) I’m now also in the process of standardizing the interface between these services and the core blog which will make standing up new dynamic features even more easy in the future.

## Have you blogged on other platforms before?

The first ideas for the blog came out of customizing a phpBB instance but IIRC that never launched. For the first couple of years I’ve used WordPress. Then I wanted to have something with a bit more “structured flexibility” and so ended up with Drupal before eventually using more or less completely custom systems.

## How do you write your posts?

Most of them are written spontaneously but I also have a list of topics that might be interesting to cover. Recently, though, I haven’t really touched that list anymore due to lack of time and focus. This is something I definitely want to get back into in the future.

The actual publishing process works like that:

- Write the draft in Obsidian. I have Obsidian installed on all of my devices and sync between them using Git, so that I can continue writing wherever I am.
- Send the Markdown file managed by Obsidian to a custom server which will create a pull request on GitHub. There I can give it one last look-over before merging the pull request
- Once that pull request is merged, automation will build the whole website and rsync it up to a server.
- Most of the times I will also manually create posts on Bluesky and Mastodon about that blog post which are automatically picked up by another service which update the blog with links to those posts.

## When do you feel most inspired to write?

“Inspired” is probably not the right word, but there are usually two situations that trigger me to write: When I learn something that solves a problem I’m having or something that has occupied my mind for a while, and when something is really annoying me 😅 Lately it has often been the latter but hopefully things will change again!

## Do you publish immediately after writing, or do you let it simmer a bit as a draft?

I wish I would sometimes hold off pushing the publish button for a bit longer. While I usually give every post at least one correction read, sometimes I should really rethink the whole premise of a post. At the same time, it just feels great to get some thoughts out as quickly as possible no matter what.

So yes, most of the time publishing happens immediately 😅

## What’s your favorite post on your blog?

I don’t have a favorite post but looking back through the years I somehow always spent a bit of time looking at my travel and conference posts. On the other hand, these are also the posts where I notice broken links or missing images which means extra work 😂

## Any future plans for your blog?

In the future I want to get back writing regularly here and also get more technical again. The last year consisted mostly of reviews and meta topics which are nice to get something off my chest. Are they useful to anybody else? I don’t know but I doubt it. In the past, writing blog posts also motivated me to play around with new frameworks and technologies (yes, weird cause and effect). I miss that time 🙂

## Who will participate next?

I would like to read answers by…

- [Dimitris](https://dsotirakis.fyi/)
- [Andreas](https://madflex.de/)
