---
title: "Introducing concourse-piper"
date: 2017-09-10T10:36:46+02:00
tags:
- concourseci
---

Recently, I started playing around with [Concourse CI](https://concourse.ci/). I
wanted to have a more stream-lined, container focused solution for our internal
build and deployment pipeline and the project sponsored
by [Pivotal](https://pivotal.io/) sounded like a good fit. Dealing with a huge
number of similar jobs and resources, though, turned out to be rather
tedious. This is
where [concourse-piper](https://github.com/zerok/concourse-piper) comes in.

## Our use-case

The project where I wanted to deploy Concourse first was basically a large
mono-repo with about 10 micro-services, each in its own directory. Whenever
someone pushes a change to the `develop` branch, the CI system should run tests
on only that project, build a docker-image just for that service, and deploy it
to our [Kubernetes](https://kubernetes.io/) cluster. These services are either
NodeJS, Go, or Java-based. In addition to that, there is also one simple nginx
container that needs its JavaScript built using webpack.

As such, we have tons of test- and build-jobs that look quite similar but just
live in different folders. That pretty much screamed for a template-based
approach for generating our `pipeline.yml` file, which is
what [concourse-piper](https://github.com/zerok/concourse-piper) is trying to
provide.

## Bringing templates to the game

Every pipeline in Concourse consists of 4 groups of elements:

- resource types
- resources
- jobs
- groups (which act as collections of a subset of the three other categories)

So, concourse-piper is looking for `resource_types`, `resources`, `jobs`, and
`groups` folders which contain templates for YAML files. Let's look at a small
example: If your pipeline has multiple source-resources for the services `A` and
`B`, you'd create a file `source.yml` within the `resources` folder with the
following content:

```
meta:
  name_template: "source-{{.Instance}}"
  instances:
  - A
  - B
data:
  type: git
  source:
    uri: ssh://git@server.com/{{.Instance}}.git
    private_key: my-private-key
```

If you now execute `concourse-piper` it will create a file called
`pipeline.generated.yml` that looks like this:

```
groups: []
resource_types: []
resources:
- name: source-A
  source:
    private_key: my-private-key
    uri: ssh://git@server.com/A.git
  type: git
- name: source-B
  source:
    private_key: my-private-key
    uri: ssh://git@server.com/B.git
  type: git
jobs: []
```

Each such template file consists of a `meta` section, which defines what
instances should be generated, and the `data` section providing the template for
the content of the generated instance. If you want to generate a job that uses
one of our generated resources, you simply put a file with that structure into
the `jobs` folder. Same for your custom resource types and groups.

If you just have a singleton job, resource, resource-type, or group, you'd
normally just have one instance fined. For this special case there also exists a
shortcut: Just use `name` instead of `name_template` and don't specify any
`instances`:

```
meta:
  name: job-name
data:
  plan:
  - get: source
  ...
```


As template engine concourse-piper
uses [Go's text/template module](https://golang.org/pkg/text/template/) which
works great for the simple use-cases we have for templating here. In addition to
the built-in functions, it provides just two helpers:

- `ite` is basically `condition ? trueResult : falseResult`, which you might
  know from languages like C or Java:

  ```
  {{ ite .BooleanValue "yes" "no" }}
  ```
- `partial`, which allows you to re-use text snippets:

  ```
  {{ partial "something.yml" 2 . }}
  ```

  This will load a template stored under `partials/something.yml`, render it
  using `.` as context and indent it with two spaces. The indentation-step is
  necessary in order to be able to generate valid YAML.

That's basically it! If you want to know more, make sure to check out the
project's README.

concourse-piper is still very new and so far I've only used it for a handful of
pipelines (including one really large one) but so far it has helped me stay sane
even with that huge pipelines. If it works for you, too, please let me know :)
