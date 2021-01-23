---
title: To use Git LFS or not
date: "2021-01-23T18:53:40+01:00"
tags:
- zerokspot
- git
- 100daystooffload
---

A couple of days ago I had the glorious idea to move all the images I had on this blog into [Git LFS](https://git-lfs.github.com/). That system only exists to manage large files within Git and so I thought it would be a good fit.

A couple of days later (and an invoice from GitHub) I’ve come to the conclusion that LFS might not be a good fit after all. First of all, it moves the actual images outside of Git, relying on an external service to store those files. Since I want to really have everything on this blog stored within just a single place if for nothing else than as a backup-solution, this should have raised some flags in my mind.

LFS - as far as I understand it - is mostly useful you have lots of large files not within the main branch. If that’s the case then there are situations in which those files are simply never downloaded. This won’t happen here as the main branch contains pretty much everything. There might be times when I have an image in a draft branch that is later removed, but that branch will be squashed into main and the image eventually garbage collected.

Another scenario where LFS is probably immensely useful is when you have a large binary file that is often changed. In this case only the new file pointer is added to the Git repository, keeping the repository itself small while duplicating the file content on the LFS store. Again, that’s simply not a use-case I have here. Sure, there might be a situations when I change a picture after I’ve published it, but these will be really rare and having everything completely managed by Git is just far more useful to me.

So today I have for the first time done a big rewrite of history on this website’s Git repository:

	$ git lfs migrate export --include 'static/**/*.jpg' --everything
	
	$ git lfs migrate export --include 'static/**/*.jpeg' --everything
	
	$ git status
	On branch main
	Your branch and 'origin/main' have diverged,
	and have 2570 and 2570 different commits each, respectively.
	
	$ git push origin main --force

I really hope that I won’t have to do such a rewrite again…
