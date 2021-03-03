---
date: '2006-06-21T12:00:00-00:00'
language: en
tags:
- macos
- rlang
title: lcc_dynanmic problem when compiling R libraries
---


Today I wanted to install the [evd](http://cran.r-project.org/src/contrib/Descriptions/evd.html) library for [R](http://www.r-project.org/) but somehow failed with such an error message:

-------------------------------



<pre class="error">

ld: can&apos;t locate file for: -lcc_dynamic

make: *** [evd.so] Error 1

ERROR: compilation failed for package &apos;evd&apos;

</pre>



It seems like -lcc_dynamic isn't really needed, though. So to disable this you have to change the /Library/Frameworks/R.framework/Resources/etc/Makeconf. Search for following lines:



<pre class="code">

LIBS =  -lcc_dynamic -lm -liconv

[...]

SHLIB_LIBADD = -lcc_dynamic

</pre>



And remove "-lcc_dynamic" from both lines. At least evd seems to work now :)
