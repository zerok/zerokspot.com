---
date: '2010-07-14T12:00:00-00:00'
language: en
rating: 2
tags:
- django
- books
- review
title: Django 1.2 E-Commerce
---


Once again I received a review-copy of a book from Packt, for which I'm again
really grateful :-) This time it was [Django 1.2
E-Commerce](https://www.packtpub.com/django-1-2-e-commerce-build-powerful-applications/book)
by Jesse Legg. The book tries to introduce you in a tutorial style to how to
write a simple e-commerce store using Django core components and some selected
3rd-party reusable apps.

------------

The author divided that into 10 chapters:

1. A basic introduction into what Django is and why you might want to use it for an online store
2. How to construct an extremely simple shop in a couple of minutes
3. User accounts and carts and the comment framework
4. Payment processors using class based views for Google Checkout and Amazon
5. Order status combined with shipment information via Google Checkout
6. Search integration
7. Reporting via ReportLab, feeds and Salesforce
8. Enhancing the UX with some AJAX
9. Selling of digital goods using post- and pre-payment as well as subscription-style payment
10. Deployment and stuff

The authors does a decent job at mentioning most of the technologies that are
relevant for building a webshop with Django if you don't want to use
something like Satchmo or LFS. The focus on external services for payment and
shipping information can be found in the Google and Amazon product catalogs.
PayPal is only mentioned under "also available" with the explanation that the
APIs are too complicated/confusing/messy (my words) and they charge
pro-actively compared to per-order.

## Django 1.2?

That said, personally, I don't think that the book really is for its intended
audience. Let's start with the title: "Django **1.2** E-Commerce". If you read
such a title and have some experience with Django (so you actually know that 1.2
introduces some new features) you might expect some of them getting used here.
But if you look for some nice multi-db integration, heavy usage of the new
i18n features to render currency values and dates in the right format
depending on your user's locale, model validation to prevent your product
import to not mess up your whole operation, you have to look somewhere else.

This book is for Django 1.x. So you can probably also get everything running
on 1.0.x. It even says so right on page 2:

> This book requires Django 1.0 or higher

So why did it get that title? Marketing?!

## E-Commerce

Ok, let's ignore the Django version for now and focus on the second part of
the title. As already mentioned, this book takes a "do it yourself" approach
for building a simple online shop with payment, shipping and basic CRM. There
is nothing wrong with this approach, but why? Why should I do everything
myself? In some of the chapters the authors takes the recommended software
development approach of DRY and re-uses apps like django-registration and
haystack but why not start with the obvious?

There are already some quite popular e-commerce apps for Django out there:

* [Satchmo](http://www.satchmoproject.com/)
* [LFS](http://www.getlfs.com/)

But none of these are even mentioned once.

Also, if you look at the chapters listed above, only 3 out of the 10 (4, 5 and
9) are really focused on e-commerce problems. The rest is just a big tutorial
on how to do stuff with Django, which is something that has already been
covered in tons of tutorials and a couple of books (like two published by
Packt). A book with such a title should really focus on e-commerce and not
general website problems.

Don't get me wrong: What's shown isn't necessarily bad, but in my opinion
simply completely off-topic. And if most of a book about a certain topic is
actually off-topic, I kind of have to wonder what the author really wanted to
achieve and why the reviewers didn't notice that or stopped him ;-)

## Lack of Focus

Part of that problem is also, that the author tries to cover quite a lot
without really focusing on use-cases that are primarily relevant for online
shops. A good example is the chapter about search integration. Here you get to
know Xapian, Sphinx, Whoosh and shortly also Solr combined with a couple of
libraries for Python and Django integration. Great, but searching is a common
feature not only of online-stores but of websites in general.

On the other hand, what a product catalog is really well suites for, and what
is widely used, is faceting. If you don't know what "faceting" is, just take a
look at Amazon's search results and there at left side you will see options to
narrow down your resultset with additional attributes of the found products
combined with the number of products matching these additional criteria.
That's faceting. And that term is mentioned only once in relation with Solr
and solango. The rest of this chapter focuses on Whoosh, Xapian and Sphinx,
even though non of them actually supports faceting directly as far as I can
tell. Xapian offers some support using MatchSpy, but that's it.

Solr was completely left out of the rest of the chapter, probably because of
that (page 100):

> To run a standalone Solr server, you'll need a container like Apache's Tomcat or the Jetty application server.

This is technically true, but the example/start.jar actually bundles a Jetty
instance so all you need is to run that to get a working Solr server.

## Review?

But sadly not only are there a quite a few flaws in the content, but even also
in the examples. A big part of the code samples has some formating or even
syntax issues. Given all these problems I have to wonder what the reviewers
where doing. It seems like no one looked twice at the code samples.

In general, this book feels like is has been rushed out of the door and
through every single life-cycle phase starting with the concept. This is the
only reason I can think about it for what made it into the PDF I got. Don't
get me wrong: Most of the information presented here is not wrong. Chapter 9
even goes a really long way of telling you the details of how to implement
various payment and delivery solutions for digital goods but even that went
overboard without giving you any diagrams so that you can keep up with the
data that has to be sent back and forth between your site and the service
providers.


## But what else?

As already said: The main problem with this book is, that it does only offers
about 3 chapters of really targeted content. The rest is more like a general
tutorial on how to do stuff with Django. But what could have been done
differently? I already mentioned one thing: Integration of already existing
shop solutions like Satchmo and LFS. But that's only part of the story. At
least for me, real-life case studies would have been great. How did real shops
solve real-world problems? What are the general legal problems when working
with other people's money? How does all that work internationally? And also
strike the "1.2" from the title if you don't need Django 1.2 for any part of
it.

For each chapter there would have been a couple of changes that would have
really improved them.

* I already mentioned faceting for the search chapter.

* Chapter 2 could have been used as a great foundation for creating more
  complicated product scenarios. Just think about product variants as you have
  them for instance on stores selling cloths like T-Shirts being offered in
  various sizes and colours which all have to have their own stock.

* Chapter 3's customer reviews could have been depended on if the customer
  actually bought the product (which is a common problem on for instance
  Amazon). Or present a cart where products of the same type are merged. This
  would have been a great example in combination with the product variants.

* Chapter 4 could have also survived with just one concrete implementation and
  with more focused on general issues when it comes to working with payment
  providers.

* Chapter 7 is again one of the chapters with just too much in it. Start with
  the practical use-cases and continue from there.

* Get rid of chapter 8. Sorry, but this is just not the place to introduce
  people to AJAX. If you're working on any kind of web project with some form of
  interaction, you should already know the basics.

* The part about the payment integration for digital goods in chapter 9 could
  have used some diagrams to explain the data flow a bit better. As it is right
  now it just goes a bit too low-level. But the first part was really great :-)

## Enough negativity...

But there was something good about this book, right? There has to be. Yes,
there definitely was. I really liked the first part of chapter 9 where the
integration with S3 was described in order to do protected downloads. Also the
chapter about the basic payment and shipping integration was really nice. Same
goes for the next chapter that provides a good overview of what's normally
going after the payment step with the order's status getting updated and the
shipping getting initializes (although, the step about charges might have been
better placed in the chapter about payment).

If you're working in an eco-system (as is normally the case with online-shops)
being able to integrate with that world is important. People often expect to
get for instance a bill receipt as PDF. For features like that it was really
great that ReportLab and other tools got a chapter.

But sadly even there it seems like some alternatives were removed due to
time-constraints. Like only offering S3 while ignoring the simple solutions
for low-filesize-content (e.g. X-Sendfile/X-Accel-Redirect).

Also the chapter about payment systems and shipment could have been structured
a bit better. It just feels kind of weird to see a chapter about SSL not being
part of the payment chapter or even in chapter 3 where the auth system is
described but of the one about order status and shipment.

-----------------------

If you're looking for a book that teaches you how to build e-commerce
solutions with Django, you probably still have to wait a little longer. This
book, while giving quite a few very detailed instructions on how to integrate
some aspects of e-commerce - like payment providers and shipping services -
into Django, does simply not provide the whole or even the better part of the
story. Most chapters are not really e-commerce specific enough to offer more
than a "How to work with Django" guide.

Content and "syntax"-wise the book just contains too many errors, too, which
makes it really hard to recommend for anyone that just wants to learn Django
in order to build an online-shop. And especially those people will probably
earn the most by getting inspiration and tools from the Satchmo and LFS
projects.

For people who already know there way around Django, this book doesn't really
offer all that much if they want to jump into making online-shops. The 3
chapters about payment and order status transitions are a great overview of
tools for these aspects of an ecommerce system, but there is just too much
other "off-topic" content around, that I can't really recommend this book :-(
