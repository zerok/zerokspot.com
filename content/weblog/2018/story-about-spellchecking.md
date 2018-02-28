---
title: "A story about spellchecking"
date: 2018-02-28T21:09:11+01:00
tags:
- writing
---

[About a year
ago](https://zerokspot.com/weblog/2017/05/21/markdown-validation/) I wrote
about how I'd love to test as much as possible about the content I produce
before pushing it out into the world. Perhaps the most important part of that
is spellchecking, for which I normally use [Aspell](http://aspell.net/). 

More recently, I needed to spellcheck some documents that should not feed back
into my main dictionary used for this blog. Because of that I started to look
into how Aspell can be customised to do just that and also to ignore certain
parts of [LaTeX](https://www.latex-project.org/) documents.

## Digging deeper into Aspell

Aspell allows you to use multiple dictionaries at the same time. For my
use-case the most convenient abstraction was using a base dictionary for German
in combination with a "personal" dictionary that would contain all the terms
specific to the project I was working on.

As I still wanted to have the base dictionary local the project, I copied the
one provided by Aspell itself into a local file:

```
$ aspell --lang=de dump master | \
  aspell --lang=de --encoding=utf-8 create master \
    ./aspell-de.dat
```

Next, I created an empty personal dictionary called `aspell-de.pws` with the
following content:

```
personal_ws-1.1 de 0
```

This header is important as Aspell would otherwise simply ignore that file.

Last but not least, I created a little configuration file with all the
customisations I wanted to have for the LaTeX mode:

```
encoding utf-8
master ./aspell-de.dat
personal ./aspell-de.pws
mode tex
add-tex-command graphicspath pp
add-tex-command autoref op
add-tex-command cellcolor op
add-tex-command columncolor op
add-tex-command bibliography op
add-tex-command bibliographystyle op
```

Voila! Now I can simply run `$ aspell --conf=$PWD/aspell.conf check
document.tex` to check the given file against the base and personal
dictionaries and skips content of things like autorefs.

One thing that took me far too long to notice, though, was that Aspell is
rather picky about how you specify the path to the configuration file. Always
provide an absolute path here to be on the safe side 🙂

On the downside, it looks like Aspell doesn't support German as well as I'd
like it to. For instance, it is not possible to whitelist terms that contain a
dash and there were also other weird tiny annoyances. Nothing major but still.

## Perhaps Hunspell?

Because of that I also thought about giving
[Hunspell](https://hunspell.github.io/) a try. It is used inside
OpenOffice/LibreOffice and should therefore incorporate even more experience
from all those office documents.

What I really like about Hunspell is that dictionaries are not just simple
wordlists but also include lexical rules and allow more precise handling of
inflections and other word-variants.

That being said, the core distribution of Hunspell comes without any
dictionaries. So the first thing I did was download a German one from [the
LibreOffice
repository](https://cgit.freedesktop.org/libreoffice/dictionaries/tree/de). The
main objective stayed the same: I want a base dictionary and an additional one
that should contain all the overrides collected while checking the target
project.

The configuration options available for a setup like this are rather similar:

```
$ hunspell -t -i utf-8 \
  -d ${PWD}/de_AT_frami \
	-p ${PWD}/hunspell-de.pws \
	document.tex
```

The main dictionary is specified using `-d` which should point to a pair of
`.aff` and `.dic` files. The first defines how words can be extended with
prefixes, suffixes, and other affixes. The syntax for that file is quite well
documented inside the project's
[manpage](https://github.com/hunspell/hunspell/blob/a7be9d32cd1b886e5334ff6ea8186bf5f3fe8118/docs/hunspell.5.md).
The latter contains the main wordlist combined with what affixes can be used
with each word.

`-p` once again allows the specification of a "personal wordlist" similar to
what Aspell has.

As you can see here, outside of specifying that the document is written in
LaTeX (`-t`) there are no options similar to Aspell's `add-tex-command`.
Outside of one or two cases, Hunspell did a good job out of the box.

Skipping the `-t` flag I also used pretty much the same setup to check this
document and Hunspell did a great job again. I will definitely give it a try
for the next couple of posts 🙂

