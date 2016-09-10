---
date: '2007-09-24T12:00:00-00:00'
language: en
tags:
- python
- scons
title: Automating stuff with SCons
---


Once again a post about a tool that I use quite frequently (or plan to). This time: *SCons*. If you know the GNU Make utility, SCons tries to solve the same problems while having a quite different philosophy. But that's just fine since I also use it for other stuff and not just compiling C code ;-)

-------------------------------

## Motivation

I mostly use SCons for compiling LaTeX documents. Actually I started using it only a week or so ago, but it just makes my workflow there so much more convenient, that I have to lose some words about it here. Previously I was writing plain Makefiles that just ran pdflatex 3 times in combination with bibtex whenever one of the sourcefiles got changed in order to produce some nice PDF file.

The problem here was, that my knowledge of the inner workings of pdflatex was quite limited, so I always ran through the complete cycle, although probably only one single call of pdflatex would sufficed to update my PDF. That's why I was looking for some simple way to automate and speed up the process without having to sift through tons of documentation for the tetex distro.

## SCons to the Rescue

Here is were SCons come in. The *Software Construction tool* `SCons`_ is quite similar to the GNU Make tool in that regard, that it tries to make compiling complex software packages as easy as possible. Basically, if you see a Makefile in some software directory, you can be quite sure that just calling `make` will compile ... something :-) Same goes for situations where you see a SConstruct file. There calling `scons` will probably give you some similar result.

A big advantage of SCons over Make in my opinion is, that it's a pure Python tool and platform independent or at least facilitates the creation of platform independent software building. Another nice aspect of SCons is, that its SConstruct files are nothing more than Python scripts themselves with some automatic loads. So you can always use things like the *sys* or *os* module in there while getting all the basic tools pre-loaded :-)

Hold a second! Yes, up there I wrote something about platform-independence. But isn't this the job of autoconf and not Make? Absolutely true, but SCons also comes with some of autoconf's functionality like [checking for headers](http://scons.org/doc/0.97/HTML/scons-user/x2827.html) and [checking for functions](http://scons.org/doc/0.97/HTML/scons-user/x2836.html).

## I <3 Examples

But back to the real topic here: How to compile stuff with it. Let's say you have a small C program that does nothing really special and therefor doesn't require any external library. Let's name it *hello.c* for now and put following into it:
    
    #include <stdio.h>
    
    int main(){
        printf("Hello, world!\n");
        return 0;
    }
    
Our first SConstruct file could then look somehow like this:
    
    import os
    env = Environment(ENV=os.environ)
    conf = Configure(env)
    if not conf.CheckCHeader('stdio.h'):
        print "stdio.h not found"
    env = conf.Finish()
    Program('hello.c')
    
As you can see here, there is nowhere a call to the gcc. That is because low-level calls are totally abstracted thanks to the built-in builder "Program", which recognizes hello.c as a sourcefile and compiles it into an executable *hello*. Now simple run scons in this directory and you should get this output:
    
    $ scons
    scons: Reading SConscript files ...
    Checking for C header file stdio.h... yes
    scons: done reading SConscript files.
    scons: Building targets ...
    gcc -o hello.o -c hello.c
    gcc -o hello hello.o
    scons: done building targets.
    
Not only did scons compile your code, but it also checked for
the requested header file. That's probably as easy as it can get. But how does it now make compiling PDFs from LaTeX any easier?


## Back to LaTeX and PDFs

Well, scons doesn't only come with built-in builders for C and C++ code, but also for LaTeX and TeX. So you create a SConstruct file like this and get away with it:
    
    import os
    env = Environment(ENV=os.environ)
    pdfOutput = env.PDF(target='presentation.pdf',source='presentation.tex')
    Depends(pdfOutput,Split('presentation.tex bibliography.bib'))
    
This simply creates a PDF out of the presentation.tex. But it also does something new here. Note the last line. There I give the pdfOutput task not only the presentation.tex as dependency, but also the bibliography.bib file. This way, whenever one of these two files gets changed, the task can get implicitly triggered.

## Resources

* [Project page](http://scons.org)
* [Download page](http://scons.org/download.php)
* [Manual](http://scons.org/documentation.php)


[SCons]: http://scons.org/