---
title: Ansible but with pull
date: "2021-01-13T18:44:12+01:00"
tags:
- 100daystooffload
- ansible
- ops
---

I really like configuring servers with Ansible. It's just explicit and magic-free enough so that I don't have to relearn it every couple of months. It's still not really a magic wand that does the whole server maintenance for you. You still have to somehow trigger that playbook somehow. A couple of months ago, I stumbled upon [this tweet by Simon Willison](https://twitter.com/simonw/status/1334238720187199488) that might make things easier, though.

Normally, Ansible operates in a push mode. This means that you as the server admin execute a playbook and Ansible will go out to all the hosts mentioned in this playbook to do your bidding. With [ansible-pull](https://docs.ansible.com/ansible/latest/cli/ansible-pull.html), on the other hand, you can turn this principle on its head by no longer pushing state onto machines but letting them pull state from a central server and updating themselves.

If you're thinking that this is basically just doing a git-pull and then ansible-playbook, you're close to the truth. ansible-pull just does a bit of wrapping around these two steps in order to make the experience a bit more reliable.

	$ ansible-pull -U git@github.com/zerok/infra-repo.git site.yaml

OK, so if you now have a single repository for multiple hosts in your infrastructure and you're using Ansible Vault to encrypt things like access keys and passwords, how can you keep host A knowing about the access keys for host B? Turns out, a Vault password doesn't need to be able to decrypt every encrypted item inside a repository but just those that are needed while applying the selected playbook. So in the end you just generate multiple passwords and inject them using different `--vault-id`s so that, for example, each host has its own password and there is another one for the shared data. Problem solved 🙂

But back to the start: why would you want to inverse the process here? A simple example would be that you often provision new machines with the same base image. With push you'd then have to trigger the playbook from somewhere else for that new host. With pull the host itself would just go to your central repo and execute that playbook for itself.

There seems to be no one-way for how you'd need to structure your repository to work with this. For my home network I have a repository that can be used for push *and* pull. All I do is setup the `/etc/ansible/hosts` inventory on each host and then pass an additional `-l hostname` flag to the `ansible-pull` call you saw above.

By default, `ansible-pull` will, if you don't explicitly specify a playbook, run a playbook based on the fully qualified hostname of the machine it is executed on. `hosta.network`would first try `hosta.network.yml`, then `hosta.yml`, and eventually fall back to a `local.yml`.  This is something I haven't done yet as I have to specify the `-l hostname` flag anyway for now since the repo should still also work for the push-scenario.

This whole setup also, naturally, has its downsides like that synchronizing updates over multiple hosts gets more complicated but for my simply home-network scenario this should do fine. Let's see if I'll still with it 🙂
