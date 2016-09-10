---
date: '2006-02-28T12:00:00-00:00'
language: en
tags:
title: Getting Jerome's Keywords plugin to work with edgeio
---


[edgeio](http://www.edgeio.com) uses the RSS/Atom/RSS2 feed of your weblog to search for new listings, but since it currently just takes one and not necessarily the Atom feed your new items might not be recognized by the edgeio parser. The solution to this is actually quite simple:

-------------------------------



Search for following line around line 637 (in the keywords_appendTags function):



<pre class="code">

if ( (!$doing_rss) || ($feed != 'atom') )

</pre>



and replace it with



<pre class="code">

if ( (!$doing_rss) )

</pre>



Now every feeditem should have the tags appeneded to it and edgeio should finally be able to use them :)