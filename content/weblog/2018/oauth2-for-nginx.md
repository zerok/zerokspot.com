---
title: "OAuth2 for Nginx"
date: 2018-06-09T11:55:58+02:00
tags:
- ops
- sysadmin
- nginx
- demo
---

Recently, I needed a way to put authentication in front of an nginx instance
that would allow logging in through oauth2/openid connect. Luckily, a
[coworker](https://github.com/kautsig) of mine had already done something
similar so I knew what components I'd need:

* [oauth2\_proxy](https://github.com/bitly/oauth2_proxy) by bitly
* [dex](https://github.com/coreos/dex) by CoreOS

I originally wanted to also create a little test-setup inside docker-compose.
Sadly, that would have required installing some sidecar components into the
nginx container or something similar. But given that my target setup was
running inside a VM anyway where I could use nginx to just proxy other local
services, I went with the following setup. 

OAuth2 will be provided by dex in this little setup simply because we can
easily "fake" a user account using its `staticPasswords` setting. You can start
a local dex instance with the following command:

```
$ /usr/bin/docker run --rm --name dex -p 5556:5556 \
  -v /var/lib/dex:/var/lib/dex \
  -v /etc/dex/dex.yaml:/etc/dex/dex.yaml \
  quay.io/coreos/dex:v2.10.0 serve /etc/dex/dex.yaml
```

This assumes that there is a `/etc/dex/dex.yaml` file with the following
content:

```
---
issuer: http://{{ fqdn }}/dex
web:
  http: 0.0.0.0:5556
storage:
  type: sqlite3
  config:
    file: /var/lib/dex/dex.db
enablePasswordDB: true
staticClients:
  - id: "oauth2_proxy"
    redirectURIs:
      - "http://{{ fqdn }}/oauth2/callback"
    name: "OAuth Proxy"
    secret: "{{ dex_client_secret }}"
staticPasswords:
- username: "{{ team_username }}"
  email: "{{ team_email }}"
  hash: "{{ team_password_bcrypted }}"
  userID: "08a8684b-db88-4b73-90a9-3cd1661f5466"
```

We've also defined a static client here to allow oauth2-proxy to be able to
connect to dex.

On the same host, start oauth2-proxy pointing to this dex installation as
backend:

```
/usr/bin/docker run --rm --name jenkins-oauth2-proxy -p 4180:4180 kautsig/oauth2-proxy:latest \
  -client-id oauth2_proxy \
  -client-secret "{{ dex_client_secret }}" \
  -cookie-secret "{{ oauth2_cookie_secret }}" \
  -email-domain "*" \
  -provider "oidc" \
  -redirect-url "http://{{ fqdn }}/oauth2/callback" \
  -oidc-issuer-url "http://{{ fqdn }}/dex" \
  -cookie-secure=false \
  -set-xauthrequest \
  -http-address "0.0.0.0:4180"
```

Now, all that's left is the nginx instance where we want to have the authentication:

```
server {
    listen 80 default_server;
    listen [::]:80 default_server;

    location /dex/ {
        proxy_pass http://localhost:5556/dex/;
        proxy_set_header   Host              $host;
        proxy_set_header   X-Real-IP         $remote_addr;
        proxy_set_header   X-Forwarded-For   $proxy_add_x_forwarded_for;
        proxy_set_header   X-Forwarded-Proto $scheme;
    }
    location /oauth2/ {
        proxy_pass http://localhost:4180/oauth2/;
        proxy_set_header   Host              $host;
        proxy_set_header   X-Real-IP         $remote_addr;
        proxy_set_header   X-Forwarded-For   $proxy_add_x_forwarded_for;
        proxy_set_header   X-Forwarded-Proto $scheme;
    }

    location / {
        auth_request /oauth2/auth;
        error_page 401 = /oauth2/start;
        # If I want, I can now forward the auth information to a proxied
        # service.
        auth_request_set $user $upstream_http_x_auth_request_user;
        proxy_set_header X-Forwarded-User $user;
    }
}
```

By now you've probably noticed all those `{{ variable_name }}` placeholders in
the code-samples above. In order to test this setup I've created a little
[Vagrant](https://www.vagrantup.com/) box that uses
[Ansible](https://www.ansible.com/) for the provisioning. You can find the
result on [Gitlab](https://gitlab.com/zerok/nginx-oauth2-demo).

Running `vagrant up` will give you a demo server on 192.168.50.4 where you can
log in using `team@team.com` as e-mail and `team` as password.

Please note that this is just a basic configuration. Cleanup etc. have been
left out in order to keep the example itself concise.
