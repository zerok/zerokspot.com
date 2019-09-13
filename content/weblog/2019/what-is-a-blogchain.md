---
title: "What is a Blogchain?"
date: 2019-09-13T21:14:27+02:00
tags:
- blogging
---

While reading through a few blog posts recently I came across a term I
wasn't familiar with: "Blogchain". With this post I want to dive bit
into some of the ideas that are related to this term.

## General idea

I first stumbled across the term in Tom Critchlow's post [Networked
Communities 2 - Blogging as a Social Act][tc1] and then, browsing
through the archives there, in [New Blogging 1 - Building
Blogchains][tc2]. From what I've read so far, the idea seems to be
similar to a post series but referencing each other and thereby
generating a trail of thoughts and content.

In a sense, every link that you put into a blog post is like a
predecessor of it. Combined with the date of each post you can then
see a clear causal connection between the posts.

From what I've seen on Tom's blog and the RibbonFarm their
implementation is more an explicit construct where every post is
explicitly marked as, let's say, the second post in a chain. This
would make their blogchains more like normal post serieses.

Let's see how Tom is actually implementing them by looking at the [most
recent post in the "Networked Communities" blogchain][tcnc3]:

```
---
layout: blog
title: Networked Communities 3 - Sidewalk Spaces and Positive Gatekeeping
subtitle:
redirect: https://www.brendanschlagel.com/2019/09/10/sidewalk-spaces-and-positive-gatekeeping/
blogchain: networked-communities
---

A holding post that should redirect to Brendan's blog.
```

In the [little snippet][tcnav] that he includes in the respective post
pages he then simply iterates through all the pages associated with
the same blogchain reverse-chronologically ordered and prints it.

I assume that the RibbonFarm does it in a similar way.


## Explicit linking

This means, though, that making a post part of a chain is an explicit
action involving setting the respective blogchain-attribute. The order
is then simply generated based on the time a post has been made.

But what would happen, if we don't do that explicit step? Linking
between posts should actually be enough to depict a
thought-chain. This would also allow us to not only have a single
parent-child path but one involving multiple parents, just as multiple
thoughts can come together to form a new one.

<figure>
<img src="/media/2019/blog-graph.svg" alt="">
<figcaption>Implicit graphs through explicit links</figcaption>
</figure>

Such directed graphs could be automatically generated and rendered but
how should they be named? Do they even need to be named? Both, Tom
Critchlow and the RibbonFarm name their chains. This has the
advantage, that you can (1) link to them and (2) that you can add some
form of description to them.

## Implementation details

This sounds actually like an entertaining project to implement here.
So, let's set a couple of constraints for a blogchain implemented
inside a Hugo-based website:

1. All elements of the chain have to be hosted on the same domain.
2. The order is purely defined by the links within posts.
3. Each post should display their predecessors and successors

As for making that whole thing addressable: Here we'd have two
options:

1. Either auto-generate within the script that builds the graph
2. Explicitly declare a name for a list of posts that should be
   considered part of a chain
   
<figure>
<img src="/media/2019/blog-graph-2.svg" alt="">
<figcaption>Posts attached to multiple sub-groups</figcaption>
</figure>

## Blog-chain or Blog-graph?

As each node can have multiple parents and itself influence more than
just a single follow-up post, I'm not sure I like the term "blogchain"
anymore. It's probably more buzz-word-compliant but, technically,
calling it "blog-graph" would be more appropriate.

Whatever the name may be, I think I like the general idea of
blogchains very much; especially so when they are done like on Tom's
blog with multiple contributors. I think I will try to implement the
more generic "graph" approach in the future here. If time permits 🙁
What I'll definitely do, though, is to crosslink more between posts
that definitely influenced each other in order to make thoughtchains
more visible.

Don't get me wrong, I don't claim this to be a new idea. To some
degree it's just the way the web works. Visualizing a thought trail on
a blog is something that is not common though and it should be fun to
implement!

[tc1]: https://tomcritchlow.com/2019/09/04/networked-communities-2/
[tc2]: https://tomcritchlow.com/2019/07/17/blogchains/
[rf]: https://www.ribbonfarm.com/2019/02/19/elderblog-sutra-4/
[tcnc3]: https://raw.githubusercontent.com/tomcritchlow/tomcritchlow.github.io/blob/c87736daeb4e69c8a945e5c90cb0690d9d39dc82/_posts/2019-09-10-brendan-blogchain-3.md
[tcnav]: https://github.com/tomcritchlow/tomcritchlow.github.io/blob/c87736daeb4e69c8a945e5c90cb0690d9d39dc82/_includes/blogchain-nav-top.html

