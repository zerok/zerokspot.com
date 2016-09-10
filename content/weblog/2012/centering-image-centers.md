---
date: '2012-04-15T12:00:00-00:00'
language: en
tags:
- javascript
- css
- webdesign
title: Keeping image centers in the center
---


If you ever need to keep an image's center centered on a website no matter
what viewport size your browser has, there is a [neat little
trick](http://stackoverflow.com/a/1344184/22312) I stumbled upon on
StackOverflow provided by [Amber
Yust](http://stackoverflow.com/users/148870/amber). The goal is that the image
is cropped not only on the right but also on the left side  to keep the center
of the image also in the center of the screen. The appoach mentioned by Amber
is to position the image absolutely with an offset of 50% from the left and
then re-adjust this offset using a negative margin of minus half the image's
actual width.

----------------------

<pre><code>
#container {
    position: relative;
}
#container img#demoimage {
    display: block;
    margin: auto;
    position: absolute;
    left: 50%;
    margin-left: -600px;
/* If 600px is half of your image's width */
}
</code></pre>

This has one caveat, though: You have to know the dimensions of the image in
advance or determine them via JavaScript and adjust the negative margin after
the fact.

<pre><code>
(function(img) {
    // This will get called, once the image has actually been
    // loaded from the server
    function updater() {
        img.style.position = "absolute";
        img.style.left = "50%";
        img.style.marginLeft = "" + -img.width / 2 + "px";
        console.log(img.style.marginLeft);
    }
    img.addEventListener('load', updater, false);
})(document.getElementById("demoimage"));
</code></pre>

If JavaScript is disabled, the margin:auto up there will at least center the
image if there is enough room for it and prevent a slight jumping *if*
JavaScript is enabled and the repositioning takes place.

For IE you will have to add some extra code to ensure your event is fired even
then the image is already available within the browser's cache, but that's
mostly it.

Another approach to this issue, which I usually prefer, is to have the image
itself be situated in the background of the controller and positoned in its
center. The upside of this solution is, that you don't need to position the
image using JavaScript if you know its height to make room for it in the
container. A downside here is, that without JavaScript this won't work if there
is an imagemap on top of the image (which was the case and the original reason
why I was looking for a different solution).
