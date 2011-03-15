Gestern war mal wieder `Webmontag in Graz`_ und ich habe auch eine kleine
Präsentation zu Webentwicklung in Python beigetragen. Die Folien selbst
onlinezustellen macht eher weniger Sinn, daher folgt im Anschluss einfach eine
Link-Liste aller Tools und Frameworks bzw. sonstiger Resourcen, die ich erwähnt
habe :-)

-----------------------------

Frameworks
==========

* Django_: Framework aus dem Zeitungsbereich und daher primär auf CMS-Projekte ausgerichtet. Kommt mit eigenem Template-, Forms- und ORM-Layer.
* Flask_: Micro-Framework mit Jinja2_ als Template-Sprache
* Pyramid_
* web.py_: Das "klassische" Micro-Webframework in Python ;-)
* Bottle_: Vergleichbar mit Flask, jedoch hat das Team zusätzlich das Ziel,
  keine externen Abhängigkeiten zu haben.
* Zope_: Das klassiche Framework für CMS-lastige bzw. corporate Anwendungen

Libs und sonstiges Grundgerüst
==============================

* Werkzeug_: Bibliothek, die das Arbeiten mit WSGI erleichtert. Werkzeug ist
  vom selben Autor wie Flask und wird dort auch als Basis verwendet.
* Jinja2_: Ebenfalls vom Autor von Werkzeug und Flask stellt Jinja2 eine
  populäre Template-Sprache dar.
* flup_: Sammlung von WSGI-Tools, wobei hiervon vor allem das Interface von WSGI auf
  FastCGI u.a. von Django verwendet wird.
* mod_wsgi_: WSGI-Modul für Apache2
* gunicorn_: Schneller und schlanker WSGI-HTTP-Server mit eingebautem Worker-Management
* virtualenv_: Sandbox-Tool für Bibliotheken
* pip_: Paket-Management-Tool
* supervisord_ ermöglicht einfaches Prozessmanagement ähnlich zu init.d unter
  Linux, jedoch plattformunabhängig. Leider habe ich ein paar Features mit
  jenen von upstart_ verwechselt. Sorry :-)
* fabric_: Tool zum einfachen Deployment auf einen oder mehrere Server. Fabric
  abstrahiert primär SSH- und Filesystem-Operationen.

Hosting
========

Irgendwie habe ich bei den Folien ganz auf Hostinganbieter vergessen, daher
nun als Nachtrag auch hierzu eine kleine Liste. VPS-Lösungen gehen natürlich
immer, also zuerst eine kurze Liste von Diensten, die ich selbst schon
verwendet habe und glücklich damit bin bzw. war:

* Linode_: VPS-Anbieter mit Servern in den USA und UK.
* Carrot-Server_: VPS-Anbieter aus Graz mit Servern primär in Deutschland

Daneben sind in den letzten Monaten eine Reihe von Hostern aufgetaucht, die
sich auf WSGI spezialisiert haben. Leider hatte ich bis jetzt noch nicht die
Gelegenheit, diese auch auszuprobieren, aber auflisten schadet ja nicht :-)

* ep.io_: Cloud-basierter Webhoster und UK mit Verrechnung entsprechend der
  tatsächlich verbrauchten Resourcen.
* gondor.io_: Noch nicht wirklich online, aber hier soll in Zukunft eine
  Hostinginfrastruktur auch für Projekte außerhalb von Eldarion basierend auf
  der Rackspace-Cloud-Infrastruktur entstehen.

Zusätzlich wird noch in vielen Diskussionen dotcloud_ erwähnt, was auch
durchaus vielversprechend aussieht, sollte der Preis stimmen :-)

.. _web.py: http://webpy.org/
.. _werkzeug: http://werkzeug.pocoo.org/
.. _jinja2: http://jinja.pocoo.org/
.. _flask: http://flask.pocoo.org
.. _django: http://www.djangoproject.com
.. _bottle: http://bottlepy.org/
.. _fabric: http://fabfile.org
.. _supervisord: http://supervisord.org/
.. _pip: http://pypi.python.org/pypi/pip
.. _virtualenv: http://pypi.python.org/pypi/virtualenv
.. _gunicorn: http://gunicorn.org/
.. _flup: http://trac.saddi.com/flup
.. _mod_wsgi: http://www.modwsgi.org/
.. _upstart: http://upstart.ubuntu.com/
.. _linode: http://www.linode.com/
.. _carrot-server: http://www.carrot-server.com/
.. _ep.io: http://www.ep.io/
.. _gondor.io: http://gondor.io/
.. _dotcloud: http://www.dotcloud.com/
.. _zope: http://zope.org/
.. _pyramid: http://docs.pylonsproject.org/projects/pyramid/1.0/index.html
.. _webmontag in graz: http://webmontag-graz.at/
