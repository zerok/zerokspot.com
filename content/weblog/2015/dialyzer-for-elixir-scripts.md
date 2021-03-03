---
date: '2015-07-18T10:27:26+02:00'
language: en
tags:
- elixir
- erlang
title: Dialyzer for Elixir Scripts
---


I'm currently in the process of making my way through the [Elixir][] tasks in
[exercism][] in order to learn the language. Sadly, these code packages are
delivered not as mix packages but as simple .exs files and all examples I've
found where around using [dialyzer][] when combined with mix for doing static
analysis.

(Please note that this is basically just a documentation of some of the
baby-steps I'm taking in Erlang/Elixir land right now. All of this is probably
obvious for people who have lived there for years.)

-------

Luckily, it seems to be quite simple to run the tool against single compiled
files (and why wouldn't it? I'm new to Erlang and Elixir 😉):

```
$ elixirc -o tmp word_count.exs
$ dialyzer tmp/Elixir.Words.beam
  Checking whether the PLT /Users/zerok/.dialyzer_plt is up-to-date... yes
  Proceeding with analysis... done in 0m0.28s
done (passed successfully)
```

If you've never executed dialyzer before (as I have) you will get a message
that it first requires something called a "PLT" (Persistent Lookup Table)
which contains the results of an analysis. In this case you need it in order to
give dialyzer a starting set of known types.

For my system I simply took the local Elixir installation as well as the Erlang
Runtime as starting point (as per Dave Thomas' awesome
[Programming Elixir book][pe]):

```
$ dialyzer --build_plt --apps erts /usr/local/Cellar/elixir/1.0.5/lib/elixir
```

I might just pack all that into a little shell script to make it easier to work
through the exercism tasks but now I'm at least able to check my `@spec`s 😊

[exercism]: http://exercism.io/
[dialyzer]: http://www.erlang.org/doc/apps/dialyzer/dialyzer_chapter.html
[Elixir]: http://elixir-lang.org/
[pe]: https://pragprog.com/book/elixir/programming-elixir
