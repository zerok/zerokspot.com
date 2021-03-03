---
date: '2006-05-14T12:00:00-00:00'
language: en
tags:
- ruby
title: ActionMailer in standalone apps
---


Today someone in #ruby-lang asked if you could use Rails' ActionMailer also in standalone ruby applications. That user left after a short time and the question stayed unanswered, although it is my opinion quite interesting :) So I started digging a little bit through the documentation and came up with a quite basic solution which actually holds no real suprises. 



-------------------------------



<pre class="code">require &apos;rubygems&apos;
require &apos;pp&apos;
require_gem &apos;actionmailer&apos;
<strong>ActionMailer::Base.template_root = &apos;templates&apos;</strong>
ActionMailer::Base.delivery_method = :test
class TestMailer &lt; ActionMailer::Base
  def test
    @recipients = &apos;re@domain.com&apos;
    @subject = &apos;Hello World&apos;
    @from = &apos;sender@domain.com&apos;
  end
end
if $0 == __FILE__
  TestMailer.deliver_test
  ActionMailer::Base.deliveries.each do |mail| 
    puts &quot;#{&apos;=&apos;*60}&quot;
    puts &quot;Subject: #{mail.subject}&quot;
    puts &quot;From: #{mail.from}&quot;
    puts &quot;To: #{mail.to}&quot;
    puts &quot;#{&apos;-&apos;*60}&quot;
    puts &quot;#{mail.body}&quot;
  end
end</pre>

As in Rails you simply create an ActionMailer subclass and do the deliver_action thing with it. The only thing that is slightly different here, is that you have to specify a template\_root directory. The rest of the configuration should be the same as for its use in a Rails webapp (for details please read [the manual](http://api.rubyonrails.org/classes/ActionMailer/Base.html)) In this case I use the 'templates' folder in the same directory as the script (actually in the pwd but anyway ;)).

So for this example you would have to have following folder structure

<pre class="output">./templates
./templates/test_mailer
./templates/test_mailer/test.rhtml
./test.rb</pre>

Running this script should give you the expected mail :)

**Note:** I haven't tested it with anything but the :test delivery\_method, so perhaps there are some hidden problems :)
