---
date: '2004-06-19T12:00:00-00:00'
language: en
tags:
- firefox
- flash
- linux
- software
title: Firefox 0.9 & Flash (Linux)
---


After some days I've now finally got MozillaFirefox to display Flash content. A big "Thank you" to <a href="http://forums.gentoo.org/viewtopic.php?t=186600&amp;highlight=firefox+flash"> cappaberra</a> on this occasion. There where probably some pieces of an old Flashblocker-Extension or something still flying around in my profile that blocked the flash plugin. Here now a short guide on how to install the flash plugin for those who also have this problem.

------------

<ol>
<li>Download the Flash plugin from macromedia.com and extract it.</li>
<li>There should now be a new folder which contains a<em>flashplayer.xpt</em> and a <em>libflashplayer.so</em>-file. Put the <em>libflashplayer.so</em> into Firefox's <em>plugins/</em> folder and the other file into <em>components/</em> folder.</li>
<li>Now comes the step where I previously failed: In your profile's folder there should be a  <em>chrome/</em> and an <em>extensions/</em> folder that are used for storing extensions et al. The simplest way would be to delete these two folders and then reinstall your extensions :-)</li>
</ol>

## Links

<ul><li><a href="http://www.macromedia.com/shockwave/download/download.cgi?P1_Prod_Version=ShockwaveFlash">Flash plugin</a></li>
<li><a href="http://forums.gentoo.org/viewtopic.php?t=186600&amp;highlight=firefox+flash">Topic about this problem on forums.gentoo.org</a></li>
</ul>
