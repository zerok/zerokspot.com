---
title: "GopherCon 2018: A quick recap"
date: 2018-09-17T19:59:01+02:00
tags:
- gophercon
- traveling
- denver
- usa
---

Thanks to my employer, [Netconomy](https://netconomy.net), a colleague and I
could attend this year's edition of [GopherCon](https://www.gophercon.com/).
I've wanted to go there for years and now I finally had a chance to actually
get to Denver for the biggest gathering of Gophers in the world 😄 So on
Sunday, August 26, we boarded the plane taking us to Frankfurt, switch machine,
and then spent another 9.5 hours in a Boing 747-400 to Denver.

<figure>
<img src="/media/2018/gophercon.jpg">
<figcaption>1Password also had a large booth at GopherCon!</figcaption>
</figure>

The usual chicken-vs-paste meals later we arrived rather exhausted at Denver
International Airport. Thanks to some construction-work going on there it's
probably the most complicated airport I've seen in recent history. It took us
quite some time and a couple of calls with our Lyft-driver to find the right
exit, but we eventually found the correct parking deck 🤪 Off we went to the
"Mile-high city"! 

Good thing we hadn't booked any tutorials for the next day as we wanted to get
to know the city and also getting some work done before focusing completely on
the main event!

Some sight-seeing, Whole Foods, and working later we arrived at GopherCon on
Tuesday morning, nearly awake and fully-motivated! By the way, putting a gift
store right next to registration booth where you could purchase Gopher swag was
just devious 😈

## The talks

For half of the session talks there was just a single track but around lunch
there were two talks/tutorials where three sessions were taking place in
parallel. These had a length of 45 minutes while the single-track sessions took
only 25 minutes. I really liked the format as it allowed to put highly
domain-specific talks onto the schedule without forcing people into them who
weren't interested. In one such slot, for instance, Filippo Valsorda presented
[best-practices around the net
package](https://www.youtube.com/watch?v=afSiVelXDTQ) while people in another
slot learnt about 3D rendering thanks to [Hunter
Loftis](https://twitter.com/hunterloftis?lang=en).

It's hard to pick favourites here but if I had to pick, I'd give the price to
Kavya Joshi for her talk about the Go scheduler (["The Scheduler
Saga"](https://www.youtube.com/watch?v=YHRO5WQGh0k)) and Anthony Starks for
["Go for Information Displays"](https://www.youtube.com/watch?v=NyDNJnioWhI).
Thanks to the parallel tracks there are still a couple of sessions, though,
like Jon Bodner's ["Go Says WAT?"](https://www.youtube.com/watch?v=zPd0Cxzsslk)
that I have to watch on Youtube.

Some take-aways from these and other talks for me were:

* Anthony Starks has created a neat little presentation tool called
  [deck](https://github.com/ajstarks/deck).
* He also created a Go-binding for
  [OpenVG](https://github.com/ajstarks/openvg), which can also be used with
  Deck to render the slides directly onto the screen without having to go
  through PNGs et al.
* [gvt](https://github.com/FiloSottile/gvt) is officially retired in favour of
  Go modules.
* If you want to debug binary data, use "%q" as it prints them in a more usable
  way than "%v".
* There is a [Go implementation of
  yacc](https://godoc.org/golang.org/x/tools/cmd/goyacc)!
* Tons of Debian tools and services are already written in Go, such as
  [manpages.debian.org](https://manpages.debian.org). 

## Social life

After the first session-day there was also a social event taking place at the
nearby [Sculpture Park](http://www.artscomplex.com/?TabId=142). I have no idea
why it is called that given we only saw a single sculpture, but never mind.

On Wednesday evening there was another event sponsored by DigitalOcean at ["The
1up"](http://www.the-1up.com/), a local barcade. We had a table reserved at
Fogo de Chao right before that event and ate so much that we couldn't stay for
long. Great location, though, with a great bar and awesome arcade machines!

As Max had forgotten his wristband which acted as invitation to the event we
also had lots of fun with the bouncer esp. as he was trying to pronounce our
names 🙂

## Community day

The two session days went by far too quickly! Luckily, though, that wasn't the
end of the event as there was still a day full of community projects left!
Sadly, I had to fix a couple of problems at home and therefore mostly spent my
time away from Go projects. That being said, Max and I at least started working
on a little pet-project of ours which should make creating release-notes based
on changed between releases easer in the future 😅

A single community day was just a bit too short when you're used to sprints at
Python conferences. There were also fewer people actively looking for help with
their projects, which left the announcement whiteboards rather empty. If I get
the chance again to attend next year then I hope I can bring something to the
table myself. This year, it was mostly [Ansible](https://www.ansible.com/),
though. As I still had fun I couldn't complain 😉

## The venue

Something else I couldn't complain about was the venue. As with the Oregon
Convention Center the [Colorado Convention
Center](https://en.wikipedia.org/wiki/Colorado_Convention_Center) was huge,
offering space for multiple large events taking place at the same time without
people running over each other. There were always refreshments available and
the queues to the lunch booths were manageable. If only there had been porridge
also on the community day for breakfast 😫

In general, the organisation was flawless (from what I can tell)! The breaks
were well-timed, there was captioning, and the air-condition was not set to
kill (only to hurt 😅) so we didn't have to leave the conference halls to warm
up as happened on another event 😉 Big thanks to everyone involved with the
event! Great job!

As I already wrote before, I really want to attend next year's GopherCon in San
Diego. I enjoyed my time in Denver so much that it was hard to leave. Next year
there is also going to be a dotGo again, so with some luck, I will be attending
two Go conferences + FOSDEM again 😍


