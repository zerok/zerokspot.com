---
title: "bradfitz/jndi: an irresponsibly bad logging library"
likeOf: "https://github.com/bradfitz/jndi"
date: 2021-12-11T13:06:08+0100
tags:
- golang
- log4j
- jndi
- CVE-2021-44228
---
For those of you who didn't have [the joy of fixing log4j issues](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2021-44228) and feel left out, this should be something for you (if you're using Go): Brad Fitzpatrick released a [little logger](https://github.com/bradfitz/jndi) that also resolves JNDI URIs and sends out HTTP requests where not appropriate. Too soon, Brad, too soon 😅