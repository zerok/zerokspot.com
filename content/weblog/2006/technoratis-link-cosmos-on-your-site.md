---
date: '2006-07-31T12:00:00-00:00'
language: en
tags:
- blogging
title: Technorati's link cosmos on your site
---


Natalie Jost has a [nice post](http://nataliejost.com/blog/quick-tip-tracking-who-links-to-you) about how to track who else is linking to your blog but also a question: "How to get this information out of Technorati?" The answer is a little bit more techie than it should have to be, but anyway:

Technorati is offering an API for getting exactly this kind of information out of it. What do you need for that?

1. An account on Technorati
2. An API key which you can get [here](http://technorati.com/developers/apikey.html)

The last thing you need is a script that actually uses this API. How should it work? That absolutely depends on what you want. I personally prefer a small script that polls that data let's say once per day and stores the output inside of a PHP file that I could then simply integrate in whatever CMS I'm currently using.



-------------------------------




A very simple script for doing something like that would be:

<pre class="code">
#!/usr/bin/env ruby
require &apos;rexml/document&apos;
require &apos;open-uri&apos;
require &apos;cgi&apos;

OUTPUT_FILE=&quot;cosmos.php&quot;
API_KEY=&quot;&quot; # your Technorati API
BLOG_URL=&quot;&quot; # your site. e.g.: zerokspot.com


class String
  def escape_single_quotes
    self.gsub(/[&apos;]/, &apos;\\\\\&apos;&apos;)
  end
end
# http://www.bigbold.com/snippets/posts/show/880

u=&quot;http://api.technorati.com/cosmos?key=%s&amp;url=%s&quot;%([API_KEY,BLOG_URL])
open(u) do |site|
  doc = REXML::Document.new(site.read)
  open(OUTPUT_FILE,&apos;w+&apos;) do |output_file|
    output_file.write(&quot;$links=array();\n&quot;)
    doc.elements.each(&quot;//item&quot;) do |item|
      puts &quot;#&quot; if $DEBUG
      out = &quot;$links[] = array(\&quot;site\&quot;=&gt;\&apos;%s\&apos;, \&quot;url\&quot;=&gt; \&apos;%s\&apos;);\n&quot;%([
        CGI::escapeHTML(item.elements[&apos;weblog/name&apos;].text.escape_single_quotes),
        CGI::escapeHTML(item.elements[&apos;nearestpermalink&apos;].text.escape_single_quotes)
        ])
      output_file.write(out) 
    end
  end
end
</pre>

All you'd have to do is change the API\_KEY and BLOG\_URL constants and it would create a cosmos.php whereever you've started this script. This php file would then contain entries in following format:

<pre class="code">
$links = array();
$links[] = array(&quot;site&quot;=&gt;&quot;My site&quot;, &quot;url&quot;=&gt;&quot;http://mysite.com/pointing_to_you.html&quot;);
</pre>

Ready for being integrated in any php script and for cron'ing :)

The code isn't all that fantastic but it should at least be a good starting point :)
