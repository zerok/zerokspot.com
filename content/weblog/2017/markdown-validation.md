---
date: 2017-05-21T07:21:57-07:00
title: Markdown validation
tags:
- blogging
- markdown
- writethedocs
---

While I'm not the biggest fan of Markdown I've been using it for some time now
for blogging and to some degree note-taking. During the most recent WriteTheDocs
in Portland Lyzi Diamond gave a talk about automated doc-testing which sounded
like something I should have integrated into my blogging workflow a long time
ago but somehow haven't yet.

Up until a couple of days ago I was just writing my posts in Emacs, previewing
them in Marked 2, spell-checking them (if I don't forget about that) using
aspell, and finally publishing them using Hugo. Then, as written in my post
about [mobile writing](https://zerokspot.com/weblog/2017/05/17/mobile-writing/),
I've now also added iA Writer and Dropbox to my tool chain to bring new ideas
more quickly into a draft format. What's missing, though, is something besides
basic spell-checking.

In her talk, Lyzi mentioned [retext](https://github.com/wooorm/retext)
and [remark](http://remark.js.org/), two tools that allow you to process
Markdown files programmatically and thereby automatically. I'm now trying to
incorporate these into my workflow to check for common patterns I've noticed
creeping into my posts:

- Historically, I've been using tons of emojis. I should probably get that under
  control and only use at most one emoji per paragraph 😉
- I tend to write really long sentences that often become so complicated, that I
  have to re-read them multiple times in order to understand them. If I as the
  author have to do that, I can only imagine what it must be like for you as the
  reader 😒
- Sometimes my headings are just too short or too long. During WriteTheDocs
  Ingrid Towey said that you should normally have at least 3 and at most 12
  words in a heading. 2 to 10 sounds more reasonable for my use-case.

In the remark ecosystem exists a tool
called [remark-lint](https://github.com/wooorm/remark-lint) which allows you to
enforce style guides on your Markdown files. This sounds like the perfect place
to start. There are already tons of plugins and rule-sets available for it but
sadly nothing that matches the three things I want to check against in my posts.

Luckily, as it turns out, it's not that complicated to create custom rules if
you know some JavaScript, so I created these three packages:

- [remark-lint-emoji-limit](https://www.npmjs.com/package/remark-lint-emoji-limit) for
  making sure that I only use one emoji per paragraph at most
- [remark-lint-write-good](https://www.npmjs.com/package/remark-lint-write-good)
  which integrates [write-good](https://www.npmjs.com/package/write-good) into
  the validation tool-chain
- [remark-lint-heading-length](https://www.npmjs.com/package/remark-lint-heading-length) that
  checks that my headings at least have 2 and at most 10 words

Now I'm using this setup:

```shell
$ yarn add -D remark-cli
$ yarn add -D remark-preset-lint-markdown-styleguide
$ yarn add -D remark-lint-heading-length
$ yarn add -D remark-lint-emoji-limit
$ yarn add -D remark-lint-write-good
```

My `.remarkrc` contains the following settings:

```json
{
	"plugins": [
		"remark-preset-lint-markdown-style-guide",
		"remark-lint-heading-length",
		"remark-lint-emoji-limit",
		"remark-lint-write-good"
	]
}
```

Whenever I publish a post I simply run aspell and remark:

```shell
$ aspell check path/to/post.md
$ ./node_modules/.bin/remark path/to/post.md
```

This will give me a good indicator of what I should improve before publishing
the post. It's not a fully-automated process yet nor do I force myself to also
apply the suggestions made by remark et al., but it's a start 😉
