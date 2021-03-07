---
date: '2009-05-09T12:00:00-00:00'
language: en
tags:
- django
- djangocon
- conferences
- traveling
- prague
title: EuroDjangoCon 2009
---


<img src="http://farm4.static.flickr.com/3568/3506612543_371b5b7b2f_m.jpg" style="float: left; margin: 0 15px 15px 0" alt="" />

Well, that's it. [EuroDjangoCon 2009][] in Prague is over and most of us are on
their way home right now, already there or will leave tomorrow. For me, tech
conferences are just great. There is always so much new stuff to learn, on so
many cool projects to work on and so many really nice and clever people to
meet with.  And then just imagine a conference in the name of the [Pony][] in
Prague with its beautiful old-town, low prices, great beer and subbed American
movies. A perfect match.

[EuroDjangoCon 2009]: http://euro.djangocon.org/
[Pony]: http://djangopony.com/


-------------------------------


But let me write this at least a little bit chronologically and perhaps get
the not-so-perfect parts out of the way first.


## About trips, hotels and pain

[Armin][], [Florian][] and I arrived on Sunday right after the only tourist info
office near the train station had closed down. So there we were, standing
close at the [Holesovice trainstation][] with some fresh money from the ATM but without any
way to buy some metro tickets with it. So we took a taxi and got ripped off
quite heavily with an 800 kc ride to our hotel. 

[armin]: http://lucumr.pocoo.org/
[florian]: http://djangopeople.net/apollo13/
[holesovice trainstation]: http://maps.google.com/maps?f=q&source=s_q&hl=en&geocode=&q=holesovice+station+,+praha&sll=50.105738,14.437505&sspn=0.004163,0.009656&ie=UTF8&z=16

And since nothing can be easy, the hotel reception couldn't find our room
reservation. Only after about 10 minutes of looking and *us* in the end
actually finding it on *their* list they agreed to give us our room. A room
without any way to open/close the window, thank you very much, since the
window handler was lying next to the window. Our complaints resulted in some
annoyed looking person knocking at our door and asking if we wanted the window
opened or closed.  Once we decided to leave it closed for now and he using an
Allen wrench to do that he went off watching his soccer match again ;-) This
was their only solution for the broken window for the whole week.

Luckily, the infrastructure of the whole vicinity was actually very nice, with a
MacDonald's right next door, a [Tesco][] grocery store (24:7) just on
the other side of the street and the next tram station only about 5 minutes
away. 

[tesco]: http://en.itesco.cz/en_cz/stores/praha/hm_praha_eden


## The conference

The next day started with a really nice breakfast incl. eggs, ham and sausages
... and naturally the conference itself. [Zed Shaw][] kicked it off with great
keynotes and [James Tauber][] reiterated some patterns actually everyone who wants
to build something with Django should really memorize. After James, [Andrew Godwin][]
convinced me that I should give [South][] another try. The afternoon for
me was mostly filled with talks by [Paul Smith][] about [GeoDjango][] and in general
how [EveryBlock][] uses the various GIS systems out there, and [Simon Willison][].
Simon's talk actually marked (once again) the absolutely highlight-talk of the
conference for me where he ranted for nearly an hour about various design
decisions in Django that turned out to be the wrong ones.

[south]: http://south.aeracode.org/
[everyblock]: http://www.everyblock.com/
[zed shaw]: http://zedshaw.com/
[james tauber]: http://jtauber.com/
[paul smith]: http://www.pauladamsmith.com/
[geodjango]: http://geodjango.org/docs/
[simon willison]: http://simonwillison.net/
[andrew godwin]: http://www.aeracode.org/

The conference as a whole covered basically all the current topics of Django
with [Honza Král][] speaking about his task for version 1.2, model validation,
[Eric Holscher][] presenting some ways of how to test your Django projects, 
[Frank Wierzbicki][] giving a status update of Django on [Jython][], James Tauber updating
his talk about the "State of Pinax" and much much more. So, I guess, there must
have been something from everyone :-)

There was also about a ton of camera equipment, so I hope we will see some
nice videos soon.

[honza král]: http://djangopeople.net/king/
[eric holscher]: http://ericholscher.com/
[frank wierzbicki]: http://fwierzbicki.blogspot.com/
[jython]: http://www.jython.org/Project/


## The sprints

Days 4 and 5 were dedicated to sprinting on various tasks mostly for Django
and [Pinax][]. Since I had wanted to get into Pinax for the last couple of
months, I focused mostly on Pinax and also a little bit on [django-piston][] 

[pinax]: http://pinaxproject.com/
[django-piston]: http://bitbucket.org/jespern/django-piston/overview/

<figure>
    <img src="http://farm4.static.flickr.com/3623/3513177706_76d6aae70b.jpg" alt="" />
    <figcaption><p>Sprint time is great ... also for sticker collectors like Jannis. And now    we even have Bitbucket stickers :D</p></figcaption>
</figure>

