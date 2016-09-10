---
date: '2008-07-01T12:00:00-00:00'
language: en
tags:
- commenting
- disqus
- weblog
title: A Look at Disqus
---


<img src="/media/2008/disqus.png" class="left" alt=""/ >I have to say, that I really start thinking about just adding [Disqus](http://disqus.com) to my site here. The problem Disqus is trying to solve is exactly the one that annoys me the most when commenting on someone's weblog: I tend to forget where I commented, so I miss the whole discussion taking place afterwards. For this Disqus offers a unified interface that shows you all your comments on Disqus-enabled weblogs. 

-------------------------------


<div class="figure"><img src="http://img.skitch.com/20080701-tr25emt2yrjmwfsbdnapyjuw3a.png" alt="" /><p class="caption">Logging in using <a href="http://clickpass.com">clickpass</a></p></div>

The user also gets a nice page were she can access all her comments made on Disqus-enabled sites, can log into her account using OpenID and can rate other comments as well as get notified on new comments via email.

From the site-maintainer's perspective this service offers a simple way to integrate comments, offering SPAM protection and centralized moderation, rating of comments and so on. If you want to, you can also access all your site's comments through a currently read-only API, which means that you won't lose your comments if you want to leave Disqus again later on. There is even a separate [documentation page](http://disqus.com/custom-css/) for all the IDs and classes available for styling.

<div class="figure"><img style="width:500px" src="http://img.skitch.com/20080701-ng66qrmiim3hp3sgeccckfe5tm.png" alt="" /><p class="caption">Easy administration of all your blog's comments</p></div>

The system is very easy to integrate (basically all you need is a small HTML-snippet) and also allows testing from localhost. On the other hand you might run into some small problem if you for instance have your comments on a separated page (like here). This is quite easy to resolve, however, but just specifying the "correct" `disqus_url` wherever you are:

@@ html @@
<script type="text/javascript" charset="utf-8">
    var disqus_url = "http://domain.com/weblog/2008/06/30/post";
</script>
@@

This seems to override whatever URL got auto-detected. On the other hand I haven't yet found an efficient way to get the number of comments for not one post but multiple through the API, which is useful for archive pages or index pages on a weblog. So far I could only find calls for single URLs.

There is also currently no way to search your comments or if you are a webmaster to import your old comments. I guess the latter is a little harder to solve here since ideally you'd have to match e-mail addresses with current Disqus accounts.

A more general problem is accessibility here. The whole commenting interface when directly integrated into your weblog is loaded via JavaScript. If JavaScript is not supported, the user gets (when using the default snippet) a link to the respective discussion-page on Disqus itself.

I guess, in the end I'm still torn. Disqus' feature-set is really nice, but it's really quite a different way for handling comments that I'm used to. With Disqus I'd still be in control of who can comment on my blog, just not 100% anymore ... 
