---
date: 2017-04-23T10:34:00+02:00
title: Getting started with govendor
tags:
- golang
- development
---

Over the last couple of months I've been slightly moving away from Glide and
over to [govendor][], mostly to give it a try and later I simply kept using
it. Compared to Glide govendor feels a bit more minimal to me but I also really
like that it can import dependencies from the global GOPATH and provides a nice
status overview for all your dependencies. I still forget the basics on how to
use it properly all the time, though, so I thought a quick getting-started guide
might be in order, discussing the most common steps 😊

## Initialising a new project

Right after creating your project's root folder, run `govendor init`. This will
create a `vendor` folder and a basic (and mostly empty) `vendor.json` in it.

```
$ govendor init
$ cat vendor/vendor.json
{
	"comment": "",
	"ignore": "test",
	"package": [],
	"rootPath": "github.com/zerok/demo"
}
```

## Adding a dependency

The whole point of govendor is managing your dependencies. It allows you to do
so from two sources:

1. A remote repository (same as `go get`)
2. Your global GOPATH

I use (1) most often as it simply replaced `go get <dependency>` for me while
inside a project. Let's say, you want to use [Dave Cheney][]'s
awesome [errors][] package. Simply run `govendor fetch github.com/pkg/errors`
inside your project directory and it will download that and all its dependencies
into the vendor folder and update the `vendor.json` file accordingly.

```
$ cat vendor/vendor.json
{
	"comment": "",
	"ignore": "test",
	"package": [
		{
			"checksumSHA1": "ynJSWoF6v+3zMnh9R0QmmG6iGV8=",
			"path": "github.com/pkg/errors",
			"revision": "ff09b135c25aae272398c51a07235b90a75aa4f0",
			"revisionTime": "2017-03-16T20:15:38Z"
		}
	],
	"rootPath": "github.com/zerok/demo"
}
```

(2) is useful if you're offline and want to get started with a project and use
your global GOPATH as something like a cache for often used dependencies. Let's
stick with the errors package for now. Imagine you know that you will be on a
plane or train in a couple of hours where you will be offline. Then run
`go get -u github.com/pkg/errors`
before going offline, which will add the errors package
to your GOPATH. Now you start your new project (see above) while you're offline
and therefore cannot use `govendor fetch` anymore. Now you can use
`govendor add github.com/pkg/errors` instead which will do pretty much the same
as the fetch-command but takes the GOPATH as a source!

## Updating a dependency

So, upstream has fixed a bug or you want a feature that has been added after you
vendored your dependency. As with adding a new dependency, updating an existing
one can be done either online or offline. If you want to update a dependency
from a remote location, use the same `fetch` command you executed in the
previous step.

For offline-updating the steps are also quite similar. While you're still
online, update your dependency in the GOPATH using `go get -u <dependency>`.
Then you can use `govendor update <dependency>` to update your vendored version
based on the code inside the GOPATH.

## Removing unused dependencies

Before looking into removing dependencies, let's first take a look at one of the
great little gems hidden inside the `govendor` command: `govendor list`. This
will provide you with an overview of all the vendored dependencies and their
status inside your project. In the empty dummy-project where I've added the
pkg/errors package as dependency, the list would look like this:

```
vu github.com/pkg/errors
```

"v" means "vendored" while "u" indicates that the dependency is actually
"unused inside the project". The `--help` output describes a couple other status
indicators:

```
Status Types

	+local    (l) packages in your project
	+external (e) referenced packages in GOPATH but not in current project
	+vendor   (v) packages in the vendor folder
	+std      (s) packages in the standard library

	+excluded (x) external packages explicitely excluded from vendoring
	+unused   (u) packages in the vendor folder, but unused
	+missing  (m) referenced packages but not found

	+program  (p) package is a main package

	+outside  +external +missing
	+all      +all packages

	Status can be referenced by their initial letters.
```

Let's say now, that we want to remove the pkg/errors package again. This can be
done using `govendor remove github.com/pkg/errors`.

But if you actually just want to remove all the unused dependencies inside your
project, you can use the status types as a shortcut!

```
$ govendor add github.com/pkg/errors

$ govendor list
vu github.com/pkg/errors

$ govendor remove +unused

$ govendor list

```

That's it! This is everything I'm using govendor for right now and it's working
really well so far 😊 I hope this little guide will also help you!

But what about [dep][]? Sadly, I haven't had time to look into it yet. Hopefully
this will change in the coming months and I will finally be able to step writing
a blog post about yet another vendor-management tool every couple of months 😉

[errors]: https://github.com/pkg/errors
[govendor]: https://github.com/kardianos/govendor
[dep]: https://github.com/golang/dep
[dave cheney]: https://dave.cheney.net/
