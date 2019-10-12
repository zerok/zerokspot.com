---
title: "Thunderbird 78 without Enigmail but with OpenPGP!"
date: 2019-10-12T09:30:31+02:00
tags:
- security
- email
- openpgp
- gnupg
- thunderbird
---

A couple of days ago, the [Thunderbird developers announced][tb] that
they want to support OpenPGP as part of Thunderbird itself in the
future. At the same time, [Patrick Brunschwig][pb], the developer
behind the popular Enigmail extension, wrote that Enigmail will not
supported for Thunderbird beyond version 68 due to changes to the
extension API. He will still maintain Enigmail for Postbox, though.


The proposed OpenPGP integration for Thunderbird will be one without
relying on GnuPG, though. This will have the consequence that it won't
be possible to share your keyring between GnuPG (and other mail
clients using that like Mailmate). Because of that, since I'm using
OpenPGP not only for encrypting e-mails, I will have to have two
keyrings that I need to keep in sync.

This custom implementation will have other effects as well:

> Will OpenPGP cards be supported for private key storage ?
> 
> Probably not, because we don't use the GnuPG software that's usually
> required to access OpenPGP smartcards.
>
> -- [wiki.mozilla.org][wm]

Right now I have my primary key on a YubiKey and always with me. Due
to that change I will most likely have to move to a sub-key of that
"master key" that is then available only on the machine running
Thunderbird. I'll have to think about how I could make such a setup
work for me, but I at least have another 12 months time for that 😉

For onboarding new users these are fantastic news, though! They won't
have to install two additional pieces of software (Enigmail and GnuPG)
after installing Thunderbird and debugging possible error will be much
easier thanks to just a single implementation.


[wm]: https://wiki.mozilla.org/Thunderbird:OpenPGP:2020
[tb]: https://blog.mozilla.org/thunderbird/2019/10/thunderbird-enigmail-and-openpgp/
[pb]: https://admin.hostpoint.ch/pipermail/enigmail-users_enigmail.net/2019-October/005493.html
