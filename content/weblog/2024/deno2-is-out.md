---
date: "2024-10-10T17:32:56+02:00"
tags:
- javascript
- typescript
- deno
- opensource
title: Deno 2 is out!
---

Yesterday the [Deno team announced version 2](https://deno.com/blog/v2.0) of their JavaScript/TypeScript runtime with an awesome [announcement video and a live-stream](https://www.youtube.com/watch?v=d35SlRgVxT8) for a community Q&A. First of all, the opening of the video announcement was just hilarious going for something that's somewhere between a Kickstarter demo and an Apple press-event. Just hilarious 😆

As for the new features in Deno 2, there are quite a few with compatibility for npm packages and NodeJS libraries being probably the most prominent here, making the deno binary also a package manager with `add`, `remove`,  `install` and `uninstall` commands. Inside the source files, npm, NodeJS and other libraries have their own prefix to keep everything transparent:

```
# Using a NodeJS core library:
import { readdir } from 'node:fs/promises';
await readdir('.');

# Using a package from NPM:
import * as chalk from 'npm:chalk'

# Using a package from JSR:
import { toCamelCase } from 'jsr:@std/text@1'
```

The [Deno standard library](https://jsr.io/@std) has now also reached its v1 milestone. It's all on JSR so these libraries are just yet another package that you import. This was also the first time that I visited [jsr.io](https://jsr.io) for more than a few seconds an I really like it with the integrated documentation and quality score. Most of that was probably there before the Deno v2 release but it's still really nice!

Another really interesting feature is that Deno now ships with a Jupyter Kernel so that you can run it directly from your Jupyter Notebooks! [Simon Willison](https://simonwillison.net/2024/Oct/10/announcing-deno-2/) already gave that a slightly deeper look on his blog. Combined with the npm support I can definitely see myself using that in the future!

I immediately went out to change the runtime for one of my Astro playground to use Deno and it worked right away! Exciting times!