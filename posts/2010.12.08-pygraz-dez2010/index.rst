PyGraz-Stammtisch vom 7. Dez. 2010
==================================

Gestern Abend fand wie geplant wieder ein PyGraz_-Stammtisch statt, diesmal
jedoch nicht im sonst frequentierten Gösser Bräu, sondern erstmals im
Realraum_. Vielen Dank nochmals an Johannes und Christian Pointner für
die Organisation :-)

-----------------------------------

Thematisch ging es - wie üblich - quer durch den Gemüsegarten: Zunächst
präsentierte Armin_ ein paar schlimme Details der Optimierungen im Jinja2_ und
allgemein ein paar Beispiele, wie aus Templates hier Python generiert wird.

Damit war der zumindest ansatzweise geplante Teil des Abends auch schon wieder
vorbei und es folgt die übliche allgemeine Diskussion mit jeder Menge
Projekt-Erwähnungen:

* ditaa_: Ein Tool zur Generiung von Graphen aus ASCII-Art

* homebrew_: Mac-User, die homebrew noch nicht kennen, sollten es sich einfach
  gleich einmal anschaun :-) Sehr komfortables OSS-Paketmanagement, was
  speziell Entwicklern (!= Sysadmins) das Leben ziemlich erleichert :-)

* octobot_: Ein Task-Queue-Worker, der z.B. an AMQP-Queues angehängt werden
  kann. Worker werden hier in JVM-Sprachen entwickelt, wofür ein einfaches
  Interface angeboten wird (effektiv einfach nur eine statische
  ``run``-Methode, die ein JSON-Objekt empfangen kann).

Daneben gab es auch einen kleinen Rückblick auf vergangene April-Scherze in
Form von PEPs:

* PEP3117_: Eine Postfix-Typen-Deklaration für Python mit Schneemännern,
  Lambdas und jede Menge Unicode

* PEP0401_: BDFL-Pensionierungspläne

Und da ich sämtliche Gerüchte über Memory-Leaks in nodejs_ als Intrigen aus
dem Erlang-Lager ansehe, werden Diskussionen disbezüglich hiermit von mir
totgeschwiegen ;-)

.. _ditaa: http://ditaa.sourceforge.net/
.. _homebrew: https://github.com/mxcl/homebrew
.. _octobot: http://octobot.taco.cat/
.. _armin: http://twitter.com/mitsuhiko
.. _realraum: http://realraum.at
.. _Johannes: https://github.com/thet
.. _Jinja2: http://jinja.pocoo.org/
.. _pep3117: http://www.python.org/dev/peps/pep-3117/
.. _pep0401: http://www.python.org/dev/peps/pep-0401/
.. _nodejs: http://nodejs.org
.. _erlang: http://www.erlang.org
.. _pygraz: http://pygraz.org
