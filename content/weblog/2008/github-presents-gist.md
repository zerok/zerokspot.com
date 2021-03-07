---
date: '2008-07-22T12:00:00-00:00'
language: en
tags:
- github
title: Github presents Gist
---


<img src="http://img.skitch.com/20080722-psajjf1rbs2xhkf63pc2kbxqem.png" class="left" alt=""/>Think about what it would look like if a versioning system like git and a pastie had a child, and you'd probably came quite close what the folks behind the project hosting service [Github](http://www.github.com) [let loose](http://github.com/blog/118-here-s-the-gist-of-it) on the web just yesterday: [Gist](http://gist.github.com), a pastie that supports versioning and much more.


-------------------------------

There are always some apps people start writing when they want to familiarize themselves with a web framework or some other web toolkit. The first thing that comes to mind here is a simple weblog. ... another is a paste service, or pastie for short. There are currently so many of these little apps out there, it's hard to get an overview, and definitely not worth anybody's time to even try. Most of them were developed around a certain community -- list [dpaste](http://www.dpaste.com) for the Django community -- or a certain IRC network, but all of them share more or less the same featureset and usecase: Allow non-registered users to paste some code snippets to be shared with the community and later on forgotten.

<div class="figure"><img src="http://img.skitch.com/20080722-rjf3hkq8f8fxgeu4xpbkyjh3ey.png" alt=""/><p class="caption">Pastie + Git = Gist</p></div>

So I was quite close to a yawn when I read the heading of [this post](http://www.techcrunch.com/2008/07/22/github-unites-version-control-with-the-pastie/) on TechCrunch this noon. I guess the word "pastie" anywhere does this to me nowadays. Anyway, the post is about [Gist](http://gist.github.com), a new pastie service by [GitHub](http://www.github.com), which has quite a nice twist to it, or actually 2:

1. It allows you to version control a paste and therefor also allows stuff like forking etc. using git
2. A paste isn't restricted to just one file, but can consist of multiple files

And since the whole service is (naturally ;-) ) build around Github's git-infrastructure, you can also checkout each paste and work offline as with any git repository and then commit your changes back (Well, you can work freely up to [a certain point](http://logicalawesome.lighthouseapp.com/projects/8570-github/tickets/708-folder-kills-paste-on-gist), but anyway).

I personally think that Gist offers some really nice ideas and additions to the whole pastie-idea, but at the same time I have to wonder, who will really use those new features. At first when I heard "versioning + pastie", I thought that this would be great for a site like djangosnippets or any other "real" snippets site (compared to a pure pastebin) where you actually produce real code. For something like this, Gist lacks a way to find such snippets, like tagging or simply a title for a paste. If Gist aims to just be a pastebin, then I'm really not sure, if the whole versioning thing isn't an overkill for the service, but perhaps someone else finds a totally cool usecase for it :-)

Another question would be: Do your pastes affect your storage limit? Hmm....
