---
date: '2011-02-15T12:00:00-00:00'
language: en
tags:
- convore
title: 'Convore: Chatty forums'
---


A couple of days ago [Convore][c], a new startup/product by Leah Culver, Eric
Florenzano and Eric Maguire, went public. So, yes: I'm once again [late](http://techcrunch.com/2011/02/09/convore-wants-to-be-the-easiest-group-communication-app-yet/) to the
"there is a new service out there, let's write about it"-party, but that's
mostly because I wanted to first use it for a couple of days and perhaps find
some use-cases for me before writing just a paragraph and be done with it ;-)

So ...

-------------------------------------------------------------------------------

## What is Convore?

Perhaps the best way to describe it -- in my opinion -- is if you switch the
taste-alerts in your brain off for a second and think about what would happen
if [Twitter][t], [IRC][i] and classic bulletin boards ([phpBB][p], [vBulletin][v], ...) had a
child. You'd probably end up with a service that emphasises short posts, that
are updated in real-time and is persistently organised in topics. And this
gets you pretty close to what Convore is. It combines the emphasis on short
messages (Twitter) with a group communication style reminiscent on forums and
the real-time flow of the conversation available with chat services.

The idea is that you either create a group or join one. If someone creates a
new post, you will see an update in the menu bar (as shown in the screenshot).

<figure>
    <img src="/media/2011/updates.png" alt="" />
    <figcaption>Updates to groups you've joined, are shown in real-time in the menu bar and a flyout shows the topics in these groups updated.</figcaption>
</figure>

But these updates are also visible right in the discussion view where posts
show up chat-style as they happen.

Compared to a forum there is right now no special markup. The only thing I
could find so far is the usual auto-link-generation and links to people's
profiles if referred to via @username.

## So what's the use-case?

At first I wasn't quite sure what this combination could be used for.
Besides, do I really need another way to communicate with people? I'm already
on Twitter, IRC, tons of mailinglists and from time to time also on some
boards.

But after using it for a couple of days I noticed that it's really hard to use
it on and off. I've so far only subscribed to a couple of groups and within
just two hours or so my unread-messages counter always hits the indicator
limit of 300 (i.e. "300+"). So I thought: OK, this is probably not a forum,
nor would I probably use this as a replacement for a general chat room.

Right now I can see two primary use-cases, at least for me .\.\. or simple
what I'd really love seeing it being used for:

First, it could really be great for smaller groups like your local usergroup.
Not for making announcement, for which mailinglists and Twitter are IMO still
the best places (or naturally your website), but for quick discussions, polls
etc. The real-time aspect of Convore just makes this much more pleasant than
a mailinglist or a forum. The first thing I did after joining, I created a
group for our small [Python usergroup in Graz](https://convore.com/pygraz/) ;-)

The second use-case is probably really weird, but I can definitely see this
for support questions in OpenSource projects. Currently the primary channels
for this are either forums, mailinglists or IRC channels. As before, the
real-time aspect of Convore beats the first two choices. But what's wrong with
chats?! Nothing if there are enough people online all the time that can
actually help with answering questions. Sadly, this is usually not the case
because people usually also have a life outside of the channel so channels are
usually filled with lurkers or people simply talking about other topics and
the original question will probably get lost if not answered right away.

Convore, on the other hand, offers a persistency to this chaos and might just
be the right middle-ground between forums/mailinglists and chats. People on IRC
are already used to using pastebins for posting code, so markup-wise Convore
is ready :-)

## But what is missing?

Feature-wise there are probably still a few things Convore could do to improve
the information overload. As I wrote about, it only takes 2-3 hours for my
update-counter to reach 300+. Quite high on my list is a way to exclude
certain topics from my updates-count (as discussed in [this topic](https://convore.com/feedback/mute-for-topics-in-subscribed-groups/)).

Another feature that IMO is required for anything real-time is for group
admins to be able to ban/block people from their groups. Some people are just
annoying ;-)

What's definitely not missing is something that I wanted for Google Groups for
a really long time: an [API][a]. I'm really curious what people will created based
on it. There are already a couple of folks teaming up to potentially create an
iOS app :-)

So yes, I'm quite excited about Convore right now, mostly because I simple
enjoy being there. I like the tech, the design and for the most part the
people there. All the ingredients for a great community :-)

[c]: http://convore.com
[a]: https://convore.com/api/
[t]: http://twitter.com
[p]: http://www.phpbb.com
[v]: http://www.vbulletin.com/
[i]: http://en.wikipedia.org/wiki/Internet_Relay_Chat
