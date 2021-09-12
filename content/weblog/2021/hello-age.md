---
title: "Hello, age"
date: "2021-09-12T19:20:00+02:00"
tags:
- security
- dataprivacy
- tooling
- blogpost
- draft
---

Getting files securely from one person is far more complicated than it should
be esp. in the age of everything behing in some cloud or another. Let's say
you're collaborating with your co-workers via a platform like Slack or
file-shares provided by Dropbox. If you now put a file (or anything for that
matter) into a Slack channel or a Dropbox share, employees of the company
behind that respective service can look at that data. Usually they don't simply
because they have better things to do, but at least some automation more often
than not does look at your data to provide you things like a preview or other
"useful" features.

If the data that you are sharing is for instance a private key for a
certificate, a private SSH key, or a password file, you need some kind of
end-to-end encryption so that really only you (the sender) and the person you
want to have that piece of data (the recipient) can actually look at it.

A simple solution for this issue that is quite common is to just zip that data
and then put a password on it. But how do you then share that password with the
recipient? This is where [public-key cryptography][pk] comes in: With this the
sender and the recipient both have their own pair of keys: a private and a
public key.  The sender encrypts the data with the public key of the recipient
and the recipient can then decrypt that data with their private key.

For the longest time, I've used GnuPG (and other OpenPGP tools) for that. The
downside of GnuPG is that it's hard to use right. There are so many commandline
flags and environment variables that you usually need a couple of tries until
you get a proper key-pair setup, stuff encrypted, decrypted, or signed.

## Hello, age

[age][age] by [Filippo Valsorda][fv] and [Ben Cox][bc] does away with most of
these options and tries to offer a very simple utility that does just one thing
and that as straight-forward as possible: encrypting/decrypting files. There
are no configuration files and really just the bare-minimum of commandline
flags.

The first thing you'll need to do when using age is to create a new key-pair
for yourself. I like to keep at least my own keys organized in a single
directory. For this purpose I'm using `$HOME/.config/age` also within this
little tutorial:

```
$ age-keygen -o ~/.config/age/me.key.txt
Public key: age1tfgpdshzh5f0zgsts2l5f6s5rcumparfmxcgy4mr44p24dw3jd8s06t46k
```

As you saw, the command already prints your public key but fear not: The
generated file also contains it alongside with your new private key:

```
$ cat ~/.config/age/me.key.txt
# created: 2021-09-12T18:36:25+02:00
# public key: age1tfgpdshzh5f0zgsts2l5f6s5rcumparfmxcgy4mr44p24dw3jd8s06t46k
AGE-SECRET-KEY-<SUPER_SECRET_KEY>
```

With this information, someone can now go ahead and send me encrypted data by
just using my public key. Let's say, I'm now someone else and want to send old
me my nvim config:

```
$ age -r age1tfgpdshzh5f0zgsts2l5f6s5rcumparfmxcgy4mr44p24dw3jd8s06t46k \
    --encrypt --output init.lua.enc \
    $HOME/.config/nvim/init.lua
```

This will generate a new file in the current directory called `init.lua.enc`.
If old me now wants to decrypt that file, he can do so using the `--decrypt`
option of the `age` command:

```
$ age --decrypt --identity $HOME/.config/age/me.key.txt init.lua.enc
```

## Organizing recipients

Entering all these public keys manually gets tidious rather quickly. For this
reason age also supports that you reference a file containing all the
recipients that something should be encrypted for.

For the purpose of this example, I've created two more test-personas that
should now also receive my nvim configuration:

```
$ age-keygen -o ~/.config/age/test-persona-1.key.txt
Public key: age1d88qqgjv2lpsxc48g0uxgk5944d3dunp7uesvevt3pr29fazjgjsf3nvy8

$ age-keygen -o ~/.config/age/test-persona-2.key.txt
Public key: age1tcne3l4smq8ew88psalqdue09vm0ev96nxh8ahqkrpehxqnjq5cqlzpdkm
```

I'll now simply put their public keys into a single text file which I can then
later one reference:

```
$ echo age1d88qqgjv2lpsxc48g0uxgk5944d3dunp7uesvevt3pr29fazjgjsf3nvy8 > ~/.config/age/test.recipients.txt
$ echo age1tcne3l4smq8ew88psalqdue09vm0ev96nxh8ahqkrpehxqnjq5cqlzpdkm >> ~/.config/age/test.recipients.txt

$ age -R ~/.config/age/test.recipients.txt -o init.lua.multienc ~/.config/nvim/init.lua
```

Now my two test-persona can decrypt that file. This is especially useful when
you repeatedly want to share files with your whole team where everyone has
their own key-pair.

## Support for SSH keys

In some use-cases having to generate yet another key-pair is too much of a
pain. For this reason age also supports the use of SSH keys. Let's say that
someone wants to send me something encrypted but only knows my GitHub account:

```
# Encrypt by taking all my public keys from GitHub:
$ curl https://github.com/zerok.keys | age -R - -o init.lua.sshenc ~/.config/nvim/init.lua

# Decrypt using my private key:
$ age --decrypt --identity ~/.ssh/id_rsa init.lua.sshenc
```

For GitLab that's a bit more complicated but not much:

```
$ curl https://gitlab.com/api/v4/users/zerok/keys | jq -r '.[] | .key' | age -R - -o init.lua.glenc ~/.config/nvim/init.lua
```

## Conclusion

I've only just started messing around with age but so far I quite enjoy it! One
thing that it doesn't do, though, is offer any kind of signing support. This is
outside of its scope and the authors in their [initial spec][sp] suggested that
one should look at tools like [minisign][ms] and [signify][sf] for that.
Personally, I think I'll keep using GnuPG for that but if time permits I
absolutely also want to at least mess a little bit around with minisign 🙂

[fv]: https://twitter.com/FiloSottile
[bc]: https://twitter.com/Benjojo12
[age]: https://age-encryption.org
[pk]: https://en.wikipedia.org/wiki/Public-key_cryptography
[sp]: https://docs.google.com/document/d/11yHom20CrsuX8KQJXBBw04s80Unjv8zCg_A7sPAX_9Y/preview
[ms]: https://github.com/jedisct1/minisign
[sf]: https://github.com/aperezdc/signify

