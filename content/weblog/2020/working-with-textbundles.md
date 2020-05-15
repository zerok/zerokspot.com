---
title: Working with TextBundles
date: "2020-05-15T17:27:40+02:00"
tags:
- blog
- textbundle
- writing
- bundle
---

While evaluating [Ulysses](https://ulysses.app/) for writing future blog posts I stumbled upon the [TextBundle](http://textbundle.org/) file format which can be created by Ulysses and other editors like Bear. 

To quote the [specification](http://textbundle.org/spec/):

> The purpose of the TextBundle file format is to simplify the exchange of various plain text files together with additional images between sandboxed applications.

Let’s take this post as an example and export it as TextBundle. The result is a single file by the name of `Working with TextBundles.textpack` which is a compressed TextBundle:

	unzip ./Working\ with\ TextBundles.textpack
	Archive:  ./Working with TextBundles.textpack
	   creating: Content.textbundle/
	   creating: Content.textbundle/assets/
	  inflating: Content.textbundle/assets/textbundle-logo.png
	  inflating: Content.textbundle/text.md
	  inflating: Content.textbundle/info.json

Now, a TextBundle contains three main components:
1. The text itself stored in `text.*`.
2. Metadata about the bundle stored in `info.json`.
3. Assets used within the text stored as separate files inside the `assets/` folder.
Metadata is mostly optional except for one thing: It has to specify the format-version of the bundle. At this point, version 2 is the latest and also used by Ulysses:

	{
	  "creatorURL" : "file:\/\/\/Applications\/UlyssesMac.app\/",
	  "transient" : false,
	  "type" : "net.daringfireball.markdown",
	  "creatorIdentifier" : "com.ulyssesapp.mac",
	  "version" : 2
	}

`type` is optional but defaults to `net.daringfireball.markdown` for backwards-compatibility. Theoretically, a bundle could contain multiple text files with different extensions but handling of such a case seems not to be specified by spec.

Since I might want to integrate Ulysses into my blogging workflow and TextBundle sounds like a decent exchange format, I started playing around with implementing a little [Go library for reading TextPack files](https://github.com/zerok/textbundle-go):

	> r, _ := textbundle.OpenReader("Working with TextBundles.textpack")
	> fmt.Println(string(r.Text))
	
	# Working with TextBundles
	...
	
	> for _, a := range r.Assets {
	>   fmt.Println(a.Name)
	> }
	textbundle-logo.png

As a little experiment I’ve also created a little `import-textbundle` command for my blog that uses this library 🙂
