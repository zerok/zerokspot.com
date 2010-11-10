Google Developer Day 2010
#########################

On November 9th the `Google Developer Day`_ road tour made a stop in Munich,
Germany, which luckily is close enough to Graz that I could attend without
giving the trip another thought.

The event then took place in the `M,O,C`_, a large conference and trade-show
location in the north of Munich. The only downside of that location was that I
couldn't find any `Etap hotel`_ nearby. That said, luckily the Bavarian capital
has a great public transport system with incredible low fares if you buy
day-passes: EUR 12.80 for three days :D

-------------------------------------------

The Prelude
===========

And since I already arrived on Monday, a 3-day-ticket was perfect :-) So once
I got to Munich and had my backpack left behind at my hotel I went back to the
central station to meet with Andreas_ and Fabian_ for some nice food at the
`Paulaner im Tal`_ were we also met quite a few other people from the
conference :-)

After some beer and great food, the way back to the hotel went sadly not as
smooth as the other way around. Not because of all the beer but because the
local public transport out to my hotel had some problems and so I had to wait
about 40 minutes before I could finally board my train.

The Intro
=========

The conference day started with a bit of a disappointment. `Google`_ had originally
announced that the first 200 attendees would get some kind of present. Being
limited to only a firth of the total number of attendees not few started
speculating that there might be something really nice to get if you managed to
get up early. So when I arrived at 0640 I was (1) really shocked to see
probably 300 people in front of me and (2) a bit disappointed that all those
people who got up even early than me got only some little bug gadget.

But if the plan had been to get people to register as early as possible to be
able to start with the regular programme on schedule, it definitely worked ;-)

So after a really great breakfast with a quite classy (salmon, turkey, ...)
sandwitches the conference could start on schedule with the keynotes.

Sessions
========

Keynotes
--------

Here a handful of Google employees with some guest appearances from `VMWare`_
and the `Technische Universität München`_, gave an overview of the current
product line-up regarding browsers, cloud computing, mobile as well as
collaborations with various partners.

The Chrome_ + HTML5_ part of the keynotes made some really nice points regarding
the evolution of the whole application ecosystem with the world before 2004
being completely desktop focused. In 2004 (+- a year) AJAX_ became widely available
and more and more systems moved to the web and the browser as their platform.
HTML5 was also a big topic with some demos.

The next part of the keynotes was all about cloud computing and AppEngine_. A
couple of new features like full map-reduce support, bulk import and export,
background servers, built-in OAuth_ and OpenID_ support as well as background
servers were announced. Also, AppEngine for Business was presented with
features like an SLA, professional support and a separate security model.

Last but not least came Android_ *without* a preview for the next releases but
with some screenshots of new `Market Place`_ features since this year's
Google I/O:

* improved search

* updates to the market publisher site

* more app statistics including crashes and logs thereof

* cloud to device messaging (push messaging)

* and a new license server

After the keynotes the conference was split into 6 tracks:

* Android

* Chrome & HTML5

* Cloud Computing

* Monetization / Social Web

* Tech Talks

* TUM & CDTM (TU München)

Programming the Web with Native Client
--------------------------------------

:Presenter: `Brad Chen`_

Brad Chen was the first to give a presentation on the Chrome & HTML5 track
about `Native Client`_. The premise here is quite nice. NaCl is basically a
re-imagination of the plugin system which brought you among other Flash and
Quicktime integration with your browser. You can basically embed a binary
(compiled from C, C++ and I think C# for now, but more seem to be in the
pipeline) using the <object /> tag into a website and then access exposed
methods from it using JavaScript.

The goal here is that the system will be as safe and portable as JavaScript
with the performance of or close to native libraries. The API on the plug-in
side also offers access to the embedding page's DOM and events et al.

Combine that with something like the local storage infrastructure in HTML5 and
the power of browser applications could reach a completely new level :D

What's new in AppEngine?
------------------------

:Presenter: `Fred Sauer`_

