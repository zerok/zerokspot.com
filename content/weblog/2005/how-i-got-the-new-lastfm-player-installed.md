---
date: '2005-08-11T12:00:00-00:00'
language: en
tags:
title: How I got the new Last.fm player installed
---


As you can already see in the title this post will describe how exactly I got Last.fm player installed on my laptop. This is not a generic howto but IMO it should also work on most other Linux systems out there.

-------------------------------



Since the player requires Qt4 and this is currently AFAIK not shipped with any of the major distributions and only hardmasked in Gentoo I had to compile it from source. So I got the bz2 package from trolltech.com or precisely from <a href="http://www.trolltech.com/download/qt/x11.html">here</a>.



After extracting the package came the compiling part:

<pre class="command">

./configure --prefix=/opt/qt4 -qt-gif

make

make install # as root

</pre>



This will compile and install Qt4 into /opt/qt4 (so that it won't mess with my Qt 3.x installation).



Now that Qt4 was installed, I downloaded the player's sourcecode from <a href="http://www.last.fm/help/player/">here</a>, extracted it and started compiling it:



<pre class="command">

export QTDIR=/opt/qt4

export QMAKESPEC=

/opt/qt4/bin/qmake

make

</pre>



After a minute I had the "player" executable in the same directory :) That's it.