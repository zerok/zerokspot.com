---
date: '2015-03-24T19:45:10+01:00'
language: en
tags:
- python
- pip
title: Dependency Trees in Python
---


One of the nice features when working with npm for io.js/node.js is the `npm list` command.
It provides you a simple tree view of all the dependencies your
package has and their dependencies and so on.

```
$ npm list
packagename@0.1.0 /Users/zerok/code/packagename
├── async@0.9.0
├─┬ chai@2.1.2
│ ├── assertion-error@1.0.0
│ └─┬ deep-eql@0.1.3
│   └── type-detect@0.1.1
├─┬ chalk@1.0.0
│ ├── ansi-styles@2.0.1
│ ├── escape-string-regexp@1.0.3
│ ├─┬ has-ansi@1.0.3
│ │ ├── ansi-regex@1.1.1
│ │ └── get-stdin@4.0.1
│ ├─┬ strip-ansi@2.0.1
│ │ └── ansi-regex@1.1.1
│ └── supports-color@1.3.0
├── es6-collections@0.5.1
├── es6-promise@2.0.1
...
```

This is very useful to find out after the fact what packages were installed and
what dependencies your project really has. If for nothing else this might help
with learning about low-level libraries you might want to use yourself in the
future...

Sadly, pip doesn't have something like that out of the box. But here something
like that would be especially useful simply because pip/setuptools/Python right
now doesn't support actual hierarchical dependency installations (similar to
npm) nor an obvious dependency resolution method. Running `pip freeze` will only
show you the final package list as installed after all the fun happened.

-------------

With "the fun" I mean that right now there is only one "final" level of
dependencies in each project as all are installed into one folder and can be
accessed directly. What about situations where you use library "a" that depends on
"a.a" at version 1 while you also depend on "b" that requires version 2 of "a.a"?
[issue988][2] has some nice details on situations like this. Basically top-level
dependencies are the most important ones and everything else is resolved more or
less in a breadth-first-way.

Luckily, there is already something similar to `npm list` out there called
[pipdeptree][1] which helps with finding possible dependency conflicts and
displaying an actual dependency hierarchy.

This would be its output for one of my smaller Django projects:

```
$ pipdeptree
Warning!!! Possible confusing dependencies found:
* django-cms==3.0.12 -> django-classy-tags [required: >=0.5, installed: 0.5.2]
  django-sekizai==0.8.1 -> django-classy-tags [required: >=0.3.1, installed: 0.5.2]
* django-debug-toolbar==1.2.2 -> Django [required: >=1.4.2, installed: 1.7.7]
  django-cms==3.0.12 -> Django [required: >=1.4, installed: 1.7.7]
  django-mptt==0.6.1 -> Django [required: >=1.4.2, installed: 1.7.7]
  django-filer==0.9.9a1.dev1 -> Django [required: >=1.4, installed: 1.7.7]
  easy-thumbnails==2.2 -> Django [required: >=1.4.2, installed: 1.7.7]
  django-classy-tags==0.5.2 -> Django [required: >1.2, installed: 1.7.7]
  cmsplugin-filer==0.10.1 -> Django [required: >=1.4, installed: 1.7.7]
* django-cms==3.0.12 -> django-sekizai [required: >=0.7, installed: 0.8.1]
  cmsplugin-filer==0.10.1 -> django-sekizai [required: >=0.4.2, installed: 0.8.1]
------------------------------------------------------------------------
cmsplugin-filer==0.10.1
  - Django [required: >=1.4, installed: 1.7.7]
  - django-cms [required: >=3.0, installed: 3.0.12]
    - Django [required: >=1.4, installed: 1.7.7]
    - South [required: >=0.7.2, installed: 1.0.2]
    ...
  - django-sekizai [required: >=0.4.2, installed: 0.8.1]
    - django-classy-tags [required: >=0.3.1, installed: 0.5.2]
      - Django [required: >1.2, installed: 1.7.7]
  - easy-thumbnails [required: >=1.0, installed: 2.2]
    - Django [required: >=1.4.2, installed: 1.7.7]
    - Pillow [installed: 2.7.0]
  ...
djangocms-style==1.5
djangocms-text-ckeditor==2.4.3
...
```

As you can see, pipdeptree warns you when you have multiple dependencies where
the versions don't exactly match. At this point it shows too many warnings for
my taste as there are actually no conflicts up there. That being said, it's a
first step.

If you want to learn more about this whole issue, I'd recommend taking a look at
[issue988][2] as well as this [great post][3] by Vincent Driessen about his and
Bruno Renié's work on [pip-tools][4].

Part of this effort is the `pip-compile` command which will hopefully soon help
resolve versioning conflicts right during creation of the
requirements.txt. Can't wait for more tools on that front!

[1]: https://github.com/naiquevin/pipdeptree
[2]: https://github.com/pypa/pip/issues/988
[3]: http://nvie.com/posts/better-package-management/
[4]: https://github.com/nvie/pip-tools
