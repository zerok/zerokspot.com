---
title: Dagger pipelines and OpenTelemetry environment
date: "2023-01-31T21:48:29+01:00"
tags:
- dagger
- cicd
- opentelemetry
- o11y
incoming:
- url: https://chaos.social/@zerok/109785810097123025
---

Recently I tried to generate tracing data out of the pipeline for zerokspot.com and feed that into an external collector. Doing that with the `ci/main.go` itself was easy: Just set the relevant environment variables, integrate the OpenTelemetry SDK and off we go.

But now I also wanted to forward the trace ID and collector configuration into the `blog`binary so that it could also generate traces and send those to the same collector. That binary is first built by the pipeline and later used for things like generating yearly archives and the blogroll for the site.

There I noticed something unexpected: Turns out that BuildKit, which is used by [Dagger](https://dagger.io), injects its own `OTEL_*` environment variables into every container execution.

	client.Container().
	  From("alpine:3.17").
	  WithExec([]string{"env"}).
	  Stdout(ctx)

... will produce something like that:

	#4 env
	PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
	OTEL_TRACES_EXPORTER=otlp
	OTEL_EXPORTER_OTLP_TRACES_ENDPOINT=unix:///dev/otel-grpc.sock
	OTEL_EXPORTER_OTLP_TRACES_PROTOCOL=grpc
	OTEL_TRACE_PARENT=00-c4b86de5410773eac8a599cfd712a7f0-be5952f342e403ed-01
	TRACEPARENT=00-c4b86de5410773eac8a599cfd712a7f0-be5952f342e403ed-01
	HOME=/root

Setting my own  `OTEL_EXPORTER_OLTP_TRACES_{ENDPOINT,HEADERS,PROTOCOL}` variables won’t work. BuildKit will always override them. This was also confirmed [on Discord](https://discord.com/channels/707636530424053791/1069249221535993916). For now, I’ve solved this by just injecting some custom environment variables and then mapping them inside the `blog` binary to what’s expected by the OpenTelemetry SDK. 

Dagger/BuildKit comes with a built-in OTEL collector but I cannot use that in my current setup, so I will most likely stick with my little workaround for now. That being said, I will definitely keep an eye on future changes to that integration!
