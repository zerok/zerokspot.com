---
title: "Dark mode on this website"
date: "2019-10-30T19:05:00+02:00"
tags:
- zerokspot
- css
- darkmode
---

Last year I wrote a little bit about incoming support for [detecting
dark mode in CSS][d]. Ever since I had added the prefers-color-scheme
media query to [servethis][s] in late September, I wanted to also play
around with it here.

While I want to eventually provide a button similar to what [Max
Böck][m] has on his website so that you, the reader, can explicitly
select either dark or light mode, this first iteration uses only the
CSS selector.

{{<figure src="/media/2019/darkmode-switch.png" caption="Dark/light switch on <a href=\"https://mxb.dev/\">Max Böck's website</a>">}}

In the next days and weeks I will probably also fine-tune the CSS a
bit as a stumble upon some rough edges. I hope, though, that you enjoy
this "new" look and find it pleasant to read in the dark 🙂

{{<figure src="/media/2019/darkmode-and-lightmode.png" caption="Dark and light mode side-by-side">}}

If not, please let me know!


[d]: https://zerokspot.com/weblog/2018/10/29/dark-mode-on-websites/
[s]: https://github.com/zerok/servethis/commit/3d8964768d11b87ddb32bc5cd15644f5bfd251e8
[m]: https://mxb.dev/
