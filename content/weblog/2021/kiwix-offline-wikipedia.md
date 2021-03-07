---
title: 'Kiwix: Offline Wikipedia'
date: "2021-01-18T19:25:32+01:00"
tags:
- opensource
- 100daystooffload
incoming:
- url: https://chaos.social/@zerok/105578140826815297
---

While listening to one of the [recent episodes of the Ö1 radio program “Digital Leben”](https://oe1.orf.at/programm/20210114/624818/Kiwix-die-Wikipedia-als-Offline-Enzyklopaedie) I learnt about the [Kiwix project](https://www.kiwix.org/en/). The primary goal of Kiwix is to make Wikipedia (and other online resources) available to people with bad to very bad connectivity.

It does so by bundling websites into content packages stored as [ZIM files](https://openzim.org/) that can then be downloaded and viewed offline using the [Kiwix viewer](https://www.kiwix.org/en/download/) available for pretty much all platforms. To make things easy, the iOS app, for instance, also provides links to a huge selection of such archives including Wikipedia, Project Gutenberg and more. If you want to create a ZIM file out of your own website, you can also do so using [youzim.it](https://youzim.it). Sadly, there currently seem to be some compatibility issues between that service and Kiwix on [macOS](https://github.com/kiwix/apple/issues/290) and [iOS](https://github.com/kiwix/apple/issues/341) due to the use of service workers there (however that is related, no idea) 😞

While the contents in these archives are [compressed](https://github.com/openzim/libzim/blob/a4330a57bcd3da6249f2ddddc71f4e456dd21c93/src/cluster.cpp#L55) (with either LZMA, bzip2, zlib, or zstd), these ZIM files are still huge - Wikipedia with it’s 6 million articles and images takes around 85 GB - and if you don’t have that at your disposal on your device, the project also offers [Raspberry PI images which create a WiFi hotspot](https://www.kiwix.org/en/downloads/kiwix-hotspot/) for you where you can get to all this data.

At this point, unfortunately, I haven’t had all that much luck with non-Wikipedia archives but I still think this project is quite interesting. At this point, though, I don’t really have an opinion if this is the right approach for making web resources available offline, but since I also have a license of [Dash](https://kapeli.com/dash) for macOS which has a similar approach, I can definitely see the appeal 😅
