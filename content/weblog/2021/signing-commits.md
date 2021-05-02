---
title: Signing commits (also with multiple GitLab accounts)
date: "2021-05-02T20:40:27+02:00"
tags:
- 100daystooffload
- git
- gitlab
incoming:
- url: https://chaos.social/@zerok/106167416815152723
---

When you look through a commit log in Git, you always see the name and the e-mail address of the author/committer. Sadly, as a consumer of such changelogs you cannot be sure that the name and address are valid or even be owned by the actual author of a commit. Nothing can stop me from doing a commit with the following command:

	$ git commit . --author "John Doe <john.doe@example.org>"

When I now run `git log` on that repo, I see that this commit was made by John Doe:

	commit ffe6bff05168a56d97e81de4ffe62cfb834f0a25
	Author: John Doe <john.doe@example.org>
	Date:   Sun May 2 18:12:01 2021 +0200
	
	    test
	

But hey! This commit was not made by John Doe! And that’s pretty much the same situation if I had used my real name and real email address. How does someone know based on the changelog that a commit was really made by the person that shows up there?

This is where “signed commits” come in.

## How to sign a commit

When you combine the power of [GnuPG](https://gnupg.org/) with Git you can create “signed commits”. Here the commit itself is signed with your private key so that people looking at the changelog can see that the owner of the key used there has created that commit.

I won’t go into details on how to create a GnuPG here as there are more than enough guides online (or even the [official manual](https://www.gnupg.org/gph/en/manual/c14.html)). 

Once you have a key that is associated with the e-mail address you’ve also configured within your Git configuration, you can sign a commit with that key using the `-s` or `--signoff` flag:

	$ git commit --signoff

If you have such commits, the signatures will show up within the changelog if you also set the `--show-signature` flag for the log-command:

	$ git log --show-signature
	commit 3436527125576685326c5af8a80108358858e315 (HEAD -> main)
	gpg: Signature made Sun May  2 20:04:06 2021 CEST
	gpg:                using RSA key C4E02BA840483A5D6B7616076F203F0D220F8E98
	gpg: Good signature from "Horst Gutmann <horst@zerokspot.com>" [ultimate]
	gpg:                 aka "Horst Gutmann <h.gutmann@netconomy.net>" [ultimate]
	gpg:                 aka "Horst Gutmann <zerok@zerokspot.com>" [ultimate]
	gpg:                 aka "Horst Gutmann <horst.gutmann@gmail.com>" [ultimate]
	Author: Horst Gutmann <zerok@zerokspot.com>
	Date:   Sun May 2 20:03:51 2021 +0200
	
	    Signed commit
	

## Verified commits on GitLab

If I now host the project of GitLab those commits will show up as “unverified”:

<figure><img src="/media/2021/gitlab-unverified-commits.png"><figcaption></figcaption></figure>

To change that, I have to do two things:

1. Add my public key to the [“User Settings → GPG Keys”](https://gitlab.com/-/profile/gpg_keys).
2. Associate the e-mail address used for the commit (and that’s a UID of the GnuPG key) with the GitLab account in the [“User Settings → Emails”](https://gitlab.com/-/profile/emails).

Note that this might not  change the status of past commits. Once it works, you will see a green “Verified” status with the following info:

> This commit was signed with a verified signature and the committer email is verified to belong to the same user. 

## Working with multiple accounts

This whole setup gets a bit more complicated if you have multiple accounts (e.g. one for work, one for personal projects). Public keys have to be unique across a GitLab instance. 

So, if you want to have signed commits for multiple accounts, you will also have to create a separate key-pair.

Once you have that, you will have to tell Git to use that key for further commits. In case you, for instance, have multiple keys associated with a specific e-mail address, then you can use the `user.signingKey` setting to specify the key explicitly:

	$ git config --local user.signingKey $KEYID

That’s pretty much it 🙂
