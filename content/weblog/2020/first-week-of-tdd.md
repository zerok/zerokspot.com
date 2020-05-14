---
title: First week of TDD
date: "2020-01-19T20:45:00+01:00"
tags:
- tdd
- agile
- testing
---

Late last year I noticed a problem with the code I was producing: I
didn't like the degree to what it was testible. Along came Robert
C. Martin's book ["Clean Agile"][ca] which reminded me of TDD. A
couple of years ago I had given it a try but never fully embraced
it. Given my dissatisfaction I now want to give it another chance and
see if I've changed enough since then for TDD to finally also work for
me.

With the new year I thought it might be a good opportunity to start
TDD'ing right when getting back to work. I had a new feature to work
on and so I started coding that in pure TDD style one micro-change at
a time... And I'm loving it so far!

Right now my performance is probably a bit lower than normally but I'm
far more confident about the code. This especially made it easier for
me to refactor and tune it and until I was finally happy with the
result.

It also made me explore the libraries and dependencies I've been using
even more than before.

## Testing Cobra commands

For instance, I've been using [Cobra][c] for years now and this was
the first time I noticed that its commands are far easier to test when
you generate new instances of them after every test:

Normally, Cobra code is somehow like this with lots of global state:

```
var someCmd := &cobra.Command{
    ...
}
```

Running tests using `someCmd.Execute()` repeatedly led to some weird
behaviour around flag parsing. Sure, there are methods like
`.ResetFlags()` but the one thing that consistently works for me right
(and probably also [Hugo][h] where I got that pattern from) is to put
commands into builder/constructor functions:

```
func newSomeCmd() *cobra.Command {
    cmd := &cobra.Command{}
    return cmd
}
```

This has the additional advantage that I can now also include flag
definitions within the closure of the constructor and so not leak
their content to other commands:

```
func newSomeCmd() *cobra.Command {
    var verbose bool
    cmd := &cobra.Command{}
    cmd.Flags().BoolVar(&verbose, "verbose", false, "Verbose logging")
    return cmd
}
```

Having constructors makes the registration of sub- and parent-commands
a bit more complicated but that's only a small price to pay (until
I've found a better way).

## Looking for testing opportunities

This weekend I also started to finally play around with [Swift][s]
again and the first thing I did was to look for ways to test the code
I'd soon be writing.

Using the same mindset I will also try to put new dependencies first
into test-cases instead of writing scripts or main-functions around
them.

I've learnt so much in just one week... my testing code is obviously
after a single week not where I want it to be but now I finally have
the confidence again that it will eventually get there!

[ca]: http://www.informit.com/store/clean-agile-back-to-basics-9780135781869
[c]: https://github.com/spf13/cobra
[h]: https://gohugo.io/
[s]: https://swift.org/
