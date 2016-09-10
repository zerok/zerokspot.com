---
date: '2008-07-31T12:00:00-00:00'
language: de
tags:
- austria
- django
- l10n
- localflavor
title: "localflavor f\xFCr \xD6sterreich in Django"
---


Kleines Update für alle Österreicher, die [Django](http://www.djangoproject.com) benutzen. Seit letztem Wochenende gibt es jetzt endlich auch ein eigenes localflavor.at-Package ([r8087](http://code.djangoproject.com/changeset/8087), [r8113](http://code.djangoproject.com/changeset/8113)). Zum aktuellen Zeitpunkt kommt es mit einer Select-Box für die neun Bundesländer, einem einfachen Formularfeld, das österreichische Postleitzahlen validiert, und einem Feld, das die hier gültigen Sozialversicherungsnummern akzeptiert, daher.

-------------------------------

Was heißt das im Klartext? Hier ein kleines Beispiel:

@@ python @@
from django.contrib.localflavor.at.forms import ATStateSelect, ATZipCodeField, ATSocialSecurityNumberField

from django import forms

class MyForm(forms.Form):
    ssn = ATSocialSecurityNumberField(label=u'Sozialversicherungsnummer')
    zipcode = ATZipCodeField(label=u'Postleitzahl')
    state = forms.CharField(label=u'Bundesland', widget=ATStateSelect)
@@

Dieses Formular erzeugt folgenden HTML-Code:

@@ html @@
<p><label for="id_ssn">Sozialversicherungsnummer:</label> 
    <input type="text" name="ssn" id="id_ssn" /></p>
<p><label for="id_zipcode">Postleitzahl:</label>
    <input type="text" name="zipcode" id="id_zipcode" /></p>
<p><label for="id_state">Bundesland:</label> 
    <select name="state" id="id_state">
        <option value="BL">Burgenland</option>
        <option value="KA">Carinthia</option>
        <option value="NO">Lower Austria</option>
        <option value="OO">Upper Austria</option>
        <option value="SA">Salzburg</option>
        <option value="ST">Styria</option>
        <option value="TI">Tyrol</option>
        <option value="VO">Vorarlberg</option>
        <option value="WI">Vienna</option>
    </select></p>
@@

Das Postleitzahl-Feld akzeptiert vierstellige Ziffernkombinationen und das Feld für Sozialversicherungsnummern mag nur Zahlen im Format xxxx xxxxxx. Diese werden soweit möglich auch auf ihre Gültigkeit überprüft. Zumindest soweit die [Informationen auf Wikipedia](http://de.wikipedia.org/wiki/Sozialversicherungsnummer#.C3.96sterreich) stimmen ;-)

Alle Bundesländernamen sind auch übersetzbar, also bitte nicht abschrecken lassen von den englischen Bezeichnungen hier.

Leider gab es bei der Entwicklung ein paar Überschneidungen, weshalb localflavor.at mehr oder weniger ein Gemeinschaftsprojekt aus 2 Tickets schlussendlich geworden ist: [#6427](http://code.djangoproject.com/ticket/6427) und [#7686](http://code.djangoproject.com/ticket/7686). Wer sich die Updates dort ansieht, wird merken, dass da doch ein bisschen parallel entwickelt wurde und mehr oder weniger das Gleiche dabei herausgekommen ist. Hauptsache, es ist jetzt alles drin, was ursprünglich drin sein sollte :-)

Ich wollte eigentlich noch ein Feld für gültige Telefonnummern drin haben, das wird sich jedoch ein bisschen verzögern, da der Austrian Numbering Plan doch ein bisschen heftig ist ;-)