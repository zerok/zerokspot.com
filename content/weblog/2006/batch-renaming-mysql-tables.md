---
date: '2006-06-25T12:00:00-00:00'
language: en
tags:
- mysql
- ruby
title: Batch-renaming MySQL tables
---


Ever felt the need to batch-rename quite a few tables in a MySQL database? Sadly phpMyAdmin doesn't offer renaming of _all_ the tables in a specific database. So I first started renaming them manually using the commandline interface but got soon quite tired of the whole process. 

A quite ugly but IMO pragmatic workaround for this limitation is to build a small query file holding all the RENAME TABLE statements you want. To generate this query listing first of all get a list of all the tables you want to rename.



-------------------------------



Since I'm lazy all I want to do is a `show tables;` and then copy this output to some variable inside of some script. Well, and then this script should generate me the statement list I will execute with the MySQL commandline client:

<pre class="code">
OLD_PREFIX = &quot;&quot;
NEW_PREFIX = &quot;newprefix_&quot;
data=&lt;&lt;-EOD
| aggregator_category_item          |
| aggregator_feed                   |
| aggregator_item                   |
| search_total                      |
| sequences                         |
EOD
tables = []
data.each_line do |line|
  if line=~/(#{OLD_PREFIX}\w+)/
    puts &quot;RENAME TABLE #{$1} TO #{NEW_PREFIX}#{$1};&quot;
  end
end
</pre>

The data variable here holds the output I got from `show tables;`. The output of this script will look like this:

<pre class="code">
RENAME TABLE aggregator_category_item TO newprefix_aggregator_category_item;
RENAME TABLE aggregator_feed TO newprefix_aggregator_feed;
RENAME TABLE aggregator_item TO newprefix_aggregator_item;
RENAME TABLE search_total TO newprefix_search_total;
RENAME TABLE sequences TO newprefix_sequences;
</pre>

Variable regex are ugly ... but sometimes they make things sooooo easy ;)