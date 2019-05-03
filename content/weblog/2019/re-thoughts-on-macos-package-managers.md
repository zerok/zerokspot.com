---
title: "RE: Thoughts on macOS Package Managers"
date: 2019-05-03T15:09:09+02:00
tags:
- reply
- macos
- sysutils
---

Saagar Jha's ["Thoughts on macOS Package Managers"](https://saagarjha.com/blog/2019/04/26/thoughts-on-macos-package-managers/) goes into quite a lot of detail what is wrong about [Homebrew](https://brew.sh/) from his point of view. There are a few technical issues but also some that a mostly related to how the project currently seems to be treating new contributions and major feature rollouts.

## Analytics...

One prominent example here is the way Google Analytics was enabled in an auto-opt-in fashion. Right now, Homebrew by default submits your package installation operations to Google. There are two options to opt-out, though:

```
$ brew analytics off

$ export HOMEBREW_NO_ANALYTICS=1
```

At least, there is a page in the documentation dedicated to this topic alone: <https://docs.brew.sh/Analytics>.

## /usr/local

On the technical side, Saagar also describes his issues with Homebrew taking over `/usr/local` and making that prefix owned by a normal user account. This is also quite high on my own personal issue-list with Homebrew. Since that prefix is on the default PATH things become ugly pretty quickly when you have more than one user account on your system or simply have other software that wants to use `/usr/local`.

## Options and user communication

Another biggie is the removal of options from core fomulae with Homebrew 2.0. That hit me by removing some options from the Emacs formula which would have been quite useful to me.

That being said, this whole topic opens up something fundamentally complicated with software: How do you get feedback from your actual users on important issues? Most desktop applications show you a little popup with the most important changes when you first launch the shiny new version. Very few CLI apps do that.

Homebrew itself has a built-in way for formulae to show some text after they're being installed but, again, very few actually do that. This is definitely an area where everyone needs to get better. Relying on your users to follow your issue tracker/developer mailing list is simply not a valid option.

## What's next?

Yes, I have some issues with Homebrew and I also share many of those Saagar Jha listed. I do not have an opinion about how users are treated by Homebrew simply because I don't contribute enough. Yes, release notes (and feature deprecations) could be handled better but they are far from being the only project lacking here.

In general, the post included some very interesting points and at least made me look at [MacPorts](https://www.macports.org/) again. I hadn't used it ever since Homebrew was first released (no idea why I switched) but perhaps I will write a follow-up post after I've re-evaluated it more thoroughly.
