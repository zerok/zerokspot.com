---
date: '2009-08-29T12:00:00-00:00'
language: en
tags:
- macosx
- snowleopard
title: I Love Snow Leopards
---


So yesterday, [Snow Leopard](http://www.apple.com/macosx/) got released to the world and right after work I went straight (with a slight detour through IKEA) to my local Apple retailer and got it. Back home, I first thought "Should I make a full backup?". But over the last years I've collected so much garbage on my Macbook, that losing it wouldn't be all that terrible. Besides, I naturally had TimeMachine backing up my documents.

-------------------------------

So I went straight ahead and ran the upgrade, which took for ages ... I managed to get from 90% to 100% in [Shadow Complex](http://www.xbox.com/games/s/shadowcomplexxbla/) by the time it was done ;-) So, now that I have 100% I'm ready to get into Snow Leopard. I won't write about how much I love [Quicktime X](http://www.apple.com/macosx/what-is-macosx/quicktime.html) (you can record with it ;-) and it now looks like MPlayer :-P), the Stacks improvement and just how great it is to just hold an icon in the Dock and get an Exposé of every window of that app, no matter on what space it is or whether it is minimized or not ... well, I'm not going to write more than this paragraph ;-)

Regarding update problems: Everything went really well. Naturally some stuff broke, though. Some means more or less all of [MacPorts](http://www.macports.org/) so I just kept the /opt/locals/var/db folder and wiped everything else from the face of my HDD in the hope of just installing [PostgreSQL](http://www.postgresql.org/) again and having all my databases for local development. Well, not so much. Seems like there is a problem regarding checksums and the 32bit-to-64bit upgrade. I wasn't really all that motivated to get down to the real issue (playing Shadow Complex for the 3rd time was kind of more interesting) so I downloaded Ubuntu 32bit, installed it with PostgreSQL into a VMWare, dumped the whole data folder from my old database installation into it and build an SQL dump which I could then re-import in OSX. Done.

Also: If you're using [LittleSnitch](http://www.obdev.at/products/littlesnitch/), get the Beta. Otherwise your syslog will get hammered with LS dying all the time. But that's about it, so far. For the last ~20 hours I've been really happy with Snow Leopard. There is nothing revolutionary about it but there are enough little details that make it well worth its price for me.