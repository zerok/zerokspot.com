---
date: '2011-09-19T12:00:00-00:00'
language: en
tags:
- textmate
- chocolatapp
- vico
- sublimetext2
title: TextMate's heirs
---


Back in 2005 the sweet topping of the cake called "Switching to OSX" for me
was [TextMate][tm]. As many back then, also I was amazed by every single
screencast that showed the power of snippets, column selection and all these
other nice little features the editor created by [Allan Odgaard][ao] had
created.

But sadly, after version 1.5 (which had been released in 2006) the core
project completely stagnated. And since then every now and then a post about
TextMate2 comes up (with the latest by Allan himself on the [official
mailing-list](http://lists.macromates.com/textmate/2010-May/030942.html) back
in May 2010). The community is still active around TextMate 1 but without
*anything* to show that TextMate2 is something worth waiting for more and more
projects are starting to pop up trying to mix the good of Allan's editor with
new features or simply missing features like split-views.

While many people (including myself) moved over to VIM/MacVIM or Emacs, these
are not the two I want to write about here, but about editors that explicitly
try to replace TextMate. So far I was able to give the following three a try:

* [Sublime Text 2][st2]
* [Chocolat App][ca]
* [Vico][vico]

I've mostly used Sublime Text 2 over the last couple of weeks but also looked
at the other, so I apologize if this post focuses a bit too much on the first
one. That said, this post is by no means a review or an exhaustive comparison
of these three editors. At least the first two are not even stable yet so this
is more like a quicklook than anything else.

-------------------------------------------------------------------------------

## Sublime Text 2

First a bit about [Sublime Text 2][st2] (ST2), which is currently in beta. I
can't really say anything about it's "stable" predecessor but version 2 has
been running quite smoothly for the last couple of weeks and I really enjoy
working with it. It offers some of the core features that, back in 2005, brought
me over to TextMate:

* column selection
* snippets
* extensibility

... and so far I was able to put nearly every TextMate bundle I wanted to use
into the "Packages" directory and it just worked. And compared to TextMate it
isn't Mac-only. In fact it was only ported to the Apple platform earlier this
year and is now available on Windows, Linux and OSX.

Another feature of ST2 that is very appealing to me is it's API which makes it
really simple to write plug-ins for if you know Python. Just last weekend I
wrote my first little plug-in for it and I was able to do that within an hour
or so without any prior knowledge of the API or infrastructure. For debugging
you can simply open a Python console within the editor where, for instance,
all the prints in your plug-in go to.

<figure>
    <img src="/media/2011/st2-devconsole.png" alt="" />
    <figcaption>The dev console is actually an interactive Python 2.6.x
    interpreter</figcaption>
</figure>

But perhaps the single most important feature of ST2 and improvement over TM for me is its
handling of settings.  You can set separate configuration options depending on
what project and what file-type you're working in. For
instance: Normally I have my CSS/SCSS files, JavaScript files etc. set up to
have lines no longer than 80 chars. But at work we have slightly more relaxed
coding guidelines where anything up to 120 chars per line is just fine. So for
my work projects I simply set the line length to 120 and just use every other
setting from my global settings files. You can even set multiple vertical
rulers if you want ;-)

<figure>
    <img src="/media/2011/st2-multiple-rulers.png" alt="" />
    <figcaption>You can even define more than one ruler</figcaption>
</figure>

The one thing that initially made me move back to VIM from TM a couple of
years ago was the ability to split the window into multiple views and have a
single file or multiple files open side by side. And guess what: ST2 can do
that too ;-) Not as flexible as VIM, but probably well enough for most people.

<figure>
    <img src="/media/2011/st2-layout-modes.png" alt="" />
    <figcaption>ST2 supports multiple layout modes to render multiple
    files side by side within a single window</figcaption>
</figure>

Right now ST2 has a few little issues like that the whole project management
from an UI standpoint doesn't really work for me but, everything considered,
this looks like a really fine editor and I personally can't wait for it to hit
a stable release. It brought me really close to the point of dropping VIM ;-)

## Chocolat App

If you want something with a bit more minimalistic feel to it, [Chocolat App][ca]
(CA) seems another really strong contender. It is probably unfair to compare
CA with ST2 or TM right now given its early stage of development, but so far
this one looks pretty good too :-) 

At time of writing this it is at version 0.0.36 (alpha stage) and therefor
lacks a large amount of features like column selection and an official API,
but thanks to the public [issue tracker][ca_issues] on Github and the [IRC
channel][ca_irc] on Freenode development is extremely transparent and its hard
not to get excited.

Compared to ST2, Chocolat App tries to solve things like configurations more
like TextMate and relies right now on a GUI quite similar to the older editor.
Actually, the preferences windows, for me taste, looks a bit too similar but
I'm pretty sure that this will change over them.

