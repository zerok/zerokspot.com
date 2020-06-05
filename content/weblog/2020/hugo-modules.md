---
title: Hugo modules
date: "2020-06-05T07:43:06+02:00"
tags:
- hugo
- tutorial
- modules
- 100daystooffload
---

Last July, [Hugo](https://gohugo.io/) [introduced](https://gohugo.io/news/0.56.0-relnotes/) a new mechanism for re-using parts of your website setup: [Modules](https://gohugo.io/hugo-modules/). Based on how Go (the programming language that Hugo is written in) handles libraries Hugo Modules allows you not only to load things like layouts but also content, i18n, and archetypes!

Since zerokspot.com was started a long time before that feature was released  I didn’t “organically” had the opportunity to work with it yet. In order to get around this, I thought I’d start a little experimental theme implemented using Hugo Modules.

Over the course of this little walkthrough I’m going to work on a mostly fictional project located in `$HOME/src/github.com/hugo-blog-quickstart` which is a Hugo site. This site should now use a theme/module located in `$HOME/src/github.com/zerok/hugo-blog-quickstart-theme`. The URL-like paths are not an accident. You will perhaps find the code produced during this article on the respective GitHub pages in the future.

So let’s get started!

## Let’s start a module

A Hugo Module is not much more than a folder that may or may not be managed using Git (or other versioning systems) that can be reached by the [module system provided by the Go programming language](https://github.com/golang/go/wiki/Modules). For our little example, we are going to create a new module inside a repository on GitHub: 

[https://github.com/zerok/hugo-blog-quickstart-theme](https://github.com/zerok/hugo-blog-quickstart-theme)

In order to work on it, I’ve put the content of this repository into `$HOME/src/github.com/zerok/hugo-blog-quickstart-theme` (aka `$THEME_ROOT`) and the blog itself into `$HOME/src/github.com/zerok/hugo-blog-quickstart` (aka `$SITE_ROOT`).

For now, let’s assume that the theme folder is actually empty. The first thing we have to do is to initialise it:

	> cd ${THEME_ROOT}
	> hugo mod init github.com/zerok/hugo-blog-quickstart-theme

This will just create a couple of folders and a `go.mod` file:

	drwxr-xr-x  - zerok  4 Jun 14:21 .
	.rw-r--r-- 68 zerok  4 Jun 14:21 ├── go.mod
	drwxr-xr-x  - zerok  4 Jun 14:20 └── resources
	drwxr-xr-x  - zerok  4 Jun 14:20    └── _gen
	drwxr-xr-x  - zerok  4 Jun 14:20       ├── assets
	drwxr-xr-x  - zerok  4 Jun 14:20       └── images

## Connect site and module

In order for Hugo to know about the module when building the site, we will have to update the `$SITE_ROOT/config.toml` and add these three lines:

	[module]
	    [[module.imports]]
	    path = "github.com/zerok/hugo-blog-quickstart-theme"

Running `hugo mod get -u` will now try downloading the module from GitHub. Since we want to work on our local version of the module, though, we will have to do one more thing:

Go module definitions can also declare that a dependency (in our case our theme) [can be found somewhere other than the URL referenced](https://github.com/golang/go/wiki/Modules#when-should-i-use-the-replace-directive) by the path. This is done inside the `$SITE_ROOT/go.mod` file using a `replace` line. Since the theme is located on our filesystem in `../hugo-blog-quickstart-theme`, we need to add the following line at the bottom of the `$SITE_ROOT/go.mod` file:

	replace github.com/zerok/hugo-blog-quickstart-theme => ../hugo-blog-quickstart-theme

## Mounting

By default, if the module contains, for instance, a `content/_index.md` and the site does not, then Hugo will consider the content provided by the module. The same applies to the other core Hugo folders like `static`, `archetypes`, `layouts`, and so on. 

But let’s say our theme has some fancy JavaScript files like `static/js/main.js` and the site already has its own `js/main.js` inside the static folder. This is where mounts come in.

We can tell Hugo that the module’s JavaScript file should be made available in `static/fancymodule/js/main.js`. For this, we will have to update the `[module]` section inside `$SITE_ROOT/config.toml`:

	[module]
	    [[module.imports]]
	        path = "github.com/zerok/hugo-blog-quickstart-theme"
	        [[module.imports.mounts]]
	        source = "static/js"
	        target = "static/fancymodule/js"

That’s pretty much it regarding features of Hugo Modules. They are a way to merge multiple directory trees into a single one right before Hugo tries to build everything. There is just one more thing that I wanted to mention:

## Embedded modules

We had originally started under the assumption that we’d have the theme in a different repository than the site. While this is great for re-use, you don’t necessarily have to take this route!

Let’s say, our theme was actually located in `$SITE_ROOT/themes/theme`, then we could also this location with the following module configuration:

	[module]
	    [[module.imports]]
	    path = "theme"

Hugo will look inside the `$SITE_ROOT/themes` folder for a matching folder. Sadly, it seems like you cannot use any other folder structure here without jumping through some hoops, but especially for the use case of migrating from classic layouts and themes to modules this should do fine 🙂
