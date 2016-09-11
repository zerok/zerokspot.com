---
date: '2016-02-05T15:28:55+01:00'
language: en
tags:
- travel
- fosdem
- belgium
- brussels
- conference
title: Back from FOSDEM
---

Last weekend it was once again time for [FOSDEM][], so I hopped onto a plane on
Friday morning to Brussels to be there. For details about the trip itself please
check my [travelogue][] which I plan to extend over the next couple of days with
some more information. I still want to write about the main event here simply
because it's highly technical in nature 😉

<figure>
    <img src="/media/2016/fosdem.jpg" alt="" />
    <figcaption><p>Just before the welcome notes at FOSDEM</p></figcaption>
</figure>

FOSDEM was once again held at the ULB and basically everything apart from the
actual talks and tracks was identical to last year's event. So yes, I once again
had far too many waffles, although it is debatable whether or not there is such
a thing. I also finally managed to make it to the beer event on the evening
before the conference. Probably won't miss it again if I can help it.

Among the more than 600 sessions taking place turing the two days it was once
again extremely hard to decide where to go. As during last year's FOSDEM I made
it into the legal and Go tracks but also visited the lightning talks and Mozilla
ones for a short while.

Contrary to last year's post I don't want to focus on all the talks I've seen,
though, but simply write a bit about the things I've learnt from them or got me
excited 😊

* Thanks to [Anand Babu Periasamy][] I now not only know about his S3 compatible
  server written in Go called [Minio][] but also found out about
  [go-bindata-assetfs][], which wraps go-bindata to be usable directly within
  Go's http server environment. I can't wait to give this a try for writing
  single-binary web applications in combination with things like React 😊

* Right after that, [Jonathan Boulle][] gave an overview about etcd, the
  distributed configuration system developed for CoreOS. It had been on my radar
  for years now but to this date I haven't yet had the change to give it a
  try. Jonathan also mentioned (among others) two tools based on or tightly
  coupled with etcd that I hadn't heard of before: [confd][], a program that
  allows you to manage classic configuration files from etcd, and [vulcand][], a
  proxy/loadbalancer configured through etcd.

* Again on the Go track I listend to [Derek Parker][] about why I should use
  [delve][] for debugging Go code. On OSX getting it to run is a bit more
  complicated but it looks well worth it if you don't want to clutter your code
  full of log statements 😉

* While being in a queue for waffles Ulrich told me about a
  [lightning talk][htoplt] by [Hisham Muhammad][] about [htop][] and it
  appearantly finally getting a proper OSX port 😃

* During a talk about dependency management [Camille Moulin][] dropped the term
  [SPDX][], which is an attempt at standardizing license-related metadata. On
  the periphery I had been aware of that topic and problems related to it but
  that talk motivated me to finally find some time to look into the tooling and
  standards available on this front.

* In my previous team we had started to struggle a bit with our Jenkins
  configuration due to far too many jobs being required for all of our
  projects. For possible alternatives I went to a talk about the
  [Jenkins Job DSL][] by Łukasz Szczęsny and Marcin Zajączkowski. I haven't
  completely understood its advantages and disadvantages yet but I'll most
  likely give it a more detailed look in the future.

* Thanks to [François Marier][fmarier] and [Raegan MacDonald][] I learnt about
  two tools by the EFF and Mozilla respectively that should help in the future
  with some privacy concerns: [Panopticlick][] and [Lightbeam][]

* The last and probably most exciting session of them all was the
  [last keynote on Sunday][] given by [Blake Girardot][] about all the work the
  [Humanitarian OpenStreetMap Team][hotosm] does and how open source helps
  them. If you ever want to see, how much good you can do by contributing to
  opensource, directly or indirectly, watch this talk!

A big THANK YOU to all the organisers and volunteers! It was, once again,
amazing and perhaps we can have some nice weather next year 😉

[blake girardot]: https://hotosm.org/users/blake_girardot
[raegan macdonald]: https://twitter.com/ShmaeganM
[fmarier]: https://fmarier.org/
[jonathan boulle]: https://github.com/jonboulle
[hotosm]: https://hotosm.org/
[minio]: http://minio.io/
[go-bindata-assetfs]: https://github.com/elazarl/go-bindata-assetfs
[last keynote on sunday]: https://fosdem.org/2016/schedule/event/keynote_crisis_response_through_open_mapping/
[confd]: https://github.com/kelseyhightower/confd
[vulcand]: https://github.com/vulcand/vulcand
[lightbeam]: https://www.mozilla.org/en-US/lightbeam/
[panopticlick]: https://panopticlick.eff.org/
[derek parker]: http://derkthedaring.com/
[delve]: https://github.com/derekparker/delve
[jenkins job dsl]: https://github.com/jenkinsci/job-dsl-plugin
[camille moulin]: https://fosdem.org/2016/schedule/speaker/camille_moulin/
[spdx]: http://www.spdx.org/
[htop]: http://hisham.hm/htop/
[anand babu periasamy]: http://www.unlocksmith.org/
[hisham muhammed]: http://hisham.hm/
[travelogue]: http://travelogue.h10n.me/journey/fosdem2016/
[fosdem]: https://fosdem.org/2016/
[htoplt]: https://fosdem.org/2016/schedule/event/htop/
