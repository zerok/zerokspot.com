---
date: '2005-08-03T12:00:00-00:00'
language: en
tags:
- development
- markdown
title: Was hat Textile, was Markdown nicht hat?
---


Während der letzten Woche habe ich mir öfters die Frage gestellt, was für eine Markup-Sprache ich für welche Applikation verwenden soll. So bietet mir zum Beispiel instiki die Auswahl aus 3 verschiedenen Sprachen:


-------------------------------


* Markdown
* Textile
* RDoc

Da ich Markdown schon von diesem Weblog hier ein bisschen kenne, hab ich es als erstes verwendet, aber dann trotzdem einfach einmal aus Neugier auch mal Textile ausprobiert und bin irgendwie zu dem Schluss gekommen, dass Textile für die Art, in der ich schreibe, doch besser geeignet ist als Markdown.

Warum? Relativ einfach: Textile gibt mir die Möglichkeit, Inline-Elemente mit CSS-Klassen bzw. IDs zu versehen.

<pre>%(class)text%
%(#id)text%</pre>

Das ist vor allem dann ziemlich nützlich, wenn man viel über Programmieren und ähnliche Themen schreibt und zum Beispiel Funktionsnamen ihren eigenen Style geben möchte. In Markdown habe ich bis jetzt leider noch nichts vergleichbares gefunden, weshalb ich jetzt vor der Qual der Wahl stehe:

1. Markdown erweitern, damit ich Klassen bzw. IDs zuweisen kann
2. oder Textile für WordPress so umschreiben, dass es pre-Blöcke komplett ignoriert bzw. sich einfach so verhält wie <a href="http://www.whytheluckystiff.net/ruby/redcloth/">RedCloth</a>

Derzeit tendiere ich mehr in Richtung (2), da mir auch sonst Textile ziemlich gut gefällt. Irgendwie kommt es mir einfach vielseitiger vor als Markdown. Mal schaun, wie die Sache ausgeht ;)
