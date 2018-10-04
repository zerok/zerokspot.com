---
title: "No worker satisfying: resource type"
date: 2018-10-04T20:43:18+02:00
tags:
- concourseci
- "ci-cd"
---

Today, while trying to restructure one of our primary pipelines in
[Concourse](https://concourse-ci.org/), I ran into a weird issue with a new
resource I was working on:

```
no workers satisfying: resource type 'snapshot'

available workers: 
  - platform 'linux'
  - platform 'linux'
```

"snapshot" was the name of the resource I was working on. That error was
confusing mostly because of its reference to the two workers labelled as
`platform 'linux'`. At first I thought I had mistaken put/get operations on a
resource with defining a task.

That led to nothing but while looking through the complete pipeline definition
again I noticed that I had forgotten to also define the `resource_type` for
`snapshot`. So, all I had there was a resource but no resource-type for it.
Once I added that, the error disappeared. In case you get the same error: Check
your resource-types 😅
