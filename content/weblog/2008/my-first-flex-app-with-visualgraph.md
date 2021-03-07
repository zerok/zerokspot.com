---
date: '2008-01-20T12:00:00-00:00'
language: en
tags:
- actionscript
title: My first Flex-App with VisualGraph
---


Yesterday, while playing around for the first time with Adobe Flex, I came across a really cool component for visualising directional and non-directional graphs. The [Flex Visual Graph Library](http://code.google.com/p/flexvizgraphlib/) does more or less everything for you. All *you* have to do, is add nodes and edges to it and select one from a bunch of layouting algorithms for positioning these nodes.

-------------------------------

A problem here is, though, that the last "release" of the component happened in August 2007 and much has changed since then (which is "a good thing"). Sure, the core concepts are still the same but for instance back then some of the layouters weren't there, and for instance `IVisualGraph` got a whole lot of new properties that make the library way more flexible and all of them show up in the official documentation since this documentation is simply part of the repository. But when you download the pre-compiled package, you will often hit walls on your way to generate your first graph on your own.

Another problem is the example code available in the [repository] (http://flexvizgraphlib.googlecode.com/svn/trunk/RelationNavigatorDemo/). Sure, it probably shows you every feature under the sun, but when you just want to get a basic graph done and don't know the library yet, it's quite overwhelming - esp. if you're also new to Flex and ActionScript ;-)

So, in order to get to know Flex and this component I wrote a tiny app that does really not all that much: It renders 3 nodes on a directional graph. Each edge also has a label. The official demo apps take all the data from either an XML datastream or some XML directly stored in the apps mxml file. Since in the end I want to be able to create nodes programmatically I also did it this way here.

<div class="figure">
<img src="/media/2008/visualgraphdemo.png" alt="Visual Graph demo"/>
<p class="caption">Not much to see. But sometimes it's better this way.</p>
</div>

The demo also features a quite barebone zoom function which basically only increases the scale everything is rendered in. This works in most cases unless you start focusing nodes or ... well, let's say that I simply haven't found a better way to zoom the graph yet.

<strong>[Download](/media/2008/Graph.mxml)</strong>

For details on how to compile it please checkout the offical guides about Flex and how to compile 3rd-party components. Again: This example is far from being perfect and definitely reflects that I'm currently trying to make my first baby-steps in the world of Flex, so please bear with me :-)
