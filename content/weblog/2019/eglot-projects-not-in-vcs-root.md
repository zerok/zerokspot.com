---
title: LSP-support with eglot in nested projects
date: "2019-11-27T22:00:00+01:00"
tags:
- emacs
- golang
---

After switching from lsp-mode to [eglot][] recently, I ran into a
weird issue today: When opening a work-project eglot+gopls basically
bailed on me. All I got was this error message:

```
Your workspace root is /path/monorepo, but your module root is /path/monorepo/project.
```

This message was emitted by gopls which is reacting to an `initialize`
call with a [root path][rp] that doesn't match the information it can
find inside the local `go.mod` file.

OK, so where is that "root path" coming from? Internally, eglot uses
the function [`project-current`][pc] which is part of the `project.el`
library. Sadly, this function, out of the box, seems to only support
detecting project root folders by looking at various VCS metadata. In
my case though, `/path/monorepo` is the one that contains the `.git`
folder while `/path/monorepo/project` doesn't.

As a quick fix I thought, I could just add a `.projectile` file (since
I'm using [projectile][] anyway) to the project directory and let
`(project-current)` detect that. project.el offers an extension point
for detecting project folders with the `project-find-functions`
hook. Luckily, Louis Roché seems to have faced a similar issue and
provided that little snippet in his [emacs.d repository][e]:

```
(defun my-projectile-project-find-function (dir)
  (let ((root (projectile-project-root dir)))
    (and root (cons 'transient root))))

(projectile-mode t)

(with-eval-after-load 'project
  (add-to-list 'project-find-functions 'my-projectile-project-find-function))
```

If I now run `(project-current)` in a file inside
`/path/monorepo/project`, I get the correct root path and gopls is
finally happy! Problem solved!

[e]: https://github.com/Khady/emacs.d/blob/52b99aa850ef354a49227b9660dca4184ef01e87/config.org
[rp]: https://github.com/joaotavora/eglot/blob/02f71fdca9b06e8eae123db80527a1cb0fcc62bd/eglot.el#L795
[eglot]: https://github.com/joaotavora/eglot
[pc]: https://github.com/emacs-mirror/emacs/blob/d0e2a341dd9a9a365fd311748df024ecb25b70ec/lisp/progmodes/project.el#L101
[projectile]: https://github.com/bbatsov/projectile
