---
date: "2023-11-14T19:32:22+01:00"
tags:
- docker
- macos
- rosetta
- applesilicon
title: Rosetta integration in Docker Desktop feels broken
---

Over the course of the last couple of weeks I've been struggling heavily with [Docker Desktop](https://github.com/docker/for-mac) on my M2 Macbook Air. I often need to cross-compile code to `linux/amd64` or just use images that were created only for that platform. Unfortunately, there seem to be some issues around the Rosetta layer in Docker Desktop that prevents me from, for instance, compile Go code inside of a `linux/amd64` container. 

>   exec go build ERROR: process "/dev/.buildkit_qemu_emulator go build" did not complete successfully: exit code: 1

The Rosetta emulation layer has now been declared as GA with 4.25.0 but that issue still persists. At least I don't seem to be the only facing that as it is already documented on [docker/for-mac#6773](https://github.com/docker/for-mac/issues/6773). I still wanted to have a quick sample program that I could use to check if that error is still present and so I created it using Dagger: https://github.com/zerok/docker-rosetta-issue

Right now the only solution seem to be to disable Rosetta in the Docker settings but that makes working with *any* amd64 code slower by an order of magnitude. I quickly tested all of this also in [Orb](https://orbstack.dev) and for some reason there everything works. This means that I'm currently at a point where I can either wait for someone to fix the issue in Docker (or whatever layer is actually responsible for this), switch Rosetta emulation of (which makes Docker mostly unusable to me), or consider alternatives like Orb. I quickly also gave [Colima](https://github.com/abiosoft/colima) a try but fan into even weirder issues there.