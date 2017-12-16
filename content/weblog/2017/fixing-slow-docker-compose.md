---
title: "Fixing slow docker-compose"
date: 2017-12-16T20:53:38+01:00
draft: true
tags:
- docker
- network
---

For the last couple of weeks I had a weird issue with docker-compose: It was
simply unbelievably slow. Turns out, docker-compose tries to do some DNS
lookups and depending on the network you're in, that can take ages or even
fail.

[#3419](https://github.com/docker/compose/issues/3419) contains lots of hints
making me try adding various entries to the /etc/hosts file.

I already had `127.0.0.1 localunixsocket.local` in there but that didn't help.
tcpdump for the rescue! Starting docker-compose now caused (among others) the
following DNS lookup:

```
$ sudo tcpdump -vv -s 0 -l -n port 53
...
08:46:15.802327 IP (tos 0x0, ttl 255, id 35049, offset 0, flags [none], proto UDP (17), length 70)
    something.59024 > something.53: [udp sum ok] 3699+ A? localunixsocket.company.local. (42)
```

So I added `127.0.0.1 localunixsocket.company.local` and all of a sudden
docker-compose was blazingly fast again!

