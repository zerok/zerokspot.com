---
date: '2010-12-08T12:00:00-00:00'
language: de
tags:
- pygraz
- python
- graz
title: PyGraz-Stammtisch vom 7. Dez. 2010
---


Gestern Abend fand wie geplant wieder ein [PyGraz][]-Stammtisch statt, diesmal
jedoch nicht im sonst frequentierten Gösser Bräu, sondern erstmals im
[Realraum][]. Vielen Dank nochmals an Johannes und Christian Pointner für
die Organisation :-)

-----------------------------------

Thematisch ging es - wie üblich - quer durch den Gemüsegarten: Zunächst
präsentierte [Armin][] ein paar schlimme Details der Optimierungen im [Jinja2][] und
allgemein ein paar Beispiele, wie aus Templates hier Python generiert wird.

Damit war der zumindest ansatzweise geplante Teil des Abends auch schon wieder
vorbei und es folgt die übliche allgemeine Diskussion mit jeder Menge
Projekt-Erwähnungen:

* [ditaa][]: Ein Tool zur Generierung von Graphen aus ASCII-Art

* [homebrew][]: Mac-User, die homebrew noch nicht kennen, sollten es sich einfach
  gleich einmal anschauen :-) Sehr komfortables OSS-Paketmanagement, was
  speziell Entwicklern (!= Sysadmins) das Leben ziemlich erleichtert :-)

* [octobot][]: Ein Task-Queue-Worker, der z.B. an AMQP-Queues angehängt werden
  kann. Worker werden hier in JVM-Sprachen entwickelt, wofür ein einfaches
  Interface angeboten wird (effektiv einfach nur eine statische
  ``run``-Methode, die ein JSON-Objekt empfangen kann).

Daneben gab es auch einen kleinen Rückblick auf vergangene April-Scherze in
Form von PEPs:

* [PEP3117][]: Eine Postfix-Typen-Deklaration für Python mit Schneemännern,
  Lambdas und jede Menge Unicode

* [PEP0401][]: BDFL-Pensionierungspläne

Und da ich sämtliche Gerüchte über Memory-Leaks in [nodejs][] als Intrigen aus
dem Erlang-Lager ansehe, werden Diskussionen diesbezüglich hiermit von mir
totgeschwiegen ;-)

[ditaa]: http://ditaa.sourceforge.net/
[homebrew]: https://github.com/mxcl/homebrew
[octobot]: http://octobot.taco.cat/
[armin]: http://twitter.com/mitsuhiko
[realraum]: http://realraum.at
[Johannes]: https://github.com/thet
[Jinja2]: http://jinja.pocoo.org/
[pep3117]: http://www.python.org/dev/peps/pep-3117/
[pep0401]: http://www.python.org/dev/peps/pep-0401/
[nodejs]: http://nodejs.org
[erlang]: http://www.erlang.org
[pygraz]: http://pygraz.org
