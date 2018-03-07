---
title: "Paging in gopkg.in/ldap.v2"
date: 2018-03-07T20:11:56+01:00
tags:
- golang
- ldap
---

Over the last couple of months I've been working more and more with the awesome
LDAP package for Go available on [gopkg.in/ldap.v2](http://gopkg.in/ldap.v2)
but only today I needed to do paging. The documentation didn't contain any
example or other hints outside of the existence of the
[ControlPaging](https://godoc.org/gopkg.in/ldap.v2#ControlPaging) so it took me
a while to get it to work.

The basic idea with that is that when you create a new search request, you can
pass a slice of so-called "controls" that act as extension points for the core
LDAP protocol-suite. LDAP core doesn't specify any sort of paging for dealing
with large result sets, but luckily
[RFC2696](https://www.ietf.org/rfc/rfc2696.txt) specifies such a control which
is implemented in the `ControlPaging` struct:

```
paging := ldap.NewControlPaging(100)
for {
    req := ldap.NewSearchRequest(baseDN, ldap.ScopeWholeSubtree,
        ldap.DerefAlways, 0, 0, false, filter,
        []string{"name", "member"}, []ldap.Control{paging})
    result, err := conn.Search(req)

    ...

    resultCtrl := ldap.FindControl(result.Controls, paging.GetControlType())
    if resultCtrl == nil {
        break
    }
    if pagingCtrl, ok := resultCtrl.(*ldap.ControlPaging); ok {
        if len(pagingCtrl.Cookie) == 0 {
            break
        }
        paging.SetCookie(pagingCtrl.Cookie)
    }
}

```

What I do here, is to first create a new `ControlPaging` object with a page
size of `100`. Next, I include this control in the `controls` parameter when
creating a new search request. Once I'm through with processing the entries
that have been returned for the "current page", I update the paging control
with the cookie returned from the server and continue to do requests until
either I don't get a new paging control back from the server or its cookie is
empty.

I will probably refine that example and create a PR for the project if time
permits.
