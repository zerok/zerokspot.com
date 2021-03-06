---
date: '2009-05-04T12:00:00-00:00'
language: de
tags:
- django
- django-docs-de
- sphinx
- virtualenv
title: Wie konvertiert man die deutschen Django-Docs nach HTML?
---


Wenn du die deutschsprachige Django-Docs verwenden möchtest, ist das leider
noch nicht so einfach, wie es sein sollte. Ich arbeite daran, aber in der
Zwischenzeit soll dir diese kurze Anleitung dabei helfen, den ganzen Prozess
trotzdem ohne größere Frustration zu überstehen.


-------------------------------

Die aktuelle Version von Sphinx (dem Dokumentationstoolkit, in dem Djangos
Dokumentation verfasst ist) ist 0.6.1. Leider ist Django damit noch nicht ganz
kompatibel, es befindest sich jedoch ein Patch dafür in der Warteschlange. Bis
dieser jedoch in trunk landet und er auch im deutschen Dokumentationsprojekt
integriert wird, benötigen wir leider noch Sphinx 0.5.2.

Wenn du nun bei der Übersetzung helfen möchtest oder einfach nur an einer
deutschen Übersetzung der offiziellen Django-Dokumentation interessiert bist,
sollte dein Workflow in etwa so aussehen:

1. **Quellcode auschecken:** Das kannst du mehr oder weniger mit ``git clone
   git://github.com/zerok/django-docs-de.git`` erledigen. Dabei wird der
   jeweils aktuelle Entwicklungsbranch ausgecheckt (also derzeit
   "releases/1.0.X").

2. Als nächstes solltest du **virtualenv installieren**, damit du einfach die
   benötigte Version von Sphinx installieren kannst, ohne dass du damit andere
   Anwendungen beinflusst:
        
        # Installiere virtualenv, falls du es noch nicht hast
        easy_install virtualenv

        cd django-docs-de

        # Erzeuge eine virtualenv fuer Sphinx
        virtualenv --no-site-packages env

        # Lade die virtualenv
        source env/bin/activate

        # Installiere Sphinx
        easy_install 'Sphinx==0.5.2'

3. Jetzt hast du alles Nötige installiert und kannst mit ``make html`` die
   aktuelle Dokumentation in HTML übersetzen. Das Ergebnis findest du dann im
   ``_build/html``-Ordner.