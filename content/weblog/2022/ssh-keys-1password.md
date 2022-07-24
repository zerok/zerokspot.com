---
title: SSH key management with 1Password
date: "2022-07-24T11:53:45+02:00"
tags:
- ssh
- security
- 1password
incoming:
- url: https://chaos.social/@zerok/108701722694651457
- url: https://twitter.com/zerok/status/1551145191628804096
---

Triggered by [this post from the official 1Password blog by K.J. Valencik](https://blog.1password.com/1password-ssh-changed-how-i-work/) I’ve been using a new feature of the product for the last couple of weeks: [SSH key management](https://developer.1password.com/docs/ssh/get-started). The feature consists of two parts: First of all you can now _store SSH keys within 1Password_ either by just importing an existing private key file or by letting the application generate a RSA or Ed25519 key for you. As usual, you can then add notes, extra fields etc. to that new key.

The second and arguably more interesting part of the feature is a _custom SSH agent_ that will use these keys directly from 1Password. For this to work you have to enable the “Use the SSH agent” setting and then refer to its sock file either with the `IdentityAgent` setting in your SSH config or the `SSH_AUTH_SOCK` environment variable. In the [official docs](https://developer.1password.com/docs/ssh/get-started#step-3-turn-on-the-1password-ssh-agent) you can find a more detailed guide.

Once configured, all SSH keys stored in your “Personal” or “Private” vault can be used automatically. For each new usage-scenario (e.g. connecting a server A in a new terminal) you will get queried for configuration which can happen using Touch ID, your password, or your Apple Watch (or other setups if you’re on a platform other than macOS). 

In my personal setup I’ve for now created a separate SSH key for pretty much every server-group that I’m working with and also different keys for GitLab, GitHub and so on. This is made easy thanks to the 1Password SSH agent supporting [hints what key you’d like to use](https://developer.1password.com/docs/ssh/agent/advanced#match-key-with-host) for which connection: Basically, export your SSH key’s public component into a file and use the `IdentityFile` setting in your SSH config to point to that file. The agent will then use the identity associated with that public key while keeping your private key stored inside the vault! This is also great if you’re working with, e.g., multiple GitLab accounts (for work, personal stuff, projects, …).

In general I’ve been really positively surprised by this feature! It has made my use of SSH keys so much more traceable for me and I just love that I can allow the use of a key using a click on my watch instead of having to navigate through a couple tea cups before reaching the Touch ID sensor of my laptop or the Yubikey attached to the same laptop🤪

The only complaint that I have right now is the limitation that keys *have* to be stored in a vault named “Personal” or “Private”. I wanted to use this feature also from project-specific vaults in order to keep project-stuff in their own vaults. Other than that, though, I love it 😍
