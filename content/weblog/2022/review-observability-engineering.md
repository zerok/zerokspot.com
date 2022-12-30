---
title: 'Review: Observability Engineering'
date: "2022-12-30T12:51:27+01:00"
tags:
- books
- reviews
- observability
- monitoring
---

A couple of days ago I decided to give O’Reilly’s subscription offering a try. While I eventually decided not to extend the trial period into a whole year, I wanted to make the best out of those 10 days and read [“Observability Engineering” by Charity Majors, Liz Fong-Jones, and George Miranda](https://learning.oreilly.com/library/view/observability-engineering/9781492076438/).

The authors manage to give a great overview about the whole debate around what’s the big difference between observability and monitoring. While I was at least partially aware of the differences, I didn’t know that the focus is so completely different, with it lying on individual users instead of the overall system.

They also succeed in making the necessity of observability very clear. Preparing metric collectors for all those things that *could* go wrong in a modern, distributed system is no longer feasible. Instead, the system should allow an observer to explore its behaviour and drill down through events and attributes to debug issues. 

There is also no longer a single user-experience. Instead, the system might work perfectly for some served by one server while for others it might seem completely broken. Also here, events/traces focus more on specific user-journeys instead of trying to put the state of the whole system into very abstract and aggregated numeric values (metrics).

But the book doesn’t stop at this high level but also includes a quick introduction to [OpenTelemetry](https://opentelemetry.io/) and guest chapters by engineers at Slack about how they make their systems observable and streamline the process using an “observability pipeline”. This and the chapter about various sampling strategies are perhaps my absolute highlights here as they answer some the question marks I’ve had flying around in my head for a long time!

For a second edition I’d like to have a bit less repetition, though. Especially the chapters about how to introduce observability at a company and the “Observability Maturity Model” mostly reiterate again and again on the primary motivators already presented in the first couple of chapters. Personally, I’d also have liked to see some more guest chapters from different companies that are not using Honeycomb. While it at least seems that the authors really tried to not focus on Honeycomb products, they only partially succeeded.

Except for these few points I really enjoyed the book and can recommend it to anyone working as software engineer.
