---
date: '2012-06-24T12:00:00-00:00'
language: en
tags:
- development
- tooling
title: Working with Foreman
---


Recently and thanks to [apollo13][a13] I've discovered [foreman][fm] for myself and added it right away to the tools I use in nearly every web-related project. So what is foreman? Well, think about following situation: You code on a slightly more complex website within your development environment. You first start your framework's development server, then some kind of queuing backend for asynchronous tasks, then perhaps some kind of dummy email server for debugging email messages, then `compass watch` in case you change the style somewhere.

Wouldn't it be nice if you could just execute a single command and all these processes would start up and print their output into a single terminal? That's what foreman does. All you have to do is create a simple "Procfile" and put your process definitions into it.

---------------------

<pre><code>web: python manage-dev.py runserver $DJANGO_PORT
compass: compass watch pygraz_website/static</code></pre>

This is, for instance, the Procfile I have for the [PyGRAZ.org][pg] website. It defines a web-process that represents the development server of Django and compass for updating the site's css if necessary.

Once this configuration is in place, simply execute `foreman start` and lean back as your processes start up and print their output into a single terminal.

For very simple setups that's about it. If your processes need to be started from within a specific folder or require some other kind of hand-holding I'd advice that you create some small wrapper scripts to do that. At least, I had some problems with getting the whole environment set up within the definition itself so I put that whole process into a simple shell-script. It kind of goes against having your Procfile as part of the project's documentation but then you could just add your shell-scripts to the project as well :-)

Another minor issue is that some environments tend to buffer output instead of just passing it right through to stdout. There is a [wiki page][wp] for that.

Foreman also supports exporting of Procfile definitions to (among others) upstart and bluepill. There even is an exporter for supervisord I intend to give a shot in the future :-) Right now, though, I'm running slightly different processes in development compared to my production-environments. A good example here is the Django devserver but also the `compass watch` process which locally runs without compression in order to facilitate debugging. Perhaps I can unify these to environments in the future or at least make them completely configuration dependent. That's usually what the ".env" file is for in the same directory as the Procfile. Here you can set environment variables that are available within the Procfile.

Cleaning up my Procfile might also be helpful if I ever decide to put some of these websites onto [Heroku][hk] ;-)

[a13]: https://github.com/apollo13/
[fm]: http://ddollar.github.com/foreman/
[hk]: http://heroku.com
[pg]: http://pygraz.org
[wp]: https://github.com/ddollar/foreman/wiki/Missing-Output
