---
date: "2019-10-16T19:10:00+02:00"
title: "direnv: folder-specific environment variables"
tags:
- productivity
- direnv
- tools
---

We recently introduced a tool at my team that I've been using for a long time and that helped kept my development setup sane like only very few: [direnv][d].

direnv allows you to set environment variables automatically whenever you change into a specific directory. Let's say, you're working on something using the GitHub API for which you'd like to use a personal access token. In this case, create a `.envrc` file in your project's root directory with content similar to this:

```
export GITHUB_API_TOKEN="my-token"
```

When you enter that root directory the next time with your terminal, direnv will complain that there is a .envrc file it doesn't know:

```
direnv: error .envrc is blocked. Run `direnv allow` to approve its content.
``` 

Once you've run `direnv allow`, direnv will load the content of the .envrc file into a Bash-subshell and expose all the variables set in there to the "host" shell. This has the additional benefit that you can have complete shell functions inside that file and they *should work*.

direnv also keeps track what variables are managed by it. If you later on move out of that project again, it will unload that `GITHUB_API_TOKEN` variable again!

The setup is also extremely simple. The following two commands are all I need on macOS with homebrew and ZSH:

``` 
$ brew install direnv
$ echo 'eval "$(direnv hook zsh)"' >> $HOME/.zshrc
```

Give it a try 😁

[d]: https://direnv.nt