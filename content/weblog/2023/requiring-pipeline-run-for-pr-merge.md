---
date: "2023-08-30T21:50:43+02:00"
tags:
- github
- development
- ci
title: Requiring a pipeline run for merging a pull-request
---

I want to configure my project on GitHub to require a whole pipeline run to pass before it can be merged. This is supported through "Branch protection rule" where you have to manipulate two settings:

- "Require a pull request before merging"
- "Require status checks to pass before merging"

While the first is quite straight forward, the second requires a bit of explanation. If you enable this flag then you have to specify one or more "statuses" that should be successful before the PR can be merged. A status is more or less just an indicator that you can attach to a commit through the [GitHub API](https://docs.github.com/en/rest/commits/statuses?apiVersion=2022-11-28).

But does this mean that I have to write a custom action in my pipeline that manipulates such a status? No! Every job name in your pipelines becomes its own status. So if I now have a workflow with the name "CI" and a job with the name "main", then I can pick "main" from the select box there and make it mandatory.

This has the side-effect that I should probably rethink how I name the jobs in my workflows...