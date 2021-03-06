---
date: '2005-09-24T12:00:00-00:00'
language: en
tags:
title: German keymap with some dead keys?
---


If you're using the german keymap of Xorg/XFree86 you've perhaps noticed that you can't write french accents. On the other hand, if you enable the nodeadkeys option you will have to hit ~ twice to get one tilde. After some googling I found this <a href="http://www.cl.cam.ac.uk/~mgk25/ucs/apostrophe.html">great article by Markus Kuhn</a> that explains the problematic in detail (I confess, I've only read the solution part ;-) ).

-------------------------------



Now I wanted to follow his first solution and alter the german keymap on Xorg just to find something really nice. The deadgraveacute-variant is already in there :-)



<pre class="code">partial alphanumeric_keys 

xkb_symbols "deadgraveacute" {

    // modify the default German layout to have only acute and grave

    // as dead keys (tilde and circumflex are needed as spacing characters

    // in many programming languages)

    include "de(basic)"

    key <TLDE> {	[ asciicircum,	degree		],

			[ notsign			]	};

    key <AD12> {	[ plus,		asterisk	],

			[ asciitilde,   dead_macron	]	};

    key <BKSL> {	[ numbersign,   apostrophe	],

			[ grave				]	};

};</pre>



So all you have to do is enable it in your xorg.conf by adding following line to your keyboard configuration ...



<pre class="code">

Option "XkbVariant" "deadgraveacute"

</pre>



... and restart your XServer. This way also Opera's "Go to Page" dialog will accept tildes ;)