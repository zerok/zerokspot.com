---
date: '2007-10-17T12:00:00-00:00'
language: en
tags:
- python
title: distutils and prefixes
---


I've tried to get around the whole *distutils* vs. *setuptools* issue ever since I started working with Python, but setuptools might have a big advantage on its list when it comes to distributing "complete" applications where I faced some problems with distutils recently. 

-------------------------------

A small example: I'm currently writing on a small utility that takes a RST document and renders it using the [Jinja](http://jinja.pocoo.org/) template engine. There I have language specific string for which I use gettext. Now I had the problem of how to bundle all these .mo files with my tool. 

With **distutils** you should normally use the `data_files` ([docs](http://docs.python.org/dist/node13.html)) parameter to specify which files should be put where if they are not Python files. This all makes sense if you want to have one single repository for files like for instance /usr/share/locale, which you can hardcode. 

<pre class="code">
    ...
    data_files=[('/usr/share/locale/',
            ['i18n/de/LC_MESSAGES/messages.mo'])
        ]
    ...
</pre>

But I wanted to have everything (code and l10n data) in one place since it makes replacing files easier and also should be way less platform dependent. 

Another problem with a central repository that should be bound to the prefix used while installing the package is, that (at least from what I can tell), there is no easy way to retrieve that prefix at runtime inside the application without guessing it using the absolute path of `__file__`. A straight forward way to solve this, would be to put everything into the same folder structure as the Python modules and install those additional files using the `package_data` ([docs](http://docs.python.org/dist/node12.html)) argument. At runtime it's now simple to reference them using the usual suspects: `join`, `dirname` and `__file__`.

There is in my opinion not really a problem with this approach if you limited it tightly bound stuff like the translations of messages in your source code. But what about things like default files such as a default set of Jinja templates with some CSS files and images? This doesn't really belong right into a Python package.

[setuptools](http://peak.telecommunity.com/DevCenter/setuptools) might be my rescue here. First of all, the EGG distribution format puts not only the Python modules in one place, but also additional data files like my .mo or .html files. And its `pkg_resources` module offers 
among other useful stuff with the `resource_filename` method a way to easily access those files without `__file__`-guessing. 

At least according to the docs. I have never used setuptools as a developer before, but since it looks like a solution to some of my biggest problems recently, I will definitely check it out intensely ;-)

P.S.: Before you say it: Yes, I know about the zip-flag and that I won't use it for this project.
