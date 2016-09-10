---
date: '2006-08-15T12:00:00-00:00'
language: en
tags:
- drupal
- howto
- python
- scripts
title: What Drupal version do you have installed?
---


How do you determine what version of Drupal you've installed? Especially with older versions this might become a little bit problematic since the VERSION constant in the system.module was only introduced with Drupal 4.7.x.

A classic method is to simply look at the changelog of your installation ... if you still have it around. If not, you will have to start with some tricks.  [iedude](http://drupal.org/user/69666) asked what to do in this case [on the Drupal board](http://drupal.org/node/78237#comment-145302) and I started digging around a little bit. 



-------------------------------



A few questions you should ask yourself:

1. How have I installed Drupal in the first place?
2. Did I have the contact.module in core?
3. Do I have the VERSION constant in the system.module?

Answering (1) will probably help the most if you're using some package management systems like apt-get and always keep your system up to date. Then you're probably using the latest available version in this repository.
If (3) is true, then you probably have 4.7.x installed. If (2) is false, then you probably have something prior to 4.6.0 installed. These questions should help you figure out, what major release you're using, but they won't help you with the patch-level. Now I think your best option is to check the $Id lines of each PHP source file of your core, to see, what was the latest change, that made it into the release. A simple script in your favorite language could do the job for you from here on. 

Let's say, you're quite sure, that you're using a 4.6.x release and your script gave you following output: "2006/06/01 21:55:58". Then you're probably using Drupal 4.6.8 since this [was released](http://drupal.org/drupal-4.7.2) on 2006/06/01.

I've written a small script myself and you can have and use it at your own risk ;) Simply start it with ...

<pre class="command">python drupal_latest_change.py path/to/drupal</pre>

It only checks the .inc and .module files and should ignore files in "contrib" folders. So if your modules folder is a mess, this script won't really help you.