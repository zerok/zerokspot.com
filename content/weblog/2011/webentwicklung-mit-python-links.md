---
date: '2011-03-15T12:00:00-00:00'
language: de
tags:
- python
- webmontag
- graz
- webg
title: 'Webentwicklung mit Python: Link-Liste'
---


Gestern war mal wieder [Webmontag in Graz][wmg] und ich habe auch eine kleine
Präsentation zu Webentwicklung in Python beigetragen. Die Folien selbst
onlinezustellen macht eher weniger Sinn, daher folgt im Anschluss einfach eine
Link-Liste aller Tools und Frameworks bzw. sonstiger Ressourcen, die ich erwähnt
habe :-)

-----------------------------

## Frameworks

* [Django][]: Framework aus dem Zeitungsbereich und daher primär auf CMS-Projekte ausgerichtet. Kommt mit eigenem Template-, Forms- und ORM-Layer.
* [Flask][]: Micro-Framework mit [Jinja2][] als Template-Sprache
* [Pyramid][]
* [web.py][]: Das "klassische" Micro-Webframework in Python ;-)
* [Bottle][]: Vergleichbar mit Flask, jedoch hat das Team zusätzlich das Ziel,
  keine externen Abhängigkeiten zu haben.
* [Zope][]: Das klassische Framework für CMS-lastige bzw. corporate Anwendungen

## Libs und sonstiges Grundgerüst

* [Werkzeug][]: Bibliothek, die das Arbeiten mit WSGI erleichtert. Werkzeug ist
  vom selben Autor wie Flask und wird dort auch als Basis verwendet.
* [Jinja2][]: Ebenfalls vom Autor von Werkzeug und Flask stellt Jinja2 eine
  populäre Template-Sprache dar.
* [flup][]: Sammlung von WSGI-Tools, wobei hiervon vor allem das Interface von WSGI auf
  FastCGI u.a. von Django verwendet wird.
* [mod_wsgi][]: WSGI-Modul für Apache2
* [gunicorn][]: Schneller und schlanker WSGI-HTTP-Server mit eingebautem Worker-Management
* [virtualenv][]: Sandbox-Tool für Bibliotheken
* [pip][]: Paket-Management-Tool
* [supervisord][] ermöglicht einfaches Prozessmanagement ähnlich zu init.d unter
  Linux, jedoch plattformunabhängig. Leider habe ich ein paar Features mit
  jenen von [upstart][] verwechselt. Sorry :-)
* [fabric][]: Tool zum einfachen Deployment auf einen oder mehrere Server. Fabric
  abstrahiert primär SSH- und Filesystem-Operationen.

## Hosting

Irgendwie habe ich bei den Folien ganz auf Hostinganbieter vergessen, daher
nun als Nachtrag auch hierzu eine kleine Liste. VPS-Lösungen gehen natürlich
immer, also zuerst eine kurze Liste von Diensten, die ich selbst schon
verwendet habe und glücklich damit bin bzw. war:

* [Linode][]: VPS-Anbieter mit Servern in den USA und UK.
* [Carrot-Server][]: VPS-Anbieter aus Graz mit Servern primär in Deutschland

Daneben sind in den letzten Monaten eine Reihe von Hostern aufgetaucht, die
sich auf WSGI spezialisiert haben. Leider hatte ich bis jetzt noch nicht die
Gelegenheit, diese auch auszuprobieren, aber auflisten schadet ja nicht :-)

* [ep.io][]: Cloud-basierter Webhoster und UK mit Verrechnung entsprechend der
  tatsächlich verbrauchten Ressourcen.
* [gondor.io][]: Noch nicht wirklich online, aber hier soll in Zukunft eine
  Hosting-Infrastruktur auch für Projekte außerhalb von Eldarion basierend auf
  der Rackspace-Cloud-Infrastruktur entstehen.

Zusätzlich wird noch in vielen Diskussionen [dotcloud][] erwähnt, was auch
durchaus vielversprechend aussieht, sollte der Preis stimmen :-)

[web.py]: http://webpy.org/
[werkzeug]: http://werkzeug.pocoo.org/
[jinja2]: http://jinja.pocoo.org/
[flask]: http://flask.pocoo.org
[django]: http://www.djangoproject.com
[bottle]: http://bottlepy.org/
[fabric]: http://fabfile.org
[supervisord]: http://supervisord.org/
[pip]: http://pypi.python.org/pypi/pip
[virtualenv]: http://pypi.python.org/pypi/virtualenv
[gunicorn]: http://gunicorn.org/
[flup]: http://trac.saddi.com/flup
[mod_wsgi]: http://www.modwsgi.org/
[upstart]: http://upstart.ubuntu.com/
[linode]: http://www.linode.com/
[carrot-server]: http://www.carrot-server.com/
[ep.io]: http://www.ep.io/
[gondor.io]: http://gondor.io/
[dotcloud]: http://www.dotcloud.com/
[zope]: http://zope.org/
[pyramid]: http://docs.pylonsproject.org/projects/pyramid/1.0/index.html
[wmg]: http://webmontag-graz.at/
