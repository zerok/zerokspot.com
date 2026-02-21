---
title: "Complex app-logic configuration with Rego"
tags:
- golang
- rego
- openpolicyagent
date: "2026-02-21T21:20:00+01:00"
---

Over the last couple of years I've written tons of little Go services that required part of their logic to be dynamically configurable. For more complex scenarios it might be time for an embedded scripting language like Lua. Especially when it comes to situations where some kind of document needs to be dynamically generated or a decision to be made based on user-defined input, there is an easier way: [OpenPolicyAgent](https://www.openpolicyagent.org/)'s [Rego language](https://www.openpolicyagent.org/docs/policy-language) and library.

(Disclaimer: Parts of the code examples were generated with Claude Code.)

## Example scenarios

Just to give you some examples for this:

- A user sends a request to a server and we need to decided based on the data if that request is allowed or not.
- Given a GitHub pull request define what labels it should receive based on the files that are modified.

Rego is a policy language and its tooling is pretty much optimized for the first scenario. What Rego actually generates when you query a policy, though, is a JSON document, so you can also use it to provide more complex answers to your input.

## Example implementation

So let's go with the the scenario where we want to get some labels that should be associated with a pull request. If a pull request has a title that starts with "fix", then it should also get a `type:fix`. Additionally, if the pull request modifies a file inside the `internal/auth` folder, then the label `component:auth` should also be set:

```rego
// Filepath: policies/pullrequest_enrichment.rego

package pullrequest_enrichment

labels contains "type:fix" if {
	startswith(input.title, "fix")
}

labels contains "component:auth" if {
	all_files := {f | some f in input.added_files} |
		{f | some f in input.changed_files} |
		{f | some f in input.deleted_files}
	some file in all_files
	startswith(file, "internal/auth/")
}
```

You can now use that policy with the `github.com/open-policy-agent/opa/v1/rego` package:

```go
import (
	"context"
	// ...

	"github.com/open-policy-agent/opa/v1/rego"
)


input := PullRequestInput{
	Title:        "fix: validate auth token expiry",
	AddedFiles:   []string{},
	ChangedFiles: []string{"internal/auth/token.go"},
	DeletedFiles: []string{},
}

ctx := context.Background()

// Create a prepared query that could be used in theory
// for multiple inputs:
preparedQuery, err := rego.New(
	rego.Query("data.pullrequest_enrichment.labels"),
	// The .rego file we created above is stored inside
	// a folder called "policies" so we can load it from
	// there:
	rego.Load([]string{"./policies/"}, nil),
).PrepareForEval(ctx)
if err != nil {
  // ...
}

// Now evaluate the input against the policy to receive
// a result set:
rs, err := pq.Eval(ctx, rego.EvalInput(input))
if err != nil {
	// ...
}
if len(rs) == 0 || len(rs[0].Expressions) == 0 {
	// No results, so no labels were returned.
}
// Resultsets are quite generic and so we need to do a bit
// of typecasting to got the results
items, _ := rs[0].Expressions[0].Value.([]any)
labels := make([]string, 0, len(items))
for _, item := range items {
	if s, ok := item.(string); ok {
		labels = append(labels, s)
	}
}
```

## Allowed?

For the scenario where we have a user request as input and would like to know if it's allowed or not, there is a helper for results that have the "allow" property which is just a boolean:

```go
rs, _ := pq.Eval(ctx, rego.EvalInput(input))
result := rs.Allowed()
```

These policies are basically configuration that you can deploy separately from your core application. These policies can get quite complex. The cool thing about Rego's tooling is that you can also create unit tests:

```rego
// Filepath: policies/pullrequest_enrichment_test.rego
package pullrequest_enrichment_test

import data.pullrequest_enrichment

test_labels_contains_type_fix_when_title_starts_with_fix if {
  "type:fix" in pullrequest_enrichment.labels with input as {
    "title": "fix: correct null pointer"
  }
}
```

After installing the OpenPolicyAgent CLI, you can then run these tests:

```
$ cd policies
$ opa test .
PASS: 1/1
```

## Part of my toolbox

Admittedly, using Rego as a document generator like it's done in the example for getting a list of labels is probably not part of its original mission, but Rego is useful for that and so many other scenarios that it has become part of my standard toolbox!