This talk basically went into a bit more detail regarding the new AppEngine
features presented during the keynotes:

* The `Channel API`_ offers an asynchronous message delivery system (1:1)
  between an application and the user's browser via JavaScript and GWT-RPC
  for the callback.

* `Mapper API`_ is for batch processing operating entirely in the user space
  with use-cases like data exports and report generation. The whole system
  can be rate-configured.

* If you serve multiple websites from a single application, the new
  multi-tenancy support/`Namespace API`_ makes it easier to partition data
  store and memcache for different sites.

* `Matcher API`_

* `Hosted SQL`_

* New features within the admin console include an improved task page (you
  can for instance pause tasks) and custom admin pages.

* Every page needs some `image`_ resizing, doesn't it? Well, now GAE has its
  own API for that at least on the Python side based on PIL.

* Custom error pages

* Increased quotas

* OpenID/OAuth support

Practical HTML5
---------------

:Presenters: `Jeremy Orlow`_, `Malte Ubl`_

OK, I'm pretty sure this is the last general HTML5 feature presentation I'll
ever attend. Two reasons: I've already been at too many and this one was quite
exhaustive. Everything was shown here with the exception of `data attributes`_
and the `local storage infrastructure`_.

Also: The room was packet. It probably had about 200 seats but when the
presentation started there was no square cm available even outside these
chairs (e.g. on the way to the exit, between the exit doors, next to the
podium...). Far too mainstream :-P

Storage, Big Query and Prediction APIs
--------------------------------------

:Presenters: `Patrick Chanezon`_, `Simon Meacham`_

Google now also has a file-`storage`_ infrastructure like Amazon's S3_ but more
tightly integrated and more expensive but with (according to the presented) a
more flexible security model. That said, it becomes interesting when you see
it as a data store for calculations done on massive data sets with BigQuery_
and the new `Prediction API`_, which allows you to train a black box with
value-result mappings and query new data using a REST API.

Big Query is intended for analyzing massive amounts of data in an SQL-like
environment.

What's new in Google Geo: Maps API V3 and Fusion Tables
-------------------------------------------------------

:Presenter: `Mano Marks`_

This was perhaps the most interesting talk for me. First of: The `Maps API
V3`_ no longer requires API keys but distinguishes services based on the
server's IPs. This means that for a site being launched in 3 countries on
different domains you no longer have to have 3 different keys for developers,
for testing servers, for stage servers and for live servers (3 x 4 keys).
Great :D

Basically the whole Maps API was rewritten with performance and mobile being
the focus. There is also no flash anymore and elevation is handled. V3 also
offers some integration with Fusion Tables.

If you've reached the limits of what Spreadsheet can do with regards to large
amounts of data, `Fusion Tables`_ is probably something for you. Each table,
which can be filled for instance from a large CSV file, can have up to 100MB
and there is currently an account cap at 250MB. Fusion Tables is quite tightly
integrated with Maps and allows things like spatial queries for bounding
boxes, radius, nearest neighbor etc.  It also sports some more sharing
settings than the classic docs applications in order to allow people for
instance to use your data but not export the dataset to CSV.

Also part of the demo was the new "Styled Maps" feature of Google Maps
V3/Fusion Tables which allows you to customize the look of your maps even more
and also easily customize certain feature types.

Authentication on the World Wide Web
------------------------------------

:Presenter: `Steven Bazyl`_

The last talk of the day for me was all about handling authentication on
websites and between them. Steven Bazyl gave a quick overview about
technologies like OpenID_ and OAuth_ as well as current work to improve them.
There was nothing new here but a healthy reminder that having a unified login
system like OpenID makes a lot of sense and should ideally make everything for
user easier.

The Outro
=========

After the last presentation there was free beer and free prezels (probably
again only for the first 200 since I didn't get one :-P) in the main hall and
those of us who took part in a quiz organized by the Chrome team got a new
notebook sleeve. While I was still on my quest for a prezel, there was a
blinkendroid_ world record attempt going on. There is even a `video on
youtube`_ of that :-)

