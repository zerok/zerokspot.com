---
title: Meetup.com changes... oh my
date: "2019-10-15T19:10:00+02:00"
tags:
- meetup
- api
- community
---

Over the last couple of months, Meetup.com has seen quite a few
changes. Sadly, though, I'll have to focus on two of them for now that
are not positive:

1. The removal of the API key system
2. An experimental new pricing model

## No more API keys

For the last couple of years I've been running a little service for
GoGraz that fetches all the signups for a particular event and
displays them directly on GoGraz.org. For the longest time, you could
simply go to https://secure.meetup.com/meetup_api/key/, grab/generate
an API key and use it in your requests.

That changed on 15 August 2019 (after having been announced on 14 June
2019 on the [mailinglist][ml]). Since then, only OAuth2 is available
as authentication method agaist the API and you need to generate an
OAuth2 consumer for your application. While the interface for
registering such a consumer is available to all members, the FAQ
basically states that only Pro subscribers can request new consumers
(all others weill be rejected right away) and even those are not
guaranteed to be gren-lit.

<figure>
<img src="/media/2019/meetup-consumers.png" alt="">
<figcaption>I do not have a Pro account yet I can request new OAuth consumers.</figcaption>
</figure>

Luckily, for my use-case I don't seem to need auth anymore, but
waiting for 3-5 business days just to be reject because I don't have a
Pro account feels like a waste of time at best and just bad UI/UX. I
shouldn't even be able to fill out that form...

## New price model?

Thanks to [this recent post on the changelog][cl] I stumbled upon
[this page][mu] on meetup.com where they describe some upcoming
changes to the payment system of a certain subset of groups. They
clearly state that it is just an experiment right now and limited in
scope but it's hard to imagine any opensource-meetup being able to
operate under these conditions:

1. Meetup organizers pay only 2 USD per month for a pro account
2. Attendees pay 2 USD each per event

<figure>
<img src="/media/2019/meetup-pricing.png" alt="">
<figcaption>Screenshot of <a href="https://www.meetup.com/lp/paymentchanges">www.meetup.com/lp/paymentchanges</a></figcaption>
</figure>

Especially for smaller meetups that are just starting this would be a
death-sentence. I also hadn't realized that WeWork had acquired
Meetup, but that explains a lot ... That being said, Meetup clearly
stated that they won't do any pricing changes in the near term. Even
having something like that on the table, though, sends a chilling
effect down my neck.

The OSS community seems to be already working on [multiple possible
replacements][r] for Meetup with the [FreeCodeCamp][f] folks just
being the latest in that list with [Chapter][c]. To be honest, though,
all I want is something like this:

- As an organizer I create a new event "somehow"
- Possible attendees get a simple form where they can enter their
  e-mail address and their RSVP status.
- They receive a mail to confirm their signup. 
- That email also contains links with which they can change or even
  cancel their signup status.
- A simple API will allow a website to render a simple widget with all
  the attendees.
  
The login-system at [changelog.com][cc] was a big inspiration for
those points. There is just something nice about getting one-time
logins via e-mail. I don't have time to work on something like that
right now, though. Perhaps in December 🤣

[mu]: https://www.meetup.com/lp/paymentchanges
[cl]: https://changelog.com/news/so-long-meetup-and-thanks-for-all-the-pizza-M3Z2
[ml]: https://groups.google.com/forum/#!topic/meetup-api/R__7mPzWJc0
[r]: https://github.com/coderbyheart/open-source-meetup-alternatives
[c]: https://github.com/freeCodeCamp/chapter
[f]: https://twitter.com/ossia/status/1183845054449930241
[cc]: https://changelog.com/
