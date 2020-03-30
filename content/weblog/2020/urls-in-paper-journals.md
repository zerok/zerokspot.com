---
title: "Dealing with URLs in paper journals"
date: "2020-03-30T19:42:55+0200"
tags:
- journaling
- bulletjournal
---

Journaling in paper notebooks was all great until I wanted to also include links in entries. It's simply not fun to write out long URLs literally by hand. To get around this I set up a little URL shortener at zerokspot.com/l which the produces URLs like these:

```
https://zerokspot.com/l/8s6jx
# -> https://zerokspot.com/
---

https://zerokspot.com/l/t7w4l
# -> https://www.oreilly.com/conferences/from-laura-baldwin.html
```

I can create those short links using a simple HTTP API which I access with...

1. an iOS Shortcut which I can access from the "share" dialog
2. a simple "short" shell command which I use on my laptop

When I now want to write a URL I create a short version of it using either of these methods and then reference just the code-part in the journal. I usually prefix that with a little symbol (<i class="far fa-external-link"></i>).

That's pretty much it but that little trick has kept me sane while even writing entries with multiple URLs in them.

Theoretically, I could have used also existing URL shorteners like bit.ly but I wanted to keep the complete URL under my control while note leaking links to third parties. So I sat down and just wrote one myself which you can now find on <https://github.com/zerok/shortlinks>. It's just a little web interface around a [SQLite](https://sqlite.org/index.html) database. Nothing fancy but it gets the job done 🙂
