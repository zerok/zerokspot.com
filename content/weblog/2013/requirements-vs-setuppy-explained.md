---
date: '2013-07-28T19:14:46+02:00'
language: en
link: https://caremad.io/blog/setup-vs-requirements/
tags:
- python
title: Requirements.txt vs setup.py explained
---


Ever wondered about when you should use a requirements.txt file or the "install_requires" section in your setup.py to specify your project's dependencies? Then [this article](https://caremad.io/blog/setup-vs-requirement/) by Donald Stufft is exactly what you need.

> This split between abstract and concrete is an important one. 

... and this split is manifested in the Python world by the existence of requirements.txt and setup.py. Also: it is just good to not feel alone anymore when thinking that Go-lang's dependency-management is rather tedious and to some degree broken since their default tool for resolving external libraries doesn't really distinguish between pinned dependencies and those with a valid version range (yet).
