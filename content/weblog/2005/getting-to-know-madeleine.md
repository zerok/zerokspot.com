---
date: '2005-07-29T12:00:00-00:00'
language: en
tags:
title: Getting to know Madeleine
---


<p>On my quest for more ruby skills I&#8217;m currently trying to get to know Madeleine. No, not a girl, although this would be quite a welcome change for me actually. I&#8217;m talking about the object persistancy layer for Ruby.</p>

-------------------------------


<p>Until now I wasn&#8217;t able to find a good tutorial for <a href="http://madeleine.sourceforge.net">Madeleine</a> so this will be probably as much an adventure for you as it will be for me. I will simply write what I found out about Madeleine so far and hope, that it will also help myself to get to know this beautiful g&#8230; system more.</p>
<p>For now let&#8217;s start with the purpose of Madeleine as far as I&#8217;ve gasped it: With Madeleine you can basically make a Ruby object persistant so that you won&#8217;t have to manually load it the next time again. Such a persistant object can be accessed for reading and writing. Both kinds have their own wrapper methods which makes it easier for Madeleine to determine when data is about to be changed. To save the data a snapshot system is used. Only changes from the last snapshot get into the new one.</p>
<p>Here also comes the first thing I&#8217;ve noticed that I want to change but haven&#8217;t found out how yet: The persistant object has the method <span class="function">take_snapshot</span> which generates a new snapshot &#8230; no matter if anything was changed at all. Anybody else smells the stench of dying inodes here?</p>
<p>Ok, enough of the theory for now. Let&#8217;s get our hands dirty a little bit. So let&#8217;s install madeleine and include it into our new <span class="file">playground.rb</span>.</p>
<pre><code>require 'rubygems'
require_gem 'madeleine'</code></pre>
<p>First we should decide what kind of data we want to make persistant. During this tutorial I will simply use a wrapper class for an array:</p>
<pre><code>class TestContainer
    def initialize
        @data = []
    end
    def each
        @data.each{|field| yield field}
    end
    def &lt;&lt; (element)
        @data &lt;&lt; element
    end
end
</code></pre>
<p>As I&#8217;ve already mentioned before, Madeleine provides a wrapper which has two methods for accessing the data:</p>
<dl><dt>execute_query</dt><dd>... is used to access the object for reading while &#8230;<br /><dt>execute_command</dt><dd>... executes a command that can also change the state of the object. So I&#8217;d guess that it holds an exclusive lock on it.</dd></dl>
<p>Both methods take a socalled <span class="class">Command</span> object as parameter. Nothing really special here: A <span class="class">Command</span> object is simply an object offering an <span class="method">execute</span> method. Since we have a data container let&#8217;s create two commands. One for iterating about the whole content and one for appending new stuff to the array.</p>
<pre><code>class AppendCommand
    def initialize(data)
        @data = data
    end
    def execute(system)
        system &lt;&lt; @data
    end
end
class QueryCommand
    def execute(system)
        system.each {|f| puts f}
    end
end

</code></pre>
<p>The <span class="method">execute</span> method has one single argument that basically holds the object as we know it. So <span class="variable">system</span> in the <span class="class">AppendCommand</span> and <span class="class">QueryCommand</span> objects will simply be a <span class="class">TestContainer</span> object.</p>
<p>Let&#8217;s bring the persistancy into all this, shall we?</p>
<pre><code>
if $0 == __FILE__
    madeleine = SnapshotMadeleine.new("storage") do
        TestContainer.new
    end
    if ARGV[0] == '-a'
        cmd = AppendCommand.new("#")
        madeleine.execute_command(cmd)
    else
        cmd = QueryCommand.new
        madeleine.execute_query(cmd)
    end
    madeleine.take_snapshot
end
</pre></code>
<p>We create or storage using the <span class="class">SnapshotMadeleine</span> class which takes as first parameter the folder where the presistant object should be stored and a block returning the object we want to make persistant. In our example a new instance of a <span class="class">TestContainer</span>. Now we can use the <span class="method">execute_query</span> and <span class="method">execute_command</span> methods on the persistant object using our two Command classes. Since we want to make this whole thing persistant after all, we execute the <span class="method">take_snapshot</span> method before exiting the script.</p>
<p>To save some inodes you should probably create a policy about when to take snapshots. For example only every hour or if a modifying command has been executed ;)</p>
<p>In instiki they seem to have an additional logfile that holds all the changes to the object that haven&#8217;t been committed into a snapshot yet. I&#8217;m not sure though, if this is a feature of madeleine oder an addition by instiki.</p>
</div>
<div id="changes" style="display: none">
  <p style="background: #eee; padding: 3px; border: 1px solid silver">

    <small>
      Showing changes from revision #1 to #2:
      <ins class="diffins">Added</ins> | <del class="diffdel">Removed</del>
    </small>
