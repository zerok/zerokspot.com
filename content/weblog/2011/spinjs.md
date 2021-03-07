---
date: '2011-08-28T12:00:00-00:00'
language: en
link: http://fgnass.github.com/spin.js/
tags:
- javascript
- css
title: 'spin.js: A pure JavaScript Spinner Library'
---


I couple of days ago [Chris Wanstrath linked on Twitter][1] to a new
JavaScript library for creating loading-spinners without any sort of image
being involved. So right after hearing about [spin.js][2] by Felix Gnass I had
to give it a try.

----------------------

I often have to put spinners in places with different backgrounds so a
solution that doesn't involve having to create a separate GIF for each and
every single one of these places definitely gets my attention. And spin.js
definitely shines in this department with no dependencies to other JS
libraries and lots of customizability so that even if you need spinners of
multiple sizes on one page, all you have to do is include spin.js, create a
new instance of the Spinner with a couple of options for each spinner and be
done with it. And if you don't know what all these options mean, [the
project's website][3] has a nice configurator for you.

<figure>
<img src="/media/2011/configurator.png" alt="" />
<figcaption>The configurator on spin.js's project page helps you test
the various options.</figcaption>
</figure>

spin.js does all this by creating multiple div-elements and using CSS
translations and animations to make them rotate. For IE it falls back to VML
objects.

spin.js is most likely not the one-size-fits-all solution right now but it's
getting really close. One downside of this compared to an animated GIF is that
it requires more CPU power especially in IE according to [this issue][4].
Especially on modern browsers it probably just a matter of time until CSS
animations have been optimized, though.

Another advantage of GIFs is that you can easily just use them as background
image (which is probably the preferred way to do it anyway). For instance if
you want to add a spinner right into a button (&lt;input type="submit"&gt; for
backwards-compatibility) using a background image is easier than extending the
padding of the button and absolutely positioning the spinner on top of that
newly created space.

But for everything else I will probably look at spin.js first the next time I
need a spinner somewhere. It looks like a great addition to the toolbox of
every web developer :D


[1]: http://twitter.com/#!/defunkt/status/104223961373618176
[2]: https://github.com/fgnass/spin.js
[3]: http://fgnass.github.com/spin.js/
[4]: https://github.com/fgnass/spin.js/issues/8
