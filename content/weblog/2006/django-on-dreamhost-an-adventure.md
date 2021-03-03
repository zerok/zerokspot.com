---
date: '2006-09-01T12:00:00-00:00'
language: en
tags:
- django
- python
title: 'Django on Dreamhost: An adventure'
---


Somehow Dreamhost's server setup must be really strange ... For the last 4 hours I tried setting up a dummy Django project on Dreamhost using basically a combination of the tutorial in the [Dreamhost WIKI](http://wiki.dreamhost.com/index.php/Django) and the tutorial on [Djangoproject.com](http://www.djangoproject.com/documentation/fastcgi/) for using Django with FastCGI ... but first things first:

At home I'm using a prefix setup for all Python libraries I'm using including Django which seems to work quite well ... at home. I now tried to mirror the setup on Dreamhost but received only 500 errors whenever the script should have been called via FastCGI. The reason for this seems to be very simple: For some reason, easy-install'd package appear not to work the way they should when called from FastCGI. Well, to solve this you have the common solutions at hand:

-------------------------------



* a nice little django.pth file in your local site-packages directory + `sys.path.insert(0, '/home/username/.python/lib/python2.4/site-packages')` in your django.fcgi
* Adding the django module's parent directory directly to sys.path in the fcgi script

None of these are particularly nice solutions, and for some reason, the first one didn't work for me. But at least the 2nd option works and so Django now finally also seems to work on Dreamhost for me :) Big thanks to the authors of the respective articles.

Another funny aspect of Django on Dreamhost is, that the first (or the first two) request(s) after changing the code for the FastCGI wrapper seem(s) to go down the next /dev/null and returns only the usual 500 error message or a nice message from project.urls about an invalid URL although the URL should be correct ;-)
