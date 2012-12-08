A-Tag 2010
##########

Der `A-Tag`_ ist eine seit 2006 jährlich stattfindende Konferenz rund um das
Thema Accessibility bzw. das barrierefreie Internet.  Stattgefunden hat das
ganze dieses Jahr am 5. November im `TechGate`_ in Wien, und zwar im 19.
Stock. Also ganz oben!  Super Räume, jedoch leider effektiv ohne Steckdosen,
aber es gibt immer Wege ;-) Und diese Räume waren mit irgendwo zwischen 160
und 190 Personen sehr gut gefüllt.

Absolute toll ist hierbei auch, dass die Teilnahme kostenlos ist, da u.a.
auch das `Bundesministerium für Wirtschaft, Familie und Jugend`_ sponsort.
Also: Web-Firmen Österreichs, schickt nächstes Jahr zumindest einen Vertreter
hin!  :-)

Vortragsmässig wurde sehr viel und sehr Diverses geboten: Von einem
allgemeinen Überblick über Accessibility auf Plattformen von PC bis iPad über
Neues im Bereich von HTML5 und CSS3 bis hin zu was Webentwickler von Cowboys
lernen können.

------------------------------------------------------------------------

Angefangen hat alles um 0930 Uhr mit einer Begrüßung durch einen Vetreter des
Bundesministeriums für Wirtschaft, Familie und Jugend, auch bekannt unter dem
Namen `Robert Lender`_, den vermutlich jeder in der österreichischen
Accessibility-, Barcamp- und Photowalk-Szene kennt. Danach ging es auch schon
voll zur Sache:

Vorträge
========

Formvollended
-------------

Im ersten Vortrag beschrieb dann `Peter Minarik`_ von Wienfluss_ gutes
Form-Design in Kombination mit einem von Wienfluss entwickeltem Framework,
das die Einbindung von Dingen wie Tooltips, `WAI-ARIA`_-Roles etc.
erleichtern soll. Er brachte auch ein paar gute Punkte vor, wie z.B.
Hilfsinformation zu Formularelementen angezeigt werden sollte, d.h. sie sollte
während des Editiervorgangs dem User zugänglich sein und nicht den
Arbeitsfluss durch modale Elemente unterbrechen.

Auch ist die Verwendung des title-Attributs für Tooltips problematisch, da
die Anzeige dessen nicht mitzoomt, wenn der restliche Seitentext vergrößert
wird.

Weiters scheint onBlur-Validierung (also die Validierung, sobald ein Element
den Fokus verliert) allgemein langsam aber sich *die* beste Form der
clientseitige Validierung zu werden. \*zustimm\*


Barrierefreiheit zum Anfassen
-----------------------------

Gleich darauf gab `Marco Zehe`_ (Mozilla) einen Überblick über die
Accessibility-Situation auf diversen Plattformen und was Apple mit der
gesamten iOS-Schiene und OSX >= 10.4 für Blinde geleistet hat.

Ich war hier ein bisschen traurig zu hören, dass Microsoft es mit Windows
Phone 7 leider nicht einmal ansatzweise geschafft hat, eine brauchbare
Plattform aus Accessibility-Sicht zu schaffen.


Wie? Mein WordPresse-Theme geht auch barrierefrei!
--------------------------------------------------

Nach einer kurzen Pause beschrieb `Sylvia Egger`_ Probleme mit dem neuen
Default-Theme von `WordPress`_ (TwentyTen) und stellt auch ihre Arbeiten an einer
verbesserten Variante vor, das sie auch rechtzeitig zum A-Tag auch in der
`Version 1.0
<http://sprungmarker.de/2010/wordpress-child-theme-for-twenty-ten-accessible-1-0/>`_
ins Netz gestellt hat.

Irgendwie echt seltsam, dass WordPress scheinbar die ganze
Accessibility-Einstellungen in TinyMCE standardmässig deaktiviert hat :-/

Relaunch wien.at
----------------

Das erste von zwei Großprojekten, das während der Konferenz behandelt wurde,
war `wien.at`_ und der kürzliche Relaunch der nun schon 15 Jahre alten Website
der Stadt Wien. Hierbei lag vor allem die Video-Komponente samt CMS im Fokus
der Präsentation, da für den Inhalt sowohl Transkripte als auch Untertitel
angeboten werden.

Das im Rahmen dieses Relaunches zusammengstellte Framework ist frei (derzeit
leider noch unter GPL) auf `accessibility.at
<http://accessibility.at/practice/barrierefreies-video-player-framework>`_
verfügbar.

Für die Erstellung der Untertitel kommt `subtitle-horse`_ zum Einsatz, welches
während der Präsentation als "open source" bezeichnet wurde. Eine
Beschreibung, die jedoch nach weiterer Diskussion scheinbar nicht
aufrechterhalten werden konnte. Eine interessante Situation ;-)

Mainstream Accessibility
------------------------

