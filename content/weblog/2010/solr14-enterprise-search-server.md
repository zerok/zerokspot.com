---
date: '2010-06-20T12:00:00-00:00'
language: en
rating: 5
tags:
- books
- review
title: Solr 1.4 Enterprise Search Server
---


If you've read my post about the latest rewrite of this site, you know that
the search engine working in the background here is
[Solr](http://lucene.apache.org/solr/), an opensource search server originally
developed by CNet. During that migration I kind of learnt the basics of
working with it on its own (without some of the fancy wrappers that do the
whole configuration for you by looking at your model layer) but the schema.xml
still was kind of scary to me. So when [Packt](http://www.packtpub.com/) gave
me with ["Solr 1.4 Enterprise Search
Server"](https://www.packtpub.com/solr-1-4-enterprise-search-server/book) by
David Smiley and Eric Pugh - once again - the chance to review a book about
something that I wanted to learn, I couldn't resist :-)

----------------------------

**Disclaimer:** This is based on a review copy generously provided by
Packt Publishing.

First of all, the title is in my opinion a bit misleading. Yes, Solr
is a so-called "Enterprise Search Server", but this shouldn't stop you
from using something like it even for small projects. Usually users
expect some kind of search facility on sites of any size and integrating
Solr with for instance Django is really easy. So just do it :-)

The book consists of nine chapters starting off with giving an introduction
into working with Solr and then steadily working through topics like how
you should design your schema and how you can index various kinds of
content up to advanced topics like scaling your search setup in various
dimensions. So let's go through the chapters:

## Chapter 1: Quickstart

This chapter gives the answer to the "What is Solr anyway" question and
where you could use it in your typical stack of technologies for your
web project.

It also describes the basic installation process and where to get the
software in the first place. The authors also don't shy away from showing
you XML dumps right in the first chapter with your first search query.

## Chapter 2: Schema and Text Analysis

One of the great aspects of working with Solr is that you can also use it
for not only your usual "user types in search term and gets result" kind
of functionality but also for many other content listing aspects. But depending
on what you want to use your search store for, you have to structure the
schema accordingly.

The second chapter explains exactly this step in detail: How to denormalize
data from your primary data store in order to make all your planed search
queries possible. Here you learn about fieldtypes, fields and many more things
that make up your Solr instance's schema.xml. The basic structure of your
schema.xml is actually pretty simple but in combination with analyzers and
filters you can do tons of things with these fields.

The analyzer chains (one for query-time, one for index-time) are, what allow
features like stemming and working with synonyms, stop-words, etc. And
contrary to what the authors wrote on page 50, "filters" is IMO not a bad name
for for the analyzers that are executed after the initial tokenizing.  Data
goes through them and ideally gets transformed :-)

But already during tokenizing you can do quite a bit. For example; I wasn't
aware that there were already special tokenizers for working with HTML
content. After reading that I immediately changed my schema.xml here and
re-indexed all posts :-P

The authors also explain a bit about what filters you'd normally do at
index-time and which at query-time, although the synonym filter IMO was
not such a good choice for an index-time filter. At least for me as newbie
it looks chapter and more practical to primarily do the synonym conversion
at query-time than rebuilding the whole index once you learn of a new
synonym.

As data source for this and all the following chapters the authors have chosen
the MusicBrainz database which included data on thousands of songs and
artists. I perhaps could be argued that this might not be the best foundation
for a supposedly full-text search service but since Solr can also be used as
some kind of denormalized frontend to RDBMS it makes sense :-)

## Chapter 3: Indexing Data

So now that you know what you want to have in your index, you have to get the
data into it. And the book goes into quite a lot of detail regarding all the
options you have there. From using your usual HTTP interface to post the
documents to let Solr itself fetch the data from a passed URL or file path
or directly from some database using JDBC. Every method was described with
an example configuration, which is really great. It never felt like just
a listing of features.

This chapter also describes Solr Cell which is a tool for indexing PDFs,
Word documents and much more.

