---
date: '2010-10-24T12:00:00-00:00'
language: en
tags:
- couchdb
- backup
- replication
- python
title: CouchDB backup through replication
---


After I put my first site using CouchDB online I also looked into ways to
create off-site backups of it, but for some reason I couldn't find any existing
scripts for something like that. Sure, I could't always just do some rsync'ing
of the actual files as suggested on
[stackoverflow](http://stackoverflow.com/questions/121599/couchdb-backups-and-cloneing-the-database)
but where is the fun in that? :-)

Yesterday afternoon I had a few minutes without anything else to do so I wrote
a small script that provides a simple historical backup using replication and
put it on [Github](http://gist.github.com/642146). Enjoy :-)

<script src="http://gist.github.com/642146.js"> </script>
