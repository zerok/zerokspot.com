---
title: Finally evaluating Obsidian
date: "2021-09-30T20:51:22+02:00"
tags:
- obsidian
- knowledgemanagement
- tooling
---

If you’re somehow interested in productivity and knowledge management, it was getting really hard over the last couple of months to not hear about [Obsidian](https://obsidian.md) every couple of podcasts or blogposts. 

If you’ve never heard of it, Obsidian is a knowledge management application that you let loose on a folder somewhere in your filesystem containing Markdown files and it will not only offers you a nice interface to edit these files but also provide you with tools for visualizing the connections between them. 

## What held me back so far

The one major problem I've had ever since I first heard of Obsidian was the topic of how to synchronize my notes between various systems. The authors in general provide two options when mobile applications are involved:

1. Using iCloud
2. Going with Obsidian Sync which also offers end-to-end encryption

Given that I want to put *anything* into my notes system I *need* something that is end-to-end encrypted. Personally, though, there are only very few companies that I'd trust with getting something like that right in a closed-source context. Technically it's not that big of a deal depending on how your application operates but never mind that.

That's the main reason why I built my own knowledge management system, trying to recreate at least some aspects of Obsidian with just basic files and folders an some custom tooling around managing those files. These files are then shared through a simple Git repository that is hosted on a raspberry pi at my home and access to that from the outside world is only possible via Tailscale.

On 2021-09-29 I then saw [a tweet by the WorkingCopy author](https://twitter.com/WorkingCopyApp/status/1442499883026419716) about how WorkingCopy now basically allows me to put any accessible folder on my phone under version control. He also explicitly gave the example of Obsidian and so I tried to replicate that setup within my own environment, using once again my local RaspberryPI for hosting yet another Git repository and then synchronizing from all my mobile devices using [WorkingCopy](https://workingcopy.app).

At this point I've put a 3-digit number of notes from my custom system into Obsidian and really like how it performs. Obviously, there are some aspects that work differently but nothing big. What *is* big, though, it what Obsidian offers me that I wasn't able to recreate yet: 

- A proper mobile experience that invites me to also interact with my knowledge base on the go.
- A good high-level view over my notes that allows me to find existing connections between them and that also invites me to look for additional ones.

## What about my old system?

No, I will not ditch it but I will adapt it to  work with Obsidian vaults and the way linking works there. I will also get away from Zettelkasten-style note-IDs as they didn't turn out to be that useful for me and got in the way more often than not.

If things continue to work out that well, then I just might to move to Obsidian for the day-to-day work and fall back to my custom tooling less frequently. If not, then I haven’t wasted any time or data since the folder full of Markdown files also using Obsidian is just that: A folder full of Markdown files 😀
