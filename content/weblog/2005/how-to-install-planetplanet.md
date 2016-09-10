---
date: '2005-05-28T12:00:00-00:00'
language: en
tags:
title: How to install PlanetPlanet
---


Since <a href="http://www.planetplanet.org">PlanetPlanet</a> is the first item on my tasklist today I also thought about writting a small tutorial on how to install it :-) (Found nothing like that on the official site except in the package so perhaps this is not completely useless ;-) 

-------------------------------



First of all download the latest PlanetPlanet package <a href="http://www.planetplanet.org/planet-nightly.tar.bz2">here</a>.

It already comes with 2 example folders which hold the configuration file and the templates for HTML, RSS and Atom output (et al.). To start playing around with it I'd suggest that you simply make a copy of the "examples"-folder. In my case I renamed the copy to "gamerslog".

Now let's edit the config.ini. In the [Planet]-section you can edit the basic settings for your new Planet like the name of it, the location of the templates etc.  Since we made a copy of the "examples" folder we have to change the location of the template files to replect this change:

<pre class="code">
template_files = examples/index.html.tmpl examples/rss20.xml.tmpl examples/rss10.xml.tmpl examples/opml.xml.tmpl examples/foafroll.xml.tmpl
</pre>
becomes
<pre class="code">
template_files =gamerslog/index.html.tmpl gamerslog/rss20.xml.tmpl gamerslog/rss10.xml.tmpl gamerslog/opml.xml.tmpl gamerslog/foafroll.xml.tmpl
</pre>

Besides the [Planet] section every single feed you want to add to your Planet has its own section. For testing only one feed should be enough so let's remove all feed sections but one and change its URL to one of your own sites:

<pre class="code">
[http://www.mysite.com/rss]
name = My little testsite
</pre>

Now comes the first time we actually start the Planet script. But first you should create the "cache" folder in your Planet-directory otherwise you will have to fetch the feed ever time you start planet.py and ever test will suck something from your monthly traffic limit ;-)

Go into the root folder of your PlanetPlanet package and execute ./planet.py gamerslog/config.ini (or whatever you called your own folder :-) ) which should give you output like this:

<pre class="output">
> ./planet.py gamerslog/config.ini 
INFO:root:Subscribing <http://www.mysite.com/rss>
INFO:root:Updating feed <http://www.mysite.com/rss>
DEBUG:root:Encoding: UTF-8
DEBUG:root:E-Tag: "someHash"
DEBUG:root:Modified: May 27, 2005 09:20 PM
DEBUG:root:URI: <http://www.mysite.com/rss>
INFO:root:Processing index.html.tmpl
INFO:root:Writing output/index.html
INFO:root:Processing rss20.xml.tmpl
INFO:root:Writing output/rss20.xml
INFO:root:Processing rss10.xml.tmpl
INFO:root:Writing output/rss10.xml
INFO:root:Processing opml.xml.tmpl
INFO:root:Writing output/opml.xml
INFO:root:Processing foafroll.xml.tmpl
INFO:root:Writing output/foafroll.xml
</pre>

Now simply open the index.html in your output folder and you will see your Planet's HTML output :-) If you like you can tweak the index.html.tmpl for example by adding a stylesheet and stuff like that. During the testing I'd recommend that you use the --offline option of planet.py so that the cached feed is used. Otherwise planet.py will always grab the latest version of the feed from your website.

All you need now is to put the output folder somewhere in your public_html folder (keeping the rest of the script un-accessible for guests of your website). You could also put the application on a competely different server and upload the output folder for example via FTP or SFTP to the destination server. This is probably what I will do for planet.gamerslog :-)