---
title: Drafts and iOS smart dashes/punctuation
date: "2022-03-02T18:42:35+01:00"
tags:
- ios
- draftsapp
---

It's weird, I've been using [Drafts](https://getdrafts.com) for a long time to write pretty much all my notes on this blog (using some shortcuts to push the data then into WorkingCopy and up to my repository) but only just recently I started running into a weird issue.

I use a simple JavaScript action to parse a draft and add some YAML frontmatter to it, setting the title, tags, and publishing date. All of a sudden, all the quotes were in a locale specific format and the triple-dashes around the frontmatter were replaced with a single em-dash.

After some debugging and messing around I was pretty sure that the fault wasn't in my script and so I looked through the settings in Drafts to confirm that "smart quotes" and "smart dashes" were disabled. Turns out, they were. What *wasn't* disabled, though, was the "Smart punctuation" setting in the iOS keyboard settings. Disabling that brought the old behavior back.

I used following script as a minimal test-case:

```go
editor.setText('---\n\n\n---\n')
```

If I just set the text to `---` the three dashes will appear but once you have something after them, things get “smart”.

This is a bit of a weird one so here’s a list of all the conditions and their respective effect:

## Behavior for normal text input

1. System setting enabled, Drafts setting disabled: No smart punctuation 
2. System setting enabled, Drafts setting enabled: Smart punctuation
3. System setting disabled, Drafts setting disabled: No smart punctuation
4. System setting disabled, Drafts setting enabled:  No smart punctuation

## Behavior for editor.setText (API)

1. System setting enabled, Drafts setting disabled: Smart punctuation
2. System setting enabled, Drafts setting enabled: Smart punctuation
3. System setting disabled, Drafts setting disabled: No smart punctuation
4. System setting disabled, Drafts setting enabled:  No smart punctuation

As you can see, (1) behaves differently when text is entered using the script instead of just manually using the keyboard. Since my use-cases falls under this scenario, I have to disable smart punctuation system-wide for now. The perfect scenario for me would be if Drafts could somehow override the smart punctuation flag for its input fields but I have no idea if that’s possible. At the very least I’m going to ask in the [Drafts forums](https://forums.getdrafts.com) about this inconsistency 🙂
