---
date: '2006-04-11T12:00:00-00:00'
language: en
tags:
title: Ugly stuff
---


<pre class="code">

-------------------------------

#!/bin/bash

for i in `find . -path './[0-9]*/[0-9]*' -type d -prune -exec dirname {} \;` ; do

	FOLDER=`basename $i`

	cd $i \

	&& find . -path './[A-Za-z]*' -type d -prune -exec rm -r {} \; \

	&& cd - \

	&& mv -f $i/${FOLDER}/* $i/

done

</pre>



Something are slow, ugly ... but work ;) Last week I messed something up while automatically copying stuff around and well, the solution is really just ugly ;)