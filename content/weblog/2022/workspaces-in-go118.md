---
title: Workspaces in Go 1.18
date: "2022-01-14T19:16:12+01:00"
tags:
- golang
- development
---

While generics are definitely the start of [Go 1.18](https://tip.golang.org/doc/go1.18), the next stable release of the language (or actual the tool) will come with another cool new feature that will make working on multiple dependent packages at the same time much easier: Workspaces!

(Please note that I’m going to use Go 1.18beta1 for the examples below.)

## Updates of dependent packages

Let’s say you have two packages: `producer` and `consumer`. You want to implement a great new feature in `consumer` but you need some code changes introduced to `producer` first. Right now, you’d probably do those changes inside the `producer` package and then add a `replace` statement to `consumer`’s `go.mod` file pointing to the local (and modified) version of the `producer` package.

Once you’re done with everything, you need to do at least the following steps:

1. Publish the producer changes
2. Remove the replace statement from the consumer’s `go.mod` file
3. Update the version of producer specified inside the consumer’s `go.mod` file
4. Publish the consumer changes

Especially step 2 is annoying here if you have to do such orchestrated changes more than once. 

## Little example

Let’s put the structure above into an actual example:

```
❯ cd ~/tmp/go-workspace
❯ exa --tree
.
├── consumer (go.h10n.me/consumer)
│  ├── go.mod
│  └── main.go
└── producer (gitlab.com/zerok/go-workspace-demo-producer)
   ├── go.mod
   └── producer.go
```

The `producer` package just exposes a dummy struct called `Producer` which is the used within the consumer’s main function. In here, I’m using a new method (`Produce`) that is attached to that struct which is not yet available in a released version:

```go
package main

import producer "gitlab.com/zerok/go-workspace-demo-producer"

func main() {
	p := producer.Producer{}
	p.Produce()
}
```

If I now go into the `consumer` package and run `go build` without that whole thing being assigned to a workspace, I get something like this:

```
# go.h10n.me/consumer
./main.go:7:4: p.Produce undefined (type Producer.Producer has no field or method Produce)
```

The clone of the producer library that is located inside the `producer` folder *does* contain that method though:

```go
package producer

type Producer struct{}

func (p *Producer) Produce() {}
```

 So next I create a workspace that makes that clone available to the consumer:

```
❯ cd ~/tmp/go-workspace
❯ go work init
❯ go work use ./consumer
❯ go work use ./producer
```

There is now a new file in `~/tmp/go-workspace` with the name `go.work` which will contain something like this:

```
go 1.18

use (
	consumer
	producer
)
```

`use` just says that the paths that are mentioned should be considered when resolving packages before anything else. Note that the consumer also has to be declared inside the workspace in order to be assigned to it and use the other packages there.

If I now go back to the consumer package and run the build command again, it will work! The Go-tool found the workspace definition and the producer package within it that should *replace* the global package.

That’s it! No fiddling around with `replace` statements or other workarounds!

I can only imagine how much easier this would make the workflow for people who are working in large mono-repos or that have more than just two packages they need to develop in orchestration.

## Learn the details

If you want to take a deeper dive, take a look at the [original proposal](https://go.googlesource.com/proposal/+/master/design/45713-workspace.md) of the feature. Additionally, there is [documentation](https://pkg.go.dev/cmd/go@master#hdr-Workspace_maintenance) available for the new `go work` command group.