In der letzten Präsentation vor dem Mittagessen zeigte `Klaus Miesenberger`_
auf, dass Accessibility leider noch immer nicht mainstream ist und unbedingt
in der Jugendkultur und den Köpfen eines jeden verankert werden muss, damit
sie auch durchgehend eingefordert und die Notwendigkeit nicht immer wieder
hinterfragt wird. Auf EU-Ebene wurde in diesem Rahmen das Projekt `eAccess+`_
ins Leben gerufen, das scheinbar als Content-Hub für Themen rund um
eAccessiblity dienen soll.

Barrierefreiheit leicht gemacht dank Mojo, Twitter und Konsorten
----------------------------------------------------------------

Nach einem ausgiebigen und verwehten (auf der Terrasse war einfach am meisten
Platz ;-)) Mittagessen ging es weiter mit einer Präsentation von `Philip
Naderer`_ über die Bestrebung von orf.at hin zur größten accessible Website
Österreichs.

Beginnend mit dem Redesign der Futurezone 2008 auf Basis der `WCAG 2.0`_ ist
derzeit Level AA bzw. sogar AAA ständiges Ziel der Weiterentwicklung des
Online-Angebotes. Aufgrund der Contentlastigkeit der Seite(n) mussten auch die
Redakteuere stark eingebunden werden und wurden mit einem rund 30 seitigen
Styleguide ausgerüstet, der 4 Seiten alleine der Barrierefreiheit widmet.

Getestet wird in JAWS 10-12, NVDA und sämtlichen aktuellen Mainstream-Browsern
bis hinunter zum IE6, wobei für diesen diverse Komponenten komplett
ausgeblendet werden, wenn sie nicht sinnvoll umsetzbar sind. In diesem
Zusammenhang wurde IE6 als "Bürobrowser" bezeichnet, da er eigentlich nur noch
zu typischen Büropausenzeiten und kurz vor Dienstschluss relevant wird.

Jetzt müssten nur noch diverse Großunternehmen davon überzeugt werden, dass es
doch sehr seltsam (und extrem schädlich für alle Beteiligten) ist, dass sie
einen Browser verwenden (und erzwingen) der vermutlich doppelt so alt ist, wie
das älteste Auto in ihrem Fuhrpark ;-)

Frische Webtechniken
--------------------

Als nächstes war `Eric Eggert`_ mit einem Vortrag über HTML5, CSS3 und mehr
(auch kurz NEWT = New Exciting Web Technologies) an der Reihe. Der Begriff
"NEWT" fasst es dann auch sehr gut zusammen :-) Nur ein paar Links:

* `html5shiv`_
* `html5accessibility.com`_
* `accessifyhtml5.js`_
* `jPlayer`_
* `jMediaelement`_

Was Webentwickler von Cowboys lernen können
-------------------------------------------

Der abschließende Talk war `Tomas Caspers`_ vorbehalten, der von einer Reise
in die USA auf eine Cowboy-Ranch berichtet und welche Erfahrungen er daraus
gewonnen hat, die man auf das Berufsleben und damit natürlich auch auf die
Arbeit als Web-Entwickler umlegen kann. Ein Beispiel hierfür war, dass
Multitasking in der Regel einfach nicht wirklich funktioniert. Manchmal ein
notwendiges Übel, sollte es jedoch möglichst vermieden werden.


Zusammenfassend...
==================

... kann ich den Organisatoren und Vortragenden nur ein Riesenlob und ein
großes Dankeschön für einen sehr informativen und unterhaltsamen Tag mit viel
Networking aussprachen. Wenn ich kann, würde ich nächstes Jahr gerne wieder
kommen.


.. _A-Tag: http://atag.accessiblemedia.at/2010/ 
.. _subtitle-horse: http://subtitle-horse.org
.. _techgate: http://www.techgate.at
.. _wienfluss: http://www.wienfluss.net/ 
.. _wien.at: http://wien.at
.. _jplayer: http://www.happyworm.com/jquery/jplayer/ 
.. _eaccess+: http://www.eaccessplus.eu/
.. _philip naderer: http://naderer.biz/ 
.. _tomas caspers: http://tomascaspers.de/ 
.. _eric eggert: http://yatil.de
.. _jmediaelement: http://www.protofunc.com/jme/ 
.. _accessifyhtml5.js: https://github.com/yatil/accessifyhtml5.js 
.. _html5accessibility.com: http://html5accessibility.com
.. _klaus miesenberger: http://www.integriert-studieren.jku.at/
.. _html5shiv: http://code.google.com/p/html5shiv/
.. _bundesministerium für wirtschaft, familie und jugend: http://bmwfj.gv.at/
.. _wcag 2.0: http://www.w3.org/TR/WCAG20/
.. _wordpress: http://wordpress.org
.. _wai-aria: http://www.w3.org/TR/wai-aria/
.. _robert lender: http://www.robertlender.info
.. _sylvia egger: http://sprungmarker.de
.. _peter minarik: http://twitter.com/pietropizzi
.. _marco zehe: http://www.marcozehe.de/
