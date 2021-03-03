---
title: Avatars without Gravatar
date: "2020-05-26T16:23:55+02:00"
tags:
- fediverse
- 100daystooffload
---

Back when [Gravatar](http://en.gravatar.com/) launched many years ago it was a game-changer. All of a sudden, I no longer had to change my profile picture on all those dozens of forums and sites I was active on but could do that in a central place. Years later I still consider it a great service but especially for internal sites problematic to have calls to some central *external* service that is not under your control that in the worst-case can see for what website an avatar has been requested. At the same time, sometimes you simply don’t want to give someone a picture of you just so that they can hand it over to someone else again.

While playing with my [Liberapay](https://liberapay.com/) account I stumbled upon a service similar to Gravatar that I hadn’t heard of before despite it already being nearly a decade old: [Libravatar](https://www.libravatar.org/). At the first glance, Libravatar is pretty much the same as Gravatar except that its source code is open. If you look a bit deeper, though, you’ll notice that it’s actually a federated system. There is an [open protocol behind it](https://wiki.libravatar.org/api/) that let’s you, for instance, run your own instance, and services that support Libravatar *should* check your own service first before doing any kind of fallback (e.g. through the central but struggling Libravatar.org server).


## Client implementation

So how should a website (or a generic client implementation) retrieve a person’s profile picture given that user’s email-address? Let’s work on a little example function that does all that:

	import hashlib
	import dns.resolver
	import sys
	
	
	def get_avatar_url(raw_email):
	    # We need the md5 sum of the lowercase'd email
	    email = raw_email.strip().lower()
	    hash = hashlib.md5()
	    hash.update(email.encode('utf-8'))
	    hash = hash.hexdigest()
	
	    handler_base_url = get_handler(email)
	    return f'{handler_base_url}{hash}'
	
	
	def get_handler(email):
	    handler_base_url = 'https://cdn.Libravatar.org/avatar/'
	    emaildomain = email.split('@')[1]
	
	    # Since browser complain about non-https content on https sites let's
	    # prefer https here:
	    srv_entries = ['_avatars-sec._tcp', '_avatars._tcp']
	    for entry_prefix in srv_entries:
	        try:
	            for answer in dns.resolver.query(
	                    f'{entry_prefix}.{emaildomain}', 'SRV'):
	                host = str(answer.target).rstrip('.')
	                if 'sec' in entry_prefix:
	                    return f'https://{host}:{answer.port}/avatar/'
	                return f'http://{host}:{answer.port}/avatar/'
	        except Exception as e:
	            pass
	    return handler_base_url
	
	
	if __name__ == '__main__':
	    mail = sys.argv[1]
	    print(get_avatar_url(mail))
	

So the first thing a website wanting to show an avatar should do is check the DNS records of the domain part of the email address. If this offers `_avatars-sec._tcp` and/or `_avatars._tcp` SRV entries, these should point to a server that implements the Libravatar protocol and can therefore be queried.

Once we have such a base URL, all we have to do is append the MD5 hash of the complete email address to it and a couple of optional parameters in order to retrieve the image itself:

- `size`/`s` for the width/height of the image to be returned (default seems to be `80`)
- `default`/`d` for the URL of a fallback image should the email not exist in the server’s database. This parameter also has a couple of reserved values for more generic fallback mechanism. You can find a complete list in the spec but here are the most useful in my opinion:
	- `404`: If the email could not be found in the database, then a HTTP 404 error should be returned instead of a fallback image
	- `mm`/`mp`: Generic silhouette icons

The default option is actually a bit more complicated since most implementations are falling back to Gravatar before even considering the `default` option simply because Gravatar also supports this parameter.

In cases where you explicitly want to retrieve the default image, use the `forcedefault=y`/`f=y` parameter in combination with the `default` parameter.

Just to make this clear: The example above is just that. An example. If you want to integrate Libravatar in your own project, please use one of the pre-existing libraries that are listed on [libravatar.org](https://wiki.libravatar.org/libraries/)!


## An experimental server implementation

Since I learn protocols best by implementing them, I also gave this one a try. You can find the complete source code on [https://github.com/zerok/microavatar](https://github.com/zerok/microavatar). This implementation is mostly intended for personal use as you have to upload avatars and set the mapping between email and avatar file manually:

	$ microavatar server \
	  --addr localhost:8888 \
	  --cache-folder /var/cache/microavatar \
	  --email "me@email.com:/path/to/file.jpg"

So far,  I have this implementation running on avatars.zerokspot.com and set the respective SRV DNS entries accordingly for my various “@zerokspot.com” addresses:

	_avatars._tcp.zerokspot.com. 60 IN      SRV     10 100 80 avatars.zerokspot.com.
	_avatars-sec._tcp.zerokspot.com. 60 IN  SRV     10 100 443 avatars.zerokspot.com.

## Liberapay not the best example

After all this work, I set my profile on Liberapay to use LibAvatar in the hope that they’d do the DNS lookup I described above. Sadly, [they don’t](https://github.com/Liberapay/Liberapay.com/blob/56bb0c2d5859126f38810dbb24836ccea7484860/Liberapay/elsewhere/_base.py#L267). They just calculate the hash and then forward it to Libravatar.org. So me writing my own server and jumping through all these hoops to be a good federated avatar citizen seems to have been in vain. At least I had fun and learnt something new 😅

## Updates:

- **2020-05-27 13:17 +02:00:** Added a link to existing libraries.
