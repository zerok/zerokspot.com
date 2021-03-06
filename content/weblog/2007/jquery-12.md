---
date: '2007-09-11T12:00:00-00:00'
language: en
tags:
- javascript
- jquery
title: JQuery 1.2
---


<img src="/media/2007/jquery.png" alt="" class="left" />Today seems to be a good day for releasing new software. At least that's what Sony and John Resig might have thought when they woke up today :-) And this is a quite special update for the popular [JQuery](http://jquery.com) JavaScript library.

-------------------------------

Why special? Because it doesn't only add new features, but also washes some stuff you might have become accustomed to from the 1.1 branch out of the core. One of these features is the XPath-like syntax for node-selection as could be use with something like

<pre class="code">$("//div/p")</pre>

Now you have to use the [XPath plugin](http://docs.jquery.com/Release:jQuery_1.2#XPath_Compatibility_Plugin).

If you are not motivated for changing all your code for this update, there is also a [plugin](http://docs.jquery.com/Release:jQuery_1.2#jQuery_1.1_Compatibility_Plugin) for getting all the features that have been removed from the code back.

On the plus side of things 1.2 comes with some new features including 2 that saw a pre-release in 1.1.4 namely the :has selector and the .slice method. And there is among others also the .map method which should be quite self-explanatory and the .content method which returns all the child elements of a specific node (including text nodes according to the [docs](http://docs.jquery.com/Release:jQuery_1.2/Traversing#.contents.28.29) ). At least these are the new features I will probably use the most.

For a complete description of all the new features and a list of removed features, check out the [release notes](http://docs.jquery.com/Release:jQuery_1.2) and the [announcement](http://jquery.com/blog/2007/09/10/jquery-12-jqueryextendawesome/) on the JQuery blog.
