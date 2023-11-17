---
date: "2023-11-17T16:16:45+01:00"
tags:
- git
- development
- til
title: Signing Git commits with your SSH key
incoming:
- url: https://chaos.social/@zerok/111426681483108534
---

When someone commits code (or anything) into a Git repository, there is a simple way to identify that author: An e-mail address. The problem with this property is that it can be freely chosen by the author/committer. I could create commits with the e-mail address of any of my co-workers or Tim Cook and the system wouldn't complain. For anybody who checks the history of a file it would appear as if Apple's CEO had all of a sudden decided that writing Go was his new passion.

This is where "commit signing" comes in. This feature allows authors to "cryptographically sign" their commits, making it possible for others to verify that the author of a commit is really the expected person.

Historically, there was one way to do that: Using a tool called GnuPG which can also be used for signing (and encrypting) e-mails and other data but that has a (mostly deserved) reputation of being extremely hard to use and complicated.

Luckily, since [some time in 2021](https://github.com/git/git/blob/master/Documentation/RelNotes/2.34.0.txt) you can now also sign Git commits using your SSH key (and also x.509 certificates, but that's a different can of worms). This is awesome simply because most people who interact with Git already have an SSH key in order to push their changes up to GitHub/GitLab/etc. 

## Simple setup

How to generate a singing key:

```
❯ ssh-keygen -t ed25519 -f $HOME/.ssh/signing-key

# Also add it to your ssh-agent for easier retrieval by the Git CLI:
❯ ssh-add $HOME/.ssh/signing-key
```

Now add that file to your Git configuration and enable SSH for signing:

```
❯ git config --global commit.gpgsign true
❯ git config --global gpg.format ssh
❯ git config --global user.signingkey "$(cat $HOME/.ssh/signing-key.pub)"
```

With this in place, whenever you do a commit, it will be signed with your new key.

## Verifying commits

Now that you have a couple of signed commits, how can you verify them? 

```
commit c56cca04c8380472c9f03381211e8f9518d1a559
Author: Tim Cook <tim@apple.com>
Date:   Fri Nov 17 14:37:47 2023 +0100

    Initial commit

```

I have this commit history inside a sample repository, but did Apple's CEO really make this commit?

```
❯ git verify-commit 73d8a19
error: gpg.ssh.allowedSignersFile needs to be configured and exist for ssh signature verification
```

And this is the core of verifying SSH-signed commits: You need a file that contains a list of public key files with some additional properties including the the e-mail addresses used in the commits. The file format is documented in the [man page](https://man7.org/linux/man-pages/man1/ssh-keygen.1.html#ALLOWED_SIGNERS) but is relatively simple:

```
zerok@zerokspot.com namespaces="git" ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIDfVABUcDd6H1O8a+niNTbcSabPcfRob4WvhVBztJGm+
```

For demo purposes, let's create such a file inside the Git repository itself and then instruct Git to use it:

```
❯ git config --local gpg.ssh.allowedSignersFile "${PWD}/allowed_signers"

# And let's verify the commit
❯ git verify-commit c56cca0
Good "git" signature for zerok@zerokspot.com with ED25519 key SHA256:LexzzNw7jvDW7/d5wi+22vSywe169EjzWC33jPnICvk

# Or also inside the log listing:
❯ git log --show-signature | pbcopy
commit c56cca04c8380472c9f03381211e8f9518d1a559
Good "git" signature for zerok@zerokspot.com with ED25519 key SHA256:LexzzNw7jvDW7/d5wi+22vSywe169EjzWC33jPnICvk
Author: Tim Cook <tim@apple.com>
Date:   Fri Nov 17 14:37:47 2023 +0100

    Initial commit

```

Huh... so the commit was actually signed by me and not Tim Cook?! No way! But the commit was at least done by someone who should have access to this repository!

But as you saw, `git log` and `git verify-commit` still indicated that the signature was "good" and didn't report an error. This is where tooling like GitHub but also others might help you discover the mismatch. In the case of [GitHub](https://docs.github.com/en/authentication/managing-commit-signature-verification/about-commit-signature-verification#ssh-commit-signature-verification), there is some verification in place that will show this commit as "unverified" as Mr. Cook doesn't seem to be on GitHub. In general, though: Don't trust the author/committer email address. The mapping mapping inside the allowed-signers file should be more trustworthy but possibly some kind of combination would be best. Single GitHub at least does some e-mail verification, this might be a good first step.

Outside of GitHub (or in some setups, the management of that allowed-signers file is crucial. The [Git manual](https://github.com/git/git/blob/facca53ac3c2e8a5e2a4fe54c9c15de656c72de1/Documentation/config/gpg.txt#L42-L76) has some hints there:

> This file can be set to a location outside of the repository and every developer maintains their own trust store. A central repository server could generate this file automatically from ssh keys with push access to verify the code against.  
> In a corporate setting this file is probably generated at a global location from automation that already handles developer ssh keys.
> 
> A repository that only allows signed commits can store the file in the repository itself using a path relative to the top-level of the working tree. This way only committers with an already valid key can add or change keys in the keyring.

Especially for teams working together on more than one project it might probably make a lot of sense to have a shared allowed-signers file in its own repository and then have a CI job verify that commits to a pull-request/merge-request are signed with keys mentioned in that file.

## Key revocation?

What should happen if a key expires or is otherwise compromised? Looks like there are two options:

1. You can either add it to yet another file configured through `gpg.ssh.revocationFile`
2. Or add life-time properties to the allowed-signers file

Adding the key to the revocation file will make verification fail also on historical commits while with the [new `valid-before` and `valid-after` fields](https://man.openbsd.org/ssh-keygen.1#valid-after) in the allowed-signers file, you should be able to more granularly handle revocations. Locally, I couldn't make this work but that is probably just a typo somewhere in my config which I haven't tracked down yet.

## Does it work with 1Password?

As I now have a few SSH keys managed within 1Password, I also wanted to see if I could use those to sign commits. Turns out, I can and it's even [documented](https://developer.1password.com/docs/ssh/git-commit-signing/)!

## Verdict so far

I think making signing commits more accessible is an important goal and not having to use GnuPG for that definitely helps! As I already hinted at in the intro, it's dangerous to not reliably know who actually authored a code change no matter how that commit then made its way into a repository.

I'm not yet sure how to handle the allowed-signers file in the long term, though. For my personal setup I have a lot of options including fetching keys from some known people from GitHub or curating that file manually. No idea yet 😄

Thanks to [Caleb Hearth](https://calebhearth.com/sign-git-with-ssh) for his article about the same topic that made me dive into SSH-signing my commits ❤️