---
title: "Turbolinks and Hugo"
date: 2017-08-17T13:41:29+01:00
tags:
- hugo
---

I'm currently working on a little side-project
called [StrangeNewPlaces.org](https://www.strangenewplaces.org/) where I'm
using [Hugo](http://gohugo.io/) as the main content-framework. That being said,
I wanted to have a few dynamic areas on that page that should be loaded via
JavaScript and eventually using React.

But having your whole React-app being re-initialised on every page-load would be
really bad. Luckily, a couple of months ago I stumbled
onto [Turbolinks](https://github.com/turbolinks/turbolinks) thanks
to [this great post](https://changelog.com/posts/why-we-chose-turbolinks) on
changelog.com. With this I hoped to keep for instance an account-box unchanged
while the rest of the content is loaded via AJAX and inserted the DOM.

Turns out, integrating Turbolinks is extremely easy but nearly immediately I
still ran into a small problem. Whenever I navigated from one page to another, I
was greeted by a blank page and something like this showing up in the browser's
JS console:

```
turbolinks.js:5 Uncaught TypeError: Cannot read property 'querySelector' of null
    at n.t.SnapshotRenderer.n.findFirstAutofocusableElement (turbolinks.js:5)
    at n.t.SnapshotRenderer.n.focusFirstAutofocusableElement (turbolinks.js:5)
    at turbolinks.js:5
```

Sadly, I couldn't find any usable information about this so I decided to
experiment with various Hugo flags. Seems like the live-reload feature is
injecting some JavaScript into the page that prevents Turbolinks from working
properly. If I just start Hugo with `hugo serve --disableLiveReload` everything
works fine.

That's not an ideal solution by any means, but for now it's enough 🙂
