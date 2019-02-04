---
title: "Tasklogs & Test-Reports"
date: 2019-02-04T19:11:23+0100
tags:
- productivity
- workflow
- orgmode
---

Whatever I do at work, one of my most important aspects and goals is
that I do it in a reproducible manner. I try to maximize the
papertrail I leave behind for two reasons:

* Whenever someone asks me down the line why I did something in a
  specific way, there is at least a good chance that I might have left
  a note somewhere about that.
* After testing the work of someone else it's helpful to be able to
  show what and how you tested it. This is especially useful if a bug
  with a feature is reported later on and you can check if the path to
  that bug has even been tested before.
  
In order to archive, whenever I start working on a task I create a
file called a **tasklog**. For instance, if the task at hand has a
JIRA ticket inside our tracker by the name of `PROJ-123` then I create
a `~/CompanyFileshare/MyFolder/tasklogs/PROJ-123.org`. What's in that
file depends on if I'm implementing or testing something. Both,
though, tend to include at least some code snippets or commandline
calls.


## Literate programming

In order to keep those snippets correct and reproducible, I'm using an
approach called *Literate Programming* which goes back to Donald
E. Knuth (1984). The idea here is that you primarily write prosaic
text and embed your source code in it. This is the inverse of how you
normally handle comments inside program code. You describe what you're
doing and then do it.

The final document has to go through a pre-processor which produces
actual source code that can be executed. For my use-cases, I describe
what I'm doing, why I'm doing it, and then add either the example code
or the command-line calls that I use to test the implementation.

When it comes to implementing literate programming, there is a huge
pool of software to choose from. For me, personally, the following
tools have worked very well in the past:

* [Jupyter Notebooks](https://jupyter.org/) is a web application that
  lets you write Markdown and embed source blocks between
  paragraphs. You can then hit a button and all the code in there is
  executed and the result is printed below the source code inside your
  browser. This approach also allows you to include diagrams and
  things like D3-based graphics inside the resulting document.
* [OrgMode](http://orgmode.org/) Babel is a package for Emacs that
  allows you to embed source blocks inside of Org documents and
  manages the execution of that code on your local machine. OrgMode
  allows you to also export the resulting document into various
  formats (e.g. HTML pages).

As you have probably already guessed from the filename used above, I'm
mostly using OrgMode. I simply prefer to work inside a text editor
compared to using a web-interface but your mileage *will definitely
differ*. Luckily, there is an export package available for producing
Jupyter Notebooks out of OrgMode files available on
<https://github.com/jkitchin/ox-ipynb>, just in case you want to
decide later to switch from one to the other.

The actual implementation isn't important, though. As long as you have
a system that allows you to embed executable code inside a normal text
document and the result can easily be viewed by others, you're
good. OrgMode is just a Markdown-like text file while Jupyther
Notebooks can easily exported into standalone HTML files.

So, what is now included in those tasklogs?


## Implementation notes and executable documentation

When I'm implementing something, I usually end up having a **TODO list
at the very top** of the file where I collect all the little bits and
pieces that I still have to do before I can open a pull-request for
the rest of the team.

I also include some kind of **motivational premise** into the header
where I explain in my own words what the outcome of the task should
be. This way I try to ensure that the acceptance criteria are clear to
me. If appropriate, this will also eventually find its way into the
project's documentation.

The main part of the **implementation notes** then includes
detail-descriptions on the problems I'm running into and how I'm
solving them. With these these three parts the tasklog for PROJ-123
could look like this:

```text
** Premise

In this task we are trying to solve abc by doing def.

** TODOs

- [ ] Implement test-case A
- [ ] Refactoring implementation of B

** Notes

*** Reasoning for going with approach C

...
```

There is usually only very little code or automated steps inside the
implementation notes except for things like examples.


## Test-report

More interesting, though, is using *Literate Programming* for creating
test-reports. Again, I start with a **short premise** what the task is
all about. Next, I formulate **a number of test-cases** followed by a
**summary** detailing the test outcome including the found issues (if
any). The structure of PROJ-123.org would look like this:

```text
** Premise

The goal of this ticket was to solve ...

** Test cases

*** TC1: Path a

...

*** TCn: ...

** Summary

Tested successfully!
```

Embedding source code works by using `SRC`-blocks:

```
#+BEGIN_SRC shell
echo hello world
#+END_SRC
```

When I hit `C-c C-c` inside the block, the code will get executed and
the result written below:

```org
#+BEGIN_SRC shell
echo hello world
#+END_SRC

#+RESULTS:
: hello world
```

In order run all the code snippets inside a document and thereby
re-run all my tests, I can hit `C-c C-v b`
(`org-babel-execute-buffer`). When I export the file to something like
Markdown, all these code blocks are also executed.


## Sharing the output

So far, we as a team haven't yet decided if or how these test-reports
should be shared. I've personally experimented with adding them
directly to the project's source repository inside the `test-reports`
folder and also adding the source of the OrgMode file as comment to
the task. I'm not sure which one of these if any will stick but if it
does then we should probably also provide a simple web-interface for
easier retrieval outside of the terminal or have JIRA somehow process
these comments for a nicer rendering output respectively.

In any case, this is something I might look into down the line, but
not now. For now, I hope this little guide will help you. Literal
programming has worked tremendously for me in the context of the
use-cases listed above and helped me work more effient. I hope it will
help you too!
