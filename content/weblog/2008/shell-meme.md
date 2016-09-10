---
date: '2008-04-11T12:00:00-00:00'
language: en
tags:
- history
- meme
- shell
title: Shell meme
---


Finally a [geeky meme](http://www.b-list.org/weblog/2008/apr/10/meme/) :D

    $ history|awk '{a[$2]++} END{for(i in a){printf "%5d\t%s\n",a[i],i}}'|sort -rn|head
       90	cd
       74	hg
       65	python
       40	ls
       39	git
       33	curl
       19	mate
       18	bash
       15	ssh
       14	mvn

Too bad `history` isn't really global, still a nice meme for a change :-)