---
date: '2008-03-15T12:00:00-00:00'
language: en
tags:
- mercurial
title: I <3 Named Branches in Mercurial
---


[Mercurial](http://www.selenic.com/mercurial/) is (as probably all decentralized version control systems) specialized in support multi-branch development. Here you have basically 2 options: Separate branches physically by just cloning a branch to a different folder and then working on it there, or using the `hg branch $name` command to create a new branch right within the same repository and give it a name (hence "named branch").

-------------------------------

While you can naturally combine these two approaches, I found esp. the naming of branches very useful. The additional benefit for me is, that even after merging such a branch back to the main development branch (named 'default' by default) you can still easily determine that for instance change foo originated in branch bar since it is stored in the log::
    
    changeset:   123:abcdef123456
    branch:      newfeature
    user:        Jonny User <jonny@user.org>
    date:        Sat Mar 15 13:06:56 2008 +0100
    summary:     Can conquer Mars now too
    

So after creating such a branch, it is automatically set as active branch, which you can easily check again with the branch-command::
    
    $ hg branch
    newfeature
    
If you want to know, what named branches are all flying around within your repository, just do a `hg branches`::
    
    $ hg branches
    newfeature           123:abcdef123456
    default              122:123456abcdef (inactive)
    
Here you can also see, that "newfeature" is the currently active branch. To change back to the default-branch, just do a `hg update default`.

But how to you now merge such a branch back after you've done everything you want to do in there? Let's work on a little example here. The main development branch has the name "default" as usual in Mercurial. Now we have been working for quite some time on some glorious new feature for our world dominating app foolog in the "newfeature" branch. Now the time has finally come for this feature and we want to merge it back into the main release branch. Let's say, we have both branches also physically separated into /path/to/main and /path/to/newfeature. Now all we have to do is go into /path/to/main and verify using `hg branch` that "default" is really the currently active branch in this directory (since this will end up being the merge target). Now just do this::
    
    hg pull -u /path/to/newfeature
    hg merge newfeature

This will merge all the changes from the "newfeature" branch situated in the /path/to/newfeature directory into the currently active branch in /path/to/main. Naturally you don't have to do the `pull -u ` when just working in one single repository ;-)

When I first saw this feature, I immediately wanted to give it a try and failed miserably, simply because I couldn't find out how to merge stuff back. Also the help within hg itself doesn't really help there::
    
    $ hg help merge
    hg merge [-f] [[-r] REV]
    
Seems like someone forgot to mention, that you can also pass a branch name there, so my work on the BlueprintCSS-integration for zerokspot was basically the first time I really used this feature and this will definitely not have been the last time I use it :-)

For more details, I'd highly recommend that you take a look at the [chapter about named branches](http://hgbook.red-bean.com/hgbookch8.html#x12-1650008.5) in the excellent "Distributed version control with Mercurial" book.
