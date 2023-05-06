---
title: Learn the image ID of a Docker image file
date: "2023-05-06T21:51:35+02:00"
tags:
- til
- docker
- dagger
incoming:
- url: https://chaos.social/@zerok/110323501433274869
---

A couple of days ago I had the problem that I needed to know the ID of a Docker image I had exported previously with Dagger. What I wanted to achieve was to retag an image that I only had available as [Docker image file](https://github.com/moby/moby/blob/v23.0.5/image/spec/v1.2.md).

```
docker load --input some-image.img
docker tag $IMAGE_ID zerok/project:latest
```

The problem was, that Dagger (when creating the export) didn’t expose the generated ID (`sha256:...`) and parsing the output of the `docker load` command would have been  quite tedious. So, I was looking at other ways to get to that ID.

Turns out, it is actually stored inside the `manifest.json` that is part of the image tarball:

```
tar -xf some-image.img
cat manifest.json | jq '.[0].Config'
```

This will print something like `"042a816809aac8d0f7d7cacac7965782ee2ecac3f21bcf9f24b1de1a7387b769.json"` 

A more or less complete solution for that using Dagger could look like the this:

```go
package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"dagger.io/dagger"
)

func main() {
	ctx := context.Background()
	dc, err := dagger.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to Dagger: %s", err)
	}

	pwd := dc.Host().Directory(".")
	idFilename, err := dc.Container().From("alpine:3.17").
		WithExec([]string{"apk", "add", "--no-cache", "jq"}).
		WithMountedDirectory("/src", pwd).
		WithWorkdir("/src").
		WithExec([]string{"tar", "-xvf", "some-image.img"}).
		WithExec([]string{"jq", "-r", ".[0].Config", "manifest.json"}).Stdout(ctx)
	if err != nil {
		log.Fatalf("Failed to get config file: %s", err)
	}

	fmt.Printf("sha256:%s\n", strings.TrimSuffix(strings.TrimSpace(idFilename), ".json"))
}

```

Yes, I could have used Go’s `tar` package to get to that value, but where would have been the fun in that?! 🤪
