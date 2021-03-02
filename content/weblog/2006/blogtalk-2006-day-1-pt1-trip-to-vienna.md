---
date: '2006-09-30T12:00:00-00:00'
language: en
tags:
- conference
- python
- ruby
- vienna
- zope
- travel
title: 'BlogTalk 2006: Day 1 (pt.1) - Trip to Vienna'
---


Long trips definitely suck, esp. when you're making them alone. Exactly for these occasions I absolutely love the battery in my little PowerBook 12": Nearly 4 hours if I treat it nicely and so I normally also take something to read with me on the trip ... on my Mac ;)

Well, OK, in my train there were enough plugs in each area, so I could have also brightened the whole thing up and done some more compiling, but I guess esp. the latter on would have burned my Jeans. 

-------------------------------

So back to reading ... this time? The ZopeBook 2.6 as PDF. I wanted to get into Zope for quite some time now esp. since I'm always in the mood to learn yet another framework ;)

Well, after 10 minutes of installing ZODB and Zope I wanted to give it a try to follow the examples in the book ... just to find out, that Zope seems not to work with Python 2.5 yet. At least there seems to be a small bug in there with the `classImplements(OverflowWarning, IOverflowWarning)` interface binding ... I always thought there was only an OverflowError (at least to me a Warning for something like that wouldn't make all that much sense to me), but I'm still too new to Python ;-) Seems like it was there back in Python 2.4 but already with a nice deprecation warning that this Warning won't exist in 2.5. Well, since I'm not really motivated to install Zope twice, I will simply make this a theoretical session :-(

Just a small break to who you how bored I was \*g\*

```
#!/usr/bin/env python
class RubyArray:
	def __init__(self):
		self.data = []
	def __lshift__(self,data):
		self.data.append(data)
		return self
	def __str__(self):
		return str(self.data)
	def __getattr__(self,attr):
		getattr(self.data,attr)

if __name__ == "__main__":
	a = RubyArray()
	a << 1 << 2
	print str(a) # [1, 2]
```
