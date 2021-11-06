---
title: "Documenting HTTP requests with Hurl"
tags:
- testing
- documentation
- http
date: "2021-11-06T20:03:00+01:00"
---

I don't really remember anymore how I stumbled upon [Hurl](https://hurl.dev/) but nevermind that, it looks awesome! It's basically a wrapper around cURL that allows you to write a handful of HTTP requests into a single plaintext file and execute them. You can then add assertions for responses and also pass data from one response to follow-up requests.

Why is that interesting? Especially when you want to document APIs and show how to test them, it's useful to provide testers a list of HTTP requests that they should execute and what kind of response is to be expected. In some situations you also need to modify follow-up requests based on the output of previous ones.

If you're using Emacs, you can use [restclient.el](https://github.com/pashky/restclient.el) for things like that. Just give your testers a file formatted for that mode and they can quickly run all the requests defined in there. But what if your testers *are not* using Emacs? Hurl is more lightweight and easy to onboard, but let me show that using a simple example: Running a chain of requests against the GitLab API.

Put the following content into a file named `gitlab.hurl`:

```hurl
# Let's first get details about the user who's associated with the given
# token.
GET https://gitlab.com/api/v4/user
Authorization: Bearer {{GITLAB_TOKEN}}

HTTP/* *
[Asserts]
status == 200
header "Content-Type" == "application/json"
[Captures]
user_id: jsonpath "$.id"


# Now that we have the current user, let's fetch all the merge requests taht
# I still have to review:
GET https://gitlab.com/api/v4/merge_requests?reviewer_id={{user_id}}&state=opened&scope=all
Authorization: Bearer {{GITLAB_TOKEN}}

HTTP/* 200

```

This [Hurl-file](https://hurl.dev/docs/hurl-file.html) contains two entries:

1. Requesting details about the current user in order to retrieve their user ID.
2. Request all opened merge requests where the user is marked as reviewer.

You can then run the requests in there with the following command:

```
hurl gitlab.hurl --variable GITLAB_TOKEN=$GITLAB_TOKEN
```

Each entry consists of a request specification (i.e. `GET ...` lines + the lines below that to declare some headers) followed by some definitions for the response. Let's take a more detailed look at the first entry:

```hurl
GET https://gitlab.com/api/v4/user
Authorization: Bearer {{GITLAB_TOKEN}}
```

Here you run a GET request and pass a variable named `GITLAB_TOKEN` as HTTP header. This is the variable we passed using the `--variable KEY=VALUE` parameter we specifed in the command above.

The response is then specified below:

```hurl
HTTP/* *
[Asserts]
status == 200
header "Content-Type" == "application/json"
[Captures]
user_id: jsonpath "$.id"
```

Basically, you just say here that you expect a HTTP response with some status code. The actual assertions are then included in the `[Asserts]` block: A status code of 200 and a content-type of `application/json`. Finally, for the next request you need the user's ID which we can capture in the `[Capture]` block using JSONPath.

What Hurl *cannot* do is to fan out to multiple requests based on a previous one. The number of requests and their broad specification has to be explicitly written down in the input file. That being said, especially for documenting and testing APIs for normal users or by human testers Hurl offers a clean and simple text format with an easy to use CLI.

And nobody's really stopping you but also using Hurl files within your CI setup 😁

