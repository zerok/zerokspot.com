---
title: "KubeCon/CloudNativeCon Europe 2019"
tags:
- travel
- spain
- barcelona
- kubernetes
- cloudnative
- cncf
date: 2019-05-27T17:34:50+02:00
incoming:
- url: https://chaos.social/@zerok/102168904089105084
- url: https://twitter.com/zerok/status/1133044346796826625
---

Thanks to [Netconomy](https://www.netconomy.net) I had the chance to attend this year's [KubeCon/CloudNativeCon Europe in Barcelona](https://events.linuxfoundation.org/events/kubecon-cloudnativecon-europe-2019/) last week. This was also my first visit to the city so there were lots of new and exciting experience to be had, and neither the city nor the conference disappointed!

{{<figure src="/media/2019/kubecon-eu-booths.jpg" caption="Sponsor showcase while everybody else was at a talk">}}

## Flight cancellations and delays

Sadly, it all started off on the wrong foot. Originally, I was supposed to arrive in the Catalan capital late on Sunday, but my flight from Graz to Frankfurt was cancelled and I got rebooked onto a flight nearly a day later. I also had a ticket for the [Continuous Delivery Summit](https://cdseu19.sched.com/) on Monday but due to that cancellation I couldn't make it. On the plus-side I could spend another night at home 🙂 The two flights on Monday (Graz to Munich and Munich to Barcelona) also got delayed but - at least this time - the connection worked. In the end I arrived in Barcelona about an hour later than expected.

{{<figure src="/media/2019/kubecon-eu-flightcancelled.jpg" caption="My first flight from Graz to Frankfurt was cancelled.">}}

Next: Getting from the airport to my hotel. Following the "preparation is everything" mantra I had purchased a 5-day ticket for the [public transport network](https://www.tmb.cat/en/home) though the TMB app. Turns out, you still have to print the ticket at a vending machine. Finding one at the airport, though, turned out to be harder than expected. Hint: Don't look for them near the actual busses but at the Metro station…

A couple of minutes before midnight I finally reached my hotel and it appeared that there bad been some problem with advance payment… Anyway, about an hour later I was able to get some sleep. Not much, though, as I wanted to collect my conference pass around 08:30 in the morning 😉

But enough of that and on to the actual conference!

## My first KubeCon

KubeCon/CloudNativeCon Europe 2019 was a multi-track event with around 7,000 attendees filling a good chunk of the Barcelona Fira. 

Every day started off with about 90 minutes of keynotes by representatives of the CNCF and high-level sponsors like Microsoft. While that might sounds like something that would be filled with marketing-talk, none of them felt like fillers but were actually really motivating or at least interesting. Probably my personal highlights were a [presentation by two CERN engineers re-running the Higgs-Boson analysis on a 25,000 node cluster](https://sched.co/MRyv) and a [quick introduction to the Operator Framework](https://sched.co/MRyx) by CoreOS/Redhat.

The time between 11:00 and 17:20 was the split into 16 tracks with talks from service meshes to monitoring stacks. There were also a couple of sessions about gRPC. 16 parallel tracks over 3 days produced far too much content and far too many new ideas to write about in a single post, so I will just tease some content and hopefully follow up later.

The biggest topic that was basically part of every talk and every keynote was that Kubernetes should not be seen as a product but **as a platform** to be built upon. Developers should integrate their domain knowledge using the building blocks Kubernetes offers right now and those that will be added in the future. Personally, I’ve already written a handful of controllers and other tools interacting with the Kubernetes API and tools like Redhat’s [Operator Framework](https://github.com/operator-framework) will go a long way of making that easier in the future. Testing is also getting easier with every new [kind](https://kind.sigs.k8s.io/) release!

All of the services, operators, and what-not have to be monitored somehow and it seems like **Prometheus** has won over pretty much the whole community. At least everyone at the event seems to be integrating their projects and products with it. Combined with the merger OpenTracing and OpenCensus into [OpenTelemetry](https://opentelemetry.io/) we as a community now seem to have a broad and unified toolset around observability (metrics, traces, logs).

Another (especially personal) takeaway from the event was that **service meshes can also be light-weight** ([Linkerd](https://linkerd.io/)) and help with debugging service connections. While Linkerd 2.x is currently lacking some of the more prominent features from Linkerd 1.x and Istio, some of the discussions and talks at the conference hinted at things to come in the future. I'm especially curious if there is going to be some first-class integration with the [OpenPolicyAgent](https://www.openpolicyagent.org/) for things like dynamic routing and auth between services.

I will most likely dedicate separate blog posts to some of this topics in the near future 🙂


## Best … conference … party … ever!!!

An event such as KubeCon isn’t all about talks and visiting sponsor booths, though. On Wednesday we also had a party where all attendees were invited. Remember: There were more than 7,000 attendees and now they all should be moved to a location somewhere within Barcelona. To manage that, the organisers had charted dozens of busses in order to get all of us to [Poble Espanyol](https://www.poble-espanyol.com/en/), which they had rented for the evening!

Given that method of transportation, I couldn't resist writing the following toot:

> All #Kubernetes available in Europe is now concentrated in a couple of busses… #kubecon #busfactor

⮑ [@zerok@chaos.social (May 22, 2019, 17:49)](https://chaos.social/@zerok/102140488836934099)

And now to the second aspect of the paragraph: They **rented all of Poble Espanyol for the evening** and placed food stands all over the village. I don't even have a broad idea how much something like that costs but it was probably the best conference event I've ever attended.

{{<figure src="/media/2019/kubecon-eu-event.jpg" caption="The main square of Pople Espanyol was filled to burst with KubeCon attendees.">}}

We were free to explore Poble Espanyol at our leisure and even the craft shops were open for us! Sure, the site probably distracted us from doing networking, but that was a small price to pay 😅

## Lots to watch in post

The CNCF's account on Youtube right now has [around 330 videos from the event](https://www.youtube.com/playlist?list=PLj6h78yzYM2PpmMAnvpvsnR4c27wJePh3) and there are still new talks being uploaded! I'm pretty sure that I'll find some more interesting topics for me while making my way through the recordings 🙂

Anyway, that's it for now about my trip to KubeCon. I really enjoyed
it and I hope this won't be my last journey to a CNCF/Kubernetes event
🙂 It also motivated me to try even more to contribute back to the
community! I also want to migrate some more of my personal projects
over to Kubernetes. Just last week
[DigitalOcean](https://www.digitalocean.com) has GA'd their Kubernetes
offering which looks quite competitive for smaller setups. If I decide
to use that, I'm probably also write a small piece about if and how
I'm happy with it.
