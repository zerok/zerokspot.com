---
date: '2007-07-18T12:00:00-00:00'
language: en
tags:
- dateutil
- python
title: Solving timezone issues with dateutil
---


During the last couple of days, one 3rd-party Python module made quite an impression on me: dateutil by Gustavo Niemeyer.

Dateutil is an extension to the datetime module that comes with the standard library in Python and offers for example some predefined Timzone classes (and factory functions) and also a way to easily generate a list of days that are between two specified dates.



-------------------------------



Just a small example: By default if you call datetime.datetime.now(), all you get is the current datetime depending on your os.environ['TZ'] setting but still without attributes like tzname and so on containing useful values. 

But if you now do something like this:

<pre class="code python">datetime.datetime.now(dateutil.tz.tzlocal())</pre>
	
The resulting datetime instance will hold the right values for your local timezone :-)

But an even cooler method is the dateutil.tz.gettz method. What does it do? It basically returns a tzinfo instance for whatever timezone name you pass to it.

<pre class="code python">datetime.datetime.now(dateutil.tz.gettz('Europe/Vienna'))</pre>
	

During my messing around with timezones on Dreamhost this tool came in verrry handy. Esp. the last snippets helped quite a lot when I wanted to convert an un-timezoned datetime object (which actually had the local time) into an UTC datetime object. Sure it is possible without gettz, but it made the code way more readable.

For more information and further examples, check out [http://labix.org/python-dateutil](http://labix.org/python-dateutil)