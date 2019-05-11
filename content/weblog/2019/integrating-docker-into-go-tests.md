---
title: "Integrating Docker into Go integration tests"
date: 2019-05-11T17:08:18+02:00
tags:
- golang
- testing
- docker
---

Testing applications that integrate with external services like a
database is always a bit complicated. Mocking those dependencies
usually results in the most performant solution but always comes with
the risk that the mock doesn't behave like the real thing. Docker made
this situation much simpler for at least some external dependencies.

Recently, I needed to test some code that relied on a PostgreSQL
database being available. Mocking would have been very complicated and
so I decided to just launch a PostgreSQL server inside a docker
container, run the migrations, prepare the relevant database structure
and then execute my tests. While setting all that up I ran into the
[dockertest][] library by [ORY](https://www.ory.sh/) which makes
launching Docker containers as part of your Go test suite relatively
straight forward:

```go
// +build integration

package integration_test

import (
        "context"
        "database/sql"
        "fmt"
        "path/filepath"
        "testing"
        "time"
        
        sqlbuilder "github.com/huandu/go-sqlbuilder"
        "github.com/ory/dockertest"
        "github.com/stretchr/testify/require"
        "project.com/pkg/db"
)

func dbURLFromResource(r *dockertest.Resource) string {
	port := r.GetPort("5432/tcp")
	return fmt.Sprintf("postgres://user:password@localhost:%s/project?sslmode=disable", port)
}

func TestSomething(t *testing.T) {
	// Create test database inside a Docker container
	ctx := context.Background()
	pool, err := dockertest.NewPool("")
	require.NoError(t, err)
	res, err := pool.Run("postgres", "11.2-alpine", []string{
		"POSTGRES_PASSWORD=password",
		"POSTGRES_USER=user",
		"POSTGRES_DB=project",
	})
	
	// The db.Setup function waits for the DB server
	// to be completely available and then also
	// runs some migrations
	conn, err := db.Setup(ctx, func(c *db.Configuration) {
		c.URL = dbURLFromResource(res)
	})
	require.NoError(t, err)
	require.NotNil(t, conn)
	defer pool.Purge(res)
	
	// Now let's set up some test data:
	ib := sqlbuilder.PostgreSQL.NewInsertBuilder()
	query, args := ib.InsertInto("users").Cols("username").Values("testuser").Build()
	_, err = conn.ExecContext(ctx, query, args...)
	require.NoError(t, err)

	tx, err := conn.BeginTx(ctx, nil)
	require.NoError(t, err)
	defer tx.Rollback()

	// And here come the tests:
	// ...
}
```

dockertest abstracts all interactions with the underlying Docker API
through a pool (which handles the API client) and resources (which
represent the launched containers).

From each resource you can then retrieve things like the local port
bindings so that you can interact with the services in them. This is
pretty much what the `dbURLFromResource` function in the example above
does: It looks up the metadata of the database container and returns a
URL that I can then use for `sql.Open`.

dockertest also offers a handful other feature like containers that
are automatically removed after a certain timeout but so far I've only
used the setup shown in the example above. Combined with [testify][]
and [sqlbuilder][] and [migrate][] for running schema migrations I
think I have a quite powerful toolbox for handling PostgreSQL (or
other DBs) integration tests 🙂

[testify]: https://github.com/stretchr/testify
[sqlbuilder]: https://github.com/huandu/go-sqlbuilder
[migrate]: https://github.com/golang-migrate/migrate
[dockertest]: https://github.com/ory/dockertest
