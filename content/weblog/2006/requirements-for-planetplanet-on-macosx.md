---
date: '2006-02-21T12:00:00-00:00'
language: en
tags:
title: Requirements for PlanetPlanet on MacOSX
---


Well, I once again wanted to play around with PlanetPlanet to improve some things over there on Gamerslog. So I downloaded the latest nightly and thought I could just execute the planet.py ... bad idea ;)



-------------------------------



<pre class="output">intrepid:~/planet-nightly zerok$ python ./planet.py --help
Traceback (most recent call last):
  File "./planet.py", line 22, in ?
    import planet
  File "/Users/zerok/planet-nightly/planet/__init__.py", line 33, in ?
    import dbhash
  File "/System/Library/Frameworks/Python.framework/Versions/2.3/lib/python2.3/dbhash.py", line 5, in ?
    import bsddb
  File "/System/Library/Frameworks/Python.framework/Versions/2.3/lib/python2.3/bsddb/__init__.py", line 40, in ?
    import _bsddb
ImportError: No module named _bsddb</pre>

This is caused, by MacOSX lacking the BerkeleyDB module and a broken bsddb module. So far I've found a quite easy solution: Installing Python-2.4 and py-bsddb from DarwinPorts. First I thought it might work if I only go with db41, which should normally be compatible with the bsddb module bundled with Tiger, but I had forgotten something: The module lacked the linking module to the BerkeleyDB altogether, so I basically had 2 options:

1. Replace the bsddb module altogether
2. or go with a 3rd party installation

Since the 2nd option is definitely the easier one, I went with it (after debugging browsing through Python and C source code for about 2 hours :-? ).

<pre class="command">sudo port install python24 py-bsddb</pre>

Short, isn't it ;)

Sorry, but I haven't yet found a way to get the dbhash module to work with the Python version that is bundled with MacOSX thanks to the bsddb module in the same bundle. If someone has found a solution for this: Please let me know :)

In the meantime: `python2.4 planet.py /path/to/config` :D

Enjoy