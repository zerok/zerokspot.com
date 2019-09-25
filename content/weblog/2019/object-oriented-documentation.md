---
title: Object-oriented documentation
date: 2019-09-25T18:49:24+02:00
tags:
- writethedocs
- documentation
---

On the first day of [WriteTheDocs in Prague][wtd] last week, [Luc
Perkins][lp] gave a [lightning talk][lt] about a workstyle he called
"object-oriented documentation". When following this method you aim to
create documentation in re-usable components that can then be mixed
and merged. Ideally, you'd this way be able to reuse them in
completely different output formats like tutorials, blog posts,
reference guides, and so on.

In the talk, Luc gave the example of "externalizing" concept
definitions that can then be reused wherever necessary, something you
might normally associated with a glossary (for instance,
[Sphinx][sg]). On second thought it's more akin to the idea of
something like [assemblies in DocBook][as], where you can combine
"topics" into completely new documents.

In general, though, I feel like my first hunch regarding
"glossary-like" content might be the most practical use-case for this
style of documentation. Pretty much anything that eventually
might be used as an aside, tooltip, or similar in a larger document
would be - for me - a prime candidate for this style of writing.

Systems like Hugo or Jekyll support different types of content. In
Jekyll these are called [collections][jc] while Hugo has [content
types][hc]. At first I thought that the concept Luc described would
perhaps better be called "component-based documentation" but that
would ignore that we are actually talking about various "classes of
content" here. Glossary-entries are just one class here while another
could be "screenshots", "keyboard shortcuts", and so on. Combined with
shortcodes or `$.Site.Data` calls these could either be embedded or just
served as standalone content.

I would love to see that being applied somewhere on a larger scale,
though. If you know of a public repository that uses this
documentation style, please let me know!


[wtd]: https://zerokspot.com/weblog/2019/09/18/writethedocs-prague/
[lt]: https://www.youtube.com/watch?v=axBTqOslb54&list=PLZAeFn6dfHpkpYchP1iFnQnc7i-2xJd0I&index=11&t=0s
[sg]: https://www.sphinx-doc.org/en/master/usage/restructuredtext/directives.html#glossary
[as]: https://tdg.docbook.org/tdg/5.1/assembly.html
[jc]: https://jekyllrb.com/docs/collections/
[hc]: https://gohugo.io/content-management/types/
[lp]: https://github.com/lucperkins
