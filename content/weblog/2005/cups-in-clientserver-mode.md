---
date: '2005-06-13T12:00:00-00:00'
language: en
tags:
title: CUPS in client/server mode
---


This weekend I found a nice little feature in CUPS that helps and probably also will help me a lot in the future: You can run CUPS also only as a client application by adding following line to your $HOME/.cupsrc

-------------------------------



<pre class="code">

ServerName yourServerHostname

</pre>



(Replace "yourServerHostname" with the IP or hostname of another machine running CUPS and having a printer plugged in ;-)  The server also has to be configured to allow connects from hosts other than 127.0.0.1)



Now CUPS will act on your user account as a client application and offer you all the printers on the other host for example in Firefox when you want to print something (or any other application that uses CUPS ;-) ). To get this behaviour systemwide add the line above to the /etc/cups/client.conf