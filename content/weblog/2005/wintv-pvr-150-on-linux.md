---
date: '2005-06-12T12:00:00-00:00'
language: en
tags:
title: WinTV PVR-150 on Linux
---


Part of my new PC is a WinTV PVR-150 which I bought for recording TV shows etc. since I don't have a recorder in my room and also don't plan to buy something like that. I've actually also come quite far in getting it to work but now I'm stuck... But first a small description of what I've done until now:

-------------------------------



First of all I got the ivtv drivers here and installed them with

<pre class="command">cd driver

make

make install

cd ../utils

make 

make install</pre>



Then I inserted the driver CD that was bundled with the card and copied the HcwFalcn.rom   and the HcwMakoA.ROM into /lib/modules. I also symlinked HcwFalcn.rom to ivtv-fw-enc.bin as was described <a href="http://ivtv.writeme.ch/tiki-index.php?page=PVR150-500Firmware&highlight=firmware">here</a>



Nothing really interesting here only that if you have the eeprom module loaded I'd recommend that you unload it before loading the ivtv module. I don't know if it really causes any problems but it at least spares you a few warning/error lines in dmesg :P



Now comes the part where the module gets loaded with <pre class="command">modprobe ivtv ivtv_std=2</pre>

The ivtv_std parameter sets the TV standard used by the driver (2 for PAL). 



This is basically where I am now. When I try to view the output of the card with <code class="command">mplayer /dev/video0</code> most of the screen is black (sometimes with a little snow ;-) ) and the rest has green fields all over it. I also went to a friend who installed the card on his Windows machine and got the same result before switching from the "cabel" mode into the "antenna" mode. My problem is now: How can I do this with ivtv? I've search the last couple of days and couldn't find a solution for this :-?