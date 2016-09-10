---
date: '2006-02-01T12:00:00-00:00'
language: en
tags:
title: MySQL5 as development system and exporting from it
---


For the last couple of weeks now I've been using MySQL5 for my local development system. Since I really enjoy using the latest stable releases (and sometimes unstable to pre-beta ;)) going with MySQL5 was a logical choice. But this also comes with a small problem: The compatiblity mode of mysqldump. Have you every exported a table that has a key with auto incrementation from a MySQL5 database and tried to import it into something < MySQL5? Fun ;) If you export it right away and ignoring the whole compatiblity settings you are nearly guaranteed to get an error message about the collation/charset settings appended to each tables DDL-statement (and if you're lucky also something about some incompatible functions or datatypes).



-------------------------------



So, the next logical step would be, to use one of the compatibility flags like "--compatible=mysql40" or "--compatible=mysql323". This should remove or adapt all the incomptible stuff, shouldn't it? 

Well, I was quite suprised, when I restored the database and ... the auto\_increment flags were simply missing. I haven't tried it yet but somehow "--compatible=no\_table\_options" looks more comptaible with MySQL4 than the mysql40 option \*g\* Isn't it strange, that the compatiblity mode actually removes features that have been present in MySQL for ages? Sure, if something like the auto incrementation gets dropped to get compatible with ANSI SQL92, no problem with that. But to get a MySQL 4.x or 3.x compatible dump?

Or does anyone out there know a better way to get this working?