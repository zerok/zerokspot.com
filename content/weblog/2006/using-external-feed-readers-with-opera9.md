---
date: '2006-06-22T12:00:00-00:00'
language: en
tags:
- bookmarklets
- feeds
- opera
title: Using external Feed readers with Opera9
---


For people who prefer using their own feed readers instead of the one bundled with Opera9 here a small hint that should get this going:

The idea is quite simple: Change the protocol of all the feeds linked in a HTML page from "http" to "feed". Then you need some way to show you the links to the feeds. Then simply bind your application to the "feed://" protocol in Opera using Preferences Â» Extended Â» Programs.



-------------------------------



You can combine the first and the second step by using [Martin Dittus' Feed bookmarklet](http://dekstop.de/weblog/2006/03/feed_links_bookmarklet/) which renders a list of all the available feeds on a webpage. The problem is, that you can't tell Opera9 to use a specific applications for feeds, this separation is only possible by protocol. So we need to modify Martin's Bookmarklet a little bit to output the feed links using a different protocol.

I've prepared this bookmarklet here:

<a href="javascript:function%20txt(str)%7Breturn%20document.createTextNode(str)%7Dfunction%20tag(n,c)%7Bvar%20e=document.createElement(n);e.style.fontFamily='Arial,sans-serif';e.style.color='%23000';if(c)e.appendChild(c);return%20e%7Dfunction%20p(c)%7Breturn%20tag('p',c)%7Dfunction%20a(href,desc)%7Be=tag('a',txt(desc));e.href=href.replace(/%5Ehttp:/,%22feed:%22);e.style.color='%2300c';e.style.textDecoration='underline';return%20e%7Dvar%20el=tag('div');el.style.zIndex=100000;el.style.position='absolute';el.style.padding='20px';el.style.top='10px';el.style.left='10px';el.style.backgroundColor='%23ffffcc';el.style.border='1px%20solid%20%23333333';el.style.textAlign='left';var%20ul=tag('ul');var%20found=false;var%20links=document.getElementsByTagName('link');for(var%20i=0,link;link=links[i];i++)%7Bvar%20type=link.getAttribute('type');var%20rel=link.getAttribute('rel');if(type&&(type=='application/rss+xml'%7C%7Ctype=='application/atom+xml')%20&&%20rel%20&&%20rel=='alternate')%7Bvar%20href=link.getAttribute('href');if(!href.match(/%5Ehttp/))%7Bvar%20path=(href.match(/%5E%5C//))?%20'/'%20:%20location.pathname;href='http://'+location.hostname+path+href;%7Dvar%20title=link.getAttribute('title');ul.appendChild(tag('li',a(href,((title)%20?%20title+'%20-%20'%20:%20'')+href)));found=true;%7D%7Dif(found)%7Bel.appendChild(p(txt('The%20current%20page%20links%20to%20these%20feeds:')));el.appendChild(ul);%7Delse%7Bel.appendChild(p(txt('The%20current%20page%20does%20not%20link%20to%20any%20feeds.')));%7Dvar%20close=a('%23','Close');close.onclick=function()%7Bel.style.display='none';return%20false;%7D;el.appendChild(p(close));function%20addFeedBox()%7Bdocument.body.appendChild(el);y=window.scroll(0,0);%7Dvoid(z=addFeedBox());">Show all feeds</a>

The only difference between the two bookmarklets is this fragment:

<pre class="code">
	e.href=href.replace(/^http:/,&quot;feed:&quot;);
</pre>

In the original this was:

<pre class="code">
	e.href=href;
</pre>

Now all you have to do is to specify whatever application you want to use for feeds using the "Programs" preferences as described above.

**Update 1:**

It seems like the whole protocol binding process it not completely consistent over different versions. I for example could only use the "feed://" binding while for other people only "feed" works. Please check out the comments in any case :)