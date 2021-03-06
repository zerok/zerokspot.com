---
date: '2009-06-09T12:00:00-00:00'
language: en
link: http://www.w3.org/Bugs/Public/show_bug.cgi?id=6999
source: http://www.zeldman.com/2009/06/08/not-safe-for-work-tag-in-html-5/
source_title: "Jeffrey Zeldman Presents  : \u201CNot Safe For Work\u201D tag in HTML\
  \ 5"
tags:
- html5
- nsfw
title: Not Safe For Work HTML5
url_title: "Bug 6999 \u2013 Adding tag <NSFW>"
---


Some feature requests are awesome, some good, some are strange and some are
just weird. I'm not yet sure where to put [this request][] for HTML5 that
proposes adding a ``<nsfw>`` ("not safe for work") tag into the language
specification. While it definitely has more semantic meaning than adding a
CSS-class "nsfw" into each relevant node, it in my opinion doesn't warrant its
own tag. If you think it through, you would also have to create tags like
``<spoiler>``, ``<shoppingcart>`` and ``<answer>``. 

In a comment on [Jeffrey Zeldman's blog][] Darcy Murphy suggested the use of
``rel="nsfw"`` which is probably a really good solution for links, but
probably doesn't help with inline-content. There using a class *is* perhaps
the best way also when it comes to semantics. In either case, native support
within browsers as perhaps part of some profile system (plus some parental
control) is required to make the whole idea useful.

[this request]: http://www.w3.org/Bugs/Public/show_bug.cgi?id=6999
[jeffrey zeldman's blog]: http://www.zeldman.com/2009/06/08/not-safe-for-work-tag-in-html-5/
