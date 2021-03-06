---
date: '2007-04-12T12:00:00-00:00'
language: en
tags:
- cookies
- netscape
- pyobjc
- python
- safari
title: Exporting Safari's cookies to Netscape cookies format
---


Remember my little [download tool for IGN](http://zerokspot.com/node/734)? Well, as I said back then, it works mainly thanks to the great cookie integration you can have between Firefox and wget. But what if you're using WebKit/Safari?



-------------------------------

Well, WebKit seems to make cookies accessible through 2 paths:

1. There is a plist file holding all the cookies in ~/Library/Cookies named Cookies.plist
2. If you don't want to parse this file on your own and don't mind getting your hands on some ObjC, there is the NSHTTPCookieStorage singleton in the Foundation framework that has some nice methods like one to get all the cookies that are relevant for a given URL.

Back in the days I wrote a small ObjC extension for Ruby to go the 2nd way but just a month later or so I ported the downloader over to Python in order to clean up the script jungle I had (and mostly still have :-?) and also switched back to Firefox (don't say it!) ... so the whole code became quite obsolete.

So now I'm using Safari once again and would still like to use my downloader again. So I needed some simple wget integration with WebKit's cookie storage. 

The actually solution is quite simple using the [PyObjC](http://pyobjc.sourceforge.net/) library (wasn't really motivated to get into Python C extensions yesterday night ;-)):

@@ python @@
import objc
import tempfile
import os
from Foundation import NSHTTPCookieStorage,NSURL

FILE_HEADER = """# HTTP Cookie File
# http://www.netscape.com/newsref/std/cookie_spec.html
# This is a generated file!  Do not edit.
# To delete cookies, use the Cookie Manager.

"""

def get_cookiefile_for_url(url):
	"""
	This function returns the path to a temporary Mozilla-style formated
	cookie file that contains all the cookies provides by WebKit's cookie
	storage	for the given URL.
	
	PLEASE UNLINK AFTER USE!
	"""
	storage = NSHTTPCookieStorage.sharedHTTPCookieStorage()
	u = NSURL.URLWithString_(url)
	fd,filepath = tempfile.mkstemp()
	try:
		os.write(fd,FILE_HEADER)
		for c in storage.cookiesForURL_(u):
			line = "%s\t%s\t%s\t%s\t%s\t%s\t%s\n"%(
				c.domain(),
				'TRUE',
				c.path(),
				str(bool(c.isSecure())).upper(),
				int(c.expiresDate().timeIntervalSince1970()),
				c.name(),
				c.value(),
			)
			os.write(fd,line)
	finally:
		os.close(fd)
	return filepath
@@

The code is pretty straight forward. get\_cookiefile\_for\_url basically creates a temporary cookie file that wget understands for a given URL and returns the path to that file. As the docstring already indicates: The function only creates the file, so you have to delete it on your own.

... damn, I really need a snippets section on django@zerokspot ;-)
