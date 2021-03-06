---
date: '2007-02-23T12:00:00-00:00'
language: en
tags:
- dreamhost
- trac
title: Correcting pretty-URLs for FCGI in Trac 0.10.x
---


I'm currently in the process of installing Trac on my Dreamhost account and came across the problem, that the server doesn't allow me to work with Aliases and ScriptAliases in .htaccess files. So I had to go back to using symlinks and RewriteRules. 

The problem here now seems to be, that if you for example have a dispatch.fcgi in your htdocs, all your URLs will look for example like this: [http://domain/path/dispatch.fcgi/wiki](http://domain/path/dispatch.fcgi/wiki). 


-------------------------------


Ugly! :-) I found in various tutorials that you have to change the trac/web/cgi_location.py to get rid of the SCRIPT\_NAME prefix. Well, it seems like the code has been a little bit refactored.

Now the relevant piece of code seems to be in trac/web/api.py:

<pre class="code python">base_path = property(fget=lambda self: self.environ.get('SCRIPT_NAME', ''),
                     doc='The root path of the application')</pre>

I'm not completely sure, but shouldn't changing this to the following solve this? At least in a test installation it worked :-)

<pre class="code python">base_path = property(fget=lambda self: os.path.dirname(self.environ.get('SCRIPT_NAME', '')),
                     doc='The root path of the application')</pre>

My main concern with this is, that I don't know what collateral damage it might cause, but I guess if you're really just using it for the setup mentioned above there shouldn't be too many risks involved :-?

To bad I had to hack the core for this to work. If anyone has found a better solution for such a situation, please let me know :-) (esp. one that doesn't involve core-modifications). 

P.S.: I also found [one tutorial](http://natmaster.com/articles/installing_trac_0.10.php) where the author simply set the base_path to '/'. Definitely easier, but it only works if you have Trac installed right in your domain's root directory.