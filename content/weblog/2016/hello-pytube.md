---
date: '2016-06-14T17:27:06+02:00'
language: en
tags:
- python
- pytube
- pyvideo
- pycon2016
- sprints
title: Hello PyTube!
---

For many years the Python community has had an awesome website called
[pyvideo.org][] which tried to aggregate all conference videos on a single
site. In January [Will Kahn-Greene][wkg], one of the awesome people behind it,
announced, that work on the project was [about to end][]. Then in March he
continued with a lengthy [status update][] announcing the [pyvideo-data][]
repository. That repository contains a data dump of all the collected videos and
should act as a starting point for other people to work with.

I hadn't heard anything about that up until [PyCon 2016][] when
[Cameron Dershem][] gave a lightning talk asking people to join him during the
sprints if they wanted to work on pyvideo and related projects. So the
conference days went and the sprint days came and all of a sudden there were
about 15 people working on pyvideo-related topics!

By that time [Paul Logston][plog] had already created a statically generated
site using the pyvideo data-set and [Pelican][] called [PyTube.org][]. Quickly
teams started to form and within an hour we had a website team that focused on
pytube.org, a scraper team, working on importing new content from conference
websites, and an archiving team, trying to come up with solutions to prevent
conference videos from disappearing from the net as did happen when blip.tv went
dark.

Over the following days pytube grew and shortly after the sprints the
maintainers of pyvideo.org even marked their own data-repository to be
deprecated in favour of the pytube-fork. A repository we had set up mainly
because at that time no-one in the room had commit rights to upstream data store
😉

Now, more than a week after the sprints have ended, there are still a handful of
us committing code every day and I'm still amazed how far we've come from the
minute we all entered that sprint room in Portland.

That being said, there is still a lot of work ahead. Tons of conferences are
still missing, for many the metadata could be improved (who gave which lightning
talk?), the website could use the screen-estate more efficiently, ...

If you want to help, simply go to the [pytube org][] on Github and pick what
interests you 🙂

Big thanks to everyone who was at the sprints and [Will][wkg] and
[Sheila][codesquid] for their support and all the work they've put into pyvideo!

[cameron dershem]: http://pinkhatbeard.com/
[wkg]: http://bluesock.org/~willkg/blog/
[pyvideo-data]: https://github.com/pyvideo/pyvideo-data
[plog]: https://github.com/logston
[pelican]: http://blog.getpelican.com/
[pytube.org]: http://pytube.org/
[pytube org]: https://github.com/pytube/
[status update]: http://bluesock.org/~willkg/blog/pyvideo/status_20160316.html
[about to end]: http://bluesock.org/~willkg/blog/pyvideo/status_20160115.html
[pyvideo.org]: http://pyvideo.org
[pycon 2016]: https://us.pycon.org/2016/
[codesquid]: https://github.com/codersquid
