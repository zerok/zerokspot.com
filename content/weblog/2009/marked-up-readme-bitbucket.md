---
date: '2009-03-08T12:00:00-00:00'
language: en
tags:
- opensource
title: Marked up README on bitbucket
---


One of the features I always liked about github was that it rendered you
README files of various formats pretty nicely when accessing the root-folder
of your project. README.rst got some rst2html-love and README.md was converted
from Markdown to HTML. When I first gave [bitbucket][] a try some months ago I
noticed, for the first time, how much I actually had grown to expect my
READMEs to get rendered as HTML, right when I couldn't find out how to get the
same behaviour on bitbucket. 

Well, there is actually a little hidden feature on the hg-hoster that lets
you do the same: Simple put your markup language in an encoding-style preamble
of your README-file and bitbucket will render it for you as HTML on the
start-page of the source-browser::
    
    -*- restructuredtext -*-

    If you put some RestructuredText here, it will get rendered as HTML.

Thanks, jespern, for this great tip :-)

So far I've tried it with "restructuredtext" and "markdown" and both seem to
work. There are just some things to note here:

* This preamble gets only interpreted if it is in the very first line of your
  README-file.

* That first line has to start with the preamble so currently you can't really
  hide it within a comment or something like that. If you want to see this
  change implemented, perhaps giving [this issue][] a little hug will help.

[this issue]: http://bitbucket.org/jespern/bitbucket/issue/423/rest-readme-should-look-for-modeline-in-first-2
[bitbucket]: http://bitbucket.org/
