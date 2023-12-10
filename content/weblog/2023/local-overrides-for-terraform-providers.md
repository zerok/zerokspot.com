---
date: "2023-12-10T21:56:11+01:00"
incoming:
- url: https://chaos.social/@zerok/111558140299027220
tags:
- terraform
- til
title: Local overrides for Terraform providers
---

A couple of days ago I had the need to debug a [Terraform](https://www.terraform.io/) provider. After some digging around I learnt that you can define local overrides in a Terraform confirmation where you tell it to search for provider binaries in a local directory without requiring any checksums.

In this particular case I wanted to check an execution path in the [GitHub provider](https://github.com/integrations/terraform-provider-github) and so I added some debug statements, built it, and then I wanted to use it in a test project. Terraform allows that by defining custom [provider installation](https://developer.hashicorp.com/terraform/cli/config/config-file#provider-installation) paths to to a global [configuration file](https://developer.hashicorp.com/terraform/cli/config/config-file) called `.terraformrc`:

```
provider_installation {
  dev_overrides {
    "integrations/github" = "/Users/zerok/src/github.com/integrations/terraform-providers-github"
  }

  direct {}
}
```

I don't want to have that globally defined, but only for that particular test project. Locally, there is also the option to define a different file using the `TF_CLI_CONFIG_FILE` environment variable

```
export TF_CLI_CONFIG_FILE=/Users/zerok/testproject/dev.tfrc
```

When I then run `terraform plan` in the test project, it will print a warning that a local override is active and then use my locally compiled provider.
