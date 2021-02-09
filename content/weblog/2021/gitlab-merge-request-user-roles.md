---
title: "Assignee, approver, reviewer, oh my!"
date: "2021-02-09T19:35:00+01:00"
tags:
- gitlab
- 100daystooffload
incoming:
- url: https://chaos.social/@zerok/105702780814787748
---

I'm working daily with [GitLab](https://gitlab.com) and absolutely love it. When it comes to all the different roles that a person can have around issues and merge requests, though, I sometimes get confused and so I thought I'd finally write them down so that even I can remember:

## Assignee

Each merge request and issue can have one or multiple users assigned to it. That assignment can change during the life-time of that item but in both contexts these are the people to contact if there are questions about that MR or issue.

## Merge request author

In the context of merge requests there is also always an author. That's the user that initially created the merge request. As far as I can tell, this attribute of an MR cannot be changed while assignees can.

## Approver/Reviewers

A reviewer is someone who you want to review your merge request. Actually, you want that they approve your merge request which makes them an “approver”.  That’s confusing… let’s work with a little example:

Let’s say you’re the maintainer of ProjectX. This means that whenever a new merge request is opened, you should be automatically assigned as reviewer so that nobody can merge it without you giving it an approval first!

To do that, you create an [“approval rule”](https://docs.gitlab.com/ee/user/project/merge_requests/merge_request_approvals.html) within your project that set the number of approvals required to “1” and add yourself to the list of “approvers”. 

OK, but that only decides who’s approval is required for an MR. If you want to explicitly request a review from someone, you can now add them to a merge request as a “reviewer”. Once you are added to an MR as a reviewer, you can also find that request in the quick menu in GitLab’s top bar:

{{<figure src="/media/2021/gitlab-to-review.png" alt="Screenshot of the GitLab top bar with expanded merge-request menu" caption="Once you are a review of an MR, you can easily navigate to it using the top bar menu">}}


## Code owner

Related to approvers are [code owners](https://docs.gitlab.com/ee/user/project/code_owners.html). These are defined inside the `CODEOWNERS` file within the repository and indicate which people are responsible for which part of the repository. GitLab can add these people to the required approvers *if* code that's under their responsibility would be changed by a MR. This feature can be configured for protected branches.

I think that’s basically it for now. Personally, I think the whole approver vs. reviewer naming is a bit unfortunate but otherwise everything is pretty clear once you’ve understood what each role is there for 🙂
