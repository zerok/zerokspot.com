---
date: '2007-07-25T12:00:00-00:00'
language: en
tags:
- bzr
- hg
- tailor
title: Ok, I'm giving up on converting repos
---


Yesterday night I once again tried to move one of my repositories over from hg to bzr. This time I took a different and more indirect approach: I first converted the whole repository into a git repository and then hoped that tailor would be able to convert this into something else (svn or bzr).

-------------------------------

Well, hoping wasn't enough, as it seems. While the conversion of the repository into git's format seems to have worked, I couldn't get tailor to move the data into a svn or bzr repository because there seems to be a problem inside of the git-support in tailor when it comes to changes that have been marked as "C" (presumably C for "Copy") operations. When such an operation should occurs, it just dies a horrible death.

The aspect where this whole problem becomes ugly is the fact, the a bug report for this already exists ... [for 1 year](http://progetti.arstecnica.it/tailor/ticket/69).

So now I could try to convert the repository to something else instead of git and not using tailor for this first step (since for some reason I can't get the hg-component of tailor to work when using it as source repository), but honestly I'm not motivated anymore. Sure, it would have been nice to keep the whole history, but for now I will probably settle with one of these options:

1. Stay with hg for this project or
2. Move *tip* over to bzr and just archive the history somewhere

I'm currently more favoring the 2nd option since I've by now customized the keyboard shortcuts of bzr exactly to my needs and the zerokspot repository is now more or less at a 1.0 milestone anyway. Let's see :-)

**Edit:** Not that this doesn't mean that I hate hg or anything like this. I was primarily trying to find out, if it's even possible to somehow convert an hg repository into an bzr repository. It works somehow. But for now the hype around hg is somehow causing people to first and foremost looking into the other direction and producing converters *from* something else *to* hg. 