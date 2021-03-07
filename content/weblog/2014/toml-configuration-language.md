---
date: '2014-09-09T22:16:53+02:00'
language: en
tags:
- development
title: TOML - The other configuration language
---


While taking a look at Heka last weekend I stumbled upon a configuration file
format I hadn't heard of before: [TOML][]. TOML (or "Tom's obvious, minimal
language") is a new entry into the field with an age of just about a year. It
was created by Tom Prescot-Werner, previously of Github, and is basically
the good old INI format but with more data structures and different
string handling.

```
# What would be described as a section in INI/CFG files are simple hashmaps
# ("tables") in TOML.
[logging]
verbosity = 4

# There are a bunch of native datatypes in here:
array = [1, 2, 3, 4]
boolean = true
date = 1981-10-11T07:08:09Z
strings = "hello world"
multiline_strings = """
If the first line after the opening quotes
is indented, it will be used as default
indentation."""

# Strings esp. have some nice helpers that should keep your configuration
# quite readable: https://github.com/toml-lang/toml/#string

# You can also nest dictionaries by putting a dot into the key name:
[logging.output]
file = "/var/log/log.log"

    # Elements can be indented without messing up the parser
    [logging.something_else]
    key = "value"

# You can create a list of tables like the list of "servers" here by
# doubling the braces around the key.
[[servers]]
url = "http://zerokspot.com"

# the second server
[[servers]]
url = "http://h10n.me"

```

I'm pretty sure that the prefixing of sub-tables might be quite annoying if
you're dealing with many such entries, so for these situations [JSON][] or
[YAML][] are perhaps better choices. But for "sane" use-cases this looks
really nice!

There are also already [libraries][lib] for more or less every major language
out there, so I wouldn't have any problem using it in my Python, JavaScript or
Go projects. That being said, at least one of the Python implementations has a
[problem with some of the ways strings can be used in TOML][pyprob]. Luckily,
that's proably the feature I would use the least here :)

A [formal specification][specpr] is current available as a PR to the main
repository and it looks like one of the few missing pieces before putting a
1.0-stamp onto the language. In the meantime there also exists a [test
suite][ts] against which implementations can be tested.

I have not used TOML in any project yet, but with one of my next ones I
definitely want to give it a try :)

[toml]: https://github.com/toml-lang/toml/
[lib]: https://github.com/toml-lang/toml#implementations
[specpr]: https://github.com/toml-lang/toml/pull/236
[yaml]: http://yaml.org/
[pyprob]: https://github.com/uiri/toml/issues/19
[json]: http://json.org/
[ts]: https://github.com/BurntSushi/toml-test
