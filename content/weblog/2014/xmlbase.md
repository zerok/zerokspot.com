---
date: '2014-08-22T21:45:47+02:00'
language: en
tags:
- xml
title: xml:base
---


As part of my rewrite of this blog I also wanted to get rid of having to put
absolute links to images and other post-local files into every rendering of a
post. While trying to achieve that for the feeds I stumbled upon a neat little
feature of XML that made this possible.

If you remember the HTML3 and HTML4-days you probably also remember the `<base>`
tag appearing virtually everywhere due to the amount of static websites. It
allows to set a base URL for every relative link or reference *within the
current document*. XML has that too but in a far more useful way: You can set
such a base path not just for the whole document but *for every node* with the
[xml:base][1] attribute:

```
<container xml:base="http://zerokspot.com/">
    <img src="image.png" />
</container>
```

---------------

If the `img[src]` attribute is interpreted as a link, then the resulting path
here is `http://zerokspot.com/image.png`.

With the rise of HTML5 replacing XHTML you might think that this doesn't work
anymore, but if you serve your [HTML5 document as XML][2] then this should work
here as well according to the specs. I didn't try that, though, since I wanted
to stick with the non-XML version of HTML5. Luckily, I only have one page that
serves the content of a post without being in the blog post's own path: The
index page. And here I just set the classic `<base>` tag for the whole page.

As I wrote above, my initial use-case for this feature was to resolve relative
links contained in my blog posts within the context of each such post. So, if
you look at the main feed right now, you will see something like this:

```
<entry xml:base="http://zerokspot.com/weblog/2014/07/14/knack-review/">
    <title>Knack reviewed</title>
    <updated>2014-07-14T18:33:33Z</updated>
    <id>
    http://zerokspot.com/weblog/2014/07/14/knack-review/
    </id>
    <content type="html">
        <!-- ... -->
        <img src="screenshot.jpg">
        <!-- ... -->
    </content>
</entry>
```

So far I've only tried that with Feedly and Firefox' feed renderer and there it
worked right away.

[1]: http://www.w3.org/TR/xmlbase/
[2]: http://www.whatwg.org/specs/web-apps/current-work/multipage/dom.html#the-xml:base-attribute-(xml-only)