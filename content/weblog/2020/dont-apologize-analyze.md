---
title: "Don't apologize, analyze"
date: "2020-05-07T08:31:02+0200"
tags:
- development
- 100daystooffload
- workflow
inReplyTo: "https://blog.danslimmon.com/2019/08/02/stop-apologizing-for-bugs/"
---

Quite some time ago, Dan Slimmon had a great post titled ["Stop apologizing for bugs"](https://blog.danslimmon.com/2019/08/02/stop-apologizing-for-bugs/
) in which he described the possible side-effects of developers apologizing bugs they've produced to other team members. While it might seem appropriate at first, keep those points in mind:

> - It reinforces the idea that any one person or piece of code can be blamed for a given failure. Short of malice, this is never the case.
> - It gives the impression that, when you wrote the code, you should have written it better. This is a counterfactual that rarely holds up to examination.
> - It positions shame as the correct emotion to feel about bugs in your code: if you were a better engineer – a better teammate – the bug wouldn’t exist.
> - If you’re a more senior engineer on your team, the effects of these anti-patterns are magnified: people see you apologizing for bugs, so they think that they should be striving to write bug-free code. They may feel ashamed if their code has bugs.

I completely agree with every single item in that list but double so with the last. Whenever a coworker found a bug in code that I had written (even if it's not critical), it felt ashamed to the point that I got angry with myself. In the past this has sometimes caused me to panic and therefore not rationally prioritise and resolve the issue but to resolve the issue no matter what, ignoring any kind of priority. In recent years this has become better but the temptation to simply panic is still there.

Trying to get past this behaviour, the following has worked well for me so far: Whenever a bug shows its head, I write down details about it (ideally on paper), forcing me to analyze what is going on and getting past the first instinctive temptations. A big part of that analysis is to find out how the issue can be reproduced and not what is causing it. The first is fact-based while the latter is speculative.

This is important as it helps me to force my brain down a certain path instead of giving it free reign. It's immensely important for me to avoid panic-mode (which I still fall into from time to time) and analysis instead of speculation helps. Additionally, writing stuff down on paper has a calming effect on me, which is always a bonus. 

