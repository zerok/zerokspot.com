---
date: '2009-12-14T12:00:00-00:00'
language: de
tags:
- webdevelopment
title: Resource packages - generische CSS Sprites
---


Irgendwie ist mir das total entgangen, deshalb &quot;Sorry&quot; f&uuml;r die Versp&auml;tung :-)&nbsp;<em style="font-style: italic; "><a href="http://limi.net/articles/resource-packages/">Resource Package</a></em>&nbsp;ist ein Vorschlag von Alexander Limi f&uuml;r ein durchaus g&auml;ngiges und absolut reales Problem: Viele Seiten verlangen vom Browser, dass er eine ganze Reihe von Ressourcen wie Stylesheets, JavaScripts und Bilder herunterl&auml;dt, bevor die Seite vollst&auml;ndig dem Benutzer pr&auml;sentiert werden kann.

Dabei entsteht jedoch das Problem, dass Browser ein Limit haben, wie viele gleichzeitige Verbindungen sie zu einem einzigen Server herstellen, einerseits um diesen nicht zu sehr zu belasten, andererseits nat&uuml;rlich auch, um die Verbindungen des Clients nicht zu &uuml;berstrapazieren.&nbsp;<a href="http://www.w3.org/Protocols/rfc2616/rfc2616-sec8.html#sec8.1.4">HTTP/1.1</a>&nbsp;spricht hier von 2 Verbindungen.&nbsp;<a href="http://www.microsoft.com/windows/internet-explorer/readiness/developers-new.aspx">IE8</a>&nbsp;geht da ein bisschen weiter und bietet 6. In Anbetracht von Seiten, die mehr als 30 Ressourcen laden wollen, ist selbst das nur ein Tropfen auf den hei&szlig;en Stein. Aus diesem Grund gilt auch schon seit L&auml;ngerem die Devise f&uuml;r die Entwicklung performanter Websites, die Anzahl der n&ouml;tigen Verbindungen m&ouml;glichst gering zu halten.

## Eine Lösung

Die Idee von Alexander Limi ist hierf&uuml;r eine L&ouml;sung, die ziemlich ins extrem geht: Ressourcen sollen in einem einzigen Zip verpackt an den Benutzer ausgeliefert werden, welches &uuml;ber einen &lt;link&gt;-Tag im Header verlinkt wird. Alles andere bliebt gleich: Bilder, CSS, JavaScript et al. werden nach wie vor direkt im Markup verlinkt und stehen somit allen Browsern zu Verf&uuml;gung. Browser die jetzt <em>Resource Packages</em> unterst&uuml;tzen, laden zun&auml;chst das Archiv herunter und fragen danach lediglich jene Ressourcen noch an, die nicht darin enthalten waren.

<pre class="code">&lt;link rel=&quot;resource-package&quot; 
    type=&quot;application/zip&quot; 
    href=&quot;/resources/package.zip&quot; /&gt;</pre>

Das Archiv selbst hat ein denkbar einfaches Format. Es enth&auml;lt lediglich die Ressourcen, sowie ein Manifest, durch das der Browser vorab die Information erh&auml;lt, welche Dateien wirklich im Zip enthalten sind, ohne dieses zuvor vollst&auml;ndig entpacken zu m&uuml;ssen. Die URLs f&uuml;r die enthaltenen Dateien werden danach vom Browser relativ zum Pfad das Packages aufgel&ouml;st, weshalb es sich anbietet, dieses auf einer Ebene mit den gepackten Ressourcen auf dem Webserver anzubieten. Also beispielsweise:

<pre class="code">/resources/package.zip
/resources/css/core.css
/resources/css/print.css
/resources/css/screen.css
/resources/img/logo.png
...</pre>


Wobei das package.zip folgenden Inhalt h&auml;tte:

<pre class="code">manifest.txt
css/core.css
css/print.css
css/screen.css
img/logo.png
...</pre>

Im Allgemeinen finde ich das wirklich eine feine Sache und ich finde es noch toller, dass es abh&auml;ngig vom Community-Feedback nicht einmal schlecht f&uuml;r eine <a href="http://hacks.mozilla.org/2009/11/a-proposal-resource-packages-to-improve-performance/">Implementierung in Firefox 3.7</a> aussieht. Durch die Verwendung von bereits existierenden Formaten und der einfachen Einbindung und Struktur ist es ein angenehm einfacher und pragmatischer Vorschlag.

## Allgemein einsetzbar?

F&uuml;r mich pers&ouml;nlich sehe ich vorerst hier eine M&ouml;glichkeit, &quot;Standard-Ressourcen&quot; einer Seite zu bundlen. Damit meine ich Stylesheets, Scripts und Bilder, die auf dem Gro&szlig;teil einer Website Verwendung finden. Hier wird man sich sicherlich ansehen m&uuml;ssen (je nach Website), ob man nicht auch s&auml;mtliche andere Unterressourcen (Stylesheets f&uuml;r die Kontaktseite etc.) gleich mit ins zip packen kann, oder ob es sich aufgrund der Zugriffszahlen etc. eher anbietet, die weiterhin getrennt zu halten.

In seinem Blogpost erw&auml;hnt Alexander Limi auch immer YouTube und die dort verwendeten Video-Thumbnails als Vorteil von Ressource Packages &uuml;ber klassische L&ouml;sungen wie CSS Sprites. Auf jeder Videoseite werden auch immer eine Reihe von Vorschl&auml;gen f&uuml;r anderen Content abgebildet, deren Previews allesamt aus einem Resource Package bezogen werden k&ouml;nnten. Auf Suchergebnisseiten w&auml;re das zwar deutlich komplexer (da die Packages noch dynamischer zusammengebaut werden m&uuml;ssten), aber sicher auch ein interessantes Problem :-)

## Tool-Support

Falls <em>Resource Packages</em> wirklich kommen, wird es hierf&uuml;r sicherlich auch recht schnell Tools geben, die es Webentwicklern einfach machen, diese schnell und dynamisch zu erzeugen, damit eben auch Suchseiten davon profitieren k&ouml;nnen. Aber auch f&uuml;r die Erstellung der Packages im allgemeinen und statischen Fall, wird sich vermutlich der eine oder andere finden, der f&uuml;r Tools wie Coda und Asset-Manager im Allgemeinen Plugins schreibt, zumal die manifest.txt ja zu Beginn des Archivs sein muss. Und Tools wie 7zip, Winzip und Freunde bieten derzeit keine einfach M&ouml;glichkeit, die Reihenfolge der Dateien in einem Zip zu manipulieren (zumindest nicht, dass ich w&uuml;sste).

<em>Zusammenfassend</em> kann ich nur sagen: Ich hoffe es kommt bald ... und wird wirklich von allen Browser-Herstellern in absehbarer Zeit unterst&uuml;tzt :-)
