---
title: "Docker Stack clashing with Docker-compose"
date: 2017-09-18T21:10:49+02:00
tags:
- docker
---

A couple of days ago I was trying to set up a new project using Docker Stack in
order to document a migration path from docker-compose. Sadly, right when I
wanted to start my services I noticed that Swarm couldn't keep a single replica
running. All I got was following error message:

```
starting container failed: User specified IP address is supported only when connecting
```

I tried pretty much everything short of re-installing Docker for Mac until I
stumbled upon [this bug report](https://github.com/moby/moby/issues/30336). Here
someone had the problem that there was a service instance left running from a
previous docker-compose execution. As I had previously launched docker-compose
with the same project that sounded like a possible cause for my issue.

In my situation, though, it wasn't a running container as I had stopped
everything before. What was left from the previous setup was just one part: the
network that docker-compose had automatically generated. Turns out,
docker-compose and docker stack are using the same network name which caused the
error. A simple...

```
$ docker-compose down
```

... resolved the issue for me and my services were finally starting up using
`docker stack deploy` 😃
