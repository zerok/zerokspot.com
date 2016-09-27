---
date: 2016-09-27T21:19:58+02:00
title: Error notifications for systemd timers
tags:
- operations
- systemd
- slack
---

{{< img-left path="2016/systemd-slack.png" width="280px" >}} Cronjobs are great;
them failing not so much. But thanks to the `MAILTO` setting it's relatively
easy to get notified if they don't complete successfully. I recently moved quite
a few jobs over to systemd timers simply because `systemctl list-timers` is just
too convenient not to use 😉

The downside? No MAILTO. Luckily, I stumbled upon [this post][nl] by Lars Ollén
where he mentions the `OnFailure` unit-setting:

    #>> /etc/systemd/system/sometimer.service

    [Unit]
    # ...
    OnFailure=status-slack@%n.service

    [Service]
    #...

This setting accepts the name of another unit that should be started with "%n"
being replaced with the name of the unit that failed.  As you can guess from the
snippet above, I actually don't want to have email notifications but something
that posts into our team-chat if a timer fails.

The status-slack unit is pretty straight forward. It is mostly just a OneShot
service that executes a shell script:


    #>> /etc/systemd/system/status-slack@.service

    [Unit]
    Description=Reports timer error to Slack

    [Service]
    Type=OneShot
    ExecStart=/usr/local/bin/systemd-slack %i

The script that is called here gathers status information using `systemctl status`
and forwards it to a little tool that forwards stdin to a Slack channel
(available on [Github][]):

    #>> /usr/local/bin/systemd-slack

    #!/bin/bash
    source /etc/default/slacksink
    UNIT=$1
    HOST=`hostname`
    MESSAGE="$UNIT failed on $HOST"
    SLACK_USERNAME="systemd-timer"
    systemctl status --full "$UNIT" | /usr/local/bin/slacksink \
    --channel="#team-channel" --message="$MESSAGE" --attachment \
    --color=danger

That's it 😊

Obviously, Slack is only one example here but it's IMHO a nice use of the
OnFailure handler. According to the [docs][] you can even list multiple services
here. Infinite fun with failing services 😉

[github]: https://github.com/zerok/slacksink
[docs]: https://www.freedesktop.org/software/systemd/man/systemd.unit.html#OnFailure=
[nl]: http://northernlightlabs.se/systemd.status.mail.on.unit.failure
