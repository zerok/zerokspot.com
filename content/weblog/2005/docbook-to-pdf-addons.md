---
date: '2005-04-26T12:00:00-00:00'
language: en
tags:
title: Docbook to PDF? Addons
---


Just some small addons to my <a href="http://weblog.zerokspot.com/posts/330/">Docbook to PDF?</a> article:

-------------------------------

<ul>

<li>In the output PDF I had some not really requested strings like "--4pc" which seem to have been ignored by pdfxmltex as input so it printed it to the output. If you pass "--stringparam passivetex.extensions 1" as additional argument to xsltproc when generating the .fo file, this output should be removed. Thanks to <a href="http://lists.oasis-open.org/archives/docbook-apps/200309/msg00199.html">Justus Piater</a> for this tip :-)</li>

<li>Another problem I noticed is the width of the header and footer section which seems to be more or less just half the page width. After some googling I found <a href="http://lists.oasis-open.org/archives/docbook-apps/200302/msg00338.html">this mail by Gary Lawrence Murphy</a> describing the same problem. In this thread someone also mentioned a possible solution for this, but I had no time to try it yet.</li>

<li>In the case of variablelists there also seem to be some fragments of the FO-element flying around. Depending on motivation and time I will also try to find a solution for this one.</li>

<li>Now that my docbook grows and grows I also multiplied the save_size by 10 again...</li>

</ul>