---
date: '2006-08-14T12:00:00-00:00'
language: en
tags:
- bash
- scripting
- shell
- unix
title: Everyday something new to learn with Unix
---


What I love about Unix? Well, that you learn something new every day ;-) (just to name one thing I love about it ...). So I'm currently in the process of learning some new frameworks and also getting more familiar with Perl (for ... I think it's now the 4th time). Today I started messing a little bit around with Django and also more or less by accident started watching some of the [webcasts posted on the Plone website](http://plone.org/about/movies). Very informative and entertaining stuff esp. if you've already at least seen some source code produced for the mentioned technologies there (and I don't mean Zope/Plone) ;-)

-------------------------------



Anyway, Sean Kelly was doing some commandline magic in his [webcast](http://oodt.jpl.nasa.gov/better-web-app.mov) comparing J2EE with things like Zope and RoR and used something I had never seen before: 

You want to rename a folder for example correcting a plural form or something similiar: 

<pre class="command">mv picture{,s}</pre>

This will rename the folder/file "picture" into "pictures". Very handy :D Actually it will simply take the word "picture" and split it up into 2 strings: One without and one with "s" appended. Just try it with something like this:

<pre class="command">echo {hello,world}</pre>

Absolutely cool, also for creating multiple folders within a shared parent directory ;-)

<pre class="command">mkdir /Users/zerok/{hello,world}</pre>

After some short checking the Bash reference manual I now know, that this feature is called ["brace expansion"](http://www.gnu.org/software/bash/manual/bashref.html#SEC27).