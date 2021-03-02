---
date: '2006-11-05T12:00:00-00:00'
language: en
tags:
- applescript
- macos
title: Download right into the Finder
---


Well, no real story behind this. The motivation just came from the need to combine downloading [dl.tv](http://dl.tv) episodes and drag'n dropping them from Finder to [Disco](http://www.discoapp.com) ;-) ). So I wrote a little script that would download me a file right into the current folder in the front most Finder window. 

<pre class="code applescript">
on do_download(dl_file, dl_folder)
	tell application &quot;Terminal&quot;
		activate
		do script (&quot;cd &quot; &amp; dl_folder &amp; &not;
			&quot;&amp;&amp; wget &quot; &amp; dl_file &amp; &quot; &amp;&amp; exit&quot;)
	end tell
end do_download
tell application &quot;Finder&quot;
	set dl_folder to the front window&apos;s target as string
end tell
set dl_folder to quoted form of POSIX path of dl_folder
set dl_file to quoted form of (the clipboard as string)
display dialog &quot;Download &quot; &amp; dl_file &amp; &quot;?&quot;
if the button returned of the result is &quot;OK&quot; then
	do_download(dl_file, dl_folder)
end if
</pre>

First of all: I'm sorry if this code is a little bit clumsy. This was basically my first AppleScript :-) And as you can see in the source code: it requires __wget__. Probably ugly, but it works ;-)

Ah, and btw.: Make sure there is something usable in your clipboard ... like and URL. Otherwise starting this may mean the end of the universe. So basically and as always: Use this at your own risk. It does _no checking_ at all simply because I normally know what's in my clipboard ... at least during the 2 seconds it takes me at maximum to get to the icon for this script in my dock after having copied an URL to the clipbaord ;-)

If you want this and don't know, what to do with all this code:

1. Open up your ScriptEditor (or something like that. I only have the German addition of MacOSX.)
2. Paste the code into it
3. Save it as "Application"

Then you can simply drop it (for example) in your Dock to make it easily accessible whenever you need it. And if you don't have wget installed yet, get [fink](http://fink.sf.net) and `fink install wget` it :)
