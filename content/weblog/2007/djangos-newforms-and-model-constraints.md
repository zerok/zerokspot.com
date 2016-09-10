---
date: '2007-05-26T12:00:00-00:00'
language: en
tags:
- django
- newforms
title: Django's newforms and model constraints
---


Today I wanted to use one of Django SVN's new features: A newforms's forum for a model instance that only uses a subset of its fields (For example if a model has the fields a, b and c, you could easily create a form that only handles a and b). 

But I found a small problem: You can use such a form, to also write the data back to the database and in the process use the model's data validation and integrity checks. For some reason, only the validation is actually performed, while integrity problems still raise an IntegrityError which you'll have to deal by yourself. 



-------------------------------



I don't know how good and if good at all my solution is, but I guess I can post it anyway:

<pre class="code python">
user = User.objects.get(pk=request.user.id)
ProfileForm = form_for_instance(user, fields=('username','email',))
if request.method == 'POST':
	form = ProfileForm(request.POST)
	if form.is_valid():
		try:
			form.save()
			set_info(request,"Profile successfully updated")
		except IntegrityError:
			message = sys.exc_info()[1].message
			if message.startswith("column "):
				parts = message.split(" ")
				form.errors.setdefault(parts[1],[]).append(" ".join(parts[1:]))
			add_notice(request,"Something went wrong. Please check the marked info for details")
	else:
		add_notice(request,"Something went wrong. Please check the marked info for details")</pre>
		
		
The for this problem relevant part of the code is the handling of the IntegrityError. I basically use the exception's message string in order to find out what form element causes the problem and update the form's error dictionary for the offending element.

It is a hack, to say the least, but since I couldn't find any other attribute in the IntegrityError for identifying the offending constraint for the model.