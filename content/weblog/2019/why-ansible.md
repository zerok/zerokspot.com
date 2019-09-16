---
title: "Why Ansible?"
date: 2019-09-16T09:52:35+02:00
tags:
- ansible
- ops
---

Over the last couple of months I've been slowly migrating all server
configuration to [Ansible][]. In some environments I wasn't using
anything similar, in some I migrated some SaltStack configurations. In
this post I just want to write a little bit about why I think working
with Ansible is so pleasant for me compared to other setups.

## Simple host inventory

Ansible has a single place where you define hosts: the inventory. In
the inventory you can define groups of hosts (e.g. jenkins-workers)
and you can override settings in there either per host or per
group. Want all jenkins-workers to install Docker CE 19.3 instead of
18.9? You’d define that override in the inventory 🙂


```
[jenkins-workers]
worker01.test.example.com
worker02.test.example.com

[jenkins-workers:vars]
docker_version=19.3

[jenkins-masters]
master01.test.example.com
```

You can, if you want, extend that simple inventory with dynamic parts
in order to, for instance, better integrate it with cloud
infrastructure, but the basics stay the same.


## Playbooks

Playbooks allow you to combine tasks that should be executed in order
to solve a particular problem. Let’s say you want to prepare a server
to become a worker node in a Jenkins cluster.

The playbook would then do at least the following steps:

1. Prepare the node
  1. install the JVM on that node
  2. Create the user the Jenkins agent should be run as
  3. Create the workspace folder for the agent
  4. Install the agent
  5. Define a systemd service to ensure that the agent is launched at
     boot time and restarts if necessary
2. Register the node on the master
3. Setup monitoring
  1. Create a scraper in Prometheus to fetch metrics from that new node
  2. Create an alerting rule to get notified of the jenkins-agent
     service is no longer running

In the simplest case, a playbook would just define a couple of tasks
that should be executed on a single host:

```
---
- hosts: jenkins-node-123
  tasks:
  - name: Create user
    user:
      name: jenkins-worker
      state: present
  - name: Create workspace
    file:
      path: /srv/jenkins-worker
      state: directory
      owner: jenkins-worker
```

In our case, though, the playbook actually needs to interact with 3
hosts:

1. the new node
2. Jenkins master
3. Prometheus 

So we just have to do three of these sections and we're done. 

```
---
- hosts: jenkins-node-123
  tasks:
    # ...
- hosts: jenkins-master
  tasks:
    # ...
- hosts: prometheus
  tasks:
    # ...
```

What I really enjoy about playbooks is that you can re-use all your
knowledge you gathered there for more advanced concepts like [roles][].

There is also no complicated eventing system where you end up in
situations where you're unsure what step would be executed at what
time. If you want to conditionally execute some step, you [register a
variable][var] and add a `when` property to the task that should be
conditionally executed. If you want to execute something at the end of
the whole playbook if a certain task has been executed, notify a
[handler][]. That's it.

With playbooks and inventories you usually already get quite
far. If you need more, there are some concepts based on these two
(like being able to import playbooks, and [roles][]) that should be be
able to solve most relevant use-cases.

## Server-less

Unlike with Salt or other systems you don't need a server component
installed on the hosts you want to manage with Ansible. As long as
[Python][py] is available there, you only have to add a host to your
inventory and you're good to go. This, combined with playbooks, have
made me convert even simple deployment scripts to playbooks for more
consistency and readability!

## From simple to slightly less so

All these points have one thing in common: Ansible allows you to work
with relatively simple building blocks. Based on them you can manage
highly complex systems but the configuration always (until you try
really hard) stays something you can look at and grasp within a couple
of minutes.

## Nothing is perfect

Sure, Ansible is not perfect. Over the months I've run into a lot of
situations where something didn't completely work as expected or I had
to create a little workaround which had a detrimental effect on
readability. The community around Ansible is large and vibrant enough,
though, that finding such workarounds is usually not an issue.

With the advent, though, of tools like [Terraform][tf] it becomes
harder and harder to draw a line where you should use Terraform and
where to use Ansible. Personally, I provision hardware using
Terraform, install the base operating system and then do the
configuration of that operating system with Ansible. I think this
combination works rather well so far 🙂


[var]: https://docs.ansible.com/ansible/2.5/user_guide/playbooks_conditionals.html#register-variables
[handler]: https://docs.ansible.com/ansible/latest/user_guide/playbooks_intro.html#handlers-running-operations-on-change
[roles]: https://docs.ansible.com/ansible/latest/user_guide/playbooks_reuse_roles.html
[ansible]: https://www.ansible.com/
[tf]: https://www.terraform.io/
[py]: https://www.python.org/
