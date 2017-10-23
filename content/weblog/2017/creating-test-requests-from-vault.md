---
title: "Creating test-requests from Vault"
date: 2017-10-23T20:16:23+02:00
tags:
- vault
---

During the last couple of weeks I've been using Hashicorp's
[Vault](https://www.vaultproject.io/) more and more also for local development.
Among other things I've added credentials to most systems I'm interacting with
to the secret store. This way, I don't have those usernames and passwords
flying around as global environment variables or plain-text config files.

For tools that support reading their configuration from stdin I've written a
little wrapper function (for [Fish](https://fishshell.com/)) around
[consul-template](https://github.com/hashicorp/consul-template):

```
function vault-template --wraps=consul-template
    consul-template -vault-renew-token=false -once -dry -template $argv | tail -n +2
end
```

Now, let's say that we have a simple JSON request containing a login and a
password field that we would like to send to an authentication service. All we
need is to create a template for that JSON payload with a few placeholders:

```
{{"{"}}{{ with secret "secret/credentials/me" }}
    "login": "{{ .Data.login }}",
    "password": "{{ .Data.password }}"
{{end}}{{"}"}}
```

Those `{{ "{" }}` and `{{ "}" }}` are sadly necessary in order to get JSON to
work within Go's template package. All that's left to do after that is to
execute it :)

```
$ vault-template login-request.json.tmpl | http post https://localhost:8443/v1/login
```

I'm not yet sure, if this will be my final approach for handling secrets in
local configurations. Maybe I will ditch consul-template for a custom
Vault-client in order to gain a bit more flexibility with default values and
prefixes. That being said, for a 1.0 the current consul-template + Fish
solution has been immensely useful so far 😊
