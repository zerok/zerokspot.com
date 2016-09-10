---
date: '2006-06-07T12:00:00-00:00'
language: en
tags:
- fonts
- macosx
- openoffice
- x11
title: Adding fonts to OpenOffice.org 2.0 on MacOSX
---


You would love to use OpenOffice2.0 but the native NeoOffice is too slow and the X11 app has too ugly fonts for your taste? Well, at least for the latter there is a small workaround :) All that is needed are some additional TTFs. Since OpenOffice.org 2.0.app uses the same foundations as the Linux versions, more or less the same tricks also apply on MacOSX.



-------------------------------



So first of all: Where do we get some fonts? Probably the easiest way would be to get the fonts from MacOSX itself, but since X11 doesn't support the native Mac fonts, we will have to do some conversion magic here.

1. Create a new folder "ttfs" in your home directory: `mkdir ~/ttfs`
2. cd into it and use [fondu](http://fondu.sourceforge.net) to convert the fonts in /Library/Fonts: `cd ~/ttfs ; find /Library/Fonts -exec fondu {} \;`
3. Just to be on the safe side, also copy all the .ttf files from /Library/fonts into your ~/ttfs folder :)

Now your ttfs folder should be full of .ttf files :)

The next step is finding spadmin. It is hidden deep in the .app folder of OpenOffice, so do the following in an X11 terminal: ` /Applications/OpenOffice.org\ 2.0.app/Contents/openoffice.org2.0/program/spadmin`

You should get a window like this:

<img src="http://zerokspot.com/uploads/ooo.fonts.spadmin.1.png" alt="spadmin's main window"/>

There hit the "Fonts" followed by the "Add" button:

<img src="http://zerokspot.com/uploads/ooo.fonts.spadmin.2.png" alt="spadmin's font adder"/>

There you can simply go to your ttfs folder and add your TrueType fonts. After restarting OpenOffice you should have some more fonts available.