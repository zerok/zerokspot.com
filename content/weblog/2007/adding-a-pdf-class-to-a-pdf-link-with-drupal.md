---
date: '2007-02-05T12:00:00-00:00'
language: en
tags:
- accessibility
- drupal
- filter
- jquery
- module
title: Adding a pdf class to a PDF link with Drupal
---


Yesterday, during the Drupal presentation (Barcamp Kärnten 07) I think it was [Maria Putzhuber](http://www.putzhuber.net/) who asked a question regarding specifically styling links pointing to for example PDF documents by adding a specific class to these links.



-------------------------------



<img src="http://zerokspot.com/uploads/pdffilter-result.png" alt="Result: A 'PDF link' show now also have a 'pdf' class" class="figure"/>

## Option 1: jQuery

The first and probably least desirable but also fastest to implement solution would probably be to simply use jQuery and add something like this to a custom JS file or simply to the page.tpl.php of your theme:

<pre class="code javascript">
	<span class="nx">$</span><span class="p">(</span><span class="nb">document</span><span class="p">).</span><span class="nx">ready</span><span class="p">(</span><span class="k">function</span><span class="p">(){</span>
		<span class="nx">$</span><span class="p">(</span><span class="s1">&#39;//a&#39;</span><span class="p">).</span><span class="nx">each</span><span class="p">(</span><span class="k">function</span><span class="p">(){</span>
			<span class="k">var</span> <span class="nx">href</span> <span class="o">=</span> <span class="nx">$</span><span class="p">(</span><span class="k">this</span><span class="p">).</span><span class="nx">attr</span><span class="p">(</span><span class="s1">&#39;href&#39;</span><span class="p">);</span>
			<span class="k">if</span><span class="p">(</span><span class="nx">href</span><span class="p">){</span>
				<span class="nx">href</span> <span class="o">=</span> <span class="nx">href</span><span class="p">.</span><span class="nx">toLowerCase</span><span class="p">();</span>
				<span class="k">var</span> <span class="nx">ext</span> <span class="o">=</span> <span class="nx">href</span><span class="p">.</span><span class="nx">split</span><span class="p">(</span><span class="s1">&#39;.&#39;</span><span class="p">);</span>
				<span class="nx">ext</span> <span class="o">=</span> <span class="nx">ext</span><span class="o">[</span><span class="nx">ext</span><span class="p">.</span><span class="nx">length</span><span class="err">-</span><span class="m">1</span><span class="p">];</span>
				<span class="k">if</span> <span class="p">(</span><span class="nx">ext</span><span class="o">==</span><span class="s1">&#39;pdf&#39;</span><span class="p">){</span>
					<span class="nx">$</span><span class="p">(</span><span class="k">this</span><span class="p">).</span><span class="nx">addClass</span><span class="p">(</span><span class="s1">&#39;pdf&#39;</span><span class="p">);</span>
				<span class="p">}</span>
			<span class="p">}</span>
		<span class="p">});</span>
	<span class="p">});</span>
</pre>

Depending on your theme you might also have to explicitly load jQuery into the page.tpl.php, but as this is the worst option there is for doing something like this, I won't really elaborate on it. If you still want to know how to get this working, simply leave a comment :-)


## Option 2: Filtering

In my opinion the best solution for this problem would be to use another of Drupal's features: Content filters. Drupal uses a number of so called "filters" to parse the input of a user and transform it to whatever the admin/user wants. These filters for example make the use of Markdown or Textile as alternatives to HTML as input language possible.

These filters are combined into "input formats" which allows also the explicit stacking of input filters in a specific order so that for example a filter parsing smilies should be run after for example the "Markdown to HTML" filter.

<img src="http://zerokspot.com/uploads/drupal-filterweights.png" alt="Rearranging Filters in Drupal using weights" class="figure"/>

So the basic idea here to have a regular expression that matches links with endpoints ending with ".php" and add to these links the class "pdf". The before mentioned weighting of filters is important here since our new filter should really only be executed after the user input has already been transformed into HTML.

To demonstrate this I made a small module having only about 50 lines (including meta data etc. ;-) ) which you can find attached to this post. __Note:__ Use it at your own risk. While it shouldn't do anything bad, you never know ;-)

To try this one out simply install the module as you'd do with any other module (extract it inside a contrib folder of the modules directory and enable it on the modules admin page), then go to the input format settings and add it for example to the "Filtered HTML" format. And don't forget to give it the highest weight of all the included filters :-)

## Other Options

There is yet another option that uses yet another nice feature of Drupal's templating system: You can manipulate every rendering aspect of the page from within the template using the template.php. This way you could use the same regex that was included in the pdffilter module to manipulate the content of the node _after_ it went through all the filters. 

I'm not sure, but I think this would also get around the whole "non-aggressive" caching system of Drupal which is why this is probably even more hacky than the first option using jQuery. 

I hope this helps :-)