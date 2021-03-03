---
date: '2010-07-03T12:00:00-00:00'
language: en
tags:
- webdevelopment
title: dotLess on a Mac
---


I'm more the sass/compass kind of guy but after having heard so much about
[doless](http://www.dotlesscss.com) from [Daniel](http://www.tigraine.at/), I
thought, I should at least try to get it running on OSX.  Infrastructure-wise
it sadly doesn't really fit into my usual projects but it's always great to
have a choice :D Luckily the two are more or less equal feature-wise. That's
sad, scoped variables in (dot)less are really nice ;-)

So let's get to it ...

------------------------------

So basically all you need in order to work with dotless on a Mac is Mono.
Sounds easily and it actually should be but there were a few architectural
issues that prevented me from getting it installed in MacPorts as well as in
homebrew. So, first I tried it with MacPorts, got frustrated, deleted
MacPorts, accidentally deleted a whole bunch of still usable databases
with it and then installed homebrew just to notice that there isn't a formula
for Mono in there and compiling it manually with glib etc. provided by
homebrew doesn't work either ...

So after fixing all my virtualenvs and scripts that still pointed to
/opt/local/bin, I simply went over to
[mono-project.com](http://www.mono-project.com) to get a binary package
(boring).

Once Mono is installed, getting dotless to work is actually a non-issue. It
just works right away as far as I can tell. I just ran a couple of test cases
provided in the source package and wrote a rather minimal example :-) Again, I
was a bit lazy and just got one of the binary packages from [their download
archive](http://www.dotlesscss.com:8081/viewLog.html?buildId=lastPinned&buildTypeId=bt3&tab=artifacts&guest=1).
After downloading and extracting the zip you get a `dotless.Compiler.exe`
which lets you compile .less files into .css files etc. I usually place 3rd
party apps somewhere in my ~/.local folder.

<pre><code>
$ cd $DOWNLOAD_DIR
$ mkdir dotless && cd dotless
$ unzip ../dotless-*.zip && cd ..
$ mv dotless ~/.local
</code></pre>


Since .exe files are not easily executable under OSX as they are under Windows
or with some binfmt magic under Linux, I also wrote a little wrapper script:

<pre><code>
#!/bin/bash
exec /usr/bin/mono ~/.local/dotless/dotless.Compiler.exe $@
</code></pre>

But, hey! I can't really make installing Mono any easier for you than
installing the official package, but at least I can provide a [homebrew
formula](http://gist.github.com/462752) ;-)

<script src="http://gist.github.com/462752.js?file=dotless.rb"></script>
