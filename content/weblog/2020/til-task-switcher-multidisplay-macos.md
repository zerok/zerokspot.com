---
title: 'TIL: Task-switcher on the wrong display (macOS)'
date: "2020-09-08T18:22:19+02:00"
tags:
- macos
- til
---

One thing that has annoyed me for many years now with the multi-display support in macOS is that the task switcher (the one which gets triggered when hitting cmd+tab) is not tied to the primary display (the one which has the menu-bar visible in the System Preferences / Displays / Arrangement dialog).

<figure><img src="/media/2020/Screenshot%202020-09-08%20at%2018.11.06.png"><figcaption>System Preferences / Displays is not the place you're looking for</figcaption></figure>

In my case, the dock is hidden by default. If I now move to the bottom of the secondary screen by accident, the dock will appear there and all of a sudden the task switcher is also bound to that screen. If I then do the same on the primary screen, the dock will re-appear there and so will the task switcher.

So, to summarise: The task switcher is bound to the display that shows the dock. You can move the dock by moving the mouse cursor to the very bottom (or whatever edge you use to show your dock) in order to move the dock to that screen. Thinking about it, it makes sense in hindsight. Still….
