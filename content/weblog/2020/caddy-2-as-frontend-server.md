---
title: Caddy 2 as frontend server
date: "2020-05-21T07:31:46Z"
tags:
- caddy
- server
- hosting
- operations
- 100daystooffload
---

For the last week I’ve been running some of my sites using [Caddy 2](https://caddyserver.com/v2) as replacement for nginx. Over the years, nginx configs have grown and grown mostly due to TLS settings not having defaults following industry best practices and due to Let’s Encrypt becoming *the* way to get certificates.

With Caddy I get two benefits out of the box:

1. Shorter configuration files
2. No longer having to run additional tools for up-to-date TLS certificates

To give you an example, this is part of the configuration I have for zerokspot.com:

	  zerokspot.com {
	      root * /srv/www/zerokspot.com/www/htdocs
	      file_server
	      encode zstd gzip
	      log {
	          format json
	          output file "/srv/www/zerokspot.com/www/logs/access.json.log" {
	              roll_size 100MiB
	              roll_keep 20
	          }
	      }
	
	      header /sass/main.min.* Cache-Control max-age=31536000
	
	      route /api/* {
	          uri strip_prefix /api
	          reverse_proxy localhost:9999
	      }
	      route /webmentions/* {
	          uri strip_prefix /webmentions
	          reverse_proxy localhost:35080
	      }
	
	      header /.well-known/openpgpkey Access-Control-Allow-Origin "*"
	      header /.well-known/openpgpkey Content-Type "text/plain"
	      redir //index.xml https://zerokspot.com/index.xml permanent
	      redir /weblog/feed/ https://zerokspot.com/index.xml permanent
	  }

The biggest difference here is that there is no huge SSL section in it. Caddy simply always exposes hosts on port 80 and 443 using [Let’s Encrypt](https://letsencrypt.org/) certificates if not [instructed otherwise](https://caddyserver.com/docs/caddyfile/directives/tls).

The configuration is also much more compact in general thanks to pretty much every directive being able to be restricted to specific URLs using so-called [matchers](https://caddyserver.com/docs/caddyfile/matchers). One example for this is the following line:

	header /sass/main.min.* Cache-Control max-age=31536000

Here I say in a single line that all files with the prefix `/sass/main.min.` should also send a very long Cache-Control max-age header in the response. Nice and concise!

Something that I hadn’t expected was that Caddy also offers some nice logging features like producing access logs with JSON statements per entry. This is going to make parsing much easier for me once I find the time to work on a little analytics generator 😅

Caddy also supports dynamically changing the server configuration through an API but I haven’t used this feature yet.

The only downside I’ve found so far is that the package installation is a bit more manual compared with something like nginx or Apache HTTPD. That’s not really an issue for me, though, as I manage my servers using Ansible.
