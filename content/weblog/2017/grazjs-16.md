---
date: 2017-06-16T22:41:48+02:00
title: "GrazJS #16"
tags:
- usergroups
- grazjs
- javascript
---

On Tuesday evening was once again
a [GrazJS meetup](https://www.meetup.com/grazjs/events/239618888/) at lab10 with
two awesome talks: an introduction to [GraphQL](http://graphql.org/)
by [Manuel Penaloza](http://manpenaloza.net/) and a presentation about
how [Timeular](https://timeular.com/) is using JavaScript
by [Manuel Zoderer](https://at.linkedin.com/in/manuel-zoderer-b0a87b78). I've
held off from getting into GraphQL for quite some time now so I was extremely
happy about that introduction 😀

## An introduction to GraphQL by Manuel Penaloza

The talk consisted of a theoretical and practical part that even included a
live-coding session. Especially thanks to the latter I think I finally have a
rough understanding of what GraphQL is and how I could use it in some of my
projects. Here are a couple of notes I took during the talk:

- GraphQL doesn't allow us to do anything we weren't able to do before. But it
  provides us with a rather nice, generic API gateway for data-centric APIs.
- It helps getting around classic issues with RESTful APIs: over-fetching,
  under-fetching, imperative approaches, API evolution.
- Used internally at Facebook since 2012 and open-sourced in 2015.
- While things get easier on the client-side, there is a considerable amount of
  work initially required on the server in order to abstract your APIs with
  GraphQL. In the long term, this pays off, though.
- While "QL" might imply that GraphQL can only be used for retrieving data from
  the server, you can also
  execute [mutations](http://graphql.org/learn/queries/#mutations) and
  subscriptions.
- [express-graphql](https://github.com/graphql/express-graphql) offers an
  express middleware for GraphQL which also allows you to explore an API through
  a debug console (if you set the `graphiql` flag.
- GraphQL exposes a tree of objects that can be queried. Each object is defined
  by a strictly typed schema with its own type-system. The existence of such a
  type-system also allows for documentation to be automatically generated and
  the graphiql console to support things like auto-completion.
- Each schema object has a `resolve` method which is called in order to fill the
  object with actual data from an underlying API. In fact, each field within the
  object can also have its own resolution mechanism. As this can easily lead to
  n+1 queries to the backend APIs there are a couple of community solutions for
  that. In general, there are quite a few tools around schema definitions
  provided by [graphql-tools](https://github.com/apollographql/graphql-tools).
- Axios is a library that abstracts http requests as promises
- The `fields` property within a schema object can also be a function in order
  to be able to reference schemata that are defined later on in the code.
- [relay](https://facebook.github.io/relay/)
  and [apollo](https://www.apollodata.com/) are popular libraries to include
  GraphQL in clients for React and other tool-chains.

After this talk I'm quite excited to find a project where I can work with
GraphQL. Now I just have to find that and the time for it, though 😉


## JS all the way by Manuel Zoderer

The second talk by Manuel Zoderer gave an overview about how JavaScript is used
by [Timeular](https://timeular.com/) for their backed, desktop, and mobile
applications around Zei. After a quick introduction of their product and it's
history (crowdfunding, partnership with a company in Munich, ...) he went right
into the nitty-gritty details! Turns out, they are
using [React Native](https://facebook.github.io/react-native/) for the latter
and [Electron](https://electron.atom.io/) for the desktop applications!
Especially interesting for me here was that the integration with native
libraries in React Native turned out to be quite easy but there are some
performance hits when switching from JS to native and vice versa.

He also mentioned, that while all applications are using Bluetooth LE, for
Windows they still have to ship a custom dongle as many slightly older laptops
and Windows versions don't support this standard. On the software-side, though,
the desktop clients for Windows and Mac share most of their code base thanks to
Electron and only differ in the driver implementation.

As for the rest of their technology stack, they are also using Java EE, Kotlin,
and RabbitMQ for the backend services and NodeJS with hapi for integrations of
3rd-party services like JIRA. It also sounds like most if not all their REST
APIs (internal and external) are documented using [Swagger](http://swagger.io/).

If I just had a Zei to play around with 😉

## ... and the rest

Sadly, as I was already extremely tired and not feeling all that well, I left
after that talk. For everyone else there was still a lot going on with at least
one round of OpenSpace-discussions and drinks and snacks.

I'm already looking forward to next month's meetup on 8 August and
the [MeetTheMeetups event](https://meet-the-meetups.org/events/graz-2017/) in
the following month!
