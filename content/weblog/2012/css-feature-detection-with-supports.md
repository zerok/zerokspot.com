---
date: '2012-11-24T12:00:00-00:00'
language: en
link: http://dev.opera.com/articles/view/native-css-feature-detection-via-the-supports-rule/
tags:
- css
- opera
title: CSS Feature Detection with @supports
---


Just stumbled upon a [nice article][art] by [Chris Miles][cm] about a new feature in CSS3 that might in the future offer an alternative to doing *feature detection* with [Modernizr][mod].

Modernizr is probably the first 3rd-party library I add to every single web project at the beginning. Way before jQuery. But most of the time I only use it for detecting CSS-feature-support in browsers. Wouldn't it be nice if CSS itself offered some kind of feature-detection functionality built natively right into the browser? Esp. if you mostly have a static page but just want to use some nice modern CSS features, adding Modernizr to the mix (a) hits your site's performance a little bit and (b) simply kind of feels redundant. Surely, you can get pretty far by letting the cascade and the fallback-behaviour of CSS work for you, but the code gets unmaintainable pretty quickly.

[CSS3 now offers a @supports conditional group rule][w3c], which you can use to check for certain declarations (e.g. "border-radius: 5px" or "display: flex") and only process the associated group if the feature is supported.

<pre><code>
section { /* create border-radius with images */ }
@supports (border-radius: 5px) {
    section {border-radius: 5px;}
}
</code></pre>

What's really nice about *@supports* is the preciseness with which features should be detectable here. Checking for "display: flex" becomes trivial here. And you get around [the whole mess that CSS feature detection is with JavaScript][jsm].

I doubt *@supports* will replace Modernizr for me in the near future, though, simply because it is not "officially" available in most browsers (Chrome >= 24, Firefox >= 17, Opera >= 12.1 according to [MDN][mdn]), but this is hopefully just a matter of time. In the meantime Modernizr will hopefully include *window.supportsCSS*-based detection soon :-)

Another really nice introduction about this topic (this time with a focus on Firefox instead of Opera) is available on [Peter Gasston's blog][pga]

[art]: http://dev.opera.com/articles/view/native-css-feature-detection-via-the-supports-rule/
[cm]: http://my.opera.com/chrismills/about/
[mod]: http://modernizr.com/
[jsm]: https://github.com/Modernizr/Modernizr/blob/master/feature-detects/css-fontface.js
[mdn]: https://developer.mozilla.org/en-US/docs/CSS/@supports
[w3c]: http://dev.w3.org/csswg/css3-conditional/#at-supports
[pga]: http://www.broken-links.com/2012/08/06/firefox-supports-supports-gets-my-support/
