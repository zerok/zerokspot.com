---
date: '2006-08-24T12:00:00-00:00'
language: en
tags:
- python
- youtube
title: 'youtube-dl: Downloading with Python magic'
---


Ricardo Garcia Gonzalez has published another tool for downloading videos from YouTube. [youtube-dl](http://www.arrakis.es/%7Erggi3/youtube-dl/) is written in Python and a commandline script (day of the commandline tools today, ey?) (Thanks [William Pramana](http://wpram.com/log/2006/08/20/download_youtub/) for that link). Besides actually downloading the video, it also has an option for only printing the video's real URL on YouTube, so that you can start the download process with whatever download manager you like. The only disadvantage here is, that the script also prints some text alongside the actual URL, so it's not you usable for some piping action. Not all that hard to patch that.



-------------------------------



When you use youtube-dl after applying this patch, you can download for example a video using something like this:

<pre class="command">
youtube-dl -q -s $url | xargs wget
</pre>

So the patch simply prints the URL even if the script was configured to be quiet, so quiet now only supresses the metadata but nothing else.

**Note:** As the patch's name already indicates: This is for the 2006.08.15 release of youtube-dl.