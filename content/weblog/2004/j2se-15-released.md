---
date: '2004-09-30T12:00:00-00:00'
language: en
tags:
- software
- java
title: J2SE 1.5 released
---


<p>Today Sun Microsystems released the new version of Java2. Ok, in the title here is written 1.5 but you'll also often find the name JDK 5.0. I'm not sure why they made this big version-jump but tbh I also don't really care ;-)</p>

-------------------------------



<p>Sun extended Java with some really nice things like (my personal favorite after looking at the changelog for about 5 minutes) <em>a foreach loop functionality</em>.

<ins><pre class="code">// Returns the sum of the elements of a

int sum(int[] a) {

    int result = 0;

    for (int i : a)

        result += i;

    return result;

}</pre></ins>

As you can see in this example (taken right from the changelog) you can now do things like <code>for(int i : a){...}</code> for iterating through an array. Same works for collections :-)</p>



<p>Java now also offers an implicit way to pass a variable number of arguments to a function/method without having to explicitly box them into a list or array or something like that. This can be achived with a declaration like this:

<ins><pre class="code">public void doSth(Object... arg){...}</pre></ins>

The three dots indicate that there is more behind it than just a simple single argument, but I haven't looked into this in detail yet :-) I just think it's nice to have it in Java.</p>



<p>For more details please read the <a href="http://java.sun.com/j2se/1.5.0/docs/relnotes/features.html">release notes</a> and or download the new J2SE from <a href="http://java.sun.com">java.sun.com</a></p>

