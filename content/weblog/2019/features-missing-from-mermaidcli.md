---
title: "Features missing from Mermaid.CLI"
date: "2019-11-05T16:22:00+02:00"
tags:
- mermaid
- diagrams
---

For pretty much all my documentation tooling I have one Docker image
or another. For instance, there is also one for Sphinx which
integrates [Mermaid][m] to allow drawing diagrams using Martín
Gaitán's [mermaid plugin][mp], which - in turn - uses
[mermaid.cli][mc] (or mmdc). Recently, a user contacted me because she
was facing a weird issue: For some reason, subgraphs didn't seem to
work. She provided me an example file that looked like this:

```
.. mermaid::
   graph LR
       subgraph A
           a((NodeA))-->b(NodeB)
       end
       subgraph B
           c{{NodeC}}-->g{{NodeD}}
       end
       b-->c
```

Putting that into the [live-editor][le] you *should* get this diagram:

<figure>
<img src="/media/2019/mermaid-expected-result.svg">
<figcaption>The diagram should looks like this...</figcaption>
</figure>

She indicated that this *should* be supported in a more recent version
of Mermaid than what our application was bundling and so I set out to
update it. Weirdly enough, now running Sphinx got stuck when
processing that definition. It kind of smelt like a Promise rejection
not being handled properly.

As my next step I tried to track down if the issue was limited to the
Sphinx integration or was there a problem with the underlying library
and CLI? I therefore exacted the diagram definition in its own file,
went directly into the Docker image ran mmdc directly:

```
./node_modules/.bin/mmdc  -i test.mm -o test.svg --puppeteerConfigFile /app/puppeteer.json
(node:647) UnhandledPromiseRejectionWarning: Error: Evaluation failed: Error: Parse error on line 2:
...ph B    c{{NodeC}}
----------------------^
Expecting 'SPACE', 'GRAPH', 'DIR', 'TAGEND', 'TAGSTART', 'UP', 'DOWN', 'subgraph', 'end', 'MINUS', '--', '==', 'STR', 'STYLE', 'LINKSTYLE', 'CLASSDEF', 'CLASS', 'CLICK', 'DEFAULT', 'NUM', 'PCT', 'COMMA', 'ALPHA', 'COLON', 'BRKT', 'DOT', 'PUNCTUATION', 'UNICODE_TEXT', 'PLUS', 'EQUALS', 'MULT', got 'DIAMOND_START'
    at Yt.parseError (file:///input/node_modules/mermaid.cli/mermaid.min.js:1:486512)
```

Huh! Turns out, mmdc cannot deal with hexagons, but the location of
that error was interesting: `mermaid.cli/mermaid.min.js`. mmdc bundles
its own copy of `mermaid.min.js` which is very old. As a quick fix
I've now simply copied the "original" mermaid library into the CLI:

```
$ cp /usr/local/lib/node_modules/mermaid/dist/mermaid.min.js \
     /usr/local/lib/node_modules/mermaid.cli/
```

This is definitely something that should be fixed upstream and so this
"fix" here is basically just a temporary workaround. Luckily, there is
already a [ticket for the mermaid.cli project][t] on GitHub by someone
with a quite similar issue 😉

[mp]: https://github.com/mgaitan/sphinxcontrib-mermaid
[t]: https://github.com/mermaidjs/mermaid.cli/issues/68
[m]: https://github.com/mermaid-js/mermaid
[mc]: https://github.com/mermaidjs/mermaid.cli
[le]: https://mermaidjs.github.io/mermaid-live-editor/
