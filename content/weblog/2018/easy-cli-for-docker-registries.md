---
title: "Easy CLI for Docker Registries"
date: 2018-12-09T17:05:30+01:00
tags:
- docker
- golang
- sideproject
---

Whenever I work with Docker images I usually have to deal with multiple registries. Some are on Gitlab instances, other on Nexus servers, and so on. As I'm administrating some of them I also have to sometimes work with multiple accounts on the same registry.

That last scenario reminded me of how [httpie](https://httpie.org/) is dealing with this: You can define so-called sessions for every host. Such a session can contain things like basic auth credentials, headers, and so on. I also wanted to get to know the [Docker Registry API](https://docs.docker.com/registry/spec/api/) in more detail. 

[container-inventory](https://gitlab.com/zerok/container-inventory) will hopefully eventually solve all of that for me.

## How to use

Let's say, you have a private registry installed on docker.company.com. The first thing I'll have to do is create a "default" session for that host:

```
$ container-inventory --registry docker.company.com \
  --session default
  --username USERNAME
  --password PASSWORD
```

Now you can list all the repositories on that registry with the `repositories list` (or for short `rs ls`) command:

```
$ container-inventory --registry docker.company.com rs ls
project/image1
project/image2
```

To get all the tags of a specific repository, you can use the `tags ls` command:

```
$ container-inventory --registry docker.company.com \
  tags ls --repository project/image1
1.0.0
1.0
latest
```

If you want to get complete repository paths that you can use, for instance, with `docker run`, use the `--long` flag:

```
$ container-inventory --registry docker.company.com \
  tags ls --repository project/image1 --long
docker.company/project/image1/1.0.0
docker.company/project/image1/1.0
docker.company/project/image1/latest
```

## v1.0.0

For now, I've provided binaries for Linux and macOS (both 64-bit). You can get them including signed checksums on Gitlab:

[Download on Gitlab](https://gitlab.com/zerok/container-inventory/tags/v1.0.0)

## What's next?

Perhaps as the next feature after 1.0 is a way to visualize what tags are actually just aliases of one another. That's probably possible by looking at the manifest of each tag, but perhaps there are also other ways. We'll see!

## Things I’ve learnt

As one of the goals of this project was to learn new things, I thought it would be useful to include all the learnings here:

- Nexus doesn’t provide a Docker-Content-Digest for manifests. That's a bit annoying as these would be quite useful when working with manifests.
- GitLab doesn’t offer a convenient way to host release binaries. That means that the container-inventory files are hosted on S3 for now.
- [goreleaser](https://goreleaser.com/customization/#S3) supports S3 for automated release uploads.
