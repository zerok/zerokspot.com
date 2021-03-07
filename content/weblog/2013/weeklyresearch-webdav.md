---
date: '2013-09-07T23:33:53+02:00'
language: en
tags:
- weeklyresearch
title: 'Weekly Research: WebDAV'
---


I've been using [WebDAV][wiki] for many years now to synchronize multiple OmniFocus installations but I never knew all that much about it apart from being like a virtual filesystem strapped onto HTTP. So this week I wanted to learn a little bit more about it.

---------------

## The premise

Back when Tim Berners-Lee authored HTTP the whole idea of it all was to have a read-write web. The latter part got quite complicated over the following years and in 1996 WebDAV was started as part of a [workshop for distributed authoring][his] on the web.

WebDAV is an extension to HTTP that adds new methods, status codes and more in order to bring better management facilities for resources to the protocol. This includes features like listing resources within a collection, storing and retrieving complex properties of resources, moving resources from one collection to another and locking.

All of this was specified in two RFCs by the IETF: [RFC2518][] in 1998 which was replaced by [RFC4918][] in 2007.

While in HTTP a resource has only a limited set of metadata represented by its HTTP headers, in WebDAV these can be much more complex. And for everything that is a little bit more complex than a simple string or number the IETF back in the late 1990s and 2000s *loved* XML; so if you work with WebDAV you work with XML.

## Data structures

First of all, to be able to move, lock or describe something we have to get down to what that something might be. In WebDAV a resource can not only be a simple resource but can also be a container for other resources, called a "collection". To stay with the file-system, think about a collection as a folder/directory. An item in a collection is called a "member" of that collection.

Each resource (be it a collection or something like an HTML page) can have multiple properties that are described using XML in all its glory (including schema definitions, DTDs etc. if you want). A property can either be "live" or "dead". While a dead property is something that the client has to maintain, a live property is one that is enforced by the server (like a file's size automatically attached to a response).

## Operations

For these data structures, WebDAV provides a handful of new HTTP methods:

* PROPFIND
* PROPPATCH
* MKCOL
* COPY
* MOVE
* LOCK
* UNLOCK

But since also the semantics of already defined HTTP methods have been changed here, I will just describe some common operations here.

To give all this a try I've configured my local Apache installation to provide me with an /uploads/ directory that supports WebDAV. All the code samples were executed against that using curl.

### Creating a resource/collection

For creating a collection, WebDAV has added a new command to HTTP: [MKCOL][]

<pre><code>$ curl -i -X MKCOL http://localhost:8080/uploads/collection/
HTTP/1.1 201 Created
Date: Sat, 07 Sep 2013 14:31:24 GMT
Server: Apache/2.2.22 (Unix) DAV/2 mod_ssl/2.2.22 OpenSSL/0.9.8x
Location: http://localhost:8080/uploads/collection/
Content-Length: 194
Content-Type: text/html; charset=ISO-8859-1

...</code></pre>

If this is successful, 201 Created is returned. If the collection already existed there, you get a 405 Method Not Allowed.

Now that we have a collection to store stuff into, let's create a test.txt file in it:

<pre><code>$ curl -i -X PUT http://localhost:8080/uploads/collection/test.txt
HTTP/1.1 201 Created
Date: Sat, 07 Sep 2013 14:27:10 GMT
Server: Apache/2.2.22 (Unix) DAV/2 mod_ssl/2.2.22 OpenSSL/0.9.8x
Location: http://localhost:8080/uploads/collection/test.txt
Content-Length: 200
Content-Type: text/html; charset=ISO-8859-1

...
</code></pre>


### Managing properties of a resource

The big feature of WebDAV is the whole property-management thing. So let's see what mod_dav stores for our newly created test.txt (I've formated the output a little bit for better readability):

<pre><code class="language-shell">$ curl -i --header "Depth: 0" -X PROPFIND http://localhost:8080/uploads/collection/test.txt
HTTP/1.1 207 Multi-Status
Date: Sat, 07 Sep 2013 14:35:38 GMT
Server: Apache/2.2.22 (Unix) DAV/2 mod_ssl/2.2.22 OpenSSL/0.9.8x
Content-Length: 902
Content-Type: text/xml; charset="utf-8"

