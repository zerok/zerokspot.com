---
date: '2009-08-10T12:00:00-00:00'
language: en
link: http://cufon.shoqolate.com/
tags:
- cufon
- fontreplacement
- javascript
- webdesign
title: "Cuf\xF3n: Font-replacement made easy"
url_title: "cuf\xF3n - fonts for the people"
---


For quite some time now I've (at least with half an eye) been looking around for a font-replacement technique that doesn't require Flash, and [Cufón](http://cufon.shoqolate.com/) by Simo Kinnunen seems to do the job quite well so far. Instead of Flash, it relies on Canvas and VML in a pure JavaScript container. 

Getting it to work is really easy, too: Just upload a font on [http://cufon.shoqolate.com/generate/](http://cufon.shoqolate.com/generate/) and you will get a JavaScript file like ``MyFont_400.font.js``. One click to the left you get the main Cufón library that does the actual font-replacement. Link both into your document and then use ``Cufon.replace('h1')`` to, for instance, replace the h1-element of your page.

    <script type="text/javascript" src="/js/cufon-yui.js"></script>
    <script type="text/javascript" src="/js/MyFont_400.font.js"></script>
    <script type="text/javascript">
        Cufon.replace('h1');
    </script>

That's all. For starters. Cufón also exposes some options for each replacement and a way to easily bind fonts to dynamically created elements through its [API](http://wiki.github.com/sorccu/cufon/api). For now I'm just playing around with it, but it really looks promising :-)