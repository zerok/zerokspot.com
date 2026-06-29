---
title: "OWASP Global AppSec EU 2026 in Vienna"
date: "2026-06-29T20:50:00+02:00"
tags:
- owasp
- security
- events
- conferences
- austria
- vienna
---

It took a long time this year but I finally made it to a conference! Last week [OWASP](https://owasp.org/) organised the [OWASP Global AppSec EU 2026](https://owasp.glueup.com/event/owasp-global-appsec-eu-2026-vienna-austria-162243/) in Vienna. My wife had convinced me and so off we went to the [Austria Center](https://www.acv.at/en/) to learn what's new in application security! 

A little disclaimer: I'm a software engineer and not a security engineer, so many of the topics discussed at the event were somehow on my radar but more like "I read it in the news". All of a sudden I was in a room with tons of people who actively work on these issues and ... I obviously learnt a ton and had lots of fun!

## AI everywhere!

As was to be expected, a large number of talks at least touched AI topics. Both keynotes also focused on the new landscape. first from a developer's point of view and how especially code reviewing is suffering from the flood of produced code and then also when seen from a researcher's point of view.

The expo hall was pretty much all AI. It felt like every company had at least some product for either integrating AI into the S-SDLC or providing tooling for monitoring agents.

## Threat modelling gamified

While browsing the OWASP shop in the expo hall, it was hard to miss the large corner of security-related board games by [CyberSec Games](https://cybersecgames.com/). From what I could tell, most of them were about helping teams doing threat modelling in various environments, from ops to dev teams with or without AI.

One of them for some reason caught my eye: [OWASP Cornucopia](https://cornucopia.owasp.org/), which focuses on generating security requirements in various categories like LLMs, authentication etc. On Friday there was even a gaming session hosted by [Grant Ongers](https://www.linkedin.com/in/rewtd/) and [Max Alejandro Gómez-Sánchez Vergaray](https://www.linkedin.com/in/max-alejandro-gomez-sanchez-vergaray/) and I had a lot of fun! Luckily, the game is also [available online](https://copi.owasp.org), so there might be a chance for me to play it one day with my squad!

## The contact games

Embedded into the conference itself there were two games, encouraging attendees to interact with each other and the vendors/sponsors:

[Punk Security](https://punksecurity.co.uk/) built a little PCB shaped like a wasp with 19 LEDs on it and sent you on a scavenger hunt. Every board had one of these LEDs shining blue. When you met someone with a blue light at a different position, you could ask them to send you their light through a little sensor near the mouth of the wasp. Once you had collected 5, you got a battery extension pack for the wasp and with 10 you got a little extra board. Some exhibitors also offered add-ons for the little wasp!

The other game was a passport where you were supposed to go to each of the main sponsor's booths and get a stamp. Once full, you'd be entitled to enter a raffle at the end of the event. That didn't work for me at all. The wasp was a completely carefree game. We wandered around and just randomly bumped into people who were looking for one mote or another. The sponsor pass vendor passport felt boring in comparison. Since some exhibitors were already onboard with the LED-wasp, it might have IMO been a better idea to do away with the passport and just integrate everything through the wasp 😅

## Personal take-aways

- A box of OWASP Cornucopia 😁
- Some companies are now getting away from requiring code reviews for certain kinds of changes. I think that's pretty much expected especially around updates from trusted sources (after certain cooldown-windows).
- OWASP [ASVS](https://owasp.org/www-project-application-security-verification-standard/) and [AISVS](https://owasp.org/www-project-artificial-intelligence-security-verification-standard-aisvs-docs/) provide standardised checklists for security requirements.
- Agentic apps need gateways, sandboxes, and observability.
- CLAUDE.md files are already used by attackers for persistence!
- PyPI offers an API to trusted parties to immediate take packages down due to security concerns.
- While agents do not necessarily produce less secure code, they produce around 10x that of a normal developer which leads to an explosion in alerts.

Global AppSec EU will return to Vienna next year. I really enjoyed the community feeling here. The price for a full ticket (even without workshops) was quite high. Luckily, I could use my learning budget for that. Still, if I managed to not sleep through the early-bird phase next year, things will be easier 🙂
