---
date: '2016-08-21T21:14:55+02:00'
language: en
tags:
- development
- documentation
title: Awesome Tutorials
---

Last week I wanted to finally automate some deployment tasks and therefore
started to dig a bit into [SaltStack][]. I've heard tons of good stuff about it
and so I started making my way through its [tutorial][]. I have to say, this is
one of the best introductory documentations I've read in a while! Big thanks to
anyone who has contributed to it over the years.

Topics as complex as server orchestration and provisioning greatly benefits from
offering some way of having a playground to really learn things while doing
them. The approach taken by SaltStack is to use a [Vagrant environment][]
starting up a master and two minion servers. Simply great and I'm now even using
that same environment while creating my first formulas 😊

If you're looking for more, you might also enjoy [Vault][]'s interactive
tutorial which gives you a quick crash course to the secret-store with a
web-based shell!

If you know of more such interactive courses, please let me know!

[vault]: https://www.vaultproject.io/
[saltstack]: https://saltstack.com/
[tutorial]: https://docs.saltstack.com/en/getstarted/
[vagrant environment]: https://github.com/UtahDave/salt-vagrant-demo
