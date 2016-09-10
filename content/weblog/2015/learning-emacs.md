---
date: '2015-01-04T19:03:30+01:00'
language: en
tags:
- emacs
title: Learning Emacs
---


[Emacs][e] and I have a weird history. It was in fact the first editor I used to
learn Java back in 2001 thanks to [Michael Kropfberger][mk] who suggest it to
his first-year students. But after just a week or so and without getting beyond
opening, saving and closing files I moved to other tools like Eclipse and VIM.

A couple of weeks ago I wanted to give it another try especially because I
wanted to mess around with Lisp-languages and having an application right out of
the box is the best way to learn anything ;-) I was also a little bored of my
current tool-chain and mixing that up is always fun!

-----------------

Usually, the first thing you hear when you tell someone you want to learn
something like VIM or Emacs is to get a pre-configured distribution like
[Janus][j] or [Prelude][2] respectively. Back when I wanted to learn more about
VIM I already ignored Janus and I think it helped me quite a lot. For me it is
simply very important to learn stuff like it was originally intended and then
slowly move to more modern settings if and only *if* I see the benefit for
me. Your mileage will definitely vary here, though. I'm a bit stubborn and
complicated in this regard ;-)


## OSX integration

One of the things that prevented me to play around with Emacs again earlier is
the keyboard layout Apple created for OSX. The META key in Emacs is extremely
important but on OSX it maps onto the alt/option key by default. This is
probably fine for the native US keyboard, but in the German and US-International
layouts that key is immensely important for things like German umlauts or
dead-key combinations when writing something in French.

After a bit of experimenting with various settings I ended up binding the META
key onto OSX's CMD key.

```
(custom-set-variables
  '(ns-command-modifier (quote meta)))
(setq mac-option-modifier 'none)
```

This has the big disadvantage that I can no longer CMD+C/V for interacting with
the clipboard. Luckily, as with VIM, you can [configure Emacs][c] to also use
the OS' clipboard:

```
(setq x-select-enable-clipboard t)
```

No, if I copy something in Chrome I can paste it into Emacs with Ctrl+y and copy
stuff out with META-w.


## Learning ...

Since re-learning some Lisp and Emacs play hand in hand for me I picked
["An Introduction to Programming in Emacs Lisp"][1] as my tour guide. It's
probably not the best book if you've seen Lisp before (yeeeears ago) or know
your way around other languages but for my purpose it was good enough.

Over the Xmas holidays I finished the book and can now at least read and to a
very limited degree write Emacs' configuration files, which is pretty much what
I wanted to achieve here :-) I'm not sure if I like the language enough to also go
beyond that but who knows...

That being said, the documentation available for any function available in Elisp
is simply awesome!


## The "killer-feature"

Another side-effect of this little adventure was that I could finally play
around with the famous [OrgMode][3]. I won't go as far as saying that you should
use OrgMode even if you don't use Emacs for anything else, but it gets very
close. This is probably the single best outlining tool I've seen so far. I'm not
sure if I will use it for some of the more advanced areas like GTD and project
management (simply because I like to have those also available on the go) but
for outlining I will definitely use it in the future.


So, after 3 weeks do I like Emacs more than VIM? No, simply because I can't really
compare these two tools. VIM feels to me with its motions like the better
editor but Emacs seems more powerful outside of that. A killer-feature of VIM is
also that it's available on any server I maintain, so even if I switch to Emacs
on my working machine VIM will still see a lot of use whenever I ssh into
anything. That being said, I really enjoy using Emacs and can definitely see
myself sticking with it at least for a couple more weeks.

A couple of other resources I found useful during the last couple of weeks:

* [Brian Carper's dotfiles](https://github.com/cdaddr/dotfiles)
* [Ryan McGeary's dotfiles](https://github.com/rmm5t/dotfiles)
* [Emacs Redux](http://emacsredux.com/)


[1]: http://www.gnu.org/software/emacs/manual/html_node/eintr/index.html
[2]: https://github.com/bbatsov/prelude
[mk]: http://www.kropfberger.com/
[3]: http://orgmode.org/
[e]: http://www.gnu.org/software/emacs/
[c]: http://stackoverflow.com/a/64558/22312
[j]: https://github.com/carlhuda/janus
