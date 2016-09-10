---
date: '2013-09-22T21:02:36+02:00'
language: en
tags:
- weeklyresearch
- json
- jsonpointer
title: 'Weekly Research: JSON Pointer'
---


While looking into [JSON Schema][jsonschema] last week I stumbled upon [JSON Pointer][rfc6901] which is an IETF RFC by Paul C. Bryan, [Kris Zyp][kz] and [Mark Nottingham][mn] and is used to reference a specific value within a JSON object. Don’t think of it as something like XPath or XQuery, though, since they effectively allow you to *select* multiple elements that match certain criteria. For this but in JSON there are other tools like [JSON Path][jsonpath] and [JSONiq][jsoniq]. This here is more like the id-references in HTML documents using the hash-name but using the nested structure provided by the object instead or artificial IDs.

-----------------

Let’s say you have following JSON object:

```
{
    “title”: “JSON Pointer”,
    “authors”: [
        {"name": "P. Bryan"},
        {"name": "K. Zyp"},
        {"name": "M. Nottingham"}
    ]
}
```

You’d use following JSON Pointer to reference the name of the first author: `/authors/0/name`. If all you want is the whole `authors` array, just use `/authors`. Don’t append a “/” to the pointer if you want the whole array or sub-object. So, object properties are simply accessed by their name without quotes and array positions by their index... that’s basically it. As I said, this is just for addressing values and since JSON is rather simple, referencing a specific value isn’t all that hard either.

There are two special characters that have to be escaped, though, according to the specification if used within a path token: “/” and “~”. The slash is escaped with “~1” while the tilde is escaped with “~0”. The main reason for this (compared to going with something like URI-encoding) seems to have been [performance-considerations][esc], but it also helps prevent some confusion if you use pointers as URI fragments since there you’d have to use URI encoding anyway. 

## Implementations

So far I found three implements:

* [Python][ipy]
* [Perl][ipl]
* [JavaScript][ijs]

Sadly, I didn’t have the time (or the use-case) to play with any of them yet :-)

## What’s coming next?

As much as I’d like to get into JSONiq and other JSON processing and querying languages right now, I think it’s time for something else. So, next week I will either take a break or write about something that is not another RFC about how to deal with JSON. I promise ;-)

[rfc6901]: http://tools.ietf.org/html/rfc6901
[esc]: https://groups.google.com/forum/#!topic/json-schema/iXHVCJk_zfQ
[jsonpath]: http://goessner.net/articles/JsonPath/
[jsoniq]: http://www.jsoniq.org/ 
[ipl]: http://search.cpan.org/~zigorou/JSON-Pointer-0.01/lib/JSON/Pointer.pm
[ipy]: https://github.com/stefankoegl/python-json-pointer
[ijs]: https://github.com/janl/node-jsonpointer
[jsonschema]: http://zerokspot.com/weblog/2013/09/15/weeklyresearch-jsonschema/
[mn]: http://www.mnot.net
[kz]: http://www.sitepen.com/blog/author/kzyp/