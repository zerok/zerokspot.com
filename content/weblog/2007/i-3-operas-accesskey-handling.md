---
date: '2007-04-16T12:00:00-00:00'
language: en
tags:
- accessibility
title: I <3 Opera's accesskey handling
---


I really like to browse the web using the keyboard only, esp. when I'm working not at home and on my <del>laptop</del> notebook. For these occassions accesskeys have always made my life much easier. 


-------------------------------


Since a couple of versions now (I don't really know when it was added though, but I guess it happened somewhen between 8.5 and 9.2 ;-) ), Opera has a new interface for supporting them. You still have to first switch into an accesskey mode (using Shift+ESC by default) and then hit the actual access key. The new aspect of it is, that you get a menu showing every single accesskey on the page after entering the accesskey mode.

Until now I always hated Opera for requiring something like a special mode for access keys (and I don't see "loving VI" as a conflict here), but not I finally can see why this should indeed make things easier :-) Finally no more guessing if and what accesskeys are supported on a site.

Here just a small example of what this menu looks like:

<a class="thickbox figure" title="Accesskeys on stikkit" href="http://zerokspot.com/uploads/operaaccess.png"><img src="http://zerokspot.com/uploads/operaaccess.small.png" alt="Accesskeys on stikkit"/></a>

As you can see, every accesskey also comes with a description of what it actually does. There is a small difference between the handling of form buttons (like submit and button) and normal links:

1. For normal links the description is taken from the title attribute It *does not* seem to handle for example the alt attribute of a single image inside of such a link that doesn't have. As some kind of fallback solution for links where no title is present, Opera seems to use the href attribute in order to give the user some idea of what to expect what might happen.
2. Form buttons on the other hand first and foremost take the value of the button, which is in my opinion good ... as a fallback like href was for links, but it should normally take the title attribute if present. But it doesn't. So perhaps something for the feature request queue on opera.com :-) 

At least that what it looks like to me, and apart from the problem with prefering \@value over \@title in form items I think this is the best accesskey implementation in a more or less mainstream browser when it comes to being accessible. Great work Opera Software :)
