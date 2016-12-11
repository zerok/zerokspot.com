---
date: 2016-12-11T12:09:28+01:00
tags:
- prometheus
- monitoring
title: Alert resolution timeouts in Prometheus
---

I'm currently using [Prometheus][] for quite a few services and esp. in
combination with Grafana and [AlertManager][] it has proven to be an extremely
handy tool. For instance, we usually have alerts for every single major
component of a service. If this component becomes unreachable, alerts are sent
to a specific Slack channel.

[prometheus]: https://prometheus.io/
[alertmanager]: https://github.com/prometheus/alertmanager

What bothered me, though, was how long it took for the resolution message to
arrive.  By default, it takes 5 minutes; luckily, though, you can customize this
in AlertManager's global settings:

```
global:
  resolve_timeout: 20s
```

This would set the timeout to only 20 seconds, feels much more usable to me
given that most of our check intervals are somewhere in the 5-15s range and the
alerts are set to something between a 10-20s range. I'm pretty sure I will be
tuning this setting in the future but for now this should do 😉

I have no idea why I didn't see this setting in the [documentation][] the first
time around, but now I'm glad I looked again (albeit using a detour through the
source code 😉)

[documentation]: https://prometheus.io/docs/alerting/configuration/#configuration-file
