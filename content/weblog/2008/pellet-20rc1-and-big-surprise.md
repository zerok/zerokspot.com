---
date: '2008-10-29T12:00:00-00:00'
language: en
tags:
- software-release
- java
- licensing
- semanticweb
title: Pellet 2.0rc1 ... and a big surprise
---


If you've ever had to work with some OWL-ontologies, you have definitely heard the name [Pellet](http://clarkparsia.com/pellet/). It is perhaps the most wildly use opensource OWL-reasoner/engine out there. Yesterday Evren Sirin [announced](http://clarkparsia.com/weblog/2008/10/27/pellet-20rc1-release/) the release of the first release candidate of the upcoming version 2.0, that (among other things) includes preliminary support for OWL 2.0. But there is also a catch ...

-------------------------------

Previous versions up to this release candidate were available under the permissive MIT license which made it possible to incorporate it into closed-source applications (or actually basically any application I can think of right now). That has changed now. Since 2.0rc1 Pellet is available under [dual-licensing](http://clarkparsia.com/pellet/dual-license) terms with custom commercial licenses being the first and the [Affero GPL Version 3.0](http://www.fsf.org/licensing/licenses/agpl-3.0.html) being the second option. From what I understand (or believe to understand) this means now that if you want to use Pellet in a closed-source environment, you either have to get into contract negotiations with Clark & Paria, or stick with version 1.5.2 (or work on some pluggable RPC/IPC-solution). Same goes for people who want to use Pellet within BSD/MIT licensed applications.

Don't get me wrong, though. I have nothing against the GPLv3 (although I'm probably not the biggest fanboy in the world if it comes to the Affero General Public Licenses) but I was definitely surprised by this drastic move. Moving from a permissive license like the MIT license to something like the AGPL, which is even more copyleftish than even the GPL ... Well, I definitely wish them luck with this new strategy. I really hope enough companies that can afford it will help sponsor the future development of this library  :-)

This move definitely made me reconsider the use of the GPL for possible future project of my own, but soon realized at least something: I can't use it for everything or it might end up as a grey-area. For example: If I want to write a recipe for zc.buildout and license it under the GPLv3, there is quite a chance that I violate the GPL by doing so (because according to [the GPL-FAQ](http://www.fsf.org/licensing/licenses/gpl-faq.html#GPLPluginsInNF) a plugin that is not executed in a forked environment is actually incorporated into the main program and therefor forms a union with it). 

These are just some observations by someone who prefers the BSD/MIT-licenses for their simplicity and who has really just the absolute minimum of knowledge about licensing, so don't take my word for it ;-) If you are a lawyer and know your way around licensing, please let me know if I misunderstood something here :-)
