---
date: '2006-04-17T12:00:00-00:00'
language: en
tags:
title: Using Ruby-Gettext installed with rubygems
---


Since English is not my native language i18n is one of the hot topics for me on and off the web. Let's see, how is it done by most open source applications? Right, [GNU gettext](http://www.gnu.org/software/gettext/). There's also a library for Ruby called [Ruby-Gettext](http://www.yotabanana.com/hiki/ruby-gettext.html?ruby-gettext). Since I love rubygems the first thing I checked when I learnd about Ruby-Gettext was if there was a gem available for it. Yes there is, in fact there are multiple versions:

<pre class="console">$ sudo gem install gettext
Password:
Attempting local installation of &apos;gettext&apos;
Local gem file not found: gettext*.gem
Attempting remote installation of &apos;gettext&apos;
Select which gem to install for your platform (powerpc-darwin8.5.0)
 1. gettext 1.4.0 (ruby)
 2. gettext 1.4.0 (mswin32)
 3. gettext 1.3.0 (ruby)
 4. gettext 1.3.0 (mswin32)
 5. gettext 1.2.0 (ruby)
 6. gettext 1.2.0 (mswin32)
 7. gettext 1.1.1 (mswin32)
 8. gettext 1.1.1 (ruby)
 9. gettext 1.1.0 (mswin32)
 10. gettext 1.1.0 (ruby)
 11. gettext 1.0.0 (ruby)
 12. gettext 1.0.0 (mswin32)
 13. Cancel installation
&gt;</pre>



-------------------------------



But how do you get it to work after installing it? It seems like all the documentation available on the author's homepage simply goes with the normal installation without using rubygems. So when I installed the gem, the first thing I tried was:

<pre class="code">require 'rubygems'
require_gem 'gettext'
puts Gettext::_('Hello world')</pre>

Doesn't really work. The whole Gettext module seems to be missing. First I thought, that I had missing an error message during the installation, but all the files were there. 

In the end it took me quite some time to find the simple solution for this problem ;)

<pre class="code">require 'rubygems'
require_gem 'gettext'
require 'gettext'
puts Gettext::_('Hello world')</pre>

It seems like, the require_gem call only appends the gettext path to the library path and doesn't auto-require anything really gettext related.