---
date: '2014-12-13T23:13:33+01:00'
language: en
tags:
- git
- development
title: Getting to know pre-commit
---

Getting your whole team to use the same set of client-side hook scripts when
working with Git is not a simple task. That's why I was really grateful when I
stumbled upon the [pre-commit][pc] project by yelp. This project offers some
tools for managing your pre-commit hooks and facilitates sharing them between
multiple projects and developers.

## Setup

The core idea is that you have all your hooks configured within the project's
root directory in a file called `.pre-commit-config.yaml` and then simply
execute `pre-commit install` once to have pre-commit inject itself into
`.git/hooks/pre-commit`. From now on, whenever you are about to make a commit,
pre-commit will launch and execute your hook scripts if appropriate.

The scripts themselves are hosted in some other repository that you specify in
that config file. To illustrate this, here a short example:

```
-   repo: ssh://git@our-repo-server.com/precommit-hooks.git
    sha: 1bf3eaef56059e168aac55393a7494cac727ebcd
    hooks:
        - id: flow-branches
        - id: grunt-test
```

This is the configuration I'm currently using for one of our work projects (as
an experiment so far). It specifies that it should execute the hookscripts
`flow-branches` and `grunt-test` from our precommit-hooks repository. The
scripts themselves are versioned with the commit ID, so updates to the scripts
repository don't break anything.

Other team members then just have to install pre-commit and execute its
`install` sub-command and have all the configured hook scripts enabled.

```
$ pip install pre-commit
$ cd /path/to/project
$ pre-commit install
```

This repository referred to in the configuration file has to contain a
`hooks.yaml` file where all the exposed hook scripts are documented. The file in
this example would look something like this:

```
-   id: flow-branches
    name: Flow Branches
    language: python
    entry: flow-branches
    files: .*
-   id: grunt-test
    name: Grunt Test
    language: python
    entry: grunt-test
    files: .*
```

The `id` signals what the script should be called in your project's
`.pre-commit-config.yaml` while the `entry` property points to the executable
script itself.


## Multi-language support

As this example indicates, you can write your hook scripts in various languages.
For Python, for instance, this will create a virtualenv and install all the
dependencies specified in the `setup.py` of our hooks repository for you. This
is for me the single best feature of this project. You don't have to tell your
team-mates what other stuff they have to install before they can use your new
fancy hook or explain virtualenvs to them if they are not Python coders but just
tell them once to install pre-commit and run `pre-commit install` within the
repository. All the rest is done by the tool itself.

If you want to write your scripts using Node.JS you can do so as well and it
will also handle all the dependencies mentioned in a `package.json` file for you.


## Only run when necessary

The example above also include a field `files` which is also worth mentioning.
With this you can specify that a hook script should only be executed if the
commit-to-be-made contains files with a specific name pattern. This way for
instance a hook for running jshint on JavaScript files would not be executed if
the commit affect any JavaScript files.

You can override this pattern within the `.pre-commit-config.yaml` if you want
to.


## The downsides

Sadly, at this point pre-commit has two major downsides which mean that I'm
still looking for alternative implementations:

1. As the name implies it is limited to handling pre-commit hooks. But what
   about when you want to also check the commit message? For that you'd have to
   install a script into the `commit-msg` hook.

2. Internally, pre-commit using some shell utilities like xargs which means it
   doesn't support Windows. There is [ticket][i159] for that but that's about it.

I've already looked at [overcommit][oc] by Causes but this doesn't seem to handle the
sharing of scripts between projects that well and also is limited to operate on
Ruby hook scripts. Especially the latter makes adapting it harder because our
current language stack doesn't include Ruby (except for Compass/SASS). From what
I've seen it also doesn't do anything about managing dependencies of hook
scripts.

[i159]: https://github.com/pre-commit/pre-commit/issues/159

[oc]: https://github.com/causes/overcommit

[pc]: http://pre-commit.com
