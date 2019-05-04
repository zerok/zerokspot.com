---
title: "Custom Repositories in MacPorts"
date: 2019-05-04T10:32:46+02:00
tags:
- sysutils
- macports
- homebrew
---

In order to distribute custom tools among friends but also to some degree at work I've created a [custom tap](https://docs.brew.sh/Taps) in [Homebrew](https://brew.sh). This is pretty straight forward as it only requires a Git repository where the audience can access it. Does [MacPorts](https://www.macports.org) have something similar?

While MacPorts focuses heavily on the main port tree you can specify multiple sources for that tree. You do that by editing `/opt/local/etc/macports/sources.conf`:

```
rsync://rsync.macports.org/macports/release/tarballs/ports.tar [default]
```

Any custom source should be added before the "default" entry in there as MacPorts will go through each source in the order they are specified here. In order to demonstrate this feature, I've created a custom Portfile for [exa][] which we are now going to distribute through a custom source.

## Signed archives

While MacPorts doesn't strictly require the ports tree to be distributed as archive, this is the only way where it also enforces the use of signatures. Since that road is only a bit longer, let's take that extra step here.

In order to sign anything, we first need a key-pair (consisting of a private and public key).

```
# Generate the private key
$ openssl genrsa -des3 -out privkey.pem 2048

# Derive the public key
$ openssl rsa -in privkey.pem -pubout -out pubkey.pem
```

I've then placed the public key file (`pubkey.pem`) inside `/Users/zerok/src/github.com/zerok/macports-ports/`. Next, MacPorts has to know that the public key can be used to verify artefacts. This is done by adding it to the `/opt/local/etc/macports/pubkeys.conf`:

```
# MacPorts system-wide public key configuration file.

/opt/local/share/macports/macports-pubkey.pem
/Users/zerok/src/github.com/zerok/macports-ports/pubkey.pem
```

## A new source

Now that MacPorts knows how it can verify our new source, we need to create an archive to distribute. For my exa Portfile I've created the following repository: <https://github.com/zerok/macports-ports>

This contains the following files:

```
$ exa -l --tree
drwxr-xr-x    - zerok  4 May 10:05 .
.rw-r--r--  562 zerok  4 May  9:59 ├── Makefile
.rw-r--r-- 1.8k zerok  4 May  9:32 ├── privkey.pem
.rw-r--r--  451 zerok  4 May  9:32 ├── pubkey.pem
.rw-r--r--  566 zerok  4 May 10:01 ├── pubkey.pem.sig
.rw-r--r--  316 zerok  4 May 10:01 ├── README.md
drwxr-xr-x    - zerok  4 May  9:26 └── sysutils
drwxr-xr-x    - zerok  4 May  9:26    └── exa
.rw-r--r-- 1.2k zerok  4 May  9:26       └── Portfile
```

The Makefile is mostly there to automate the creation of the ports-tree archive and its signing:

```
all: dist/ports.tar.bz2 dist/ports.tar.bz2.rmd160

dist:
	mkdir -p dist

dist/ports.tar.bz2: dist
	mkdir dist/ports && \
	cp -R sysutils dist/ports/ && \
	cd dist && tar -cjvf ports.tar.bz2 ports && \
	rm -rf ports

dist/ports.tar.bz2.rmd160:
	openssl dgst -ripemd160 -sign privkey.pem -out dist/ports.tar.bz2.rmd160 dist/ports.tar.bz2.rmd160

clean:
	rm -rf dist

release:
	rsync -avz dist/* zerokspot.com:/srv/www/h10n.me/www/htdocs/macports-ports/ && \
	rsync -avz pubkey.* zerokspot.com:/srv/www/h10n.me/www/htdocs/macports-ports/

.PHONY: all clean release
```

Running `make clean all release` now generates an archive, signs it, and uploads it including the public key (and its signature) to `https://h10n.me/macports-ports`.

Hurray, we now have a ports archive that we can finally use. Finally, update the sources.conf inside `/opt/local/etc/macports/` to look like this:

```
https://h10n.me/macports-ports/ports.tar.bz2
rsync://rsync.macports.org/macports/release/tarballs/ports.tar [default]
```

This tells MacPorts that it should look first inside the h10n.me source for a port before falling back to the default one.

```
# Download the sources:
$ sudo port sync

# Show the details of the new exa port:
$ port info exa
exa @0.8.0 (sysutils)

Description:          exa is a replacement for ls with additional features like color-coding for file types, Git integration, and extended attributes.
Homepage:             https://github.com/ogham/exa

Extract Dependencies: unzip
Platforms:            darwin
License:              MIT
Maintainers:          Email: zerok@macports.org
```

## The good and the not so good

While MacPorts approach for offering multiple sources of packages is a bit more complicated to set up than Homebrew's, I really like that archives *have to be signed* (unless you just expose a folder somewhere). Personally, I would have preferred GnuPG signatures over openssl/rmd160 simply because they are a bit more modern *and* include more meta-data about the owner of the signature. At least for the distribution of the source's public key I'd still opt to sign that with GnuPG which is why there is also a `pubkey.pem.sig` in my example repository.

I'm relatively new to MacPorts so I don't know if what I've described here even remotely follows best practices. You've been warned 😅

[exa]: https://the.exa.website/

