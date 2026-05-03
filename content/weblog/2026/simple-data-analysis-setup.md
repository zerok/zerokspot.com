---
date: "2026-05-03T11:30:00+02:00"
incoming:
- url: https://chaos.social/@zerok/116510092124723114
tags:
- python
- jupyter
- data
- duckdb
title: Simple data analysis setup
---

Over the last couple of months I've had to do a bit more number-crunching than usual in order to write design docs. In a recent example I had to work with a datasource that was slow to query (GitHub's REST API) and eventually needed to work with 100,000+ data points.

In order to demonstrate my setup there, I'll work on a slightly contrived example: Get data about all PRs in the [grafana/grafana](https://github.com/grafana/grafana/) project within the last 6 months. Then try to find out which of these include documentation *and* code changes (basically anything but changes to the docs folder).

## Core setup

For me, the most convenient way to think about larger amounts of data is using [Jupyter](https://jupyter.org/) notebooks. I have a local Jupyter Lab set up where I collect all my analysis into separate folders.

```sh
cd notebooks
uv init
uv add jupyterlab
uv add ipywidgets
uv add pygithub
uv run jupyter lab
```

## Setup

Before being able to access GitHub data, I need to log in. My [gh cli](https://cli.github.com/) is already set up, so I'm going to use that instead of creating yet another PAT that would be just lying around on my system:

```
from github import Github
from github.Auth import Token

token = !gh auth token
token = token[0]
auth = Token(token)

# Create a GitHub client using that token and
# with a page-size of 100 (default: 30):
gh = Github(auth=auth, per_page=100)
```

## Data gathering

Each analysis gets its own `data` folder. Even if I have to reproduce something with a tiny dataset, I use this structure.

The data folder is important since it allows me to not only keep my analysis reproducible but also speeds up iterations. Sometimes getting the data is actually the slow part that can take (depending on the data source) minutes, hours, or days. During the gathering phase I try to store as many chunk files as possible. Here I like to use [JSONLines](https://jsonlines.org/) files (one JSON document per line) where I just open the output file in append-mode. If my gathering script runs into an issue, I lose at most the last entry. At the same time, I don't generate thousands of files as would be the case if I put one datapoint into its own file.

Especially for long runs, it helps to have some kind of progress bar available that is more than just continuously printing `n/45678 elements processed`. I've stumbled upon the [tqdm](https://github.com/tqdm/tqdm) package a while back and am now using it pretty much everywhere:

```python
# uv add tqdm

import datetime
from tqdm.notebook import tqdm
import pathlib
import math
import json

data_folder = pathlib.Path('data')
data_folder.mkdir(exist_ok=True)

range_start = datetime.datetime(2025, 11, 1, tzinfo=datetime.UTC)
range_end = datetime.datetime(2026, 5, 1, tzinfo=datetime.UTC)

repo = gh.get_repo('grafana/grafana')

# Unfortunately, the GH API for pull requests doesn't allow 
# for time-based filtering, so we actually have to iterate
# through everything 
pull_requests = repo.get_pulls(state='closed')

with tqdm(total=pull_requests.totalCount) as progress:
    for pr in pull_requests:
        progress.update()
        if not (range_start <= pr.created_at < range_end):
            continue
        # We don't want a single file per pull request and also not everything in one.
        # Let's split it up by pr.number / 1000
        state_file = data_folder / f'{math.floor(pr.number / 1000)}000.jsonl'
        with state_file.open('a+') as fp:
            files = pr.get_files()
            json.dump({
                'title': pr.title,
                'id': pr.id,
                'number': pr.number,
                'user_login': pr.user.login,
                'user_type': pr.user.type,
                'files': [file.filename for file in files],
                'body': pr.body,
                'created_at': pr.created_at.isoformat(),
                'merged_at': pr.merged_at.isoformat() if pr.merged_at else None,
            }, fp)
            fp.write("\n")

```

This snippet can take hours and hours (around 20 on my previous run). The problem is that GitHub doesn't let me filter PRs based on a date range. That's only available view the search API which is limited to 1000 items.

## Data analysis

Now that I have around 10 JSONL files flying on my system, I can feed them into a [DuckDB](https://duckdb.org/) and treat everything like one large database table!

```python
# uv add duckdb

import duckdb

duckdb.sql("SELECT * FROM 'data/*.jsonl'")\
    .to_view('pull_requests')
duckdb.sql("SELECT count(*) FROM pull_requests")\
    .show()
```

This gives me the following output:

```
┌──────────────┐
│ count_star() │
│    int64     │
├──────────────┤
│         8292 │
└──────────────┘
```

The dataset isn't all that large so I can just put it into a view. The count-query will just do a `SELECT * FROM ...` internally without caching. If the dataset were larger, I'd use `to_table` instead to properly materialize the data.

Now that all the JSON files are in a database, I can find out how many of the PRs created over those 6 months actually contained documentation. I'm going with a naive approach here and just consider anything that touched a file in the `docs/` folder to be a documentation update:

```python
# Creating a view again since we will need the filter a couple
# more times:
duckdb.sql("""
    SELECT *
    FROM pull_requests
    WHERE len(array_filter(files, lambda file: starts_with(file, 'docs/'))) > 0
""").to_view('docs_pull_requests')

duckdb.sql("""
SELECT 
    count(*) as docs_count,
    100 * docs_count / (SELECT count(*) FROM pull_requests) as docs_percentage
FROM docs_pull_requests
""").show()
```

Turns out, 1033 PRs included docs changes. That's around 12.46%. Who created all these PRs?

```python
duckdb.sql("""
	SELECT user_login, count(*) as count
	FROM docs_pull_requests
	GROUP BY user_login
	ORDER BY count DESC
	LIMIT 10
""").show()
```

Most came from an the grafana-delivery-bot, which makes sense since it is used to backport fixes into specific release branches:

```
┌───────────────────────────┬───────┐
│        user_login         │ count │
│          varchar          │ int64 │
├───────────────────────────┼───────┤
│ grafana-delivery-bot[bot] │   291 │
│ urbiz-grafana             │    86 │
│ imatwawana                │    84 │
│ lwandz13                  │    61 │
│ jtvdez                    │    40 │
│ JohnnyK-Grafana           │    22 │
│ irenerl24                 │    19 │
│ macabu                    │    19 │
│ stephaniehingtgen         │    19 │
│ knylander-grafana         │    15 │
└───────────────────────────┴───────┘
```

Now back to the original question: "How many PRs contain docs and code changes?". Again, I'm going with a naive approach here and consider any modifications to `.json`, `.ts`, or `.go` files to be code changes:

```sql
ITH code_pull_requests AS (
    SELECT 
        *
    FROM
        pull_requests
    WHERE
        len(array_filter(files, lambda file: ends_with(file, '.json') or ends_with(file, '.ts') or ends_with(file, '.go'))) > 1
)
SELECT
    count(*) docs,
    count(*) FILTER (id IN (SELECT id FROM code_pull_requests)) AS docs_and_code,
    100 * docs_and_code / docs AS percentage
FROM 
    docs_pull_requests
```

```
┌───────┬───────────────┬────────────────────┐
│ docs  │ docs_and_code │     percentage     │
│ int64 │     int64     │       double       │
├───────┼───────────────┼────────────────────┤
│  1033 │           224 │ 21.684414327202322 │
└───────┴───────────────┴────────────────────┘
```

That's pretty much it! Jupyter + tqdm + DuckDB + JSONLines has become my personal dream team for any kind of quantitive analysis. While I never dreaded writing a design doc, now I consider the analysis part even fun. With every single of them I've learnt something new and awesome which made my setup better!

## PS.: Support restarts

I work primarily on a laptop and tend to move around a bit every day. This means that my data gathering steps need to support things like me closing my laptop, going offline, and resuming work somewhere completely different. Normally, you'd use some kind of state file for this. In this scenario I can do a little hack, though, since DuckDB is so fast when it comes to ingesting data:

```python
seen_ids = set([])
try:
    seen_id = set([row[0] for row in duckdb.sql("SELECT number FROM 'data/*.jsonl'").fetchall()])
except:
    pass
```

With that in place, I can check in the PR loop if I have to also fetch the files again or not. I still have to loop through every pull request page, but without the files query it gets much faster. Thanks to that, I can cancel and restart a gathering run and only lose a few hours ... not tens of hours 😉

For other API endpoints that allow more detailed querying, that window gets even shorted to the point where I just cancel and restart all the time.
