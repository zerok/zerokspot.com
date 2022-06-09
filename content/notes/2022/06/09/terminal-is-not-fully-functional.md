---
title: "terminal-is-not-fully-functional"
date: "2022-06-09T20:30:37+0200"
tags:
- macos
- homebrew
- tmux
- git
---

After updating tmux and zsh last week, I started to see some weird errors like the one below:

```
❯ git log
WARNING: terminal is not fully functional
Press RETURN to continue
```

This [seems](https://github.com/tmux/tmux/issues/2262) to be caused by `/usr/bin/less` using an old ncurses version which lacks some color profiles. To fix this I've now installed less from Homebrew and configured Git to use that instead

```sh
brew install less
git config --global core.pager "/usr/local/bin/less -FX"
```