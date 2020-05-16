---
title: GitHub Actions to build this website
date: "2020-05-16T11:13:36+02:00"
tags:
- blog
- automation
- github
- ci
- 100daystooffload
---

Two days ago I finally managed to get my GitLab CI configuration to the point where it did everything that I wanted. It builds the site, generates the blogroll, updates the search indexes on Algolia, and finally sends out webmentions for all changed articles.

For the last weeks or so I spent more time in [GitHub Actions](https://github.com/features/actions), though, thanks to projects like [webmentiond](https://github.com/zerok/webmentiond) and [covid19-aut-stats](https://github.com/zerok/covid19-aut-stats/) which I have running on GitHub in order to make them more accessible for contributors (among other reasons). None of them has a particular elaborated workflow and so I thought I should try porting the build pipeline I have for this website over.

Turns out, the experience was much smoother than I had anticipated. The biggest challenge was to remember that I had to explicitly inject secrets using `${{ secrets.SECRET_NAME }}`. As you might expect, the [final workflow](https://github.com/zerok/zerokspot.com/blob/ff201c793e18f7bbea390a7736ee30c18a90525d/.github/workflows/main.yml) looks pretty similar to what I had on GitLab and they also take about the same time: ~ 4m on both systems.

To get there, though, I had to fine-tune one particular aspect of the GitHub workflow: Since I pass the `public` folder that Hugo generates, around between dependent jobs, I use the [upload-artifact  action](https://github.com/actions/upload-artifact) quite a lot. Turns out, this one really doesn't like working with that many files and so I'm now packaging the folder before uploading it:

	    - run: tar -cJf public.tar.xz public
	    - uses: actions/upload-artifact@v1
	      with:
	        name: public.tar.xz
	        path: public.tar.xz

This small change has speed up the upload from 3m34s to 28s + 18s. Sure, compressing and decompressing adds some overhead, but 46s is still far better than 214s 😉 I'm pretty sure, I'll be able to get a couple more seconds out of the whole workflow once I feel more familiar with GitHub Actions.

This also means, that for the time being, [https://github.com/zerok/zerokspot.com](https://github.com/zerok/zerokspot.com) is the canonical repository of this website 
