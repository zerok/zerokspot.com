---
date: '2005-07-11T12:00:00-00:00'
language: en
tags:
- ubuntulinux
title: Ubuntu > Laptop
---


Now I've finally come to the point where I don't want to stress my poor laptop with having to compile everything so I've removed Gentoo from it and installed Ubuntu tonight. This doesn't mean that I don't like Gentoo anymore or something like that. I still use it on my server and it's still my first choice if I have high performant hardware at my disposal.

-------------------------------



A main reason for installing Ubuntu on the laptop is that I want to play around with apt and also to have two distributions instead of just one at my disposal for testing software. First I also thought about using Debian but then I noticed that Debian Sarge is still using XFree86 and not X.org. Since I'm quite a fan of X.org and I also prefer a 6month-release-cycle above a n-years-release-cycle Ubuntu seems to be the right choice for me.



Since I had no Hoary DVD or CD at home I installed 4.10 and then updated to 5.04. No big problems there except a small problem with postfix which was resolved by removing and re-installing ubuntu-base and postfix. The next problem where the old version of Firefox and Thunderbird as well as unison in the Ubuntu repository. Sure, Ubuntu always offers the latest security fixes but in the case of Firefox an update of the version number would be better considering the version check on http://addons.mozilla.org . But it's not really a problem. I simply downloaded Firefox and installed it manually while keeping the original package to not stress my friendship with apt on the first day by violating some dependencies with yelp ;-)



Also the problem with unison was quickly resolved by simple installing unison manually on both server and client. This was necessary since unison seems to be a little bit fussy when it comes to using different versions on the client and the server.



Tomorrow ... today ... whatever will probably see the installation of MPlayer on this system which should basically mark the last big installation. After that I'll perhaps just replace metacity with openbox again and other minor customizing stuff.