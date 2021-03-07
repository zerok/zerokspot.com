---
title: First steps at getting webmentiond listed
date: "2020-07-04T09:12:54+02:00"
tags:
- webmentiond
- web
- community
- 100daystooffload
---

Thursday was quite a bad day for me. For some reason I was completely down and nothing was enjoyable. Perhaps I simply had too much to eat the evening before when we were at a BBQ event with the family. I was also deep in preparing for an exam for next week and therefore didn’t get any immediate successes that usually lift my spirit. Whatever the reason, I knew that I should get some coding in and so I finally sat down (after having planned to do so for months now) to go through the [implementation report template](https://github.com/w3c/webmention/blob/master/implementation-reports/TEMPLATE.md) for getting [webmentiond](https://webmentiond.org) listed on [webmention.net/implementations](https://webmention.net/implementations).

When I first started webmentiond, I just wanted to learn about the webmention protocol and simply have some fun at it. It quickly grew into a system that I’d use for my own websites and projects and eventually also became something that I saw myself maintaining in the long run also for a broader user base than just myself 😄

From a user’s point of view, the implementation listing provided by the webmention.net has the advantage that they can quickly determine what features are or are not supported by a given implementation. The implementation report template also requests information regarding a set of test-cases and how the implementation handles these. 

While webmentiond already had an extensive test suite, more cases are always welcome and in this scenario also helped me uncover a couple of bugs in my implementation. After a bit of back and forth between running one of the test-cases and fixing unexpected behaviour I was able to submit [my implementation report](https://github.com/w3c/webmention/pull/105) on the day day. In that process I also created lots of new tests including [a suite that was designed to emulate some parts of the official set](https://github.com/zerok/webmentiond/blob/3e2fc6a0d73596e7eab606d28bfac85d217c78c8/pkg/server/conformance_test.go) 🥳

I now also finally had my success for the day and was up again 😊 Sure, filling out the report was just a first step but IMO an important one. Now, I’m looking forward to getting some feedback on it and hopefully get it merged eventually. 
