---
date: 2017-05-26T15:32:46-07:00
title: PyCon US 2017
tags:
- pycon2017
- travel
- portland
- conference
---

PyCon usually stays in one city for two years to make it possible for organizers
to re-use the gathered resources for a conference at least once. Sadly, this
means that this year
is [the last PyCon in Portland](https://us.pycon.org/2017/). Or at least the
last one for the next couple of years.

Still, even though I now do most of my work in NodeJS and Go, it was extremely
interesting learning what has been going on in the Python community over the
last 12 months. Sounds like asyncio is finally gaining traction which means I
might be able to switch back from NodeJS for a couple of use-cases 😊

But before the event can welcome 3300 attendees, some things have to be
prepared…

## Bag-stuffing/stuff-bagging

Since I'm usually not the workshop-guy, PyCon starts for me with the
bag-stuffing event, where nearly 100 attendees meet the afternoon before the
main conference to prepare the goodie-bags that are handed out the following
morning to all attendees.

Over the years we have become quite proficient at that. There is one long row of
tables with all the goodies on it, people standing next to that to hand out the
goods and then two rows of people running from top to bottom with an open bag in
order to receive them and handing the filled bag over to people managing the
"final storage".

<figure>
<img src="/media/2017/pycon-us-2017-bagstuffing.jpg" alt="">
<figcaption><p>The queue for filling all these bags has become a well-oiled
machine over the years!</p></figcaption>
</figure>

And we are also getting faster and faster with every year! This time around it
only took 2 hours to fill the 3000+ bags!

## The conference

Following the bag stuffing and evening reception are three days of talks in up
to five parallel tracks. From social talks, over asyncio, to things like Docker
and Kubernetes there should have been something for anyone.

My personal favorites were these:

- [Design secure APIs with state machines](http://pyvideo.org/pycon-us-2017/ashwini-oruganti-mark-williams-designing-secure-apis-with-state-machines-pycon-2017.html) (Ashwini
  Oruganti from Docker): Out of this talk will most likely come another blog
  post where I look into various state-machine libraries in Python and perhaps
  also in Go.
- [Python @ Instagram](http://pyvideo.org/pycon-us-2017/keynote-pythoninstragram.html):
  Two days before that talk I had re-activated my Instagram account after many
  years of slumbering so this was especially interesting to me. It mostly
  covered Instagram's strategy for moving from Python 2.x to 3.6 and Django 1.3
  to 1.8. If you have a large code-base then this talk might interest you. The
  next large-scale update/refactoring *will* come!
- [Async/await and asyncio in Python 3.6 and beyond](http://pyvideo.org/pycon-us-2017/asyncawait-and-asyncio-in-python-36-and-beyond.html) (Yury
  Selivanov): A really nice overview of what the new language features regarding
  asyncio bring to the table. I cannot wait to play around with it!
- [The dictionary even mightier](http://pyvideo.org/pycon-us-2017/brandon-rhodes-the-dictionary-even-mightier-pycon-2017.html) (Brandon
  Rhodes): Again a great overview about some of the more hidden features of the
  Python dictionary and the way Python 3.6 might help you with getting more
  reproducible behavior.
- [Kubernetes for Pythonistas](http://pyvideo.org/pycon-us-2017/kelsey-hightower-keynote-pycon-2017.html) (Kelsey
  Hightower): This was the absolute highlight of the last day for me. I had
  already known a bit about Kubernetes but it's always such a joy to listen to
  Kelsey Hightower!

There were, sadly, also a handful of talks and lightning talks that made me
leave the room, but overall the talks were great and the lightning talks are
still my absolute highlight of any PyCon 😍

## The dinner party

Every year PyCon also organizes two dinners in special locations around the city
where people can get to know each other better. This year, Ulrich and I managed
to get tickets for the second dinner right on the last day of the main
conference. The event was hosted by our good
friend [Adrienne Friend](https://codingwithknives.com/) at the gorgeous Portland
Art Museum. Delicious food, great company, what more can you want?!

<figure>
<img src="/media/2017/pycon-us-2017-dinnerparty.jpg" alt="">
<figcaption><p>The conference dinner at the Portland Art Museum offered enough
space to get to know people without being too crowded.</p></figcaption>
</figure>

## The sprints

During the sprints, I mostly worked on [PyVideo](http://pyvideo.org)-related
things. I had wanted to work on a search-feature for the site for many months
now but only found time for it now. You can learn more about that
in
[a different post](https://zerokspot.com/weblog/2017/05/23/search-for-pyvideo/)
😊 The deployment of that was a bit more complicated than expected and required
another server-move (thanks [Ulrich]()) but other than that it was a very smooth
process and I hope that feature proof to be useful in the future.

After that I started to play a little bit around with asyncio but, sadly, didn't
get that far as one of my PRs for WriteTheDocs kept on coming back to me. That
being said, I think it was still a very productive sprint despite me not writing
more than a handful of lines of Python. At least I now
have [Pythonista](http://omz-software.com/pythonista/) on my iPad 😉

## Goodbye, Portland

This is probably the most emotional end of a trip I've ever had. Over the last
two year's I've fallen in love with Portland and its people. There is just so
much to love here, from the beers, the food, the donuts, the bridges, the river,
the MAX… It was just the perfect place for PyCon and WriteTheDocs and I hope
that I will get a change to return here one day.

<figure>
<img src="/media/2017/pycon-us-2017-portland.jpg" alt="">
<figcaption><p>Portland from above is even more beautiful. View from the New
Relic offices.</p></figcaption>
</figure>

Big thanks to everyone who made these two events possible and everyone in the
city we came across during the last two years ❤️
