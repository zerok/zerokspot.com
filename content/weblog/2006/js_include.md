---
date: '2006-05-09T12:00:00-00:00'
language: en
tags:
- include
- javascript
- webdevelopment
title: js_include
---


[J Wynia](http://www.wynia.org/wordpress/) has written a nice little [include function for JavaScript](http://www.wynia.org/wordpress/2006/05/04/javascript-includes/) basically embedding the included-script nicely into the head element of the displaying HTML page. So all you have to do is puts something like `js_include('path/to/file.js');` into one of your already "installed" JavaScript files. Definitely much nicer than dothing something

-------------------------------

ugly like `document.write('<script type="text/javascript" src="http://domain.com/j/newlib.js"></script>');`.