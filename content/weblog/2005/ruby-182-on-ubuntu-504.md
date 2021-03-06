---
date: '2005-07-13T12:00:00-00:00'
language: en
tags:
- ubuntulinux
title: Ruby 1.8.2 on Ubuntu 5.04
---


Since yesterday there is a new version of Ruby On Rails available: 0.13.1. It seems like this version has introduced a version-check on Ruby which makes it hard to get Rails installed on Ubuntu 5.04. The problem here is, that Hoary uses a pre-release of Ruby 1.8.2 as base which was made 2 days before the final release.

-------------------------------



There are no real replacement packages available so I installed Ruby 1.8.2 (with all the security patches) manually gave its binary a higher priority in my path than the Ubuntu binary. I'm not sure so far if this breaks anything but I hope not :-) But first things first: I will now describe what exactly I've done :-)



<ol>

<li>Get Ruby 1.8.2 from <a href="http://www.ruby-lang.org">ruby-lang.org</a> together with the security patches</li>

<li>Extract the tarball and apply the patches</li>

<li><pre class="command">./configure --prefix=/opt/ruby-1.8.2

make

sudo make install</pre></li>

<li>Now let's make this new Ruby a brighter one: <pre class="command">export PATH=/opt/ruby-1.8.2/bin:$PATH</pre>

Simply put this in your own ~/.bashrc and the one of root.</li>

<li>After refreshing your environment installing gems and rails and using rails should be no problem anymore :-)</li>

</ol>



Since this appears to be a small bug in Ubuntu (having a pre-release from 2 days before the stable release) I still hope that someone will fix this also for Ubuntu 5.04 :-)



<https://bugzilla.ubuntu.com/show_bug.cgi?id=12613>