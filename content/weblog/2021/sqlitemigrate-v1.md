---
title: SQLiteMigrate v1.0.0
date: "2021-01-30T17:31:26+01:00"
tags:
- golang
- library
- sqlite
- 100daystooffload
incoming:
- url: https://chaos.social/@zerok/105645672755483959
---

During a [GoGraz](https://gograz.org) meetup somewhen in 2019 [Matthias](https://twitter.com/matthias_endler/) mentioned that SQLite had a pragma (basically an attribute of the database) called `user_version` which might be used for storing the “version” of a database. Could this perhaps be useful for a migration library to store at what version the database scheme was at?

Sadly, I couldn't find any migration library for Go that used that feature and so I just wrote my own as part of the Geotrace project and eventually moved that into it's own project on GitHub called [SQLiteMigrate](https://github.com/zerok/sqlitemigrate).

Usage is pretty simple:

	# You create a registry where your migrations should live.
	reg := mig.NewRegistry()
	
	# New, create a migration which consists of SQL statements
	# for upgrading and downgrading a scheme.
	reg.RegisterMigration([]string{
		`CREATE TABLE users (id integer primary key autoincrement)`,
	}, []string{})
	
	db, _ := sql.Open("sqlite3", "test.sqlite")
	defer db.Close()
	
	# Finally, you apply the registry to a SQLite database:
	if err := reg.Apply(ctx, db); err != nil {
		log.Fatal("Failed to apply migration: %s", err.Error())
	}

So far I've only used that within Geotrace but I want to also have it in webmentiond before that project hits 1.0.0!

Anyway, I haven't had a new to touch the code in quite a while so I just wanted to finally put a v1.0.0 on it. That's what I've now finally done yesterday afternoon! I don't anticipate any large API changes on that tiny project as it works for my own use-cases for now. Perhaps this is also something that might be useful to you, too 😊
