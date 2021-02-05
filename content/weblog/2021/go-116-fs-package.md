---
title: 'Go 1.16 feature: The fs module'
date: "2021-02-05T20:46:39+01:00"
tags:
- golang
- 100daystooffload
---

Embed isn’t the only new package that made it into [1.16](https://tip.golang.org/doc/go1.16). There’s now also `io/fs` which contains lots of functionality that was previously part of `io/ioutil` or something like [spf13/afero](https://github.com/spf13/afero): A general abstraction on top of file-system operations and navigation which is currently limited to read-operations.

The `io/fs` package provides primarily the new `FS` interface:

	type FS interface {
	    Open(name string) (File, error)
	}

There are also a couple of other interfaces that a file-system might implement like `StatFS`, `SubFS`, or `ReadDirFS` which allow implementors adding things like more performance stat access or directory handling.

All of that would be rather pointless if there wasn’t also a way to actually create something that implements those interfaces. For working with an actual OS file-system, the `os` package offers the new `DirEnv(path string) (fs.FS, error)` function.

If you want to have an instance of such a filesystem, you can use the new `os.DirEnv(path string) (fs.FS, error)` function. Sadly, the FS instance that is returned here really only implements the `Open` method. That being said, the fs-package offers some helper methods for iterating over directories or just listing all their entries that can work with this:

- `fs.WalkDir(fsys fs.FS, root string, fn fs.WalkDirFn) error`
- `fs.ReadDir(fsys fs.FS, name string) ([]fs.DirEntry, error)`

So where can you now use `fs.FS` outside of the new package? Within the standard library I’ve found the following places so far:

- You can mount a FS into a http.Server using the new `http.FS` function.
- `text/template.Template` now also has a `ParseFS` method so that you can parse templates from a file-system (same for `html/template`).

In order to test such setups there is also a new `testing/fstest` package which provides a map-based file-system with `MapFS`. 

In general, this looks really nice and might eventually make [spf13/afero](https://github.com/spf13/afero) obsolete for my use-cases 🙂 One thing that’s missing right now is any kind of write-support. There are helpers like `fs.ReadFile` and so on but the `fs.File` interface right now only offers reading methods. 

At this point there is also no context support. I’m currently assuming that write-support will make its way into the `io/fs` interfaces but context-support according to [this ticket](https://github.com/golang/go/issues/41190) sounds like a bigger issue.
