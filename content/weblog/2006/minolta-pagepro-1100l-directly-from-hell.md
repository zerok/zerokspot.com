---
date: '2006-08-05T12:00:00-00:00'
language: en
tags:
- windows
- linux
title: Minolta PagePro 1100L directly from hell?
---


<img src="http://zerokspot.com/uploads/mpp1100l.jpg" alt="Minolta PagePro 1100L" class="left"/>While I had military service about 6 years ago, my parent's printer died and they went to the local retailer of Minolta to get a new one. When I came home I saw a brand new PagePro 1100L ... as I had feared: GDI, so no chance to get it work with GNU/Linux, BSD, whatever system that isn't Windows.

Now my parents are finally considering to get a new computer as a replacement for their nine year old AMD K6-200MMX. Since I'm not really motivated to help them with Windows problems, I'm currently on a small campaign to convince them to get a Mac. As a first task (at least for me) I wanted to get the Minolta printer to somehow work inside of a network. Therefor I simply wanted to first of all, get it to work with Windows XP. Well ... there are drivers hidden [somewhere on the Minolta server](ftp://ftp.minolta-qms.com/pub/crc/out_going/win2000/pp11l_eu.exe), but they are more or less useless as for every successful print you get one error message saying, that the cover of the printer is open while it isn't.

-------------------------------



So printing via WindowsXP is not really an option anymore. The other problem is, that - from what I understand about CUPS - the CUPS server has to have the driver for the printer shared by the Windows machine. Very funny!

The only solution I could come up with for now (and that at least would work in my head) is writing a small server script for the Windows machine to poll a certain folder and print everything that hasn't been printed yet ... yeah, stupid idea, but it should at least work.

If someone has a better idea where I could simply share the Windows printer to a Linux server but let the sharing Windows machine handle the printing process, I'm all ear :)
