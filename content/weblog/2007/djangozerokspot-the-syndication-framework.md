---
date: '2007-03-30T12:00:00-00:00'
language: en
tags:
- django
- syndication
- zerokspot
title: 'django@zerokspot: The syndication framework'
---


Welcome to the 2nd installment of my little series where I write about what I learnt today when trying to implement zerokspot.com using Django. This time it will be all about Django's contributed syndication framework



-------------------------------



First of all: What is it? django.contrib.syndication is a framework consisting of a single view as well as a couple of helper classes and functions that allow developers to easily extend their websites with RSS and Atom feeds.

So what you'd normally have to do is just write an URL pattern pointing to the framework's `feed` view and pass to it a Feed class that actually fetches the data you want to have in your feed. For detail on the general usage of this module take a look at the [official documentation](http://www.djangoproject.com/documentation/syndication_feeds/) :)

Coming from Drupal I now want to have a remotely similar URL scheme for feeds. Something like /weblog/tags/lala/feed/atom/ would for example return a feed of the latest weblog entries tagged with "lala".

At first I thought, Django couldn't do it and would require some format like /weblog/feed/tags/lala/atom, basically with the relevant part of the content selection after an URL separator like "feed" or "feeds" (or whatever you prefer). So I thought about it a little bit and started writing a wrapper for the syndication.view `feed` in order to pass additional arguments like the tag and the feed type ("atom" in this case).

The URL configuration for this looked now somewhat like this:

<pre class="code python"><span class="n">urlpatterns</span> <span class="o">=</span> <span class="n">patterns</span><span class="p">(</span><span class="s">&#39;&#39;</span><span class="p">,</span>
<span class="p">(</span><span class="s">r&#39;^tags/(?P&lt;tag&gt;.+)/feeds/(?P&lt;url&gt;.*)/$&#39;</span><span class="p">,</span>
	<span class="s">&#39;zerokspot.weblog.views.tag_feed&#39;</span><span class="p">),</span>
<span class="p">)</span>
</pre>

with the actual view now looking this way:

<pre class="code python"><span class="k">def</span> <span class="nf">tag_feed</span><span class="p">(</span><span class="n">request</span><span class="p">,</span><span class="n">url</span><span class="p">,</span> <span class="n">tag</span><span class="p">,</span> <span class="n">feed_dict</span><span class="o">=</span><span class="p">{}):</span>
	<span class="sd">&quot;&quot;&quot;</span>
<span class="sd">	Wraps django&#39;s feed generating view in order to pass additional stuff</span>
<span class="sd">	to the feed class.</span>
<span class="sd">	&quot;&quot;&quot;</span>
	<span class="c"># Prepare classes for the given tag and add them to the feed_dict</span>
	<span class="n">rss</span> <span class="o">=</span> <span class="n">feeds</span><span class="o">.</span><span class="n">createTagFeed</span><span class="p">(</span><span class="n">tag</span><span class="p">,</span><span class="s">&#39;rss&#39;</span><span class="p">)</span>
	<span class="n">atom</span> <span class="o">=</span> <span class="n">feeds</span><span class="o">.</span><span class="n">createTagFeed</span><span class="p">(</span><span class="n">tag</span><span class="p">,</span><span class="s">&#39;atom&#39;</span><span class="p">)</span>
	<span class="n">feed_dict</span><span class="p">[</span><span class="s">&#39;rss&#39;</span><span class="p">]</span><span class="o">=</span><span class="n">rss</span>
	<span class="n">feed_dict</span><span class="p">[</span><span class="s">&#39;atom&#39;</span><span class="p">]</span><span class="o">=</span><span class="n">atom</span>
	<span class="k">return</span> <span class="n">django</span><span class="o">.</span><span class="n">contrib</span><span class="o">.</span><span class="n">syndication</span><span class="o">.</span><span class="n">views</span><span class="o">.</span><span class="n">feed</span><span class="p">(</span><span class="n">request</span><span class="p">,</span><span class="n">url</span><span class="p">,</span><span class="n">feed_dict</span><span class="p">)</span>
</pre>

As you can see here, I'm using a custom factory method that creates a Feed class for every tag and ('rss','atom',) combination out there on the fly. 

<pre class="code python"><span class="k">def</span> <span class="nf">createTagFeed</span><span class="p">(</span><span class="n">_tag</span><span class="p">,</span><span class="n">_type</span><span class="p">):</span>
	<span class="c"># Extract the tag from the request_path</span>
	<span class="n">_types</span> <span class="o">=</span> <span class="p">{</span><span class="s">&#39;rss&#39;</span><span class="p">:</span><span class="n">Rss201rev2Feed</span><span class="p">,</span><span class="s">&#39;atom&#39;</span><span class="p">:</span><span class="n">Atom1Feed</span><span class="p">}</span>
	<span class="n">tag</span> <span class="o">=</span> <span class="n">get_object_or_404</span><span class="p">(</span><span class="n">Tag</span><span class="p">,</span><span class="n">name</span><span class="o">=</span><span class="n">_tag</span><span class="p">)</span>

	<span class="k">class</span> <span class="nc">klass</span><span class="p">(</span><span class="n">Feed</span><span class="p">):</span>
		<span class="n">title</span> <span class="o">=</span> <span class="s">&quot;tag:</span><span class="si">%s</span><span class="s">&quot;</span><span class="o">%</span><span class="p">(</span><span class="n">tag</span><span class="o">.</span><span class="n">name</span><span class="p">,)</span>
		<span class="n">link</span> <span class="o">=</span> <span class="s">&quot;/weblog/tags/</span><span class="si">%s</span><span class="s">/&quot;</span><span class="o">%</span><span class="p">(</span><span class="n">tag</span><span class="o">.</span><span class="n">name</span><span class="p">,)</span>
		<span class="n">description_template</span> <span class="o">=</span> <span class="s">&#39;feeds/latest_description.html&#39;</span>
		<span class="n">feed_type</span> <span class="o">=</span> <span class="n">_types</span><span class="p">[</span><span class="n">_type</span><span class="p">]</span>
		<span class="k">def</span> <span class="nf">items</span><span class="p">(</span><span class="bp">self</span><span class="p">):</span>
			<span class="k">return</span> <span class="n">TaggedItem</span><span class="o">.</span><span class="n">objects</span><span class="o">.</span><span class="n">get_by_model</span><span class="p">(</span><span class="n">Entry</span><span class="p">,</span><span class="n">tag</span><span class="p">)</span><span class="o">.</span><span class="n">order_by</span><span class="p">(</span><span class="s">&#39;-datetime&#39;</span><span class="p">)[:</span><span class="mi">10</span><span class="p">]</span>
	<span class="k">return</span> <span class="n">klass</span>
</pre>

While perhaps a funny idea, it's ... strange and quite obscure. Esp. since I consider stuff like that a quite common use-case for feeds and therefor couldn't really believe than some quite flexible URL handling wasn't somewhere in the syndication framework. 

So I kept looking and eventually found it (I have to admit that I was blind when I looked for the first time and missed it ;-))

Then I went looking a little bit more and found a hint in the [cab-project](http://code.google.com/p/cab/)'s source code that indicated some get\_object method in order to handle more complicated URLs. And thanks to some input from jdunck_ on #django I finally found my way around that Feed class ;)

I guess the easiest way now to achieve feeds behind such an URL-scheme would be something like this:

<pre class="code python"><span class="n">urlpatterns</span> <span class="o">=</span> <span class="n">patterns</span><span class="p">(</span><span class="s">&#39;&#39;</span><span class="p">,</span>
<span class="p">(</span><span class="s">r&#39;(?P&lt;url&gt;tags/[^/]+)/feed/&#39;</span><span class="p">,</span><span class="n">feed</span><span class="p">,</span>
	<span class="p">{</span><span class="s">&#39;feed_dict&#39;</span><span class="p">:</span> <span class="p">{</span><span class="s">&#39;tags&#39;</span><span class="p">:</span><span class="n">TagFeed</span><span class="p">}}),</span>
<span class="p">)</span>
</pre>

This does no longer require a custom view (the `feed` function referenced here _is_ the one offered by the syndication framework) but just need the TagFeed class:

<pre class="code python">	<span class="k">class</span> <span class="nc">TagFeed</span><span class="p">(</span><span class="n">Feed</span><span class="p">):</span>
		<span class="n">description_template</span> <span class="o">=</span> <span class="s">&#39;feeds/latest_description.html&#39;</span>
		<span class="n">feed_type</span> <span class="o">=</span> <span class="n">Atom1Feed</span>
		<span class="k">def</span> <span class="nf">get_object</span><span class="p">(</span><span class="bp">self</span><span class="p">,</span><span class="n">bits</span><span class="p">):</span>
			<span class="n">tagname</span> <span class="o">=</span> <span class="n">bits</span><span class="p">[</span><span class="mi">0</span><span class="p">]</span>
			<span class="n">tag</span> <span class="o">=</span> <span class="n">get_object_or_404</span><span class="p">(</span><span class="n">Tag</span><span class="p">,</span><span class="n">name</span><span class="o">=</span><span class="n">tagname</span><span class="p">)</span>
			<span class="k">return</span> <span class="n">tag</span>
		<span class="k">def</span> <span class="nf">title</span><span class="p">(</span><span class="bp">self</span><span class="p">,</span><span class="n">obj</span><span class="p">):</span>
			<span class="k">return</span> <span class="s">&quot;tag:</span><span class="si">%s</span><span class="s">&quot;</span><span class="o">%</span><span class="p">(</span><span class="n">obj</span><span class="o">.</span><span class="n">name</span><span class="p">,)</span>
		<span class="k">def</span> <span class="nf">link</span><span class="p">(</span><span class="bp">self</span><span class="p">,</span><span class="n">obj</span><span class="p">):</span>
			<span class="k">return</span> <span class="s">&quot;/weblog/tags/</span><span class="si">%s</span><span class="s">/&quot;</span><span class="o">%</span><span class="p">(</span><span class="n">obj</span><span class="o">.</span><span class="n">name</span><span class="p">,)</span>
		<span class="k">def</span> <span class="nf">items</span><span class="p">(</span><span class="bp">self</span><span class="p">,</span><span class="n">obj</span><span class="p">):</span>
			<span class="k">return</span> <span class="n">TaggedItem</span><span class="o">.</span><span class="n">objects</span><span class="o">.</span><span class="n">get_by_model</span><span class="p">(</span><span class="n">Entry</span><span class="p">,</span><span class="n">obj</span><span class="p">)[:</span><span class="mi">10</span><span class="p">]</span>
</pre>

Much better, isn't it? ;-) This one just offers Atom1.0 feeds for now, but you should get the idea. Thanks jdunck_ for saving me from a breakdown :)
