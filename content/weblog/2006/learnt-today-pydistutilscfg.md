---
date: '2006-09-03T12:00:00-00:00'
language: en
tags:
- distutils
- learnt-today
- python
title: 'Learnt today: .pydistutils.cfg'
---


Are you a Python developer? And do you have tons of libraries installed somewhere into a non-standard prefix? Since I started coding in Python again, I wanted to avoid installing libraries globally so that I could easily manage them. So I now install everything with a --prefix=~/.python . Distutils setup.py makes this very easy since all you have to do is append this option to the `python setup.py install` call and change the PYTHONPATH environment variable accordingly. 

But distutils has something that makes this procedure even easier: If you're tired of typing --prefix=whatever again and again you can also set a default value for each user. If you're under MacOSX/Linux/Unix/blob create $HOME/.pydistutils.conf and add following lines:

-------------------------------



<pre class="config">[install]
prefix=~/.python</pre>

With this file you can basically pre-set each option the setup.py offers for each of its command. Very handy :D