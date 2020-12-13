---
title: Testing Prometheus alerts
date: "2019-11-25T20:03:00+01:00"
tags:
- prometheus
- testing
incoming:
- url: https://www.reddit.com/r/PrometheusMonitoring/comments/e1l6ml/getting_started_with_testing_prometheus_alerts/
---

One of [Prometheus][p]' core features is that you can not collect metrics but also declare alerting rules. You'd normally configure Prometheus to connect with [Alertmanager][a] in order to get notified if one of these rules triggered an alert. You could test these rules but just waiting for the moment when they match and then wait for an alert to show up in your inbox. There is a more practical and less time-consuming way, though:

promtool, the CLI utility that comes with Prometheus, has a `test rules` command which you can use to run [test-cases against alerting rules][d]. For zerokspot.com, for instance, I have a node-exporter running and when. that exporter has been unreachable for more than 5 minutes, I want to get notified. To achieve that I have added the following rule into a `alerts/zerokspot.alerts.yaml` file:

```
groups:
  - name: zerokspot
    rules:
      - alert: zerokspot--sys--down
        expr: up{job="zerokspot-sys"} == 0
        for: 5m
```

Alongside that file I've also created another file which contains tests for this rule: One to make sure that the alert fires after 5 minutes and one that checks that it doesn't fire before:

```
rule_files:
  - zerokspot.rules.yaml

tests:
  - input_series:
      - series: up{job="zerokspot-sys"}
        values: 0

    alert_rule_test:
      - eval_time: 4m
        alertname: zerokspot--sys--down
        exp_alerts: []

      - eval_time: 5m
        alertname: zerokspot--sys--down
        exp_alerts:
          - exp_labels:
              alertname: zerokspot--sys--down
              job: zerokspot-sys
```

With `rule_files` I can define, for which rules files the `tests` apply. A test then consists of test data inside the `input_series` property and a number of `alert_rule_test`-cases. 

The first test simulates waiting for 4 minutes (the counterpart to the `for` property of an alerting rule) and does not expect any alerts (`exp_alerts: []`). The second test waits for 5 minutes and then expects a single alert instance with two labels: `alertname: zerokspot--sys--down` and `job: zerokspot-sys`.

These tests can then be executed with the following command:

```
$ promtool test rules alerts/zerokspot.tests.yaml
```


## Testing PromQL expressions

You can also use the same facility to test PromQL expressions. On the same level as the `alert_rule_test` you can also specify `promql_expr_test` cases with a similar structure:

```
    promql_expr_test:
      - expr: up
        exp_samples:
          - labels: '{__name__="up", job="zerokspot-sys"}'
            value: 0
```

Taking the same input series as before, this test checks what data is included when just `up` is queried. A single sample should be the result with the value 0.


## Testing time

There are a handful of functions inside PromQL that let you deal with timular data. For instance, if I'd like to just raise an alert if zerokspot-sys were down on Mondays, I could do it like this:

```
- alert: zerokspot--sys--down__monday
  expr: up{job="zerokspot-sys"} == 0 and on() (day_of_week() == 1)
```

According to [#4817](https://github.com/prometheus/prometheus/issues/4817) you should be able to set the output of the `time()` function using the `eval_time` property of `prom_expr_test` or `alert_rule_test`. Let's give that a try:

```
rule_files: []

tests:
  - promql_expr_test:
      - expr: time()
        eval_time: 10s
        exp_samples:
          - labels: '{}'
            value: 10
      
      # eval_time allows you to time-travel from 1970-01-01, which 
      # was a Wednesday
      - expr: day_of_week()
        eval_time: 10s
        exp_samples:
          - labels: '{}'
            value: 4
      
      # Moving the clock forward by 4 days we will reach the next Monday:
      - expr: day_of_week()
        eval_time: 96h
        exp_samples:
          - labels: '{}'
            value: 1
```

Because of that I would have expected the following check to work:

```
- input_series:
    - series: up{job="zerokspot-sys"}
      values: 0

  alert_rule_test:
    - eval_time: 96h
      alertname: zerokspot--sys--down__monday
      exp_alerts:
        - exp_labels:
            alertname: zerokspot--sys--down__monday
            job: zerokspot-sys
```

For some reason, I couldn't get the alert to fire, though:

```
Unit Testing:  rules/zerokspot.tests.yaml
  FAILED:
    alertname:zerokspot--sys--down__monday, time:96h0m0s,
        exp:"[Labels:{alertname=\"zerokspot--sys--down__monday\", job=\"zerokspot-sys\"} Annotations:{}]",
        got:"[]"
```

I haven't yet found the fault in my setup but, luckily, Prometheus also supports custom [recording-rules][r], which can be used to create "pre-calculated" records inside the datastore. In my case I use those to generate a new series called `my_current_day` which is calculated inside rules/times.test.yaml:

```
groups:
  - name: times
    interval: 1h
    rules:
      - expr: day_of_week()
        record: my_current_day
```

This also requires that I change the alert's expression but it also makes the tests more readable IMO:

```
# Alert:
#-------
- alert: zerokspot--sys--down__monday
  expr: up{job="zerokspot-sys"} == 0 and on() (my_current_day == 1)

# Tests:
#-------
# Note that I DO NOT INCLUDE the records file into the tests since
# I want to set the generated records to different values.
rule_files:
  - zerokspot.alerts.yaml
  
## Day 4 is not Monday
- input_series:
    - series: up{job="zerokspot-sys"}
      values: 0
    - series: my_current_day
      values: 4

  alert_rule_test:
    - alertname: zerokspot--sys--down__monday
      exp_alerts: []

## Day 1 IS not Monday
- input_series:
    - series: up{job="zerokspot-sys"}
      values: 0
    - series: my_current_day
      values: 1

  alert_rule_test:
    - alertname: zerokspot--sys--down__monday
      exp_alerts:
        - exp_labels:
            alertname: zerokspot--sys--down__monday
            job: zerokspot-sys

```

If anybody could tell me what I'm doing wrong with regards to `day_of_week()`, though, I'd really appreciate it! That being said, my little record-workaround seems to work and I kind of like what it's doing for the tests...

I'm currently in the process of adding more and more test cases to my alerting rules. As my alerts become more and more complex, testing becomes even more crucial. The more "state" I can simulate with simple series-data the better, IMO. That being said, I really want to know what I'm doing wrong, but I repeat myself ;-)

[d]: https://prometheus.io/docs/prometheus/latest/configuration/unit_testing_rules/
[p]: https://prometheus.io
[a]: https://prometheus.io/docs/alerting/alertmanager/
[r]: https://www.prometheus.io/docs/prometheus/latest/configuration/recording_rules/

## Update

elbuenodefali has posted [some instructions on GitHub][gh] on how to improve
the test series for resiliently mocking days across a week. Thank you!

[gh]: https://github.com/zerok/zerokspot.com/issues/160
