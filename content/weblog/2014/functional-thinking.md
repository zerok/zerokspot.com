---
date: '2014-08-02T22:06:38+02:00'
language: en
rating: 2
tags:
- books
- review
- development
- java
title: Functional Thinking by Neal Ford; O'Reilly Media
---

<img src="/media/2014/thumb.gif" alt="" style="float:left; margin: 0 1em 1em 0"/>
In this 180 pages book Neal Ford tries to introduce the reader to functional
programming and the core concepts that distinguish it from other programming
language paragigms. He does this by first explaining why you as a developer
should care about functional methodologies, then moves on to introducing
the core functional building blocks that can be found in languages like
Scala, Java 8, and Groovy with a focus on filtering, mapping and reducing.
After that the book is all about actual language constructs like closures,
memoization, and lazyness (as in lazy evaluation of expressions) that are
rather hard to implement in non-functional environments without making them
appear clumsy.

**Disclaimer:** I got this book as a review-copy from O'Reilly as part of their
reviewer programme.

In general, I think the author's primary goal was to help object oriented
developers working with the JVM to take a step outside their comfort zone and
explore functional programming without leaving the JVM completely. Introducing
Java developers to mapping, filtering, and reducing now that Java 8 finally has
some less verbose ways to handle that is definitely a good idea, but the book
kind of just does some cherry picking here.

For example, you have to wait until page 137 for this sentence:

<blockquote>Functional architectures embrace immutability at the core level, leveraging it as much as possible. Embracing immutability is high on the list of ways to think like a functional programmer.</blockquote>

... which should perhaps have been in the first chapter. The whole idea of
a functional actually being side-effect-free is also only hinted upon
throughout some chapters but that's it. Same goes for the term "higher-order
function" which is used throughout the book without but only actually described
on page 122.

The second big problem I have with this book, though, are the examples.
The author tries to come up with simple examples but then always repeats them
for Scala and Groovy and if possible Java. At first this seemed like a good
approach but already after the second chapter it felt rather tedious and
distracting. Scala and Closure are just such different beasts from Java that
(despite the author trying very hard to help here) you definitely benefit from
knowing at least the basics of each language for some of the examples. Sticking
with the JVM throughout the book might have not been the best choice after all.

Chapter 2 which introduces map, reduce, and filter also suffers from quite
a lot of repetition. Basically these basic functions are described multiple
times throughout the chapter which really harms the reading flow.

In the last two chapters the author tries to look at functional methodologies
again from a rather high point of view (architecture, web applications, ...),
which was nice. But again, in a look at the different type systems in multiple
languages only a very small selection was made that only included JVM and .NET
languages and JavaScript, despite Rake (a build-tool written in Ruby) being
mentioned on the same page in a different context (p. 154).

Book's website on [oreilly.com](http://shop.oreilly.com/product/0636920029687.do)

-------------

While I think the book had good intentions for me it is too focused on the JVM
environment. The last chapter indicated that the author also had experience with
other things outside the JVM and .NET world so it would have been nice to see
some of this here. That, the unfortunate order of introducing functional
characteristics and the length of the mostly repeating examples that distract
from the actual functional concepts while focusing too much on the
implementation make this really hard to recommend.
