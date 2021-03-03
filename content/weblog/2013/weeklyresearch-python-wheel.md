---
date: '2013-09-29T19:40:36+02:00'
language: en
tags:
- python
- weeklyresearch
title: 'Weekly Research: Python''s Wheel Format'
---


People coming from the Java world know software packaging in the form of .jar files, simple ZIP files with a specific structure that indicates what part has to be "installed" where in order for the VM to make use of it. In Python we've had things like that too with the plural here being used intentionally. You can use source-distribution files, eggs, and probably some more that I don't remember. Over the last couple of years quite a lot of standardization in the form of a handful of PEPs has happened, though, which should clean up this mess.

[Wheel][pep 427] is a replacement for one of these formats: Egg. Eggs were introduced with [setuptools][] which in itself tried to solve some of the shortcomings of [distutils][] (like dependency management, name-spacing, ...). An egg by itself is just a ZIP file with the ".egg" file extension that has a specific structure (just like a .jar), but given its age it doesn't support all the newer standards that were created after it.

---------------------------------------------------------------------------------------------

Eggs and now Wheels are binary distributions that help you keep your deployment environment as simple as possible by not requiring for instance a compilation system being available on the target system. This is especially important on Windows where compiling C-extensions is extremely complicated but also everywhere else binary distributions help keep the amount of time it takes to install a project's dependencies at a minimum.

Just like Egg Wheel is a ZIP file but this time with the ".whl" extension and a slightly different structure based on [PEP 376][] ("Database of installed Python distributions") and [345][pep 345] ("Metadata for Python software packages 1.2") (or newer). Let's walk through this structure using an already existing whl-file. Originally, I had planned to use [pyzmq][] here but sadly found some inconsistencies which I couldn't explain. So I've created my own wheel for [Django][] (a process I will explain later on).

## Filename and structure

The file I'm working with here has the name `Django-1.5.4-py2.py3-none-any.whl` which tells me that this package works on on Python 2 and Python 3 and is not limited to any platform ("any"). Also, the [ABI][] tag has been set to "none", meaning it is basically a source package. The Python-version, ABI and platform tags are specified in more detail in [PEP 425][].

After extracting it using your favorite unzip-tool, you get following structure:

```
|
+- Django-1.5.4.dist-info
|    +- DESCRIPTION.rst
|    +- LICENSE.txt
|    +- METADATA
|    +- pydist.json
|    +- RECORD
|    +- top_level.txt
|    +- WHEEL
+- Django-1.5.4.data
|    +- scripts
|        +- django-admin.py
+- django
```

The .dist-info folder is where all the metadata about the package is kept. METADATA contains information about the software itself (like the author's name, classifiers etc.). `pydist.json` holds basically the same information but formatted in JSON. The `WHEEL` file holds similar information about about the package itself, telling the tool that should work with this file, what version of the format to expect and what platforms etc. are supported by it.

The `RECORD` file has an entry for every single file in the package (with some exceptions like the RECORD-file itself) combined with a checksum, e.g.

```
django/__init__.py,sha256=zXXWJdRfxNNrF84trwiYplAOfE3KK74dRY-Q6geXj-s,269
```

A bit of a surprise was the existance of the `top_level.txt` file in the .dist-info folder, because it actually comes from the [egg-specification](http://svn.python.org/projects/sandbox/trunk/setuptools/doc/formats.txt) and is nowhere mentioned in the PEP 427:

> This file is a list of the top-level module or package names provided by the project, one Python identifier per line.
> Subpackages are not included; a project containing both a ``foo.bar`` and a ``foo.baz`` would include only one line, ``foo``, in its ``top_level.txt``.
> This data is used by ``pkg_resources`` at runtime to issue a warning if an egg is added to ``sys.path`` when its contained packages may have already been imported.
> (It was also once used to detect conflicts with non-egg packages at installation time, but in more recent versions, setuptools installs eggs in such a way that they always override non-egg packages, thus preventing a problem from arising.) 

The .data folder usually contains a subset of following subfolders:

* `purelib`: If the wheel is a pure-Python package, this should contain all the package's content
* `platlib`: Same as above, but for platform dependent packages
* `headers`: Header file for instance for C-extensions
* `scripts`: Executable scripts
* `data`: Non-Python data

While `purelib` and `platlib` exist for historical reasons, the specificiation itself recommends to just put their content into the root of the package (as referenced by the `top_level.txt` file).

In our case, all we have is the scripts folder which contains the `django-admin.py` script. If you take a closer look at the file's content, you will notice the `#!python` like at the very top. Any Wheel-installer should look for this line and replace it with the actual path to the interpreter during the installation process.

## Building a Wheel

Now that we are through the rough intro to the format itself, how did I build my whl-file of Django? Thankfully, this was pretty simple thanks to pip:

```
pip wheel --wheel-dir=~/django-package Django
```

This will download the source dist of Django (if you don't have it installed yet), build it and create a .whl file out of it.

If you're working with the source folder of a project, pip works here too:

```
pip wheel --wheel-dir=~/django-package .
```


## Installing Wheels

Now that you have a whl file, just run `pip install package.whl`to install it. If you want to install something from PyPI and would prefer to have whl files where available, run `pip install —-use-wheel packagename`.

pip also lets you create whl files for all your project's requirements by simply replacing "packagename" in our previous example with a reference to your `requirements.txt`. For more details on using pip for Wheels please take a look at the project's [documentation](http://www.pip-installer.org/en/latest/cookbook.html#building-and-installing-wheels).

## Who is offering Wheels right now?

Sadly, on this front I didn't have much success. So far I found only whl files for [pyzmq][]. [Django](https://code.djangoproject.com/ticket/19252) will probably get them with 1.6 and the code has been ported to the 1.5 branch but it seems like they are not yet part of the release process. [Pillow][] getting it would be kind of [like X-mas](https://github.com/python-imaging/Pillow/issues/310) :-)

[pep 376]: http://www.python.org/dev/peps/pep-0376/
[pep 345]: http://www.python.org/dev/peps/pep-0345/
[pep 427]: http://www.python.org/dev/peps/pep-0427/
[pep 426]: http://www.python.org/dev/peps/pep-0426/
[pep 425]: http://www.python.org/dev/peps/pep-0425/
[abi]: http://stackoverflow.com/questions/2171177/what-is-application-binary-interface-abi
[setuptools]: https://pypi.python.org/pypi/setuptools
[distutils]: http://docs.python.org/2/distutils/
[pyzmq]: https://github.com/zeromq/pyzmq
[django]: https://www.djangoproject.com/
[pillow]: https://pypi.python.org/pypi/Pillow
