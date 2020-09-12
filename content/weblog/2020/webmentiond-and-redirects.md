---
title: Webmentiond and redirects
date: "2020-09-12T17:43:27+02:00"
tags:
- webmentiond
- bughunt
---

For the last months I’ve been getting tons of mentions from [Hyde](https://lazybear.io/) as we comment on each other’s thoughts but something about them made [webmentiond](https://webmentiond.org) always mark them as “invalid” on my side. I was immediately quite certain that it had to have something to do with redirects as the mentions where coming from “http” while his blog usually serves content via HTTPS, but I simply couldn’t find the issue and, to be frank: The pain wasn’t big enough. I simply created the mentions manually 🙃

Yesterday, I again got two mentions and again they were marked as invalid but this time I went bug-hunting. The redirect counter had a little bug. Whenever webmentiond finds a link, it follows it before analysing the underlying document. In order to avoid endless redirects I added a little counter. If the number of redirects exceeds that counter, the mention is marked as invalid and we’re done.

There are some situations, though, where I simply don’t care about the number of redirects and in those cases I set the counter to `-1`. The bug was now, that I had forgotten to also handle that special value:

	client := &http.Client{}
	client.CheckRedirect = func(r *http.Request, via []*http.Request) error {
		if len(via) > cfg.MaxRedirects {
			return errors.New("too many redirects")
		}
		return nil
	}
	

Nothing in that little code snippet from `pkg/webmention/verify.go` tells the code to handle `-1` somehow different from any other value. Getting a redirect and the value is -1? Too bad! One is more than -1 and so it’s invalid. 

If you’re using a build that was made on 2020-09-12 or later, this should be fixed now. That being said, I actually don’t know why I wouldn’t care about the number of redirects and so I’ve now set it to 10 by default (you can override it using `--verification-max-redirects <VALUE>` if that value doesn’t work for you 🙂
