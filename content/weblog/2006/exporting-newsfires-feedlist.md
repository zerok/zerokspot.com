---
date: '2006-01-18T12:00:00-00:00'
language: en
tags:
- macosx
title: Exporting NewsFire's feedlist
---


If you want to automate exporting your feedlist to put it onto your homepage or something like that, you normally use something like AppleScript under MacOSX to tell the application to create an export using its own libraries. If your reader doesn't offer any AppleScript you have to find another reader or simply try to get into the data structures used by the news reader. As I've already written yesterday, I'm now using [NewsFire](http://www.newsfirerss.com/) which offers nearly everything I want except a simple way to export the feedlist from the shell. It offers no AppleScript but there are other ways around this problem (as the word "simple" in the last sentence should have already indicated ;)).



-------------------------------



So after searching a little bit I found where NewsFire is storying its (1) messages and later (2) also the feedlist itself so I could started to try to find a way to get the data out of these files. Now, after a few hours of work I think I'm more or less done with it and plan to post it here somewhen tonight after some testing. It's basically a 150lines Ruby script that uses and extends [Patrick May's plist parser](http://www.narf-lib.org/2006/01/plistxml-parser-for-ruby.html) which will also support exporting of groups (which is quite experimental but still should do it for now).

Besides testing I also want to first play a little bit around with my XSL that transforms the OPML file to HTML and I also want to make sure, that what the script is producing is in fact OPML ;)