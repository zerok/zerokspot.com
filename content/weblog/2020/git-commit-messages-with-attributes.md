---
title: Git commit messages with attributes
date: "2020-10-24T10:25:10+02:00"
tags:
- git
- til
---

While looking through the [ConventionalCommit spec](https://www.conventionalcommits.org/en/v1.0.0/) I stumbled upon a feature of Git I somehow always thought was there but hadn’t realize was actually part of the Git toolchain itself: Trailers!

In Git, trailers are key-value pairs written inside of commit messages that, for instance, make things like references to JIRA tickets programmatically accessible. Let’s take a look at the following commit message:

	Awesome new feature
	
	Some details about this commit
	
	JIRA-Ref: PROJ-123

The last line is an example for a trailer. It consists of a token (`JIRA-Ref`), a separator (`:`), and a value (`PROJ-123`).

With the [`interpret-trailers`](https://git-scm.com/docs/git-interpret-trailers) the Git command offers a subcommand for extracting and manipulating these trailers. Using various configuration options, (among others) the following operations can be performed:

- Add a trailer if it is not present (using a shell command)
- Remove a trailer if present
- Modify a trailer if present (using a shell command)

A call to `git inspect-trailers` could, for instance, be integrated into a commit-msg hook script in order to do some modifications or using pre-commit to prevent commit messages not including, for instance, certain trailers or explicitly adding trailers with information from the context of a commit.

In general, you’ll probably most often need to implement a bit of scripting around trailers, the git subcommand just makes it a bit easier so that you don’t have to extract the trailers yourself:

	$ git config trailer.jiraref.key "JIRA-Ref"
	$ git config trailer.jiraref.ifexists replace
	$ git config trailer.jiraref.command "echo https://jira.company.com/secure/\$ARG"

Let’s now take the same commit message that we create above and interpret it. `git inspect-trailers` will then replace the JIRA-Ref line like this:

	$ cat message.txt | git interpret-trailers
	Awesome new feature
	
	Some details about this commit.
	
	JIRA-Ref: https://jira.company.com/secure/PROJ-123

## Trailers only

You can also use inspect-trailers to get just the token/value pairs:

	$ cat message.txt | git interpret-trailers --only-trailers
	JIRA-Ref: https://jira.company.com/secure/PROJ-123

I haven’t yet tried to script my way around trailers but I’d assume I’d do something like this if I need more complex logic for updating the trailers:

1. Extract just the trailers using the `--only-trailers` flag
2. Massage the values into the for that I want
3. Modify the commit message by passing the new trailers using something like `cat message.txt | git interpret-trailers --trailer JIRA-Ref=my-new-ref`

In such a scenario, I’d probably also remove the `trailer.jiraref.command` option that we set before simply to have all the processing logic inside (2) and not being split between that and the git configuration.

## Use case

Coming to think of it, trailers might be a nice way to include a list of all the reviewers that checked a pull-request/merge-request before it was merged! Implementing that as some kind of hook/action for GitLab or GitHub might also be a nice coding dojo project 😄
