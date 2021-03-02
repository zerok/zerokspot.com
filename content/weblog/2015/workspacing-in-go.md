---
date: '2015-07-09T19:40:56+02:00'
language: en
tags:
- golang
title: Workspacing in Go with gb and wgo
---


Last week I was writing a small server application that I wanted to use as a
playground for finally giving [gb][] a real try. I've heard great things about
it and given that my application was about to become a consumer of multiple
libraries rather than a library itself it kind of sounded like a good idea.

I've used [godep][] previously but didn't like the idea of import path rewriting
hence the check for alternatives.

------------

gb makes quite a few things about working on projects extremely easy. Out of the
box it will build the whole project and install the resulting binary into a
clean `bin` folder when running `gb build` no matter where. When combined with
the vendor plugin it also solves dependencies quite nicely.

Sadly, I also ran into a few problems when wanting to integrate it into
[go-mode][] for Emacs. At first I thought setting the GOPATH accordingly to
reflect the distinction between the project's source code and dependencies
should do the trick:

```sh
$ export GOPATH=$PROJECTROOT:$PROJECTROOT/vendor
```

Or do the equivalent in Emacs:

```elisp
(defun zerok/setup-gb-gopath ()
  (interactive)
  (make-local-variable 'process-environment)
  (let ((srcPath (_zerok/get-gb-src-folder buffer-file-name)))
    (when srcPath
      (let* ((projectPath (string-remove-suffix "/" (file-name-directory srcPath)))
             (vendorPath (string-remove-suffix "/" (concat projectPath "/vendor")))
             (gopath (concat vendorPath ":" projectPath)))
        (message "Updating GOPATH to %s" gopath)
        (setenv "GOPATH" gopath)))))
(add-hook 'go-mode-hook 'zerok/setup-gb-gopath)

(defun _zerok/get-gb-src-folder (path)
  (let ((parent (directory-file-name (file-name-directory path)))
        (basename (file-name-nondirectory path)))
    (cond ((equal "src" basename)
           (string-remove-suffix "/" path))
          ((equal "/" parent)
           nil)
          (t
           (_zerok/get-gb-src-folder parent)))))
```

This helped with basically everything I was using from go-mode (goimports, ...)
except for the autocompletion support that is built around [gocode][]. After a
lot of debugging while waiting for one or the other plane to take me home from
Poznan I noticed the issue here being that gb places all dependency binaries not
in `$PROJECTROOT/vendor/pkg` but in `$PROJECTROOT/pkg`. Something that seems to
confuse either the go-tool or gocode itself quite a bit.

To get around this I wrote a [little helper script][1] (this repo also contains
the elisp snippet from above) that I'd from then on use right after changing
anything about my project's dependencies. What it does is it copies the compiled
binaries around into vendor that would make it look more like a `GOPATH` than it
actually is. Admittedly quite a hackey approach but it seems to work.

While talking with a couple of folks in the Gophers community on Slack John
Asmuth mentioned his own workspace tool for Go called [wgo][] which surprisingly
has a folder structure nearly identical to gb's but treats the vendor folder as
its own GOPATH segment. As it basically just wraps the go-tool, binaries will
then be placed in `vendor/pkg` which solves my integration issue with gocode.

The downside is that wgo offers less convenience. For instance to build my
little tool I had to do this:

```sh
$ cd src/cmd/myapp && wgo install
```

instead of just doing

```sh
$ gb build
```

On the other hand, gb doesn't yet support an easy path for [cross-platform
compilation][2], which was also a big requirement for me here.

So right now wgo looks like the better fit for me. It lacks some of the
comfort-features of gb, but it tries to keep everything as compatible to a
multi-GOPATH-setup as possible. That, in turn, makes integrating it with other
tools of the go-toolchain a lot easier 😊

That being said, I really hope that integration-story of gb improves, as it
looks like a great tool and I'd really love to give it a try again in the future!

[wgo]: https://github.com/skelterjohn/wgo
[gb]: http://getgb.io/
[gocode]: https://github.com/nsf/gocode
[go-mode]: https://github.com/dominikh/go-mode.el
[godep]: https://github.com/tools/godep
[1]: https://github.com/zerok/emacs-golang-gb
[2]: https://github.com/constabulary/gb/issues/31
