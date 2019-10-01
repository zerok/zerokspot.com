---
title: "Some updates about tpl"
date: 2019-10-01T18:40:36+02:00
tags:
- tpl
- sideproject
- releasenotes
---

[tpl][] is a little project I started [nearly two years ago][orig] out
of the need to dynamically generate a couple of files based on a
simple template and some secret data coming from a [Vault][] server.

Since then it has been extended to read environment and eventually
execute external shell commands and use their output during the
rendering phase. Well, but even that was now nearly 12 months ago with
not much happening therafter.

Until last week when a [coworker][kautsig] of mine discovered the tool
and started to use it. He also [opened a PR][pr1] for allowing tpl to
read the source template from stdin instead of from a file:

```
$ cat template.tpl
{{ index .Env "HOME" }}

$ cat template.tpl | tpl -
/Users/zerok
```

After I merged this, I released [tpl v2.6.0][v26] 🙂

While working on the same project that had caused [kautsig][] to
extend tpl, I felt the need for something that I love when
using [Hugo][]: [Data files][df]. Wouldn't it be nice to use JSON or
YAML data without first having to jump through some hoops (i.e. using
the system command to `cat` the file and the using some parsing logic
in the template)?

Let's say, you have a users.yaml file that you generated in some other
step containing all the users you have in your system:

```
- zerok
- admin
- superadmin
- bossofsuperadmin
```

The template should now generate a new YAML file that contains the
username but also the home-directory of that user:

```
{{ range .Data.users }}
- name: {{ . }}
  home: /home/{{ . }}
{{ end }}
```

That's pretty much what was added with [v2.7.0][v27] later that day:

```
$ tpl --data 'users=users.yaml' template.yaml.tpl

- name: zerok
  home: /home/zerok

- name: admin
  home: /home/admin

- name: superadmin
  home: /home/superadmin

- name: bossofsuperadmin
  home: /home/bossofsuperadmin
```

Let's see if it takes another 12 months for someone to come up with
another feature for tpl. I'm not sure yet but perhaps something
related to [nushell][] perhaps? 😉

[tpl]: https://github.com/zerok/tpl
[pr1]: https://github.com/zerok/tpl/pull/16
[v26]: https://github.com/zerok/tpl/releases/tag/v2.6.0
[v27]: https://github.com/zerok/tpl/releases/tag/v2.7.0
[kautsig]: https://github.com/kautsig
[orig]: https://zerokspot.com/weblog/2017/10/29/templating-all-the-things/
[vault]: https://www.vaultproject.io/
[hugo]: https://gohugo.io
[df]: https://gohugo.io/templates/data-templates/
[nushell]: http://www.nushell.sh/

