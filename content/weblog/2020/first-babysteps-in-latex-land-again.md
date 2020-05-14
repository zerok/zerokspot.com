---
title: "First baby steps in LaTeX-land (again)"
date: "2020-05-08T07:44:51+0200"
tags:
- latex
- tex
- writing
- 100daystooffload
---

I haven't used [LaTeX](https://www.latex-project.org/) since I handed in my master thesis more than 10 years ago. Now that I'm attending some university classes again, though, I thought it would be fitting to also start writing homework assignments that way again. Obviously, the basics haven't changed all that much but there are still a couple of things that I've learnt over the course of writing the first couple of texts again. My requirements for LaTeX also changed:

- I need to write multiple small documents in parallel instead of one huge.
- All documents should use a single bibliography.
- All documents should have a similar style leaning on the classic article document class.

## The setup

Since I hadn't used LaTeX for many years, I didn't have it installed on my laptop. So the first thing I did was to install [MacTeX](http://www.tug.org/mactex/) using Homebrew:

```
$ brew install mactex
$ export PATH=/usr/local/texlive/2020/bin/x86_64-darwin:$PATH
```

There is also a non-GUI version of that formula but I went with the default one for now. VIM supports TeX documents out of the box so I didn't have to install any additional packages for it.

For every class I'm attending (and for which I have to write essays et al.) I created a folder which contains a central bibliography file and a Makefile. More on both later.

## biblatex/biber has replaced bibtex

One of the main reasons why I enjoy working with LaTeX is the way citations are handled. The system that is used for that, bibtex, is ancient but works quite well. Sadly, getting things like German umlauts working in it is a bit tedious, though. Custom styling is also not something you just do because you feel like it.

[biblatex using biber](https://ctan.org/pkg/biblatex?lang=en) as backend seems to be the new preferred approach for dealing with both. `.bib` files can contain UTF-8 and styling happens through normal LaTeX commands.

## Bibliography as footnotes

For books and longer reports it's definitely nice to have all your sources listed in a dedicated section, especially you usually use a source more than once. For articles, though, putting them into footnotes is often enough. Turns out, you can do that using `\footcite` instead of `\cite` 😅

## Single-column IEEEtran

In general I really like the [IEEEtran](https://ctan.org/pkg/ieeetran?lang=en) document class but for the classes I required something that supports generating single-column documents. Turns out that IEEEtran actually allows that using the `onecolumn` option 🙂

## Custom documentclass

Since I have to write on different documents at the same time I also finally moved all the settings I apply for a specific class into a custom documentclass based on IEEEtran:

```
\NeedsTeXFormat{LaTeX2e}
\ProvidesClass{classnamearticle}
\LoadClass[12pt,onecolumn,a4paper]{IEEEtran}
\usepackage{url}
\usepackage[style=verbose]{biblatex}
\usepackage{csquotes}
\usepackage[ngerman]{babel}
\usepackage{titling}
\usepackage{setspace}
\usepackage{fullpage}
\setstretch{1.3}
\setlength{\droptitle}{-6em}
\setlength{\parskip}{0.3em}
\renewcommand\maketitlehookc{\vspace{-1.5em}}
\usepackage[margin=1in]{geometry}
\renewcommand{\UrlFont}{\small\tt}
```


## Makefile per class

[Makefiles](https://www.gnu.org/software/make/) are for me the central interface I expect when entering a folder that contains files that need to be processed. Since `.tex` files fall in this category, I created a Makefile for each class. There's actually nothing fancy in there but it just builds all the tex documents in the current directory including bibliography:

```
docs := $(shell find . -name '*.tex')
bases := $(shell echo $(docs) | sed s/.tex//g)
pdfs := $(shell echo $(docs) | sed s/.tex/.pdf/g)

all: $(pdfs)

%.pdf: %.tex classname.bib classnamearticle.cls
	latex $(subst .tex,,$<)
	biber $(subst .tex,,$<)
	latex $(subst .tex,,$<)
	pdflatex $<

clean:
	rm -f $(pdfs) *.aux *.dvi *.bbl *.blg *.bcf *.log *.run.xml

.PHONY: clean
.PHONY: all
``` 

That's it for now but every time I work with LaTeX I learn something new 😄
