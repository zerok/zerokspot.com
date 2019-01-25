---
title: "Emacs: Back and Forth"
date: 2019-01-25T18:17:59+01:00
tags:
- emacs
- productivity
- setup
---

Nearly exactly 4 years ago I wrote [a little post](https://zerokspot.com/weblog/2015/01/04/learning-emacs/) about why I wanted to learn [Emacs](https://www.gnu.org/software/emacs/) again. Late last year I then moved back to VIM using [Neovim](https://neovim.io/) simply because I was working more often than not on remote servers and therefore wanted to reuse the same VIM-configuration no matter where I was opening files on.

{{<figure src="/media/2019/emacs-screenshot.png" caption="Emacs with Magit">}}

While I was using [evil-mode](https://github.com/emacs-evil/evil) for my Emacs setup, it simply wasn't the same. There were also some situations inside Emacs where working with evil-mode kind of broke apart, which further motivated me to give a pure VIM setup a try again. For some reason, though, (or I was simply missing magit) I couldn't really let go and now, after just a couple of months, started working on a complete fresh start with Emacs using a plain configuration file without (or at least disable by default) evil-mode.

So far, I've only enabled a handful of packages:

* [magit](https://magit.vc/)
* [dockerfile-mode](https://github.com/spotify/dockerfile-mode)
* [go-mode](https://github.com/dominikh/go-mode.el)
* [dracula-theme](https://draculatheme.com/emacs/)
* [company-mode](http://company-mode.github.io/)
* [ivy + swiper](https://github.com/abo-abo/swiper/)
* [avy](https://github.com/abo-abo/avy)
* [projectile](https://github.com/bbatsov/projectile)

Since I'm still working quite a lot on remote servers, I will also try to get to know Emacs' [tramp](https://www.gnu.org/software/tramp/) better. Last time around I mostly forgot about it. But first I have to get to grips with not falling back to my VIM muscle-memory 🤪 So far, I'm still struggling a bit with using bindings like "C-n" and "C-p" to move from line to line and not typing "v" every time I want to select *anything*. 

No matter what my future setup will be like, I also want to add it again to my [dotfiles repository](https://github.com/zerok/dotfiles) which I've neglected for far too long.



