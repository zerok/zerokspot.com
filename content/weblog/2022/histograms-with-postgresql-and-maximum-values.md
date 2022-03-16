---
title: Histograms with PostgreSQL and maximum values
date: "2022-03-16T17:32:12+01:00"
tags:
- postgresql
- sql
---

A couple of days ago I wanted to implement a basic histogram for a range of values I had in a PostgreSQL database. I basically wanted to create 10 equally sized buckets with the first bucket starting at the lowest value of my dataset and the last bucket ending with the highest value.

```sql
CREATE TABLE sample_data (value double precision);
INSERT INTO sample_data (value) VALUES
    (1.0),
    (2.0), (2.2),
    (3.0),
    (4.0),
    (5.0),
    (6.0),
    (7.0),
    (8.0), (8.5),
    (9.0),
    (10.0);

```

I’ve already sorted the values in a way to show you what kind of histogram I’d expect coming out in the end.

For things like that, PostgreSQL offers the `width_bucket` function:

```sql
width_bucket(value, min, max, num_buckets)
```

You basically give it the minimum value, the maximum value, the number of buckets to consider and it will then tell you in which bucket (1-indexed) your given value will show up.

So I started with this query:

```sql
SELECT
    width_bucket(
        value,
        (SELECT min(value) FROM sample_data),
        (SELECT max(value) FROM sample_data),
        10) as bucket,
    count(*)
FROM sample_data
GROUP BY bucket
ORDER BY bucket;

```

… which produces the following result:

```
 bucket | count
--------+-------
      1 |     1
      2 |     2
      3 |     1
      4 |     1
      5 |     1
      6 |     1
      7 |     1
      8 |     1
      9 |     2
     11 |     1
(10 rows)
```

OK, this looks mostly fine except for the very last row: Why is there all of a sudden an 11th bucket and why are there 2 items in the 9th bucket and not in the 8th? For some reason I then just stackOverflow’d, set the `num_buckets` parameter to 9 and the result looked:

```
 bucket | count
--------+-------
      1 |     1
      2 |     2
      3 |     1
      4 |     1
      5 |     1
      6 |     1
      7 |     1
      8 |     2
      9 |     1
     10 |     1
(10 rows)
```

But somehow I couldn’t explain that `9` to a colleague when he ran into an issue with that calculation two days later. For some reason, one of his results turned up in the wrong bucket which caused some other issues. Somehow my test-data was different from his and so I never noticed that in my integration tests 😔

Some reading and experimenting later and I understood that the max value above was actually not part of the last bucket anymore! `width_bucket` places values that are outside of the last bucket into one with the number `last_bucket + 1`. When the number of buckets was 9, that fictional bucket number was 10 and so it looked right to me but actually the bucket boundaries were off (since everything was calculated based on 9 instead of 10 buckets).

In the end, I solved it like this:

```sql
SELECT width_bucket(
  value,
  (select min(value) from sample_data),
  (select max(value) from sample_data) + 0.01,
  10) as bucket
  count(*)
FROM sample_data
ORDER BY bucket
GROUP BY bucket;
```

I set the maximum value of the bucket calculation to a tiny bit higher than my maximum value. In my case `0.01` is subtle enough but your milage may vary here. 

So, when using that form of the `width_bucket` function keep in mind, that the maximum parameter *is not* part of the last bucket! Big thanks to [Ellis Valentiner](https://ellisvalentiner.com/post/discretizing-data-in-postgres-with-width-bucket/) who explains the various forms on his blog a couple of years ago 🙂 
