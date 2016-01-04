# TIL: New String Formatting in Python

For a language that claims to focus on one right way to handle most situations,
string formatting in Python is quite a buffet and it's getting more diverse by
the year. Starting with [Python 3.6][py36] we now have three ways to format
strings (outside of simple concatinating things or using `string.Template`):

* using the % operator
* str.format
* interpolated strings

(If you are not motivated to read all this, I will also give a slightly extended
lightning talk about this at the next
[PyGraz meetup in February 2016][pygraz1602] with probably a few more examples
😊)

-----------------


## %-formatting

%-formatting has been part of the language since at least version 1.0. If you've
done anything with Python prior to Python 3 you know it.

```
"%s %s" % ('Hello', 'World',)
```

It is more or less the equivalent to C's `sprintf` with all its ups and
downs. It gets the job done but you have to do quite a bit of work around it.

Since this only supports a limited set of types you have to convert your custom
objects into one of these before passing it into the string formatter.

Many years later, the native string datatype was extended with a `format`
method:


## str.format()

... was added in Python 2.6 in October 2008 alongside things like
context-managers. As has been detailed in [PEP-3101][pep3101] it tries to solve
some of the shortcoming of the old % binary operator like it only supporting a
limited set of types or actually its handling of the right-side component of the
whole expression which had some special cases that easily lead to errors.

```
>>> "%s" % ("lala",)
'lala'
>>> "%s" % "lala"
'lala'
```

Since `.format` is a method and not an operator (which is mapped to a binary
method), handling of arguments has become more explicit. If you pass a string,
it is interpreted as a string. If you pass a tuple containing just one string,
it is interpreted as a tuple containing one string:

```
>>> "{}".format("lala")
'lala'
>>> "{}".format(("lala",))
"('lala',)"
```

Compared to % it also supported giving your parameters names without having to
work with dictionaries out of the box:

```
"{firstname} {lastname}".format(firstname="Horst", lastname="Gutmann")
```

Initially, this was intended as a complete replacement for the %-operator (it
*was planned* to deprecate the old-style formatting functionality with
[Python 3.1](https://docs.python.org/3.1/whatsnew/3.0.html)) but that never
fully happened. The core features of this string formatter are mostly the same
as with the old %-operator, but the syntax is a bit different and IMHO more
intuitive. Actually, because of that Ulrich and I created [pyformat.info][pfi]
to help people migrate to the new system.

But, obviously, PEP-3101 didn't stop at just cleaning up the old feature-set. It
also introduced a protocol that allows for a more versatile interaction with
custom classes:

```
class Country:
    def __init__(self, name, iso):
        self.name, self.iso = name, iso

    def __format__(self, spec):
        if spec == 'short':
            return self.iso
        return self.name

country = Country("Austria", "AUT")

print("{}".format(country))
print("{:short}".format(country))
```

You can think of the `__format__` method as a `__str__` for string formatting
that you can pass options. The moment you have a `__format__` method in your
object, it will be used instead of `__str__` when you're using the format-method
(unless you do something like `"{!s}".format(country)"`).

You can actually find a nice example for how to use that within the
[datetime.date class in Python 3.4][dtfmt]:

```
class date:
    ...
    def __format__(self, fmt):
        if len(fmt) != 0:
            return self.strftime(fmt)
        return str(self)
```

This allows you format dates directly within the "parent" string format so that
you no longer have to first convert your date into a string and then pass that
into the string formatter:

```
import datetime
print("Today is {:%A}".format(datetime.datetime.now()))
# Today is Thursday
```


## PEP-0498: String Interpolation

While it's right now the recommended way to doing string formatting, `.format`
is quite verbose:

```
a = "Hello"
b = "World"
"{} {}".format(a, b)
# vs.
"%s %s" % (a, b,)
```

PEP-0498 tries to improve this situation by offering something that has been
common to other languages like Ruby, Scala and Perl for quite some time:
Interpolated strings. Here expressions can be integrated directly into the
string itself which means you don't have to call any additional functions
explicitly.

ES2015 introduced this feature to the JavaScript-world "recently" where it is
referred to as "template strings":

```
const username = "Horst";
const welcomeMsg = `Hello, ${username}!`;
```

In Python backticks have a bit of
[history](https://docs.python.org/3.0/whatsnew/3.0.html#removed-syntax) up to
Python 3.0 so they are not available. Introducing them again would also once
again affect the basic syntax of the language. Instead, another literal-prefix
was introduced: `f`.

```
a = "Hello"
b = "World"
f"{a} {b}"
f"{a + ' ' + b}"
```

You no longer need to explicitly call the `.format()` method of a string but
simply mark the format with the `f` prefix and inline the expressions you want
to have included in the final string. Otherwise they are supposed to offer the
same functionality as what you get out of `.format()`. These formatted strings
are also referred to as "f-strings" in the documentation.

That actually looks pretty nice but as Python 3.6 is slated for release in
another 12 months you will have to wait a little longer. That being said, the
code is already there so you could just grab a Python 3.6 pre-release or tip
using something like [pyenv][] and give it a go 😊

And there is more. There is another PEP ([0501][pep0501]) which wants to
introduce i-strings that result in lazily evaluated string so that you can for
instance do i18n or security checks on them before the final evaluation. While
that proposal has been deferred until further discussions have taken place, it
looks like a neat idea.

But back to f-strings: If you want to know more about why string interpolation
was solved the way it has been, take a look at [PEP-0502][pep0502] which
includes a detailed discussion of the motivation behind and inspiration from
other languages for this feature.

[pep3101]: https://www.python.org/dev/peps/pep-3101/
[py36]: https://docs.python.org/3.6/whatsnew/3.6.html
[pep0502]: https://www.python.org/dev/peps/pep-0502/
[pep0501]: https://www.python.org/dev/peps/pep-0501/
[pygraz1602]: https://pygraz.org/meetups/2016-02-02
[pfi]: https://pyformat.info/
[dtfmt]: https://hg.python.org/releasing/3.4/file/tip/Lib/datetime.py#l725
[pyenv]: https://github.com/yyuu/pyenv
