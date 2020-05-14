---
title: "TIL: SSH key comment determines login in GCP"
date: "2020-04-23T10:06:00+02:00"
tags:
- til
- gcp
- ssh
---

When you create a new VM in [GCP][] and add a SSH key to it then the VM will
create a new login based on the comment of that SSH key. For instance...

```
ssh-rsa LOTSOFBASE64 username
```

The public key shown above would result in a new login "username" on the VM.
While this is documented in the [compute docs][cd] it surprised me since Azure
explicitly asks for the username that should be used when creating an
administrative account.

[gcp]: https://cloud.google.com/
[cd]: https://cloud.google.com/compute/docs/instances/adding-removing-ssh-keys?hl=en_US
