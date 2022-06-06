---
title: GraphQL-support in OpenPolicyAgent
date: "2022-06-06T12:39:10+02:00"
tags:
- openpolicyagent
- features
---

Since [0.41](https://github.com/open-policy-agent/opa/releases/tag/v0.41.0), OPA now also has some [built-in functions for working with GraphQL](https://www.openpolicyagent.org/docs/latest/policy-reference/#graphql) requests! These allow you to parse requests, parse schemas, and verify that a request adheres to a schema:

- `graphql.is_valid(query, schema)`
- `graphql.parse(query, schema)`
- `graphql.parse_and_verify(query, schema)`
- `graphql.parse_schema(schema)`
- `graphql.parse_query(query)`

With this you can now, for instance, enforce that queries have a specific argument before they even reach the actual GraphQL server for processing or that they don’t go beyond a certain complexity!

Let’s say, I have a tiny schema that expose a `user(id:String)` query and I want to enforce that the ID is only allowed to be `myname`, then I could do something like this using those new functions:

```Rego
package demopolicy.gql

schema := `
type User {
    id: String
}
type Query {
    user(id: String): User
}
`

# Allow only queries for the user "myname"
allow {
	parsed := graphql.parse_and_verify(input.query, schema)
	is_valid = parsed[0]
	is_valid

	query := parsed[1]

	# There query is an object with one key: "Operations"
	op := query.Operations[_]
	op.Operation == "query"
	selection := op.SelectionSet[_]
	selection.Alias == "user"
	selection.Arguments[0].Name == "id"
	selection.Arguments[0].Value.Raw == "myname"
}
```

And for completeness’ sake, here’s the testing code:

```Rego
package demopolicy.gql

test_gql_parsing {
	allow with input.query as "query {user(id:\"myname\") {id}}"
	not allow with input.query as "query {user(id:\"someone-else\") {id}}"
}
```
