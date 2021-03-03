---
date: '2011-09-25T12:00:00-00:00'
language: en
tags:
- vim
title: 'UltiSnip: Snippets for VIM'
---


A couple of days ago I stumbled upon [UltiSnip][ultisnip], a TextMate-like
snippet system for VIM, via Alexey Bezhan's post [Why UltiSnip?][why]. For
quite some time now, I've been using [snipMate][snipmate] which does
essentially the same but seems to have been abandonned by its original
maintainer. There *is* a [fork][snipmate.garbas] maintained by [Rok
Garbas][garbas] but, I guess, I was ready to try something new that day and so
I removed snipMate from my vim config and installed UltiSnip instead.

-------------------

So far, this works out pretty nicely for me. One of the issues I had with
snipMate was how reloading snipppets was done, which, in its original
implementation, wasn't. With UltiSnip you basically edit a snippets-file and
your changes work right away. I also prefer the different syntax for creating
snippets which doesn't rely on indentation but has an explicit closing tag.

<pre><code>snippet msg "Shortcut for spring:message" i
&lt;spring:message code="${1}" /&gt;${0}
endsnippet</code></pre>

You can also define a handful of options with each snippet which, for
instance, restricts the snippet to only work at the beginning of a line or
doesn't require the trigger to be surrounded by whitespaces.

Another really nice feature is that you can have one snippet library extend
another one. For instance, you probably want to use all your HTML snippets
within a JSP file. So all you have to do is add following line to your
jsp.snippets file:

<pre><code>extends html</code></pre>

Right now I've only scratched the surface here and I'm still trying to find
the time to read through the documentation as a whole, but so far I really like
UltiSnip :-) If I have enough custom snippets I will probably also create a
little repository for them and put them online.

[why]: http://fueledbylemons.com/blog/2011/07/27/why-ultisnips/
[snipmate]: http://www.vim.org/scripts/script.php?script_id=2540
[garbas]: https://github.com/garbas
[ultisnip]: http://www.vim.org/scripts/script.php?script_id=2715
[snipmate.garbas]: https://github.com/garbas/vim-snipmate
