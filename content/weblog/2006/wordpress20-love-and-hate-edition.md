---
date: '2006-01-06T12:00:00-00:00'
language: en
tags:
- wordpress
title: 'Wordpress 2.0: The Love or Hate edition'
---


It's strange. Normally when a new version of a web application enters the net you see basically a general opinion about it. With Wordpress 2.0 you get both, good and bad reviews judging from what perspective the reviewer looks upon it.

When I first tried it, I actually found nothing really special about it. Good, it had TinyMCE integrated, it has the upload form now directly under the posting textarea and other neat stuff, but nothing that I thought would be useful for my own blogging needs and habits. Then I suddenly really liked the TinyMCE integration and the upload feature just to disable it altogether just a few days later. 



-------------------------------


Since its unofficial release on the 26th December the net slowly but steadily fills with reviews like the [article by Joe 'Zonker' Brockmeier on newsforge](http://internet.newsforge.com/article.pl?sid=06/01/04/2324241&from=rss) today which is definitely one of the longest reviews I've seen so far. It's also a mostly positive review while others out there focus more on the stuff that doesn't work or isn't all that great. A very prominent part of this review is the new posting interface which is IMO quite well done. While Elliotte Harold raised some quite valid points on the official mailing list (for example that it now takes you longer to post if you're using categories since they are by default collapsed) I still think, that it makes the whole interface less stuffed. If you want to use categories, simply hit the "+" sign of its box and the box will stay open for your next visits. 

I general I like the JavaScripting of the whole category-selection since you now are able to select them in a "tagging style" (simply enter the category names as a comma separated listing) which will also create not-existing categories on-the-fly. While this is still IMO not the perfect solution it's a step in the right direction for me as someone who in general prefers the way of tagging.

Back to the editor and its upload component, though. As [Michael Arrington said on TechCrunch](http://www.techcrunch.com/2005/12/28/wordpress-20-the-good-and-the-bad-2/) 

> The new image uploader is a train wreck. 

Yupp. While the idea of putting the upload form right below the posting form is definitely a good idea it works a little bit too much in the dark. It for example generates thumbnails of uploaded pictures and only offers you those when you drag and drop them into the WYSIWYG editor. Here it would have been better to offer both or at least import a thumbnail that is linked to the original image. It's also not possible to configure this whole behaviour in any way thought the admin interface. It would be for example nice to be able to set some kind of upload limit for users of specific user groups, disable uploads altogether for them or simply change the size of the thumbnails without having to change the respective lines in the source code. Generally speaking this feature looks like it was a little bit rushed into the release. This is actually the only thing I don't like about WP2.0.

Ok, the TinyMCE editor is a little bit hard to extend but since I don't use it at all, it's not that much a tragedy anyway. Actually, I can't really write anything useful about the other new enhancements because I don't really use them. For the backend changes: Anything that speeds WP up is a good thing ;)

Judging from the interface changes alone I'm still not sure it they justify the jump from 1.5 to 2.0. Perhaps it's just that hate-love thing again. Many people out there will probably love the new features while I for example since 2 days only use the positng textarea for the final posting. For writing my posts I'm using [TextMate](http://www.macromates.com) in combination with [Markdown](http://daringfireball.net/projects/markdown/). Perhaps just another step for me to get away from doing things online to save some bandwidth ;)

Anyway, WP is still my first choice for traditional blogging since it offers thanks to its post metadata feature an easy way to extend your posts with things like ratings, "Now listening to" stuff and similiar things. Some people also use it as a full CMS, which is something I still wouldn't do. Esp. for community sites I think Drupal is simply the better solution out there. For personal sites it might be enough but on the other hand it also might be an overkill. For example: You have a static page which only includes your imprint. If you manage it using Wordpress it requires much more file accesses than if you would use a static HTML page which would do the job nearly as well. Ok, editing the site would become harder, but at least I don't change my contact information every few days, so going thought the extra-work of editing the file and replacing it on the server isn't such a tragedy either, is it? ;) Before I get more OT here, I probably should submit this post \*submit\*