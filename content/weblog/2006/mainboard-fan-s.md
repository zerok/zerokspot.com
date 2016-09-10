---
date: '2006-05-07T12:00:00-00:00'
language: en
tags:
title: Mainboard fan :S
---


<pre class="code">

-------------------------------

#!/usr/bin/env ruby

while true

  `sensors`.each_line do |l| 

    if l=~/^M\/B Temp:\W*\+(\d*) C/

      if $1.to_f &gt; 45.0

        system(&apos;halt&apos;)

      end

      break

    end

  end

  sleep(10)

end

</pre>



No, I don't trust my mainboard fan.



No, I don't like sed.



Don't do this at home, kids.