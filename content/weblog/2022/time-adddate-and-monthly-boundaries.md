---
title: Time.AddDate and monthly boundaries
date: "2022-03-31T20:01:06+02:00"
tags:
- golang
- til
---

A couple of days ago I ran into a weird issue: I needed to run some test-cases with dates from different months as input. Everything worked fine but then, on March 29, it all of a sudden didn’t anymore.

The code looked like this:

```go
now := time.Now()
previousMonth := now.AddDate(0, -1, 0)
```

On March 28 `now.Month()` was `3` and `previousMonth.Month()` was `2`. Midnight came and the results were `3` and `3`.  This happened because, as documented, `time.Date` normalises values:

> AddDate normalizes its result in the same way that Date does, so, for example, adding one month to October 31 yields December 1, the normalized form for November 31.
> — [time.go](https://cs.opensource.google/go/go/+/go1.18:src/time/time.go;l=906)

So what did happen exactly in my case? Go internally created this new Time object:

```go
time.Date(2022, 2, 29, ...)
```

2022-02-29 doesn’t exist but Go tries to normalise this value [by calculating a UNIX timestamp](https://cs.opensource.google/go/go/+/refs/tags/go1.18:src/time/time.go;drc=90462dfc3aa99649de90bb587af56a9cb0214665;l=1402) and the result is actually 2022-03-01 which is, obviously, in the same month as March 29.  

So what was the workaround in my case? My taste-dates are now always on the first day of the month and not based on `time.Now()` 😅
