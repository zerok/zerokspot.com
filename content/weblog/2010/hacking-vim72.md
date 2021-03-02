---
date: '2010-06-02T12:00:00-00:00'
language: en
rating: 4
tags:
- books
- review
- vim
title: 'Book review: Hacking Vim 7.2'
---


First a short **disclaimer** : I received a review copy of this book by
[Packt](http://www.packtpub.com/), which is really nice of them.

For the last couple of years I've been using
[Vim](http://vim.sourceforge.net/) more or less constantly without actually
getting all that deep into it. My whole configuration is a collection of (a)
stuff I found on the web and (b) what I could hack together myself after
looking at some examples and the documentation for less than an hour. For some
reason I simply never could find the time to actually read the documentation.

So when Swati Viswanathan of Packt Publishing asked if I'd like to review a
book about hacking up Vim, at first I was pretty undecided. I have virtually
no knowledge about the extensibility of that awesome piece of software except
for what is in my configuration. But then I read the description of [Hacking
Vim 7.2](https://www.packtpub.com/hacking-vim-7-2/book) by Kim Schulz and
thought this might be finally the right time for getting into it for real
;-)

-------------------------

As can be deduced from the title, this book isn't really a beginners' guide
(it probably would be called "Learning Vim" if this were the case) but instead
focuses on people who have a working knowledge of Vim but want to know more.
So don't expect a rehash of vimtutor.

There is quite a lot in there that I didn't know or simply didn't realize
about Vim. Basically that just about *everything* can be customised. From the
generation of the status line to just how much auto-completion support is in
there.

## The structure

The book starts with helping you personalize your experience with aspects like
the look and feel of the editing area by working with colour schemes and the
highlighting system as well as how tabs look like and how to extend GVim with
additional menus and buttons for frequently used actions.

Over the course of four chapters the author describes more and more
complicated customisations to improve your way to get around files and
projects up to setting up an automated coding style. And is some really
amazing stuff in here like how to customise autocompletion but also quite a
few things that I probably won't ever user (i.e. low-level session handling
for views). All these techniques are presented in a recipe-like format.

During these chapters the author most of the time implements those
customisations using small to mid-length functions which shows you that just
about everything can somehow be controlled using the scripting engine of Vim.
The scripting language itself and some addons to it (like the option to also
code in languages like Perl and Python) are described in chapters 6 and 7.
These offer a nice introduction and also show you methods to distribute your
own plugins.

I'm personally not really sure about the structure and the amount of chapters
customising received compared to the actual developer tools. But, I guess,
that's why this book has the word "Hacking" in its title and not "Programming"
;-)

At least with the eBook version there was also a bit of weirdness with
regards to the chapters about vimdiff which was for some reason placed within
the chapter about folding. So, yes, the book also introduces vimdiff, which is
something I only recently stumbled upon quite by accident :D

## Don't reinvent the wheel

That said, the book in just about every chapters refers to already
existing scripts provided by community members that improve on functionality
introduced there. I was delighted to even find a short tutorial for
[snipMate](http://www.vim.org/scripts/script.php?script_id=2540), one of my
favorite plugins.

A suggestion for the next edition of the book: The [MRU
extension](http://www.vim.org/scripts/script.php?script_id=521). It offers a
TextMate-like window for jumping between your most recently used files.

## GVim is not everything

But plugins are not the only diverse part of the Vim-ecosystem. It's also the
platform it runs on and the frontend it runs in. Sadly the author focuses
nearly exclusively on GVim for graphical frontends and ignores OSX as platform
altogether :-( So is `has("gui_running")` not really a reliable way to detect
GVim since it also matches MacVim which doesn't offer the same features as the
popular GUI.

---------------------

Concluding, if you want an introduction into what is possible with Vim beyond
what vimtutor has to offer, this is something for you. [Hacking Vim
7.2](https://www.packtpub.com/hacking-vim-7-2/book) offers some great tips on
how to customise your experience and also introduces you to its scripting
engine in order to publish your own extensions.

The book is kind of a combination of recipe-like tips and tutorials which
works rather well, although some more depth with some of the chapters would
have been nice, considering that the language extensions for Perl, Ruby and
Python got whole chapters, each.

If you're interested in working with Vim, I'd definitely suggest that you take
a look at the sample chapter.
