---
date: '2015-10-21T19:32:07+02:00'
language: en
tags:
- cookiecutter
- yeoman
- project-templates
title: Project templating with Cookiecutter and Yeoman
---


Two weeks ago I finally started learning a bit about [Redux][r] in combination
with React. Since I liked it so much I decided last weekend to put it into a
template that I'd from then on use for every project I'd start with this stack
and evolve it depending on what I've learnt.

There are tons of tools out there for handling project scaffolds. Rails, Django,
they all have their own, but I wanted to use one tool I could use for basically
every stack I was working with (which is actually quite a few).

I also wanted to use this first template as a proofing ground for two tools I've
heard a lot of and partially used before: [Yeoman][yo] and [Cookiecutter][cc].

---------

First of:

## Cookiecutter

Getting started here is extremely easy. Create a folder for our template. Then
create a `cookiecutter.json` file alongside a `{{ cookiecutter.repo_name }}`
folder and fill the latter with all the things you want to see in your file name
project. Every file in there is treated as a Jinja template with the context
variables coming out of the `cookiecutter.json` file.

If you don't need anything else, you don't have to write a single line of
code. That's it!

If you do want more, you can also create hook scripts that are executed either
before or after the copying of the generated files and you can also create
additional context variables to be used with the templates.


## Yeoman

If you need even more, Yeoman (or Yo) might be something for you. Here you have
to write code right from the get-go as templates are just another feature. Yo's
process consists of multiple phases called "priorities" (initializing,
prompting, configuring, default, writing, conflicts, install, and end) which
house JavaScript functions in which you define what your generator should
do. Create files, install packages, ask some weather API for data to put into
your generated files, ...

Contrary to what the website might look like, it's not only useful for web
applications but thanks to everything just basically being a bunch of JavaScript
functions you can do more or less whatever you want with a generator. No one
will stop you from create a template for your favorite Go project structure with
it 😉

It's definitely harder to get started, though, as everything has to be done
explicitly rather than implicitly.


## Packaging and distribution

Another big difference between these two is how they are packaged and
distributed. Cookiecutter is just a collection of JSON and Jinja template files
and not a full Python package. This means that there is no easy way to have hook
scripts have additional dependencies and things like that. If you want to use
[requests](http://docs.python-requests.org/en/latest/) in your post-project
hook, you have to install it manually.

Yo generators on the other hand are just plain node modules with a package.json
and with whatever dependencies you like.

Yo expects generators to be installed globally, so the easiest way for
distribution is by publishing your generator to npm. The downside of this is
that you can't really have "your flavor" of a generator without giving it a
somewhat awkward name. You obviously can just `npm link` your project (as has
also been discussed in the [yeoman issue track][yo348]) but that's more a
workaround, IMHO.

Cookiecutters, by contrast, simply takes local folders, git and mercurial URLs
as an argument and works based on that. This way I can have my own
cookiecutter-react-redux template without interfering with anybody else's.


## Which to pick?

Honestly? No idea. I've now written my project template in both and will, for
the time being, probably maintain both. Given the naming issue only the
cookiecutter one has made it onto [Github][ccr] yet.

Yo is far more powerful, though, so that might give it an edge for some of the
things I might end up needing templates for in the long run, but for now,
Cookiecutter's simplicity is a big advantage.

[cc]: http://cookiecutter.readthedocs.org/en/latest/
[yo]: http://yeoman.io/
[yo348]: https://github.com/yeoman/generator/issues/348
[r]: https://github.com/rackt/redux/
[ccr]: https://github.com/zerok/cookiecutter-react-redux
