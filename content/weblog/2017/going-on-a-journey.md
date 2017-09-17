---
title: "GOing on a Journey"
date: 2017-09-17T10:39:52+02:00
tags:
- golang
- meet-the-meetups
---

**Note:** This the long version of a talk I was giving on Sept 15
at [Meet-the-Meetups 2017](https://meet-the-meetups.org/events/graz-2017/) in
Graz about why I'm focusing on Go and what the language and ecosystem has to
offer that makes it so special for me.

## The special mix

Up until about 4 years ago, Python and NodeJS were my main working
environments. For everything where I needed a bit more speed I went with NodeJS
while I used Python as my go-to language for most automation tasks. Since I was
mostly doing web-based projects I could usually solve pretty much anything I
came across with either of those.

Then I started to learn Go and simply fell in love. If you look at Go from a
language-feature perspective, it doesn't appear to be all that modern. Instead,
the focus during its design phase was on being **pragmatic** and **easy to
learn**, something many other and perhaps more modern languages are clearly
lacking. For me, though, these two were just part of the reason that got me
hooks. Let's look at the others:

### Go's approach to concurrency

While languages like C, Rust, Erlang et al. offer means to somehow easily work
with threads and multiple
processes, [go-routines](https://golang.org/doc/effective_go.html#goroutines)
and [channels](https://golang.org/doc/effective_go.html#channels) as well as
what is in the [sync package](https://golang.org/pkg/sync/) of the Go
standard-library makes concurrent solutions for your problems even easier.

```go
package main

import "fmt"
import "sync"

func main() {
	wg := sync.WaitGroup{}
	// Make sure that we wait for 2 "Done" calls in the end.
	wg.Add(2)

	for i := 0; i < 2; i++ {
		go func(idx) {
			for j := 0; j < 10; j++ {
				fmt.Printf("Routine %d: %d\n", idx, j)
			}
			wg.Done()
		}(i)
	}

	// Let's wait until both goroutines are done with their
	// printing.
	wg.Wait()
}
```

Go-routines abstract just enough of the underlying implementation away that you
normally just have to think about how your application-threads should
communicate. You can either use channels or go with some of the gems in the sync
package like `sync.WaitGroup` and `sync.RWMutex`.

### A powerful standard library

But `sync` is just one example for a great package inside the
standard-library. You will also find these (among many many other modules):

- a bare-bone **logging** implementation
- a **H2-enabled HTTP server and client** module
- a **JSON** parser
- a simple **commandline-argument parser**
- an **abstraction for relational databases**
- ...and so much more

Let's look at a small example: How to decode a simple JSON document like this:

```json
{
    "name": "Horst",
    "favorite_color": "blue"
}
```

We first have to define a type into which this JSON object should be parsed into
followed by using the `json` package to do the actual parsing:

```
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type profile struct {
	Name          string `json:"name"`
	Country       string `json:"country"`
	FavoriteColor string `json:"favorite_color"`
}

func main() {
	file, err := os.Open("cmd/json-parser/demo.json")
	if err != nil {
		log.Fatalf("Failed to open test file: %s", err.Error())
	}
	defer file.Close()
	var p profile
	if err := json.NewDecoder(file).Decode(&p); err != nil {
		log.Fatalf("Failed to decode data: %s", err.Error())
	}
	fmt.Printf("name: %s\ncountry: %s\nfavorite color: %s\n", p.Name, p.Country, p.FavoriteColor)
}
```


### Interfaces

While in other languages, interfaces are something that your implementation has
to explicitly reference, in Go they are **implicit**. This removes the
additional coupling between the two without removing any of its benefits.

They are also **used through the standard library** for common operations like
copying data from an `io.Reader` into an `io.Writer`.

```
// io.Copy(io.Writer, io.Reader) simply copies everything from a reader
// into a writer.
sourceFile, err := os.Open("interfaces_io_test.go")
if err != nil {
	t.Fatalf("Failed to open source file: %s", err.Error())
}
fileContent, err := ioutil.ReadFile("interfaces_io_test.go")
if err != nil {
	t.Fatalf("Failed to read source file: %s", err.Error())
}

// A writer can for instance be stdout:
io.Copy(os.Stdout, sourceFile)
sourceFile.Seek(0, io.SeekStart)

// A writer can be a buffer
var sink bytes.Buffer
io.Copy(&sink, sourceFile)
assert.Equal(t, sink.Bytes(), fileContent)

sourceFile.Seek(0, io.SeekStart)
sink.Reset()

// Write to multiple outputs simply because the MultiWriter also implements
// the io.Writer interface and can therefore be used in io.Copy.
mw := io.MultiWriter(os.Stdout, &sink)
io.Copy(mw, sourceFile)
assert.Equal(t, sink.Bytes(), fileContent)
```

## Built-in tooling

The package wouldn't be complete with the tooling that comes along when you
install Go.

## Cross-platform compilation

One of the most important features for me was the ability to build Linux or
Windows binaries right from my Mac without having to install anything else.

```sh
# Build a linux binary
$ GOOS=linux GOARCH=amd64 go build -o myproject

# Build a windows binary
$ GOOS=windows GOARCH=amd64 go build -o myproject.exe
```


### Platform-specific files

You can also provide platform-specific implementations by just following the
simple naming pattern:

```sh
# Common parts
main.go

# Include only on Linux
main_linux.go

# Include only on OSX
main_darwin.go
```

And that's just the beginning as you can assign build-tags to each file and
build only those that match a certain tag-selection.


### Standard code formatting

Go has rather quickly put an end to age-old discussions like "tabs vs. spaces"
thanks to providing a community standard for how Go code should look like and
backing it with a tool to **automatically normalise your source code**
accordingly.

```sh
# Formats all files in the current package:
$ go fmt .
```

### Testing with go test

While languages like Java heavily rely on the community to provide testing
frameworks (e.g. JUnit), the Go comes with a basic yet powerful testing-harness
built-in. Simply name your files containing tests with the **suffix "_test"
(e.g. `store_test.go`)** and run **`go test ./...`** to execute the tests in the
current and all nested packages (except for the vendor-folder since Go 1.9).

### Documentation through godoc

Using `go doc` you can browse source level documentation of anything that is
available on your `$GOPATH`. You can even start a little HTTP server to have
that experience in the browser:

```sh
# Browse information about the sync package
$ go doc sync

# Install godoc and start an HTTP server
$ go get -u golang.org/x/tools/cmd/godoc
$ godoc -http=:6060 -v
```


## Community tooling

This focus on awesome tooling has also spread through the community.

### goimports

Based on `go
fmt` [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) also tries
to resolve **missing imports** and sorts your imports:

```go
package main

func main() {
	fmt.Println("hello world")
}
```

becomes

```go
package main

import "fmt"

func main() {
	fmt.Println("hello world")
}
```


### gomodifytags

Going back to the JSON example from above, writing all these tags and changing
them afterwards can be rather tedious, esp. for larger structs. Writing tags is
also quite error-prone as their syntax is not enforced during compile-time.

[gomodifytags](https://github.com/fatih/gomodifytags) comes to the rescue here.

```go
package main

type Address struct {
	Line       string
	City       string
	PostalCode string
	Country    string
}

type Person struct {
	Name      string
	Addresses []Address
}
```


With *gomodifytags* we can now easily add camel-cased JSON tags to the Address
struct:

```sh
$ gomodifytags -file modifytags.go -line 4,7 -add-tags json -transform camelcase
package main

type Address struct {
	Line       string `json:"line"`
	City       string `json:"city"`
	PostalCode string `json:"postalCode"`
	Country    string `json:"country"`
}

type Person struct {
	Name      string
	Addresses []Address
}
```

### delve

With [delve](https://github.com/derekparker/delve) the community has provided a
quite powerful debugger that has by now also be integrated into editor
like [VSCode](https://code.visualstudio.com/) or [Atom](https://atom.io/).

```sh
$ dlv debug .

# list all source files
> sources

# set a breakpoint
> breakpoint /absolute/path/to/file/go:123

# continue to the breakpoint
> continue

# print all local variables
> locals
```


# Just a few examples

These are just a few examples for the great tool-ecosystem that exists for
Go. Some of this has definitely benefitted from Go's standard library coming
with modules for parsing Go source files into ASTs but also printers for ASTs
into output files. This has made creating source code generators, for instance,
much simpler.

In the following example we rewrite that very same program not replace the
function "lala" with "dummy":

```
package main

import (
	"flag"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
)

func lala() {
}

type visitor struct{}

func (v *visitor) Visit(n ast.Node) ast.Visitor {
	switch node := n.(type) {
	case *ast.FuncDecl:

		if node.Recv == nil && node.Name.Name == "lala" {
			node.Name.Name = "dummy"
		}
	case *ast.CallExpr:

		switch fvalue := node.Fun.(type) {
		case *ast.Ident:
			fvalue.Name = "dummy"
		}
	}
	return v
}

func main() {
	flag.Parse()
	fname := flag.Arg(0)
	if fname == "" {
		log.Fatal("No filename specified.")
	}
	log.Printf("Parsing %s", fname)
	fset := token.NewFileSet()
	pmode := parser.Mode(0)
	f, err := parser.ParseFile(fset, fname, nil, pmode)
	if err != nil {
		log.Fatalf("Failed to parse %s: %s", fname, err.Error())
	}
	ast.Walk(&visitor{}, f)
	lala()
	printer.Fprint(os.Stdout, fset, f)
}
```
