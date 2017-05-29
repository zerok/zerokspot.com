---
date: 2017-05-29T23:08:30+02:00
title: Backing up with restic
tags:
- restic
- backup
- tooling
- golang
---

[restic](https://restic.github.io/) is a tool with the goal of making it easy
for you to back up your files and directories in an encrypted fashion. Think of
Apple's TimeMachine but on a more granular level. But it not only encrypts the
backed-up data but also signs it so that any tampering can be detected. The
whole idea of restic is that the local system is trustworthy, the location where
the backups are stored is not, though. Something like Dropbox or Google Drive
comes to mind here.

And, indeed, restic is really easy to set up:

```
$ restic --repo ~/tmp/restic-backup init
enter password for new backend:
enter password again:
created restic backend 5fc9a7e3c1 at /Users/zerok/tmp/restic-backup

Please note that knowledge of your password is required to access
the repository. Losing your password means that your data is
irrecoverably lost.
```

Now that I have a repository in place, I will simply add some of my documents to
it:

```
$ cd ~/Documents/travels
$ restic --repo ~/tmp/restic-backup backup pycon2017
enter password for repository:
scan [/Users/zerok/Documents/travels/pycon2017]
scanned 1 directories, 6 files in 0:00
[0:00] 100.00%  0B/s  381.108 KiB / 381.108 KiB  7 / 7 items  0 errors  ETA 0:00
duration: 0:00, 13.82MiB/s
snapshot 1c2f0cb7 saved
```

With the `snapshots` command I can inspect all the save-points I have for my
data. Unsurprisingly, there is only one so far 😉

```
$ restic --repo ~/tmp/restic-backup snapshots
enter password for repository:
ID        Date                 Host        Tags        Directory
----------------------------------------------------------------------
1c2f0cb7  2017-05-29 22:50:23  mercury                 /Users/zerok/Documents/travels/pycon2017
```

Next, let's say I remove this folder by accident. I can then use the `restore`
command to get my data back given the snapshot ID:

```
$ rm -rf ~/Documents/travels/pycon2017
$ restic --repo ~/tmp/restic-backup restore 1c2f0cb7 --target ./
enter password for repository:
restoring <Snapshot 1c2f0cb7 of [/Users/zerok/Documents/travels/pycon2017] at 2017-05-29 22:50:23.933081627 +0200 CEST by zerok@mercury> to ./
```

Sadly, it looks like a I cannot simply restore a directory over its existing
state. Instead, I have to remove it first. For 99% of the use-cases this is
absolutely fine with me.

I can now create multiple snapshots of my data and also delete old snapshots if
I like. There is even a command that allows me to remove snapshots according to
a policy:

```
$ restic --repo ~/tmp/restic-backup forget --keep-weekly 2
```

The command above will keep for the last 2 weeks only the latest snapshot and
remove all older ones. There are flags for weeks, months, days, and even hours
or years. If I want to test-run your policy, I can do so using the `--dry-run`
flag with the `forget` command 😊 This will then print all the snapshots that
would be kept and those that would be removed.

One especially neat feature I want to mention before ending this post is that
repositories support more than just one password. I can add and remove passwords
basically at will using the `keys` command and its sub-commands `add`, `rm`, and
`passwd`:

```
$ restic --repo ~/tmp/restic-backup key list
enter password for repository:
 ID          User        Host        Created
----------------------------------------------------------------------
*e473e1c9    zerok       mercury     2017-05-29 19:59:48
```

To sum it up, restic looks like a really nice tool esp. when you're working with
storage systems you might not entirely trust. If you want to keep a backup of
something like your financial documents in something like Dropbox, this looks
like a really fine tool to evaluate 😊
