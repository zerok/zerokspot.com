---
date: '2005-02-04T12:00:00-00:00'
language: en
tags:
- nucleuscms
title: NP_TechnoratiTags 0.1...
---


OK, here it is, my first plugin for NucleusCMS. This small little something adds a new field to the AddItem and EditItem dialogs where you can enter tags in a "space-seperated" format. The tags are then stored in the sql_table('plug_technoratitags').

-------------------------------



The plugin also listens to the PreItem and PostItem events to add the tags to the $data['item']->body field so that no templates have to be modified to integrate the tags. I had no time so far to write a small plugin option to modify the look and feel of the output there so you will have to modify this in the event_PreItem function.



<strong>Note:</strong> This plugin is not intended to be used on a productive system. It was more or less just a playground for me to get to know the NucleusCMS plugin API so don't expect to much so no real code optimization or even nice code at all in there :-) This is also the reason why I don't want to submit this plugin to the NucleusCMS wiki. I also take no responsiblity for data loss or damage on your system ;-)

I also had no real chance to test this plugin because I currently don't have a NucleusCMS site online :-(



<a href="http://www.zerokspot.com/uploads/NP_TechnoratiTags.phps">NP_TechnoratiTags.phps</a>