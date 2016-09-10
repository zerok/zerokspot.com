---
date: '2007-07-22T12:00:00-00:00'
language: en
tags:
- bzr
- hg
title: 'Hg to Bzr: A rocky road'
---


While Git seems to get quite some [buzz](http://daringfireball.net/linked/2007/july#thu-19-git) in the Mac community over the last couple of days I personally am more leaning towards bzr as I already wrote a couple of days ago. 

Two days ago I now wanted to make the switch and move the repository for this site from Mercurial over to bzr first of all just to see if it is possible at all. 

-------------------------------

So I discovered, that there are 2 options for such a step right now:

1. The [bzr-hg](https://launchpad.net/bzr-hg) plugin for bzr which allows you to read-only access and pull a Mercurial repository
2. and [tailor](http://progetti.arstecnica.it/tailor) which is more or less a generic conversion tool for all the better known RCS out there (svn, cvs, bzr, hg, ...)

And in this order I also gave them a try. But right after installing the bzr-hg plugin I faced the first problem: With version 0.9.4 the Mercurial developers introduced a new module loading system immediately found some "friends" among 3rd-party developers that wanted to integrate hg but faced some problems as described in [this](http://www.selenic.com/mercurial/bts/issue605) bug in the Mercurial bug tracker. Thanks to Tim Hatcher this problem was swiftly [resolved](https://bugs.launchpad.net/bzr-hg/+bug/127181) resolved within bzr-hg, so I could continue :-)

After finally being able to pull my whole repository, I made some cleanup and then committed the changes and afterwards wanted to push it to my primary backup-repository. Well... as EGM's [Shane Bettenhausen](http://www.1up.com/do/my1Up?publicUserId=1002415) likes to say: "It's nice to want something."

Somehow it seems like the whole pulling didn't went as successful as I'd thought. Whenever I try to push (or branch), I just get something like this:

> bzr: ERROR: Revision {hg:c2e0a7156b4216b7c01f47d72382324a5b42e330} not present in KnitVersionedFile(file:///Users/zerok/tmp/test/.bzr/repository/knits/38/hg%253adjango%253azerokspot%253ashared%253atemplates%253azerokspot%253ashared).

So I guess, the first path (bzr to hg via bzr-hg) has reached a dead end for now. On to the next option: tailor

To make a long story short: No success on this front either. When I tried following config ...

	[project]
	source=hg:source
	target=bzr:target

	[bzr:target]
	repository=./target

	[hg:source]
	repository=/path/to/zerokspot/

In order to try out some things I downgraded to hg 0.9.3 which gave me more or less the same error as mentioned [here](http://progetti.arstecnica.it/tailor/ticket/102). Then I upgraded again to 0.9.4 and got this one:

>  Common base for tailor exceptions: 'bzr' is not a known VCS kind:
> 'module' object has no attribute 'util'

(Note that this is the verbose output ;-) )

So it seems that the problem here shifted from tailor not being able to play nicely with hg to not being able to even notice that it actually should support bzr (at least according to the docs).
