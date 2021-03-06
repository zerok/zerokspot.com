---
date: '2005-04-07T12:00:00-00:00'
language: en
tags:
title: 'Re: An observation about tech support'
---


Jason Fried posted on <a href="http://37signals.com/svn/archives2/2005/04/an_observation.php">signal vs. noise</a> a very interesting entry about his view on tech support in general and where it could perhaps be optimized. I'm somehow biased here since I'm (1) on a support team and (2) being a developer myself, but I'll still try to be as open-minded as possible ;-) 

-------------------------------



From my point of view there is a quite simple problem on tech support: How do you know (as a supporter) who tech savvy the customer is? Simple answer: You don't. So normally you should start the whole dialog on a quite basic level which could make handling this request a quite lengthy task. Depending on what you actually support you could offer the customer some kind of form where he/she can give some basic information about the kind of problem he/she is experiencing. In my opinion it shouldn't be the question if you offer your customers such a form or template but how and what kind of. 

Such a request method can have positive and negative effects. While it can help people bringing some structure in their support request, it can also lead to confusion on what should be entered in which field. If the form is well documented (descriptive labels etc.) it should at least help the customer to write more. On the other side it also helps the support stuff by giving them more background information and also provides them with some more information about the "tech savvy" of the customer ... at least the customer shows that he/she knows what and what version of it he/she's using.

The form should IMO only include quite basic stuff. If you're supporting some web application it would be helpful for the support staff to know some details about the server setup. Since webapplications should normally installed by people who are also willing to administrate them some documentation linked in the labels about how to get the environment information should not be too much for the user. If the user does not want to administrate the webapp you have a problem ... I won't go into details here simply because it's simply getting to deep into the project philosophy but depending on the kind of service you provide you have a quite wide range of options on how to deal with such people. Just some examples:

* Rethink your support approach
* Suggest an alternative application that is easier to maintain (for example completely managed environments)
* Rethink what you can do to make your software easier maintainable

I selected these options for a purpose: Support is not a separated part of a project. Support and Development have to co-exist otherwise you offer a software for developers (and even they will probably ignore you if you don't offer at least the basics of support)... but back to the form/template suff ;-)

As already indicated the form should offer the user to specify some details about the environment the product is being used in. The more detailed the better, but it's probably easier if the user has only some very general options. They will help the support but should not frustrate the user.

Another very basic field would be the version of the product the user is using. In rapidly evolving software like it's produced by many opensource/Free software projects minor releases appear in a quite short intervals so the problem could be already fixed in newer versions (which is in my experience the case in about 25% of the support requests that at least looks somehow like a bug somewhere ;-) ). Here is it also important that the version number can be easily accessed by the user.

These are in my opinion the most important fields for such a form or template... just to clear that up: I'm talking about templates simply because I'm used to forum based support where the support staff could suggest (or require) that the user uses some kind of structured forum post. Something like this can be for example <a href="http://www.phpbb.com/phpBB/viewtopic.php?t=128123">here</a> on phpBB.com.

<blockquote>What do you think? How do you feel as a customer when you have a problem with a product? Do you prefer the simple support email address, or do you prefer an elaborate ticketing system with multiple levels of issue classification?</blockquote>

Since Jason asked for people's opinion from the user side I will also try to give my opinion here. It seems to be a general of non-FreeSoftware/ope.... non-FSOS companies/projects' tech support, that they can't really deal with tech savvy people. I can only write about some of my experience with some IT companies (ISPs, software companies...) but I wasn't really happy with any of them. Here I had the problem, that e-mail support is only very limited and telephone support was more a joke. Some supporters seem to have problems accepting information provided by the user as correct or at least partially reliable. It's true that trust is good but controll is better, but still some basic level of trust should be established between the tech supporter and the user, otherwise the whole process will probably become quite a pain.

Here esp. the first response by the supporter is important. If the user more or less indicates that he/she's just a user without any special tech competence the supporter has probably not really a problem if he/she can explain everything in simple terms. But what to do if the user indicates some tech savvy? Should the supporter suggest some at least mid-tech-level dialog? But what if the user isn't as tech savvy as he/she indicated? Quite a dilemma for the supporter. In my opinion it is also partially the user's problem here. If he/she's tech savvy the user should indicate this strongly, so that it's clear for the supporter that some techspeach is possible here.

I had for example once the problem that I made a support request but received more or less your automated answer asking for more details and doing the basics steps that should be done by any user. Nothing wrong here - only the context. I already provided most of the information requested in this answer. In this example the supporter made 2 mistakes:

* Answering in the wrong speach level since I already indicated due to my request style that I'm tech savvy
* Not reading my request completely

Both shouldn't happen or at least not often. They happen. It's natural that they happen simply because support is a quite hard job. Burning out once in a while is normal for supporters... ok, I'm probably off-topic once again :-? Sorry, was a long day.

If something like that can be really prevented by using for example support forms is the other question. IMO yes, since they can easily be automatically or at least semi-automatically filtered which makes the job once again easier for the supporter.

Still some people are more comfortable with e-mail instead of forms (I am an example for such a person ;-) ) simply because it doesn't force the user to use a completely new environment again.

Ok, I think I'm done now... otherwise it will become even more a view on tech support as a whole :-? Sorry everyone for another quite lengthy article :-?