Afterwards a small group including `Fabian`_, `Constantin`_, Thomas and myself
(if you were part of this group and don't see your name here, please comment
or drop me a mail!) went out to a nice restaurant called "`Alter Simpl`_" for
some great Bavarian food :D


Conclusion
==========

Concluding I have to say that this was an awesome conference with tons of
great new stuff (for me) and tons of  great people. The only real downside I
could see was the really bad WIFI. EuroPython 2010 in Birmingham had its
problems with the connectivity, but the GDDDE took this to a whole new level.
I think on the whole conference day I was only for an hour or two online. Kind
of ironic for an online-company like Google ;-)

But except for this issue, the conference was really great. Big thanks to
Google and the local organizers as well as to everyone else involved. I would
really love to come again next year :-)

.. _Google Developer Day: http://www.google.com/intl/de_ALL/events/developerday/2010/munich/index.html
.. _Maps API v3: http://code.google.com/apis/maps/documentation/javascript/3.0/reference.html
.. _openid: http://openid.net
.. _oauth: http://oauth.net
.. _fusion tables: http://tables.googlelabs.com
.. _mano marks: http://randommarkers.blogspot.com/
.. _chrome: http://www.google.com/chrome
.. _s3: https://s3.amazonaws.com/
.. _storage: http://code.google.com/apis/storage/
.. _html5: http://www.w3.org/TR/html5/
.. _matcher api: http://groups.google.com/group/google-appengine/msg/40021537e2e58962
.. _google: http://www.google.com
.. _mapper api: http://googleappengine.blogspot.com/.../introducing-mapper-api.html
.. _appengine: http://code.google.com/appengine/
.. _bigquery: http://code.google.com/apis/bigquery/
.. _android: http://code.google.com/android/
.. _m,o,c: http://www.moc-muenchen.de/
.. _image: http://code.google.com/appengine/docs/python/images/usingimages.html
.. _data attributes: http://ejohn.org/blog/html-5-data-attributes/
.. _local storage infrastructure: http://diveintohtml5.org/storage.html
.. _fred sauer: http://twitter.com/fredsa
.. _vmware: http://www.vmware.com
.. _market place: http://www.android.com/market/
.. _brad chen: http://www.google.com/research/pubs/author37895.html
.. _steven bazyl: http://www.google.com/profiles/sqrrrl
.. _jeremy orlow: http://www.linkedin.com/in/jeremyorlow
.. _prediction api: http://code.google.com/apis/predict/
.. _simon meacham: http://twitter.com/simonmeacham
.. _patrick chanezon: http://www.chanezon.com/pat/
.. _malte ubl: http://www.xing.com/profile/Malte_Ubl
.. _hosted sql: http://code.google.com/appengine/business/
.. _namespace api: http://code.google.com/appengine/docs/python/multitenancy/overview.html
.. _native client: http://code.google.com/p/nativeclient/
.. _channel api: http://bitshaq.com/2010/09/01/sneak-peak-gae-channel-api/
.. _ajax: http://en.wikipedia.org/wiki/Ajax_(programming)
.. _etap hotel: http://www.etaphotel.com/
.. _technische universität münchen: http://www.tum.de
.. _constantin: http://twitter.com/consti
.. _andreas: http://twitter.com/mfandreas
.. _alter simpl: http://maps.google.at/maps/place?um=1&ie=UTF-8&q=alter+simpl+m%C3%BCnchen&fb=1&gl=at&hq=alter+simpl&hnear=M%C3%BCnchen,+Deutschland&cid=10726857813957190943
.. _fabian: http://twitter.com/fabian7t
.. _paulaner im tal: http://www.paulaner-im-tal.de/
.. _video on youtube: http://www.youtube.com/watch?v=MhXRAaRFK-o
.. _blinkendroid: http://code.google.com/p/blinkendroid/
