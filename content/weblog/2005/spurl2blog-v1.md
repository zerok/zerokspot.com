---
date: '2005-05-05T12:00:00-00:00'
language: en
tags:
- development
- python
title: Spurl2Blog v1
---


Ok, after testing it a little bit it's enough for myself so here it is: Spurl2Blog is a small Python script using PyXML and the builtin xmlrpclib of Python to fetch your latest spurls and post them to your weblog. The script "caches" the links included in the last fetch and tries to determine what links are new by iterating through the new links from oldest link to newest. As long as it only finds "old" links, the will be removed. The first "new" link stops this iteration and all "newer" links will be posted to your weblog. As you might guess, this cache is quite imporant, so I'd suggest that you never touch the storage/linkdump file (which is created after successfullly executing the script for this first time) unless you know what you're doing ;-)

-------------------------------



I'd also suggest that you first try the script on a demo weblog :-)



To configure this script (for example for setting your weblogs XMLRPC URL etc.) open the spurl2blog.py in a text editor and edit the appropriate lines. I tried to comment most of them so that it's easy to learn, what has to be changed and where. After configuring spurl2blog, simply run it with `python spurl2blog.py` in its folder :-)



Another short notice: I wrote this script primary for myself so if you ask me for help with it, I will try to help if I have the time. I can't promise anything. And I'm also not responsible if this script damages your weblog. I only tried it with WordPress 1.5 so far and it worked for me. This doesn't imply, that it will 100% work for you. I wrote this script for and under Linux. Thanks to Python and the work of the PyXML developers I <em>think</em> that this script should also work under Windows but I haven't tried it.



In the future I will perhaps add some features or fix bugs (which I'm quite sure exist in there ;-)). This updates (as long as they are for v1) will be released as patches.



<h4>Tip</h4> 

To run the script daily using "cron" is perhaps the best way ;-) For details please read the manpage of crontab :-)



<h4>Credits</h4>

I want to say "Thank you" to Jörg Kantel for describing how to use the xmlrpclib of Python in <a href="http://www.server-wg.de:8080/schockwellenreiter/webworking/weblogtool3/index_old_html/view">this post</a>



<a href="http://www.zerokspot.com/spurl2blog/spurl2blog.tar.bz2">Download</a>