&lt;?xml version=&quot;1.0&quot; encoding=&quot;utf-8&quot;?&gt;
&lt;D:multistatus xmlns:D=&quot;DAV:&quot;&gt;
    &lt;D:response xmlns:lp1=&quot;DAV:&quot; xmlns:lp2=&quot;http://apache.org/dav/props/&quot;&gt;
        &lt;D:href&gt;/uploads/collection/test.txt&lt;/D:href&gt;
        &lt;D:propstat&gt;
            &lt;D:prop&gt;
                &lt;lp1:resourcetype/&gt;
                &lt;lp1:creationdate&gt;2013-09-07T14:35:35Z&lt;/lp1:creationdate&gt;
                &lt;lp1:getcontentlength&gt;0&lt;/lp1:getcontentlength&gt;
                &lt;lp1:getlastmodified&gt;Sat, 07 Sep 2013 14:35:35 GMT&lt;/lp1:getlastmodified&gt;
                &lt;lp1:getetag&gt;&quot;1d9906e-0-4e5cc11689bc0&quot;&lt;/lp1:getetag&gt;
                &lt;lp2:executable&gt;F&lt;/lp2:executable&gt;
                &lt;D:supportedlock&gt;
                    &lt;D:lockentry&gt;
                        &lt;D:lockscope&gt;&lt;D:exclusive/&gt;&lt;/D:lockscope&gt;
                        &lt;D:locktype&gt;&lt;D:write/&gt;&lt;/D:locktype&gt;
                    &lt;/D:lockentry&gt;
                    &lt;D:lockentry&gt;
                        &lt;D:lockscope&gt;&lt;D:shared/&gt;&lt;/D:lockscope&gt;
                        &lt;D:locktype&gt;&lt;D:write/&gt;&lt;/D:locktype&gt;
                    &lt;/D:lockentry&gt;
                &lt;/D:supportedlock&gt;
                &lt;D:lockdiscovery/&gt;
                &lt;D:getcontenttype&gt;text/plain&lt;/D:getcontenttype&gt;
            &lt;/D:prop&gt;
            &lt;D:status&gt;HTTP/1.1 200 OK&lt;/D:status&gt;
        &lt;/D:propstat&gt;
    &lt;/D:response&gt;
&lt;/D:multistatus&gt;
</code></pre>

