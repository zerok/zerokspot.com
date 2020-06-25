---
title: Auto-approval in webmentiond
date: "2020-06-25T13:25:08+02:00"
tags:
- webmentiond
- webmention
- 100daystooffload
- feature
- policy
---

Today I started implementing a new feature for [webmentiond](https://github.com/zerok/webmentiond) that I’ve wanted to have for quite some time now: [auto-approval](https://github.com/zerok/webmentiond/commit/c8bf8c1240ac695565c46da2a8fd4e3f88ed763f). The idea here is that I’d like to auto-approve mentions from certain domains (or following specific URL patterns) in order to cut down the time it takes for them to show up on the target website. For instance, I’m pretty sure that [Jan](https://jlelse.blog/) and [Hyde](https://lazybear.io) aren’t going to spam me with mentions and so I want their mentions to go right through the approval process without me having to manually click on the “approve” button.

Right now, this feature doesn’t have a UI yet so these “policies” have to be added manually into the database:

	INSERT INTO url_policies VALUES ('^https://jlelse.blog/', 'approve', 1);
	INSERT INTO url_policies VALUES ('^https://lazybear.io', 'approve', 1);

The server checks that table every 20 seconds for new policies and then acts accordingly *after* a mention has been verified. Each policy consists of three parts:

1. A URL pattern the mention’s source is matched against
2. The action that should taken if the pattern matches
3. The weight determining the priority if multiple policies match a given URL. Lower weights cause policies to be evaluated earlier.

Let’s say, that a mention is incoming from `https://lazybear.io/something` and really links to my site. The server then verifies the mention and automatically approves it since it matches the second policy listed above.

I’ve already deployed this feature on my site and am currently testing it out under live-conditions to find out what kind of UI might work here. If this sounds like something for you, please give it a try and let me know how you like it and what could be improved 🙂

And before you ask: Documentation will come later today or tomorrow 😅 
