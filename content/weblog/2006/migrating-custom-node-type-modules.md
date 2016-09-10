---
date: '2006-11-02T12:00:00-00:00'
language: en
tags:
- 5-0
- development
- drupal
- module
- php
title: Migrating custom node type modules
---


Today I finally started porting the modules I had written for zerokspot.com from Drupal 4.7 to 5.0. So far everything has worked out just fine apart from one little thing:

I added the modulename.info file to my book review module and there it was in the admin panel. Everything's great. Then I wanted to create a new node with this type and noticed (1) that it doesn't appear in the main listing but only in the navigation sidebar and (2) that when I try to submit the new node, I get an error similar to this one:

<pre class="output">array_merge_recursive() [&lt;a href=&apos;function.array-merge-recursive&apos;&gt;function.array-merge-recursive&lt;/a&gt;]: Argument #2 is not an array in /opt/wwwdev/htdocs/drupal-5.0/modules/node/node.module on line 1916.</pre>

-------------------------------



So I checked my modules hook\_form implementation, but couldn't really find where the $form array should _not_ be an array ;-)

Then finally I got the idea to search the forums on drupal.org for a solution and found [this one](http://drupal.org/node/91892#comment-168077). I would never have searched in the hook\_node\_info implementation. Thanks [yodadex](http://drupal.org/user/18468) :)

The next time I should really try to apply _all_ the changes mentioned in the [module conversion guide](http://drupal.org/node/64279) before getting desperate ;)