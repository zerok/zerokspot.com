---
date: '2006-05-09T12:00:00-00:00'
language: en
tags:
- ruby
title: Require to execute
---


Have you ever wondered, how it works, that by simply putting something like ...

-------------------------------



<pre class="code">

require 'test/unit'

class MyTest < Test::Unit::TestCase

	def test_sample

		assert(false,";)")

	end

end

</pre>



... you can make a unit test in ruby and simply call that file and everything works? Coming from Java I first thought, perhaps also Ruby classes get have some main methods in them to trigger execution if really nothing else is there in the main-scope to execute. But then I found this code in the test/unit.rb:



<pre class="code">

at_exit do

  unless $! || Test::Unit.run?

    exit Test::Unit::AutoRunner.run

  end

end

</pre>



And yes: [at\_exit {block}](http://ruby-doc.org/core/classes/Kernel.html#M002953) is actually something like that. You can register there things that should be run when the main-scope exits. Sweet :)