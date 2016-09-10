---
date: '2006-01-28T12:00:00-00:00'
language: en
tags:
- newsfire_export
title: newsfire_export v1.1.0
---


Three days ago Patrick May let me know, that his <a href="http://www.narf-lib.org/2006/01/plistxml-parser-for-ruby.html">plist parser for Ruby</a> is now also available as gem in the gems.rubyforge.org repository. So I saw a reason to look into gems once again and make my small newsfire_export script available as gem. This way installing it and its requirements (well, plist is the only one ;)) should be far easier now. 



-------------------------------



So how to install this new release now? Simply download the file linked below and run following command in the same directory where you've downloaded that file to:

<pre class="command">sudo gem install plist
sudo gem install newsfire_export-1.1.0-powerpc-darwin.gem</pre>

If you've the previous version of newsfire_export installed on your system, I'd recommend to remove it now. You can also undo the $RUBYLIB path changes I've suggested last time, if you don't need them by now for anything else ;) The above command will install newsfire_export as well as the plist gem if you don't have it on your system yet. It will also put the script itself in your $PATH (default: /usr/bin/)


For those of you who already know and use RubyGems: <s>Sorry for not putting it on rubyforge yet. The project registration pending though :)</s> Damn, this was fast. The newsfire_export script is now also available on rubyforge.org and will hopefully soon make it into their gems archive. The moment it will make it into this repository you can also install newsfire_export with following command:

<pre class="command">sudo gem install newsfire_export
</pre>

Updating works with `sudo gem update newsfire_export`.

To the changes:

* You can now specify an output file using the -o option. If this is omitted $stdout is used.
* -i can be used for specifying an alternative path to NewsFire's config plist.
* A little bit more debugging output if something goes wrong

**[Download](http://rubyforge.org/frs/?group_id=1312&release_id=4117)**

<div class="update">
<p>Sorry, forgot something: If you don't know what RubyGems is: It's a package management system for Ruby, somehow in between apt-get and CPAN (Perl). For details on how to install and use it please check out <a href="http://docs.rubygems.org/">the RubyGems homepage</a> :)</p><span class="date">2006.01.27 21:55</span></div>