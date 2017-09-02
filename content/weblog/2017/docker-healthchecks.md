---
title: "Docker Healthchecks"
date: 2017-09-02T20:24:26+02:00
tags:
- til
- docker
---

While looking for ways to improve our service-startup in docker-compose I
stumbled upon a nice feature introduced with Docker 1.12, which lets images
expose a command to check that the created container is still working properly.

```
HEALTHCHECK --interval=5s --timeout=5s --start-period=15s --retries=5 CMD curl http://localhost:8080 || exit 1
```

The above line configures a healthcheck for a HTTP service running on
`localhost:8080` inside the container. This check does a HTTP GET request every
5 seconds (`--interval`). If the check passes, the status of the container will
be set to `healthy`. If it fails, it will try again 5 times (`—retries`) and
then set the health-status to `unhealthy`.

Services that take longer to start up (like most Java based applications) are
given a grace-period (`--start-period`) in which healthcheck failures don’t
count to the retry-limit until either the period is over or the healthcheck has
passed at least once.


## Healthchecks with docker run

Let’s say, we have a simple Postgres container that we want to monitor. The
official image doesn’t support this feature by default, but luckily the `docker
run` command has all the options discussed above exposed as command-line flags:

```
--health-cmd string
	Command to run to check health

--health-interval duration
	Time between running the check (ms|s|m|h) (default 0s)

--health-retries int
	Consecutive failures needed to report unhealthy

--health-start-period duration
	Start period for the container to initialize before starting
	health-retries countdown (ms|s|m|h) (default 0s)

--health-timeout duration
	Maximum time to allow one check to run (ms|s|m|h) (default
	0s)
```

Taking our Postgres container as an example, you can add healthchecks like this:

```
docker run --rm \
  -e POSTGRES_DB=testdb \
  -e PGDATA=/var/lib/postgresql/data/pgdata \
  -v $PWD/postgres_data:/var/lib/postgresql/data/pgdata \
  --health-cmd "pg_isready -U postgres" \
  --health-interval 5s \
  postgres:9.6
```

If you now inspect the global docker events you can see lots of `container
exec_start` and `container health_status` entries:

```
...
2017-09-02T19:14:00.033560520+02:00 container exec_create: /bin/sh -c pg_isready -U postgres 015f289346c02e3ac131948f39e0ca044a41641b9b41c23951862806db888c29 (image=postgres:9.6, name=pedantic_heisenberg)

2017-09-02T19:14:00.033701347+02:00 container exec_start: /bin/sh -c pg_isready -U postgres 015f289346c02e3ac131948f39e0ca044a41641b9b41c23951862806db888c29 (image=postgres:9.6, name=pedantic_heisenberg)

2017-09-02T19:14:00.165324872+02:00 container health_status: healthy 015f289346c02e3ac131948f39e0ca044a41641b9b41c23951862806db888c29 (image=postgres:9.6, name=pedantic_heisenberg)

2017-09-02T19:14:05.170512928+02:00 container exec_create: /bin/sh -c pg_isready -U postgres 015f289346c02e3ac131948f39e0ca044a41641b9b41c23951862806db888c29 (image=postgres:9.6, name=pedantic_heisenberg)
...
```

You can also see the healthiness of each running container in the output of
`docker ps`:

```
CONTAINER ID  IMAGE         COMMAND                 CREATED        STATUS                   PORTS     NAMES
015f289346c0  postgres:9.6  "docker-entrypoint..."  5 minutes ago  Up 5 minutes (healthy)   5432/tcp  pedantic_heisenberg
```

It is also available through `docker inspect`:

```
$ docker inspect 015f289346c0 | jq '.[0].State.Health.Status'
"healthy"
```

## Doing healthchecks in docker-compose

Getting back to my original use-case, healthchecks are also supported within
docker-compose (since version 2.1 of the configuration file format):

```
version: '2.1'

services:
	db:
		image: postgres:9.6
		healthcheck:
			test: ["CMD", "pg_isready", "-U", "postgres"]
			interval: 5s
			retries: 10
			start_period: 10s
			timeout: 2s
	app:
		# ...
		depends_on:
			db:
				condition: service_healthy
```

In the example above, the service “app” will wait for the db-service to be
healthy before starting up. Normally, it would only wait until the container has
been launched.

Sadly, the `condition` option within `depends_on` has been removed with version
3 (See [docker-compose (version 3): depends_on contains an invalid type, it should be an array](https://github.com/moby/moby/issues/30404)
for details).

According
to
[this comment by shin-](https://github.com/docker/compose/issues/4305#issuecomment-276527457) the
reason behind this is that Docker seems to be slowly deprecating docker-compose
in favour of `docker stack` and for that a unified approach is preferred.

## What about docker-compose and v3?

If you cannot easily migrate up to docker stacks but need a way to start a
service only if another service has reported that it’s ready, there are
currently only some workarounds available. As described in the official
documentation
([Controlling startup order in Compose](https://docs.docker.com/compose/startup-order/))
you should use some kind of wrapper script to wait for a resource to become
available within your own container.

Personally, I will most likely take a look at `docker stack` next as I want to
play around with the
new [secrets manager](https://docs.docker.com/engine/swarm/secrets/) in Docker
Swarm anyway 😊
