---
title: "go-querystring: Convert a struct into url.Values"
tags:
- golang
- libraries
- development
date: "2021-10-31T09:33:00+01:00"
---
While making my way through the excellent [xanzy/go-gitlab](https://github.com/xanzy/go-gitlab) library I stumbled across little gem I hadn't seen before: [google/go-querystring](https://github.com/google/go-querystring), a library that can marshal a struct into a URL querystring. 

```go
package main

import (
	"fmt"

	"github.com/google/go-querystring/query"
)

type QueryOptions struct {
	Size            int    `url:"size,omitempty"`
	Filter          string `url:"filter,omitempty"`
	ShouldBeIgnored int    `url:"-"`
}

func main() {
	o := QueryOptions{
		Size:            100,
		Filter:          "all",
		ShouldBeIgnored: 1000,
	}
	vals, _ := query.Values(o)
	fmt.Println(vals.Encode())
	
	// Output:
	// filter=all&size=100
}

```

As with things like `enocding/json` you can customize the marshalling process by attaching tags to the struct you want to convert into `url.Values`.  The example above shows the use of the `url` tag including the `omitempty` option. If you don't want to include a field in the output, you can use the `-` value for the `url` tag.

## Times

A common use-case is that you want to encode a `time.Time` instance as part of a querystring. By default, encoding will use RFC3339 as format. If you want to change that to follow a custom layout, you can use the `layout` tag:

```go
package main

import (
	"fmt"
	"time"

	"github.com/google/go-querystring/query"
)

type QueryOptions struct {
	CreatedAt time.Time `layout:"2006-01-02"`
}

func main() {
	o := QueryOptions{
		CreatedAt: time.Now(),
	}
	vals, _ := query.Values(o)
	fmt.Println(vals.Encode())
	
	// Output:
	// CreatedAt=2021-10-31
}
```

In case you just want to get a timestamp out of it, use `url:",unix"`, `url:",unixmilli"` or `url:",unixnano"`.

## Slices

Slices/arrays in querystrings are always something weird. Some webframeworks expect them to be sent using special keys like `param[]=1&param[]=2`, while others expect just a single value that is then separated with spaces or commas or even just the normal parameter name just repeated. go-querystring has some flags for most of these.

Let's assume that you have this kind of struct:

```go
type Data struct {
	Values []string
}
```

... and then assign a simple array with three strings to it (`["a", "b", "c"]`). Depending on the formatting option, you'd get the following output:

```
// default (no option)
Values=a&Values=b&Values=c

// numbered
Values0=a&Values1=b&Values2=c

// space
Values=a+b+c

// comma
Values=a%2Cb%2Cc

// brackets
Values%5B%5D=a&Values%5B%5D=b&Values%5B%5D=c
```

## Booleans

By default, booleans are encoded as `true` or `false` but you can change that by setting the `int` option inside the `url` tag:

```go
package main

import (
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

type Data struct {
	Flag bool `url:",int"`
}

func main() {
	var vals url.Values
	vals, _ = query.Values(Data{Flag: true})
	fmt.Println(vals.Encode())
	
	// Output:
	// Flag=1
}
```

## Custom values

When you have custom data structures inside your main struct, then you'll most likely want to also have some kind of custom encoding for it. This can be done by having the struct implement the [`query.Encoder`](https://pkg.go.dev/github.com/google/go-querystring/query#Encoder) interface:

```go
type Encoder interface {
	EncodeValues(key string, v *url.Values) error
}
```

```go
package main

import (
	"fmt"
	"net/url"

	"github.com/google/go-querystring/query"
)

type CustomField struct{}

func (v *CustomField) EncodeValues(key string, values *url.Values) error {
	values.Set(key, "CUSTOM")
	return nil
}

type Data struct {
	Flag *CustomField
}

func main() {
	var vals url.Values
	vals, _ = query.Values(Data{Flag: &CustomField{}})
	fmt.Println(vals.Encode())
	
	// Output:
	// Flag=CUSTOM
}
```


## Conclusion

That's pretty much every use-case I've run into so far with go-querystring and I really liked the way it allowed me to handle them!  There are also some options that I haven't mention here simply because I haven't used them yet. You can find all of them in the documentation for the [`query.Values`](https://pkg.go.dev/github.com/google/go-querystring/query#Values) function.
