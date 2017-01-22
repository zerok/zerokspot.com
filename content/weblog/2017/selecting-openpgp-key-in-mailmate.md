---
date: 2017-01-22T19:35:31+01:00
title: Selecting an OpenPGP key in MailMate
tags:
- gnupg
- mailmate
---

For pretty much all my not-work-related mail [MailMate][] has become my favorite
tool. It is fast, lightweight, supported OpenPGP through GnuPG 2.x, and has
decent shortcuts. I also really like that even if the author hasn't the found
time to put a GUI around a feature's settings, *there are settings*.

In a recent example I wanted to make sure that for one particular e-mail address
a certain OpenPGP encryption key was used. MailMate doesn't have any kind of GUI
for managing keys but you can still choose which one to be used for which e-mail
address through the `~/Library/Application Support/MailMate/Security.plist`
file. In there you basically just create a simple mapping like this:

```
{
	map = (
		{
			address = "<e-mail address>";
			userID = "<key id>";
		}
	);
}
```

You can learn the details deep down in the [hidden-preferences section][] of the
MailMate manual. Perhaps the only slightly complicated part here is learning the
value you should put into the `userID` field. The example in the documentation
uses the old long-form key ID, which is a bit complicated to get from most
GUIs. Here it's much easier (and also recommended by the gpg2 manpage) to go
with the fingerprint (`--fingerprint` option) without spaces. As a fallback (for
whatever reason) you could also use the 8-char key ID normally printed right in
front of your key when using `--list-keys`.

[hidden-preferences section]: https://manual.mailmate-app.com/hidden_preferences#openpgp--smime
[mailmate]: https://freron.com/
