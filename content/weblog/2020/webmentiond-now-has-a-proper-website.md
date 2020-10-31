---
title: 'webmentiond now has a proper website '
date: "2020-10-31T17:35:06+01:00"
tags: []
---

Ever since I decided I want to maintain [webmentiond](https://github.com/zerok/webmentiond) in the long run I wanted to have a small website for that project outside of the repository on GitHub. I had already purchased [webmentiond.org](https://webmentiond.org) some months ago but for the longest time it was just a redirect to the GitHub repository.

<figure><img src="/media/2020/webmentiond-website.png"><figcaption>Webmentiond.org website</figcaption></figure>

Last week then I tried a couple of systems with the goal to have the project's documentation as the center-piece of that website. I eventually picked mkdocs simply because it didn't require lots of changes to the structure of the docs folder and was easy to deploy. Updates to the website or just as any changes to the zerok/webmentiond repository and are deployed with the same CI workflow that's also updating the project's package and Docker image.

As always, if you find anything that could be improved about the project (code, documentation, ...) please create a ticket inside the GitHub project or use the webmentiond subreddit 🙂