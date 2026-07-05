---
title: Additional checkbox-types in Obsidian
tags:
- obsidian
- bulletjournaling
date: "2026-07-05T20:29:00+02:00"
---

Thanks to the current heatwave I'm experiment with moving parts of my bullet journal temporarily away from paper and into Obsidian. While looking for ways how others handle the daily migrations when using checkboxes, I stumbled upon this [video by Rob Conery](https://www.youtube.com/watch?v=k66W_wOQKVc) where he describes his setup and mentions that various themes support additional checkbox styles! Luckily, the [Minimal](https://minimal.guide/checklists) theme is one of these.

But what if your favourite theme does not support them or you have an idea for another type? So I was curious how this was done under the hood. Obsidian is an Electron app, so all the rendering HTML + CSS in the end *and* the app not only exposes the usual `:checked` state of a checkbox but also what is inside the `- [ ]` braces. If you put a `>` in there, the input element has a `data-task=">"` attribute. And all of a sudden, things like this should be possible as the Minimal theme demonstrates in its [checklist-icons.scss](https://github.com/kepano/obsidian-minimal/blob/master/src/scss/features/checklist-icons.scss):

```css
input.task-list-item-checkbox[data-task='>']:checked,
    border: none;
    background: none;
    --checkbox-marker-color: transparent;
    background-image: url("data:image/svg+xml,$ENTITY_ENCODED_SVG");
}
```

What's needed now is a way to add custom CSS to a Vault independent of the used theme. Luckily, Obsidian supports this through [CSS Snippets](https://obsidian.md/help/snippets). With this in place, I can do bullet journaling with Obsidian in whatever theme I choose. There are a couple (potentially minor) downsides to this whole setup, though:

- It's Obsidian-specific since none of this is in any Markdown or Commonmark specification. On the plus-side, extending, for instance, [goldmark](https://github.com/yuin/goldmark) shouldn't be that complicated if I wanted to create any custom tooling around this.
- Support in the mobile app is similar to the one in the desktop app: You have to manually enter the respective sign and moving there on a virtual keyboard is tedious.

