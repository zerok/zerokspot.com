---
title: "Preventing connection leaks with database/sql"
date: 2017-07-31T18:56:00+02:00
tags:
- golang
- sql
---

Over the last week I've had tons of fun working with Go's database/sql package
for interacting with a PostgreSQL database (using the
native [pq package](https://github.com/lib/pq)). It's just a simple,
straight-forward abstraction over the native drivers without any major
surprises. Simply pure joy! Just as with the net/http package, though, you can
accidentally leak resources if you are not careful.

The database/sql package internally manages a connection pool in a completely
transparent way to your application. All you have to do is create a `sql.DB`
object once and call `Close()` on it when your application is shutting down. In
a web application that is using database connections, for instance, you don't
have to set up connections for each request but can re-use a single global
`sql.DB` instance. The standard library does all the connection management for
you.

What you have to keep in mind, though, is that the underlying connection stays
active as long as there is data to be consumed. This means that whenever you
have a `sql.Rows` object, you have to call `Close()` on it; otherwise you will
leak connections. These connections will build up over time and eventually your
server will tell you that it doesn't accept any more connections. That's similar
to the net/http package's HTTP client where you leak goroutines if you don't
close the response body once you've finished reading from it. There is some
implicit closing going on when you're iterating through the result set using the
`Next()` check until the very end, but it's still safer to explicitly close the
`sql.Rows` instance.

Especially when you only want a single row, you can avoid the `Query` method
altogether in favour of
the [QueryRow](https://golang.org/pkg/database/sql/#DB.QueryRow) method. This
doesn't return a `sql.Rows` object but instead can be directly chained with a
`Scan` method. Here you don't have to worry about forgetting to close anything.

Debugging this issue turned out to be rather easy thanks to
the [Stats](https://golang.org/pkg/database/sql/#DB.Stats)-method, which
provides you the number of open connections to the backing server. Once I had
fixed my connection leak, I immediately exposed that counter to our Prometheus
instance 😁
