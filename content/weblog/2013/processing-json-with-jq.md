---
date: '2013-07-18T22:31:11+02:00'
language: en
tags:
- json
- tooling
title: Processing JSON with jq
---


In case you have to process a large JSON file on the command-line for instance to extract a certain element or transform it into a new, simpler format, take a look at Stephen Dolan's tool [jq][]. My initial use-case to finally (after having it on my list for quite a while now) using it was that I had a huge but compacted JSON file and I needed to check a certain element in there.

<pre>$ cat file.json | python -m 'json.tool' > file.clean.json
$vim file.clean.json</pre>

... and then looking through the file manually might have been an option, but the file was large enough to be rather tedious to walk through manually. So I needed basically grep for JSON.

---------------

Just something about the data I worked with here: The file more or less consists of a huge array of objects. Each object representing a store location has a "name" property by which I wanted to filter. jq is basically a language that consists of filters on input data to convert it into a specific output. Think about grep+sed+awk but for JSON-structured data. Since we know that we are working with a huge array, the first filter has to be '.[]' which passes each element of that array down to the next filter.

But before we get to that, let's first check how many stores are actually in there:

<pre>$ cat file.json | jq 'length'
1134</pre>

Ok... more than expected and so really something I don't want to wade through manually. But let's take a look at all the store names first, just for curiosity:

<pre>$ cat file.json | jq '.[] | .name'
"Graz"
"Baden"
"Zürich"
...</pre>

This is also a small example of how pass data from one filter to the next. The first filter receives an array object and is instructed to pass each element on its own down to the next one. So '.name' now receives the store object as input and send the value of the "name" property out.

But I actually wanted all the data for a store with a specific name. For this, jq offers the select filter:

<pre>$ cat file.json | jq '.[] | select(.name=="Graz")'
{
    "name": "Graz",
    "openingHours": [...]
    "assortments" {...}
    ...
}
</pre>

Awesome!

At this point I had everything I wanted, but there is naturally far more to jq than just that. Extracting all the keys of an object, sorting arrays, temporarily assigning values to variables to use in later filter,... Just take a look at the [manual][] which offers tons of examples for your experimenting-pleasure :-)


[jq]: http://stedolan.github.io/jq/
[manual]: http://stedolan.github.io/jq/manual/
