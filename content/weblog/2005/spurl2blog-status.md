---
date: '2005-05-05T12:00:00-00:00'
language: en
tags:
- development
- python
title: spurl2blog status
---


<a href="http://www.zerokspot.com/gallery/image.php?id=60&action=viewfull" class="left" title="View a larger version of this screenshot"><img src="http://www.zerokspot.com/gallery/image.php?id=60&action=viewthumb" alt="Screenshot"/></a> After todays <cite>Charmed</cite> episode I started working on spurl2blog again and I'm making quite good progress. Fetching the feed and posting it to a WordPress weblog already works (as can be seen in the screenshot which is <strong>not</strong> a mockup ;-) ). I had no time for adding some error handling code or something like that which currently keeps me from releasing the code. The whole configuration for now has to be done within the spurl2blog.py itself which shouldn't be too complicated IMO. I'm currently lacking enough motivation to use optparse for handling commandline arguments, sorry :-)

-------------------------------



Before releasing the code I want to add some options to at least partially allow the user to change the look of the posts without having to search the whole source file for the code creating the post ;-) Perhaps tomorrow :-)