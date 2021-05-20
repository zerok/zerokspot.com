---
title: What’s new in tpl 4.4.0
date: "2021-05-20T19:35:22+02:00"
tags:
- 100daystooffload
- opensource
---

Today, [Maximilian updated lots of the dependencies](https://github.com/zerok/tpl/pull/24) we have in [tpl](https://github.com/zerok/tpl) and I then pushed those updates out as [v4.4.0](https://github.com/zerok/tpl/releases/tag/v4.4.0). So what’s in the package this time around?

Thanks to upgrading [sprig](https://github.com/Masterminds/sprig) from [2.16 to 3.2.2](https://github.com/Masterminds/sprig/releases) it got tons of new template functions and functionality:

- [durationRound](https://github.com/Masterminds/sprig/pull/187)
- [toRawJson](https://github.com/Masterminds/sprig/pull/193)
- [get support for dicts](https://github.com/Masterminds/sprig/pull/197)
- [seq](https://github.com/Masterminds/sprig/pull/205)
- [duration filter](https://github.com/Masterminds/sprig/pull/224)
- [htpasswd](https://github.com/Masterminds/sprig/pull/225)
- [randInt](https://github.com/Masterminds/sprig/pull/211)
- [fromJson and mustFromJson](https://github.com/Masterminds/sprig/pull/223)
- [bcrypt](https://github.com/Masterminds/sprig/pull/242)
- [randBytes](https://github.com/Masterminds/sprig/pull/253)
- [dig function for dicts](https://github.com/Masterminds/sprig/pull/254)
- [regexQuoteMeta](https://github.com/Masterminds/sprig/pull/257)
- [osBase, osDir, osExt, osClean, and osIsAbs](https://github.com/Masterminds/sprig/pull/261)
- [addf, add1f, subf, divf, mulf, maxf, and minf](https://github.com/Masterminds/sprig/pull/181)
- [chunk](https://github.com/Masterminds/sprig/pull/265)
- [various certificate functions](https://github.com/Masterminds/sprig/pull/270)

Actually, all the changes you should see are coming from that one upgrade. Maximilian also updated pretty much every dependency but these are mostly under-the-hood.

We should really keep those dependencies updated at a more regular interval. There’s some awesome stuff in sprig v3 that I completely missed! Big thanks to Maximilian for the update 😀
