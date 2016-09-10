---
date: '2012-06-17T12:00:00-00:00'
language: en
tags:
- django
- supervisord
- deployment
title: Managing site processes with Supervisord
---


Managing a project's processes can be hard work. Usually there is not just your main web app process (or processes) but things like a task queue manager (e.g. Celery), some kind of backend/broker for that (e.g. RabbitMQ or Redis) and a bunch of workers. First I wrote simple init scripts for handling all these processes, later I moved to upstart simply because it saved me from all that boilerplate code SysV init scripts usually contain and I actually really like the whole eventing idea there. But when it comes to actually deploying an application, even upstart to some degree gets in your way. So in recent months (and thanks to [Steph][steph]) I've come to really enjoy the combination of [upstart][upstart] and [supervisord][supervisord].

The main motivation behind this combination was, that I needed a way to restart the webserver process (in my case gunicorn) during deployment without requiring a user switch (i.e. sudo). For this reason each project now has its own supervisord instance, executed by a user that has access to the project's directories and not much more. 

A small **disclaimer**: I'm currently using this or a similar setup on two sites I maintain myself and adapted it from a 3rd site I'm heavily contributing to.

--------------------

Let's work on a small example site using Django and gunicorn, called "djangosite", here. All of djangosite's files reside in /srv/www/djangosite. There is a "htdocs" folder for normal static files, an "app" folder for the django application and so on. This site is managed by a system user by the name of "www-djangosite" (you could naturally split this up even more with more users but let's keep it simple here). Now all the sites processes should be managed by a single supervisord instance run as that user in order to be able to restart relevant processes easily during deployment. For this we create a simple config file inside the /srv/www/djangosite folder:

<pre><code>[supervisord]
childlogdir = /srv/www/djangosite/logs
logfile = /srv/www/djangosite/logs/supervisord.log
logfile_maxbytes = 5000000
logfile_backups = 5
loglevel = info
pidfile = /srv/www/djangosite/supervisord.pid
umask = 022

[unix_http_server]
file = /tmp/supervisord.www-djangosite.sock
chmod = 0700

[supervisorctl]
serverurl = unix:///tmp/supervisord.www-djangosite.sock

[rpcinterface:supervisor]
supervisor.rpcinterface_factory=supervisor.rpcinterface:make_main_rpcinterface

[program:django]
autostart = true
directory = /srv/www/djangosite/app/current
command = /srv/www/djangosite/app/current/env/bin/python manage.py run_gunicorn 127.0.0.1:9011
</code></pre>

This so far only defines a single process called "django" which is actually [gunicorn][gunicorn] serving the application (static files are served by an external nginx instance) as well as access to this supervisord instance via a file socket. (/srv/www/djangosite/app/current/env/ is a site-specific virtualenv.)

I installed supervisord globally since there will probably be more than just this one site on the system and updating a whole bunch of supervisord installations would be rather pointless :-)

To start and manage the project's supervisord instance, we create a very simple upstart script and place it in /etc/init/www-djangosite.conf:

<pre><code>start on [2345]
stop on [06]
chdir /srv/www/djangosite
exec supervisord -u www-djangosite -c /srv/www/djangosite/supervisord.conf -n
</code></pre>

Running `start www-djangosite` will now start the supervisord instance for djangosite by using its config file and being run as its user. The `-n` just indicates that by no means this process should fork itself into the background (which would interfere with upstart). 

For deployment all there is to do after replacing the actual application files is this little Fabric task:

<pre><code>@task
def restart():
    run('superviord -c {cfg} restart django'.format(
        cfg=env.supervisord_cfg))
</code></pre>

... with `env.supervisord_cfg` representing the server path to /srv/www/djangosite/supervisord.conf

Before ending up with this setup I also experimented with putting supervisord itself into the same virtualenv that is used by the rest of the site which made it possible not to explicitly set the path to the config file all the time. The downside was the overhead of having multiple installations of the same tool without any apparent benefit.

While I'm quite happy with my current setup there are always new tools entering the scene. So perhaps I will give [circus][circus] a try once it's a bit matured ;-)

What are you using for managing your site's processes right now?

[supervisord]: http://supervisord.org/
[upstart]: http://upstart.ubuntu.com/
[steph]: http://sjaekel.com/
[gunicorn]: http://gunicorn.org/
[circus]: http://pypi.python.org/pypi/circus
