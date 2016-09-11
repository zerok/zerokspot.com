---
date: '2014-01-19T12:50:06+01:00'
language: en
tags:
- fastmail
- email
- gmail
title: From Google To Fastmail
---

I've been using *Google Apps for Domains* for [nearly 7 years][2] now and it has served me well over the years. Except for some minor issues it was my go-to location for when I bought a new domain and wanted to put hosted emails behind it. Then, in December 2012 Google announced that they would [discontinue the free version][1] of this service most likely in favor of pushing more people towards the simpler default GMail service.

## Looking for alternatives

Since then I've been looking for an alternative. My main requirement (obviously besides the core e-mail package) is that I want to manage also the e-mails for my domains with it. Eventually I ended up with two options:

* [Fastmail][6]
* [Runbox][7]

Both of them are feature-wise more or less on the same level with both offering 15GB plans with custom domain support. The main advantage of Runbox is that their servers are in Europe while Fastmail's are located in the USA. Given the whole NSA spy scandal, Fastmail even wrote a [lengthy post][3] about what this means for its users. The advantage of Fastmail is the price and the (according to [comparable screenshots from Runbox][8]) better webmail client. For what I wanted out of it, Fastmail only costs USD 40 per year compared to USD 79 at Runbox. For the last couple of months I've been reading many discussions on [emaildiscussions.com][9] about both of them and (as with everything else on the net) both have enough haters and lovers to remove this a relevant criterium.

One thing that made me nervous, though, was that Runbox seems to run [multiple versions of their web frontend][10] at the same time with different security requirements. They also seem to not yet support setting up two-factor authentication.

So after thinking quite a bit about it and using the trial for a couple of days I decided to go with Fastmail.


## The migration

I did the migrations in two steps. First, I changed the DNS entries for my mailserver over to Fastmail's mail server so that new mails would arrive either in my old Google Apps inbox or already in my Fastmail inbox. Initially, I was afraid that FM would only support full DNS hosting because the first page you hit then looking for it in the documentation is one that tells you to [change the DNS servers of your domain to Fastmail's][4]. Good thing there is also the ["advanved" approach][5] where you can just change just the MX entries and be done with it.

Once that was in place I started the IMAP migration tool FM provides and let it run for about 2-3 hours. Once the tool has done all the fetching and merging of folders it will tell you the details in an e-mail, marking the end of the migration.

That was the whole migration process for me since I didn't use filters on GMail all that much and wanted to get rid of those that I had anyway.

## Postbox setup

On the desktop I use mostly Postbox to manage my emails. Getting it to work with Fastmail obviously basically consists of creating a new IMAP account there, with one small exception.

<figure><img src="/media/2014/postbox-advanced-settings.png" alt=""></figure>

For some reason the IMAP server directory was left empty which results in your folders not being automatically listed in your account. Setting that to "INBOX." fixed that for me.


## Current status

So far everything looks good here. There some tiny UI issues I have with the way the settings work, but that's all I can complain about right now.

Oh, and Fastmail has a referral system. So if you want to give them a try, please use [this link][11] :-)


[1]: http://googleenterprise.blogspot.co.at/2012/12/changes-to-google-apps-for-businesses.html
[2]: http://zerokspot.com/weblog/2007/12/08/google-apps-for-your-domain/
[3]: http://blog.fastmail.fm/2013/10/07/fastmails-servers-are-in-the-us-what-this-means-for-you/
[4]: https://www.fastmail.fm/help/quick_tours_setting_up_domain.html
[5]: https://www.fastmail.fm/help/domain_management_setup.html
[6]: https://fastmail.fm/
[7]: https://runbox.com/
[8]: https://runbox.com/features/email-services/preview/
[9]: http://www.emaildiscussions.com/
[10]: http://www.emaildiscussions.com/showthread.php?t=67296
[11]: http://www.fastmail.fm/?STKI=11979233
