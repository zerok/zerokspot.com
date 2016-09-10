---
date: '2010-08-29T12:00:00-00:00'
language: en
tags:
- python
- windows
- mingw
title: Seeing a Python through your Windows?
---


OK, the title might be a bit vague so let me first describe what this is all
about: I've been working on Unix and Unix-like systems more or less
exclusively for the last nearly 10 years. I switched from Windows to Linux
right after the new millennium started and only started to look back a bit
thanks to Windows 7. So now, for about three weeks, I can call a
Windows-PC my own again (look what Starcraft2 can do to you). Since I'm still
very much a Python guy (although I really like quite a few other languages) I
naturally want to continue coding Python even on Windows.

This works pretty well if you don't do one of two things: (1) Prefer
Powershell over cmd.exe or (2) require some C-extension for Python (like
keyring's Windows integration). Getting all this working will definitely take
some time, so in this post I want to write about some of my first steps and
impressions.

----------------------

## Powershell integration

Powershell actually isn't the real problem. You can easily use Python scripts
in it to do basically everything you could do anywhere else. The problems
start when you try to use [virtualenv](http://pypi.python.org/pypi/virtualenv)
on top of Powershell. Right now, virtualenv only ships with some .bat files to
do the activation for you, but the changes to the environment don't seem to
[propagate from the batch file into the
shell](http://stackoverflow.com/questions/1365081/virtualenv-in-powershell).
I guess, you could get around this by creating your own activation script
but for Powershell or wrap the activation inside a PS-script that
propagates the environment changes (I have that one working here thanks to
some snippets I found all over the net, but it's still kind of a mess). If
time permits, I really want to investigate this further and hopefully
provide a patch to virtualenv for that.

That and learning Powershell scripting as a bonus :D

## C-Extensions

Installing C-extensions for Python at first looked like a real big problem to
me. I generally thought I had about 3 options here:

1. Going straight with [cygwin](http://www.cygwin.com/) and so basically giving up on Windows for C-extension-related development altogether or
2. Getting a license of Visual Studio since I had read somewhere that that is still used for building the official binaries.
3. Go with [MinGW/MSYS](http://www.mingw.org/)

Since I had never used MinGW before and kind of wanted to stick to "native"
solutions this time around, my plan was to go with something like Visual
Studio Express and see how far I could get.

Well, not that far, it turned out. Importing the project files in PC/PCbuild
failed and so, being a total VS newbie, I reached a dead end after only a
couple of minutes. Going with cygwin was beside the point so next I tried to
install MinGW. And once I found the [right
documentation](http://www.mingw.org/wiki/Getting_Started) it was pretty
simple:

1. Download the auto installer (current alpha, but it works pretty well) and extract it into C:\MinGW
2. Add that folder's bin directory to your %Path%
3. Run ``mingw-get install gcc mingw32-make`` to install gcc and MinGW's make version

After that I configured distutils to use mingw's compiler. Thankfully
distutils is quite customizable in that regard. All you have to do is create
a pydistutils.cfg file in your home directory and add following lines to it:

<pre >[build]
compiler = mingw32
</pre>

When you next encounter a source page that includes C-extensions, distutils
will use MinGW's gcc to compile it. At least, I can finally install keyring...
just to find some issues with it that indicate some very limited testing on
Windows but that's another story. I'm also not sure yet, if the compiled
modules really work. I just tried it with some calls to them and at least they
didn't die. I just take this as a good sign ;-)

## Feelings so far

Well, so far I have a more or less working setup but for the near future I'll
probably get most of my work on in the VM instead of the underlying OS (Win7).
Installing MinGW was definitely easier than expected but I keep running into
issues mostly related to the libraries I want to use (mongoengine and keyring)
so this journey is far from being over :-)
