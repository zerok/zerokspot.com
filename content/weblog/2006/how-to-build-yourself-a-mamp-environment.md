---
date: '2006-01-05T12:00:00-00:00'
language: en
tags:
- development
- macosx
- web
title: How to build yourself a MAMP environment
---


I normally have a server somewhere standing here which I use as my development environment for all webbased stuff (with a focus on PHP/MySQL) so I had not yet the pleasure to install a \*AMP system on my new Powerbook. But a few days ago I was at the University and simply wanted to do some minor testing, but couldn't becasue my server was at home and behind some other machines and firewalls. So I decided to give it a try and install a MAMP system (MacOSX Apache MySQL PHP ;) ) on my laptop which should offer only the real basics. I don't need SSL here because the whole thing should only be available from localhost. OK, there is for example [XAMPP for MacOSX](http://www.apachefriends.org/de/xampp-macosx.html) which is ... kind of outdated. And since I don't really fear the shell (actually I decided to get a Mac instead of giving Windows a try again after 5 years of Linux *because* I could have a full UNIX system in the background with all its shell magic ;) ) I decided to go with doing it using the source packages and not something someone else has compiled ;) In the following short article I will describe exactly what I did.



-------------------------------



**First of all:** This tutorial will use quite a few libraries from the fink repository, so you'll need to have [fink][] installed to be able to proceed.

**A small warning here:** This is just my personal playground environment which is quite raw and without all the nice whistles and candy you can expect from a complete package (like it is provided on a decent webserver). So please don't use this to setup your public webserver ;) It just includes what I currently need with probably more to be added later when at least this basic setup is working.

**A bigger warning here:** I take no responsibility when following this tutorial breaks your system or in any other way costs you time, money or anything else. This simply describes how I did it and doesn't necessarily mean it will work for anyone else. You follow this tutorial at your own risk.

## Requirements

* Apache (2.2.0)
* PHP (5.1.1)
* MySQL (5.0.16)
* libpng
* libjpeg
* libtiff
* zlib
* FreeType

... not to forget the developer tools which you can find on your MacOSX Tiger discs.

## Installing MySQL

This is definitely the easiest part of this whole endevour:

1. Get the [standard binary package for MacOSX 10.4](http://dev.mysql.com/downloads/mysql/5.0.html)
2. Install it like you would any other app for Tiger. This package also includes an extension for the SystemPreferences which makes starting and shutting down the server much more comfortable ;)
3. Well, nothing really left for the server part but it's always nice to have a client so you could for example go with [CocoaMySQL](http://cocoamysql.sourceforge.net/) or the [MySQL Administrator](http://dev.mysql.com/downloads/administrator/1.1.html) or install [phpMyAdmin](http://www.phpmyadmin.net) after installing PHP.

## Installing Apache

Since Apache doesn't provide binary packages for Unix/BSD we will have to compile it ourself. Sure, Apache (1.3.3) is bundled with Tiger, but we want to stay on the bleeding edge of releases, don't we ;) So let's get started.

From now on everything that we compile and install ourselves will we put in a separated directory so that it won't mess with other already installed software. For this tutorial I've used /opt/wwwdev as the directory which will from now on house Apache, PHP and also the main htdocs directory for the httpd.

So let's get the source code and start compiling ;)

<pre class="command">
sudo mkdir /opt/wwwdev
wget http://apache.4any.org/httpd/httpd-2.2.0.tar.bz2
tar -xjf httpd-2.2.0.tar.bz2
cd httpd-2.2.0
./configure --prefix=/opt/wwwdev/ --enable-mods-shared=all \
&amp;&amp; make \
&amp;&amp; sudo make install
cd -
</pre>

This will install Apache HTTPD 2.2.0 with all modules (where ./configure determined it was possible to build them) built in a way so that you can easily enable and disable each of them by just editing the httpd.conf file.

To start your httpd simply run `sudo /opt/wwwdev/bin/apachectl start`.

## Installing PHP

PHP was by far the most annoying part of getting this basic MAMP installation working mostly because I wanted GD support and also at least limited FreeType support. For this little joy a few extra libraries are needed which mostly can be installed using [fink][] .

<pre class="command">
sudo apt-get install libpng3 libpng3-shlibs libpng
sudo apt-get install libjpeg libjpeg-bin libjpeg-shlibs
sudo apt-get install libtiff-shlibs
</pre>

I somehow couldn't find the zlib in the fink repository so this will also has to be compiled manually:

<pre class="command">
wget http://www.zlib.net/zlib-1.2.3.tar.bz2
tar -xjf zlib-1.2.3.tar.bz2
cd zlib-1.2.3.tar.bz2
./configure --prefix=/opt/wwwdev &amp;&amp; make &amp;&amp; sudo make install
</pre>

The last requirement for our little dev setup is FreeType2. This *is* available in fink but I somehow couldn't make PHP recognizing it. So let's get back to the manual way here as well:

<pre class="command">
wget http://download.savannah.nongnu.org/releases/freetype/freetype-2.1.10.tar.bz2
tar -xjf freetype-2.1.10.tar.bz2
cd freetype-2.1.10
./configure --prefix=/opt/wwwdev &amp;&amp; make &amp;&amp; make install
cd -
</pre>

Now that all the dependencies should be installed, let's get to PHP itself.

<pre class="command">
wget http://at2.php.net/get/php-5.1.1.tar.bz2/from/this/mirror
tar -xjf php-5.1.1.tar.bz2
cd php-5.1.1
./configure  --prefix=/opt/wwwdev/ --with-mysql=/usr/local/mysql \
	--with-apxs2=/opt/wwwdev/bin/apxs --with-gd --with-jpeg-dir=/sw/ \
	--with-png-dir=/sw/ --with-freetype-dir=/opt/wwwdev/ \
	--with-tiff-dir=/sw/ --enable-cli --enable-mbstring --enable-sockets \
	--with-xpm-dir=/usr/X11R6/ --with-zlib-dir=/opt/wwwdev/ \
&amp;&amp; make &amp;&amp; sudo make install
cd -
</pre>

Thanks to apxs `make install` will already enable mod_php in your httpd's httpd.conf but it won't bind the .php and .phps file extensions to the PHP module so we will have to do this manually. Therefor open /opt/wwwdev/conf/httpd.conf and add following 2 lines at the end of the config file:

<pre class="config">
AddType application/x-httpd-php .php 
AddType application/x-httpd-php-source .phps
</pre>

Now, to check if PHP is really working create a info.php file in /opt/wwwdev/htdocs with your usual `phpinfo()` call in it:

<pre class="code">
&lt;?php
phpinfo();
?&gt;
</pre>

And open it in your webbrowser:

<pre class="command">
open http://localhost/info.php
</pre>

You should now see a nice listing of your PHP installation's setup and **not** the code you put into the info.php file ;)
[fink]: http://fink.sf.net "Fink Project"