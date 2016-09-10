---
date: '2005-06-04T12:00:00-00:00'
language: en
tags:
title: Show PL/SQL compile-time errors in ISQL
---


Developing PL/SQL in a webinterface is a pain... ok, developing anything in a webinterface is a pain but if you still one of those lucky people out there and could kill ISQL for not displaying any compile-time errors, here is a small hint I found after googling a little bit:

-------------------------------



<pre class="code">SHOW ERRORS;</pre>



From what I can see it displays all errors the last query/definition/whatever produced :-)