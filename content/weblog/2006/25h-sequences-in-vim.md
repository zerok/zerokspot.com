---
date: '2006-07-29T12:00:00-00:00'
language: en
tags:
- freebsd
- term
- vim
title: 25h sequences in VIM?
---


I'm currently evaluating FreeBSD as some kind of alternative to Gentoo for one of my machines I naturally also installed VIM right away. A little bit later I noticed a quite annoying problem: In the upper left corner the string "25h" which seemed to be the rest of some loose escape sequence. The same string also got appended to the rendering of any character I wrote.



-------------------------------



What caused it? Well, I also copied over my .vimrc from my PowerBook and this held following line:

<pre class="code">
set term=xterm-color
</pre>

Well, nice idea, but does only work if you're using xterm-color ;) So after resetting the terminal to "cons25" everything worked again, but it took me quite some time finding the real culprit. I already knew that it must have had something to do with $TERM (thanks to some excessive googling for a few hours), but I couldn't really find where I had set it incorrectly ;)