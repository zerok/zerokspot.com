---
date: '2004-12-07T12:00:00-00:00'
language: en
tags:
- development
title: XSL parameters with PHP5
---


Occording to the <a href="http://www.w3.org/TR/xslt#variable-values">XSLT 1.0 specification</a> parameters and variables in XSL(T) are basically more like constants than like variables. Ok, bad naming then but let's get to the point of this post: I'm currently writing a small template engine based on XSL for a project. Here I use parameters for passing i.e. small language specific texts to the stylesheets. I use a small hash for caching these parameters to that I still have my re-declaration. 

-------------------------------



Then I thought, perhaps PHP5 is caching these parameters before sending it to the stylesheet by itself through the XsltProcessor::setParameter(...) method. And yes, it seems to do this :-) First I thought it doesn't thanks to some caching of my template function :-)