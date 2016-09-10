---
date: '2005-08-09T12:00:00-00:00'
language: en
tags:
title: BlogCounter.de
---


Thanks to <a href="http://web43.server-drome.net/blog/">Marcel Adrian</a> I found a new service to try: <strong><a href="http://www.blogcounter.de">Blogcounter.de</a></strong>. Here you can register for free to get quite detailed statistics about the visitors of your weblog. All you have to do is enter your (hopefully still available) username and e-mail address to get access to this service. No realname, no address. Actually, it isn't all that new having as far as I can see all known german bloggers listed but it's (as far too often for my taste) new for me :) 



-------------------------------



<img src="http://www.zerokspot.com/uploads/bc_styles.png" alt="Available buttons" class="left"/> After registering you have to add a small javascript and image to your weblog so that the stats are counted. This image is available in 6 different styles so there should be something for everyone.

BlogCounter.de collects quite detailed statistics including the google search string that brought some users to your site up to the screen resolution and color depth as well as the used browser. While this is nice it would be quite useless if it would also count your own visits to your site, wouldn't it? For this case you have 2 options: You can ignore specific IPs or you can install a cookie in your browser that will disable the stats-image when you visit your weblog. After finishing the initial setup both options can be found in the "HTML-Code" section instead of the "Einstellungen" section which took me a few minutes to realize but anyway....

The site offers the user 4 sections for viewing the collected statistics:

* <em>Statistik</em>: Here the user sees single like the top refs, top google search strings and things like that.
* <em>Grafische ßbersicht</em> offers 2 graphics showing the visitors per day and per hour.
* <em>Browser / OS</em> shows... guess what: Statistics about which browsers and operating systems your visitors are using.
* <em>Log</em> offers you a listing of your visitors hostmasks and times of arrival and which document they accessed.

<div class="figure"><img src="http://www.zerokspot.com/uploads/bc_logs.png" alt="Log view"/><p>Log view</p></div>

While I think this service is great I some security concerns:

* When you register or change your password, the system mails you this new password in a not-encrypted email. Given the open nature of the whole e-mail system this isn't such a good idea and it also indicates that the password is perhaps stored in plaintext (or at least not using a secure oneway-hash) in the database. If this is the case, I'd suggest (more than normally) to use for this service a password other than what you use anywhere else.
* The URL of the image to add to your weblog also holds your username. The user-id or something similiar would be better since it won't give a possible attacker the first half of the login information required to access an account.