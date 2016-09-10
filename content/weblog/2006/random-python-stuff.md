---
date: '2006-09-13T12:00:00-00:00'
language: en
tags:
- development
- python
title: Random Python stuff
---


Some random Python hints I learned over the last week while writing on my first Django website (Note: All things mentioned here are in the official docs, but I haven't noticed them before looking at code others have written):

If you want to have a tuple, better play it safe and write ("hello",) instead of ("hello"). Since the braces can also be used to control the precedence of a statement, adding the extra "," will make sure, that you really get a tuple, no matter what number of elements is in there.

[Named groups](http://docs.python.org/lib/re-syntax.html) is in my opinion one of the coolest things I've seen with RE done so far. I don't know, if this also works in other languages, but it just makes regex much more readable if you have something like this:

-------------------------------



<pre class="code">import re
m = re.match(r&apos;(?P&lt;name&gt;.*)&apos;,&apos;hello&apos;)
print m.group(&quot;name&quot;)</pre>

Something similar is possible within string substitutions:

<pre class="code">print &quot;Hello %(name)s&quot;%{&apos;name&apos;:&apos;Horst&apos;}</pre>

You give %s basically the name "name" and assign its value using the passed dict instead of a tuple. Details about this can be found [here](http://docs.python.org/lib/typesseq-strings.html).