---
date: 2016-09-24T09:21:29+02:00
title: Link management in AsciiDoctor
tags:
- documentation
- asciidoc
- asciidoctor
---

I'm currently in the process of writing a couple of concept papers and thought I
might give [AsciiDoctor][ad] a try here. Sphinx still feels better for
multi-document projects and using docutils alone is always a bit weird 😉 One of
the features I've grown to love in both, RST and Markdown, though, is that I can
specify a link's URL somewhere else in the document. This allows the actual
source documents to be still very pleasant to read (Markdown example below):

    This is a paragraph where I want to [link][l] to something else but don't
    want the actual URL to disturb the whole flow.

    [l]: https://mydomain.com

While AsciiDoc doesn't have this feature explicitly listed somewhere a quick
question in the [WriteTheDocs Slack channel][sl] and great answers from
[Jared Morgan][jm], [Brian Exelbierd][be], and Jonatan Jäderberg later I had
what I wanted:

    :l: https://mydomain.com

    This is a paragraph where I want to {l}[link] to something else but don't
    want the actual URL to disturb the whole flow.

This uses the attributes feature in combination with substitutions only
mentioned in the writer's guide chapter about
[document attributes][da]... Basically, you define your URL as an attribute and
then reference that where the URL would normally be used. The only aspect about
this that I don't really like is that I have to put these attribute definitions
*before* the actual usage while in Markdown and RST I can put them at the end of
a paragraph or the whole document.

That being said, from the chat we had in the Slack channel it seems like quite a
few people are collecting all links in their own files that they then
include into their documents:

    = My long article

    include::links.txt[]

    ...

This sounds like a great way to deal with share links between multiple documents
but would also make automatically checking these URLs much easier 😊


[jm]: https://twitter.com/jaredmorgs
[be]: http://www.winglemeyer.org/
[da]: http://asciidoctor.org/docs/asciidoc-writers-guide/#document-attributes
[ad]: http://asciidoctor.org/
[sl]: http://slack.writethedocs.org/
