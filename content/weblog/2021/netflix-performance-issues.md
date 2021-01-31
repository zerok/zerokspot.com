---
title: Netflix performance issues
date: "2021-01-31T12:48:00+01:00"
tags:
- 100daystooffload
- netflix
- streaming
- qos
- customersupport
incoming:
- url: https://chaos.social/@zerok/105650192622701318
---

Content-wise I really love Netflix. At least compared to Disney+, AppleTV+, and even Amazon Prime I’ve had a lot of fun with the exclusive content provided there. Sadly, not everything is that great with them, though.

For the last months we had some weird issues with the image quality. IIRC it started before we switched ISPs from UPC/Magenta Austria to A1 with lots of hiccups and buffering phases while streaming. Once we switched to A1 everything was fine again but in December something else happened: On many (if not most) evenings after about 18:00 or 19:00 the streaming quality completely degrades. While we have a 300Mbps downstream connection (also benchmarked during that time-window) and a Premium UltraHD subscription with Netflix the streaming quality goes down to something that looks more like a 1996 FMV sequence in Command & Conquer (meaning: really bad, lots of compression fragments,…). This happens while only a single stream is active with pretty much no other traffic going on (ignoring normal web-browsing and IMAP here).

Frustrated, we contacted our ISP and they told us that they don’t do any QoS-levelling on their end and from what they can see our downlink is fine but that Netflix might be throttling on their end. So, next we contacted Netflix and they too said that they weren’t throttling anything.

OK, Netflix didn’t actually say that since all I got was writing back and forth with first-level support personell who insisted that it had to have something to do with my home network or with my ISP. Since this issue is only limited to Netflix and not to Amazon Prime, Disney+, or AppleTV+ that sounds quite unlikely but they ignored that despite me trying to make that point repeatedly… I don’t know how many guides I’ve already received from Netflix support on how to restart my devices 😣

At this point I’m not sure what to do. Perhaps we will just move down one tier to HD streaming again but then we would limit ourselves to only two parallel streams and two devices for downloads. Given that we are 3 people in our household, each with different preferences, that might be rather painful. It’s also really sad since I just like their content and the overall app quality.

If nothing else, this issue has shown me, that I just need more monitoring in my network so that I can at least quickly check what devices are currently producing how much traffic. But that’s a project for another weekend 😅
