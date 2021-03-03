---
date: '2006-08-06T12:00:00-00:00'
language: en
tags:
- javascript
- zerokspot
title: Now work in IE
---


As you might or might not have noticed, the site should now look way better (apart from the top menu) and also the dropdown menus in the sidebar should now finally work. Sorry for that, but I hadn't noticed that problem until yesterday. 

What was the problem with the sidebar? Quite simple. I build the dropdown menus using DOM. There I had something like this to set the onchange eventhandler:

<pre class="code">selectbox.setAttribute("onChange","redirect(this.value)");</pre>



-------------------------------



For some reason Internet Explorer 6.x seems to ignore this event handler. Perhaps because it was set as a simple attribute and perhaps IE interprets event handlers as something different. Anyway, replacing it with following code now seems to work in all maojor browsers. At least I've tested it now with IE, Firefox, Opera9 and Safari.

<pre class="code">selectbox.onchange=function(){redirect(selectbox.value);};</pre>

If you find some other problems esp. with Internet Explorer, please let me know using the comment form below or the contact form :)
