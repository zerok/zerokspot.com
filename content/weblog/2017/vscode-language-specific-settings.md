---
date: 2017-04-11T21:31:02+02:00
title: Language-specific settings in VSCode
tags:
- editor
- vscode
- tooling
---

Over the last couple of days I've come to really
love [Visual Studio Code][vscode]. The amount of polish and tiny little gems in
there is just amazing! Something I stumbled upon this morning is that you can
override specific settings for specific file types.

Inside the editor this is called "Language specific settings" or "Language
based settings" and is documented in detail [here][docs]. It basically boils
down to something like this within your workspace's `settings.json` (or your
user-specific settings):

```
{
  "[markdown]": {
    "editor.wordWrap": "wordWrapColumn",
    "editor.wordWrapColumn": 80
  }
}
```

In this example, `[markdown]` as a configuration key indicates that everything
within that object should be specific to Markdown files. In this case I override
the word-wrap column by setting it to 80 characters and activate word-wrapping
at that column.

<figure>
<img src="/media/2017/vscode-language-settings.png" alt="" />
<figcaption><p>UI for accessing language-specific settings</p></figcaption>
</figure>

There is also quite a lot of UI available to generate that settings-block. You
can either click on the language button in the lower right corner when you have
a file open, or hit cmd+p and execute "Preferences: Configure Language Specific
Settings...". In general, I'm amazed by the way VSCode handles
preferences. Nearly everything has contextual help and auto-completion despite
being presented as a simple text file.

In [#1587][] there was even some discussion about letting you define settings
for specific filename patterns. That could be quite handy if you want to apply,
for instance, different word-wrap settings for Markdown files in different
folders for some reason That's not yet supported, though, nor could I find
another ticket about something like that. Luckily, I usually don't need
something like that 😉

[docs]: https://code.visualstudio.com/docs/getstarted/settings#_language-specific-editor-settings
[#1587]: https://github.com/Microsoft/vscode/issues/1587
[vscode]: https://code.visualstudio.com/
