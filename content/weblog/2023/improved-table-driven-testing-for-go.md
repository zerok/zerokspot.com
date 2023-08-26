---
date: "2023-08-26T10:42:25+02:00"
tags:
- golang
- testing
title: Improved table-driven testing for Go
---

Table-testing/table-driven testing in Go has been a thing for a very long time, yet there are still  articles coming out about it that add some new ideas at least for the way I use them. Exactly this happened a couple of days ago when I stumbled (thanks to a co-worker) upon Dominik Braun's article titled ["Writing Clean and Efficient Table-Driven Unit Tests in Go"](https://semaphoreci.com/blog/table-driven-unit-tests-go) over on [SemaphoreCI.com](https://semaphoreci.com/).

First of all, he gives a great rundown of what table-testing/table-driven testing is and how to get started. He then also goes into a few guidelines on how to make them more efficient and readable over the long run. I use pretty much all of them already but two which I will from now on most likely try to incorporate into my process:

## Maps instead of slices

Traditionally, I've been putting my test cases into a slice of structs. This kind of follows the semantics of a table by having "rows" in fixed order. One thing that's always annoying there, though, is that I also have to pass some sort of "name" property to my test structs so that I can distinguish the test runs.

Dominik now mentions the use of maps instead of slices here.

```
tests := map[string]struct {
	Input          string
	ExpectedOutput string
}{
	"simple": {
		Input:          "hello",
		ExpectedOutput: "world",
	},
}
```

This not only makes adding the extra name-property redundant but also makes them run not necessarily in the order you've defined, adding some extra randomness to the execution.

## Helper structs

The other little annoyance I always have with table-testing is that my test structs tend to get messy especially for more complex parameter-setups. For this he suggests to create helper structs to better group the input parameter and the expected output aspects of a run:

```
type Input struct {
	Value string
}
type ExpectedOutput struct {
	Output string
}
tests := map[string]struct {
	Input          Input
	ExpectedOutput ExpectedOutput
}{
	"simple": {
		Input:          Input{"hello"},
		ExpectedOutput: ExpectedOutput{"world"},
	},
}
```

This is especially nice if you also test functions that may or may not return errors or have lots of parameters (for which you could also use prototype functions).
