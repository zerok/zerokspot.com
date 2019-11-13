---
title: "What is TaskPaper?"
date: 2019-11-13T16:10:29+01:00
tags:
- til
- productivity
---

Listening to various productivity apps and throughout the OmniFocus UI
a certain name popped up that I didn't know anything about:
TaskPaper. [TaskPaper][t] is a small macOS app that allows you to work
GTD-style in a plain-text format. That format is defined [here][h] but
let's look at a small example: Let's say, I want to work on a project
called "Write a blogpost about TaskPaper" with some tasks:

1. Find out what TaskPaper is
2. Find out what's different to something like todo.txt
3. Find out how the format is integrated in OmniFocus

In TaskPaper's format the project would look somehow like this:

```
Write a blogpost about TaskPaper:
- Find out what TaskPaper is @wohnung
- Find out what's different to something like TodoTXT @wohnung
- Find out how the format is integrated in OmniFocus @wohnung
```

Projects are defined with lines that end with a colon, tasks below
such a line are associated with the project and have to start with a
dash. The example above also defines tags for the tasks. A tag is
indicated by an at-sign and optionally can also have a parameter which
is useful for things like due-dates or priorities:

```
- This has to be done by tonight @due(2019-11-14 22:00)
- This has a high priority @priority(1)
```

Basically, if you've ever seen a list in a plain-text file, you can
understand TaskPaper lists.


## What about todo.txt?

Another plaintext format for TODOs that has been going around for many
years now is [todo.txt][tt]. Here all information about a task
including the project's name is stored in the task's line:

```
X Find out what TaskPaper is +"Write a blogpost about TaskPaper" @wohnung
```

This has the advantage that you can just parse a single line and you
know everything there is about a task, including what project it is
associated with. The downside is that you have tons of
info-duplication which becomes annoying when you have projects with
lots of tasks and just want to fix a typo in the project's name. When
I was using this system this aspect of it led me to favor short
project names that left out sometimes crucial information.

TaskPaper is better here simply because you don't have to duplicate
project names. This makes those files more readable even when you
don't have any kind of tooling support. On the other hand, todo.txt
allows you to associate tasks with multiple projects. I haven't had a
use-case for that so far, though. I also like the idea what tags can
be parametrized so that things like priorities are not special cases:

```
# In TaskPaper
project:
- some task @priority(1)

# In todo.txt
(A) some task +project
```


## TaskPaper in OmniFocus

OmniFocus [starting with version 2.7 on macOS and 2.14 on iOS][s] supports
exporting single tasks or projects in a TaskPaper compatible
format. All you have to do is to right-click whatever item you want to
export and select "Copy as TaskPaper" to get a TaskPaper-formatted
version of that item in your clipboard.

Going with the example from above, OmniFocus would export it like
this:

```
- Write a blogpost about Taskpaper @parallel(false) @autodone(false)
	- Find out what TaskPaper is @parallel(false) @autodone(false) @context(Wohnung) @tags(Wohnung) @due(2019-11-14 22:00)
	- Find out what’s different to something like todo.txt @parallel(false) @autodone(false) @context(Wohnung) @tags(Wohnung)
	- Find out how the format is integrated in OmniFocus @parallel(false) @autodone(false) @context(Wohnung) @tags(Wohnung)
```

OmniFocus, for some reason, exports the project not as project but as a
nested tasklist. A bit unexpected but that shouldn't be a big issue
when integrating it into your workflow.

Importing works in a similar way: Just put a TaskPaper-formatted list
into your clipboard and click CMD+V inside OmniFocus. The new projects
and tasks should appear in your inbox.


[t]: https://www.taskpaper.com
[h]: https://guide.taskpaper.com/getting-started/
[s]: https://support.omnigroup.com/omnifocus-taskpaper-reference/
[tt]: http://todotxt.org/
