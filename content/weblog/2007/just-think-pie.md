---
date: '2007-02-18T12:00:00-00:00'
language: en
tags:
- commandline
- perl
title: Just think PIE!
---


While I was listening to a recent [Lullabot Podcast](http://www.lullabot.com/audiocast/drupal_podcast_no_31_drupal_development_tools) someone meant that TextWrangler has soooo nice multifile string-replacement features. Well, definitely nice :-)

But if you don't have TextWrangler or simply prefer the commandline way of life, this might be quite handy:

<pre class="code bash">perl -p -i.old -e 's/lili/lala/g' test.txt test2.txt</pre>



-------------------------------



So if text.txt (and text2.txt) looks originally somehow like this:

<pre>lili
lili</pre>

Then it would after running the above command look like this:

<pre>lala
lala</pre>

The nice thing about this is, that thanks to the `-i.old` parameter, you even get backup files for all the originals.

3 parameters might be a p.i.t.a., but at least for me, thinking about pie might help here ;-)

Thanks to [Hubert Chen](http://programming.newsforge.com/article.pl?sid=06/03/08/1456241&amp;from=rss) for this one.

(I at least hope I haven't messed up anything here. I'm normally not using Perl for ... anything :-) )