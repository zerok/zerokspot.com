---
date: '2007-03-12T12:00:00-00:00'
language: en
tags:
- python
title: for-loop with index in Python
---


Everyday something new to learn :-) For quite some time now I'm coding mostly small scripts but also bigger stuff in Python (and now I'm also trying to get really into Django) and from time to time Python's for-each preference to the C-style for loops sometimes got a little bit in my way.

For example: I have a list and I want to iterate over it but also know where I am during this loop. So until today I would have gone this road:

<pre class="code python"><span class="n">data</span> <span class="o">=</span> <span class="p">(</span><span class="s">&#39;a&#39;</span><span class="p">,</span><span class="s">&#39;b&#39;</span><span class="p">,</span><span class="s">&#39;c&#39;</span><span class="p">)</span>
<span class="n">i</span> <span class="o">=</span> <span class="mi">0</span>
<span class="k">for</span> <span class="n">f</span> <span class="ow">in</span> <span class="n">data</span><span class="p">:</span>
	<span class="k">print</span> <span class="s">&quot;</span><span class="si">%d</span><span class="s"> =&gt; </span><span class="si">%s</span><span class="s">&quot;</span><span class="o">%</span><span class="p">(</span><span class="n">i</span><span class="p">,</span><span class="n">f</span><span class="p">)</span>
	<span class="n">i</span><span class="o">+=</span><span class="mi">1</span>
</pre>



-------------------------------



But today, while reading a litte bit in the Python Cookbook again, I discovered the `enumerate` object which make's my own counter completely redundant :-)

<pre class="code python"><span class="n">data</span> <span class="o">=</span> <span class="p">(</span><span class="s">&#39;a&#39;</span><span class="p">,</span><span class="s">&#39;b&#39;</span><span class="p">,</span><span class="s">&#39;c&#39;</span><span class="p">)</span>
<span class="k">for</span> <span class="n">i</span><span class="p">,</span><span class="n">f</span> <span class="ow">in</span> <span class="nb">enumerate</span><span class="p">(</span><span class="n">data</span><span class="p">):</span>
	<span class="k">print</span> <span class="s">&quot;</span><span class="si">%d</span><span class="s"> =&gt; </span><span class="si">%s</span><span class="s">&quot;</span><span class="o">%</span><span class="p">(</span><span class="n">i</span><span class="p">,</span><span class="n">f</span><span class="p">)</span>
</pre>
	
I love those small things that make my life sooo much easier :D
