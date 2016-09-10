---
date: '2014-05-31T15:09:55+02:00'
language: en
tags:
- golang
title: Integrating godep into my toolchain
---


Coming from Python and more recently NodeJS I've come to appreciate tools like
[npm][] that allow you to easily manage all your project's dependencies into
specific vendor directories or environments to prevent issues where one project
requires one version of a library while another might need something different.

In Go this becomes especially important because the whole [versioning-
issue][vers] has not been solved here yet. So "vendoring" is actually the
[recommended approach][govend] and rather simple to pull off by specifying a
GOPATH for each project. Simple but tedious. That's pretty much where [godep][]
comes into play:

`godep` is a small tool that you can wrap all your `go` calls around to work
on a project-specific `$GOPATH`. Since I only started working with it a couple of
days ago, I'm still not completely sure if I'll use it for every project and
how deeply I will integrate it into my toolchain. That being said, this is how
far I've got.

Oh, and I know about [gom][], but simply haven't had yet the time to look into
it yet. godep simply came first for me :)

-----------------

## How to integrate it with GoSublime

For this you will have to have a SublimeText project set up for every project
where you want to use godep. The settings here are rather minimal since
GoSublime allows you to change the environment variables used for it's go
subprocess:

```
{
    "settings": {
        "GoSublime": {
            "env": {
                "GOPATH": "$HOME/path/to/your/project/Godeps/_workspace"
            }
        }
    },
    "folders":
    [
        {
            "follow_symlinks": true,
            "path": "."
        }
    ]
}
```

As a quick helper you can get the GOPATH godep would use internally by running
`godep path` in your project's directory. I'm not yet all that happy about
having to put an absolute path in here, but given Go's tendency to absolute
import paths it feels at least consistent :)


## How to integrate it with GoConvey

For testing I really like [goconvey][] which provides a webinterface to the
current state of coverage and test execution. That webinterface is started by a
commandline tool that let's you specify which go binary should be used for
running the tests and determining the coverage. This is also where I thought I
might be able to inject godep into the process.

At first I thought I could just set the `-gobin` parameter to something like
`godep go` and be done with it. Sadly, this doesn't work because the flag really
requires *one* file. But that's what short shell scripts are good for:

```
#!/bin/sh
exec godep go $@
```

I put this onto my `$PATH` and all of a sudden all I have to do is start
goconvey like this:

```
$ goconvey -gobin='godepgo'
```

This seems to work pretty much as expected so far. And if it stays that way
I will most likely roll that setup out to other projects.

[gs]: https://github.com/DisposaBoy/GoSublime
[govend]: http://golang.org/doc/faq#get_version
[godep]: https://github.com/tools/godep
[goconvey]: http://goconvey.co/
[npm]: https://www.npmjs.org/
[gom]: https://github.com/mattn/gom
[vers]: http://www.goinggo.net/2014/01/go-package-management-for-2014.html