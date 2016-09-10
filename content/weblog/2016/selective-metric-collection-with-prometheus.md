---
date: '2016-05-26T17:02:04-07:00'
language: en
tags:
- prometheus
- ops
- monitoring
- golang
title: Selective metric collection with Prometheus
---

On my way to getting to know Prometheus I've noticed that many exporters for it
expose tons of data, most of which I'd probably never use in my monitoring and
alerting setup. The [node exporter][], for instance, also includes data about
the exporter itself (like how many go-routines are active etc.). I guess
collecting all that data for no practical reason won't be all that beneficial
with regards to memory and disk usage so I started to look at option to get
Prometheus to only collect what I actually want to work with later on.

Ignoring some weird proxying approaches there are probably only the following
two ways to get what I want:

1. Prevent the exporters to even expose the data or
2. prevent Prometheus from storing them

The first approach completely depends on how configurable the exporters I use
are. For many of my tools I'll probably use the official
[Go client library][goclient], which also exposes generally useful metrics from
the underlying run-time like the number of go-routines or GC behaviour. These are
enabled by default but I found a slightly hacky way to get rid of them:

```
// ...
func main() {
	http.Handle("/metrics", prometheus.UninstrumentedHandler())
	http.ListenAndServe("127.0.0.1:9999", nil)
}

func init() {
	prometheus.Unregister(prometheus.NewGoCollector())
    // Register your own collector
}
```

With other collectors (like if they are already used by an app like the node
exporter or are simply not as configurable as the Go one) I will have to go with
option 2. For this, Prometheus offers the [`metric_relabel_configs`][mrc]
setting for each scraper, where the labels for the metrics that
should be collected can be manipulated:

```
global:
  scrape_interval: 5s
scrape_configs:
  - job_name: "myservice"
    target_groups:
      - targets:
          - "localhost:9999"
    metric_relabel_configs:
      - source_labels: [__name__]
        regex: go_(.*)
        action: drop
      - source_labels: [__name__]
        regex: go_(.*)
        action: drop
```

This tells Prometheus to drop the metrics with names starting with `go_` or
`http_`. The label `__name__` is an internal label of every metric that
represents the name of that metric.

I hope that with a combination of these two approaches I will be able to keep
the number of unused data within Prometheus to a minimum 😊

[mrc]: https://prometheus.io/docs/operating/configuration/#metric_relabel_config
[goclient]: https://github.com/prometheus/client_golang
[node exporter]: https://github.com/prometheus/node_exporter
