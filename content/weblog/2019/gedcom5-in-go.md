---
title: "GEDCOM 5 in Go"
date: 2019-01-06T12:29:34+01:00
tags:
- genealogy
- brotherskeeper
- golang
- gedcom
---

During the most recent Chaos Communication Congress in Leipzig I started a new side-project. Before I go into that, though, a bit of background: Over the course of the last decade or so my father has created a family tree inside a software called [Brother's Keeper](https://bkwin.org/). Sadly, BKWin hasn't aged all that well. While it has been in active development for about 30 years, the UI has seen only minor improvements over the last 15 years. It still feels like a basic Windows 95 app that has been ported over from Windows 3.11 that has been ported over from whatever.

Then, last year, I stumbled upon [Dgraph](https://dgraph.io/) (a graph database) and thought it might be a nice datastore for something like a family tree. But before I can start working on my own little genealogical software I need to solve something first: How can I use all that data that my father and my uncle have already collected when I don't want to install BKWin in VM? Luckily, Brother's Keeper can export data using the [GEDCOM](https://en.wikipedia.org/wiki/GEDCOM) 5 format. When I looked for it in early December 2018 I couldn't find any Go library for decoding and encoding this format, though, so I set out to create one!

Meet <https://gitlab.com/zerok/go-gedcom5>!

There isn't all that much there, yet. For now I can only decode some information about individuals and families but I hope I find the time to continue working on it over the next couple of months. If you have a GEDCOM file, you can parse it using this snippet:

```go
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	gedcom5 "gitlab.com/zerok/go-gedcom5"
)

func main() {
	flag.Parse()
	fp, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer fp.Close()

	var gfile gedcom5.File

	err = gedcom5.NewDecoder(fp).Decode(&gfile)
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, record := range gfile.Records {
		fmt.Printf("%v\n", record)
	}
}
```

I've also added a little executable to the project inside the `cmd/gedcom5` package to make experimenting with the library easier. To use it, you have to run `make install` inside the project's root folder. This will create `bin/gedcom5` and copy that binary into `$GOPATH/bin`. If you now run `$ gedcom5 your-family-tree.ged` you will see the parsing result of the given file 🙂

