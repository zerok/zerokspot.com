---
date: '2008-08-18T12:00:00-00:00'
language: en
tags:
- geoserver
- java
- java6
- openjdk
- ubuntu
title: Geoserver with OpenJDK on Ubuntu? I guess not
---


Or at least not for now. Today I had to install [Geoserver](http://geoserver.org) on a new [Ubuntu](http://www.ubuntu.com/) 8.04 server within a current [Tomcat](http://tomcat.apache.org) (6.0.18) for a colleague and was greeted by a nice error message related to casting and the javax.imageio package (or earlier on with a ClassDefNotFound exception related to the same package). The problem here seems to be that the stable Geoserver 1.6.x is not really compatible with [OpenJDK](http://openjdk.java.net/) yet. 

So for now the fastest way to get it working again (that is, if you don't absolutely require OpenJDK) to move back to the old Sun Java 6 package by first installing "sun-java6-jdk" and then by switching to it using `update-alternatives --config java`.

At first I thought I had to install some external libraries and miserably failed to install the jai\_imageio binary package thanks to the package being [broken](http://forums.java.net/jive/thread.jspa?messageID=282271) and naturally Sun has to have internal checksums so that you can't easily fix this. But luckily this has become a non-issue with the move back to Java6 for now, but I'm really curious if I just messed up something there or the missing jai\_imageio package was really the problem here. I guess, this is a problem for later, now that I have a workaround in place ;-)