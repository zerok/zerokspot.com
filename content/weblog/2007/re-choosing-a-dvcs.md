---
date: '2007-12-31T12:00:00-00:00'
language: en
tags:
- dscm
- dvcs
title: 'RE: Choosing a Distributed Version Control System'
---


Dave Dribin recently wrote a great [article](http://www.dribin.org/dave/blog/archives/2007/12/28/dvcs/) comparing 3 of the most prominent
distributed version control systems (git, bzr and mercurial) followed up 
yesterday by his reasoning for [choosing Mercurial](http://www.dribin.org/dave/blog/archives/2007/12/30/why_mercurial/).



-------------------------------

Especially following sentence simply nails it:

> There really is no run-away winner for me. 

That's the whole problem. I've been moving from svn to hg to bzr to hg to bzr 
to hg ... and I also gave git a try recently, just to jump back to hg again.
The world would just be that much nicer if git would have better Windows
support or Mercurial would support renaming better and directory 
versioning at all ... not to mention fix that damn [rename bug](http://www.selenic.com/mercurial/bts/issue883). I won't even
start with bzr becoming faster, although props to the team at Canonical 
and all the other contributors for making it significantly faster over the 
last year. 

I just want to make some small ammendments to Dave's articles since he didn't
mention some of the pros and cons of git and Mercurial:

*   For me personally, Mercurial's web interface was one of the major 
    selling-points since it not only supports CGI but can easily extended
    to work like any other WSGI application and therefor works so far
    quite nicely through mod_wsgi and FastCGI.
    
*   This is also my problem with git. Their web interface seems to be 
    limited to classic CGI for now.
    
*   I also can't really reproduce Dave's problems with installing bzr 
    on MacOSX. I've been using it on Tiger and Leopard now and never had
    any bigger problems. 
    
*   Another point I don't completely agree (yet I see where he's coming from)
    is his complaint about Git using SHA-1 hashes to identify revisions. 
    Yes, this is annoying, but since you can use shortcuts here (for example
    only the first 3 or so chars if they are still unique in your repo)
    or use relative numbers like HEAD^^ (for the HEAD's grandparent)
    it was far less a problem for me.
    
*   "Mercurial, in contrast, does not have the concept of local branches." is
    also not completely correct since they actually [do exist](http://hgbook.red-bean.com/hgbookch8.html#x12-1650008.5), 
    they are just not as prominent as in Git.
    
So in the end, Mercurial is also for me currently the winner, although I 
will probably also keep and eye on bzr and esp. on git to see if the first
becomes competitive in the speed-department and the latter regarding
Windows-compatibility. 

[via [DaringFireball](http://daringfireball.net/linked/2007/december#sun-30-dribin)]