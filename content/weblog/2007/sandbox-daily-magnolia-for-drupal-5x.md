---
date: '2007-02-23T12:00:00-00:00'
language: en
tags:
- dailymagnolia
- drupal
- ma-gnolia
- module
- sandbox
title: 'Sandbox: Daily Ma.gnolia for Drupal 5.x'
---


It's been a little bit quiet here for the last couple of days. This is mostly because of the exam I had today and I needed to learn quite a lot for it. Other factors were _Okami_ for the PS2 and a small module I'm currently writing on which I want to present here :-)

First of all, what is this module all about? It's a tool that will integrate into Drupal's cronjob mechanism, fetch your recent bookmarks from the social bookmarking site [Ma.gnolia](http://ma.gnolia.com) and post them as new nodes on a Drupal site. A few of the features and limitations are now listed below:


-------------------------------


* It is optimized for single-user websites like zerokspot.com where you basically have one person posting the content and having ultimate control of the site. Therefor the security aspects of this feature will be limited to simply checking whether someone has the right to administer nodes or not.
* It stores the links as XML into the node and displays them using a custom filter in order to reach a maximum of themability.
* _Daily Ma.gnolia_ can be configured to use the tags associated with each link, aggregate them and add them to the node. This requires that the admin selects a vocabulary in the module configure that allows free tagging and multi-selects.
* The whole rendering of the node can be customized.
	* For the title field the module requires the _token_ module to allow the admin to specify how the date in the header should be formated, where it should be placed etc.
	* As already said above the actual content of the node gets filtered using a custom filter which also involves two theme functions depending on whether the teaser or the full page is rendered.

Last but not least something about the "Sandbox" title: This module will be highly specific and therefor probably not really useful for the masses (considering Drupal site admins as "masses" here :-) ). Therefor I see this module more as some kind of foundation for other people to hack on and customize for whatever _they_ need. Therefor the code will probably be released under the MIT license and there will be a Mercurial repository on this server and a Trac installation for it (all that depends on whether I finally get rid of my release-phobia).

Some screenshots :-)
<a rel="gallery_dailymagnolia" href="http://zerokspot.com/uploads/dailymagnolia_teaser.png" class="thickbox figure" title="Teaser view"><img width="320px" src="http://zerokspot.com/uploads/dailymagnolia_teaser.png" alt="Teaser view"/></a>
<a rel="gallery_dailymagnolia" href="http://zerokspot.com/uploads/dailymagnolia_page.png" class="thickbox figure" title="Full page view"><img width="320px" src="http://zerokspot.com/uploads/dailymagnolia_page.png" alt="Full page view"/></a>