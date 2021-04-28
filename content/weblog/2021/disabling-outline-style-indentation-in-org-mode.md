---
title: Disabling outline-style indentation in Org Mode
date: "2021-04-28T21:44:09+02:00"
tags:
- emacs
- 100daystooffload
---

One thing that has driven me mad over the last months was the default indentation that [Org Mode](https://orgmode.org/) applies whenever you write something below a headline. Basically, content is always indented to visualise the hierarchy of the content. This is useful for outlining but I mostly use to Org for general note-taking (and even long-form writing from time to time) and there this default behaviour simply didn’t fit my writing style 🙂

So this would be Org’s default behaviour:

	* Chapter
	  level 1 indentation
	  
	** Deeper
	   level 2 indentation

And what I want most of the time is this:

	* Chapter
	some content
	
	** Deeper
	no indentation here

This just works better for me especially when dealing with things like src-blocks etc. which have slightly different semantics when being indented (whitespaces FTW!).

Luckily, other folks have also had this issue and documented it for instance on [StackExchange](https://emacs.stackexchange.com/questions/41220/org-mode-disable-indentation-when-promoting-and-demoting-trees-subtrees). Turns out that you can toggle this automatic indentation using the `org-adapt-indentation` flag:

	" enable indentation
	(setq org-adapt-indentation t)
	
	" disable indentation
	(setq org-adapt-indentation nil)

This variable can also have other values but `nil` is doing exactly what I want right now 😅 That being said, I might add a quick keyboard-toggle for this flag just in case I need to do more outlining again in the future…
