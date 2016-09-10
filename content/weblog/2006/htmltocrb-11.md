---
date: '2006-05-07T12:00:00-00:00'
language: en
tags:
title: htmltoc.rb 1.1
---


[What is htmltoc.rb?](http://weblog.zerokspot.com/posts/633/)

-------------------------------



Changes in this version:



* You can now specify how deep the ToC should be with the -l LIMIT (or --limit LIMIT) flag. For example if you don't want to include h4,h5 and h6 elements in the ToC, simply append --limit 3 

* The ToC doesn't include empty ul-tags anymore if all children are marked as hidden



Another change is where you can get htmltoc.rb from now on:



**[http://zerokspot.com/code/htmltoc.rb/](http://zerokspot.com/code/htmltoc.rb/)**



This is also a darcs repository so if you want to get older versions, you need darcs :)