---
title: Development containers with Visual Studio Code
date: "2020-07-15T16:28:41+02:00"
tags:
- 100daystooffload
- vscode
- docker
- development
---

I’m very late to the party here but while trying to set up a coding environment on my gaming PC I remembered a recent episode of [The Changelog in which Jessica Kerr](https://changelog.com/podcast/398) talked about setting up sharable development environments using Docker containers inside of Visual Studio Code.

Before remembering that show I had already planned to create a little Docker image that would bundle all the tools I’d need for doing Go development in order to do just that on one of my remote VMs while my laptop was getting repaired. If I could all of a sudden use that image also within VSCode, it would be awesome!

The result of this little experiment is [available on GitHub](https://github.com/zerok/dev-environments).

## What do you need?

Let’s work on a little example here: I want to work on zerokspot.com, which is hosted on GitHub, inside of Visual Studio Code on Windows (with WSL2 and Ubuntu already installed). For this, I’ve installed the following things:

- Docker Desktop for Windows
- Visual Studio Code
- GitHub Desktop (mostly because I’m really lazy)

Inside VSCode I’ve installed the following extensions:

- Docker
- Remote - Containers

## Open the project

Next, I need to check out the latest code for zerokspot.com from GitHub. I do that using GitHub Desktop. Once everything is there, I open the new folder inside VSCode just to be greeted with following message:

<figure><img src="/media/2020/devcontainer-notification.png"><figcaption>VSCode detects existing dev-container configuration</figcaption></figure>

VSCode noticed that the project contains a `.devcontainer/devcontainer.json` file. So what’s in there? 

	{
		"name": "dev-environment-base",
		"context": "..",
		"dockerFile": "..\\Dockerfile.devcontainer",
		"settings": { 
			"terminal.integrated.shell.linux": null
		}
	}

Turns out, not that much! It’s basically just a reference to a Dockerfile that should be used to set up a development environment.

Back in VSCode, I click on the “Reopen in Container” button and a new application window is shown with (at first) only one difference: In the lower left corner I now see a green badge with the label “Dev Container: dev-environment-base”. This tells me that VSCode is now connected to the container that was mentioned inside the `.devcontainer/devcontainer.json` file.

Let’s now open a terminal (Terminal -\> New Terminal) and take a look around. The first thing you’ll notice is that it starts in a folder called `/workspaces/zerokspot.com` and that the shell indicates that I’m running all that as `root`. Since the OS that VSCode is running on right now is Windows 10, this indicates that I’m already inside a container.

And that container comes with everything I need!

	$ node --version
	v12.18.2
	
	$ python3 --version
	Python 3.8.2
	
	$ go version
	go version go1.14.5 linux/amd64

This also means that if an extension requires external tools, these can also be pre-installed using this environment. I did so, for instance, for the Go extension and goimports and gopls in particular.

## Exposing ports

For zerokspot.com I usually want to test changes before commiting them. On my usual dev machine I’d therefore just run ` hugo serve -D` and then just look at http://localhost:1313. 

This also works inside the devcontainer! Inside the terminal I just start hugo and then expose the port using the “Forwarded ports” feature:

<figure><img src="/media/2020/devcontainer-forwardedports.png"><figcaption>Managing forwarded ports</figcaption></figure>

A quick click on the “1313” entry and I’m done  😍

## How does all of this work?

In the background, VSCode starts a Docker container and installs a vscode-server into it. This server then handles all the tasks that would normally be the operating system’s job.

At this point, such a  server-setup is available for Docker containers, SSH, and WSL. Given my usual environment the latter two won’t see all that much use from me but the Docker-way is just wonderful! I’ve probably had far too much fun setting all that up today 😁

I still do all the Git operations from outside this setup simply because I haven’t yet set up the credentials forwarding properly. Let’s see if my laptop is back before I do that 😅
