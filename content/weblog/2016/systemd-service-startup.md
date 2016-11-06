---
date: 2016-11-06T12:56:17+01:00
title: Service startup in Systemd
tags:
- systemd
- operations
- til
---

Due to some issues I recently had on one system I wanted to learn more about how
systemd handles dependencies and especially when it actually considers a service
to be active/started. This is by no means a complete or highly detailed guide
but simply gives an overview of what I've learnt so far 😊

The original issue was that I had one service (A) depending on another service
(B) and a port it was binding to. Sadly, B didn't bind the port immediately but
only after a couple milliseconds while systemd was already starting A. A
couldn't connect to B's port and exited.

While due to some constraints I opted in the end to just go with something like
`ExecStartPre=-/bin/sh -c 'sleep 1'` I was still curious what options I'd have
in other situations where the services are more under my control.


## Status transitions

So how does systemd know that a service has been started successfully and
therefore moves it from the status "activating" to "activated"?

According to the [systemd.service][] manpage, for "simple" services (those
service units with the type "simple") that's pretty straight forward. The
ExecStart command is started and right afterwards all depending services are
launched. No timeouts or something like that to see if the service keeps running
before continuing. If the start itself doesn't fail, everything is done.

For "forking" services (classic Unix daemons), depending services are launched
right after the configured commands exits successfully.

If you need a bit more control over that transition, you can use the "notify"
service type, which requires the service to send back a signal that it is
ready. That sounds pretty useful, so let's dig a bit deeper.


## Notifying you're ready

With a notifying service unit the whole service stays in the "activating" state
until the process signals through the `NOTIFY_SOCKET` a `READY=1` message:

The protocol for communications through that socket is simple. All you have are
newline-seperated key-value pairs like these:

    MAINPID=12345
    READY=1

The first datasets helps systemd to determine which service is saying it is
ready, while the second does just that.

There are also other messages that can be sent through that connection like when
the service wants systemd to know that it is restarting, reloading its
configuration, stopping, that it failed, ..., and also a free-form `STATUS`
value.

You can find out more on the [sd_notify][] manpage.

Setting the `MAINPID` seems only to be necessary if you send the notification
from something like a subprocess, not the process that is directly managed by
systemd. This is probably done using some low-level domain socket headers and
peer credentials. Haven't checked 😊


## A really tiny demo

Let's give all of that a try:

First, let's create a little service unit for something that should start up but
wait a couple of seconds before confirming that it is done:

    [Unit]
    Description=Systemd Notify Demo

    [Service]
    ExecStart=/tmp/systemd-notify-demo
    Type=notify

    [Install]
    WantedBy=multi-user.target

The actual service code looks like this:

    package main

    import (
        "errors"
        "net"
        "os"
        "time"
    )

    var ErrNoSocket = errors.New("No NOTIFY_SOCKET set")

    func notify(msg string) error {
        socketPath := os.Getenv("NOTIFY_SOCKET")
        if socketPath == "" {
            return ErrNoSocket
        }
        addr := &net.UnixAddr{
            Name: socketPath,
            Net:  "unixgram",
        }
        conn, err := net.DialUnix(addr.Net, nil, addr)
        if err != nil {
            return err
        }
        defer conn.Close()
        _, err = conn.Write([]byte(msg))
        return err
    }

    func main() {
        notify("STATUS=Process started. Setting up...")
        time.Sleep(5 * time.Second)
        notify("READY=1")
        notify("STATUS=Setup complete. Enjoy!")
        for {
            time.Sleep(1 * time.Second)
        }
    }

What's really neat is that you can set the status even before the service is
marked as ready. This way you can have in one line displayed some kind of
progress information. That status is displayed right below the "Main PID" in the
output of `systemctl status systemd-notify-demo.service`:

    ● systemd-notify-demo.service - Systemd Notify Demo
    Loaded: loaded (/etc/systemd/system/systemd-notify-demo.service; disabled)
    Active: active (running) since Sat 2016-11-05 15:46:47 UTC; 8min ago
    Main PID: 5145 (systemd-notify-)
    Status: "Setup complete. Enjoy!"
    CGroup: /system.slice/systemd-notify-demo.service
            └─5145 /tmp/systemd-notify-demo


## What else?

In case you don't want to do the socket communication yourself, systemd also
comes with a little command called [systemd-notify][] which lets you send
notifications, for instance, from your shell scripts.

[RabbitMQ][] among others considered it for their startup system but eventually
opted for a native implementation as systemd-notify turned out to
be [unreliable][] when executed by a non-root user.

For me personally, I'd probably try to go down to the socket whenever I can
simply because it has only a very limited amount of overhead compared to
launching a subprocess.

In any case you might also want to look into the [NotifyAccess][] service
setting which controls, what processes within the service's cgroup can send
status updates.

## Conclusion

While this hasn't really helped me solve my initial problem, the systemd
notification system looks like a way I might use in the future when I have
systems that require, for instance, an HTTP service running. I'm not yet sure,
though, how applicable all of this is when combined with Docker. I've
seen [systemd-docker][] but haven't had the time, yet, to give it a try. Perhaps
something for the next weekend 😉


[sd_notify]: https://www.freedesktop.org/software/systemd/man/sd_notify.html
[systemd.service]: https://www.freedesktop.org/software/systemd/man/systemd.service.html#Type=
[unreliable]: https://github.com/systemd/systemd/issues/2739
[rabbitmq]: https://github.com/rabbitmq/rabbitmq-server/issues/664
[notifyaccess]: https://www.freedesktop.org/software/systemd/man/systemd.service.html#NotifyAccess=
[systemd-docker]: https://github.com/ibuildthecloud/systemd-docker
[systemd-notify]: https://www.freedesktop.org/software/systemd/man/systemd-notify.html
