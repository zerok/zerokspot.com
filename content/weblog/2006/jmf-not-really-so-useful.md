---
date: '2006-11-16T12:00:00-00:00'
language: en
tags:
- java
- jmf
- mpeg-1
- sun
title: JMF not really so useful...
---


Is it just me or is JMF yet another example why Java isn't really platform independent? Or actually another example of Sun really making it useful no matter what platform you're on. For example: I now have to write a small program using JMF. I thought: Great, should be no problem using a normal MPEG-1 stream as input source. Yeah, right ... According to the [list of supported formats](http://java.sun.com/products/java-media/jmf/2.1.1/formats.html) the really platform independent JMF package doesn't support even MPEG-1. This feature is reserved for the platform-dependent binaries. I also tried to register the MP3 plugin and failed there too, so perhaps I understood something the wrong way (but at least I also read about [other people](http://weblogs.java.net/blog/jonathansimon/archive/2004/11/suns_mp3_plugin.html) having the same problem).

-------------------------------



Or is this perhaps just yet another license issue? Actually I don't really care. I can encode and decode MPEG-1 streams with FFMPEG under Linux, MacOSX and probably also Windows (haven't tried it with Windows but I guess it should work since MPlayer is also available for Windows ;-) ). 

The least Sun could do, would be to also make a platform-optimized package for MacOSX since there seem to be really some license issues for Sun to release for example the MPEG-1 codec [as source](http://www.sun.com/software/communitysource/jmf/download.xml). Please, Sun.