---
date: '2006-04-26T12:00:00-00:00'
language: en
tags:
title: iTunes top 10 ... how to deal with boredom
---


... or basically a downtime of my 1st-tier ISP (bottom-up since I don't know how many tiers my traffic goes through). Did you ever want to export to your Top10 of all times from your iTunes library to show off on your homepage (for example) or are you just bored like me?

-------------------------------



<pre class="code ruby">

#!/usr/bin/env ruby

LIBRARY_PATH = File.expand_path(&quot;~/Music/iTunes/iTunes Music Library.xml&quot;)

require &apos;rubygems&apos;

require_gem &apos;plist&apos;

class Plist

  class PDate &lt; PTag

    def to_ruby

      text

    end

  end

  class PData &lt; PTag

    def to_ruby

      text

    end

  end

end

d = Plist::parse_xml(LIBRARY_PATH)

top10=d[&apos;Tracks&apos;].values.sort{|a,b| (a[&apos;Play Count&apos;] ||= 0) &lt;=&gt; (b[&apos;Play Count&apos;] ||= 0)}.reverse[0..10]

top10.each do |t|

  puts &quot;#{t[&apos;Name&apos;]} by #{t[&apos;Artist&apos;]}&quot;

end

</pre>



Just to be on the safe side, better first backup your "iTunes Music Library.xml" or operate on a backup to begin with ;)



As always: __Use this at your own risk.__