---
date: '2013-09-11T23:07:10+02:00'
language: en
tags:
- tool
- writing
- blogging
- markdown
- editorially
title: Writing Markdown with Editorially
---


[Markdown][md] has been around for ages but, all of a sudden, there is quite a surge in attention for this lightweight markup format. OK, not really all of a sudden, but more over the course of the last 1-2 years perhaps thanks to tools like iWriter and static blog engines like Jekyll. And since February there is a new tool available for everyone who like the format which has also just become an [open beta][beta]: [Editorially][edi].

---------

What makes Markdown so great is easily explained: It is a simple syntax for formatting plain-text files that doesn’t distract you from *what* you’re going to write by giving you tons of options on *how* to write. More and more people seem to finally jump onto that bandwagon thanks to tons of new tools that make it easier and easier to work with it (even more so than it was before!) . There is even a short book out there [introducing Markdown to writers][bk] (no idea if it is any good).

This simplicity is precisely why I switched to writing mostly exclusively in Markdown a couple of years ago: Being able to focus on the content instead of the style (which is also why I chose LaTeX for writing my master thesis back in the days ;-)). 

## Tool-chain over the years

Over the years I’ve worked with [VIM][], [TextMate][tm], [Sublime Text 3][slt] to write my posts with or without tools like [Marked][ma] to provide some great preview support. All of them make up for quite a powerful set of tools with features like snippets, auto-bracketing, syntax highlighting and real-time previews, but there was always something missing.

## It is all about drafts

Don’t get me wrong. I’ve been getting tons of writing done over the last couple of months esp. thanks to the combination of Sublime Text and Marked, but there was always one problem: *How to deal with drafts?*

At first I just put them into a folder called “drafts” alongside my normal posts. Which is pretty much where they ended up being until I accidentally deleted that folder and was not motivated enough to restore them from a backup simply because the ideas in there weren’t all that great to begin with.

Then, thanks to [Shawn Troop][st], I gave [iWriter][iw] by [Information Architects][ia] a try. The great thing here is, that it easily integrated with Dropbox, so it is (1) much harder to delete the drafts order by accident and (2) I can improve them no matter where I am. The downside of iWriter was that I didn’t really enjoy the writing experience all that much. Also, the document management in Dropbox was kind of clumsy.

Next came [Evernote][en]. Evernote, by itself, is a great tool. It helps immensely to know where you dumped all your research on a topic before you started writing about it. But for writing articles or anything longer than a tweet the UI is just too noisy.

## Enter Editorially

And this is where Editorially comes in. I’ve just discovered it a a couple of days ago but so far it looks like exactly the kind of authoring tool I’ve been looking for to flesh out my drafts. The UI is simply and clean, I can use it from anywhere where I’m online and it offers a simply document management feature which helps me keep an overview of all my drafts. It supporting Markdown is nice, too ;-)

## Core editing features

And at that it is really great. Headings and other inline formats defined by the original Markdown specification are styled in a really nice way and that's all I want anyway. Embedded images are a nice touch but anything more fancy rendering-wise would simply be too much.

Speaking of embedding images: If you like to that in your posts you still have to combine it with another host or tool since Editorially doesn’t host anything but your text content. In my case this other tool would be Evernote together with [Skitch][sk] simply because I already use them for virtually anything else anyway.

## Collaboration

Editorially being online makes another feature (outside of being available whenever I might need it) possible: You can invite other people to help you with your drafts, for instance by commenting on specific sections. I haven’t tried that one yet but it sounds far more useful than sharing your Evernote-post and the communicating view e-mail where you usually end up forgetting about half of the comments.

You can also directly collaborate on articles by making someone else an “editor” which includes edit-permissions.

## Versioning

To document all these changes that were made because of comments or simple refactoring of a draft Editorially also supports versioning your documents. 

<figure>
    <img src="Version%20in%20Editorially.png" alt="">
    <figcaption><p>Navigating versions</p></figcaption>
</figure>

For instance, your post’s history could look like this:

1. Initial draft with structure
2. Added section about drafts
3. Cleaned up conclusion

You can then compare versions in order to find out in detail what you’ve changed. This is nothing new if you’ve used Google Docs before or just manage files with versioning tools like Mercurial or Git, but it is an essential feature for me so I didn’t want it to go unmentioned. The implementation here is also very simple in the way that you just have to hit a button and your version is saved, compared to the usual `git commit -a -m “...”` dance when working with something like Git.

<figure>
<img src="Changed%20in%20Editorially.png" alt="" />
<figcaption><p>Changes are color-coded</p></figcaption>
</figure>

Alongside the document revisions you can also set a status for your document, moving it over time from “draft” over “"reviewing” and other stati all the way to the “final”.

<figure>
    <img src="Document%20status%20in%20Editorially.png" alt="" />
    <figcaption><p>Setting your document’s status</p></figcaption>
</figure>

This status is also shown in the document manager (color-coded and/or as a separate column if you choose the list-view).

## Perhaps finally the right tool

I’ve just started to mess with Editorially and there is much left to explore (as the section about collaborating above indicated), but so far I really like it. The idea right now is to use Editorially for new posts until I’m done with them content-wise and then move over to give them their final touches in Sublime Text and Marked. I simply prefer to have images formatted in a slightly different way then is common in Markdown itself.

As a next step I will hopefully get a chance to try the collaboration-features with an internal technical article I want to write. I’m really curious how this will work out compared to something like Confluence or Google Docs.

## Wishlist

Editorially is in beta right now and even though the product is already awesome, I still have a couple of features wishes on my list:

* Version-timeline should be easier to access, right now you have to make one click, move out of the drop-area of the menu and then you finally see it.
* The support for the iOS version of Safari could use some more love esp. when it comes to cursor positioning, which gets messed up a little bit too easily.
* A drop-down or something similar that allows to navigate between sections of a document.
* I really like Markdown’s link-reference syntax where you can refer to a link-location (`[label][ref]`) but sometimes I forget to also provide a mapping of a reference to an actual URL. Here Editorially could provide a check and highighting if it detects a missing reference.
* This kind of brings me to my last feature request for now: A manager for references, which autocompletes them if you start writing the respective syntax for it.

[edi]: http://editorially.com
[bk]: https://www.smashwords.com/books/view/342055
[en]: http://evernote.com
[tm]: http://macromates.com/
[vim]: http://vim.org
[slt]: http://sublimetext.com/
[md]: http://daringfireball.net/projects/markdown/syntax
[beta]: http://blog.editorially.com/post/60769345200/doors-are-open
[ma]: http://markedapp.com/
[ia]: http://ia.net/
[iw]: http://www.iawriter.com/
[st]: https://alpha.app.net/shawnthroop
[sk]: http://evernote.com/skitch/
