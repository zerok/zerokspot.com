---
date: '2005-02-05T12:00:00-00:00'
language: en
tags:
- nucleuscms
title: NP_TechnoratiTags 0.2
---


<h4>What is it?</h4>

-------------------------------

NP_TechnoratiTags is a small playground for me for getting to know the Nucleus Plugin API. It adds a new field to the AddItem and EditItem forms where you can enter Technorati Tags in a space-separated format. 



<h4>New in this version</h4>

<ul>

<li>You can now manipulate the look&feel of the tags through the plugin panel</li>

</ul>



<h4>Planned for the next release</h4>

<ul>

<li>Finally some useful database access :P Now the playtime should be over so I'll start to put some useful code into this. Currently there is one db-query per post. I first have to check if and how I can access additional columns through Nucleus itself. If I can do this, there will probably be a database change moving the tags away from the separated table and into a new column of the items table. Let's see :-)</li>

</ul>



<h4>Note/Warning</h4>

This plugin is not intended to be used on a productive system. It was more or less just a playground for me to get to know the NucleusCMS plugin API so donât expect to much so no real code optimization or even nice code at all in there :-) This is also the reason why I donât want to submit this plugin to the NucleusCMS wiki. I also take no responsiblity for data loss or damage on your system ;-) I also had no real chance to test this plugin because I currently donât have a NucleusCMS site online :-(



<a href="http://www.zerokspot.com/uploads/NP_TechnoratiTags-0.2.tar.gz">NP_TechnoratiTags-0.2.tar.gz</a>