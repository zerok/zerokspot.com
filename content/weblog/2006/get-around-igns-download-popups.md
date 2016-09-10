---
date: '2006-08-24T12:00:00-00:00'
language: en
tags:
- ign
- scripts
title: Get around IGN's download popups
---


First of all, I'm a wget junkie. I think I haven't used the download manager in Firefox for quite some time now simply because with wget everything is simply "better". Ok, the major advantage is that it allows to resume downloads ;-) So since I'm now a quite pleased IGN reader, I also wanted to use wget there to download all these nice videos.



-------------------------------



If you have an Insider account (their way of saying "you give us money, we give you hidef-videos and additional content") you can straight away use wget for most of the downloads that are available to you. Most, because for some you have to be logged in since they want to know your age. And wget even makes this possible: Simply tell it, were your Firefox's cookies file is by putting something like this into your ~/.wgetrc:

<pre class="config">
cookies=on
load_cookies=/Users/zerok/Library/Application Support/Firefox/Profiles/som4funnyh45h.default/cookies.txt
</pre>

This would still leave you with that stupid popup IGN uses to give you the actual download link but also a stupid flash ad. And exactly this flash ad was what I want to get around since Flash9 for some reason is very prone to simply die everytime I'd actually mind. So to get around this, simply look at the link that is producing the video popup: It's basically a JavaScript function call with an URL in it. Well, so all your have to do, get that linked .html file and start parsing for the download link to the video. Boring ;-)

I've now written a very simple script that does exactly that using a little bit of wget and ruby magic. The usage is very simple: Download the file and put it somewhere in your $PATH (and make it executable ;-)). Then simply open a terminal and call this script with the JavaScript link you saw on IGN:

<pre class="command">
ign_movie_downloader.rb "javascript:popSizedWinProtected2('http://media.psp.ign.com/media/783/783630/dl_1643717.html','1643717',260,490,460,0)"
</pre>

Now you can download both, the normal and the insider content (the script will simply recognize a URL starting with http:// as insider content) using commandline scripts ;-) 

**Note:** This script uses only wget for downloading content, so it's not really all that platform independent. But this makes it much easier to integrate it into the cookies system of Firefox.