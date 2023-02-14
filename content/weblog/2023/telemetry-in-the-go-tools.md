---
title: Telemetry in the Go tools
date: "2023-02-14T17:13:32+01:00"
tags:
- golang
- telemetry
- o11y
---

Over the last weekend quite a few people of the Go community made a lot of noise about [a proposal that Russ Cox had made about adding telemetry](https://research.swtch.com/telemetry) to the core Go tools (like the compiler, gopls, gofmt, etc.). That proposal comes not only as a ticket on GitHub but also with three blog posts explaining the overall idea, to some proposed implementation details, all the way to some example use-cases.

One thing to keep in mind about all of this: This is just a **proposal**. It hasn’t been implemented yet and is current just in discussion. On a personal note, this is just a snapshot of my opinion on this proposal. It has changed it the past, it might do so again in the future. So what is this proposal all about?

In general terms (assuming that I understood the proposal correctly), the Go team would like to know which features developers are using in the language and the toolchain and which are not used anymore in order to prioritize development efforts and also detect some issues without requiring developers to explicitly open a ticket.

To answer those questions, metric points could be added to the Go tools. When the Go team wants to know, if people are using certain parts of the compiler, they could add a configuration in a central place and the Go tools would fetch that config in order to start preparing an answer in an aggregated report spanning a week (so one single number per metric per week). Now, based on the actual number of installations, yours may or may not also be triggered to submit that report at the end of the week. This depends on how many reports the Go team wants to have in order to get a statistically significant answer with a defined error margin. This also means that it would be rare that your installation would actually make a submission.

When I first read about this proposal, I was shocked but decided to read through those posts before coming up with an actual opinion. I’m now at a point where I actually agree with pretty much everything in there (with one big exception). The proposal seems to try really hard to prevent abuse of the system and to collect really only the bare minimum of what would be needed to answer some very high-level questions without including any actual identifiers and with working only with aggregates. Not every installation would send such telemetry but only more or less randomly chosen ones and only as many as are needed to provide some statically significant data in a defined error margin. 

I’m even at the point where I say that *I* as a normal developer using Go want the Go team have that kind of data. But now to the one thing about this proposal that I don’t like: It’s opt-out.

If I understood the proposal correctly, every installation would have a couple of weeks before even considered for collection; so if the developers know about this feature, they have ample time to switch it off by using a documented environment variable. That’s in my opinion a good way to deal with an opt-out system, but it’s still opt-out and requires a lot of trust. As Go developers we already trust the Go team (and Google for that matter) a great deal by (1) using the language and tools and (2) going through things like the sumdb and the module proxy. With this telemetry approach, the Go team opens up another path *by default* that requires trust at pretty much the same level as (2).

Unfortunately, it’s pretty hard to split the Go team from Google. If, let’s say, the Rust team had made such announcements, the outcry might have been a completely different one but here, with Google involved, it just has a strong and bad aftertaste. As I said above, I like the overall idea here. I believe, though, that it should be opt-in at this point in time and may or may not be ok to become opt-out eventually if absolutely necessary and the approach has earned the community’s trust.
