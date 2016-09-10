---
date: '2008-08-03T12:00:00-00:00'
language: en
tags:
- nerdtree
- plugins
- snippetsemu
- vim
title: MacVIM as TextMate replacement
---


[TextMate](http://macromates.com/) is one of the coolest editors around thanks to its Snippet-feature, ease of extensibility and community support. But more and more I saw myself using VIM instead of TextMate for one reason or another, but mainly because I'm just so used to searching with / and using :s for search-and-replace. And [MacVIM](http://code.google.com/p/macvim/) helped quite a lot too, thanks to its pretty nice integration into OSX.

Today I tried two plugins for vim that move it pretty close to what TextMate has to offer (not that vim is less powerful than TM or something, but some gems are just too well hidden ;-) ): snippetsEmu for snippets and NERD\_tree as a replacement for the project-drawer.


-------------------------------

## Snippets in VIM

In my quest for making VIM a true replacement for Textmate, I today gave snippetsEmu a try. snippetsEmu is a plugin that gives you more or less the power of snippets in Textmate. Pretty cool, I thought, but was appruptely stopped by the first hurdle: How to install it.

That plugin comes as a vimball bundle -- think .dmg for VIM plugins. To use this, you first require the [vimball plugin](http://vim.sourceforge.net/scripts/script.php?script_id=1502), which in turn requires a VIM version 7.1-299 or higher. And this is where the problem lies: Leopard comes with a veeeerry old VIM 7.0 release at patch-level 234. So the first thing you have to do, is install a newer VIM. This is pretty easy thanks to [MacPorts](http://www.macports.org/).

The next step then is, to install the vimball plugin. Just download it from the URL given above and extract it within your ~/.vim folder. This should put everything into the right places (.vim in the plugin folder etc.) and you should be ready to use vimballs from now on.

So now just download the 2 snippetsEmu files, "snippy\_plugin.vba" and "snippy\_bundles.vba", from [vim.org](http://www.vim.org/scripts/script.php?script_id=1318). Then open them with vim and load them using `:so %`. That's it. Now you should have the vimball installed.

If you're now editing let's say a .py file and want to use the snippets, the easiest way I've found so far is to execute `:emenu Snippets.python` which will basically just execute the menu entry for the Python snippets. In the end you'll probably want to bind the this to an AutoCmd ;-) After the snippets are loaded, snippetsEmu behaves mostly identically to what you've used to from TextMate. Just enter some text, hit <tab> and off you go. As in Textmate, you can also define certain variables within a snippet.

For details about how to define your own snippets, please take a look at `:help Snippet`.

## Filebrowser

Another neat feature of TextMate is the drawer that gets attached to the editor window once you've opened more than just one file or a directory. This is a pretty slick way to move between files within a project.

Right when checking out the snippetsEmu website, I noticed a screencast that showed a quite nice filebrowser for vim, which later on also got recommended to be my [Jannis](http://jannisleidel.com/), called [NERD_tree](http://www.vim.org/scripts/script.php?script_id=1658) by Marty Grenfell. This one doesn't only support browsing and opening files in the same window, but also supports opening documents in tabs, bookmarking files and so on. Really neat.

<div class="figure"><img src="http://img.skitch.com/20080803-r9tpiptakaxeb2r1pa2x846jum.png" alt="" /><p class="caption">NERD\_tree plugin for VIM in MacVIM</p></div>

I guess with a couple of shortcuts added, this might end up even more powerful than what TextMate has to offer. 

## What's still missing?

Now the only thing that I'm still missing is a good replacement for the CMD+t shortcut in TextMate, which let's you navigate through the files within a project sorted by the last-recently-opened file first. I guess you can achieve this with some search-path magic, but simply haven't found the time to really look into it so far :-)