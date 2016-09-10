---
date: '2006-11-24T12:00:00-00:00'
language: en
tags:
- 5-0
- d5
- drupal
- theming
title: Theming primary links in Drupal 5.0
---


I'm currently in the process of porting over the theme here to Drupal 5.0 and noticed a slight difference between 4.7 and 5.0. While I was able to use something like this in my template.php before:

<pre class="code php">
function mytheme_menu_item_link($item, $link_item) {
	static $menu;
	$attribs = isset($item[&apos;description&apos;]) ? 
		array(&apos;title&apos; =&gt; $item[&apos;description&apos;]) : array();
	if (&apos;&lt;front&gt;&apos; == $link_item[&apos;path&apos;] &amp;&amp;
		(
			arg(0)==&apos;node&apos; &amp;&amp; (
				(is_numeric(arg(1)) &amp;&amp; arg(1) != 213)
				|| arg(1)==&apos;&apos;
			)
		)
	){
		$attribs[&apos;class&apos;] = &apos;active&apos;;
	}
	return l($item[&apos;title&apos;], $link_item[&apos;path&apos;], $attribs);
}
</pre>



-------------------------------



...which is then integrated with `theme('links',$primary_links)` in the page.tpl.php.

This is now longer possible since the theme\_links function now [expects a structure](http://drupal.org/node/64279#menu-links) for each link instead of the complete HTML link. Now you have to override the whole **theme_links** function instead since you won't otherwise be able to access the structure of the generated links in this menu before they're getting rendered.

Not really an optimal solution for me but I guess it falls under "good enough" since the extra class I add to the link shouldn't be all that bad. Anyway: If someone has found a different solution, please let me know :-)