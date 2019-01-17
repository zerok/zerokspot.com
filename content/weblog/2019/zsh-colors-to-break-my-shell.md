---
title: "ZSH colors to break my shell"
date: 2019-01-17T20:28:29+01:00
tags:
- zsh
- prompt
---

For the last couple of months I've had a shell prompt with two lines. One for displaying things like the task I'm currently booking time on, one for the PWD and the actual command input.

It had one problem, though: Whenever I had a command that filled the whole line or on some edit commands the whole prompt fell apart: The input-line no longer showed the complete command, sometimes only a couple of characters anymore. In addition to that, using Ctrl+r for searching the history added a weird space in front of the resulting command.

[![asciicast](https://asciinema.org/a/Vv5IAtJAOLcEb8MDJwZoMUZBE.svg)](https://asciinema.org/a/Vv5IAtJAOLcEb8MDJwZoMUZBE)

After a bit of testing I noticed that the problem was somehow related to the variables provided by the `colors` module, or actually my usage of them. I had something like this:

```
autoload -Uz colors
colors
export PROMPT="$fg[blue]\$$reset_color "
```


Turns out, inside a prompt variables have to be wrapped inside `%{...%}` stanzas. 

> %{...%}
>
> Include a string as a literal escape sequence. [...]

... according to [zsh manual](http://zsh.sourceforge.net/Doc/Release/Prompt-Expansion.html#Visual-effects). So the following prompt works as expected:

```
autoload -Uz colors
colors
export PROMPT="%{$fg[blue]%}\$%{$reset_color%} "
```



