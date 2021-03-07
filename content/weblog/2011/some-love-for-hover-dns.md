---
date: '2011-08-08T12:00:00-00:00'
language: en
tags:
- web
title: Some love for Hover's DNS management
---


Recently, Kenneth Reitz wrote a little [blog post][5] about a bunch of tools he's
using. One of the items on his list was [DNSimple][4] which looks like an awesome
tools for managing DNS entries for multiple domains with a single interface
and single DNS infrastructure independent of your registrar or hosting
service. 

The funny thing is, up to a couple of weeks ago I'd have gladly spent a couple
of USDs on something like that but then I accidentally went into my current
ROC's (registrar of choice = [hover.com][2]) backend and remember that I never really
played around with their new UI before. And I've missed quite a lot ...

------------------------------------------

Previously, [hover.com][2] had one of the worst yet still luckily simplistic
interfaces for DNS management in a registrar I've ever seen. I have to
emphasize on the "simplistic" here since the worst interface IMO is still the
one provided by [domainfactory][1], in my opinion ;-) The biggest issue I had with
hover's previous interface was that the main menu was restricted to a
"currently managing" domain, a fact displayed only in the smallest of readable
fonts.

<figure>
<img src="/media/2011/oldhover-managing.png" alt="" />
<figcaption>The old managing interface for domains on hover had something like
an "active" domain, but hid that fact pretty well...</figcaption>
</figure>

The next little obstacle in the course towards managing your DNS was that you
had to basically confirm that you had at least the slightest idea of what you
were doing. I still don't get why they deemed something like that necessary.

<figure>
<img src="/media/2011/oldhover-overview.png" alt="" />
<figcaption>In order to actually get to the DNS management interface you had
to hit another confirmation link...</figcaption>
</figure>

Luckily this all changed a couple of months ago with a new interface being
made available in some kind of pre-roll-out similar to [#newtwitter][3] back in the
days. And boy, tons has changed:

* No more stupid domain changing before the main menu has any kind of meaning
* No more confirming that you know what DNS means
* No more ugly 1999-table-designs

What you get now is a rather slick interface with tons of aesthetic bling but
without really trying to jump you into your face. For instance the account
settings for a domain or now easily accessible through a simple form with a
handful of collapsed subforms. No unlocking or something. It's like they trust
you to know what you're doing if you've made it this far:

<figure>
<img src="/media/2011/hover-manage.png" alt="" />
<figcaption>All domain settings in one place.</figcaption>
</figure>

Also the DNS interface is now far cleaner with some filters for finding your
entries and ... it just looks way nicer.

<figure>
<img src="/media/2011/hover-dns.png" alt="" />
<figcaption>New DNS manager on hover</figcaption>
</figure>

And if you just have one or two domains, there is even a "global" DNS
interface for all your domains.

I have to say, I'm extremely pleased with these changes ... otherwise this
post would probably not exist. That said: Over the long run I will probably
need something like DNSSimple anyway. As much as I like hover so far and I
definitely feel like I get my money's worth here, the idea of having one DNS
managing all my domains independent of host and registrar is a big seller for
me. A couple of months ago I also looked at [PointHQ][6], but I somehow didn't
like the UI. I totally understand that they want to upsell you to their
premium account, but this was a bit too much for me :-) Anyway, this and/or
DNSimple is probably something for another post.

[1]: http://df.eu
[2]: http://hover.com
[3]: http://twitter.com/newtwitter
[4]: https://dnsimple.com/
[5]: http://kennethreitz.com/i-use-this.html
[6]: https://pointhq.com/
