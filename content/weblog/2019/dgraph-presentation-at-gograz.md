---
title: "Dgraph presentation at GoGraz"
date: 2019-09-10T21:58:19+02:00
tags:
- dgraph
- gograz
- usergroup
- graz
---

Yesterday evening I gave a quick presentation at our local [Go
usergroup in Graz][gg] about the graph database [Dgraph][]. In the
previous months I've played a little bit around with this system in
the context of a genealogical software I'm working on right now and
now that I got familiar enough with it I thought it might be time to
share what I've learnt. Sadly, I couldn't really get into some of the
details on how to work with Dgraph from the point of view of a Go
application or it's features as a distributed service but I hope the
presentation was still a good appetizer and enough to get new folks
(1) interested in graph databases in general and (2) them started 🙂

If that might like something for you, then you can find the complete
material as Jupyter Notebook on [Github][gh]. I tried to keep the
Python in that notebook as hidden as possible in order to demonstrate
things like querying and mutating using Dgraph's [HTTP
API][ha]. Sadly, I couldn't get a [Go kernel for Jupyter][gn] to work
reliably in the availble time so I just fell back to old habits 😉

I hope you can still enjoy what is there!

[dgraph]: https://dgraph.io
[gg]: https://gograz.org/meetup/2019-09-09/
[gh]: https://github.com/zerok/dgraph-intro-presentation/blob/master/presentation.ipynb
[ha]: https://docs.dgraph.io/clients/#raw-http
[gn]: https://github.com/gopherdata/gophernotes

