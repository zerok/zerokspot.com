---
title: "Logging in Docker"
date: 2017-11-17T18:56:18+01:00
tags:
- til
- docker
---

The more services you have, the more logging data you will have to manage. And
with Docker it has become even simpler to run services. 

Luckily, it comes with lots of options when it comes to managing your logging
data. The main mean of controlling how logging works inside Docker is through
so-called "logging drivers". There are a couple of built-in drivers and since
17.05 you can also provide custom drivers through plugins.

By default, logs are written as JSON and can be accessed using the `docker
logs` command for each container. If you want to use a journald installed on
the host-system, you can do so by setting `--log-driver journald` when starting
a container.

That's probably also the most interesting option for the environment I'm using
Docker in right now. The driver exposes a handful of fields with each log
statement that help you to find them:

* `CONTAINER_NAME`
* `CONTAINER_ID` and `CONTAINER_ID_FULL`
* `CONTAINER_TAG`
* `CONTAINER_PARTIAL_MESSAGE` for handling large log statements

Now I could, for instance, get all the log statements of the "my-app" container
using journalctl:

```
$ journalctl CONTAINER_NAME=my-app
```

There are also drivers available GELF, Syslog etc. You can find out more on
[docs.docker.com](https://docs.docker.com/engine/admin/logging/view_container_logs/). 


## Custom logging plugin

In case there is a target system that is not yet supported by Docker, you can
also create your own logging plugin which is specified again inside the [docker
documentation](https://docs.docker.com/engine/extend/plugins_logging/). Such
plugin is itself just another Docker image that offers a simple HTTP API.


## Disabling logging and more...

If you set `--log-driver none` for a container, you can suppress logging for
that container. This might be useful for systems where just too much log data
is generated and you cannot control  it at the source. 

In general, Docker offers pretty much anything you'd want when it comes to
logging. Luckily, the documentation is well written. Now I just have to find
out, what system I should reconfigure first 🙂
