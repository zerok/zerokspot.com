---
date: '2007-01-19T12:00:00-00:00'
language: en
tags:
- stikkit
- web2-0
title: 'First impressions: Stikkit'
---


The idea behind __[Stikkit](http://stikkit.com)__ is actually brilliant: A simple PIM and planning tool using text parsing to set variables (like deadlines, affected contacts etc.). Just a short example:

<pre>
	Let's do a small meeting today 

	share with Horst
</pre>

This would create a new event for "today" and inform "Horst" about it by parsing your input on-the-fly using JavaScript, a process called marked with "Thinking...". This way you can also create todo elements, contacts and so on. 


-------------------------------


<img src="http://zerokspot.com/uploads/stikkit-variables.png" alt="Variables" class="figure"/>

Stikkit also offers various listings to give you an overview about future appointments, events in general, todo elements, contacts etc.

<img src="http://zerokspot.com/uploads/stikkit-menu.png" alt="The main menu" class="figure"/>

Another nice idea is, that you don't have to activate your account right away, which comes in quite handy if you have a slow ISP when it comes to receiving and forwarding mails ;-) For these occasions you get a temporary account and can use the basic features of Stikkit right away until you log out. After logging out you will have to finally activate your account by clicking a link you got right after the registration.

<img src="http://zerokspot.com/uploads/stikkit-temp.png" alt="Temporary accounts" class="figure"/>

The site has a big problem, though: You smell the BETA (or even ALPHA) everywhere (apart from the design). 

* For now it seems like you can't change your password and e-mail address
* "Thinking" limited to English input
* Apart from the timezone no l10n at all
* Variables can only be set by "thinking", not with explicit fields
* <s>When you delete a stikkit, it will be removed by the system using an AJAX request, but not removed from the listing before the next _real_ request.</s> (seems to have been fixed already)

But don't get me wrong. The idea is great enough, to warrant coping with those problems. But it definitely depends on what your normal workflow is. I, for example, normally use [OmniOutliner](http://www.omnigroup.com/applications/omnioutliner/) and [TextMate](http://macromates.com) for 100% of my notes and articles. And for coordination I use e-mail nearly exclusively. Changing this workflow to integrate Stikkit for group coordinations is not possible yet for me, since it would force English even on completely German groups. And if I had to use different tools for different groups it would basically kill the purpose of Stikkit for me.

__Note:__ These are just my impressions after using Stikkit for a few hours.

But the moment Stikkit also support other languages apart from English for their "Thinking" JavaScript, it would absolutely become an alternative for me :) So I'll definitely keep an I on it.