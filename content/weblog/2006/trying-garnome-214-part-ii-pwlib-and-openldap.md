---
date: '2006-03-17T12:00:00-00:00'
language: en
tags:
title: 'Trying Garnome 2.14 Part II: pwlib and openldap'
---


The problem I've mentioned [here]() can easily [be solved by updating to an OpenLDAP 2.3.x release](http://www.mail-archive.com/garnome-list@gnome.org/msg02087.html). I also tried the other method mentioned there (changing the CXXFLAGS and CPPFLAGS) but this somehow didn't work. So let's head to the next problem ;)

-------------------------------



<pre class="output">configure: error: You must not compile firefox with the "typeaheadfind" extension enabled!make[1]: *** [configure-work/main.d/epiphany-2.14.0/configure] Error 1</pre>