---
title: Advanced features of Chezmoi
date: "2021-12-04T17:12:06+01:00"
tags:
- til
- chezmoi
---

I’ve been using [chezmoi](https://www.chezmoi.io/) for a very long time only with its most basic feature set: Basically managing specific files as they are and nothing more. While browsing through the docs I noticed that I’ve only scratched the surface so far of what the tool *actually* can do!

## Templates
The first thing I noticed is that chezmoi supports [Go templates](https://www.chezmoi.io/docs/templating/)! One example where this comes in handy is my `startship.toml` where I need to set the path to binary that may or may not be on the path. 

Here I can simply have something like this:

	format = """
	$username\
	$hostname\
	$directory\
	$git_branch\
	$git_status\
	${custom.kubeselect}\
	${python}\
	$line_break\
	$status\
	$shell\
	$character"""
	
	[gcloud]
	disabled = true
	
	[custom.kubeselect]
	when = "true"
	command = "{{ env "HOME" }}/bin/kubeselect status"
	style = "bold blue"
	format = "⛵️ [($output)]($style) "

Within the chezmoi-root this would be stored as `dot_config/starship.toml.tmpl`.

Chezmoi’s support for templates doesn’t end at simple environment variables. You can even fetch state from secret stores like Hashicorp’s Vault:

	some_secret = "{{ (vault "vault-key").data.data.value"

## Encrypted files
If you don’t want to treat only parts of a file a secret but actually the whole thing, you can also run the [whole file through either GnuPG or Age](https://www.chezmoi.io/docs/how-to/#encrypt-whole-files-with-age). For this you first have to specify what key should be used. Let’s say I want to have `$HOME/secret.txt` encrypted using Age and so I added the following entry to my configuration file:

	encryption = "age"
	[age]
	    identity = "/Users/zerok/.config/age/me.key.txt"
	    recipient = "age1tfgpdshzh5f0zgsts2l5f6s5rcumparfmxcgy4mr44p24dw3jd8s06t46k"
	

Now I can add an encrypted version of that file to chezmoi:

	chezmoi add --encrypt ~/secret.txt

The fill will now be stored as `.local/share/chezmoi/encrypted_secret.txt.age`. 

There are also ways to execute functions around the application of changes but I haven’t used them yet. Perhaps this will be something for another post 😅
