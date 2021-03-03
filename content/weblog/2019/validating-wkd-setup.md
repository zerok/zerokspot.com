---
date: 2019-09-28T12:16:48+02:00
title: Validating my WKD setup
tags:
- security
- privacy
- openpgp
---

Right after writing [Signed GPG keys and keyservers][sk], [Wiktor][]
contact me, suggesting that using a [Web Key Directory][wkd] might be
the better approach to solving my signed-keys-exchange problem. He
also noticed that I had uploaded my GPG key in the wrong format, so
receiving it via WKD only worked for GnuPG users.

Because of that I thought about playing a little bit around with how
URLs are calculated for WKD and create a little validation utility:

```
$ wkd-validate 'horst@zerokspot.com'

$ wkd-validate 'horst1@zerokspot.com'
11:45AM FTL Failed to validate WKD content. error="received status code 404 while fetching key via WKD"
```

You can find the implementation [on Github][gh]


[sk]: https://zerokspot.com/weblog/2019/09/12/signed-gpg-keys/
[wkd]: https://zerokspot.com/weblog/2019/03/31/web-key-directory/
[Wiktor]: https://metacode.biz/@wiktor
[gh]: https://github.com/zerok/wkdtools
