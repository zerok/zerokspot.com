---
date: '2012-09-16T12:00:00-00:00'
language: en
tags:
- jamjs
- javascript
- assets
title: Managed JavaScript with JamJS
---


No matter if your project is large or small, the moment you do anything with JavaScript on the web you usually end up not with one or two external libraries but tons of them. You start with something like jQuery for abstracting core functionality like DOM manipulation or AJAX and eventually end up with libs for things like lightboxes, object sliders and so on. Managing these dependencies becomes quite a hassle esp. when the time comes for updating them.

But also when just prototyping some ideas having to go to ten different sites to download each library and then integrating them into your project kind of gets in the way of ... well, prototyping. Right now there are a couple of tools out there for helping you managing such assets. The first I want to look at is [JamJS][jam]. Last weekend I played a little bit around with it as part of my Backbone-learning-trip. I didn't get very deep here, though, but anyway :-)

-------------------

## The basics

Let's say you're working on a small project where you need jQuery, Backbone.js and Handlebars. Just enter following command and you're set:

<pre><code>jam install jquery backbone handlebars</code></pre>

This will create a new folder within your current working directory called "jam" and put these libraries in there. Now to include those libraries and build upon them, add this for instance to your index.html files:

<pre><code>&lt;script type=&quot;text/javascript&quot; src=&quot;/jam/require.js&quot;&gt;&lt;/script&gt;
&lt;script type=&quot;text/javascript&quot;&gt;
    require([&#039;jquery&#039;, &#039;backbone&#039;, &#039;handlebars&#039;], function($, Backbone, Handlebars) {
        // ...
    });
&lt;/script&gt;</code></pre>

This is basically what the core tutorials tells you. Let's improve on this in two areas: First I wanted to put everything that is loaded by JamJS into a folder "static/js/vendor" (since my playground app uses Flask and I'm simply too lazy to register another static file handler ;-)), and second, since coding within the templates themselves is just a bad idea, let's add our own JavaScript file that then uses jQuery, backbone and handlebars.

## Choosing your own path

By default JamJS puts everything you install into a folder named "jam". There you will find a configured require.js file, a configuration for it as well as all the installed JS libraries. To be able to configure your own path for the dependencies I created a file called "package.json" within my project's root directory and added following settings:

<pre><code>{
    "name": "testproject",
    "jam": {
        "baseUrl": "static/js",
        "packageDir": "static/js/vendor",
        "dependencies": {
            "jquery": ">1.7",
            "backbone": ">0.9",
            "handlebars": null
        }
    }
}</code></pre>

The important parts here are the "baseUrl" and "packageDir" settings, which tell JamJS where your JavaScript files are located in general and the installed dependencies in particular.

I also used this opportunity to set specify my dependencies for this project here to able to run "jam install" or "jam upgrade" without having to specify each dependency explicitly on the command-line.

## Wrapping dependencies

Usually, you shouldn't include any JavaScript code in your markup (with the exception of perhaps some kind of initialization call). To get there I created my own little app file in the form of a main.js file within "static/js/testproject". All that is included there for now is a basic init function which leaves it open for future extensions:

<pre><code>define('testproject/main', ['jquery', 'backbone', 'handlebars'], function($, Backbone, Handlebars) {
    return {
        init: function() {
        }
    };
});</code></pre>

... and the markup only receives this piece of code:

<pre><code>require(['testproject/main'], function(app) {
    app.init();
});</code></pre>

## Compile for production

JamJS also comes with a "compile" command that combines and [uglyfies][ugfy] your JavaScript files. Since we have our own little app that specifies all the dependencies, we can do something like this:

<pre><code>jam compile -i testproject/main -o core.js</code></pre>

If you want to compile just the packages defined in the package.json file, just leave out the -i options.

## What I like about it and don't

That's basically as far as I got with JamJS there. From what I've seen so far I really like most of it. I like that it's based around [Require.js][rjs]/[AMD][amd] and that it focuses exclusively on JavaScript and doesn't try to be a one-size-fits-all packaging solution.

What I don't like about it, though, is that the whole repository solution is kind of weird. Coming from [homebrew][brew] where the package itself is a rather transparent object, I was a little bit surprised when I noticed that the JamJS repository is basically a mostly centralized [CouchDB][couch] application for which there doesn't yet seem to be an easy way to role out your own intra-net clone.

Part of this is also that you can't really see in advance what got changed on a package to get it included in the JamJS repository. Let's take the [jQuery package](http://jamjs.org/packages/#/details/jquery) as an example:

<figure><img src="/media/2012/jam-jquery.png" alt="">
    <figcaption><p>So what makes this a package?</p></figcaption>
</figure>

The package page only includes the most basic information about the package. Who contributed it, when and at what version it is at. You can also download the tar.gz which is also used by jamjs itself. Perhaps I'm missing something but I'd really like to see the package's content without having to download a tar.gz first. Perhaps a straight integration with Github or Bitbucket might do the trick.

I'm also not really convinced that naming the main configuration file for a JamJS project/package "package.json" is such a good idea given the possible clash/confusion with npm. Here Bower's "component.json" makes more sense, IMO.

As I mentioned, I really like the basic idea behind JamJS and also its UI. But somehow the whole repository system kind of feels weird right now. I guess, I will just have to keep looking for that perfect solution that most likely doesn't event exists ;-)

[jam]: http://jamjs.org
[brew]: http://mxcl.github.com/homebrew/
[couch]: http://couchdb.apache.org/
[rjs]: http://requirejs.org/
[amd]: https://github.com/amdjs/amdjs-api/wiki/AMD
[ugfy]: https://github.com/mishoo/UglifyJS
