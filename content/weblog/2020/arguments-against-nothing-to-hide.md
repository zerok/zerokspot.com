---
title: "Arguments against \"nothing to hide\""
date: "2020-04-30T07:43:50+0200"
tags:
- privacy
- 100daystooffload
- covid19
---

This is a reply to Jan-Lukas Else's post titled ["Nothing to hide"](https://jlelse.blog/thoughts/2020/04/nothing-to-hide/).

During the current crisis and especially over the course of the last couple of weeks more and more initiatives popped up with proposals for so-called "contact tracing apps". The idea here is to shorten the time it takes to find all the people someone who has been positively tested with COVID-19 might have infected. There are a couple of different approaches such apps can take but one important aspect is whether they store the collected contacts in a centralised (on a server somewhere) or decentralised (only on the smartphones themselves) fashion. I don't want to go into the details of this particular discussion, though, as others like the [CCC](https://www.ccc.de/en/updates/2020/contact-tracing-requirements) and [nyob](https://noyb.eu/en/data-protection-times-corona) have already done so.

What I want to focus on here, though, is an age-old argument that comes up whenever privacy and its protection is relevant (hint: always) and that causes me immense head-ache whenever I hear it: "But I have nothing to hide!". In the context of contact tracing the relevant data is mostly related to your position relative to other people and your current health status (are you infected/do you have symptoms).


## Really nothing?

First off, the term "nothing" is most likely wrong for pretty much anybody on the planet. At least I don't think there is anybody who wants their medical history released for the whole world to see.

So, let's narrow it down to information about who you were close over the course of the day. This piece of information might appear to be quite harmless at first glance. After all, everyone with an active phone has their position transmitted to the cell towers of your provider whenever you have reception. The level of detail here is not high enough, though, to really be certain that you were in close proximity of another human but let's just assume, this were solved somehow to a certain degree of precision and that kind of information were now at your disposal.

But isn't the entity that I send this data to one of the "good people"? It doesn't really matter.


## Data without expiration date

So (continuing our assumptions), you now have information about who you've been in contact with over the course of the last month stored on a central server. The service you are using also stores information that makes identifying each tracked person possible (through phone number, IMEI, or even directly with your name and address).

If everything goes as planned then these collected contacts will expire after a month or so. We are still talking about software here, so you should not assume that everything goes according to the plan, but anyway. Everyone in your country or even on your continent with a smartphone uses this service. This makes it quite a valuable target for attackers. If someone were able to extract snapshots of the collected data your data all of a sudden (1) would no longer expire and (2) could be used outside of the original use-cases despite you only giving it to the "good people".


## But the data was encrypted!

(1) also has the implication that even if the service provider encrypted all the data at rest (incl. snapshots etc.) that encryption might no longer be effective against some nefarious uses of the data.

Remember: The data no longer expires, so if the attacker has a bit of time (and money), they can wait until a weakness in the used encryption system is found or even brute-force it (depending on the time they have). Some kind of data might be relevant even years after it has been collected. The people who you interacted with fall in that category.


## But how critical are contacts?

Let's just name a few things that an attacker might deduced about you from the people you interacted with:

- Did one of them commit a crime around the time you were in contact with them? Are you perhaps complicit in that?!
- Why were you in contact with that other people last night when you actually told me you were on a business trip?!
- Since you talked with those people at that time, are you a member of a certain party? What's your political leaning?
- One pastor, two ministrants: You were at a church. What's your religious leaning?
- At 10:00am you were in close proximity to 3 employees of that grocery store chain. You see them every 2 days. This helps deduce a bit of your shopping habits.

Keep in mind, we are talking here about situations where the data that you provided thinking you'd help someone is actually abused either by that entity or someone else entirely.

Once you've released control over certain data, you can no longer get it back. Once your data is out there, it's out there and there is nothing you can do. So, even if you have "nothing to hide" *now*, you actually lose that choice for as long as that data might be relevant to someone.

That's why privacy is important.
That's why it's important to only share as little data is is necessary to solve a given problem.

Please think about that the next time your insurance company tells you they'll give you a discount if you use that [hot new fitness tracker](https://www.theverge.com/2018/9/26/17905390/john-hancock-life-insurance-fitness-tracker-wearables-science-health). The scenarios I wrote about are mostly worst-case ones yet not unrealistic IMO.

