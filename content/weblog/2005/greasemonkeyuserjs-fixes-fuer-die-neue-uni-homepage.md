---
date: '2005-10-30T12:00:00-00:00'
language: en
tags:
title: "Greasemonkey/UserJS-Fixes f\xFCr die neue Uni-Homepage"
---


Nachdem die neue Homepage der Uni Klagenfurt jetzt seit einiger Zeit online ist und ich ansich seit dem ersten Tag ein paar kleine Fixes für mich eingerichtet habe ... naja, vielleicht will sie ja noch wer verwenden ;)

-------------------------------



* Da anscheinend alles auch über https erreichbar, warum als nicht verwenden?! :)

* Das Login-Form öffnet ein komplett neues Fenster fürs ZEUS. Ich mag keine neuen Fenster ;)



<pre class="code">

// ==UserScript==

// @name Fix-Sammlung für www.uni-klu.ac.at

// @description (1) Da https verfügbar ist, sollte es verwendet werden ;)

//              (2) Das LoginForm sollte kein neues Fenster aufmachen.

// @include http://www.uni-klu.ac.at/*

// @include https://www.uni-klu.ac.at/*

// @namespace http://zerokspot.com

// ==/UserScript==

(function(){

	if (window.location.protocol == 'http:'){

		document.location = 'https://'+window.location.host+':'+window.location.port+window.location.pathname;

	}

	else {

		// Fix the login form

		document.getElementById("login").getElementsByTagName("form")[0].setAttribute("target","_self");

	}

})();

</pre>

<strong><a href="http://www.zerokspot.com/userjs/zk.uniklu.misc1.user.js">zk.uniklu.misc1.user.js</a></strong>



Das ist mal Version 1.0. Mal schaun, ob ich noch andere Dinge finde, die mich an der neuen Homepage stören ;)



Getestet mit Firefox 1.4.x + Greasemonkey 0.6.2 und Opera 8.5