---
title: "News without archive"
date: 2019-08-23T10:26:51+02:00
tags:
- news
- archival
---

Sorry, a bit of frustration-venting coming up:

One of the most popular if not even *the* most popular website in
Austria is <https://orf.at>, the website of Austria's public
broadcasting company. It has usually around 80,000,000 visits each
month according to
[ÖWA](https://der.orf.at/medienforschung/online/index.html) and is
also a quite popular source for news articles. Basically, if someone
lives in Austria, the chances are quite high that this person visits
orf.at daily to get their news.

While the website is extremly speedy and rich on high-quality content,
it has two big issues:

## The "ORF Act"'s 7-day-clause for online content

> The individual elements of news overviews shall be provided only for
> as long as they are topical but no longer than seven days from the
> date they were first provided forviewing on the platform of the
> Austrian Broadcasting Corporation.
>
> [...]
>
> Programmes shall be provided for viewing or listening without the
> possibility of saving them (with the exception of podcasts) for a
> period of up to seven days after they were broadcast
> 
> -- "Special mandate for an online service", §4e (2 and 4), [ORF Act](https://zukunft.orf.at/rte/upload/isabelle/federal_act_on_the_austrian_boradcasting_corporation.pdf)


These two paragraphs more or less prohibit the ORF from making
content associated with their primary content (e.g. news coverage) or
even recordings of their podcasts available for more than 7 days after
the original publication date.

If I want to listen to an episode of the [Ö1 show
"Matrix"](https://oe1.orf.at/matrix) from two months ago, I
can't. It's simply not there anymore. The same is true for pretty much
every single Podcast or recording that the ORF is offering on their
website.

{{<figure src="/media/2019/orf-matrix.png" caption="No real podcast archive">}}

Same goes for links tweeted by the various ORF accounts. Ö1's primary
news show, for instance, regularly posts a short blurb + a link to the
broadcasting. On 12 June they made [this
tweet](https://twitter.com/oe1journale/status/1138751725811195904) and
linked to <https://oe1.orf.at/player/20190612/556250/120636000>
through bit.ly. Since it refers to an audio stream that they are no
longer allowed to provide, the website simply redirects to the player
itself. They don't even send a 30x status code or anything. The
content is simply gone.

You might have some luck contacting their [customer
service](https://der.orf.at/kundendienst/service/servicenummern100.html)
but it's probably better to create your own archives if you like some
of their podcasts...


## Hidden archives

The other big issue with [orf.at](https://orf.at) is that content that
is no longer on the frontpage, is either no longer accessible at all
or hidden on topic-specific subdomains like <https://science.orf.at>
without direct archive-links from the frontpage. If you don't find
this subdomain through another article or through the sitemap, you
won't know that it's even there.

{{<figure src="/media/2019/orf-subdomains.png" caption="Subdomains only in the sitemap">}}

Some content items aren't living within their own thematic subdomain,
though: <https://orf.at/stories/3131884/>. You can still access this
article even though it was created on 29 July, but you won't find it
through a pure clickpath from the frontpage. Luckily, orf.at is quite
well-indexed by external search engines like
[DuckDuckGo](https://duckduckgo.com/?q=site%3Aorf.at+Erste+Generalprobe+f%C3%BCr+Berliner+Flughafen+BER&t=ffab&ia=web)
but not having a working archive of news articles on one of the
primary sources for news in Austria is really bad.

It makes researching news about a given time-frame hard to impossible
and forces you to use other sources instead of the biggest one in the
country.

Just yesterday, my partner wanted to send me a link to an article she
had stumbled upon in the morning about [mathematics and
online-hate](https://science.orf.at/stories/2990262/). By the time she
got to it, it had already vanished from the frontpage. We were still
lucky to find it on the science-subdomain a day later, but if it had
been a news-story, it would have been lost.

## Dev team?

In recent months, the whole website received a bit of tuning without
more pleasant fonts, more whitespacing etc.. That and other
navigation-related usability aspects haven't been touched, though, as
it seems. At this point I'm not sure how large the development team
behind the site is. [dev.orf.at](https://dev.orf.at) hasn't been
updated since Jan 2017 but at least their [Twitter
account](https://twitter.com/devorfat) is still active. On the
[imprint](https://orf.at/stories/impressum/) they also link to a
project called Helga, but the link is broken.

I really like a lot of the content being produced by ORF and Ö1 and
FM4 in particular so I really want the site to be good. Right now (and
for the recent history) it hasn't been, though. It's servicable, it
gets the job done, but it doesn't shine while doing so. At this point
I get the impression, that their website offering (not their on-demand
streamin service) has an extremely low priority. I couldn't even find
any related full-time job offerings on their [jobs
page](https://der.orf.at/jobs/).


## Implications

The current state of ORF.at has a few implications:

- When I want to know what happened during a specific time-window in
  Austria, I now usually visit other news outlets instead of the
  biggest and most popular one, simply because it doesn't have a
  working archive.
- Older audio programs are no longer accessible after 7-30 days. I've
  therefore started to archive my favorite podcasts from them on my
  own. Given that every Austrian citizen with a TV/Radio is supposed
  to contribute financially to the ORF, a publicly accessible archive
  o their recordings should have been a no-brainer. I don't understand
  why I, as an individual, have to do that. That being said, it's
  awesome that they publish recordings as normal podcasts so that I'm
  able to archive them! Unlike the previous point, this one is not meant
  against the ORF but only against the cited clauses of the ORF Act.
  
At this point I'm not sure, what I as an individual can do to
help. Should I contact some members of the parliament? Should I talk
to the ORS? If you're working for the ORF and can tell me, please let me
know!
