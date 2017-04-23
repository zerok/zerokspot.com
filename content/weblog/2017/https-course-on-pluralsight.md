---
date: 2017-04-23T09:19:08+02:00
title: Learning about HTTPS from Troy Hunt
tags:
- learning
- pluralsight
- video
- security
- development
---

While it might seem like handling HTTPS is easy now thanks to Let's Encrypt,
especially migrating sites to it is far from it. In addition to that, a lot of
web developers still don't even know the basics about HTTPS and cling to some
old myths about performance and the costs involved. Luckily, [Troy Hunt][th] has
published a new course on [PluralSight][] a couple of weeks ago
called [What every developer must know about HTTPS][course], which addresses
these concerns.

<figure>
<img src="/media/2017/pluralsight-https.png" alt="" />
<figcaption><p>The overview page for the course.</p></figcaption>
</figure>

I've spent the last week watching all 3.5 hours of this course and cannot
recommend it enough to anyone doing web-development! I'm pretty sure, *everyone*
will learn something new here, even if you've already worked with HTTPS on a
couple of projects. This course won't teach you about all the different
crypto-suits and algorithms and when to enable which to support specific clients
(that's what [Mozilla's SSL Configuration Generator][scg] is there for),
though. Instead, it will provide you with the knowledge you will need to get
going and progress from there. It will also introduce a handful of tools so that
you can check yourself that you're doing it the right way 😊 For me personally,
the part about CSP was the most interesting as it will teach you how to migrate
a site properly and all the options you have for that.

This was also the first time I attended a PluralSight course and I absolutely
loved it. The segments had just the right length and the presentation as
top-notch. In addition to that, there is a short test for after you've finished
the course to test your knowledge. It's simply a great experience!


[th]: https://www.troyhunt.com/
[course]: https://app.pluralsight.com/library/courses/https-every-developer-must-know/table-of-contents
[pluralsight]: https://www.pluralsight.com
[scg]: https://mozilla.github.io/server-side-tls/ssl-config-generator/
