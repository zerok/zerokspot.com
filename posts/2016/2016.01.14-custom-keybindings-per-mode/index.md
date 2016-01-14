# TIL: Custom keybindings per mode

Emacs makes it extremely easy to define custom key bindings depending on what
major mode you're in. A couple of days ago I wrote myself a little helper for
the awesome [jedi][] plugin that allows me to select which virtualenv should be
used for completions.

Since this is only relevant for Python code, I wanted to have that binding to be
only available while in python-mode. As I wrote above, getting a local
keybinding is very easy using the `define-key` function:

```
(define-key python-mode-map (kbd "C-c C-v") 'jediselect)
```

But I'm using [evil-mode][] and so I actually want something that feels more
natural in that environment. For command like that I prefer leader
keys. Luckily, [evil-leader][] also has a way to create mode-local key bindings
using `evil-leader/set-key-for-mode`:

```
(evil-leader/set-key-for-mode 'python-mode "v" 'jediselect)
```

Note that this won't work for buffers where python-mode is already active. You
will either have to reopen these buffers or re-enable the major mode (didn't try
the latter option).

[evil-mode]: https://bitbucket.org/lyro/evil/wiki/Home
[jedi]: https://github.com/tkf/emacs-jedi
[evil-leader]: https://github.com/cofi/evil-leader
