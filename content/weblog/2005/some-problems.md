---
date: '2005-05-01T12:00:00-00:00'
language: en
tags:
title: Some problems
---


Sorry, but I'll have to reset the priority of hacking the wp-admin panel thanks to some data loss and some other problems with my current development server. I'm currently not completely sure what caused the data loss but I only see 2 possibilities right now:

-------------------------------



* Reiser4 FS corruption

* NFS failure



As I said, I'm not sure, but I think I will move my data partition back to Reiser3 or Ext3. The other problem is my webserver setup. To make it more flexible I will probably replace the official Gentoo ebuilds with custom builds. Since I see no real point in using Gentoo on the server anymore I'm also thinking about moving it to UbuntuLinux or Debian for easier maintenance. But first things first: First I will transform the Reiser4 partition into a Reiser3 or Ext3 partition, then I'll setup my webserver so that I can get the wp-tagging done and then perhaps the thing with Ubuntu/Debian ;-)