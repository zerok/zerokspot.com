---
title: "A maintained YAML library for Go again!"
date: "2026-02-07T18:30:00+01:00"
tags:
- golang
- yaml
- opensource
---

For 14 years Gustavo Niemeyer has maintained the [de-facto standard library](https://github.com/go-yaml/yaml) for working with YAML in Go. As is so often the case, life just happened and so he eventually [marked the repository of go-yaml as unmaintained](https://github.com/go-yaml/yaml/blob/944c86a7d2/README.md). After a bit of uncertainty OpenSource worked: [The library was forked](https://github.com/yaml/go-yaml) and is now managed by the folks who also own the YAML specification and so of the original downstream users.

The “old” v1, v2, and v3 branches have been frozen, to only receive security fixes, ensuring an easy upgrade from `gopkg.in/yaml.vX`. All the new stuff is happening in v4 (in the `main` branch)!

## New configuration API

The biggest change of v4 is a [new configuration API](https://github.com/yaml/go-yaml/pull/212) that should fit better with the YAML vocabulary. While `Marshal` and `Unmarshal` are very common in Go libraries, [PyYAML](https://pyyaml.org/wiki/PyYAMLDocumentation) and others primarily us `load` and `dump`. So that’s what is now also in `‌go.yaml.in/yaml/v4` (alongside the “old” API):

```golang
# func Load(in []byte, out any, opts ...Option) error
var data []Datum
err := yaml.Load(
	someBytes,
	&data, 
	// To load all documents from the input
	yaml.WithAllDocuments(true),
)
//...

# func Dump(in any, opts ...Option) (out []byte, err error)
output, err := yaml.Dump(
	data,
	yaml.WithExplicitStart(true),
	yaml.WithExplicitEnd(true),
	// If the input is a slice, multiple documents
	// will be created in the output
	yaml.WithAllDocuments(true),
)
//...

```

For streaming from `io.Reader` and to `io.Writer` objects, there are replacements for `Decoder` and `Encoder` with `Loader` and `Dumper`.

For more details see the [v3-to-v4 migration guide](https://github.com/yaml/go-yaml/blob/main/docs/v3-to-v4-migration.md#recommended-use-new-api). V4 is currently in the pre-release phase with the 4ths release candidate.

## Adoption

While the fork “started” only about a year ago, a lot of libraries have already switched to it. Most prominently probably [Kubernetes](https://github.com/kubernetes/kubernetes/issues/132056) and [Prometheus](https://github.com/prometheus/common/pull/834). So... I should finally get my act together and update my own stuff. While I’m at it, I can also go to the new API 😂
