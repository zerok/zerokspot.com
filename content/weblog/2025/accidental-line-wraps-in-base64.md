---
date: "2025-07-03T14:58:08+02:00"
tags:
- til
- linux
- macos
title: Accidental line-wrapping in Base64
---

This is one of those findings that are just stupid. There was a bug in some CI script where we just manually generate the Basic-Auth HTTP header using something in a secret and the `base64` utility:

```
auth_header=$(echo -n “$USER_ID:$TOKEN” | base64)
```

That doesn’t look completely wrong, does it? Turns out, base64 on Linux has a different behavior compared to, let’s say, MacOS: It will by default use line-wrapping at 76 characters:

```
USER_ID=hello_world
TOKEN=01234567890123456789012345678901234567890123456789012345678901234567890123456789
auth_header=$(echo -n “$USER_ID:$TOKEN” | base64)
echo $auth_header
aGVsbG9fd29ybGQ6MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0
NTY3ODkwMTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODk=
```

Yes, there’s a newline here when using that 80 character token. If you had run that script on MacOS, no such line-break would be part of the output.

Especially for scripting my assumption is that most of the time you actually do not want line-breaks in your Base64 output. So if you’re on Linux, always make sure to use `base64 -w 0` unless you want that wrapping. 
