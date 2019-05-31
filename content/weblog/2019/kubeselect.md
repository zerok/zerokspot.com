---
title: "KubeSelect: KUBECONFIG switching made easy"
date: 2019-05-31T16:08:25+02:00
tags:
- kubernetes
- productivity
- utils
---


I love working with Kubernetes but there has been one aspect of the way it treats configurations that bugged me: It's really hard to work with various kubectl configuration files (and the multiple contexts you can define in them). A co-worker recently mentioned that he has a shell script that allows him to quickly pick a configuration file from a specific folder and set the KUBECONFIG environment variable for the current shell based on that.

Sadly, that didn't work for me because I’m using contexts quite heavily. I usually structure my kubeconfigs like this:

* 1 kubeconfig file per cluster
* 1 context for each namespace/user-pair that I use more than once

So I needed something that would not only help me picking a KUBECONFIG environment value but also set the `--context` flag in kubectl. I have a certification exam next week and so I'm basically spending most of my time learning for that but in order not to go completely insane, I thought it would be nice to tackle this little issue of mine. The result of a bit of coding yesterday is [kubeselect](https://gitlab.com/zerok/kubeselect).

`kubeselect select` scans your `$HOME/.kube` folder for files ending with .yml, .yaml, or .conf and looks for contexts defined in them. It will provide you with a listing using [go-fuzzyfinder](github.com/ktr0731/go-fuzzyfinder). When you now pick an option, it will return environment variable commands for your shell:

```yaml
$ kubeselect select
export KUBECONFIG=$HOME/.kube/some-config.yaml
export KUBECTX=context-name

# To also set those inside your current terminal,
# run eval on it
$ eval $(kubeselect select)
```

Since I'm lazy, I've defined an alias for that:

```
$ alias ks='eval $(kubeselect select)'
```

The second part of kubeselect is a simple wrapper around kubectl that uses the `KUBECTX` environment variable to set the `--context` flag:

```
$ kubeselect run -- get nodes

# is the same as...
$ kubectl --context some-context get nodes
```

Again, aliases FTW:

```
$ alias k='kubeselect run -- '
```

## How to install

At this point there isn't a binary-release available but you can either take a binary produced by one of the [build jobs](https://gitlab.com/zerok/kubeselect/-/jobs?scope=finished) or install from the master branch using go-get:

```
GO111MODULE=off \
go get gitlab.com/zerok/kubeselect/cmd/kubeselect
```

## What's missing?

What's missing right now is a simple status-command that shows what configuration a call to `kubeselect run` would be using. This command could then, for instance, also be integrated into shell-prompts etc.

I hope I'll be able to get to this in the next couple of days 🙂
