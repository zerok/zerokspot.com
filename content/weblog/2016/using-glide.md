---
date: '2016-05-17T22:32:35+02:00'
language: en
tags:
- golang
- development
title: Using Glide for vendoring
---

Disclaimer: This is basically the long version of a lightning talk I gave
recently at our local [Go user-group in Graz][gograz].

I really enjoy writing Go. The primary things I love are interfaces, the really
great standard library, and how you distribute your applications. Not much can
beat having a single binary that you can simply hand to a co-worker and let it
do its magic there, no matter what operating system they are on.

Sadly, nothing is perfect and so I'm always looking jealously at Rust, Node, or
Elixir (just to name a few) that have a far better dependency management story
than Go does. Things have improved over the years thanks to first-class
vendoring support initially as an opt-in experiment with 1.5 and now as opt-out
feature in 1.6. That being said, there is still no single-best solution out
there for how to manage that vendor folder, how to explicitly declare your
dependencies and their versions.

Tools like `go get` are still too focused on the `$GOPATH` and most of the time
leave out the `vendor` folder.

This is something [Glide][] is trying to solve. It all starts off very simple:
Create your Go project like you'd normally do on your `$GOPATH`. Then run...

[glide]: http://glide.sh/

```
$ glide init
```

This will create a single `glide.yaml` file which holds your project's metadata
as well as high-level information about its dependencies. Let's say we have a
little tool that needs a bit more command-line argument parsing than the `flag`
module provides out of the box. Jeremy Saenz' [cli][] library is quite popular
here, so let's add that to our project.

[cli]: https://github.com/codegangsta/cli

```
$ glide get https://github.com/codegangsta/cli
```

This will download cli and put it right under the project's `vendor` folder:

```
$ tree vendor
vendor
└── github.com
    └── codegangsta
        └── cli
            ├── CHANGELOG.md
            ├── LICENSE
            ├── README.md
            ├── ...

5 directories, 29 files
```

Both, `glide.yaml` and `glide.lock`, now hold references to this dependency:

```
$ cat glide.yaml
package: github.com/zerok/lala
import:
- package: github.com/codegangsta/cli

$ cat glide.lock
hash: 4ef4d86e8c6c6831ef597baca9ae65b38db380a2449070f868b6c5b4d1157a1d
updated: 2016-05-08T09:24:36.81347661+02:00
imports:
- name: github.com/codegangsta/cli
  version: d3a4d5467b9d41ee3b466a3d6684cdc36482e471
devImports: []
```

Let's say there was a bug in cli that got fixed upstream and you want to
update your vendor content to that version?

```
$ glide update
```

This will fetch the latest version that still matches the version requirements
specified within the `glide.yaml` file (we haven't specified a version there, so
it will just fetch the latest). And this is also what sets Glide apart from
other vendoring tools I've tried so far: It supports version ranges. If a
library creates releases according to semantic versioning, Glide will probably
be able to handle it.

When you fetch a dependency, you can also specify a version. Let's say we want
cli at version 1.x so that we don't run into any planned backwards-incompatible
changes during an update:

```
$ glide get github.com/codegangsta/cli#~1
```

The resulting `glide.yaml` file will look quite similar to what we had before
but with a little note on the dependency's version:

```
package: github.com/zerok/lala
import:
- package: github.com/codegangsta/cli
  version: ~1
```

So if version 1.20 is released and we last executed `glide get` or `glide update`
with version 1.19, this dependency will be updated to 1.20. If you want
to know more about the syntax Glide supports for specifying versions and version
ranges, this has its own chapter in the [documentation][doc].

[doc]: http://glide.readthedocs.io/en/latest/versions/

OK, that was all about `glide.yaml` but what's with that other file that Glide
manages?  `glide.lock` points to the actual commit that was fetched for the
requested version so that any `glide install` would allow you to restore the
version you got when you first fetched that dependency. This is only updated
when you either run `glide get` or `glide update`.

But how does Glide even determine a library's version given that Go doesn't
really provide any metadata related to that out of the box? Let's stick with the
cli example a bit longer.

This package doesn't have any metadata files in its repository but there are a
handful of tags, each prefixed with a "v" that represent versions. Looking
through Glide's code a bit it it seems like it uses names of tags and branches
to determine versions. Take a look the `SemVerRegex` within the [semver][]
package for details.

[semver]: https://github.com/Masterminds/semver

This way Glide simply leverages a common best-practice (to tag your releases) to
get all the metadata it requires. Nice 🙂

Given that the whole dependency-story is still an open issue in the Go
community, Glide probably won't be the last attempt at solving it. For now,
though, it looks like a nice compromise of using what is provided by the
language itself (vendoring within a project's directory tree) and putting a
little semantical layer on top of that for more convenience and
reproducability. [govendor][] looks like another tool with a similar approach
but I haven't tried it yet. It seems to also support versions but to a bit
smaller degree than Glide does.

Exciting times 😊

[gograz]: http://gograz.org
[govendor]: https://github.com/kardianos/govendor
