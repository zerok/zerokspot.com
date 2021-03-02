---
title: 512KB club and FontAwesome
date: "2020-11-29T18:38:31+01:00"
tags:
- zerokspot
- performance
- challenge
---

Yesterday I had a bit of fun trying to make my website here pass the rules for the [512KB club](https://512kb.club/). For those of you who don’t know that, it’s basically a challenge to get the size of your website below 512 KB. I’ve known for a very long time that my use of FontAwesome was a big issue here, coming in at around 1.5MB.

I then took a quick look at what icons I actually needed. I still love that icon set and therefore wanted to keep using it. Turns out that FontAwesome icons always consist of a single path that you can just extract and the build your own little SVG generator around it.

So my solution now consists of two steps:

1. Generate an `icon.js` file based on those icons that I want to have ([source](https://github.com/zerok/zerokspot.com/blob/7a7999b990c88a3c19e7e5abd98bc470ce270b7e/cmd/blog/buildicons.go)). Whenever I want to use a new icon, I just run `blog build-icons` inside the root folder of the project and upload the resulting `icons.js` onto my server.
2. At runtime scan through the whole HTML document, look for `i.icon` elements and then use the paths from the `icons.js` file to generate SVGs inside the document ([source](https://github.com/zerok/zerokspot.com/blob/45a701a2a52dc1f58f3cdddd4d3724a4ec9805d0/assets/js/main.js))

Due to me using FontAwesome Pro here I cannot check in the `icons.js` file, though, but that’s only a minor issue as I had to have the whole FontAwesome folder on my server before, too.

Thanks to that, my front-page now comes it at around 160KB! Mission accomplished 😉
