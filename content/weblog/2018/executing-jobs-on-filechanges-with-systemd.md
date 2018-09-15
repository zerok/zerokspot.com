---
title: "Executing jobs on file-changes with Systemd"
date: 2018-09-15T10:41:31+02:00
tags:
- ops
- development
---

For the last couple of months I've been using [Facebook's
watchman](https://facebook.github.io/watchman/) for executing a deployment
script when my CD pipeline pushed a file onto the target system. Watchman is a
bit weird when it comes to running it as a service, though, so I started
looking around thinking that this should be something that Systemd should be
able to handle all by itself. And to actually nobody's surprise I found the
feature I was looking for:
[systemd.path](https://www.freedesktop.org/software/systemd/man/systemd.path.html).

Similar to a timer you can tell Systemd to activate a service if something
happens to a file-system path. A sample setup would look like this:

```
$cat /etc/systemd/system/app-update.service
[Unit]
Name=app-update executor
After=network.target

[Service]
Type=oneshot
ExecStart=/srv/app/scripts/update.sh
```

```
$ cat /etc/systemd/system/app-update.path
[Path]
PathModified=/srv/app/etc/main.yml

[Install]
WantedBy=multi-user.target
```

Here `/srv/app/scripts/update.sh` is triggered when `/srv/app/etc/main.yml` is
modified. If you don't want to trigger the service whenever the file is
modified but only when the file descriptor is closed again after a
modification, use `PathChanged` instead.

Just make sure to enable **and** start the path file when you're done.
Otherwise nothing will happen when the file changes 😅

```
$ systemctl enable app-update.path
$ systemctl start app-update.path
```
