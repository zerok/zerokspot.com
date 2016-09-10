---
date: '2006-05-05T12:00:00-00:00'
language: en
tags:
title: Sometimes LaTeX is just annoying
---


Not really, but sometimes I simply come across smaller problems that are in my opinion simply not logical. As part of a course I had to write excerpts of 4 different papers and put them all together into one big document. First of all I wrote the excerpts having their own article-class file so that I could also generate PDFs for each excerpt without having to go to the one big conglomerate. Every excerpt has an abstract. No problem so far.



-------------------------------



But when you then want to \include or \input the separate articles you get at least 2 problems:

1. Including/Inputing a whole article is a big NoNo because the merging process also processes the LaTeX headers, so you would end up with articles inside of a report... Perhaps someone could explain to me, why the \include command isn't simply dropping the offending header? So I had to write a small Ruby script for creating temp. documents without those headers.
2. If the included document holds an abstract-environment the whole page numbering will go nuts. Suddenly all the separate chapters will have their very own page numbering. So the first page of each chapter has the page number ... 1. 

Stupid, but at least now I know, what I sould have to write scripts for ;)