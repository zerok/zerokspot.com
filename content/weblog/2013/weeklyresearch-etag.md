---
date: '2013-08-31T22:58:42+02:00'
language: en
tags:
- etag
- caching
- http
- weeklyresearch
title: 'Weekly Research: ETag'
---


There are always little things you use all the time (or at least rather often) but probably don't know enough about for your own liking. For me, this was the reason to add a new fix-point to my week's schedule: "Weekly Research". In this I want to look into quite specific topics and techniques and summarize the most important aspects from my point of view.

Last week I had chosen HTTP entity-tags, better known as [ETags][http11etag], as the topic of the week. Starting with the entry on [Wikipedia][wiki] I made my way through quite a lot of documentation and mailinglists to fill some of my knowledge-gaps about this HTTP header.

------------

For those of you who've never heard of that before, entity-tags are part of HTTP's caching system. For caching in HTTP you have two main goals:

1. Prevent that the user makes a subsequent request on a resource already fetched before by indicating for how long that resource can be expected to stay the same. There the server sends an expiration date before which the client may choose not to ask for that resource again once it has been received and stored.

2. If a resource doesn't lend itself to having an expiration date let the client indicate what version of a resource it already fetched and don't send the whole resource back down again if it hasn't changed and thereby save bandwidth and ideally computational resources by not having to generated or procure the resource once again.

ETags handle scenario 2. They are basically a checksum or version-indicator that the server sends with a resource. Think about a large image being requested by a client as part of a website. Before completing the response, the server generates an ETag (for instance the MD5 sum of the file) and puts it into the response's header. The client then caches the image together with the entity-tag.

Next time, when the client wants to the same image, it sends a request to the server with the stored checksum attached as "If-None-Match" header. The server then checks the value of this header against the ETag it would have generated for the image. If they match, that means that the client still has the right resource cached and the server only responds with a HTTP 304 Not modified status code, saving bandwidth.

## How to use

For static resources like images you usually don't have to do much. Apache and nginx both offer ETag generation out of the box by just flipping a switch.

nginx has since 1.3.3 the [etag-setting][etagnginx]:

<pre>etag on;</pre>

[Apache][etagapache], too, has built-in support for generating ETags and goes even one step further by letting you customize the way they are generated. You can, for instance, put following line into your config ...

<pre>FileETag INode MTime Size</pre>

... and the ETag will be generated based on the file's i-node, last modification time and file-size in bytes. Depending on your setup, removing the INode flag will probably solve a couple of [issues if you operate behind balancers][abal] but in general I think it's great that the option is there.

Also, most web frameworks come with some degree of support for ETags. [Django][dj], for instance has some view decorators and [Rails][ra] has a helper on a similar layer. [Spring also comes with a filter][spfil] that handles ETags but it looks rather limited compared to the other two.


## 1 resource, 1 entity-tag ... what about gzip?

As long as a resource doesn't change it's ETag shouldn't change either. This is also necessary for being able to continue byte-range requests based on the previously received tag. gzip is kind of a problem here since it changes the representation of the resource and you therefor can't continue a byte-range request if the server all of a sudden switches from gzip to non-gzip representation for this resource between requests while still serving the same entity tag. Probably for that and other reasons [nginx][nxnogzip] doesn't automatically attach an ETag to something that has been gzip-compressed. Apache, on the other hand, seems to append the string "-gzip" to the tag in order to distinguish the two representations.


## Close enough

The issue above could for instance also be avoided by ignoring subrange GET requests. If you don't care about "continuing" on previous data-packages, a gzipped representation and a not-gzipped one of the same resource might be semantically identical in your use-cases. For things like that there are actually two kinds of entity tags that are distinguished by [how they are validated][val]:

1. Those supporting strong validation: <code>ETag: "12345"</code>
2. and those that only support weak validation: <code>ETag: W/"12345"</code>

The HTTP/1.1 specs have a nice [distinction][valdist] between these two:

>  In order to be legal, a strong entity tag MUST change whenever the associated entity value changes in any way. A weak entity tag SHOULD change whenever the associated entity changes in a semantically significant way. 

Personally, I interpret that in this way: Think about a blog post page which also has a list of banners on the bottom that are automatically rotated. If your ETag is based on the last-modification-time of your blog software (say its version number) and the last modification time of your blog post, then you could use a weakly validated entity tag if you consider the banners not relevant for the semantics of the resource. Otherwise you'd have generate a new entity tag whenever the banners should be rotated.

Combine that with the issue that the server you have in front of your application may or may not gzip the app's response data, going here with a weak ETag sounds more and more like a good approach to me (please let me know if I'm wrong here!).

## Enough for round one

I hope you got something useful out of this. The plan is, to write one of these posts every week if time permits.

[etagnginx]: http://nginx.org/en/docs/http/ngx_http_core_module.html#etag
[nxnogzip]: http://nginx.2469901.n2.nabble.com/etag-support-td7585448.html
[abal]: http://joshua.schachter.org/2006/11/apache-etags.html
[etagapache]: http://httpd.apache.org/docs/2.2/mod/core.html#fileetag
[http11etag]: http://www.w3.org/Protocols/rfc2616/rfc2616-sec14.html#sec14.19
[val]: http://www.w3.org/Protocols/rfc2616/rfc2616-sec13.html#sec13.3.3
[valdist]: http://www.w3.org/Protocols/rfc2616/rfc2616-sec13.html#sec13.3.4
[spfil]: http://static.springsource.org/spring/docs/3.2.x/spring-framework-reference/html/mvc.html#mvc-etag
[dj]: https://docs.djangoproject.com/en/dev/topics/conditional-view-processing/
[ra]: http://api.rubyonrails.org/classes/ActionController/ConditionalGet.html#method-i-fresh_when
[wiki]: http://en.wikipedia.org/wiki/HTTP_ETag
