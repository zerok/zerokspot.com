---
date: "2023-09-10T21:02:02+02:00"
tags:
- docker
- containers
- til
title: Docker images for multiple platforms
---

Ever since I got an M2 Mac, I noticed more and more warnings around the images I wanted to use to not be native to my platform. Well, I was aware that Docker images and their content were platform/architecture-specific, but I've never really "felt" the effects of that before except when I played around with Docker on Raspberry PIs 🤦‍♂️ So now that I have both an M2 Mac and an old Intel-based one I started to look into options for how to bundle my applications up for both target systems.

As far as I can tell, there are two options:

1. Create a separate image for every target like `myapp:latest-amd64` and `myapp:latest-arm64`
2. Have a single image (or at least make it look like a single image) that includes all the architectures that you want

The first approach is not all that user-friendly as users now have to be aware of what platform they are on. It's just easier for the user to just pull `myapp:latest` and let the Docker runtime handle the rest. 

I was curious how some of the tools I regularly use solved this and so I took a look at [goreleaser](https://hub.docker.com/r/goreleaser/goreleaser/tags). If you look at the `goreleaser/goreleaser:v1.20.0` entry on Docker Hub, you'll notice that it supports two platforms: `linux/amd64` and `linux/arm64`. Both of these have image digests listed so there seems to be a completely standalone image also available for each of them:

- `goreleaser/goreleaser:v1.20.0-arm64`
- `goreleaser/goreleaser:v1.20.0-amd64`

These two images are then somehow bundled into one entry for easier consumption. To find out more, I used `skopeo inspect --raw docker://goreleaser/goreleaser:v1.20.0`

```
{
   "schemaVersion": 2,
   "mediaType": "application/vnd.docker.distribution.manifest.list.v2+json",
   "manifests": [
      {
         "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
         "size": 2632,
         "digest": "sha256:b2c31539d194880b293399c7ac98bc032cb9bbb098c512b0abaa55c8c1541d34",
         "platform": {
            "architecture": "amd64",
            "os": "linux"
         }
      },
      {
         "mediaType": "application/vnd.docker.distribution.manifest.v2+json",
         "size": 2632,
         "digest": "sha256:8bbc939928d0c06c05208004c0f013957ff4d7669f25b5b850f63524f8c67e99",
         "platform": {
            "architecture": "arm64",
            "os": "linux"
         }
      }
   ]
}
```

This entry is actually a list of so-called [manifests](https://github.com/opencontainers/image-spec/blob/main/manifest.md) (one for each of the platforms). The runtime will then pick the right manifest for the platform you're on (linux/arm64 for M1/2 Macs, linux/amd64 for Intel in my case).

## How to create a multi-arch image?

So, how do you now create such a multi-platform image? The [Docker manual](https://docs.docker.com/build/building/multi-platform/) has a detailed guide with multiple options but I wanted to create a minimal example using this `Dockerfile`:

```
FROM alpine:3.18
ARG TARGETPLATFORM
ARG BUILDPLATFORM
RUN echo "Building on $BUILDPLATFORM for $TARGETPLATFORM" > /log
```

In this example, I want to have an image on DockerHub with the name `zerok/multiarch-test:latest` that I can use on linux/amd64 *and* linux/arm64.

To make platform specific images, I'll use buildx:

```
docker buildx build --platform linux/amd64 \
    --tag zerok/multiarch-test:latest-amd64 \
    --output=type=docker,dest=latest-amd64.tar .

docker buildx build --platform linux/arm64 \
    --tag zerok/multiarch-test:latest-arm64 \
    --output=type=docker,dest=latest-arm64.tar .
```

This will write the generated images to separate `.tar` files. For generating the manifest list I'll need to have these also available within the "normal" Docker engine:

```
docker load --input latest-arm64.tar
docker load --input latest-amd64.tar
```

Finally, with those images in place, I'll create a single manifest file and push that to DockerHub:

```
docker manifest create \
    --insecure zerok/multiarch-test:latest \
    --amend zerok/multiarch-test:latest-amd64 \
    --amend zerok/multiarch-test:latest-arm64

docker manifest push zerok/multiarch-test:latest
```

Actually, all of these steps can also be combined into a single command which is is handy if you don't need to process the images before pushing them to the registry:

```
docker buildx build \
    --platform linux/amd64 \
    --platform linux/arm64 \
    --tag zerok/multiarch-test:latest \
    --push .
```

## Some resources

Thanks to this little experiment/investigation I found a lot of useful resources around manifests and related topics:

- [skopeo](https://github.com/containers/skopeo/) is a tool for working with remote registries but also quite handy to "convert" images between formats (or just copy them around)
- The [OCI Manifest specification](https://github.com/opencontainers/image-spec/blob/main/manifest.md)
- [Official manual](https://docs.docker.com/build/building/multi-platform/) about multi-platform images
- [manifest-tool](https://github.com/estesp/manifest-tool) for working with manifests