---
title: Excluding files from completion in VIM
date: "2020-06-17T12:18:06+02:00"
tags:
- vim
- usability
---

Ever since working with LaTeX again for university, I’ve had one pain point in my setup: The autocompletion for opening files listed also all those files that I normally don’t want to open in VIM anyway like PDFs and `.aux`-files.

Today I looked a bit around and learnt a way to hide those:

	set wildignore+=*.pdf,*.aux,*.run.xml,*.blg,*.bcf,*.bbl

[Stack Exchange](https://vi.stackexchange.com/questions/2426/how-do-i-exclude-certain-file-extensions-from-completion#2427) for the win 😅
