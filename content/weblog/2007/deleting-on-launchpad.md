---
date: '2007-09-27T12:00:00-00:00'
language: en
tags:
- bzr
- launchpad
title: Deleting on Launchpad
---


There you've ignored [Launchpad](http://launchpad.net) for 2 months and when you take a look at it again, and all of a sudden it includes all the features you complained were missing the last time you checked. Perfect.

-------------------------------

A couple of months ago, when I was in the dilemma of what DVCS to use for a small project of mine, I originally wanted to give Launchpad a try, since it offered me a direct way to integrate bzr into it instead of going the dual-way of using hg/bzr locally and then publishing the actual code using svn, as I'm now doing it on Google Code.

Back then I had some trouble importing my code into Launchpad's repository system and ended up with a dead repository that I couldn't delete. The reasoning is in my opinion quite understandable, since you can end up with tons of dependencies between the branches and therefor you should only be able to delete branches, no other branch is depending upon. I then asked in the official chatroom about it and someone told me, that this was planned for one of the next releases. But since I wanted to get the code out of the door as soon as possible, I simple went over to Google Code and switched to SVN for it. 

Now that [bzr](http://bazaar-vcs.org/) finally seems to close in on the 1.0 release with 0.9x releases for the last month, I also took another look at Launchpad and noticed, that you can now finally also remove branches. It's just a shame that this update already [happened in August](http://news.launchpad.net/releases/launchpad-118-released) and I didn't notice it :-(

With this feature finally integrated, I guess Launchpad is feature-wise ahead of Google Code, yet the whole usability is IMO still a little bit lacking. Most or the dialogs are simply too complicated and resemble for my taste way too much those of Sourceforge. But it's hard to overlook, that Launchpad has some killer features (at least in my eyes) when it comes to community integration. Who else has a built-in translation system? ;-)

