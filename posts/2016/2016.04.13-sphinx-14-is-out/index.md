# Sphinx 1.4 is out

Somehow I totally missed this
one! [Sphinx 1.4.0](http://www.sphinx-doc.org/en/stable/changes.html#release-1-4-released-mar-28-2016)
was released right before I left for DjangoCon Europe on March 28 while
[1.4.1](http://www.sphinx-doc.org/en/stable/changes.html#release-1-4-released-mar-28-2016)
was pushed out the door just yesterday.

Looking through the release notes there are tons of changes and new feature but
I haven't yet seen anything that would prevent me from upgrading or even make
the process harder than a change in the requirements.txt 💖

The only thing that I will probably forget the first couple of times is that
[sphinx\_rtd\_theme](https://pypi.python.org/pypi/sphinx_rtd_theme/) is now an
optional dependency and that I therefore have to install it manually ...
everywhere 😉

On other fronts, I like the move away from the strftime-style of declaring date
formats towards things like the
[Locale Date Markup Language](http://unicode.org/reports/tr35/tr35-dates.html#Date_Format_Patterns). I'm
using Sphinx mostly outside of the classic Python context and operating with
more general formats makes explaining usually much easier. So far I didn't have
to mess with date formats here, though 😉

With [#1970](https://github.com/sphinx-doc/sphinx/issues/1970) the basic theme
also received keyboard shortcuts for jumping from one topic to the next (and
vice versa). I'm really tempted to port this feature to the
sphinx\_rtd\_theme as this is what I'm using everywhere. The code looks to be
highly portable 💖
