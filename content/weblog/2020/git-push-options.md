---
title: Git push options
date: "2020-09-28T21:05:01+02:00"
tags:
- git
- til
---

While working on a CI workflow last week I learnt about something I hadn’t known before: [Push options](https://git-scm.com/docs/git-push#Documentation/git-push.txt---push-optionltoptiongt). Using the  flag `-o=<string>` or `--push-option=<string>` you can pass an additional string to the pre-receive/post-receive on the repository receiving the push. Why would you want to do that? Well, [GitLab](https://docs.gitlab.com/ee/user/project/push_options.html) for instance allows you, for instance, inject additional environment variables into a CI pipeline or skip the pipeline altogether:

	# Do not trigger CI
	$ git push --push-option=ci.skip 
	
	# Inject another environment variable into CI
	$ git push --push-option=ci.variable="name=value"

 If you want to provide more than a single option then just use the flag multiple times.

How does that work under the hood? To find out I’ve created a little bare repository (`repo.git`) with just a single commit in it to which I then added the following pre-receive hook script:

	#!/bin/bash
	echo "BEGIN ENV"
	env | grep GIT_PUSH
	echo "END ENV"
	exit 1

If I then run `git push origin master --push-option hello_world` on it, the receive hook will print the following output before rejecting the push:

	remote: BEGIN ENV
	remote: GIT_PUSH_OPTION_0=hello_world
	remote: GIT_PUSH_OPTION_COUNT=1
	remote: END ENV

As you can see from that you can access the push options from `$GIT_PUSH_OPTION_n` environment variables with `n` being a zero-indexed counter with the maximum value being `$GIT_PUSH_OPTION_COUNT - 1 `. That’s pretty much it. The rest is completely up to the hook script.

## fatal: the receiving end does not support push options

Initially I got the error shown above which means that the repository I wanted to push to didn’t advertise its support for push options. That was easily fixed with the following commands:

	$ cd repo.git
	$ git config receive.advertisePushOptions true

## Playground setup

In case you want to play around with that feature, I’ve prepared a little setup script that automates everything I’ve written above:

	#!/bin/bash
	set -e
	
	# Clean slate
	rm -rf repo
	rm -rf repo.git
	rm -rf clone
	
	# Create simple repository
	mkdir repo
	cd repo
	git init
	touch README.md
	git add README.md
	git commit --no-gpg-sign -m "Initial commit"
	cd ..
	
	# Make the repo a bare one so that we can directly push to it:
	git clone repo repo.git --bare
	rm -rf repo
	
	# Let's create a simple pre-receive hook
	cp pre-receive.sh repo.git/hooks/pre-receive
	chmod +x repo.git/hooks/pre-receive
	cd repo.git
	git config receive.advertisePushOptions true
	cd -
	
	# Now let's clone that repository
	git clone repo clone
	cd clone
	touch otherfile
	git add otherfile
	git commit --no-gpg-sign -m "otherfile added"
	git push origin master -o "hello_world"
