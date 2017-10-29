---
title: "Templating all the things!"
date: 2017-10-29T08:45:36+01:00
tags:
- tooling
---

Over the course of the last months I’ve grown to love Go’s built-in templating
language and started to apply it to all sorts of problems: First to more
complex [ConcourceCI](https://concourse.ci/) pipelines and recently also to
docker-compose files. The latter has ended up in a little CLI-tool called
[tpl](https://github.com/zerok/tpl), which simply takes a template, various
data points from the world around it and writes the rendering to stdout. It is
basically an extension to what I wrote about in ["Creating test-requests from
Vault"](https://zerokspot.com/weblog/2017/10/23/creating-test-requests-from-vault/)
a couple of days ago.

The initial use-case was that I had a [Vault](https://www.vaultproject.io/)
instance running on my local machine and wanted to expose its secrets to
containers within Docker. For that I needed to provide an external IP address
to docker stack:

```
version: "3"
services:
  core-service:
    external_hosts:
    - "vault:{{ .Network.ExternalIP }}"
    ...
```

That's the template I wanted to render to docker-compose. *tpl* to the rescue!

```
$ tpl docker-compose.yml.tpl | docker stack deploy --compose-file - project
```

(Just make sure to execute `docker stack deploy` in the same folder as the
docker-compose.yml.tpl file is located in order to make volumes mounts find
their folders.)

I also have tons of applications that could benefit from taking their
credentials from a secure store. These apps shouldn't have to know about Vault,
though, but, again, should be able to read their configuration from stdin:

``` 
$ cat app.yml.tpl
credentials:
  username: {{ vault-secret "secrets/app/creds" "username" }}
  password: {{ vault-secret "secrets/app/creds" "password" }}

$ tpl app.yml.tpl | app --config -
```

If you want to give it a spin, you can find the code on
[Github](https://github.com/zerok/tpl) but also on brew if you're on macOS:

```
$ brew tap zerok/main https://github.com/zerok/homebrew-tap
$ brew install zerok/main/tpl
```
