---
date: '2008-12-10T12:00:00-00:00'
language: en
tags:
- export
- pownce
title: Pownce export and files
---


<img src="http://zerokspot.com/uploads/pownceexport-20081210-004809.png" class="left" alt="" />It took a little bit over a week not tonight I finally got my export of my [Pownce](http://pownce.com/) data, an API-dump of all my posts (or if I want all my posts combined with all the posts I've seen from all the other people I've followed). But the dump does *not* include the actual files uploaded to the service. They are still on their S3 account (or on the server they used prior to using S3) and only the links to those files are included in the export. 

-------------------------------

If you move over to Vox, their importer also seems to at least do some importing of the files but only of thumbnails. The original-size images, for example, only seem to be linked to Pownce's servers -- as can be seen for example [here](http://zerok.vox.com/library/post/6a00c2251f87a6604a0109d0724395000e.html) -- which is kind of ... useless. Esp. if it's obvious that the file is stored on one of Pownce's old servers the importer should try to really copy them over.

Don't get me wrong, though: I totally see why they don't offer one big n GB-download for all your posts including files, but at least the importer to Vox should be smart enough to do that for you. Or for people out there who actually don't want to move to Vox a small tool for extracting all the files would be pretty nice. It's also kind of weird to see that some of the [competition](http://www.soup.io/pownce) actually promotes that they will also also import all your files ... 

For myself I created a small Python-script that simply downloads all the
files mentioned in the XML dump and stores them according to their internal
name. Maybe it's also useful to others, so here's [the link](http://gist.github.com/34160). Use at your own risk.