## Chapter 4: Basic Searching

The forth chapter then goes into how you can get data out of the index again
with "basic" search queries. And I was really surprised how many options
you have here. Solr also offers a way to debug your search result in order
to understand why a certain result-set was generated. And the authors seem
to absolutely love this feature, using it quite a lot in this and the next
chapters :-) And it definitely comes in handy when explaining the results
after some boosting, a term that should have perhaps already been described
in the first chapter, but anyway.

Ah, and if you're looking for a simple explanation of the query syntax (for
the standard query parser), you can find it here too :-)

## Chapter 5: Enhanced Searching

If you want to dynamically influence the score documents get on certain
queries this chapter presents function queries which offer exactly that. These
seem to be interesting for situations where you for example need the age of a
document to influence its score, but not completely overwhelm other criteria.

Another nice feature is the dismax query parser, which offers a simplified
search syntax; something that you might want to expose directly to the
user front-end. What kind of confused me here was the authors said
that the dismax request handler was deprecated. After reading it again and
looking it up on the [wiki](http://wiki.apache.org/solr/DisMaxRequestHandler)
it was clear that this only means the request handler and not the query
parser, but IMO this information should have belonged in some info box
here.

Also part of this chapter is faceting, which probably should have been given
its own chapter. It's for me personally just one of those features, that make
Solr (and other solutions offering this feature) so interesting compared to
basic full-text search service and another reason why I wanted to read this
book: To learn more about how to use faceting/how to better use it. The
authors also give a really nice example in combination with a pattern
tokenizer and synonym filter to provide facets for ranges of characters (like
A-D, E-H, etc.). Very nice :-)

## Chapter 6: Search Components

Search components are Solr's extension mechanism, which includes the already
mention facet search component. In this chapter the authors present a handful
of other core and 3rd-party components like highlighting and geo-spatial
searching that can be attached to a query handler.

What I really liked about this chapter was the inclusion of those 3rd-party
components. I so want to use the LocalSolr (geo-spatial) component right now
somewhere ;-)

Part of this chapter is also the "More like this" component which is great to
list related content like related blog posts etc. The spell-checking component
looks great, too, and, once again, I really have to mention the amount of
detail the authors provide with each example.

## Chapter 7: Deployment

Chapter 7 contains information about how to deploy Solr either using your
favorite servlet container or the bundled jetty. A great tip here is to
give all your search interfaces their own handler for abstraction.

My personal highlight of this chapter was the part about Solr cores. Having a
staging core that is active during indexing and then swapping it with the
current live-core really looks like a nice approach to provide consistent and
fast search results even during high-load updates.

The last part of this chapter was all about security and how you can secure
Solr to some extend using your container's or frontend-server security
facilities.

## Chapter 8: Integrating Solr

Next comes a whole chapter of tips and tools helping you to integrate
your web application with Solr. The authors focus here on projects using
Java, JavaScript, PHP and Ruby. There are again some really nice examples here
but the Ruby On Rails part was a bit too much for my taste. This part could
have been a bit shorter in favour of a listing of tools for other languages
and/or frameworks.

## Chapter 9: Scaling Solr

The last chapter is all about scaling Solr wide, high and deep. The authors
even went so far as to provide a disk image for Amazon EC2 in order to let
your try Solr's wide-scaling features easily for yourself. Simply a great last
chapter :-)

------------------------------

I really love this book. It showed me so many features of Solr that I didn't
know before that I still can't even find the right order I want to integrate
some of them into my projects :D

The authors didn't just list features and gave some short explanation for all
the configuration options but also always gave examples (using the MusicBrainz
database as original data store) on what to use those features for. What's
also great is that the book isn't really limited to what is bundled with the
Solr distribution but also mentions quite a lot of components that are
provided by 3rd parties. With the MusicBrainz database and working on it there
is a common thread throughout the whole book which makes it really pleasant to
read.

If you're looking for a book about how to work with Solr, this is definitely
something for you. For me personally reading this book was just a joy :D
