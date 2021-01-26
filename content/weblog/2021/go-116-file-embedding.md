---
title: 'Go 1.16 feature: file embedding'
date: "2021-01-26T18:39:06+01:00"
tags:
- golang
- 100daystooffload
---

When you develop applications in Go that you want to deliver as a single binary file but that require, for instance, image files then you need to somehow embed those into the binary such that the Go application can access them at runtime. One example for such a scenario is an application that wants to provide you with a JavaScript/web GUI for management.

Over the years there have been tons of tools appearing left and right on the web that do just that:

- [go-bindata](https://github.com/go-bindata/go-bindata)
- [packr](https://github.com/gobuffalo/packr)
- [pkger](https://github.com/markbates/pkger)
- [statik](https://github.com/rakyll/statik)
- …

With version [1.16](https://tip.golang.org/doc/go1.16) Go itself now incorporates that functionality through a new package called “embed”. Let’s work with a slightly smaller and simpler example here: An application where we want to embed the `LICENSE.txt` file so that we can easily print that to the user:

	```
	$ ls -1
	main.go
	LICENSE.txt
	
	$ go build -o app
	
	$ ./app --license
	This is a really long license!

OK, so how do we now get that `LICENSE.txt` into the binary? With the `embed` package we have two choices:

## Embedded file-system

`go build` can now embed whole directories trees from your file-system into the binary using a new `//go:embed <glob-patterns>` directive. Let’s say, we want to keep our implementation flexible and in the future perhaps also embed other text files into the binary. In that case, our implementation could look somehow like this:

	package main
	
	import (
		"embed"
		"fmt"
	
		"github.com/spf13/pflag"
	)
	
	func main() {
		var showLicense bool
		pflag.BoolVar(&showLicense, "license", false, "Show the license")
		pflag.Parse()
	
		//go:embed *.txt
		var fs embed.FS
	
		if showLicense {
			data, _ := fs.ReadFile("LICENSE.txt")
			fmt.Printf("%s\n", data)
		}
	}

When the compiler sees that `//go:embed *.txt` directive right in front of an `embed.FS` variable declaration, it will populate that variable at compile time. Such a virtual file-system is pretty simple and has only three methods:

- `Open(name string) (fs.File, error)`
- `ReadDir(name string) ([]fs.DirEntry, error)`
- `ReadFile(name string) ([]byte, error)`

For our purposes we can just take `ReadFile` and print the result to stdout.

You can also specify multiple glob patterns with the embed-directive or even have multiple such directives in a raw before the `embed.FS` declaration!

## File as variable

If you know that you want to just access a single file and have its data embedded either as a `string` or a `[]byte`, you can also do something like this:

	//go:embed LICENSE.txt
	var licenseContent string

## But what about websites?!

As the API for the `embed.FS` struct already hinted at, there’s a new `io/fs` package which basically comes with tons of stuff that was previously associated with the `io/ioutil` or the `os` package.

I’ll leave most of that package out for another time but just keep in mind that we now have that package and also a new `fs.FS` interface with just one method:

	type FS interface {
	    Open(name string) (File, error)
	}

The `http.FileSystem` looks also just like that and, lo and behold, `embed.FS` implements it 😬

So, all you have to do is to create an `embed.FS` as shown above and then mount that file-system using `http.FS(embeddedFS)` and you’re done.

If we just add this line at the end of our first example, we can fetch the `LICENSE.txt` from `http://localhost:9999/LICENSE.txt`!

	http.ListenAndServe("localhost:9999", http.FileServer(http.FS(fs)))
	

I really cannot wait for 1.16 to be released 😀 If you’re already curious and want to see the whole documentation, take a look at the [tip docs](https://tip.golang.org/pkg/embed/) and/or grab the latest 1.16 release ([beta1](https://golang.org/dl/#go1.16beta1) at the point of writing this)!

Oh btw.: In the examples above I’ve always used a simple text file. In practice, though, there’s nobody stopping you from, for instance, compressing your file with xz, then embed it, and then de-compress it when you need it at runtime 🤪
