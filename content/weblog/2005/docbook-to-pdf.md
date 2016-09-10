---
date: '2005-04-19T12:00:00-00:00'
language: en
tags:
title: Docbook to PDF?
---


Docbook is IMO simply the best documentation tool for software projects. Normally you just want to export it in some online-readable format like HTML. This is quite simple using `xsltproc`. But there are always times when you want to create a PDF out of your documentation. First I thought there should be an easy way using openjade, but after some problems I found <a href="http://lists.ethernal.org/cantlug-0211/msg01084.html">this mail</a> by Michael JasonSmith describing some alternatives.

-------------------------------



 The procedure is quite simple and I will try to list all the packages required for it under Gentoo:

* passivetex
* tetex (>= 3.0-r2)

Why tetex >= 3.0-r2? Previously I think I had 3.0-r1 installed and somehow ended up without pdftex which I now have after updating, so perhaps it was really missing back then :-?

I won't list the requirements for the "docbook to xhtml" process here, simply because I think if you work with docbook you should know what you need for it ;-)

In the first step we have to generate a .fo file using the print-stylesheet for docbook and xsltproc:

<pre class="command">
xsltproc /usr/share/sgml/docbook/xsl-stylesheets-1.66.1/fo/docbook.xsl index.xml > index.fo
</pre>

Now that we have the index.fo, we can transform it into a PDF:

<pre class="command">
pdfxmltex index.fo
</pre>

Now you should have a PDF version of your docbook. The quality isn't really great, but it's IMO good enough for what I need :-)

<h3>Troubleshooting</h3>
After the generation of the fo-file, I tried pdfxmltex to transform it into a PDF but code following error:

<pre class="error">! TeX capacity exceeded, sorry [save size=5000].</pre>

Here I thought about just changing this value in the /etc/texmf/texmf.cnf which is quite futile under Gentoo since <a href="http://forums.gentoo.org/viewtopic-t-192860-highlight-texmf+cnf.html">this file isn't used</a>. If you want to change any settings better try it in <em>/usr/share/texmf/web2c/texmf.cnf</em>. For my documentation I tried it with save_size=10000 which worked.