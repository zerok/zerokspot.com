---
title: This was GoGraz in September
date: "2021-09-22T21:05:40+02:00"
tags:
- golang
- gograz
- meetup
---

Last Monday we once again had our [monthly meetup](https://gograz.org/meetup/2021-09-20/) and this time we *didn’t* have to cancel it at the last minute 🙂 Due to Delta and the whole situation “outside” we opted to go 100%-remote again. To make up for that, I prepared a little presentation but we also had lots of other stuff to talk about:

- The first topic was a quick introduction into [GraphQL](https://graphql.org/) going through some elements of schema definitions and the type system but also how operations are processed. As an example, I provided a little sample application for querying meetup events and sessions associated with them written using [gqlgen](https://gqlgen.com/). I’ll try to find the time to clean up the slides in the next couple of days and then share them together with the source code of the demo application 🙂 **Update:** You can find the slides and source code of the demo application on [GitLab](https://gitlab.com/zerok/gograz-graphql-demo).
- Based on that, Stephan mentioned [FieldMasks in Protobuf](https://netflixtechblog.com/practical-api-design-at-netflix-part-1-using-protobuf-fieldmask-35cfdc606518) which cover at least some of the data-selection of aspects of GraphQL but for gRPC/Protobuf.
- Since gqlgen does most of its work using code-generation, we also talked about some other use-cases for that approach. [mockgen](https://github.com/golang/mock) was mentioned as was the good old [stringer](https://pkg.go.dev/golang.org/x/tools/cmd/stringer) generator. 
- Finally, we had a little demo of [age](https://age-encryption.org/) for doing file-encryption and decryption. I just wish I could use that little tool more… anyway, I posted my thoughts about it [here](https://zerokspot.com/weblog/2021/09/12/hello-age/).

We nearly filled 2 whole hours this time around and I think everyone had lots of fun (at least I hope so) 😅 As for what we’ll do next month, I don’t know yet. Since Delta is probably not going away anytime soon, it’s at least very likely that  the meetup will be done remotely again. Other than that… if you want to give a talk, ping me (Webmention, Mastodon, Twitter, Reddit, e-mail, meetup.com, …) 🙂
