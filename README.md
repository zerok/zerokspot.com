# zerokspot.com

This repository contains pretty much all that is reachable publically on
<https://zerokspot.com/>. Under the hood this project is based on the following
technologies:

- Hugo (as the static site generator that produces all the HTML pages that are
  then served)
- [Vue.js][vue] (for implementing the search frontend)
- [Webpack][] (for building the search- and webmentions frontend)
- [SASS][]
- Go (for implementing the search backend)


## How to build

You will need to have the following tools installed:

- [Hugo][] (v0.64 or newer)
- [Go][] (1.13 or newer)
- [NodeJS][] (12 or newer)
- [Yarn][] (1.x)
- [GNU Make][make]

Once you have all that installed, run the following commands to get the website
up and running locally (for now without search or webmentions):

```
# Installs all the NodeJS modules required to build the JS and CSS parts:
$ make prepare

# Build the JS and CSS:
$ make frontend
```

Everything is ready now! The next command will build the HTML and serve it (by
default via <http://localhost:1313>).


## Search

Search is done using a simple Go backend which proxies calls to [Algolia][].
The search index that is stored there is updated by a CI job after every push
to the master branch.

[algolia]: https://algolia.com
[hugo]: https://gohugo.io/
[nodejs]: https://nodejs.org/
[yarn]: https://yarnpkg.com/
[make]: https://www.gnu.org/software/make/
[vue]: https://vuejs.org/
[webpack]: https://webpack.js.org/
[sass]: https://sass-lang.com/
[go]: https://golang.org
