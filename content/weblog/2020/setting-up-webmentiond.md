---
title: Setting up webmentiond
date: "2020-06-14T12:36:00+02:00"
tags:
- webmentiond
- tutorial
- indieweb
- webmention
- 100daystooffload
---

Since [Yarmo asked this morning](https://fosstodon.org/@yarmo/104341371114206112) about how to use [webmentiond](https://github.com/zerok/webmentiond) behind a proxy I noticed that I had completely forgotten to provide a proper getting-started guide. I’m not yet sure how I’ll organise documentation for the project in the long run so I’ll just give you a quick tutorial here using my own setup as example 🙂

## Goal

The goal of this guide is that you have a webmentiond instance running on your server (in this example `yoursite.com`), can log into the management interface, and other people can discover your webmention endpoint on your website.

## Environment/requirements

In my own setup I use [Caddy 2](https://caddyserver.com/) as proxy server but you can use pretty much anything there. The only really hard requirements of webmentiond are that you have **Docker** running on your server and that your server can connect to an **SMTP server** (I really like the service offered by [Postmark](https://postmarkapp.com/)) in order for webmentiond to send out login/authentication tokens via e-mail.

In our case, webmentiond should be made available on `https://yoursite.com/webmentions/` and I can log into its admin interface through the fictional email address `login@yoursite.com`.

## Step 1: Setting up webmentiond as service

Since I use systemd to handle pretty much all services on my services, let’s also use it for webmentiond. The service will be run as the user `webmentiond` and store all its data into `/var/lib/webmentiond` belonging to that user:

	$ adduser --home /var/lib/webmentiond webmentiond
	
	# Get the UID of the newly created user:
	$ id webmentiond

Next, I’d suggest pulling the `zerok/webmentiond:latest` image in order to make sure that Docker is set up properly:

	$ docker pull zerok/webmentiond:latest

Finally, you have to create a service definition (i.e. `/etc/systemd/system/webmentiond.service` with the following content:

	[Unit]
	Description=Webmentiond
	After=network-online.target
	StartLimitInterval=0
	
	[Service]
	ExecStart=/usr/bin/docker run --rm \
	    -e "MAIL_USER=..." \
	    -e "MAIL_PASSWORD=..." \
	    -e "MAIL_HOST=..." \
	    -e "MAIL_PORT=..." \
	    -e "MAIL_FROM=no-reply@yoursite.com" \
	    -v /var/lib/webmentiond:/data \
	    -p 35080:8080 \
	    -u UID_OF_WEBMENTIOND_USER \
	    zerok/webmentiond:latest \
	    --addr 0.0.0.0:8080 \
	    --allowed-target-domains yoursite.com \
	    --auth-jwt-secret SOME_RANDOM_SECRET_STRING \
	    --auth-admin-emails login@yoursite.com \
	    --public-url https://yoursite.com/webmentions
	Restart=always
	RestartSec=5
	
	
	[Install]
	WantedBy=multi-user.target
	

Once that file is in place, start the service:

	$ systemctl daemon-reload
	$ systemctl enable webmentiond
	$ systemctl start webmentiond

Now check if the service was able to start up:

	$ journalctl -f -u webmentiond
	Jun 14 09:50:09 ubuntu-512mb-fra1-01 systemd[1]: Started Webmentiond.
	Jun 14 09:50:10 ubuntu-512mb-fra1-01 docker[70940]: 9:50AM INF UI path served from /var/lib/webmentiond/frontend
	Jun 14 09:50:10 ubuntu-512mb-fra1-01 docker[70940]: 9:50AM INF Listening on 0.0.0.0:8080...

If you see something else, please make sure that you’ve replaced all those placeholders in the service file 🙂

## Step 2: Update reverse proxy config

In order to make webmentiond available through `https://yoursite.com/webmentions` I’ve added the following lines to the host configuration in my Caddyfile:

	# Prevent people from grabbing the exposed Prometheus
	# metrics:
	respond /webmentions/metrics 404
	
	# Forward /webmentions/*:
	route /webmentions/* {
	    uri strip_prefix /webmentions
	    reverse_proxy localhost:35080
	}

Now the UI should be available through `https://yoursite.com/webmentions/ui/`:

<figure><img src="/media/2020/Screenshot%202020-06-14%20at%2012.05.28.png"><figcaption></figcaption></figure>

## Step 3: Try to log in

Now that you have the UI available, try to log in using the email you set in the service definition (in this case `login@yoursite.com`). You should receive a login token within the next minute or so that you can redeem on the authentication page linked to from the login page. If you didn’t receive a mail, make sure your email settings are correct and that the mail wasn’t flagged as spam or something like that.

## Step 4: Link to the /receive/ endpoint

In order for folks to be able to actually send you mentions, they have to know where to send them. The workflow goes something like this:

1. Another blog post with the URL `https://a.com/post` mentions `https://yoursite.com/post`.
2. The server at a.com (or another service altogether) checks `https://yoursite.com/post` looking for a link-element in the markup that looks like this: `<link rel="webmention" href="https://yoursite.com/webmentions/receive">` .
3. It finds it, it will send a simple HTTP request to it indicating that `https://a.com/post` mentioned `https://yoursite.com/post`.

In our case, let’s make sure that we have a working receive endpoint:

	$ curl -i https://yoursite.com/webmentions/receive
	HTTP/2 405
	[...]

Looking good 🙂

Now you have to add the following line to your blog’s head-section:

	<link rel="webmention" href="https://yoursite.com/webmentions/receive">

With this done, people should be able to send you mentions 🙂 One thing, though: Any mention that is sent to the receive-endpoint is first checked for validity (i.e. that the source of the mention really actually links to its target) and only then does it show up in the UI. Once it’s there, you have to explicitly *approve* a mention before it can be shown on your website. This is there in order to prevent people abusing your blog as link-heaven.

## Step 5: Display mentions

Webmentiond also comes with a little widget that you can embed in your website for rendering mentions:

	<div class="webmentions webmentions-container"
	    data-endpoint="https://yoursite.com/webmentions"
	    data-target="https://yoursite.com/url/to/post"></div>
	<script src="https://yoursite.com/webmentions/ui/dist/widget.js"></script>
	

This *should* be it. I’m pretty sure that I’ve forgotten a thing or two or that something is completely unintelligible so please let me know 😅
