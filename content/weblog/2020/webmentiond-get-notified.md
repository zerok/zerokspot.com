---
title: 'Webmentiond: Get notified!'
date: "2020-06-22T16:16:54+02:00"
tags:
- webmentiond
- feature
---

[Last Friday I implemented](https://github.com/zerok/webmentiond/commit/edd07281ec655a55ba1d93b87314383e717b99a7) a little feature into [webmentiond](https://github.com/zerok/webmentiond) that will notify you, the administrator or an instance, of newly verified mentions. When a new mention is received by the service, it will first go out and check if the source really links to one of your articles. If it does, then the service marks the mention as “verified” inside the database. Right after that, a notification email is sent out to every address specified with the `--auth-admin-email` flag. 

<figure><img src="/media/2020/Screenshot%202020-06-22%20at%2016.12.25.png"><figcaption>Notification email with links to source and target</figcaption></figure>

Please note that this feature has to be explicitly enable using the `--send-notifications` flag.

I wanted to add this simply because I tend to otherwise refresh the admin ui in my mobile browser far too often 😅  If you want to give it a try, too, just download the latest release (binary or Docker) and set that flag when starting the server 🙂
