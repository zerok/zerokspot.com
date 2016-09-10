---
date: '2005-08-11T12:00:00-00:00'
language: en
tags:
title: Last.fm relaunched and new approach?
---


Yesterday <a href="http://last.fm">last.fm</a> was relaunched with a brand new design and many new features like tagging and journaling. They've also changed the way their radio streaming works, they now require your browser to handle the lastfm:// protocol which is quite easy to do in Firefox: Simply add following lines to your prefs.js:

-------------------------------



<pre class="code">

user_pref("network.protocol-handler.app.lastfm", "/usr/bin/xmms");

user_pref("network.protocol-handler.expose.lastfm", true);

user_pref("network.protocol-handler.external.lastfm", true);

</pre>



This is exactly how far I've got. When xmms gets started it doesn't know how to handle the path, so that's it. I really currently don't know if it's even possible anymore to use 3rd party players with last.fm or if we all have to wait for the official player to be released for Linux :-(