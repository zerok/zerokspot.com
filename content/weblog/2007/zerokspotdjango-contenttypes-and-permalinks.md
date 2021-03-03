---
date: '2007-03-29T12:00:00-00:00'
language: en
tags:
- django
- zerokspot
title: 'zerokspot+Django: contenttypes and permalinks'
---


As some kind of project for the next couple of months I've now started to work on a new implementation of zerokspot.com no longer using a pre-built CMS like WordPress or Drupal but trying to do it completely on my own using [Django](http://www.djangoproject.com) and several generic apps for it.



-------------------------------



While I'm not sure yet, if this system will really succeed the current Drupal site, it will at least do a couple of things for me:

1. Improve my knowledge about Django and how to work with it
2. Perhaps some style improvements
3. Offer me a general playground for ideas

I also thought for a while to go even deeper and just start from scratch using components like [CherryPy](http://cherrypy.org), [SQLObject](http://www.sqlobject.org/) and [Mako](http://www.makotemplates.org/), but I guess this would have been a little bit too hardcore or simply would have taken too much time. I simply prefer to always see a light at the end of the tunnel ;)

So last Wednesday I started writing some code and also some general features I'd like to see in this new implementation like Tagging, RSS/Atom etc. and in this and the following entries I'd like to write a little bit about my progress (whatever this will be worth in the end) and new things I learnt about Django.

## Content types

The first things I learned about Django came from taking a look at the [Django-tagging](http://code.google.com/p/django-tagging)'s code: There is a generic way to store model associations for a certain object. Think about following situation: You want to offer a commenting system with your project but don't want to limit it to , let's say,  the Article model. 

This is where the contenttypes app comes in handy. It contains a Model (and therefor a table in your project's database) that manages all the Models used by your project and let's you reference them using the content type as foreign keys.

<pre class="code python"><span class="k">class</span> <span class="nc">Comment</span><span class="p">(</span><span class="n">models</span><span class="o">.</span><span class="n">Model</span><span class="p">):</span>
	<span class="n">body</span> <span class="o">=</span> <span class="n">models</span><span class="o">.</span><span class="n">TextField</span><span class="p">()</span>
	<span class="n">content_type</span> <span class="o">=</span> <span class="n">models</span><span class="o">.</span><span class="n">ForeignKey</span><span class="p">(</span><span class="n">ContentType</span><span class="p">)</span>
</pre>

ContentType's manager now offers with the get_for_model now a way to receive the respective foreign keys for a given Model :-)

Something that I haven't found out yet is, how I could use this, to efficiently have a listing of - for example - the latest 10 content elements created on my site (like reviews, entries etc.). So far I would go about it somehow like this:

1. Extend for example the journal entry model's `safe` method with code to automatically create an entry into a global content table. Let's call this model Node for now.
2. Then I'd fetch the latest n elements from this Node class using the normal `Node.objects.order_by('-pub\_date')[:5]`
3. Now comes the real problem: How to get to the actual elements without creating tons of queries. I guess the only solution for this for now would be to create temp. lists of IDs for each content type and then do queries on them.

I think I will implement something like this tomorrow just to find out if and how it really works. Might be a nice addition to django.contrib.contenttype's Manager if it doesn't exist yet and I just missed it :-)

## Permalinks

Another thing I learnt yesterday was, that there also exists a decorator to facilitate the creation of Model specific URLs. One of the first things I learnt about Django was, that there seems to be a convention that model's should offer a `get_absolute_url` method in order to explicitly address model instances with an URL. 

While a great idea, it becomes a pain to maintain and it would be nice to be able to bind such an object not to an URL per se, but to a method that would be called in order to represent such an object to the user at the other end of the pipe.

This would now be a job for the `permalink` decorator:

<pre class="code python"><span class="k">class</span> <span class="nc">Entry</span><span class="p">(</span><span class="n">models</span><span class="o">.</span><span class="n">Model</span><span class="p">):</span>
	<span class="c"># ...</span>
	
	<span class="nd">@models</span><span class="o">.</span><span class="n">permalink</span>
	<span class="k">def</span> <span class="nf">get_absolute_url</span><span class="p">(</span><span class="bp">self</span><span class="p">):</span>
		<span class="k">return</span> <span class="p">(</span><span class="s">&#39;myproj.myapp.views.view_entry&#39;</span><span class="p">,[],</span><span class="nb">dict</span><span class="p">(</span><span class="nb">id</span><span class="o">=</span><span class="bp">self</span><span class="o">.</span><span class="n">id</span><span class="p">),)</span>
</pre>

This will now do a reverse lookup in the URLs configuration and generate an URL for the given view function.
