---
title: Creating my own personal information management system
date: "2021-01-19T19:27:17+01:00"
tags:
- pim
- orgmode
- zettelkasten
- emacs
- productivity
- 100daystooffload
---

After reading [“How to take smart notes” by Sönke Ahrens](https://takesmartnotes.com/) I somehow wanted to give the whole personal information management thing another try. Just like many others I’ve used Evernote in the past but then rather quickly moved on to just put stuff into files and folders on a Dropbox. The problem with all of that was, as always: Naming. How do you name your stuff or even worse: how to do you group and organise it? Something about the Zettelkasten system with just giving notes numbers instead of names and creating your indexes as you needed them was extremely appealing to me.

Perhaps I’m just the kind of person who prefers many small files over few large files with good internal structure. Who knows…

Anyway, so I set out looking for options but eventually came to the conclusion that I’d want to implement that myself. Giving that beast the name “datasphere" after the information network in various Peter F. Hamilton books, the original idea was to just use Markdown files and organise them into a single folder with various utilities built around for quick retrieval and adding of notes. I use Markdown virtually everywhere but, for some reason though, they it didn’t fit here. 

After a bit of tinkering and thinking (more often than not in that order) I’ve come to the following setup:

- I put all notes into a single folder (with perhaps some buckets being generated later on for performance reasons).
- Every note is put into a single file.
- The filename consists of a random string as prefix (for now the first part of a UUID) followed by a short description of the content. This way I can still learn from the filename what’s in the file but the order is no longer affected. The name just becomes a hint.
- The notes themselves are written in [OrgMode](https://orgmode.org/).
- The folder has a README file which acts as the entry point of the whole repository, linking to “major” topics.
- Index notes act as overview over specific topics, linking to notes about this topic.

Right now I’m in the process of setting that up properly within Emacs. It will most likely consist of at least the following parts:

- Search using ripgrep and rendering the output using helm for quick navigation between the notes
- A custom link type where I can just write `ds:<prefix of filename>` and it will get rendered as a normal link inside Emacs with click behaviour etc. ✅
- Support for multiple root folder so that I can have one datasphere for work and one for personal stuff ✅
- Create a shortcut for creating a new notes within the currently active datasphere ✅
- Create a side-buffer which shows me what other notes link to the one that is focused right now.

I’m already at a point where I can work with it and each datasphere is slowly but steadily growing. It’s still very work-in-progress, though! Once I have also search and navigation nailed and put some polish on the creation shortcut with templates etc. (perhaps I will simply re-use the functionality available in OrgMode here)  I’m pretty sure I will have lots of fun with this system 😀
