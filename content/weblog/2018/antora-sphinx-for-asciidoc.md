---
title: "Antora: Sphinx for Asciidoc"
date: 2018-03-25T22:34:37+02:00
tags:
- documentation
- asciidoc
---

For most projects, both at work and in my sparetime, I'm using a mix of
Markdown and Sphinx/ReStructuredText for creating READMEs, reference
documentation, and topic guides. While this setup works for me, I'm not the
biggest fan of some aspects of ReStructuredText and have therefore started
looking more and more at asciidoc. The downside this is that, though, I hadn't
found any framework similar to Sphinx that allows to work with a collection of
asciidoc-documents. Well, and then I stumbled upon [Antora][] this week.

While Sphinx is focused on taking a specifc folder full of rst files and
converting it to HTML, TeX, PDF, what-have-you, Antora is a more generic
framework for mixing documentation stored in multiple locations and converting
it to HTML. Out of the box it supports building documentation for multiple
releases of a project based on tags, branches, and so on.

The output is based on a playbook defining what modules should make up the
final documentation and how they should be converted into the desired output
format. Each module can define its own navigation and in turn can contain
multiple components with pages and assets like images. The modules referenced
in a playbook don't have to reside in the same code repository. For instance,
there could be one team working on the user guide in their own repository while
another team focuses on the developer guide in another.

Playbooks now instruct the converter what UI bundle (zip files with markup
templates, stylesheets etc.) to use to generate HTML out of these modules and
what navigation should be rendered on the final output.

Coming from Sphinx it definitely took me some time getting comfortable with
this structure. In general, Antora by default expects a certain file system
structure (detailed in [Organizing Your Documentation for
Antora](https://docs.antora.org/antora/1.0/component-structure/) and right now
can only deal with "packaged" bundles. For the latter there already exists [a
ticket](https://gitlab.com/antora/antora/issues/150) on Gitlab 🙂

Antora in general looks despite its young age extremely useful and I'm pretty
sure I will integrate it in one of my next projects. I already tried to convert
one with only a couple of pages and really liked the resulting structure.
Perhaps we finally have a real competitor to Sphinx in the market 😄

[Antora]: https://antora.org/

