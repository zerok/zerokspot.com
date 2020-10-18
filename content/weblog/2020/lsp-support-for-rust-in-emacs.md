---
title: LSP support for Rust in Emacs
date: "2020-10-18T16:33:16+02:00"
tags:
- rust
- emacs
- lsp
---

Thanks to [this post by Mike Stone](https://mikestone.me/revisiting-emacs) I opened up Emacs for the first time in months. Since I’m still somehow trying to make my way through the Rust curriculum on Exercism I wanted to get code completion in there. With VSCode I already had a working setup and so I started messing around with my init.el file.

Back when I stopped using Emacs I had a more or less functional configuration involving eglot but when I opened a Rust source files I got basically no completion. At first I thought this might be a problem with my eglot + company setup and so I switched to [lsp-mode](https://emacs-lsp.github.io/lsp-mode/). Only after checking with VSCode I noticed that the underlying [rls](https://github.com/rust-lang/rls/) didn’t completely all that much there either!

That was weird. I hadn’t noticed that back when working through the first couple of exercises but once I explicitly started looking it was hard to miss. So I started messing around with the config in VSCode, enabling [rust-analyzer](https://rust-analyzer.github.io/), and all of a sudden I got pretty the degree of completion I was expecting!

Recreating that functionality in Emacs would have been quite simple if not for my own stupidity. In case you don’t suffer from that, here is the simplified setup that I have with all the roadblocks already removed:

	(use-package rust-mode
	  :ensure t)
	
	(use-package lsp-mode
	  :ensure t
	  :hook (rust-mode . lsp))
	
	(setq lsp-prefer-capf t)
	(setq lsp-completion-provider :capf)
	(setq lsp-completion-enable t)

As I wrote above, it sadly took me a whole day to get to this. Let’s run down all the steps that contributed to this:

## 1. Download rust-analyzer into the path

The first thing I did was to download the rust-analyzer binary into my `$HOME/.local/bin` folder, assuming that this would be in my path. Turns out, it wasn’t. All I had was this:

	(setenv "PATH" "~/.local/bin:/Users/zerok/.cargo/bin:~/bin:/usr/local/bin:/usr/bin")
	(setq exec-path (append exec-path '("/Users/zerok/bin" "~/.local/bin" "/Users/zerok/.cargo/bin")))

Emacs doesn’t automagically expand `~` and so lsp-mode had no way to find the rust-analyzer binary in my path. All it did was report that it would use `rls` since that was in my PATH… 

## 2. lsp-rust-analyzer-server-command

Before checking my PATH settings, though, I also tried setting the `lsp-rust-analyzer-server-command` custom variable inside my init.el with something like this:

	...
	:custom (lsp-rust-analyzer-server-command "/Users/zerok/.local/bin/rust-analyzer")
	...

This didn’t help either which led to even more confusion and debugging. Was lsp-mode for some reason simply not looking at this variable? No, lsp-mode is fine! I just overlooked that this variable should be set to a list and not just to a string:

	...
	:custom (lsp-rust-analyzer-server-command '("/Users/zerok/.local/bin/rust-analyzer"))
	...

Given the settings for gopls and other language servers I had kind of assumed that variable to also be just a string but it was actually written in the documentation and all examples indicated that it would actually be a list. I just had overlooked this fact.

So, long post, short verdict: My struggles with LSP + Emacs + Rust were basically just a mix of unclear expectations (what experience I got from VSCode compared to Emacs) and lots of tiny bugs in my setup. As always, simplify your configuration to produce the smallest and simplest configuration that still shows the unexpected behaviour. This should make finding the bug either in your config or in whatever libraries you’re using much simpler! 😅

On the plus-side, this little trip down multiple rabbit holes at least helped me discover a little typo in lsp-mode's documentation and fix it, making this my first proper contribution for hacktoberfest 🙃
