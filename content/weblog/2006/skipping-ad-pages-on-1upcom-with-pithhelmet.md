---
date: '2006-01-15T12:00:00-00:00'
language: en
tags:
- macosx
- web
title: Skipping Ad-pages on 1up.com with PithHelmet
---


Honestly I couldn't visit sites like [1up.com](http://www.1up.com) anymore without tools like [PithHelmet](http://www.culater.net/software/PithHelmet/PithHelmet.php) thanks to the thousands of Flash ads that nearly kill basically any browser out there. Still, classic ad-blocking can't really help you with dedicated ad-pages that 1up.com triggers every now and then. But there's a quite easy way to get around this using PithHelmet's MacheteScript feature. This simple Perl script parses the passed page and adds a ne reload element into it if the page looks like one of the ad-pages:

<pre class="code">#!/usr/bin/perl -w
while(&lt;&gt;){
        if(/CONTINUE TO 1UP/){
                print &apos;&lt;meta http-equiv=&quot;refresh&quot; content=&quot;0&quot;/&gt;&apos;;
        }
        print $_;
}</pre>

(Sure, this is not really a valid approach, but it works.)



-------------------------------



## Installation
Simply store this script somewhere on your HD with the .pl extension and make it executable using Finder or chmod in the commandline.

Now add a new ruleset in PithHelmet with the Wildcard match "*1up.com*" (sure, this will also hit other pages like 21up.com etc. but I'm too lazy for a decent regex). Enable the "default site preferences", go into the advanced tab and select the Perl script you've saved in the previous step.

Ah, and if you haven't enabled MacheteScripts in the PithHelmet preferences, you have to do this now and restart Safari.