---
date: '2012-05-17T12:00:00-00:00'
language: en
rating: 4
tags:
- books
- review
- javascript
title: Maintainable JavaScript
---


When you're working on your own on a project with JavaScript you usually know how you like your code formatted,
structured and organized. But once you are part of a team, be it on a spare-time project or during office hours, there
needs to be consensus about this not only for you but also for the rest of the team. Only very few things are a bigger
waste of time then ending up with badly formatted or structured code because everyone was adhering to their own 
personal coding style.

[Nicholas C. Zakas][nz] with his new book ["Maintainable JavaScript"][book] tries to prevent this by comparing multiple
coding style-guides and describing common JavaScript coding practices as well as how to automate things like testing
against these best practices and he does a pretty good job at that. I bought this book as in the "Early release"
version, so some of the points I make down below might not be applicable anymore.

------------------------

The first part compare different style-guides and does some cherry-picking here and there. The
result is very informative and luckily comes pretty close to the style I personally prefer. Also, the author does
mention JSHint and JSLint right from the get-go. This first part might perhaps seem a bit basic for most readers but
the comparison between the JQuery coding guidelines and those of other projects make it quite an informative read.

The second part describes common coding practices like avoiding global variables, loose coupling, exception handling 
and so on. Personally, this was basically the part of the book I was expecting when reading the title and I was not
disappointed. There was some very good stuff in there like how to do exception throwing and handling in a cross-browser
compatible manner. Also the drawbacks and advantages of browser-detection vs. feature-detection where described very
clearly.

For some of the topics discussed in this part there are already tons of solutions in the popular JS frameworks like
JQuery, Dojo and YUI available. Where appropriate the author includes examples from those libraries.

One thing I would have preferred reading more of here was the whole modularization of JavaScript development. The
author goes into that when explaining how to get rid of global variables but this and how to clearly structure a 
larger project would have been a killer-chapter for me. There is also no mention of the whole CommonJS-approach on 
modularization which goes hand-in-hand with a total lack of server-side discussions. There is a short chapter on AMD 
and also RequireJS gets mentioned but more would have been greatly appreciated.

Perhaps the third part could have lost some pages to put those into part two ;-) This third and last part of the book 
deals with automating what was taught before and making it ready for prime-time via deployment. The tool of choice 
here is Ant, which receives IMO a bit too much focus here, but given that quite a lot of tools around the JS ecosystem 
are Java-based this makes sense. That said, it kind of made me skip most of this part since I have already worked with 
most of the tools mentioned in this chapter before.

----------------------------------

If you're looking for a book about best practices and coding style-guides for client-side JavaScript I can really 
recommend this one. Everything is well explained and well structured. There are some rough edges here and there but overall it is an enjoyable read esp. thanks to the second part.

[nz]: http://www.nczonline.net/
[book]: http://shop.oreilly.com/product/0636920025245.do
