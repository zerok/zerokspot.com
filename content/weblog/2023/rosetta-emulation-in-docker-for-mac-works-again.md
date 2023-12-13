---
date: "2023-12-13T19:39:11+01:00"
tags:
- docker
- macos
- applesilicon
- rosetta
title: Rosetta emulation in Docker For Mac works again
---

Just a quick one: [Last month](https://zerokspot.com/weblog/2023/11/14/rosetta-docker-desktop-feels-broken/) I wrote about an issue I had with Go builds failing inside of linux/amd64 Docker containers when Rosetta was enabled. This issue has been fixed and should no longer appear with [Docker For Mac 4.26.0](https://docs.docker.com/desktop/release-notes/#4260) 🥳 

Big thanks to [David Gageot](https://github.com/dgageot) and presumably other folks from Docker who investigated [docker/for-mac#6773](https://github.com/docker/for-mac/issues/6773) and fixed it.