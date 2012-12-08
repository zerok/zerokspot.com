Deutsches Übersetzungsprojekt für die offzielle Django-Dokumentation
####################################################################


So großartig die offizielle Django-Dokumentation auch ist, es wäre manchmal
hilfreich, wenn sie auch auf Deutsch verfügbar wäre. Sei es, weil man
sichergehen möchte, dass man ein Konzept wirklich verstanden hat oder auch
nur, um Personen ohne guten Englischkenntnissen dein Einstieg zu erleichtern. 

-------------------------------


Geschichtliches
---------------

Aus vermutlich diesen Gründen startete Jannis Leidel vor knapp 15 Monaten ein
`Projekt <http://django-de.org/>`_ mit dem Ziel einer vollständigen deutschen
Übersetzung von Djangos Dokumentation. Dann kam das große Refaktoring in der
englischen Dokumentation und leider mussten wir feststellen, dass sich die
Struktur dort dermassen stark verändert hat, dass wir nur sehr wenig einfach
portieren konnten (also ohne ~ 30% der Übersetzung zu ändern). Gleichzeitig
war das Projekt in einen Winterschlaf gefallen, aus dem es irgendwie nicht
mehr so recht aufwachen sollte. Alle halten irgendwie so viel zu tun, dass
kaum mehr Zeit zum Übersetzen blieb.

2008 war auch das Jahr, in dem DVCS-Lösungen dank Diensten wie `Github
<http://github.com/>`_ und `Bitbucket <http://bitbucket.org/>`_ immer 
attraktiver wurden. Viele von uns nutzen git/mercurial schon
weitestgehend exklusiv, weshalb wir auch viel überlegten, wie und ob wir auf
git migrieren sollten.

Ende 2008 eröffnete ich dann auf Github ein `kleines Repository
<http://github.com/zerok/django-docs-de/>`_, um den ganzen
Entscheidungsprozess ein bisschen zu erleichtern und fing auch an, einzelne
Einführungsdokumente erneut zu übersetzen. Dabei hat sich dann auch gezeigt,
dass dank der zusätzlichen ReST-Direktiven, die für die englischsprachige
Dokumentation implementiert wurden, die Portierung der bisherigen
Übersetzungen verhältnismässig aufwendig würde, weshalb ich sie immer wieder
vor mich her geschoben habe. 

Aktueller Stand
---------------

Derzeit ist in etwa ein Drittel der Dokumentation für Django 1.0.X übersetzt,
somit gibt es noch genug zu tun (wobei sich die Dokumentation zwischen 1.1 und
1.0.X bisher nur minimal geändert haben dürfte).

Ungefähr 90 Tickets sind im Projekt-Ticker noch offen, was derzeit 1:1 auf
noch zu übersetzende Dokumente mappt.

Wie kann kannst du helfen?
--------------------------

Mithelfen ist relativ einfach:

1. Schau in den `Bugtracker <http://github.com/zerok/django-docs-de/issues>`_
   und such dir ein Ticket aus, für das du eine Übersetzung bauen möchtest.

2. Kommentiere in dem Ticket, damit alle anderen Wissen, dass das Ticket
   bereits bearbeitet wird. Sobald ich das Kommentar sehe, tagge ich das
   Ticket entsprechend.

3. Falls du noch keinen Fork von `django-docs-de
   <http://github.com/zerok/django-docs-de/>`_ auf Github angelegt hast, tu
   das jetzt. Entsprechende Dokumentation hierfür findest du in den
   `Github-Guides <http://github.com/guides/home>`_.

4. Übersetze das Dokument und committe es in dein Repository (und pushe es
   auf Github).

5. In der Regel bekomme ich das mit (falls nicht, bitte einfach in #django-de
   auf Freenode ansprechen), schaue die Änderungen durch und baue deine
   Übersetzung in das "offizielle" Repository ein.

Das wars auch schon :-) Nur noch ein paar Kleinigkeiten:

* Falls du bei der Übersetzung auf eine Phrase stößt, wo du nicht ganz sicher
  bist, wie du sie übersetzen sollst, wirf auch einen Blick in unseren
  `Glossar <http://wiki.github.com/zerok/django-docs-de/glossar>`_.

* Die ganze Übersetzung ist derzeit informell gehalten. Wir sind doch alle
  implizit per du ;-)

* Die Entwicklung findet derzeit im "releases/1.0.X"-Branch statt. Bitte achte
  darauf, wenn du deine Änderungen eincheckst!

Wenn du mithelfen möchtest, würden ich mich sehr freuen :-) Falls du gerade
auf der `EuroDjangoCon09`_ in Prag bist, können wir ja vielleicht einen
OpenSpace organisieren, falls das Interesse da ist.

.. _eurodjangocon09: http://euro.djangocon.org/
