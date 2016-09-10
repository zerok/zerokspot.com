---
date: '2006-10-20T12:00:00-00:00'
language: en
tags:
- apple
- e-mail
- macosx
- mail-app
- mozilla
- thunderbird
title: Mail.app as alternative to Thunderbird?
---


<img src="http://zerokspot.com/uploads/inbox_horror.png" class="left" alt=""/>For the last couple of days (actually since I noticed, that Mail.app is now finally able to work with Socks5 proxies) I started considering Apple's Mail.app as an alternative for Mozilla Thunderbird for my daily mail usage. 



-------------------------------



Here a short list of things that made me consider switching away from Thunderbird:

* I'm not sure, I like the format my mails are stored in. From what I can see in my profile folder, every folder is stored in one single file. With an inbox holdering around 10,000 mails (or more) those single files are getting pretty big. And what happens if this one file gets corrupted? Sure, I could start trying to fix it, but the situation would be way easier, if Thunderbird had some kind of threshold for the number of messages per file (_not per folder_ ;)). Actually there is a [feature request](https://bugzilla.mozilla.org/show_bug.cgi?id=58308) for maildir support in the queue since 2000, but not much seems to have happened so far :-(
* No integration into the MacOSX AddressBook, which is in my opinion a more flexible solution than what Thunderbird offers.

And now a few things I really like about Mail.app:

* It highlights messages that are part of a thread also without using the threaded view.
* It is always clear from what mail account a mail came thanks to a simple division of the main inbox into subfolders for each account.
* I like Mail.app's sidebar more than the one of Thunderbird. It would be nice, if you could somehow expose the "Local Folders" when you don't use any other top-level folders for accounts.

But there are still some thing I really like about Thunderbird that Mail.app isn't offering me:

* Like Firefox compared to Safari it has a huge development community offering tons of add-ons. 
* The fishing protection and basically spam protection in Thunderbird feels a little bit more solid to me. It for example blocks images in mails from senders, you don't have in your address book. In Mail.app you can only disable images in general and then load them explicitly for each mail.
* Well, it's opensource.

Somehow in the gray area for me is Mail.app's GnuPG support. I found [GPGMail](http://www.sente.ch/software/GPGMail/English.lproj/GPGMail.html) but I haven't yet tried it (installed but not tried). 

As the length of the first point of my list of complaints with Thunderbird indicates: The whole situation of Thunderbird not supporting a maildir-like format is quite terrifying to me. Since GnuPG _is_ working with Thunderbird, the only thing that is holding me back so far, <s>is the "weaker" fraud/spam-protection in Mail.app</s> is that I don't know yet, where to find what in Mail.app. For example what can be controlled using an app-wide CSS file, but articles [like this](http://www.macosxhints.com/article.php?story=20040219094626558) definitely help pointing me into the right direction on how to get Mail.app to use such a CSS file in the first place ;-)


So far I'm really not sure what to do. I've been using Thunderbird for many years, but Mail.app is getting more and more an alternative for me. The only _real_ disadvantage I can see so far with using Mail.app, is that I'd now finally be bound to MacOSX. But since I will still keep all my mail on the respective servers and since there are tools like [emlxconvert](http://www.cosmicsoft.net/emlxconvert.html) I probably wouldn't have any bigger problems moving back. Another thing is Apple's AddressBook application. It seems not to have any built-in export facilities (except the one for creating a VCard for the currently selected contact) but there is at least the C Framework available as well as some AppleScript libs. So I guess ... all green on this front too. Let's see, perhaps I'm really motivated for this :-)
