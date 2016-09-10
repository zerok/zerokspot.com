---
date: '2006-03-17T12:00:00-00:00'
language: en
tags:
title: 'Trying Garnome 2.14 Part I: GLIBC 2.4 and Firefox'
---


Well, today Gnome 2.14.0 was released, and guess what: I wanted to give it a try right away. Considering that it normally takes quite some time for new gnome packages making their way out of the dungeons of hard-masking in Gentoo I decided to go with Garnome (once again). Sure, there are always some smaller problems involved with not going with packages of your distribution, but for me that's at least part of the fun and thrill that comes with trying to keep up with the bleeding edge of software development ;)

-------------------------------



The first bigger problem I had to face was compiling Firefox (probably as a dependency for some application that wan't to use its rendering engine ... Epiphany?). It error'd out inside xpcom/base with an error message involving the JB\_BP constant. Well, a few days ago Gentoo-unstable un-hardmasked the new GLIBC 2.4 and this seems to be part of the problem here. But thanks to some googling and some browsing the Mozilla bugtracker I found [this patch](https://bugzilla.mozilla.org/show_bug.cgi?id=323853#c47) which also seems to have made its way into the MOZILLA\_1\_8\_0\_BRANCH.



So simply get this patch, download it, go into the garnome-2.14.0/bootstrap/firefox/work/main.d/mozilla/xpcom/base folder, and add the patch using for example `cat path/to/patch | patch -p0`. **Note:** Follow this at your own risk :)



This was just the first problem, but hopefully the biggest ;) Now I got another one with missing OpenLDAP support in the PWLIB that is built by garnome, but this is hopefullly just a minor issue with the OpenLDAP version used by Gentoo :)