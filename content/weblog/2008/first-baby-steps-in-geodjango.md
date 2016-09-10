---
date: '2008-08-09T12:00:00-00:00'
language: en
tags:
- django
- geodjango
- postgis
title: First baby-steps in GeoDjango
---


For a couple of days now, [Django](http://www.djangoproject.com) trunk finally also [includes GeoDjango](http://code.djangoproject.com/changeset/8219). A geo-spatial extension module for Django that adds -- among other things -- model fields and managers to easily use backends like the geo-spatial extensions for PostgreSQL, MySQL and Oracle. Let's see, how hard it is to get it going on my little MacBook.


-------------------------------

Since I have some ideas for a couple of projects that could definitely use some GIS-enhancements, I wanted to spend some time this weekend on at least getting a really tiny app working with [GeoDjango](http://geodjango.org/), just to see what's all required to get it working (using a PostgreSQL database). And compared to other parts of Django, GeoDjango has quite a few dependencies, that are thankfully all listed on [the wiki](http://code.djangoproject.com/wiki/GeoDjangoInstall). On OSX and thanks to [MacPorts](http://www.macports.org/) the process is pretty simple. All you basically have to do, is to install postgis (`sudo port install postgis`), which will normally (or at least did in my case) pull PostgreSQL 8.3, proj and geos. The only item left for you to do, is install gdal again using MacPorts.

Next thing I did, was to create a database and load the postgis data (lwpostgis.sql and spatial_ref_sys.sql, both can be found in /opt/local/share/postgis/) into it. When you now run `python manage.py syncdb` after having created some basic models like ...

@@ python @@
from django.db import models
from django.contrib.gis.db import models as geomodels

class PointOfInterest(models.Model):
    name = models.CharField(max_length=255)
    description = models.TextField()
    coords = geomodels.PointField()

    objects = models.Manager()
    geoobjects = geomodels.GeoManager()
@@

... GeoDjango will probably complain that it doesn't know what geos library it should use:

> django.contrib.gis.geos.error.GEOSException: Could not find the GEOS library (tried "geos_c"). Try setting GEOS_LIBRARY_PATH in your settings.

Since MacPorts already installed geos for us, solving this is pretty straight forward. Simply add this line to your settings module, and syncdb should work:

@@ python @@
GEOS_LIBRARY_PATH='/opt/local/lib/libgeos_c.dylib'
@@

So far I've done just some basically playing with it like entering some of my favorite restaurants into the DB and letting GeoDjango give me the centroid there and the list of restaurants within a certain radius around my home. So for it's just been great and remarkably easy. Just the documentation (also the docstrings) could use some more examples in my opinion, but this is just a matter of time. I really can't wait to actually do something with this new module. Big kudos to everyone involved with the project :-)

For more details, take a look and keep an eye on [geodjango.org](http://geodjango.org/), where you can find among other things the current version of the project documentation (that is until it is added to Django's documentation as well ;-) )