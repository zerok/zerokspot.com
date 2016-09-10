---
date: '2004-06-29T12:00:00-00:00'
language: en
tags:
- software
title: Adding a new mime-type in Gnome 2.x
---


Well, if you know me then you know that I always switch from one WM to another after a week max. a month. I always quite liked Gnome 2.x because of its simplicity. I don't need 1000 options in a config dialog. I only need the important ones and the rest somewhere in a config file or as in Gnome's case in some kind of registry. "So, yeah... you like Gnome, then why don't you use it and keep us from reading all this?" Quite simple: Currently the whole filetype handling in Gnome seems to me from only complicated to nearly unusable, because I can't add a new mimetype and bind it to an application if I want to.

After some weeks of playing and searching I've now finally found a ways through the freedesktop.org mime-type definition/handling that is used by Gnome 2.x.

---------------

In this work-in-progress tutorial we will create a new mimetype for phpBB 2.0.x's template files. It's not the most useful example because Gnome should recognize it as HTML file out of the box but anyway.

First of all, you need a <strong>.local/share/mime/packages/</strong> folder in your $HOME directory. If you don't have one, then create it :-) In there you need an <strong>Override.xml</strong> where the new mime-type gets defined.

Now open that file and put following into it:

<pre class="code">
&lt;?xml version="1.0" ?&gt;
&lt;mime-info xmlns="http://www.freedesktop.org/standards/shared-mime-info"&gt;
&lt;mime-type type="text/x-phpbb-template"&gt;
   &lt;comment&gt;phpBB template&lt;/comment&gt;
   &lt;glob pattern="*.tpl"/&gt;
&lt;/mime-type&gt;
&lt;/mime-info&gt;
</pre>

The first two lines and the last one are just the base of this file. The more interesting part is

<pre class="code">
&lt;mime-type type="text/x-phpbb-template"&gt;
   &lt;comment&gt;phpBB template&lt;/comment&gt;
   &lt;glob pattern="*.tpl"/&gt;
&lt;/mime-type&gt;</pre>

mimetype is defined.

These lines define the new mimetype , a comment in the system's default language (in my case English) and the filename-pattern that should match this new mime-type.

After saving the file run <strong>update-mime-database ~/.local/share/mime</strong> in a terminal. This will merge you local mime-definitions with the system definitions.

Now open the <strong>gnome-file-types-properties</strong> dialog. Where you create a new Mimetype and assign an application to it. Specifying a filename extension is not necessary because you've already done that in your Override.xml.

<img class="figure" src="http://weblog.zerokspot.com/wp-content/gnome-filetp.jpg" alt="Gnome Filetypes Properties dialog"/>

Now simply restart all your Nautilus instances (or kill Nautilus if you also have it rendering your desktop) and the new filetype for phpBB (2.0.x) template files should be assigned to GVim :-)

## Links

<ul><li><a href="http://freedesktop.org/Standards/AddingMIMETutor" title="Adding MIME-Type tutorial on Freedesktop.org"/> Adding MIME-Information into the database (freedesktop.org)</a></li></ul>
