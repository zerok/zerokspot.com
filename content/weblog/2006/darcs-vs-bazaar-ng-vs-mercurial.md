---
date: '2006-04-29T12:00:00-00:00'
language: en
tags:
title: Darcs vs. Bazaar-NG vs. Mercurial
---


Just a short comparision of the two systems based on what I've read on the web and what I like and dislike about each of them. Note: I haven't really worked with any of them yet, but I definitely want to use a distributed versioning system for some stuff. This is just a really out of the hand listing of things I noticed while playing around with those systems, so it's ... informal would definitely be an understatement.



-------------------------------



## Darcs

[Darcs homepage](http://darcs.net/)

### Pros

* Offers a send system for directly mailing changes on a repository to someone else
* Offering an HTTP repository without the need for CGI support on the server (only the darcs binary has to run somewhere for the ssh checkin)
* After playing a little bit with it I like the patch-based approach
* It seems like you don't need Cygwin nor GHC installed for [Windows](http://darcs.net/DarcsWiki/CategoryBinaries#head-c7910dd98302946c671cf63cb62712589b392074)

### Cons

* Installing Haskell via DarwinPorts on my Powerbook was ... a pain. It took for ages.
* The installation of Darcs via DarwinPorts was even more a problem since ghc (Haskell) was barely motivated and segfault'd every few minutes. Going with binary released like [this](http://www.carpetcode.org/get-carpet-darcs.html#darcs) or the one from fink might be better here (although to me fink appears far less actively developed than DarwinPorts).
* Requires darcs installed on the server when you want to push your work (But there is a single-file statically linked binary for Linux so it's not really a problem)


## Bazaar-NG

[Bazaar-NG homepage](http://bazaar-vcs.org/)

### Pros

* Easily extensible
* Written in Python (sorry, but I know this language while Haskell is also after seeing a presentation of it totally obscure to me)

### Cons

* I somehow always have a strange feeling when a company is backing an opensource project. But I guess in this case and with the company being Canonical I can my an exception ;)
* There is nothing like `darcs send` to simply send patches to another maintainer
* Many Python requirement and it still somehow doesn't work on one of my webservers for no obvious reason.


After some additional searching and some chatting with [Keltia](http://www.keltia.net/) on irc://freenode.org/textmate I found out about [Mercurial](http://www.selenic.com/mercurial/) which appears to me somehow like a Python equivalent to Darcs. I just have a few problems with it:

<ul><li>If you rename a file, it basically removes and adds the same file with a different name. So the patch between these two revisions would actually hold the binary data of the new/old image, while Darcs (thanks to it's patch dependancy approach) will simply produce a patch with something like <pre class="code">[rename zerok@local**20060428080113] {<br/>move ./quicksilver.jpg ./quick.jpg<br/>}</pre>Really neat esp. if you move a binary (like in this case).</li>
<li>Exporting for binaries doesn't work. It produces a completely broken patch or basically ignores the binary file altogether. At least that's what came out for me, but perhaps I've just messed up something ;) As sad as this sounds: This is a show-stopper for me. The things I wanted to use a DVCS for now always involve the addition of new binary files and code exchange via e-mail (not as the primary way but as a good secondary).</li>
<li>It requires CGI enabled on the webserver ideally with suExec. There is also the so-called old-http option but it's not recommended according to the official website.</li></ul>

Esp. the binary patch problem in hg is currently holding me back from using it and probably going with Darcs at least for now. I heard some stories about Darcs not being able to handle bigger repositories all that well, but since I don't plan to use it for something like that, it shouldn't be a problem at all :) Bazaar-NG is currently not an option because the SubmitByMail feature as described on their [Wiki](http://bazaar-vcs.org/SubmitByMail) is not yet available. When bzr gets it, it might become an option for me thanks to its extensibility and that I somehow doubt that it will have problems with bigger repositories since Ubuntu is using it. On the other hand Mercurial looks even more promising thanks to its "lightweightness" and basic support for everything I need. It's just still a little bit too rough about the edges (esp. handling binary data) for now, but hopefully this will change in the future :)