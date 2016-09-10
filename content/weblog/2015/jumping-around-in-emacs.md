---
date: '2015-01-07T20:37:17+01:00'
language: en
tags:
- emacs
title: Jumping around in Emacs
---


I’m often in a situation where I start writing a function and notice that I’d
like to use a helper function that doesn’t exist yet. When this function should
be created in a different file, no problem: Open a new buffer, visit and edit
the file there and kill the buffer again. Emacs will restore me to the previous
point in the other buffer.

But what about when I want that function to exist in the same file and I want to
jump right back where I left off once I have that function? Emacs supports this
for instance through the so-called [mark][m]. You’ve probably already used the
mark before when you selected text in order to change or remove it...

-----------------------

## The Mark

Whenever you create such a region you basically operate between two positions:
the pointer (where your cursor ends up being when you’re done) and the mark
(usually where the cursor was at the start of the selection). During the
selection process the mark is “active” and you see the region being
highlighted. Whenever you create a mark it is added to a ring of positions
called the “mark ring”. Setting the mark is done with the `C-<SPC>` shortcut (or
the `set-mark-command` command) which starts the selection process. But since we
don’t want to create a selection here, we can deactivate the mark right away
again hitting `C-<SPC>` once more.

So now we have our mark on the mark-ring and it is not active, so changing your
position inside the buffer won’t create a region (and potentially highlight
it). After moving around for a while (for instance to create that other function
we wanted to work with at the previous position) we want to get back were we
left off. The previous position is just one `C-u C-<SPC>` away!

```
;; Create a mark and deactive it right away
C-<SPC> C-<SPC>

;; Jump back to the head of the mark-ring and pop it afterwards
C-u C-<SPC>
```

Always keep in mind that the mark is also used for creating
regions/selections. So whenever you select text, the current head of the
mark-ring is changed and therefore jumping back might return you to a position
you didn't necessarily expect. Just hit the jump-back key-combo repeatadly ;-)


## Register storage

Because of that I’ve also started to look at the
[point-to-register and jump-to-register][r] functions which allow you to store
positions outside of the mark-ring inside a register with a single character
name.

```
;; Set a new mark
C-x r <SPC>

;; Jump to a register
C-x r j
```

Especially the key-map for the setter felt rather tedious to me so I've remapped
it to something that is closer to the one for setting the mark:

```
(global-set-key (kbd "C-c C-<SPC>") 'point-to-register)
```

You have to enter the name of the register the position should be stored in so
the key-chain becomes a little longer than for simply using the mark. Because of
that I've also experimented with having another shortcut that used these methods
on a specific register but quickly abandonned that approach because it ended up
to confusing due to the position not being restricted to a single buffer.


## Bookmarks

Another option I stumbled across was using [bookmarks][bm]. They kind of feel like a
more permanent version of `point-to-register` in that all your bookmarks are by
default stored into `~/.emacs.d/bookmarks`. You can give those bookmarks
arbitrary names when you create them. By default the name of the current file is
used which should be enough for my initial use-case but it might become
confusing if you have multiple files with the same name in your project.

```
;; Set a bookmark
C-x r m

;; List bookmarks
C-x r l

;; Jump back to a bookmark
C-x r b
```

I haven't yet had a look at [Bookmark+][bp] but it's better default-name
handling might be a good reason to do that in the future :)


## Now what?

With these three options Emacs solves my initial use-case quite nicely out of
the box. I'm not yet sure which approach I'll eventually prefer but I'd guess it
would either be the register-storage or bookmarks. Especially when working with
a mouse messing up the mark-ring seems far too easy to me.

I'm not yet sure if bookmarks are the right thing either but I definitely have
other use-cases for them!

[bp]: http://www.emacswiki.org/emacs/BookmarkPlus

[bm]: https://www.gnu.org/software/emacs/manual/html_node/emacs/Bookmarks.html

[r]: http://www.gnu.org/software/emacs/manual/html_node/emacs/Position-Registers.html#Position-Registers

[m]: http://www.gnu.org/software/emacs/manual/html_node/emacs/Mark.html#Mark
