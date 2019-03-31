---
title: "GnuPG: What is a Web Key Directory"
date: "2019-03-31T11:21:32+02:00"
tags:
- gnupg
- openpgp
- til
---

A couple of weeks ago I was looking for a way to manage multiple if
not all SSH keys within an organization. My primary requirement was a
simple way to make all these keys retrievable to everyone inside the
organisation. At first I thought about hosting a key server for that
organization but then
[Wiktor](https://mastodon.social/@wiktor/101698016892305192) on
Mastodon mentioned something called “Web Key Directory” which I hadn’t
heard about before so I thought I should give it a look 🙂

Web Key Directory or WKD for short is an alternative GnuPG
distribution/publishing implemented by GnuPG 2.1.12. Compared to a
tradition key-server it allows key owners to publish their keys on
their own domain and users to retrieve them from there without having
to look them up on some third party.

It does that by exposing public keys through a well-known URL
schema. Let's take the email "zerok@zerokspot.com" as example. A
client that supports WKD would now look for a public key to that
e-mail address on the following URL:

```
https://zerokspot.com/.well-known/openpgpkey/hu/<name-hash>

name-hash := zbase32(sha1("zerok"))
```

You can also get the name-hash by running the `--list-keys` command
with the additional `--with-wkd` option:

```
$ gpg --list-keys --with-wkd zerok@zerokspot.com

...

uid           [ultimate] Horst Gutmann <zerok@zerokspot.com>
              mgaqw3gxp6grwymorjyk99yanaxk1kbu@zerokspot.com

...
```

All that's left to do is to export the public key to a file using the
naming scheme defined above and upload it it:

```
$ gpg --export KEY_ID > mgaqw3gxp6grwymorjyk99yanaxk1kbu
$ scp mgaqw3gxp6grwymorjyk99yanaxk1kbu \
    server:/var/www/htdocs/.well-known/openpgpgkey/hu/
```

You can find further details about WKD on the [official GnuPG
wiki](https://wiki.gnupg.org/WKD) and even more so on [Matt Rude's
website](https://keyserver.mattrude.com/guides/web-key-directory/). There
is even a WKD checker available on
[metacode.biz](https://metacode.biz/openpgp/web-key-directory) :-D
