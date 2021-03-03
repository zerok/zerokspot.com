---
title: Finding a good SSH client for iOS
date: "2020-07-27T18:20:27+02:00"
tags:
- 100daystooffload
- development
- ios
---

Since my laptop is still at the repair shop I had a bit of time and need to test some SSH clients for iOS for daily use. I started off with Prompt, which I had purchased a couple of years ago simply because I just wanted an SSH client without any bells and whistles and without any kind of in-app purchase in order to make it work properly. 

With daily use, though, Prompt turned out to be less reliable than I had hoped, kicking me back to the connect screen after just having been in the background for a minute or so while I was looking up some code elsewhere. Looking a bit around I found [Blink Shell](https://blink.sh) which, so far, has been wonderful to work with. As a bonus it also supports Mosh which should help with the majority of flaky connections. It’s also highly customizable and supports multiple windows on iPad OS which allows you to have multiple SSH clients open in parallel, something that Prompt doesn’t offer at all as far as I know.

<figure><img src="/media/2020/0D1E22FA-1385-4FA8-9E5C-DD1E2919E006.png"><figcaption>Split window with empty shell on the left and help dialog on the right</figcaption></figure>

The one thing that genuinely surprised when when I first launched the app, that really *everything* is trigged from a commandline interface. If you want to access the settings, you have to enter “config” 😅 Apart from that, it’s extremely intuitive and also the handling of multiple tabs and windows with cmd+o being used to jump from one window to the next and cmd+shift+right to do the same but for tabs required only a very short time before entering my muscle memory.

Oh, and Blink Shell is under GPL which completely sealed the deal for me. It’s not free. In fact, it’s € 21.99, but you get quite a lot of value here AND you support free software 🥰

OK, now that the client-side is sorted out, I should probably order yet another RaspberryPI4 in order to play around with Ubuntu Server/Core on ARM64 😂
