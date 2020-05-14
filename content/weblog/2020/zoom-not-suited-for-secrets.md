---
title: "Zoom: Not suited for secrets"
date: "2020-04-03T13:23:11+0200"
tags:
- zoom
- privacy
- encryption
- security
- china
---

This week, US-company Zoom has come under fire after [The Intercept published a story][I] indicating that Zoom's claim of End-to-End encryption is mostly bogus:

> But when reached for comment about whether video meetings are actually end-to-end encrypted, a Zoom spokesperson wrote, “Currently, it is not possible to enable E2E encryption for Zoom video meetings. Zoom video meetings use a combination of TCP and UDP. TCP connections are made using TLS and UDP connections are encrypted with AES using a key negotiated over a TLS connection.”

Last night, Citizen Lab at University of Toronto has followed up with [publishing a report][c] on their own after taking a long, hard look at Zoom. According to them, audio and video content in Zoom meetings ist only encrypted with **a single AES-128 key operating in ECB mode**. The key itself is sent to the participants *from* the central Zoom servers and even with no participant being in China, **Chinese servers** seem to be used here. The Citizen Lab's report closes with this chapter title:

> 5. Conclusion: Not Suited for Secrets

Honestly, at this point I think everyone should just throw some money and humanpower at open projects like [Jitsi][j] and try to minimise their (legal) losses... 

[I]: https://theintercept.com/2020/03/31/zoom-meeting-encryption/
[c]: https://citizenlab.ca/2020/04/move-fast-roll-your-own-crypto-a-quick-look-at-the-confidentiality-of-zoom-meetings/
[j]: https://jitsi.org/
