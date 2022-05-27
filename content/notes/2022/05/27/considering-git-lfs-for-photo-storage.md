---
title: "Considering git-lfs for blog photos"
inReplyTo: "https://lazybear.io/notes/re-photos-on-a-git-blog/"
date: 2022-05-27T10:26:58+0200
tags:
- blogging
---
Some time ago I used git-lfs for storing photos within the Git repository of my blog but it came with one downside while not really improving my overall workflow: You have to pay extra for LFS with GitLab and GitHub even if you have pro accounts on them.

One thing I also wanted to have improved in my workflow was that I could simply upload an image from my phone and have it auto-resized on the server into the profiles I need for my blog. The approach that I'm currently working on is just building a little API that does just that and having the photos backed up using Restic into GCP or something like that. With some luck this will go live this weekend unless I get distracted somehow.