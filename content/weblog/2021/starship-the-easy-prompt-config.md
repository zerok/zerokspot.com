---
title: "Starship: the easy prompt config"
date: "2021-04-09T20:58:00+02:00"
tags:
- tools
- 100daystooffload
- rust
incoming:
- url: https://chaos.social/@zerok/106036954835899622
---

[Starship.rs](https://starship.rs/config/#custom-commands) is a little tool
that I stumbled upon thanks to [Jeff Triplett](https://jefftriplett.com/) which
generates PROMPT variables for various shells based on a simple and easy to
read configuration file. The goal here is to have the same configuration
available independent of what shell you're using, be it ZSH, BASH, or something
else entirely.

All you have to do, is to install starship (e.g. using Homebrew if you're on
macOS) and then add it to your shell:

    $ eval "$(starship init zsh)"

Configuration then happens through `~/.config/starship.toml`. You basically
configure a set of module (e.g. the Directory module or the Git module) and
also tell starship through the `format` setting where in the prompt the output
of that module should be rendered:

    format = """
    $username\
    $hostname\
    $directory\
    $git_branch\
    $git_status\
    ${custom.kubeselect}\
    $line_break\
    $status\
    $shell\
    $character"""
    
    [gcloud]
    disabled = true
    
    [custom.kubeselect]
    when = "true"
    command = "/Users/zerok/bin/kubeselect status"
    style = "bold blue"
    format = "⛵️ [($output)]($style) "

This is, for instance, my current setup. I'm using pretty much the default
configuration with two small changes:

1.  I don't want to have my gcloud configuration show up in the prompt
2.  I use kubeselect to manage my kubectl configuration

The second part is also a good example for how you can get custom data into
your prompt. You define a [custom
command](https://starship.rs/config/#custom-commands) and tell starship when it
should be considered (in my case always), what command should be executed, and
finally how the output should be styled and formatted.

The goal of being cross-platform comes at a small price, though: If you're on
ZSH and used to have part of your prompt being on the right side (using the
`rprompt` variable) you're out of luck for now. That being said, there exists a
[draft](https://github.com/starship/starship/pull/2425) right now for adding
right-side prompts in the future.

But even without rprompt support, I really like starship. It's fast, it's easy,
and it's extensible. At this point I'm pretty close to just installing it per
default also on all my servers just so that I get a decent default prompt no
matter where I go 😄

