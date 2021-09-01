---
title: "Massive COVID-19 test result leak in Tyrol"
date: "2021-09-01T17:33:00+0200"
tags:
- austria
- news
- covid19
- dataprivacy
---

The number of COVID-19 fails in Austria is getting longer and longer. This morning, news broke that more than 24,000 positive test results were leaked through an Excel sheet sent via e-mail. The information that I’m posting here is based on an article on [ORF.at](https://tirol.orf.at/stories/3119233/) and [DerStandard.at](https://www.derstandard.at/story/2000129310910/tausende-tiroler-pcr-testergebnisse-mit-namen-und-daten-geleakt).

Ralf Herwig, the former CEO of the company HG Lab Truck in Tyrol, had sent that file to an IT technician for “back up purposes”. HG Lab Truck was previously tasked by the state of Tyrol to manage PCR testing and the leaked dataset probably contains most if not all positive test results in the state from January to June.

There are just some many WTFs and open questions I don't really no where to start. Why does the CEO of a company have access to this kind of data? I guess because as a doctor he had some kind of double role in the company. But he resigned in May so at least after that he shouldn’t have had it.

Why did he sent that data via e-mail as Excel sheet to an IT technician who has been fired by the company that *previously* did the software for HG Lab Truck? He says that he did that as a back-up. Let’s ignore the sensitivity of the data because even for “normal” data sending something around via e-mail is not even remotely a good backup strategy. In this case we are talking about health data which has to be especially protected under the GDPR (and other regulations) so that idea just got a lot worse.

Judging by the media coverage my assumption is that the dataset wasn’t even encrypted, so some IT technician all of a sudden had direct access to 24,000+ test results.

Did he leak it? No idea. The former HG Lab Truck CEO stated that the leak probably happened due to a hacking attack:

> Herwig bestätigt auf Anfrage, dass er dies getan hat, und erklärt das Leck damit, dass er Opfer eines Hackerangriffs geworden sei. ([DerStandard.at](https://www.derstandard.at/story/2000129310910/tausende-tiroler-pcr-testergebnisse-mit-namen-und-daten-geleakt))

If you sent such stuff around via unencrypted e-mail to someone who shouldn’t have it while you yourself shouldn’t even have that data, there's just no need for a hacker to get involved. There is just so much stupidity in this whole story, I’ll take that as a more likely explanation.
