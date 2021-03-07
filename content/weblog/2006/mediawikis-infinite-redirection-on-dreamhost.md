---
date: '2006-12-20T12:00:00-00:00'
language: en
tags:
- development
title: MediaWiki's infinite redirection on Dreamhost
---


Just a small note for people trying to use MediaWiki on Dreamhost. I first made a local installation which had this as $wgArticlePath:

<pre class="code php">
$wgArticlePath      = "$wgScript/$1";
</pre>



-------------------------------



For reason, that I haven't yet found (and probably won't find thanks to other more important stuff on my list ;-) ) this doesn't work. It actually ends up in generating infinite URL redirections. 

After switching to the more ugly but at least working ...

<pre class="code php">
$wgArticlePath      = "$wgScript?title=$1";
</pre>

The infinite redirections finally seem to have stopped :)
