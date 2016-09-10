---
date: '2005-12-06T12:00:00-00:00'
language: en
tags:
- english
title: Linking Fortran objects into C code using GNU tools
---


I'm currently attending a course about programming languages used for mathematics you I also inevitable got to know Fortran. Everything worked like a charm under Linux but now that I've moved to MacOSX as client system I also wanted to get it working under this OS. 

-------------------------------



Well, let's say you have a library in Fortran named testlib.f and a main program written in C with the name main.c. Normally you would make a Makefile more or less like this:

<pre class="code">
main: main.o testlib.o
	g77 -o main main.o testlib.o
main.o: main.c
	gcc -c main.c
testlib.o: testlib.f
	g77 -c testlib.f
</pre>

Looks OK and would probably work under Linux. But not so on Tiger. There you will get something like this as error:

<pre class="output">
/usr/bin/ld: Undefined symbols:
_fprintf$LDBLStub
_fscanf$LDBLStub
_printf$LDBLStub
collect2: ld returned 1 exit status
make: *** [interpolation] Error 1
</pre>

This is probably caused by different versions of the compiler used for compiling the Fortran library and the C main program...
<pre class="output">$ gcc -v
gcc version 4.0.0 20041026 (Apple Computer, Inc. build 4061)</pre>

<pre class="output">$ g77 -v
gcc version 3.4.2</pre>

Indeed. Solving this is actually quite simple since gcc-3.3 (which is also bundled with Tiger) seems to be compatible with g77 of the gcc-3.4.2 package :-) So simply replace the gcc calls in the Makefile with gcc-3.3 calls, and it should work :)