As with TM there is also a bundle editor that looks also looks very much much
like TM's and doesn't really offer any of the main editor's features in its
text-areas either. I just hope these text-areas will eventually be replaced by
instances of Chocolat itself for things like syntax highlighting et al. when
editing a plug-in/command written in something like Python or Ruby.

<figure>
    <img src="/media/2011/choc-library.png" alt="" />
    <figcaption>Chocolat sports a bundle editor similar to the one provided by
    TM</figcaption>
</figure>

Splitting views is right now a bit awkward with only vsplits being possible and
[no apparent way to actually close the split again][3].

All things considered, Chocolat App looks really promising. It basically takes
UI elements of TextMate and improves upon that feature-set with things like a
new bundle installer and split-able views. Right now, though, it's not really
usable enough to out-do it's "competition". That said, I will definitely keep
my eye on it and I'm pretty sure there is much greatness to come :-)

## Vico

[Vico][vico] is another really nice app that appeared out of nowhere in
[April][vico_init] and is by now at version 1.2. It is basically a mix of
TextMate and VIM + perhaps a bit of XCode with all its search-panels:

<blockquote>
<p>Vico is the result of a personal itch. My eyes want a beautiful looking, modern Mac text editor. But my fingers just want vi.</p>
<quote><a href="http://blog.vicoapp.com/2011/04/Introducing-Vico">Martin Hedenfalk</a> in the initial announcement</quote>
</blockquote>

The TextMate part comes into play with a "go-to-file" shortcut and support for
TM bundles. I have no idea, though, how much of the bundle API is actually
supported, though. But I'd guess snippets and syntax files are a safe bet.

Similar to TextMate's project handling Vico also offers some status icons in
the project drawer for when a file has been edited as has been the case with
the setup.py file in the picture.

<figure>
    <img src="/media/2011/vico-modified.png" alt="" />
    <figcaption>Vico's sidebar gives some indication on what files in the
    project have been modified.</figcaption>
</figure>

Regarding VIM look & feel Vico brings the classic hjkl navigation and hitting
":" even opens a command-line as in VIM's normal mode.

<figure>
    <img src="/media/2011/vico-commandline.png" alt="" />
    <figcaption>Hitting ":" even opens a commandline.</figcaption>
</figure>

While this all sounds great, I'm not so sure that sticking to close with that
classic editor is a good idea. If I want VIM, I want VIM with all its numerous
scripts and really every mode. Having the visual mode emulated is fine but one
out of five times I also need the visual block mode. This kind of breaks the
illusion that the rest of the navigation system provides.

I guess, Vico just isn't for me.

## What about others?

I also at first considered adding Kod to this list sadly, 0.0.3 is simply too
far away to be usable and with the [last commit][2] to the project at this stage
being more than a month ago, it looks kind of dead too :-(

## One-man-shows

One thing that's kind of scaring me off a little bit with at least Sublime
Text 2 and Vico is that they are both one-man shows. Judging form the outside
this was what basically killed TextMate and what makes pure open-source
editors so attractive despite any shortcomings they might have. With them at
least you can be sure, that if the current maintainer doesn't want to improve
the app anymore and *someone* is interested in the project, there will be a fork.

Chcoolat App is the exception here because there were from the get-go two
developers involved.

TM demonstrated quite well that even a seemingly dead project can live on
through sheer power of the community who still puts quite a lot effort into
extensions and bundles but in the end even that, in my opinion, can only help
to some degree if the actual core gets completely outdated. Luckily it
probably hasn't happened with TM yet. There even is [support for Lion's
fullscreen][tm_fullscreen] feature provided through a 3rd-party plug-in ;-)

In the end, I hope that this amount of competition will create a lot of new
ideas for how editors might or even should help a developer's work-flow. TM
showed back in 2004 that there is still quite a lot of room for innovation and
I highly doubt this room has been filled up completely by now :-)

[st2]: http://www.sublimetext.com/2
[ca]: http://chocolatapp.com/
[ca_issues]: http://github.com/fileability/chocolat-public/issues
[ca_irc]: irc://irc.freenode.net/%23%23chocolatapp
[tm]: http://macromates.com/
[tm_fullscreen]: https://github.com/enormego/EGOTextMateFullScreen
[ao]: http://twitter.com/sorbits
[vico]: http://www.vicoapp.com/
[vico_init]: http://blog.vicoapp.com/2011/04/Introducing-Vico

[2]: https://github.com/rsms/kod/commit/6043dcb673c3431e5f4cbcad3e3254560f7ca3bc
[3]: https://github.com/fileability/chocolat-public/issues/295
