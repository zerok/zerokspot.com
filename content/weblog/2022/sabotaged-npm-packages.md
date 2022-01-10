---
title: Sabotaged NPM packages
date: "2022-01-10T21:16:56+01:00"
tags:
- sabotage
- development
- opensource
incoming:
- url: "https://chaos.social/@zerok/107600077629480894"
---

A couple of days ago, Marak Squires released version 1.4.44-liberty-2 (and later also 1.4.1 and 1.4.2) of the [colors](https://www.npmjs.com/package/colors) package onto NPM. This is a simple utility library that makes it easy to write output onto a terminal in different colors and formats. It’s widely used as dependency throughout the node eco-system including (for instance)...

- forever
- redux-i18n
- yeoman

...as well as by various tools by companies like Microsoft or Atlassian. 

So far so good... but the latest changes actually didn’t fix any bugs nor improve anything. Instead, they included endless loops basically printing garbage to the terminal and therefore blocking the code that is using the library. Sadly, that was not a bug but most likely intentional given the comments by the author himself in the code and especially [here](https://github.com/Marak/colors.js/issues/285).

I have no idea what happened prior to that. Was he just annoyed that others including large corporations were [using his code and not giving anything back](http://web.archive.org/web/20210704022108/https://github.com/Marak/faker.js/issues/1046)? Did something tragical happen in his life recently? No idea, but there are some thing you simply don’t do as a software developer and especially not as an open source maintainer. Intentionally sabotaging code that you yourself basically donated to the world and that is used by others is quite high on that list.

Whatever happened, [DABH](https://github.com/Marak/colors.js/commits?author=DABH), the person who has been maintaining the repository since 2018 (and yes, there was no mainline commit to this repository by Marak Squires since 2018), is now trying to somehow rescue what is left of the code that was in the code.

Some folks now argue that Marak Squires had every right right change the code and do what he did. Legal? Perhaps, but I’m not a lawyer. Morally, I don’t think so. The fallout is just huge, the effect on other people either developing at a large corporation or just doing it in their spare time is immense. Acting as if this all is just there to make people aware of what happened to [Aaron Swartz](https://en.wikipedia.org/wiki/Aaron_Swartz) makes it even worse. Everything about this is just wrong.

Yes, open source has a huge funding problem and it’s not like we didn’t have enough reminders of that over the course of the last couple of years. Going out, actively and knowingly sabotaging other people’s work whose only fault it was to trust a piece of code that you shared years ago and that had been well-maintained by others for multiple years is just a bad move. Bad for everyone involved.

Especially since there would have been other ways like changing the license. The original author of the *colors* package may certainly have the attention of parts of the developer community, but I guess everyone will in the future think twice about giving him any serious job.

One good thing that might come out of all of that is that hopefully more people will look harder at what’s included in updates to packages not changed for 2 years and pin their dependencies. But that’s about it. Securing (financially *and* technically) is hard and we are just doing a pretty bad job there all around, but you shouldn’t make it worse. Make the situation better, instead!

(I originally stumbled upon this thanks to this article on [bleepingcomputer.com](https://www.bleepingcomputer.com/news/security/dev-corrupts-npm-libs-colors-and-faker-breaking-thousands-of-apps/).)
