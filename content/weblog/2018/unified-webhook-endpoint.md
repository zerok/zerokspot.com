---
title: "Unified webhook endpoint"
date: 2018-06-11T18:57:30+02:00
tags:
- zerokspot
- toolbox
---

Last weekend I finally wanted to invest the time to automatically deploy
[zerokspot.com](https://zerokspot.com) whenever there was a change in its
source repository. Up until then, I basically had to open my laptop and execute
a simple deployment script to build all the pages using Hugo and update the
search index with the latest content.

What I ended up doing was to first put the repository onto Gitlab in order to
integrate their CI/CD system. I didn't want to give Gitlab SSH access to my
primary webserver, though, so I decided to look into what webhooks Gitlab
offered to notify my server that there was a new build to be deployed. Turns
out, there is a collection of [pipeline
events](https://docs.gitlab.com/ce/user/project/integrations/webhooks.html#pipeline-events),
one for every state of a pipeline. (I only cared about the with
`status=success`.)

Next, I started implementing my own little HTTP endpoint that would do the
actual deployment once it got called with the right secret. It should download
all the artefacts created by the pipeline, move them into the right folders,
and restart the search server.

Luckily, I stumbled upon [Adnan Hajdarević's webhook
project](https://github.com/adnanh/webhook/), which already does the first part
in a generic fashion. It allows you to define endpoints for your webhooks,
filter calls against those endpoint for matching attributes (within the body,
URL, or HTTP header), and call an executable.

The whole configuration can be put into a simple JSON or YAML file (I don't
like JSON for configuration files, so YAML it is…), which looked somehow like
this for my setup:

```yaml
---
- id: "zerokspot"
  execute-command: "/path/to/script.sh"
  command-working-directory: "/working/dir/"
  pass-arguments-to-command:
    - source: string
      name: "--access-token"
    - source: string
      name: "<GITLAB API TOKEN>"
  pass-file-to-command:
    - source: entire-payload
      name: payload
  trigger-rule:
    and:
      - match:
          type: value
          value: pipeline
          parameter:
            source: payload
            name: "object_kind"
      - match:
          type: value
          value: success
          parameter:
            source: payload
            name: object_attributes.status
      - match:
          type: "value"
          value: "<GITLAB SECRET>"
          parameter:
            source: "header"
            name: "X-Gitlab-Token"

```

Here I've defined a webhook for "zerokspot" (which means that it will be
reachable on `http://domain/hooks/zerokspot`) and will trigger
`/path/to/script.sh` but **only if** the request has the JSON properties
`object_kind=pipeline` and  `object_attributes.status=success` in its
body/payload and `<GITLAB SECRET>` as the value of the `X-Gitlab-Token` header.

The last part still required a bit of custom coding in order to download the
artefacts and move the around, but Adnan's tool still helped a lot and thanks
to it I could roll automated deployment of my site out last night 😃
