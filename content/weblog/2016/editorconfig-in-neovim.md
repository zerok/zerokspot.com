---
date: '2016-07-10T08:38:15+02:00'
language: en
tags:
- tooling
title: EditorConfig in NeoVIM
---

I'm currently playing a little bit around with [NeoVIM][] - not as a replacement
for Emacs but simply for some situations where I just need to open a file in a
shell, change a word and be done with it again. I've used VIM for that before
but I've heard great things about NeoVIM recently and so I thought I should give
it a try.

Anyway... whatever editor I use, it has to support `.editorconfig` files. No
excuses. I work on to tons of different projects with a lot of different coding
styles and [EditorConfig][] is the one thing that helps keeping me sane when
jumping from 2- to 4-space indentations all day long.

For VIM there exists an [official plugin][] that's also supposed to work on
NeoVIM, so I've added it to my `init.vim`:

```
Plug 'editorconfig/editorconfig-vim'
```

Turns out, there might be some compatibility issues between that plugin, my
globally installed Python, NeoVIM, and probably half a dozen other components:

> function <SNR>37_InitializePythonBuiltin[23]..provider#python#Call[9]..remote#host#Require[13]..provider#pythonx#Require, line 15
> Vim(if):Channel was closed by the client
> Failed to load python host. You can try to see what happened by starting Neovim with the environment variable $NVIM_PYTHON_LOG_FILE set to a file and opening the generated log file. Also, the host stderr will be available in Neovim log, so it may contain useful information. See also ~/.nvimlog.
> Press ENTER or type command to continue

Googling for solutions here only turned up some Python-setup related things that
I actually didn't want to get into.

Luckily, the plugin also supports a globally installed C-version of EditorConfig
which is what I wanted to use anyway. All you have to do is tell
editorconfig-vim about it:

```
let g:EditorConfig_exec_path = '/usr/local/bin/editorconfig'
let g:EditorConfig_core_mode = 'external_command'
```

That's it. After a restart nvim finally starts without the warning above and
also recognises my `.editorconfig` files out of the box 😊

[neovim]: https://neovim.io/
[official plugin]: https://github.com/editorconfig/editorconfig-vim/
[editorconfig]: http://editorconfig.org/
