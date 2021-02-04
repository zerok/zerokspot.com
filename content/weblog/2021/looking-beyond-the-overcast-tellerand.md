---
title: "Looking beyond the Overcast Tellerrand"
date: "2021-02-04T18:14:00+01:00"
tags:
- 100daystooffload
- podcastapp
- ios
- overcast
- castro
incoming:
- url: https://chaos.social/@zerok/105674150034706813
- url: https://twitter.com/zerok/status/1357379974760693761
---

For the last 5 years, [Overcast](https://overcast.fm/) has been on my iPhone as the one app that I did never consider switching away from. It just works for me. That being said, its feature set hasn’t really changed all that much since then. Sure, sharing has been improved and the support of chapters is more than just welcome, but these were the last big and visible new features. That being said, there are some less visible ones like the privacy screen that tells you about each podcast if they’re tracking you but at its core Overcast hasn’t changed much since 2015.

Triggered by [this post by Vincent Ritter](https://vincentritter.com/2021/02/02/19-43-24) I thought of what I’d like to see improved in Overcast and actually couldn’t come up with anything. It’s good software but is it perfect? Most likely not and so today I started playing around with two of the better known alternatives: PocketCasts and [Castro](https://www.castro.fm/).

Sadly, I quickly gave up on PocketCasts again. While it was my primary player while using Android many years ago, the current interface is just not right for me. It does pretty much the same things as Overcast but with a less interesting look. So I crossed it off my list and jumped over to Castro.

## How’s Castro?

A small disclaimer before getting started here: Since I’m an Overcast subscriber, I won’t make the distinction between Castro and Castro Plus here. Yes, the Overcast subscription costs €10 per year while Castro Plus is €19 but I would take that anyway and so only considering the free features here would be unrealistic for me. I just expect an ad-free application and also one that supports the trimming of silence. Both are locked behind Castro Plus and so the free version is not an option for me.

Castro is conceptually slightly different than Overcast. You usually play episodes through your “queue” instead of using one of multiple playlists. Basically, that queue is one big playlist which you make your way through, eventually.

If you have lots of subscriptions, that won’t work as that queue would just become one big mess without having some kind of gatekeeper that prevents stuff from being added there automatically. That’s where the “inbox” comes in. Whenever one of your subscriptions has something new to offer, these new episodes will show up in your inbox. There you can select if you want to skip these new items or move them either to back or the front of your queue (or start listening to them right away).

<div class="figure-group">
{{< figure caption="From the inbox you can move an episode to your queue, skip it, or start playing right away." src="/media/2021/castro-inbox.jpg" >}}
{{<figure src="/media/2021/castro-skip-inbox.jpg" caption="New episodes can skip the inbox either by default or per show." >}}
</div>

If you have some shows that you always want to listen to next (or eventually), you can configure them to skip the inbox altogether, having them queued up directly as you see fit. For example, I know that I always want to listen to [“The Pen Addict”](https://www.relay.fm/penaddict) right away so I’ve configured it to skip the inbox and go straight to the top of my queue. Same for [“Wissen aktuell”](https://oe1.orf.at/collection/581950) and other German daily shows.

The third and last major area of the app is the “library”. Here you see all your subscriptions and can add single episodes either to your inbox or directly to your queue. This is especially useful if you’ve just subscribed to a new show but want to also listen to some parts of its back-catalog. Additionally, the library also has a “Starred” section which contains all the episodes you’ve given a start to and a “History” section where you can see what you’ve listened to in the past.

<div class="figure-group">
{{<figure src="/media/2021/castro-library.jpg" caption="Castro's library also includes starred episodes and a history!">}}
{{<figure src="/media/2021/castro-chapters.jpg" caption="Chapters can be unchecked so that they're skipped.">}}
</div>

Once you start listening to an episode, the interface looks very similar to Overcast with two exceptions:

1. The show notes are shown in a different screen that is removed from the player.
2. If an episode contains chapters, you can access a list of these chapters and uncheck them so that the player just skips them once you get there! That’s a nice improvement over what Overcast has right now.

## Feature wish: Auto-skip of chapters per podcast

This also brings me to a feature that I’d really like to see in Overcast or Castro on top of that skipping of chapters: I would like to keep a list of chapter name patterns for each show that should be skipped without me having to do that unchecking manually whenever an episode starts.

Let’s say I have a podcast like Methodisch Inkorrekt, a German podcast about physics and science in general. Every episode they have a chapter where they do a little experiment and also at least one chapter with some music playing. I’d love to always skip those chapters (not because they are bad but simply because I’m not into them) and so I’d set these chapter name patterns to be skipped by default:

- `Musik`
- `Experiment.*` 

Once the player then starts the episode it would mark the chapters “Musik” and “Experiment 1” and “Experiment 2” by default for skipping.

## What now?

OK, but do I like Castro more than Overcast? It’s too early to tell but I will try to make up my mind until that Castro Plus trial expires 🙂 The whole idea behind the queue and inbox matches my working-style very well but I want to experiment with it for a couple more days.
