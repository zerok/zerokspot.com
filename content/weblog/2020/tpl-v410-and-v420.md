---
title: tpl 4.1.0 and 4.2.0
date: "2020-10-15T18:10:01+02:00"
tags:
- tpl
- utility
- opensource
- release
---

It took quite some time but today I made two new releases for [tpl](https://github.com/zerok/tpl/) with new features that I personally needed and that are hopefully also useful to you 🙂

You can find  v4.2.0 on [https://github.com/zerok/tpl/releases/tag/v4.2.0](https://github.com/zerok/tpl/releases/tag/v4.2.0).

## Reading files

The first new change is a new addition to the `.FS` element: Previously, you could only use `{{ .Shell.Output "cat path/to/file" }}` in order to include the content of a file inside the generated output. This required the `--insecure` flag and was therefore a bit tedious. To simplify that you can now use `.FS.ReadFile FILEPATH`:

	{{ .FS.ReadFile "path/to/file" }}

## Output buffering

Another change that I made this afternoon affects how tpl generates the final output. Previously, output was directly streamed which resulted in partial files being generated in case an error happened. This was especially annoying when used in combination with GNU Make. Let’s say you have a `targetfile.tpl` that contains a bug and the following rule:

	targetfile: targetfile.tpl
	  tpl targetfile.tpl > targetfile

Even as tpl fail, `targetfile` will be created and therefore returning that target won’t work until the file is deleted.

To fix this, output is now buffered and only printed if no error occurred. There is now also a new `--output FILE` flag which only touches/creates the output file if no error happened. This way, if you have following Make rule and tpl fails, the target file won’t be modified:

	targetfile: targetfile.tpl
	  tpl --output targetfile targetfile.tpl
