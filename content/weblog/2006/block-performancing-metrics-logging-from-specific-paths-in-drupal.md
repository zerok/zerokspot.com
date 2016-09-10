---
date: '2006-07-23T12:00:00-00:00'
language: en
tags:
- drupal
- metrics
- performancing
title: Block Performancing Metrics logging from specific paths in Drupal
---


Well, I guess you don't want Performancing Metrics logging all your jumping around in the admin panel, right? Solving this is actually quite simple. I assume that you have added Metrics' &lt;script&gt; in your theme's page.tpl.php. The cool thing about PHP template engines like phptemplate is, that you can probably most of the functions available through the API of whatever CMS you're using.



-------------------------------



In Drupal's case this also means, that the arg(int) function works. So to block Metrics from the admin panel, you could simply replace your previous script-reference with something like that:

<pre class="code">
&lt;?php if(arg(0) != &apos;admin&apos;):?&gt;
&lt;script id=&quot;stats_script&quot; type=&quot;text/javascript&quot; src=&quot;http://metrics.performancing.com/drupal.js&quot;&gt;&lt;/script&gt;
&lt;?php endif; ?&gt;
</pre>

Nothing really fascinating about this, but it's simply nice to have all this stuff also available in the templates and this is just IMO a quite good use for it since Metrics has nothing to do in my admin section.

With arg(int) being available  you could naturally do the same thing with other paths like for example blocking it from specific nodes with code like:

<pre class="code">
&lt;?php if(!(arg(0) == &apos;node&apos; &amp;&amp; arg(1)==&apos;2&apos;)): ?&gt;
...
&lt;?php endif; ?&gt;
</pre>