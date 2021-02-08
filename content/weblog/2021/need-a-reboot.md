---
title: Need a reboot?
date: "2021-02-08T16:32:46+01:00"
tags:
- prometheus
- monitoring
- ops
- 100daystooffload
---

Over the weekend I wanted to improve my monitoring stack for my personal systems. One thing that I kept missing is when a server would needs a reboot due to an update. While there are options to also configure unattended reboots along unattended updates, I just wanted to get notified by Prometheus and then schedule the actual reboot manually.

To get that state into Prometheus I created a [little exporter](https://gitlab.com/zerok/reboot-required-exporter) that creates the `reboot_required` metric based on `/var/run/reboot-required` file on Debian.

That exporter can be run in two ways: (1) as a standalone HTTP server or (2) as a command that just prints the metric to stdout which can be redirected into, for instance, node-exporter’s `textfile.d` folder.

Since I already have a node-exporter running on pretty much every server I'm using the latter mode for now:

	$ reboot-required-exporter --one-shot > /etc/node_exporter/textfile.d/reboot_required.prom
	
	…waiting…
	
	$ curl http:// localhost:9100/metrics
	…
	reboot_required 1
	…

If that sounds useful to you, please give it a try too 🙂
