---
date: 2016-12-16T16:53:15+01:00
tags:
- emacs
title: Automatic mode selection in Emacs
---

Depending on which file you're opening in Emacs, you normally want to have
different behaviour or different highlighting patterns to be loaded. This is
something where you'd normally use the `auto-mode-alist` variable to bind a
file name pattern to a major mode:

```
(add-to-list 'auto-mode-alist '("\\.js\\" . js2-mode))
```

Now I was in the situation where I wanted to use js2-mode pretty much everywhere
close to JavaScript except for files that contained JSX. There I prefer web-mode
or js2-jsx-mode for now. Turns out, the whole mode-selection process is
extremely customizable and is explained in detail [here][].

At first I thought I might add file-local variables to each JSX file. That's not
really a good option, though, as it would force my mode-selection onto basically
everyone else working on such files:

```
// -*- mode: web -*-
```

For what I had in mind, it should be enough to just take a look at the first
line of a JavaScript file. If it contains something like `import React ...`
web-mode should be used. That's pretty much what the magic-mode-alist allows.

```
(add-to-list 'magic-mode-alist '("^import React" . js2-jsx-mode))
```

That's obviously not really bulletproof (as that import might be in a complete
different line or I'd be dealing with an CommonJS import), but for the most part
it should do what I want 😊

During playing with these settings I finally also got a chance to give Wilfred
Hughes excellent [refine][] command a try. It made modifying the various
something rather long alists extremely easy!

I also learnt about the `normal-mode` command which allow you to reset the
mode-selection in a buffer back to what would be picked automatically.

[here]: http://www.gnu.org/software/emacs/manual/html_node/emacs/Choosing-Modes.html#Choosing-Modes
[refine]: https://github.com/Wilfred/refine