As the example request indicates, you can also specify a "depth" for which properties should be returned. This is relevant for collections where you can fetch requests not only for a single resource but for multiple members of a collection. [According to the specs](http://tools.ietf.org/html/rfc4918#section-9.1) this is a required parameter with the value 0, 1 or "infinity", but Apache will assume an infinite depth request if you omit it. And if you do that on a collection, you get a nice info that the server doesn't like it one bit:

<pre><code class="language-shell">$ curl -i --header "Depth: infinity" -X PROPFIND http://localhost:8080/uploads/collection/
HTTP/1.1 403 Forbidden
Date: Sat, 07 Sep 2013 18:09:48 GMT
Server: Apache/2.2.22 (Unix) DAV/2 mod_ssl/2.2.22 OpenSSL/0.9.8x
Content-Length: 235
Content-Type: text/html; charset=ISO-8859-1

&lt;!DOCTYPE HTML PUBLIC &quot;-//IETF//DTD HTML 2.0//EN&quot;&gt;
&lt;html&gt;&lt;head&gt;
&lt;title&gt;403 Forbidden&lt;/title&gt;
&lt;/head&gt;&lt;body&gt;
&lt;h1&gt;Forbidden&lt;/h1&gt;
&lt;p&gt;PROPFIND requests with a Depth of &quot;infinity&quot; are not allowed for /uploads/collection/.&lt;/p&gt;
&lt;/body&gt;&lt;/html&gt;</code></pre>

But just requesting info about the direct members of the collection (Depth: 1) works. Turns our that there is a [setting](http://httpd.apache.org/docs/2.2/mod/mod_dav.html#davdepthinfinity) for that which prevents infinite-depth requests by default.

But back to the properties returned with our initial request. Here nothing out of the ordinary for something that a file-system would report over a file shows up. Permissions, last modified time, creation date, file-type and size. There is also information about supported locking methods which we will get into a little bit later.

The [PROPFIND][] method also let's you request only a subset of properties associated with a resource. Let's say all we want is the file's size. Then we would have to submit following request body:

<pre><code class="language-xml">&lt;?xml version=&quot;1.0&quot; encoding=&quot;utf-8&quot; ?&gt;
&lt;D:propfind xmlns:D=&quot;DAV:&quot;&gt;
    &lt;D:prop&gt;
        &lt;D:getcontentlength /&gt; 
    &lt;/D:prop&gt;
&lt;/D:propfind&gt;</code></pre>

<pre><code class="language-shell">$ curl -i --header "Depth: 0" --header "Content-Type: text/xml" --data @propfind-size.xml -X PROPFIND http://localhost:8080/uploads/collection/test.txt
HTTP/1.1 207 Multi-Status
Date: Sat, 07 Sep 2013 18:33:23 GMT
Server: Apache/2.2.22 (Unix) DAV/2 mod_ssl/2.2.22 OpenSSL/0.9.8x
Content-Length: 365
Content-Type: text/xml; charset="utf-8"

&lt;?xml version=&quot;1.0&quot; encoding=&quot;utf-8&quot;?&gt;
&lt;D:multistatus xmlns:D=&quot;DAV:&quot; xmlns:ns0=&quot;DAV:&quot;&gt;
    &lt;D:response xmlns:lp1=&quot;DAV:&quot; xmlns:lp2=&quot;http://apache.org/dav/props/&quot;&gt;
        &lt;D:href&gt;/uploads/collection/test.txt&lt;/D:href&gt;
        &lt;D:propstat&gt;
            &lt;D:prop&gt;
                &lt;lp1:getcontentlength&gt;0&lt;/lp1:getcontentlength&gt;
            &lt;/D:prop&gt;
            &lt;D:status&gt;HTTP/1.1 200 OK&lt;/D:status&gt;
        &lt;/D:propstat&gt;
    &lt;/D:response&gt;
&lt;/D:multistatus&gt;</code></pre>

Now let's add a "dead" property to all these live ones using [PROPPATCH][]:

<pre><code class="language-xml">&lt;?xml version=&quot;1.0&quot; encoding=&quot;utf-8&quot; ?&gt;
&lt;D:propertyupdate xmlns:D=&quot;DAV:&quot; xmlns:dc=&quot;http://purl.org/dc/elements/1.1/&quot;&gt;
    &lt;D:set&gt;
        &lt;D:prop&gt;
            &lt;dc:language&gt;en&lt;/dc:language&gt;
        &lt;/D:prop&gt;
    &lt;/D:set&gt;
&lt;/D:propertyupdate&gt;</code></pre>

<pre><code class="language-shell">$ curl -i --header "Depth: 0" --header "Content-Type: text/xml" --data @proppatch-set.xml -X PROPPATCH http://localhost:8080/uploads/collection/test.txt
HTTP/1.1 207 Multi-Status
Date: Sat, 07 Sep 2013 18:42:36 GMT
Server: Apache/2.2.22 (Unix) DAV/2 mod_ssl/2.2.22 OpenSSL/0.9.8x
Content-Length: 322
Content-Type: text/xml; charset="utf-8"

&lt;?xml version=&quot;1.0&quot; encoding=&quot;utf-8&quot;?&gt;
&lt;D:multistatus xmlns:D=&quot;DAV:&quot; xmlns:ns1=&quot;http://purl.org/dc/elements/1.1/&quot; xmlns:ns0=&quot;DAV:&quot;&gt;
&lt;D:response&gt;
&lt;D:href&gt;/uploads/collection/test.txt&lt;/D:href&gt;
&lt;D:propstat&gt;
&lt;D:prop&gt;
&lt;ns1:language/&gt;
&lt;/D:prop&gt;
&lt;D:status&gt;HTTP/1.1 200 OK&lt;/D:status&gt;
&lt;/D:propstat&gt;
&lt;/D:response&gt;
&lt;/D:multistatus&gt;</code></pre>

The same structure can also be used to remove a property:

<pre><code class="language-xml">&lt;?xml version=&quot;1.0&quot; encoding=&quot;utf-8&quot; ?&gt;
&lt;D:propertyupdate xmlns:D=&quot;DAV:&quot; xmlns:dc=&quot;http://purl.org/dc/elements/1.1/&quot;&gt;
    &lt;D:remove&gt;
        &lt;D:prop&gt;
            &lt;dc:language /&gt;
        &lt;/D:prop&gt;
    &lt;/D:remove&gt;
&lt;/D:propertyupdate&gt;</code></pre>


### Updating a resource

Now, if we want to change the content of our test.txt resource, we just execute a PUT request on it:

<pre><code>$ curl -i -X PUT --header "Content-Type: text/plain" --data "some content" http://localhost:8080/uploads/collection/test.txt
HTTP/1.1 204 No Content
Date: Sat, 07 Sep 2013 18:52:17 GMT
Server: Apache/2.2.22 (Unix) DAV/2 mod_ssl/2.2.22 OpenSSL/0.9.8x
Content-Length: 0
Content-Type: text/plain</code></pre>

### Fetching a resource

To fetch the data associated with a resource, simply do a GET request on its URL:

<pre><code>$ curl -i http://localhost:8080/uploads/collection/test.txt
HTTP/1.1 200 OK
Date: Sat, 07 Sep 2013 18:52:22 GMT
Server: Apache/2.2.22 (Unix) DAV/2 mod_ssl/2.2.22 OpenSSL/0.9.8x
Last-Modified: Sat, 07 Sep 2013 18:52:17 GMT
ETag: "1d9906e-c-4e5cfa7707a40"
Accept-Ranges: bytes
Content-Length: 12
Content-Type: text/plain

some content</code></pre>

When it comes to collections, the GET method has no standardized behavior here. Apache, by default, will just send you a 403 error. And even if you enable "Options Indexes" in your configuration, you only get a listing as HTML, not something you'd use as a directory listing. So how do you find the members of a collection?

[Turns out][dir] this isn't done with GET or some other collection specific command but just with PROPFIND, which we already looked at before. Setting the request depth to a value other than 1 seems to be the way to go here.

### Deleting a resource

Deleting a resource is done using a simple DELETE request:

<pre><code>curl -i -X DELETE http://localhost:8080/uploads/collection/test.txt</code></pre>

If you do that to a collection the collection as well as all its members is removed. If a member cannot be deleted (for instance because someone else has a lock in it) then non of its parents are allowed to be deleted either, because otherwise the namespace containing that member would end up being messed up.

### Copying and moving

The [COPY][] method copies a resource from its URL to another defined using the "Destination" header:

<pre><code>$ curl -i -X COPY \
--header "Destination: http://localhost:8080/uploads/collection/test2.txt" \
http://localhost:8080/uploads/collection/test.txt</code></pre>

The [MOVE][] command combines a COPY and a DELETE into one atomic operation. If it succeeds you either get a 201 Created or a 204 No Content depending on the existence of a resource at the target location before the move.

### Locking and unlocking

As mentioned above, WebDAV also supports locking of resources so that for instance other clients may not change it. The protocol so far only supports one [lock-type][locks]: write. So while you can request a lock-scope of either shared or exclusive, the only lock itself is a write lock.

The point of a shared lock is, that multiple principals can have one on the same resource.

<pre><code>$ curl -i -X LOCK --data @lock.xml http://localhost:8080/uploads/collection/test.txt
HTTP/1.1 200 OK
Date: Sat, 07 Sep 2013 19:52:19 GMT
Server: Apache/2.2.22 (Unix) DAV/2 mod_ssl/2.2.22 OpenSSL/0.9.8x
Lock-Token: &lt;opaquelocktoken:e1471c1c-5e91-460a-aa9c-1fc7c2ab6540&gt;
Content-Length: 378
Content-Type: text/xml; charset="utf-8"

&lt;?xml version=&quot;1.0&quot; encoding=&quot;utf-8&quot;?&gt;
&lt;D:prop xmlns:D=&quot;DAV:&quot;&gt;
&lt;D:lockdiscovery&gt;
&lt;D:activelock&gt;
&lt;D:locktype&gt;&lt;D:write/&gt;&lt;/D:locktype&gt;
&lt;D:lockscope&gt;&lt;D:exclusive/&gt;&lt;/D:lockscope&gt;
&lt;D:depth&gt;infinity&lt;/D:depth&gt;
&lt;D:timeout&gt;Infinite&lt;/D:timeout&gt;
&lt;D:locktoken&gt;
&lt;D:href&gt;opaquelocktoken:e1471c1c-5e91-460a-aa9c-1fc7c2ab6540&lt;/D:href&gt;
&lt;/D:locktoken&gt;
&lt;/D:activelock&gt;
</code></pre>

To unlock the resource, take the Lock-Token returned as a response header and send it back up with the UNLOCK method:

<pre><code>$ curl -i -X UNLOCK --header "Lock-Token: &lt;opaquelocktoken:e1471c1c-5e91-460a-aa9c-1fc7c2ab6540&gt;" http://localhost:8080/uploads/collection/test.txt
HTTP/1.1 204 No Content
Date: Sat, 07 Sep 2013 19:52:38 GMT
Server: Apache/2.2.22 (Unix) DAV/2 mod_ssl/2.2.22 OpenSSL/0.9.8x
Content-Length: 0
Content-Type: text/plain</code></pre>

As we can see in the original LOCK response, you can also set a depth and a timeout for the LOCK request. Contrary to the PROPFIND command, "Depth" only supports the values 0 and "infinity" here. 0 only locks the given URL while "infinity" locks the resource as well as all members, sub-members and so forth.

The timeout is either specified as a number of seconds or set to "infinity" as documented in the [specs](http://tools.ietf.org/html/rfc4918#section-10.7).

## Extensions

As the whole properties-system and the preparation for different lock types indicates, WebDAV is not complete but has open doors for extensions. Here are just a few of them:

* SEARCH in [RFC5323][]
* Delta-V for versioning in [RFC3253][]
* Access control lists in [RFC3744][]

I didn't have the time to also look into these but they sound like a perfect match :-)


## Compatibility

The WebDAV specs also [protocol compliance][compl] classes. Class 1 acts as a base layer including all the "MUST" requirements in the specification. Class 2 more or less contains locking and Class 3 describes *all* requirements of the specification except for locking. The idea here is that you can have an implementation that supports the protocol to the letter except for locking. In this case it could be advertised as supporting classes 2 and 3.

On the other hand you could have an implementation that supports the basics as well as locking (1 and 2). Apache's [mod_dav][moddav] claims to be of that kind.

## Tools support

What makes WebDAV so interesting is that it's so widely supported. I already mentioned mod_dav which, for instance, is used for sharing your [SVN](http://svnbook.red-bean.com/en/1.7/svn.ref.mod_dav_svn.conf.html) repositories (as part of another extensions).

[nginx][ngdav] also has a module for WebDAV but it only implements a really small subset of the specification and not even supports properties. Because of that it not even supports class 1 operations.

Regarding client support, I hardly know where to begin. Basically, wherever you have some file-system support, WebDAV is probably also there :-)

[wiki]: http://en.wikipedia.org/wiki/WebDAV
[his]: http://lists.w3.org/Archives/Public/w3c-dist-auth/1996JulSep/0001.html
[rfc4918]: http://tools.ietf.org/html/rfc4918
[rfc2518]: http://tools.ietf.org/html/rfc2518
[rfc5323]: http://tools.ietf.org/html/rfc5323
[rfc3253]: http://tools.ietf.org/html/rfc3253
[rfc3744]: http://www.webdav.org/specs/rfc3744.html
[moddav]: http://httpd.apache.org/docs/2.2/mod/mod_dav.html
[dir]: http://stackoverflow.com/a/7400224
[locks]: http://tools.ietf.org/html/rfc4918#section-14.15
[compl]: http://tools.ietf.org/html/rfc4918#section-18
[ngdav]: http://nginx.org/en/docs/http/ngx_http_dav_module.html
[mkcol]: http://tools.ietf.org/html/rfc4918#section-9.3
[propfind]: http://tools.ietf.org/html/rfc4918#section-9.1
[proppatch]: http://tools.ietf.org/html/rfc4918#section-9.2
[copy]: http://tools.ietf.org/html/rfc4918#section-9.8
[move]: http://tools.ietf.org/html/rfc4918#section-9.9
