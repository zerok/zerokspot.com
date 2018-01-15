---
title: "Yubikey auto-lock, the duct-tape way"
date: 2018-01-15T22:01:48+01:00
tags:
- security
- yubikey
- macos
- python
---

Ever since I registered my [YubiKey](https://www.yubico.com/) with my Mac, I wanted to have one feature: Whenever I unplug it at work, I want the device to be locked. Therefore, I looked at implementations like [yubiswitch](https://github.com/pallotron/yubiswitch) but for one reason or another I decided to solve this little problem by myself.

At first I wanted to give Swift a try and found some functionality inside IOKit that might help me detect when a certain USB device got disconnected. Sadly, I don't have enough time do learn Swift + Foundation + IOKit right now. So I've decided to go with the duct tape approach at this point: 

* A simple shell script that checks the connected USB devices for one named `Yubikey` and if it first found one and all of a sudden not anymore, it triggers macOS to launch the screensaver.
* And a LaunchAgent to launch that script once the user logs in.

You can find the current implementation [here](https://gist.github.com/zerok/96ff78adbbf2105dd116b20aadc07718). I really don't like this approach but I simply don't have the time for something involving OS-level notifications right now. I still want to do it the proper way, though, some time in the future 🙂

