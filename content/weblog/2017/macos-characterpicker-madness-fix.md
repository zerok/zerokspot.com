---
date: 2017-01-18T20:45:25+01:00
title: "Fixing macOS's CharacterPicker madness "
tags:
- apple
- macos
---

Ever since updating to macOS Sierra I had a weird issue: the character picker
("Emoji & Symbols") behaved erratically. I could only open it every couple of
minutes and then it might appear out of nowhere whenever I focused a native
text input field. Today I was close to simply re-installing the whole machine as
I was also facing some other issues but decided to check first if the behaviour
could be reproduced in a fresh account.

But before that I went through the list of third-party kernel extensions to see
if anything might be relevant. `kextstat | grep -v com.apple` is really nice for
that 😉 Just don't `kextunload` any extension directly 😉

After removing some old extensions of apps I had previously uninstalled the
manual way (and restarting as I definitely kextunloaded things I shouldn't
have), I tested if I could still reproduce the issue. Yep, CharacterPicker was
still drunk. So on to creating a new dummy account and another
reproduction-attempt.

Turns out, it was gone. Plans for operation "Nuke and rebuild" were cancelled
and I started sifting through `~/Library`. I even moved that whole folder but
"Emoji & Symbols" was still hanging in the Activity Monitor. I dug a bit deeper
and learnt that the `CharacterPicker.app` (which is just localized with that
name) was launched or at least controlled by a background service:

```
launchctl list | grep Char
605	0	com.apple.CharacterPicker.FileService
```

Whenever I had that issue, that service would die a horrible death after a
minute or so with the exit code -9. But during that time there was another
service called `com.apple.CharacterPaletteIM.<some number>` visible in that
listing. Googling for `com.apple.CharacterPaletteIM` on DuckDuckGo lead me to a
support article by apple
titled
["Mac OS X Leopard: Keyboard Viewer, Character Palette does not appear"][sup].

Hm... there are most likely no traces of that OSX on my current machine simply
because I had a completely different backup-strategy back then. Anyway, I opted
to follow the steps mentioned there after inspecting the mentioned folders
first:

```
sudo rm /System/Library/Caches/com.apple.IntlDataCache*
sudo rm /var/folders/*/*/-Caches-/com.apple.IntlDataCache*
```

The latter wasn't relevant as I didn't have any matching files, though. After a
reboot, everything was working fine again 😄 And yes, there was definitely some
"Migration Assistant" in the past of my current setup. I'm just not sure if I
used it for Sierra or El Capitan.


[sup]: https://support.apple.com/en-us/HT203221
