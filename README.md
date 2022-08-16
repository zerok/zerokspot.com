# zerokspot.com

This repository contains pretty much all that is reachable publically on
<https://zerokspot.com/>. Under the hood this project is based on the following
technologies:

- Hugo (as the static site generator that produces all the HTML pages that are
  then served)
- Go (for implementing the search backend)


## How to build

You will need to have the following tools installed:

- [Hugo][] (v0.64 or newer)
- [Go][] (1.13 or newer)
- [GNU Make][make]

Once you have all that installed, run the following command to get the website
up and running locally:

```
$ make run
```

Everything is ready now! The next command will build the HTML and serve it (by
default via <http://localhost:1313>).


[hugo]: https://gohugo.io/
[go]: https://go.dev/
[make]: https://www.gnu.org/software/make/


## Limitations

These instructions don't cover the Webmention and Mastodon integration. For
these please take a look at the [webmentiond][] and [retoots][] projects.

[webmentiond]: https://webmentiond.org/
[retoots]: https://github.com/zerok/retoots
