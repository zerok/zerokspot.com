---
title: About main and master branches
date: "2020-07-08T09:55:04+02:00"
tags:
- development
- git
- inclusivelanguage
- 100daystooffload
---

In recent weeks the discussion about the “master” branch name in Git has been [gaining some new traction](https://twitter.com/leahculver/status/1269109776983547904), probably related to the #BlackLivesMatter movement.

## What is the master branch?

When you create a new repository in Git, a branch is auto-created for you so that you can start committing changes right away. That initial branch has the name “master”. Based on this branch all other work starts and most likely all other branches are created at one point or another. It is the most important branch in the vast majority of Git repositories.

## Why is the term bad?

The term “master” in this context has a negative co-notation as it most likely stems from “master/slave”. I think I’ll skip the history lesson here since I think we can all agree that slavery was, is, and will always be a bad thing. Putting trigger words for people who at least somehow got in contact with it either within their or their parents generation all over the place is not necessary nor nice, to put it mildly.

If you want to learn more details, take a look at [this awesome post by Fred Hebert about inclusive language](https://ferd.ca/inclusiveness-in-language-for-outsiders-looking-in.html).

## Where is “master” coming from in Git?

According to some digging that [Bastien Nocera](https://mail.gnome.org/archives/desktop-devel-list/2019-May/msg00066.html) did last year in Git “master” probably derives from how Bitkeeper. One might argue that perhaps the term isn’t based on the “master/slave” concept here. 

The [bitkeeper documentation](https://github.com/bitkeeper-scm/bitkeeper/blob/master/doc/HOWTO.ask#L290) indicates otherwise, though:

> Now let's push the changes you have made in the slave repository to
> the master repository:

I have no idea what’s the origin for the term within Bitkeeper but my guess is that it is at least somehow based on the use of the “master/slave” concept in [ATA](https://en.wikipedia.org/wiki/Parallel_ATA) and other hardware elements. As a side-note, turns out that ATA-2 replaced master with “device 0” and slave with “device 1”.

## So let’s just rename it!

OK, so the term “master” for a branch is really a bad idea if you want your project to be inclusive. Renaming it should be relatively easy but in reality has most likely some side-effects in most projects.

Your CI/CD pipeline probably has a couple of steps that are targeted explicitly at that branch name. But even if not, historically many tools treat the master branch slightly differently since it is the default branch.

All these possible roadblocks have to be evaluated and mitigated when you decide you want to get rid of the master branch name. None of these should stop you, though, from just renaming the branch. “main” is just the better branch name anyway as it is the main development branch for your project anyway.

## Steps

1. Create a new `main` branch off of the `master` branch.
2. Update the software you’re using to manage your repository (e.g. GitHub) to use the `main` branch as “Default branch”.
3. Check your CI pipeline and other tooling to make sure that the new `main` branch gets the same treatment as the old `master` branch.
4. If possible, configure your management system to prevent pushes to the master branch (even by admins).
5. If possible, notify all people who work on the repository about the branch change so that they can switch to the new mainline branch and update their own tooling if they have one.

In general, this shouldn’t take all that long to implement, perhaps half a day at maximum with all the side-effects. It’s a small change after all. Just do it, please 🙂
