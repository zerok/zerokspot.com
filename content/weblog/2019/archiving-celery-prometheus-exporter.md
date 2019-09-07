---
title: "Archiving celery-prometheus-exporter"
date: 2019-09-07T16:13:06+02:00
draft: true
tags:
- software
- oss
---

When I started the [celery-prometheus-exporter][cpe] project back in
[May 2016][init] I needed a way to monitoring the worker queue in an
internal project at work. Over the year's, though, that project and
therefore the exporter has faded into the background so I lacked time
to keep it properly maintained.

{{<figure src="/media/2019/cpe-archived.png" caption="celery-prometheus-exporter is in read-only mode">}}

The next iteration of the project I was using CPE for might not even
have Celery or Django in it anymore so I decided to archive the
exporter. I want to take this opportunity to thank those who
contributed over the years:

- [Rupert Angermeier][rangermeier]
- [Gabriel Abdalla Cavalcante][gcavalcante8808]
- [Brett Randal][javabrett]
- [Roberto Barreda][robertobarreda]
- [Dimitris Rozakis][dimrozakis]
- [Dmitriy][zokalo]
- [Neil Chudleigh][nchudleigh]
- [Gautier MATHON][gmathon-ovh]
- [Kostya Esmukov][KostyaEsmukov]

If the project has been useful to you, please fork and improve it 🙂
There are still tons of things that could be worked on but I simply
lack the time to even coordinate the work there. Hence I see archiving
this repository as the only responsible way. I will keep this
repository online as long as possible! Sorry for any inconvenience
this action causes and thanks for your understanding!


[init]: https://github.com/zerok/celery-prometheus-exporter/commit/7f15f269f3d215ded95098e596045e7c72668c1c
[cpe]: https://github.com/zerok/celery-prometheus-exporter/
[rangermeier]: https://github.com/rangermeier
[gcavalcante8808]: https://github.com/gcavalcante8808
[javabrett]: https://github.com/javabrett
[robertobarreda]: https://github.com/robertobarreda
[dimrozakis]: https://github.com/dimrozakis
[zokalo]: https://github.com/zokalo
[nchudleigh]: https://github.com/nchudleigh
[gmathon-ovh]: https://github.com/gmathon-ovh
[KostyaEsmukov]: https://github.com/KostyaEsmukov
