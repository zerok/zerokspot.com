---
date: '2008-02-09T12:00:00-00:00'
language: en
tags:
- java
title: Default value handling with Json-lib
---


As part of a project I'm working on I'm currently moving tons of data between servers and clients using JSON. Since all components are written in Java I was looking for a simple yet customisable JSON library and eventually ended up using [Json-lib](http://json-lib.sourceforge.net/) which let's you easily serialise a JavaBean into a JSON-string. While looking through the network traffic yesterday, I noticed a small problem, though: When you convert for instance an Integer property of the bean and this Integer is currently a null-pointer, Json-lib will make a 0 out of it. While in most cases probably a good idea, in my particular case it's simply the wrong way of handling the null-reference since 0 has a different semantic in that particular bean.

-------------------------------

Up to [version 2.2](http://json-lib.sourceforge.net/changes-report.html#2.2) of the library there was no way to solve this problem by handling Integer-nulls differently in a nice way. Sure, you could still go through your JSONObject or JSONArray afterwards and replace the values one by one, but that would more or less render the whole library useless to me. This situation changed with 2.2, though. Now you can associate a so-called `DefaultValueProcess` (and a `DefaultValueProcessorMatcher` for that matter) with a JsonConfig instance and let for instance String-null be represented as `null` instead of `''` or Integer null values represented as `null` instead of `0` (which might be invalid depending on how you interpret the JSON specs but sometimes it's the easiest solution). Here a small example:
    
    import net.sf.json.JsonConfig;
    import net.sf.json.processors.DefaultValueProcessor;
    
    ...
    
    JsonConfig conf = new JsonConfig();
    conf.registerDefaultValueProcessor(Integer.class, 
        new DefaultValueProcessor(){
            public Object getDefaultValue(Class type){
                return JSONNull.getInstance();
            }
        });
    
This example does exactly that, it registers a new `DefaultValueProcessor` for the Integer class and returns `null` if the Integer instance is a null-pointer. Note that you can basically ignore the `type` argument of the `getDefaultValue()` method since the class this processor is bound to is already set in the `registerDefaultValueProcessor()` method. 

The following example would now take a null String and replace it with "I am empty!":
    
    conf.registerDefaultValueProcessor(String.class, 
        new DefaultValueProcessor(){
            public Object getDefaultValue(Class type){
                return "I am empty!";
            }
        });
<!--
You can find a complete example [packaged here](/media/2008/jsonlibpg-defaultvalues-9891cf56f2c4.tar.bz2>) and [versioned here](http://hg.zerokspot.com/jsonlibpg.defaultvalues/).
-->
