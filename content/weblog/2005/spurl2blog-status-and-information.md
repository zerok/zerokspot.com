---
date: '2005-05-04T12:00:00-00:00'
language: en
tags:
- development
- python
title: Spurl2Blog status and information
---


As mentioned in one of my previous posts and on the Spurl forum I'm currently writing a small Python script that should fetch the latest spurls and post them to a weblog using the MetaWeblog API (since I'm using WordPress ;-) ). Yesterday I only had about half an hour of free time so I didn't really get that far. I only implemented parts of logic that determines, if a link is new or not. Today will come the actual fetching and perhaps also the posting to the weblog.

-------------------------------



I want to write it modular enough that other people can add other weblog APIs in the future. It will be possible to specify the API with a commandline parameter (just like the username and password for the weblog as well as its URL).



What will you need to run this script? You will need Python and the <a href="http://sourceforge.net/projects/pyxml/">PyXML</a> library, which is needed for parsing the RSS feed. If you want to run it on a daily basis I'd also recommend that you execute the script on a GNU/Linux server (or BSD etc.) using Cron.