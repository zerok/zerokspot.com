# Using Launchd for Cronjobs

Recently I needed to schedule some re-occurring tasks on one of our office-Mac
Minis. While I had used [launchd][] before for the obvious service tasks (think
SysV init.d scripts) I hadn't tried it's cron-like features. Seems like it might
safe me some time in the future as using launchd here turned out to be quite
straight forward! Not that cron isn't, but I wanted to try something different
here 😉

-----------------

The process is basically identical to what you do for programs that you want to
keep running indefinitely with launchd. The only real difference is that you
also set either the `StartCalendarInterval` or `StartInterval` property and
disable things like `RunAtLoad`, `KeepAlive`, and `WatchPaths`.

A simple yet very annoying example job would say "Hello World" every morning at
09:51:

```
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple Computer//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
  <key>Label</key>
  <string>me.h10n.helloworld</string>
  <key>ProgramArguments</key>
  <array>
    <string>/usr/bin/say</string>
    <string>Hello World</string>
  </array>
  <key>StartCalendarInterval</key>
  <dict>
    <key>Hour</key>
    <integer>9</integer>
    <key>Minute</key>
    <integer>51</integer>
  </dict>
</dict>
</plist>
```

(Put that into `~/Library/LaunchAgents/me.h10n.helloworld.plist` and execute
`launchctl load ~/Library/LaunchAgents/me.h10n.helloworld.plist` to experience
it in all its beauty...)

`StartCalendarInterval` offers everything you're used to from cron. It supports
fields for `Hour`, `Minute`, `Month`, `Day`, and `Weekday` of a date and the
values are integers with the same semantics as the integer variants in cron. If
no value is set, it's the equivalent to the wild-card character.

Sadly, it looks like defining ranges involves putting all matching intervals
into a big array in order to keep the annoyance limited to weekends:

```
<key>StartCalendarInterval</key>
<array>
  <dict>
    <key>Hour</key>
    <integer>9</integer>
    <key>Minute</key>
    <integer>51</integer>
    <key>Weekday</key>
    <integer>6</integer>
  </dict>
  <dict>
    <key>Hour</key>
    <integer>9</integer>
    <key>Minute</key>
    <integer>51</integer>
    <key>Weekday</key>
    <integer>7</integer>
  </dict>
</array>
```

`StartInterval`, on the other hand, allows you to define intervals in
seconds. So if I wanted to hear "hello world" every 10 seconds, this would do:

```
<key>StartInterval</key>
<integer>10</integer>
```

Pretty much the most annoying thing ever and an instant `launchctl unload` 😉

As I wrote earlier, everything else is the same compared to usual services. The
distinction between the LaunchAgents and LaunchDaemons folders still applies and
you can still put shared files into `/Library/{LaunchAgents,LaunchDaemons}` or
their `/System` counterparts.

If you want to learn more of the details here, take a look at most excellent
[launchd.info][info] or `man launchd.plist` 😃


[info]: http://launchd.info/
[launchd]: https://en.wikipedia.org/wiki/Launchd
