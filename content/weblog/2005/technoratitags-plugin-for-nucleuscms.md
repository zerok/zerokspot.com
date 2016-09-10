---
date: '2005-02-03T12:00:00-00:00'
language: en
tags:
- development
- nucleuscms
- software
- web
title: TechnoratiTags plugin for NucleusCMS
---


During the last two days I've written on a small plugin for NucleusCMS that adds a new field to the EditItem and AddItem fields in the admin panel to enter Technorati Tags. <a href="http://www.leftoftheweb.com">Stefan Koopmanschap</a> has installed the plugin today on his <a href="http://www.stefankoopmanschap.nl/weblog/">dutch weblog</a>. Because I currently don't have a NucleusCMS-powered weblog I see this as a testcase for this little plugin ;-)

-------------------------------



If everything works out, I will probably release this plugin tonight (after perhaps adding some additional preferences). Currently you still have to change the markup of the tags in the event_PreItem function. Ugly, but it should work ;-)



<strong>Note:</strong> The plugin doesn't ping Technorati so you need one of the <a href="http://wakka.xiffy.nl/PingPong">PingPong plugins</a>.