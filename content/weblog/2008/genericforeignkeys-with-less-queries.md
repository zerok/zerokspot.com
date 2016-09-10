---
date: '2008-08-13T12:00:00-00:00'
language: en
tags:
- contenttypes
- django
- generic
- genericforeignkey
title: GenericForeignKeys with fewer queries
---


When working with [generic relations](http://www.djangoproject.com/documentation/contenttypes/) in [Django](http://www.djangoproject.com) you have to be quite careful not to end up with n+1 queries for a simple fetch of n elements. The reason for this is that internally a generic relation is not really a true foreign key (naturally) but just an id combined with a foreign key to a content-type. But there are some ways around this problem. Among them a quite simple one: Doing the actual content-loading by yourself.


-------------------------------

Inspired by [Ryan Berg's](http://ryanberg.net/blog/2008/jun/24/basics-creating-tumblelog-django/) tutorial about how to build a small tumblelog with Django, let's adapt this example a little bit. If you're using generic foreign keys you will probably end up with a structure like this:
    
    class Item(models.Model):
        pub_date = models.DateTimeField()
        content_type = models.ForeignKey(ContentType)
        object_id = models.IntegerField()
        content_object = generic.GenericForeignKey('content_type', 'object_id')
    
    class Post(models.Model):
        title = models.CharField(max_length=255)
        pub_date = models.DateTimeField()
        content = models.TextField()
        item = generic.GenericRelation(Item)

    class Link(models.Model):
        title = models.CharField(max_length=255)
        pub_date = models.DateTimeField()
        description = models.TextField(null=True, blank=True)
        url = models.URLField()
        item = generic.GenericRelation(Item)

    class Photo(models.Model):
        title = models.CharField(max_length=255)
        pub_date = models.DateTimeField()
        item = generic.GenericRelation(Item)

Combined with a simple signal handler that gets triggered every time you change a Post, Link or Photo instance and that updated the `Item` of that instance, this gets the job done pretty nicely from the writing-point-of-view.
    
    from django.db.models import signals

    def update_item(instance, raw, created, **kwargs):
        if created:
            item = Item()
            item.content_type = ContentType.objects.get_for_model(type(instance))
            item.object_id = instance.id
        else:
            item = instance.item.all()[0]
        item.pub_date = instance.pub_date
        item.save()
    
    signals.post_save.connect(update_item, Post)
    signals.post_save.connect(update_item, Photo)
    signals.post_save.connect(update_item, Link)

When it comes to reading that data, you'd normally not want to use something like this::
    
    Item.objects.select_related().all()

... for the simple reason that it would pull every related object in its own query. But since there is probably a much lower number of content types in your model structure than items, you could lower that count from n+1 queries (where n is the number of items) to something more like 1+m (where m is the number of models queried through the `content_type` property of the `Item`-class).

The trick is pretty simple: Don't do `.select_related()`, but be a bit more specific what related object you actually want to have. For now, all that is actually needed is (as the whole m-thing from above indicated) the content-type of each item. So a `.select_related('content_type')` is enough. With this we end up having all the references and content-types within one query (at least after the content-type's have been cached). Now all that is left to do is to do one query for each content type to get the actual Posts, Items and Photos that were referenced in the Items:
    
    items = Item.objects.select_related('content_type').all()
    model_map = {}
    item_map = {}
    for item in items:
        model_map.setdefault(item.content_type, {}) \
                [item.object_id] = item.id
        item_map[item.id] = item
    for ct, items_ in model_map.items():
        for o in ct.model_class().objects.select_related() \
                .filter(id__in=items_.keys()).all():
            item_map[items_[o.id]].content_object = o

In order not to repeat myself here, I simply put that snippet into a simple manager and associated it as secondary manager with the `Item` class.

With something like this in place, GenericForeignKeys are once again quite high on my list of features I really like about Django. Sometimes, as nice as it is to have, model inheritance simply isn't what you want and for something like a tumblelog where you just want to have a meta-object that helps you basically merge queries, they are IMO simply still the way to go. And with < n+1 queries for a simple page, all the better ... ;-)