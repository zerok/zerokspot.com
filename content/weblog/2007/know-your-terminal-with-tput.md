---
date: '2007-03-04T12:00:00-00:00'
language: en
tags:
- programming
title: Know your terminal with tput
---


From time to time I have to write shell programs that should print big amounts of text to the terminal. In these situations it would be nice to know at least the horizontal dimension of the terminal to be able to mess for example with the indentation of the lines. So far I always simply took 80 (the default with of most terminal), but thanks to a recipe in the [_Python Cookbook_](http://www.oreilly.com/catalog/pythoncook2/) (Recipe 1.25 by Brent Burley and Mark Moraes) I now know the `tput` command :)


-------------------------------


While tput is a general tool for accessing an manipulating all kind of information about a terminal, one argument is especially important for me: "cols". A small example:

<pre class="shell">$ tput cols
80</pre>

Well, nothing new so far, but if you now change the dimensions of your terminal, tput's output will also change accordingly :D

Another very interesting entry in the terminfo database is "it", which tells you in how many spaces a tab is expanded to by default.

<pre class="shell">$ tput it
8</pre>
