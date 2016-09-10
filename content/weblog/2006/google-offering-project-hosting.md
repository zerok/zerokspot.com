---
date: '2006-07-28T12:00:00-00:00'
language: en
tags:
- devjavu
- google
- hosting
- opensource
- sourceforge
title: Google offering project hosting
---


The first thing I read this morning, was [Marshall Kirkpatrick's post on TechCrunch](http://www.techcrunch.com/2006/07/27/google-challenges-sourceforge-in-open-source-project-hosting/) informing the world about Google's new idea of [becoming a SourceForge competitor](http://code.google.com/hosting/). Since I'm one of those people who experienced quite a lot of problems with SourceForge in the past, I went over there and created a small project (mainly for testing purpose) right away.

The registration of a project is very straight forward. First of all you need to have a Google account (GMail or whatever). After you've logged in, simply hit the "Create a new project" link, give your project a name, short description, description and license (tags are optional) and you're done.

-------------------------------



code.google offers every project an SVN repository for all the data but appearantly no SSH+SVN access for developers. For other people this might not be a problem, but for me it is, since so far I couldn't get HTTPS+SVN to tunnel through my proxy. They also don't offer a decent WebSVN viewer were you can read the changelog of each file but instead rely exclusively on Subversions WebDAV integration.

Also most of the other sections look very stripped down. From what I've seen so far, there is absolutely no project statistic and I also have yet to see an option to "release" packages so that people don't have to use SVN to download a new release ;)

Apart from the gmail-like issue tracker, there are also no means of team communication available. No mailing lists, no forums. Apropos issue tracker: Everyone who wants to make a bug report on any project, has to have a Google account. **Very stupid.**

I also haven't found yet a way to inform google about bugs or problems with their service other than a [discussion group](http://groups.google.com/group/codesite-discuss) and an email address. Perhaps using their own issue tracker for something like that wouldn't be such a bad idea ;)

So if stability is Google's only selling point on this, I think they will have a very hard time on the market esp. with packages like [DevjaVu](http://devjavu.com/) out there, that are offering Trac+SVN.