---
date: '2006-04-17T12:00:00-00:00'
language: en
tags:
- structuredblogging
- wordpress
title: Giving StructuredBlogging for WordPress a try
---


Another entry into the "Giving (.*?) a try" series ;) This time's topic is the StructuredBlogging plugin for WordPress. 

[StructuredBlogging](http://www.structuredblogging.org) is an initiative bringing to you means to blog in a more semantic way giving automated aggregators some new ways to categorize your content. For example if you want to review something, you write a review that can now also be recognized as such by automated aggregators.



-------------------------------



To achieve this, StructuredBlogging - in addition to using the hReview microformat -  appends basically an XML chunk to each of your posts including the metadata instead of going with the Microformats idea of using plain and simply HTML in combination with classes to get the semantics into the posts. This way it's definitely easily parseable but it also adds quite a few bytes to your post's output compared to the Microformats approach. For more information about this you might want to checkout the [developer ressources page](http://www.structuredblogging.org/resources.php).

A short note about the installation though: I'm trying this on my Mac, so the installation instructions don't completely work since `ß¬nstall -D` doesn't exists on Darwin.

<pre class="code">
for f in `find . -type f`; do install -D $f /path/to/wordpress/installation/$f; done
</pre>... doesn't work

But this should basically do the same: <pre class="code">
for f in `find . -type f`; do install $f /path/to/wordpress/installation/$f; done
</pre>

So my first test is writing a small bookreview. For things like that the plugin offers fields like "Book title", "Category", "Author", etc. All these forms are generated as an overlay to the normal posting form which looks quite nice this way. But already in this form I see the first problems: 


* Why doesn't it use the normal categories used by WordPress? 
* What if the book has multiple authors? 
* When you want to store a posting as a draft, you only can do this once. When you edit it later on, all you get is a save button, which will also public the post no matter what.
* You can kill their parser by simply adding &lt;!--more--&gt; into the description field ;)

Another problem at least for me is, that I can't use Markdown with this plugin. At least the points above are merely bugs (perhaps except for the first one) and the last problem could probably be fixed by moving the rendering into the templates and away from the plugin. But apart from this, the whole idea in my opinion looks really nice, and if it's adopted by the masses, it could give the whole blogging thing another push into a more structured direction, where you could simply tell Technorati, that you want to read a review of something and not get dozens of advertizing posts without anything reviewy about them :)

But to bring this post back to the StructuredBlogging plugin: I doubt I'll use it for this weblog here or my [booksblog](http://booksblog.zerokspot.com) in the near future, but esp. for the booksblog adding the [hReview microformat](http://microformats.org/wiki/hreview) would be at least an idea since it's all about semantics, isn't it? ;) 

And in this segment currently the plugin seems to fail a little bit (but given their different approach this is probably not really a problem although a pain for semantics evangelists ;)):

<pre class="code">&lt;p&gt;&lt;b&gt;Author&lt;/b&gt;:
	 Dave Stern&lt;/p&gt;
</pre>

Am I the only one missing the semantics here? Perhaps something like that would make more sense:

<pre class="code">&lt;p&gt;&lt;span class=&quot;label&quot;&gt;Author(s):&lt;/span&gt;
	&lt;span class=&quot;authors&quot;&gt;Dave Stern&lt;/span&gt;&lt;/p&gt;</pre>

And achieving something like that esp. with the quite powerful template systems of TXP and WP shouldn't be that much a problem :)

Just my 2 eurocents :)

**Note:** My opinion of the StructuredBlogging WordPress plugin is based on 90 minutes playing with it and only trying the book review so far (this timespan also include the writing of this post) ;)