---
date: '2010-04-22T12:00:00-00:00'
language: en
tags:
- mercurial
- hg
- theming
- styling
title: Styling the hgwebdir index the hard way
---


One of the most prominent (and rightly so) features of
[Mercurial](http://mercurial.selenic.com/) is its built-in web-interface that
not only lets you pull from your repositories via plain old HTTP but also offers
a simple history and graph browser (among other things).

But if you want to use hgweb(dir) within a company you probably also want to
have it styled in your CI. For the most part creating your own style is rather
simple. But while trying it out for a simple CI-style I ran into a small issue.

-------------------

Creating a style that you want to use for your repository is rather
simple:

1.  Create your map-file and all the relevant templates (or just copy them over
    from for example the "paper" style) in some arbitrary location that can
    be accessed by your hgwebdir user. The
    [wiki](http://mercurial.selenic.com/wiki/Theming) has some information about
    what comes into each file.

2.  Use the web.templates and web.staticurl settings within your hgweb.config
    to tell Mercurial where to find your style and the URL your static content
    is reachable at:
    
    <pre><code>[web]
    templates = /path/to/your/templates-folder/new-style
    staticurl = http://media.server.com/hg-static
    </code></pre>

This works for all the pages that are associated directly with a *single 
repository*. But once you get to the repository-listing, the whole "in some
arbitrary directory" part no longer seems to work. hgwebdir will always fall
back to "paper" (the default style).

From what I can tell, the reason for this is that hgwebdir for the index
relies exclusively on the style-name, which is something you can't really
influence with the web.templates setting. And styles resolve exclusively to
folder in your mercurial/templates folder within your Mercurial installation.

So far I've only found 2 workarounds for this and neither is really all that
appealing:

*   Symlink or copy your new template folder into mercurial/templates

*   Monkey-patch the mercurial.templater module from within the WSGI script
    to also look within your own template-repository.
    
    <pre><code>from mercurial import templater
    templater.path[0:0] = ['/path/to/your/templates-folder/']
    </code></pre>

In either case you no longer have to set web.templates but can just select
your style using web.style. IMO the second solution is a bit more flexible
since you don't have to re-link your style if you update Mercurial and your
update-tool decides to first remove the mercurial package altogether.

Now I'm just curious if there is a cleaner solution out there :-) There has 
to be a cleaner way to manipulate the search path and that I have missed.