Within Pinax especially the new task manager of the `cpc-website project`_
received a lot of love like an attachments feature and a pastebin by Martin_
and various performance improvements by Stephan_. Since there were tons of
open documentation tickets and some other parts that hadn't been documented at
all yet, I used the first day of focusing on those tasks while the 2nd day saw
some baseline Microformats_ support for the profile application within Pinax
and Piston getting unittests with the help of Buildout.

[cpc-website project]: http://github.com/pinax/code.pinaxproject.com/tree/master
[martin]: http://mahner.org/
[stephan]: http://www.sjaekel.com/
[microformats]: http://microformats.org/

<a href="http://www.flickr.com/photos/zerok/3516676250/" style="float: left; margin: 0 15px 15px 0"><img src="http://farm4.static.flickr.com/3335/3516676250_1897f48a5d_m.jpg" alt="Centrum Holdings" /></a>

The location in Honza's company [Centrum Holdings][] was probably the best
sprint-location ever: Large tables, modern furniture, broadband Wifi
connection (broadband in the *real* sense), sponsored drinks, food and
enough toilets. What else do you need :D

[centrum holdings]: http://www.centrumholdings.com/en/

Well, perhaps some faster pizza services, but that's about it.


## The entertainment

It's not like learning new stuff and coding in between isn't entertaining but
if you're in Prague, you really should get to know some of the local bars and
restaurants.

For the first evening [Robert][] organized a party at a local beer bar, 
[The Pub World][]. What made it so special besides the guests? Every table had its own
beer taps and a high-score listing with entries for each table. I so want
something like that in Klagenfurt ;-) The bar also provided 2 small buffets:
One with cold, one with warn dishes, but both really good.

[robert]: http://www.siudesign.co.uk/
[the pub world]: http://www.thepubworld.com/

<figure>
    <img src="http://farm4.static.flickr.com/3539/3506609905_239c34f453.jpg" alt="" />
    <figcaption>
   <p>Beer break. Imagine how hard it is to take a picture of people in a 
   beer bar <em>not</em> drinking beer ;-) (on <a href="http://www.flickr.com/photos/zerok/3506609905/">Flickr</a>)</p>
</figcaption>
</figure>

The second evening had no event planned, so we all just went off to get
some food and later on perhaps meet up in some Jazz bar. Getting the food was
actually, despite out group being about 20 people IIRC, not a problem. With the
first restaurant we found a simple yet good, friendly and extremely cheap
place to eat (sorry but I can't remember the name anymore).  Just an example:
A large beer, Gulasch (so hot!) with some dumplings and a dessert for only
about 4 EUR! Finding that Jazz bar afterwards turned out to become an odyssey
through the Praha 1 Stephan and I could at least take some nice pictures. But
that [jazz bar][] was really nice. At that time I regretted it at least a bit
that I had already eaten. Some of the menus looked really delicious.

[jazz bar]: http://www.malyglen.cz/

<figure><img src="http://farm4.static.flickr.com/3343/3507418716_ae689d6304.jpg" alt="" />
    <figcaption>
        <p>Finding a place when you have no clue where it is and or don't have a real
        address but a map can be quite entertaining. The sites compensated that,
        though. (on <a href="http://www.flickr.com/photos/zerok/3507418716/">Flickr</a>)</p>
    </figcaption>
</figure>

On the third day after the last conference session we all went with tram 7 to
the west side of Prague just to end up in front of an already full restaurant,
so we had to split up. Some could still find a place in there, some went
somewhere else and most of [#django-de][] with some friends went to some 
[nice Pizzeria][].  Probably not the cheapest place in town, but definitely good :-)

[#django-de]: irc://irc.freenode.net/#django-de
[nice pizzeria]: http://www.cortediangelo.cz/

Sprint days are mostly self-entertaining, yet on the 2nd sprint day we still
tried to first find some cheap food and then a nice bar for some cheap beer.
While we failed on both tasks (one bar even wanted 149 kc for a beer!!!)
Arthur, Bastian, Jannis, Jesper, Martin, Philipp, Stephan and me had least had
some fun in the [Palace Cinema][] watching the [new Star Trek movie][].

[palace cinema]: http://www.palacecinemas.cz/
[new star trek movie]: http://www.startrekmovie.com/

What's kind of weird in Prague is, though, that most of the restaurants close
really early. The one we found on day 2 wanted to close an hour after we had
arrived (i.e. 21:00). 


## Conclusion

I'm really sad that these 5 days are already over. I met a lot of people I had
only seen on IRC before and we actually got a chance to make a small
German-users-meetup on the 2nd day. If it's possible I'd really love to meet
all of you again in Birmingham and Portland later this year for the
[EuroPython][] and the [DjangoCon][]. I really hope that I'll be able to make it at
least to one of these two conferences.

[djangocon]: http://www.djangocon.org/
[europython]: http://www.europython.eu/

A big "thank you" goes to Robert and Honza (and everyone else involved with
the organization) for making this event possible. It was great :D

