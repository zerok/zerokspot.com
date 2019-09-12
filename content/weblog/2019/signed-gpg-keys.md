---
title: "Signed GPG keys and keyservers"
date: 2019-09-12T15:39:38+02:00
tags:
- crypto
- gnupg
- openpgp
- security
---

Recently, a couple of coworkers and I had a little [key signing
party][ksp] in order for everyone on the team to be able to send and
receive credentials and other sensitive data via insecure channels
like Slack or e-mail.

Once we had all our keys generated and started signing them I noticed
something weird. After checking with a collegue I uploaded his key
that I had just signed onto [keys.openpgp.org][koo] so that (1) he and
(2) all the other people in the room could get the newly signed
version easily. There was a problem, though: Turns out, the key that
we all downloaded again from the keyserver didn't contain my
signature. In fact, it didn't contain any signature except for the
original one.

Looking through the FAQ I didn't immediate saw it but [Vincent
Breitmoser][vb], the admin of keys.openpgp.org, quickly pointed me to this
little section in the FAQ detailing their [handling of so-called
"third-party signatures"][k3]:

>  Do you distribute "third party signatures"?
>
> Short answer: No.
> 
> A "third party signature" is a signature on a key that was made by
> some other key. Most commonly, those are the signatures produced
> when "signing someone's key", which are the basis for the "Web of
> Trust". For a number of reasons, those signatures are not currently
> distributed via keys.openpgp.org.
> 
> The killer reason is spam. Third party signatures allow attaching
> arbitrary data to anyone's key, and nothing stops a malicious user
> from attaching so many megabytes of bloat to a key that it becomes
> practically unusable. Even worse, they could attach offensive or
> illegal content.
> 
> There are ideas to resolve this issue. For example, signatures could
> be distributed with the signer, rather than the
> signee. Alternatively, we could require cross-signing by the signee
> before distribution to support a caff-style workflow. If there is
> enough interest, we are open to working with other OpenPGP projects
> on a solution.

How is another popular keyserver, [keyserver.ubuntu.com][kuc],
handling that?  To try that out, I exported my own signed key from my
keychain, uploaded it there, and inspected what the server returns:

```
$ curl --silent https://keyserver.ubuntu.com/pks/lookup\?op\=get\&search\=0xc4e02ba840483a5d6b7616076f203f0d220f8e98 | gpg -v                                                                                                                                                                [k:taa-master.conf // minikube]
gpg: WARNING: no command supplied.  Trying to guess what you mean ...
pub   rsa4096 2017-01-15 [SC]
      C4E02BA840483A5D6B7616076F203F0D220F8E98
uid           Horst Gutmann <horst@zerokspot.com>
sig        6F203F0D220F8E98 2017-01-15   [selfsig]
sig        6F203F0D220F8E98 2017-01-15   [selfsig]
sig        6F203F0D220F8E98 2018-12-20   [selfsig]
sig        7CAF0EEF0B8DBFB6 2019-07-24   Robert van der Stel <robert@vanderstel.at>
sig        E8620BE896A1864F 2019-09-05   Maximilian Zollneritsch <m.zollneritsch@netconomy.net>
uid           Horst Gutmann <horst.gutmann@gmail.com>
sig        6F203F0D220F8E98 2017-01-15   [selfsig]
sig        6F203F0D220F8E98 2018-12-20   [selfsig]
sig        7CAF0EEF0B8DBFB6 2019-07-24   Robert van der Stel <robert@vanderstel.at>
sig        E8620BE896A1864F 2019-09-05   Maximilian Zollneritsch <m.zollneritsch@netconomy.net>
uid           Horst Gutmann <zerok@zerokspot.com>
sig        6F203F0D220F8E98 2017-01-15   [selfsig]
sig        6F203F0D220F8E98 2018-12-20   [selfsig]
sig        7CAF0EEF0B8DBFB6 2019-07-24   Robert van der Stel <robert@vanderstel.at>
sig        E8620BE896A1864F 2019-09-05   Maximilian Zollneritsch <m.zollneritsch@netconomy.net>
uid           Horst Gutmann <h.gutmann@netconomy.net>
sig        6F203F0D220F8E98 2017-01-15   [selfsig]
sig        6F203F0D220F8E98 2018-12-20   [selfsig]
sig        7CAF0EEF0B8DBFB6 2019-07-24   Robert van der Stel <robert@vanderstel.at>
sig        E8620BE896A1864F 2019-09-05   Maximilian Zollneritsch <m.zollneritsch@netconomy.net>
sub   rsa4096 2018-06-17 [A]
sig        6F203F0D220F8E98 2018-06-17   [keybind]
sub   rsa4096 2017-01-15 [E]
sig        6F203F0D220F8E98 2017-01-15   [keybind]
sig        6F203F0D220F8E98 2017-01-15   [selfsig]
sig        6F203F0D220F8E98 2018-12-20   [keybind]
sig        6F203F0D220F8E98 2018-12-20   [keybind]
```

OK, so that keyserver allows third-party signatures. What about
[pgp.mit.edu][pme]? Despite running into lots of timeouts and proxy
errors I eventually managed to run the same test there with a similar
result. The MIT server also returns third-party signatures.

So while the MIT and the Ubuntu servers "work", I can absolutely get
behind the reasoning behind openpgp.org's decision not to support that
feature. So for now I will simply request from anyone signing my
key that they send me the signed key instead of uploading it
somewhere.

That being said, I'm looking forward to a future where the spam
problem has been fixed somehow...



[k3]: https://keys.openpgp.org/about/faq#third-party-signatures
[vb]: https://twitter.com/Valodim
[koo]: https://keys.openpgp.org/
[ksp]: https://en.wikipedia.org/wiki/Key_signing_party
[kuc]: https://keyserver.ubuntu.com/
[pme]: https://pgp.mit.edu/
