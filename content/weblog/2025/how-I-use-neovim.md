---
date: "2025-09-22T10:49:03+02:00"
in-reply-to: https://lazybea.rs/vim-carnival-sept-2025/
tags:
- vim
- vimcarnival
title: How I use (Neo)VIM
---

This is a reply to Hyde’s VIM Carnival topic for September: [How do you use VIM?](https://lazybea.rs/vim-carnival-sept-2025/)

With a few interruptions, since my first contact with VIM in the early 2000s, it (and its derivates) has been my primary text editor for everything from writing code and documentation to blog posts. While some folks like to tune their experience with things like AI integration and fancy UIs, I prefer a minimal setup. The only thing I want in there (at least right now) is some nice syntax highlighting and code completion - everything else I like to do in a separate shell with command-line utilities.

Code completion especially made switching to [NeoVIM](https://neovim.io) a couple of years ago very attractive to me. All of a sudden I had more sane defaults out of the box *and* first-class LSP support without having to configure go all in on VIMScript. Nowadays I code mostly in Go, Python, and TypeScript. While my setup works well for the first two, I’ve so far been too lazy to also properly tune it for TypeScript outside of Deno. That’s pretty much the only time when I open VS Code.

I also write a lot of documentation and, again, I’ve kept a very minimal setup within VIM. One day, I will integrate a spell-checker and perhaps also try to integrate stuff like [Grammarly](https://grammarly.com). I'm not there yet, though. I think the only long-form text that I write outside of VIM is content for this blog here, for which I use Obsidian. Using VIM on a virtual keyboard is just too annoying while on the go  🙂

There you pretty much also have the things that I have on my TODO list regarding my setup:

1. Improve TypeScript support so that I can also use NeoVIM in general NodeJS projects
2. Add various tools to spell-check documents in a way that I can use at work and for personal projects