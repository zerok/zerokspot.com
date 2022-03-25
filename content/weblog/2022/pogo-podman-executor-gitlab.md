---
title: 'Pogo: A GitLab Runner executor using podman'
date: "2022-03-25T17:20:41+01:00"
tags:
- golang
- gograz
- podman
- pogo
- gitlab
---

Last Monday, [Christoph Lipautz](https://twitter.com/lipdaguit) presented a custom GitLab Runner [executer](https://docs.gitlab.com/runner/executors/) that is using [podman](https://podman.io/) under the hood at our [monthly Go meetup in Graz](https://gograz.org/meetup/2022-03-21/): [Pogo](https://github.com/eyeson-team/pogo). The executor is the component in the GitLab CI system that actually runs jobs passed to it. For that it handles 4 life-cycle steps: (1) config, (2) prepare, (3) run, and finally (4) cleanup . The runner provided by GitLab for use in your own environment comes with the following executors:

- SSH
- Shell
- VirtualBox
- Parallels
- Docker
- Kubernetes

If you want to have something different, then you have to hook a so-called custom executor which is what Pogo now provides. In general, the solution provides pretty much the same level of abstraction as the Docker executor but does not require a service running as root. Especially for on-premise systems this might be quite interesting and you cannot simply provision and de-provision machines that easily but also don’t have to fear a job somehow gaining root access due to an issue within Docker.

Another nice feature of Pogo is that you can define custom volumes that should be mounted into the job’s execution environment depending on what tags are set for the job. Let’s say that you have a job tagged with `data-media` then you could automatically mount a `media` folder into the job using the following configuration of the executor itself:

```yaml
mounts:
  - volume: "/mount/shared-media:/mount/media"
    tags:
      - "data-media"
```

I’ve been thinking again and again about moving my repositories off of GitHub and over to GitLab again. When (not if) I finally decided to do that and once again use my own runners, then I’ll definitely take a look at Pogo! But first I should probably finally spend some time with podman, nerdctl and all the other non-Docker container stuff that is available on Linux 😅
