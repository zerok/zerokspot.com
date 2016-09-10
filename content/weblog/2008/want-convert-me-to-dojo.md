---
date: '2008-11-08T12:00:00-00:00'
language: en
tags:
- dojo
- javascript
title: Want to convert me to Dojo?
---


Since I first heard about jQuery back when the discussion started if it should integrated into Drupal's core distribution, I've been using it on nearly every website I've been working on since then. At the recent Barcamp in Munich [Wolfram Kriesing](http://wolfram.kriesing.de/blog/) presented some of the features the [Dojo toolkit][dojo] has to offer that might make me a switcher. I still like jQuery quite a lot, but I simply wanted to take a look at some of Dojo's features and found a couple of them that I really enjoy using on a project right now :-)

-------------------------------

## sprintf for JavaScript

First of all, coming from languages like C, Java and Python I could never really understand, why JavaScript hasn't something like ``sprintf`` in its standard library. Dojo's [dojo.string.substitute][] offers similar functionality and takes a Ruby-esque approach allowing you to specify named patterns to work with during the substitution:
    
    dojo.require('dojo.string.substitute');
    var newstring = dojo.string.substitute(
        "Hello ${name}, it's ${time} o'clock", 
        {
            name: "Horst", 
            time: new Date()
        }
    );
    

## Custom context for event handler

It kind of always ended up as a small mess when I worked with tons of event handlers. For some reason I tended to forget saving the context explicitly and therefor ended up losing it altogether. To help in this regard, Dojo's [dojo.connect][] function offers a nice and easy way to build your own context from which event-handlers are then called:
    
    var myctx = {
        somevariable: "123",
        handler: function(ev){
            alert(this.somevariable);
        }
    }
    dojo.connect('elem', 'onclick', myctx, 'handler');
    
The code above registers a handler to the click-event on the element with the id "elem". When this event happens, the function "handler" from *within* ``myctx``-object will be called. This way it is very simple to know your context's state when for example registering handlers in a loop always with a different set of variables you would like to use from within the handler.

## It's as big as you want it to be

Dojo comes with a pretty sophisticated module system that allows you to [dynamically load libraries](http://dojotoolkit.org/book/dojo-book-0-9/part-3-programmatic-dijit-and-dojo/modules-and-namespaces) as you need them. This is done using the [dojo.require][] function, which you've already seen being used in the first example. This way you can, for example, use a base-package of Dojo that offers you the whole CSS-selector-toolkit. And on a subpage, were you want to use some nifty animation, just require it while still using the base package from the user's cache.

But what if you're using some functions like dojo.string.substitute all the time? No problem. Dojo comes (in its source-package) with a [build-utility](http://dojotoolkit.org/book/dojo-book-0-9/part-4-meta-dojo/package-system-and-custom-builds) that lets you combine all the modules that you'd like to use in just one JavaScript file, that you then simply load the old-fashioned way. 

## Convert your form to JSON

While working with jQuery on a project a couple of months ago, one of the main complains I had was the lack of full JSON support. Sure, you can easily treat a dataset received via a XHR as native JavaScript-dataset, but without an addon I couldn't easily take a form and send it as JSON to the backend.

Without any additional require-statements you can use the [dojo.formToJson][] function, which will go through each input element within a form and build a json-object out of it:
    
    #----- 8< ----- markup ----- 8< -----
    
    <pre id="code"></pre>
    <form id="myform" method="post">
        <input type="text" name="textfield" />
        <select name="selectfield">
            <option value="0">0</option>
            <option value="1">1</option>
        </select>
        <button type="submit">toJSON</button>
    </form>
    
    #----- 8< ----- submit handler ----- 8< -----
    
    <script type="text/javascript">
        dojo.addOnLoad(function(){
            dojo.connect(dojo.byId('myform'), 'onsubmit', 
                    null, function(ev){
                ev.preventDefault();
                dojo.html.set(dojo.byId('code'), 
                    dojo.formToJson('myform'));
            });
        });
    </script>

All this naturally comes at a price: Dojo's base package is compressed around 26KB (compared to the 15KB of jQuery), but luckily the folks behind Dojo are not only using Google's JavaScript library CDN, but also AOL's, which doesn't require you to do anything than the usual ``<script/>``-loading :-)

[dojo.formToJson]: http://api.dojotoolkit.org/jsdoc/dojo/1.2/dojo.formToJson
[dojo.string.substitute]: http://api.dojotoolkit.org/jsdoc/dojo/1.2/dojo.string.substitute
[dojo.require]: http://api.dojotoolkit.org/jsdoc/dojo/1.2/dojo.require
[dojo.connect]: http://api.dojotoolkit.org/jsdoc/dojo/1.2/dojo.connect
[dojo]: http://dojotoolkit.org/
