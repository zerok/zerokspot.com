---
title: Testing Promtail pipelines
date: "2023-01-25T20:04:20+01:00"
tags:
- ops
- grafana
- logging
- observability
incoming:
- url: https://chaos.social/@zerok/109751559102541892
---

If you want to get data into Grafana Loki then [Promtail](https://grafana.com/docs/loki/latest/clients/promtail/) probably the easiest way to do so. In my case, I wanted to forward Caddy’s access logs to Loki. There are some fields, though, that I don’t want to send over to Loki, like a user’s IP or port number.

There are two ways I could do that: Either configure Caddy to not even log this fields in the first place, or not submitting them to Loki in Promtail. Since I wanted to play around with [Promtail’s pipeline feature](https://grafana.com/docs/loki/latest/clients/promtail/pipelines/) anyway, I chose the latter option for this little experiment.

Promtail has a couple of flags that make experimenting with configuration changes quite convenient. The primary ones here are `-stdin` and `-dry-run`. Let’s say I have a little configuration file located in `./promtail.yaml` and want to test some dummy data located inside `test.log` against that. Then I can run the following command to pipe that data through Promtail *without* updating any state files or submitting data to Loki but instead logging it to the terminal:

```
cat test.log | promtail -stdin -dry-run -config.file promtail.yaml
```

This will give me the transformed log statements like this:

```
Clients configured:
----------------------
url: http://localhost:30000
batchwait: 1s
batchsize: 1048576
follow_redirects: false
enable_http2: false
backoff_config:
  min_period: 500ms
  max_period: 5m0s
  max_retries: 10
timeout: 10s
tenant_id: ""
stream_lag_labels: ""

2023-01-25T18:32:28.3600988+0100	{level="info"}	{"level":"info","request_host":"zerokspot.com","request_method":"GET","request_path":"/tags/licensing/","response_status":"200"}

2023-01-25T18:32:34.4030662+0100	{level="info"}	{"level":"info","request_host":"zerokspot.com","request_method":"GET","request_path":"/weblog/2005/03/09/links-on-delicious/","response_status":"200"}
```

If I then want to see more of what’s actually happening in each stage of the pipeline, I can also pass the `-inspect` flag which will give me something like this:

```
[inspect: timestamp stage]:
{stages.Entry}.Entry.Entry.Timestamp:
	-: 2023-01-25 19:30:34.767041 +0100 CET
	+: 2023-01-25 18:32:28.3600988 +0100 CET
2023-01-25T18:32:28.3600988+0100	{level="info"}	{"level":"info","request_host":"zerokspot.com","request_method":"GET","request_path":"/tags/licensing/","response_status":"200"}
```

That’s pretty much all I needed to iterate on my configuration in order to just submit the data that I wanted to Loki 😊  And just for completeness’ sake, this is the pipeline I’m currently experimenting with:

```yaml
# Dummy client for local testing
clients:
  - url: http://localhost:30000

scrape_configs:
  - job_name: access_log
    pipeline_stages:
      # Extract all the fields I care about from the
      # message:
      - json:
          expressions:
            "level": "level"
            "timestamp": "ts"
            "response_status": "status"
            "request_path": "request.uri"
            "request_method": "request.method"
            "request_host": "request.host"
            "request_useragent": "request.headers.\"User-Agent\""
      # Promote the level into an actual label:
      - labels:
          level:

      # Regenerate the message as all the fields listed
      # above:
      - template:
          # This is a field that doesn't exist yet, so it will be created
          source: "output"
          template: |
            {{toJson (unset (unset (unset . "Entry") "timestamp") "filename")}}
      - output:
          source: output

      # Set the timestamp of the log entry to what's in the
      # timestamp field.
      - timestamp:
          source: "timestamp"
          format: "Unix"

```
