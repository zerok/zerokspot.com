---
date: '2006-06-10T12:00:00-00:00'
language: en
tags:
- development
- ruby
- web
title: Sometimes expire fields in cookies just suck
---


If you can read this, I just made my first blog post using my little commandline tool for blogging ;) But I can tell you, writing this was just a pain. Ok, not at first where I could rely on normal XMLRPC APIs like the metaWeblog API. But then I wanted to create new categories on the fly and this is exactly what none of the available APIs for WordPress seems to be able to do. The quite obvious solution for this is simply do a POST call on the wp-login.php to receive the cookies for an admin and then continue with creating categories this way. 



-------------------------------



The problem just was, that (1) I couldn't remember the specs for cookies exactly anymore and that someone had the great idea to put a nice and verbose date format (feature complete with "," after the day name) into the expire field :-?. After solving this I naturally forgot that cookies in the sending request are separated by semicolons ;)

<pre class="code">cookies = get_cookie.split(&quot;,&quot;).map{|c| c.split(&quot;;&quot;)[0].strip}.reject{|e| e=~/^[0-9]/}.join(&quot;;&quot;)</pre>

Ugly, but it works ;)

Another problem is, that none of the APIs supports the specification of custom fields, so using tags is quite a problem. And the solution?

Well, I don't really have a good, but at least a working one: For a week or so now I'm using categories and tags at the same time to prepare the blog for dropping tags altogether. No, not externally ;) Just internally. I will simply use categories like tags from now on which is also why I wanted to be able to specify categories in my commandline tool on the fly (just like the admin panel in WP2.0 allows it).

The other problem I had while testing this little tool, was some stupid whitespace character that turned up at the end of my input. Thanks to some Performancing Publish for Firefox I finally saw this when copy'n'pasting it ;)

About the tool itself: If you're curious, let me know :) Currently it is quite optimized for what _I_ need: WordPress + Markdown. But it's at least planed to add support for the new TextPattern XMLRPC interface.

And if you've subscribed to my feed: Sorry for the accident a few minutes ago ;)