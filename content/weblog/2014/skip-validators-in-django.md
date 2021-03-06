---
date: '2014-02-10T19:30:25-00:00'
language: en
tags:
- django
- development
title: Skipping validators in Django forms/models
---


For a couple of versions now Django has had a nice little API for [registering validator functions][2] for specific fields in either forms or models that are checked using the model validation logic:

```
class UserProfile(models.Model):
    avatar = models.ImageField(
        upload_to='avatars',
        validators=[
            check_format,
            check_minimum_resolution
        ])
```

What's great about this API is that it encourages the creation of tiny, self-contained value-checks that are easily testable. Both, `check_format` and `check_minimum_resolution`, can be tested independently of each other without running into tests that might mix up cases from the other.

The downside of the current implementation in Django is, that *all* validators are always executed, no matter if a previous one already reported an error and therefor the field's value can no longer end up being. Normally, this isn't really an issue but it might become one if one of you validators is rather expensive.

In our previous example we had two checks: One for the format of a file and one for the image's resolution. While the first validator can (in a rather naive implementation) be as simple as checking the file's extension, the latter really requires that we open the file and look at its content. Ideally, we'd want to avoid this if the format is not supported anyway.

By default, there is no easy way to just register a validator that is skipped if a previous one already raised an exception. Luckily, Django internally just takes the list of validators you provided, adds that to a list of default validators and then iterates over it. If an exception is raised, it will be added to a list of errors and then it's the next validator's turn:

```
# django.forms.fields.Field
def run_validators(self, value):
    if value in self.empty_values:
        return
    errors = []
    for v in self.validators:
        try:
            v(value)
        except ValidationError as e:
            if hasattr(e, 'code') and e.code in self.error_messages:
                e.message = self.error_messages[e.code]
            errors.extend(e.error_list)
    if errors:
        raise ValidationError(errors)
```

(More or less the same code also exists in `django.db.models.fields.Field`.)

Since this code only used a simple iterator, I tried to get around this limitation by creating my own little iterator that notices when something throws an error and then just skips the return of the next item if it was marked as "skip on error". I haven't worked on that level with iterators for years but luckily the API is simple enough that within a couple of minutes of the [DjangoWeekend sprint in Cardiff][3] I had some working code, could test it and integrate it into a project I'm currently working on. The end result of this morning-session can be found in the [django-validatorchain][1] package.

If we apply the code there to our original example, we would end up with something like this:

```
from validatorchain import ValidatorChain

class UserProfile(models.Model):
    avatar = models.ImageField(
        upload_to='avatars',
        validators=ValidatorChain()
            .add(check_format)
            .add(check_minimum_resolution, skip_on_error=True)
    )
```

Now the `check_minimum_resolution` validator is no longer executed if `check_format` threw an error before.

If you had similar requirements, perhaps this little package will help you :-)

[1]: https://github.com/zerok/django-validatorchain
[2]: https://docs.djangoproject.com/en/1.5/ref/forms/validation/#validators
[3]: http://djangoweekend.org/
