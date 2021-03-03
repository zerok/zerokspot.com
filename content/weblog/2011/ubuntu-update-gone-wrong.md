---
date: '2011-02-27T12:00:00-00:00'
language: en
tags:
- linux
title: Ubuntu update gone wrong
---


Wow, I finally had my first update problem with Ubuntu yesterday :-) After
doing an update recently, all of a sudden I could no longer connect to my
server. Port 22 was for some reason not reachable. Right away, I booted into a
rescue system and went on checking the log files. At first I suspected there
might be a problem with some fail2ban rules but everything appeared to be fine
on that front.

Luckily, [carrot-server][1]'s support team had already someone with similar
issues and pointed me at [issue 634900][2] which describes a problem with the
new ssh upstart config in OpenVZ containers. So after removing the "oom never"
line from /etc/init/ssh.conf sshd finally started again.

-----------------------

But that was only the first problem. For some reason I couldn't yet log in
with anything other than my root account. Whenever I tried to connect with a
different user I got this:
    
    Cannot execute /bin/bash: Permission denied

Or a slight variation when using "su":
    
    Unable to cd to '/home/myusername'

Virtually every topic discussing a similar issue points at /bin having wrong
permissions, but those were fine. What was wrong in my case was not /bin but /
itself. So the fix was pretty simple:
    
    chmod a+rx /

and also:
    
    chmod a+rx /var

which was wrong too. No idea why, but finally I could log in again and
restart all necessary services :-)


[2]: https://bugs.launchpad.net/ubuntu/+source/upstart/+bug/634900
[1]: http://www.carrot-server.com