</p>
<p>On my quest for more ruby skills I&#8217;m currently trying to get to know Madeleine. No, not a girl, although this would be quite a welcome change for me actually. I&#8217;m talking about the object persistancy layer for Ruby.</p>
<p>Until now I wasn&#8217;t able to find a good tutorial for <del class="diffmod">Madeleine </del><ins class="diffmod"><a href="http://madeleine.sourceforge.net">Madeleine</a> </ins>so this will be probably as much an adventure for you as it will be for me. I will simply write what I found out about Madeleine so far and hope, that it will also help myself to get to know this beautiful g&#8230; system more.</p>
<p>For now let&#8217;s start with the purpose of Madeleine as far as I&#8217;ve gasped it: With Madeleine you can basically make a Ruby object persistant so that you won&#8217;t have to manually load it the next time again. Such a persistant object can be accessed for reading and writing. Both kinds have their own wrapper methods which makes it easier for Madeleine to determine when data is about to be changed. To save the data a snapshot system is used. Only changes from the last snapshot get into the new one.</p>
<p>Here also comes the first thing I&#8217;ve noticed that I want to change but haven&#8217;t found out how yet: The persistant object has the method <span class="function">take_snapshot</span> which generates a new snapshot &#8230; no matter if anything was changed at all. Anybody else smells the stench of dying inodes here?</p>
<p>Ok, enough of the theory for now. Let&#8217;s get our hands dirty a little bit. So let&#8217;s install madeleine and include it into our new <span class="file">playground.rb</span>.</p>
<pre><code>require 'rubygems'
require_gem 'madeleine'</code></pre>
<p>First we should decide what kind of data we want to make persistant. During this tutorial I will simply use a wrapper class for an array:</p>
<pre><code>class TestContainer
    def initialize
        @data = []
    end
    def each
        @data.each{|field| yield field}
    end
    def &lt;&lt; (element)
        @data &lt;&lt; element
    end
end
</code></pre>
<p>As I&#8217;ve already mentioned before, Madeleine provides a wrapper which has two methods for accessing the data:</p>
<dl><dt>execute_query</dt><dd>... is used to access the object for reading while &#8230;<br /><dt>execute_command</dt><dd>... executes a command that can also change the state of the object. So I&#8217;d guess that it holds an exclusive lock on it.</dd></dl>
<p>Both methods take a socalled <span class="class">Command</span> object as parameter. Nothing really special here: A <span class="class">Command</span> object is simply an object offering an <span class="method">execute</span> method. Since we have a data container let&#8217;s create two commands. One for iterating about the whole content and one for appending new stuff to the array.</p>
<pre><code>class AppendCommand
    def initialize(data)
        @data = data
    end
    def execute(system)
        system &lt;&lt; @data
    end
end
class QueryCommand
    def execute(system)
        system.each {|f| puts f}
    end
end
</code></pre>
<p>The <span class="method">execute</span> method has one single argument that basically holds the object as we know it. So <span class="variable">system</span> in the <span class="class">AppendCommand</span> and <span class="class">QueryCommand</span> objects will simply be a <span class="class">TestContainer</span> object.</p>
<p>Let&#8217;s bring the persistancy into all this, shall we?</p>
<pre><code>
if $0 == __FILE__
    madeleine = SnapshotMadeleine.new("storage") do
        TestContainer.new
    end
    if ARGV[0] == '-a'
        cmd = AppendCommand.new("#")
        madeleine.execute_command(cmd)
    else
        cmd = QueryCommand.new
        madeleine.execute_query(cmd)
    end
    madeleine.take_snapshot
end
</pre></code>
<p>We create or storage using the <span class="class">SnapshotMadeleine</span> class which takes as first parameter the folder where the presistant object should be stored and a block returning the object we want to make persistant. In our example a new instance of a <span class="class">TestContainer</span>. Now we can use the <span class="method">execute_query</span> and <span class="method">execute_command</span> methods on the persistant object using our two Command classes. Since we want to make this whole thing persistant after all, we execute the <span class="method">take_snapshot</span> method before exiting the script.</p>
<p>To save some inodes you should probably create a policy about when to take snapshots. For example only every hour or if a modifying command has been executed ;)</p>
<p><ins class="diffins">In instiki they seem to have an additional logfile that holds all the changes to the object that haven&#8217;t been committed into a snapshot yet. I&#8217;m not sure though, if this is a feature of madeleine oder an addition by instiki.</ins></p>