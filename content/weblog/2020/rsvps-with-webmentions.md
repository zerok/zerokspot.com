---
title: RSVPs with Webmentions
date: "2020-06-08T17:34:11+02:00"
tags:
- webmentions
- indieweb
- rsvp
- webmentiond
- 100daystooffload
---

The last big Webmention feature that I wanted to support on both, my website *and* webmentiond, were [event RSVPs](https://indieweb.org/rsvp). With these participants of, for instance, a user group meetup can indicate whether or not they plan to attend in a decentralised fashion as opposed to something like Meetup.com where both the event’s website *and* the participants’ RSVPs are stored on a single, centralised platform.

RSVPs using Webmentions are basically extensions to comments. If I want to indicate that I’m going to be at [this month’s GoGraz meetup](https://gograz.org/meetup/2020-06-16/), then I just have to publish something like this and ping the GoGraz website:

	<div class="h-entry">
		<a class="u-author" href="https://zerokspot.com">Horst Gutmann</a> RSVPs with
	    <span class="p-rsvp">yes</span> for 
	    <a class="u-in-reply-of" href="https://gograz.org/meetup/2020-06-16/">GoGraz June 2020 meetup</a>.
	</div>

The only “new” element here is the `p-rsvp` which can have 1 of 4 distinct values:

- `yes`, I’m going to attend
- `no`, I’m not going to attend
- `maybe` I’m going to be there
- `interested` as in I’ll keep an eye out for this event but otherwise don’t know yet if I even consider attending.

The spec allows for two variants how that RSVP property can be set:

1. Using the node’s text value directly as value for the RSVP: `<span class="p-rsvp">maybe</span>`
2. Or using the `data` element in case you want the actual visible text to be something other than one of those 4 values: `<data class="p-rsvp" value="yes">I'm going to be there!</data>`.

At this point I’ve implemented the sending-part of RSVPs on zerokspot.com and most of the receiving-side on webmentiond. For the next couple of days I want to focus on fine-tuning that and ideally also ship a simple widget with webmentiond that sites can integrate in order to easily render received mentions. The first user of this widget will most likely be GoGraz.org so that [my RSVP](https://zerokspot.com/notes/2020/06/07/rsvp-gograz-june/) finally shows up there 😅
