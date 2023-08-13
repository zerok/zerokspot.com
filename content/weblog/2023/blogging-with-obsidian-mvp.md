---
date: "2023-08-13T11:15:55+02:00"
tags:
- blogging
- zerokspot
- obsidian
title: 'Blogging with Obsidian: The initial implementation'
---


Despite not having too much time for working on my little vacation project, I managed to write a little [Obsidian plug-in](https://github.com/zerok/obsidian-http-publish) for sending a single note to an HTTP endpoint and a [small server that receives](https://github.com/zerok/zerokspot.com/pull/415) that note and creates a pull-request for my blog on GitHub.

At this point half of the implementation is basically just a big mess but I'm going to iterate on both parts of it in the near future! Since the Obsidian plug-in is my first one, it's the bigger mess of the two of them but at least it works for me right now. Its state is also the reason why I'll probably not submit it to the official registry anytime soon. For now it's a plug-in for me alone and the code is online just so that others can learn from my mistakes 😅