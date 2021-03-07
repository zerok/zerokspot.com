---
date: '2008-12-28T12:00:00-00:00'
language: en
tags:
- semanticweb
title: Making OWL files readable
---


While working on some projects I started also looking into transforming ontology files (and OWL/XML in particular) into something more pleasing to the eye -- such as HTML. 

I guess, the obvious first choice here is XSL(T) simply because it's what you want to use if you need to transform one XML file into another (or something different). So I looked around and found among tons of other people doing it some [really great work](http://www.kanzaki.com/ns/ns-schema.xsl)) by Masahide Kanzaki (whose [Exif ontology][] I absolutely love) as well as some other XSL-stylesheets. Since I hadn't worked with XSLT in ages I also started tinkering around with a stylesheet on my own. The result of that is nothing really great, so far, but if you absolutely have to take a look you can find it within my [owltools][] repository on github. I've only tried it with the OWL-files that Protégé 4.0.x produces and also with Masahide Kanzaki's Exif ontology, so I can at least be fairly sure, that it works ;-)

[owltools]: http://github.com/zerok/owltools/
[exif ontology]: http://www.kanzaki.com/ns/exif

There still remains the problem, though, that XSLT might not be the ideal tool to render OWL or any kind of RDF application for that matter given the syntax-independence there. For now this XSLT does what I want it to do but eventually the repository might also see some Jena/Pellet-based tool that will treat OWL files more like RDF graphs. I guess that's again something for my life after the final exam ;-)
