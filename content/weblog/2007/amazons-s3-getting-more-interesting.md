---
date: '2007-10-09T12:00:00-00:00'
language: en
tags:
- amazon
- s3
- sla
title: Amazon's S3 getting more interesting
---


... and now according to [Jeff Barr](http://aws.typepad.com/aws/2007/10/amazon-s3-at-yo.html) even an SLA, which should be something for all those people who thought about using S3 for mission critical storage but were turned down by the lack of some kind of clear statement regarding the reliability of the service.

-------------------------------

The new [SLA](http://www.amazon.com/gp/browse.html?node=379654011) includes in my opinion exactly such a statement with a 10% credit should the monthly uptime percentage fall below 99.9% (and even 25% for less than 99% uptime). This agreement won't be really all that interesting to private customers, though, since it requires that you yourself protocol when according to you S3 had an outage and notify Amazon about this issue. 

But S3 thanks to its architecture is more suited for professional in the first place, anyway. 

[via [Fred Oliveira](http://blog.webreakstuff.com/2007/10/amazon-s3-gets-a-sla-exhale/)]