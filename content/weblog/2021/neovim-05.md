---
title: Neovim 0.5
date: "2021-07-16T19:35:20+02:00"
tags:
- 100daystooffload
- neovim
- vim
- development
---

A couple of days ago, [Neovim 0.5](https://neovim.io/news/2021/07) was released. If you’ve never heard of it, it’s a fork of the by now 30 year-old editor VIM (which itself is an improvement of vi which was released in 1976). Neovim now tries to add new features to the editor while staying to Vim wherever possible.

What makes VIM, Neovim, Emacs, and some other editors so special is that they can be completely customised. There are thousands of extends/modules/packages/whatever-you-call-them out there to make the editor fit perfectly to the way you want to work. Historically, VIM has used a language called Vimscript to write such plugins. I’m not a fan of that language. I tried to learn it multiple times and basically just couldn’t find any love for it which was one reason why I eventually moved over to Emacs. I want to customise my editor but I also want to have some fun doing it.

Neovim makes all that interesting again since plugins can be written in pretty much any language that can talk via msgpack RPCs. With version 0.5, Neovim went one step further and Lua is becoming more and more a first-class language within the editor *alongside* Vimscript. This means that you can not only write plugins directly within Lua directly accessing the Neovim API (no msgpack required) but you can also write your `init.vim` file in Lua and call it `init.lua`!  This was done by embedding LuaJIT directly with the project so a Lua compiler *is part of* Neovim.

With the enhanced support for Lua Neovim now also offers LSP functionality through a client library written in Lua.

The last of the major changes of 0.5 is that adds experimental support for a library called “tree-sitter” which will in the future allow syntax-tree based operations where previously the content of the buffer was only parsed with regular expressions.

All in all, I’m really excited about Neovim 0.5! So far I’ve tried to create a little `init.lua` configuration file based on my old config and had quite a bit of fun! The documentation is still IMO not as good in VIM as it is in Emacs but … I had fun 😅 Good thing that I’m using evil-mode in Emacs which is basically a VIM emulator, so I can use Neovim and Emacs without having to rethink everything all the time!
