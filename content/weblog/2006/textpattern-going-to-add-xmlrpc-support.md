---
date: '2006-01-23T12:00:00-00:00'
language: en
tags:
- textpattern
- xmlrpc
title: Textpattern going to add XMLRPC support
---


A week ago Pedro Palazß³n of the TXP core developer team <a href="http://textpattern.com/weblog/171/ask-a-dev-when-will-xml-rpc-show-up">announced support for XMLRPCs</a> for blogging on Textpattern where he also stated, that it will include complete support for the MovableType API. Now I just hope this also means, that the mt_keywords fields of MT's extension of the metaWeblog-API will be supported which would be a reason for me too look into TXP once again. 

-------------------------------



Why that? Quite simple: When I've got the time and I'm not motivated for doing anything else I'm still sometimes writing on my little Ruby script that will make blogging from the commandline possible. The problem here is, that Wordpress doesn't support this field of the XML request since it doesn't store keywords for a post. So I had to start messing around with the xmlrpc.php to extend the respective methods to support this field which (as always) would make updating a pain ...



I'm currently not really looking for any alternatives to Wordpress but given my low custom-fields usage during the last couple of weeks (thanks to writing basically everything in TextMate) and with TXP now on the road to support XMLRPCs there'd be at least a new option for me ;)
