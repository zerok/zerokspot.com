---
title: "Go 1.16: Retracting module versions"
tags:
- golang
- releasemanagement
- 100daystooffload
date: "2021-02-03T19:00:00+01:00"
incoming:
- url: https://chaos.social/@zerok/105668658378606006
---

Some time ago thanks to a blogpost (sorry, cannot remember anymore which one 😩) I stumbled upon another feature in Go 1.16 that I’m looking forward to: As a module author you can now mark releases as not being recommended for use. This is interesting for situations where you know that there was a really bad bug in a certain version or some security flaws in a whole range of releases up to a certain point.

The [release notes](https://tip.golang.org/doc/go1.16) only mention this feature with these two sentences:

> retract directives may now be used in a go.mod file to indicate that certain published versions of the module should not be used by other modules. A module author may retract a version after a severe problem is discovered or if the version was published unintentionally. 

Sadly, I couldn’t find anything about that (yet) inside the official documentation itself but then I stumbled upon [this comment by Jay Conrod](https://github.com/golang/go/issues/24031#issuecomment-597263309) in which he outlines the way this feature has been implemented. The following paragraphs are based purely on that since I haven’t had the time to use that feature on an actual dependency of mine.

So how do you use retractions? As an author you add one more multiple `retract ...` directives to your project’s `go.mod` file and make that available for the “@latest” version of your library.

That directive can be added in various forms:

	// Really bad by introduced in this patch
	retract v1.2.3
	
	retract (
	    v1.3.0 // Broken go.mod file
	    [v1.3.1, v1.3.4] // Security issue introduced
	)

Here we have (1) a single `retract` directive in a dedicated line and (2) two such directives combined with a syntax that looks similar to the usual `require (...)` directive.

`[v1.3.1, v1.3.4]` here indicates that all the versions from v1.3.1 up to *and including* v1.3.4 have been retracted. According to [this comment](https://github.com/golang/go/issues/24031#issuecomment-747585995) open intervals or partially open ones might be supported in the future.

When a user now wants to use your library, you can use `go get` with the  “@latest”, “@upgrade”, and “@patch” versions and these will respect and therefore ignore retracted versions.

At this point I’m curious if there will also be some standardisation around the comment for each retraction. For instance as a user I might be especially interested if something was retracted due to security issues found for those versions. If those comments could then even link to CVEs that would even be more awesome. Let’s see 😀
