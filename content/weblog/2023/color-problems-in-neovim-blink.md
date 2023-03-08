---
title: Color problems in neovim + Blink.sh
date: "2023-03-08T19:11:16+01:00"
tags:
- mosh
- development
- til
---

While trying to configure a RaspberryPI as a development machine for use with my iPad, I ran into a weird issue: When I connected using [Blink.sh](https://blink.sh) and Mosh, NeoVIM for some reason had completely broken coloring with background colors being extremely intense green/blue instead of ... grey 😑

<figure><img src="https://zerokspot.com/api/photos/2023/03/08/IMG_0019.jpeg?profile=800"><figcaption>Broken colors in neovim</figcaption></figure>

Turns out, truecolor support was only added with [Mosh 1.4](https://mosh.org/mosh-1.4.0-released.html) last October. Ubuntu 22.02 didn’t have that (obviously) and so colors via SSH looked fine but were broken via Mosh. Luckily, there’s a [PPA](https://launchpad.net/~keithw/+archive/ubuntu/mosh-dev) available that makes dev-builds available which fixed my problem!
