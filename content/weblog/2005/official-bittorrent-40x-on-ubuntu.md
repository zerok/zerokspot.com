---
date: '2005-07-11T12:00:00-00:00'
language: en
tags:
- ubuntulinux
title: Official BitTorrent 4.0.x on Ubuntu
---


Since Ubuntu Hoary only offers Bittorrent 3.x updating is a good idea if you want some of the newer features. So here comes a little step-by-step guide on how to get Bittorrent 4.x running:

-------------------------------



<ol>

<li>First download the Linux deb package from <a href="http://www.bittorrent.com/">bittorrent.com</a></li>

<li>Install it using <code>sudo dpkg -i $packagename</code></li>

<li>Ok, so Bittorrent is now installed. But there is a problem: It puts all the libraries into the python-2.3 paths. The problem is, that Ubuntu 5.04 uses Python 2.4 as default ;-) So open the /usr/bin/btdownloadgui.py in your favorite text editor and change the first line from

<pre class="code">#!/usr/bin/python</pre>

to

<pre class="code">#!/usr/bin/python2.3</pre></li>

<li>Enjoy ;-)</li>

</ol>