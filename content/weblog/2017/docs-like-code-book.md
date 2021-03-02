---
date: 2017-05-19T20:14:05-07:00
title: "Docs Like Code: The book: A review"
tags:
- books
- review
- documentation
---

Docs-like-code is a relatively new term, but the idea behind it has been
implemented by the open source community for many years now: Writing
documentation using the same systems that are used for writing, reviewing, and
deploying code. From using version control systems
like [Git](https://git-scm.com/) or [Mercurial](https://www.mercurial-scm.org/)
all the way up to using continuous integration and deployment tools
like [Jenkins](https://jenkins.io/index.html)
or [TravisCI](https://travis-ci.org/).

[This book](http://docslikecode.com/book/) by Anne Gentle now tries to collect
all the best-practices around this method. Due to Anne's and her co-authors'
history with the [OpenStack project](https://www.openstack.org/), most of the
examples are based on that project. The book starts with a quick introduction of
the topic as a whole and then goes on covering the tool-chain from how to set up
repositories all the way to publishing the result
using
[Jekyll](https://jekyllrb.com/), [Sphinx](http://www.sphinx-doc.org/en/stable/),
or other tools. There are also large chapters about automated reviews, reviewing
in general, and continuous integration/deployment.

*Docs-like-code* mostly focuses on using Git, Restructured Text and Markdown for
the editing part and Github for publishing and collaborating. Sadly, the reasons
behind some of these choices and excluding other tools are not really explained
explicitly and the book therefore partially feels a bit biased. For instance,
AsciiDoc is only mentioned once with the tooling around it not receiving any
kind of coverage. Same goes for code hosting solutions other than Github. It
feels as if it were the only horse in town, with self hosted products
like [Gitlab](https://about.gitlab.com/)
or [Bitbucket](https://www.atlassian.com/software/bitbucket) only receiving only
the briefest of mentions.

Other parts of the book might be downright confusing to people new to the
method. Sometimes Git and Github are used interchangeably. Chapter 2 starts off
with a glossary on Git related terminology without introducing the actual topic
of that chapter ("Plan for docs like code") first. Terms like containers are
introduced without properly explaining them. Another time the author did suggest
rebasing in a chapter that should help a beginner make their first edits. As
rebasing is considered an advanced feature in Git this felt completely out of
place. Same as using vi for editing source files where the author went as far as
giving key-stroke by key-stroke introductions instead of just telling people to
use the text-editor of their choice.

A chapter that I would have expected earlier in the book was one about comparing
Wikis to documentation within a code repository. This should IMHO have been part
of the first chapter. But even down there were some weird points: The chapter
made it sound like you couldn't automate checks within a Wiki while systems
like [Confluence](https://www.atlassian.com/software/confluence) are extensible
enough to allow for these things. They are not easy, mind you, but usually
possible nonetheless.

There are also some technical aspects that are simply wrong. In one chapter the
author tries to introduce containers as part of the build tool chain. Good idea,
but in the same paragraph "micro services" are mentioned without a real reason
and containers are described as something that you need to have multiple of in
order to get an actual "service". There is also a chapter about programming
language considerations (since according to the author the primary language in a
project kind of dictates what documentation tool-set you should use) where eggs
are still described as the primary packaging format in the Python community.

In general, the book, sadly, kind of feels like an alpha version. The
copy-editing is thorough but the technical one is spotty at best and painfully
lacking in other
places. Perhaps
[lulu.com](http://www.lulu.com/shop/anne-gentle/docs-like-code/ebook/product-23064405.html) just
gave me an outdated version of the book but the one I got is simply nothing I
can recommend. As much as I'd love to have a book about the docs-like-code
approach that I can send around the office, this is not the one.

Big thanks to [Ulrich](https://twitter.com/ulope) for proof-reading 😊
