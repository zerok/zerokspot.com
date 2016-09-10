---
date: '2005-04-01T12:00:00-00:00'
language: en
tags:
- web
- wordpress
title: Wordpress.org lost google page rank
---


I don't know all the facts yet, but it seems like google removed Wordpress.org's page rank after violating some policy which could be interpreted as page rank spamming. 

-------------------------------

Wordpress.org has following section in its source code which seems to be the stumbling block:

<pre class="code">&lt;div style="text-indent: -9000px; overflow: hidden;"&gt;
&lt;p&gt;Sponsored 
&lt;a href="/articles/articles.xml"&gt;Articles&lt;/a&gt; 
on &lt;a href="/articles/credit.htm"&gt;Credit&lt;/a&gt;,
 &lt;a href="/articles/health-care.htm"&gt;Health&lt;/a&gt;,
 &lt;a href="/articles/insurance.htm"&gt;Insurance&lt;/a&gt;,
 &lt;a href="/articles/home-business.htm"&gt;Home Business&lt;/a&gt;,
 &lt;a href="/articles/home-buying.htm"&gt;Home Buying&lt;/a&gt; 
and &lt;a href="/articles/web-hosting.htm"&gt;Web Hosting&lt;/a&gt;&lt;/p&gt;

&lt;/div&gt;</pre>
(Reformated to be better readable)

This code can be called cloaking (<a href="http://www.google.com/webmasters/faq.html#cloaking">details in Google rules</a>) which means, that this search engine optimized code to increase the page rank. Not really a clever move considering the Wordpress (and other software developers) where teaming up with google and technorati to fight spam. 

Another problem here is, that Wordpress was the main haven for people searching a new "home" after MT went commercial. This move could make them reconsider this decision. While I can understand that a popular project like Wordpress needs sponsors it should IMO be the top priority not to alienate the core userbase. I just hope that this whole situation will be resolved rapidly by removing the offensing code and coming to terms again with google. But I still fear that some users are already quite disappointed out there. I learned about this whole thing by browsing through some of my favorite blogs so I've already read some of the reactions:

* <a href="http://hugo.muensterland.org/2005/03/31/wordpress-websites-search-engine-spam/">musterland.org</a>
* <a href="http://www.perun.net/archiv/2005/03/31/wordpressorg-und-suma-spam/">perun.net</a>
* <a href="http://www.waxy.org/archive/2005/03/30/wordpres.shtml">waxy.org</a>

... just to name a few.

There were also quite a few comments which indicated that users shouldn't confuse matt's action with wordpress... uhm.... The last time I checked Matt was the lead developer of Wordpress and the whole thing started because of the code above being present on the official wordpress website <http://wordpress.org>. While wordpress is still a quite community driven project it's still IMO valid to take project related actions by leading team members as project actions.

Another thing that amused me a little bit was the voices who raised the question if and how opensource projects should be financed... I found no comments about Free Software so this discussion probably only applies to Open Source Software ;-) (I know the joke was lame, but it's late and I'm tired ;-) ). Every work has to be financed somehow. If someone wants to get some revenue by some google ads on his/her project site or finds perhaps some company donating some hardware/software/beer/whatever where should be the problem? The only problem I see here are methods that are against some official or inoffical rules like in this case like the google ruleset in this case.

.... Ok, I'm typing far too slow or I'm simply late as always: It seems like at least some of the linked offensing pages have been removed from wordpress.org (I'm too lazy to check them all ;-) ). So it seems like a first step was taken to resolve this whole situation while Matt is still in Europe (perhaps someone in Italy told him about some of these postings all around the world ;-) ).

I just really hope, that wordpress won't take any longterm damage because of this whole thing :-)