---
date: '2006-03-25T12:00:00-00:00'
language: en
tags:
title: Apple's Ruby and Python version getting more annoying ...
---


Somehow it looks to me, like Apple is fixing everything about MacOSX except those things that annoy me the most \*g\* Being a developer guess what's bothering me ;) Broken programs bundled with MacOSX. Just to name a few problems:

* Broken Ruby libraries and components:
	* Now readline support for irb
	* mkmf appears to be completely broken (dying right away with "Can't find header files")
	* Other build-system related stuff mentioned for example [here](http://wiki.rubyonrails.com/rails/pages/HowtoInstallOnOSXTiger#4)
* Broken Python libraries as described [here](http://weblog.zerokspot.com/posts/587/)



-------------------------------



As long as "updating" Ruby and Python to actually working setups is a easy as running `port install ruby python24` I don't really care all that much, but the situation gets a little bit more complicated if you don't just problem for you and yourself as customer but for other people out there. Is really everyone running a self-compiled or fink/DarwinPorts-based Ruby/Python? I somehow doubt that. But thanks to Apple's obvious lack of motivation to fix those quite well known issues I simply don't see a realy option here. If there is a good reason for not fixing these things I'd really love to here it :)

Sorry, if this sounds a little bit edgy but that's exactly how I'm currently feeling. Failing to install a Ruby application and then in the end noticing that a part of the core library is causing the problem here is quite annoying.

Now I just hope, that switching to the DarwinPorts version of Ruby will actually solve all of my problems. Ok, I will have to re-install all my gems and remove the old ones from the /usr prefix. It somehow still doesn't get less annoying ;) And can you really expect your users to use not the version of a base component that is actually bundled with their operating system but manually update? And yes, IMO having to install a 3rd-party software repository like fink or DarwinPorts just to be able to the update qualifies as "manual" for me ;)