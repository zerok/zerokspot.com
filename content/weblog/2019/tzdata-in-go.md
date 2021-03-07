---
title: "Timezone data in Go"
date: "2019-12-05T20:48:00+01:00"
tags:
- golang
- development
incoming:
- url: https://twitter.com/zerok/status/1202678538098401281
- url: https://chaos.social/@zerok/103256941965229740
- url: https://dev.to/zerok/timezone-data-in-go-2ej3
- url: https://www.reddit.com/r/golang/comments/e6moyx/where_does_the_timezone_data_in_go_come_from/
---

After reading Jon Skeet's excellent [blog post about issues with storing datetimes in UTC][js] inside applications I wondered, how Go was dealing with updates to timezones. In Python this is done using the [pytz package][pytz] which is updated frequently. Go goes a slightly different way:

On Unix systems it tries to load information about the current timezone from one of the following places ([src/time/zoneinfo_unix.go:21][z21]):

```
var zoneSources = []string{
	"/usr/share/zoneinfo/",
	"/usr/share/lib/zoneinfo/",
	"/usr/lib/locale/TZ/",
	runtime.GOROOT() + "/lib/time/zoneinfo.zip",
}
```

`/usr/share/zoneinfo` is a folder that is usually provided by the tzdata package of whatever Linux distribution you're using. This in turn is built based on data provided by the tz database project which is maintained by the IANA. You can get the "original" data and project code on <https://www.iana.org/time-zones>. If you want to get notified when something changes there, there is even an [annoucement mailing list][a]!

If you want to stay up-to-date here (and you definitely want!) then it's probably easiest to just stick with the package provided by your operating system. Looking at, for instance, Debian's [tzdata package for Jessie][d], it still gets updated and is currently at 2019c which is the latest release of the tz database at the time of writing this.

At the other end of the spectrum you can also force Go to explicitly use a path of your choosing by setting the [`ZONEINFO` environment variable][z]. You could even roll your own `zoneinfo.zip` file: The lib/time folder inside Go's source tree contains an `update.bash` file which should come in handy there.

To summarize: If you want to stay up-to-date with your timezone information in Go applications, Go makes that pretty simple by sticking close to what the operating system provides. For edge cases you are able to roll your own version of the tz database, though.

[pytz]: https://pypi.org/project/pytz/#history
[js]: https://codeblog.jonskeet.uk/2019/03/27/storing-utc-is-not-a-silver-bullet/
[z21]: https://github.com/golang/go/blob/50535e6b422ac6b0195f9d3a83607326401cee0b/src/time/zoneinfo_unix.go#L21
[a]: https://mm.icann.org/mailman/listinfo/tz-announce
[d]: https://packages.debian.org/jessie/tzdata
[z]: https://github.com/golang/go/blob/50535e6b422ac6b0195f9d3a83607326401cee0b/src/time/zoneinfo.go#L294
[u]: https://github.com/golang/go/blob/20bf6a495eabad79b7b275d46fc3e11c620b8212/lib/time/update.bash
