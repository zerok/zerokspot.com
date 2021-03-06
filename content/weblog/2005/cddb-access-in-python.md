---
date: '2005-01-08T12:00:00-00:00'
language: en
tags:
- python
title: CDDB access in Python
---


Today I played around with some modules and libraries and Python, so I want to use this as a starting point to fill the "Development" category here a little bit :-)

In the first ... let's call it a "Mini-Intro" I want to give a small introduction into Py-CDDB and how you can use it to get to know the CDDA in your CD-ROM driver better.

-------------------------------

## Requirements

* <a href="http://www.python.org">Python</a>
* <a href="http://cddb-py.sourceforge.net/">CDDB-Py</a>
* 1 or more CD-ROM drives
* 1 CDDA

## What and how?

In this small intro you will get to know 2 modules provided by the CDDB-Py package:
1. DiscID
2. CDDB

You can use these two modules to access data about your audio CD stored for example in the database on <http://www.freedb.org> which is also the default database queried by the CDDB module. 

Another cool thing here is, that the CDDB module uses the urllib and so also uses the $http_proxy environment varialble if it is set :-)

The homepage of the CDDB-Py project already gives a nice little example on how to use these two modules so I will just re-use it and go through it step by step.

As the first step we need to import the two modules:

<pre class="code">import CDDB, DiscID</pre>

Now that we have the modules at our disposal let's open the CD-ROM device and get the ID of the disc:

<pre class="code">cdrom = DiscID.open("/dev/hdc")</pre>

In the original example Disc.open() was used without the optional device-parameter. I use it here so that I get a working example on my own machine :-)

Ok, we now have the CD and can get its DiscID:

<pre class="code">disc_id = DiscID.disc_id(cdrom)</pre>

We can now use the retrieved DiscID to query the FreeDB server:

<pre class="code">(query_status, query_info) = CDDB.query(disc_id)</pre>

CDDB.query() has some parameters for choosing for example which server to query and with what account. Only the first parameter is required as it holds the ID of the disc. If everything goes well, the query_info element will now hold a list of matching CDs. The original example doesn't seem to use a list here, so I will adapt the example for the last few lines a little bit.

<pre class="code">(read_status,read_info) = CDDB.read(query_info[0]['category'],query_info[0]['disc_id'])
</pre>

With this step we finally get some "plaintext" info about the CD like the artist, the tracknames etc.

We can now use this to generate a small list of all the tracks of the disc including their titles and close the cdrom at the end.

<pre class="code">for i in range(disc_id[1]):
	print "Track %.02d: %s" % (i, read_info['TTITLE' + `i`])
cdrom.close()
</pre>

This will procude this little listing for the CD I have in my CD-ROM drive right now:

<pre>Track 00: Microphone Fiend
Track 01: Pistol Grip Pump
Track 02: Kick Out The Jams
Track 03: Renegades of Funk
Track 04: Beautiful World
Track 05: I'm Housin'
Track 06: In My Eyes
Track 07: How I Could Just Kill A Man
Track 08: The Ghost of Tom Joad
Track 09: Down on the Street
Track 10: Street Fighting Man
Track 11: Maggie's Farm
</pre>

Again, this small intro heavily uses a code sample available on <http://cddb-py.sourceforge.net/>. For more details on CDDB-Py please read the documentation available in the source package and on the homepage :-)