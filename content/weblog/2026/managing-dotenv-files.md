---
date: "2026-01-24T19:55:00-08:00"
incoming:
- url: https://chaos.social/@zerok/115953753470427177
tags:
- development
- security
title: Managing .env files
---

Injecting secrets into applications usually happens through environment variables. For local development it has become common to use `.env` files (or `.envrc` files if you’re using [direnv](https://direnv.net/)) to manage those.

While having all your environment variables defined inside a text file is simple to set up, it comes with some well documented downsides; first and foremost that secrets in there are not encrypted and can accidentally be committed to a versioning system.

## 1Password’s solution

A couple of months ago [1Password](https://1password.com/blog/1password-environments-env-files-public-beta) announced that they’d roll out their own solution for this problem. When you enable the “Developer experience” in your settings, you can now define so-called “Environments”. These are associated with your account (or team) and contain a set of key-value pairs. These are just secrets inside your Vault.

You can then also define a “Destination” for such an environment. This can be either a local `.env` file or an AWS Secret Manager instance.

OK, so 1Password allows you to automatically write a set of key-values pair into a `.env` file. What’s does that solve exactly?

## Not a normal file

What 1Password mounts into your filesystem is actually not a normal file but a Unix-named pipe! When applications read from that pipe you get a popup asking if that’s ok for you (similar to when you use 1Password’s SSH key manager). The file itself never contains those secrets, though. They are just passed through the pipe after your confirmation.

Additionally, tools like Git don’t commit named pipes. `git add .env` is pretty much a no-op.

You can learn more in the [1Password docs](https://developer.1password.com/docs/environments/local-env-file/).

## Sharing with your team

Since the “Environments” are tied to your 1Password account, they can be shared with others. This is especially handy for small teams where everyone should quickly get a setup with a test system without just checking stuff into Git.

## Snowflake

Unlike SSH keys in 1Password, environments still feel like a little snowflake right now. They are hidden behind their own section in the UI and also in other aspects feel like something completely different from your normal secrets.

I hope this will eventually change once they move out of beta but I like this feature, otherwise! Previously, I used a local Hashicorp Vault server in combination with direnv to load secrets without putting them into plain text files, but this approach looks even nicer to me!
	
