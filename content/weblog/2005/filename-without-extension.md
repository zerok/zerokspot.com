---
date: '2005-07-26T12:00:00-00:00'
language: en
tags:
title: Filename without extension
---


From time to time I have to write some small bash scripts (since writing them in Perl/Ruby/Python/... would simply be an overkill). Most of the time this involves some kind of directory-listing-parsing which is quite simple using find, grep and sed.

-------------------------------



Since I couldn't sleep tonight I decided to work through and there I needed a script that checks if a folder holding a .mp4 file also holds a .tgz file with the same basename. Not being really good at bash scripting googling helped me quite a lot with this nice tutorial site: <a href="http://www.splike.com/howtos/bash_faq.html">Bash Scripting FAQ on splike.com</a>



But there was only described a way to get the extension of a file. Guessing that getting the other part of the file shouldn't be all that different I opened the ABS and checked the <a href="http://www.tldp.org/LDP/abs/html/refcards.html#AEN16857">string operations reference</a>. ${string%.*} should normally do what I want. So here is the small script. Hopefullly it will be useful for someone else :)



<pre class="code">

#!/bin/bash

for video in `find . -name '*.mp4' `; do

        basename=${video%.*}

        if [ ! -f "${basename}.tgz" ]; then

                echo $basename

        fi

done

</pre>