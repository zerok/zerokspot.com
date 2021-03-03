---
title: "npm outage due to Cloudflare"
date: "2020-02-18T19:20:00+0100"
tags:
- javascript
- news
---

Yesterday at 13:26 CET  a coworker contacted me asking what this error was all about:

> npm ERR! 429 Too Many Requests - GET https://registry.npmjs.org/yourpackagename.tgz

He presumably got it either on his local development machine or on one of our CI servers. This kind of error message smelt like something that might be related to some DDoS prevention system/WAF (Web Application Firewall) and so I forwarded him to our internal expert on the topic.

A short while later [a ticket](https://github.com/npm/cli/issues/836) appeared in npm's issue tracker and lots of users were flooding Twitter with reports of that same error I had received from my co-worker. Quickly it became clear that there was something weird going with Cloudflare which was eventually confirmed by David Kitchen:

> Hello and profuse apologies from Cloudflare, a post-mortem of sorts directly in your issue comments.
>
> I am the engineering manager for the DDoS protection team and this morning at 11:06 UTC we tweaked a rule that affected one of our signals. The signal relates to the HTTP referer header, and we have a piece of code that looks at invalid referer headers. In this case we tweaked it to include not just "obvious garbage" but "anything that does not conform to the HTTP specification"... i.e. is the referer a URI? If not then it contributes to knowledge about bad traffic.
>
> So... why did this impact npmjs.org? It turns out that a lot of NPM traffic sends the referer as "install" which is invalid according to the HTTP specification. As NPM is also a heavily trafficked site this resulted in the DDoS systems picking this up and treating the traffic as a HTTP flood and determining that a rate-limit should be applied.
>
> When we noticed that NPM was seeing an increase in HTTP 429s (as seen on Twitter) we contacted NPM and started an internal investigation. As soon as we identified the root cause we reverted the change, which was at 13:00 UTC.
>
> We'll note that NPM and 1 other site use the referer for purposes outside the HTTP spec and we'll update our systems to ensure that this does not happen again. Additionally we'll improve our monitoring around changes of this nature so that we can discover impact sooner and roll back automatically.
>
> -- [buro9](https://github.com/npm/cli/issues/836#issuecomment-587019096)

So, it was npm's fault based on the HTTP Referer header? "install" is not a valid value there? Looking at [RFC 2616 section 14.36](https://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.36) it's not that black or white:

1. Referer can contain either absolute or relative URLs. "install" would be a relative URL and therefore syntactically valid.
2. **But** a source that does not have its own URI **must not** send a Referer header.

So, in a sense, Cloudflare is right that npm is not setting a proper referrer here but taking that as an indicator for malicious content is still ... interesting.

The way that Cloudflare communicated this new rule and its affects on the npm infrastructure was also problematic at best which was nicely summarized by npm's CTO:

> @buro9 we would appreciate it if you respond to our tickets and our internal slack comms, before posting to a public issue, we still have not gotten a post-mortem report for the last two outages.
> 
> as for pointing at HTTP specifications, considering this behaviour has been in place for years, I would ask to review what change was pushed in CF today that caused this sudden "compliance with HTTP Specification" result?
>
> I'll ask again to please follow up with our open tickets and report back to to us on the post-mortem for the last two outages, we'd rather learn about this from you directly, than seeing it in an issue on github ..
>
> -- [Ahmad Nassri](https://github.com/npm/cli/issues/836#issuecomment-587040922)

As of yesterday 13:07 UTC [that incident](https://status.npmjs.org/incidents/5qxhhcx0vtlc) has been resolved.
