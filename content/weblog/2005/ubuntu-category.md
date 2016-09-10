---
date: '2005-01-06T12:00:00-00:00'
language: en
tags:
- ubuntulinux
title: Ubuntu category
---


I've now also added an Ubuntu category simply because I'm currently very interested in this new Debian/Gnome based distributions and have also installed it on my laptop. I'll perhaps write quite a few posts about it, so this will should keep my weblog tidy :-)

-------------------------------



<a href="http://weblog.zerokspot.com/posts/151/">As I've already written before</a> I've switched to the Hoary branch of Ubuntu. This was a week ago and I really like it there. I've noticed no real problems so far (except for not being able to alter the menus because of the gnome-menu switch and a dead desktop right after booting, which can be solved quite easily but more on this later) which came as a suprise to me. To be honest, I expected more problems running the development (and pre-beta) branch of Gnome, but so far, so good :-)

As for the blank desktop problem I've mentioned above: This seems to be caused by some problems with gnome-vfs in Gnome 2.9.2. Nautilus doesn't start, the gnome-panel stays empty etc. I've solved this for me with adding an xterm instance into the gnome-session and then executing following:
<pre class="code">killall trashapplet gnome-vfs-daemon nautilus gnome-panel ; nautilus</pre>
If you have some more gnome-vfs depending applets running you perhaps want to kill them before gnome-panel too. 