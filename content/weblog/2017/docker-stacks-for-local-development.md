---
title: "Using Docker Stacks for local development"
date: 2017-09-03T15:22:55+02:00
tags:
- docker
- til
---

While looking into the rather
new
[healthcheck feature of Docker](https://zerokspot.com/weblog/2017/09/02/docker-healthchecks/) I
noticed that docker-compose seems more and more like a legacy tool
with [docker stack](https://docs.docker.com/get-started/part5/) acting as its
replacement. So I set out to see, if I could use that for local development with
the same ease I got from docker-compose.

Let's work on a small example application here containing three components:

1. An application backend
2. A frontend application
3. [Traefik](https://traefik.io/) as a reverse proxy (which makes it easier to
   add new services in the future)

You can find the whole setup
on [Github](https://github.com/zerok/docker-stack-localdev).

For this, we already have a simple `docker-compose.yml` file which we now want
to use with Docker Stack:

```
version: '3'

services:
  backend:
    build: ./backend
    volumes:
      - ./backend/backend:/app/backend
    labels:
      - "traefik.backend=api"
      - "traefik.enable=true"
      - "traefik.frontend.rule=PathPrefixStrip:/api/"

  frontend:
    image: nginx:1.13.3-alpine
    volumes:
      - ./frontend:/usr/share/nginx/html
    labels:
      - "traefik.backend=frontend"
      - "traefik.enable=true"
      - "traefik.frontend.rule=PathPrefix:/"

  reverse:
    image: traefik:1.3.7-alpine
    command: "--configfile=/etc/traefik/traefik.toml"
    volumes:
      - ./traefik/traefik.toml:/etc/traefik/traefik.toml
      - /var/run/docker.sock:/var/run/docker.sock
    ports:
      - 8088:8080
      - 8080:80
```


## Getting ready

The first thing you have to do, is to make your local machine act as a Docker
Swarm manager. Just initialising a new swarm locally does the trick here:

```
$ docker swarm init
```

After that you can create a new `docker-compose.yml`-based stack and deploy it
onto your local Docker Swarm with the following command:

```
$ docker stack deploy demo \
  --compose-file docker-compose.yml
```


## Moving from docker-compose

Most of our original file should work out of the box with Docker's new stack
feature, except for one little detail:

You can no longer only define a `build` property, which was used to tell
docker-compose where to find a Dockerfile for your custom container. If you
still use that, you will receive something like that as error:

```
$ docker stack deploy demo \
  --compose-file ../docker-compose.yml
Ignoring unsupported options: build

Creating network demo_default
Creating service demo_reverse
Creating service demo_backend
failed to create service demo_backend: Error response from daemon: rpc error: code = InvalidArgument desc = ContainerSpec: image reference must be provided
```

Instead, you will have to create the image beforehand and just use the `image`
property:

```
services:
  backend:
    image: docker-stack-backend:latest
    labels:
      - "traefik.backend=api"
      - "traefik.enable=true"
      - "traefik.frontend.rule=PathPrefixStrip:/api/"
```

Now we can start the stack again and it should work:

```
$ cd backend
$ make docker
$ cd ..
$ docker stack deploy demo \
  --compose-file docker-compose-for-stack.yml
```


## Updating the configuration

If you make any change to your `docker-compose.yml` file you can simply execute
the deploy-command again to update your stack.


## How to restart a single service?

Something that was very handy while when working with docker-compose was that
you could update/restart specific services using the `up` and `down` commands.

With Docker Stacks you'll have to use the scaling feature of Docker Swarm. If we
make a small change to the `main.go` of our backend service, we can now deploy
that by setting the number of replicas of the backend first to 0, replace the
binary, and then scale it back up to 1.

```
$ cd backend
$ make backend
$ docker service scale \
  demo_backend=0 \
  demo_backend=1 \
  --detach=false
```

## Concluding

... I have to say that I really like the new workflow and the options that it
provides. Sure, there are probably a few areas that still need some work, but
for my use-cases `docker stack` already works quite well right now 😀
