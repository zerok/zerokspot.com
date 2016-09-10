---
date: '2005-01-09T12:00:00-00:00'
language: en
tags:
- python
title: Accessing a MySQL DB in Python
---


And another small intro into a MySQL API. This time we want to fetch the phpbb_config table within a small Python script.

-------------------------------



Because Python doesn't enforce try-except blocks like Java I will leave them away this time so that the example code stays better readable.

<h3>Requirements</h3>
* On Ubuntu/Debian you need the python-mysqldb package which should pull the right version of the library for your version of Python as well as all other dependencies you need for accessing a MySQL database with Python
* On Gentoo Linux you need the mysql-python package which should also install all necessary packages through the dependencies

<h3>How?</h3>
First step: Import the MySQLdb package and open a connection to the server:
<pre class="code">
import MySQLdb
conn = MySQLdb.Connection(host='192.168.0.1',user='myuser',passwd='mypassword',db='phpbb20')
</pre>

Now that we have the connection we want to query the database. The Connection objection already has a query() method but we will use a different way which will give us some additional options:
<pre class="code">cur = conn.cursor()</pre>
Here we receive a Cursor object from the Connection class which we can now use to execute our already known query:
<pre class="code">cur.execute("SELECT * FROM phpbb_config")</pre>

If the query was successful, we can fetch the result from it:
<pre class="code">res = cur.fetchall()</pre>

We don't need the connection to the database anymore so we can close the cursor and the database connection:
<pre class="code">cur.close()
conn.close()</pre>

Final step: Let's iterate through the result set:
<pre class="code">for row in res:
	print row</pre>

If you are following these steps, you should now see a tuple of tuples on your screen. That's because we used the "normal" Cursor class. But what if we want to have the rows stored in Dicts instead of Tuples? As the word "normal" should have indicated, there are some other subclasses of the BaseCursor class available and one of them stores the rows as Dicts. Simply replace the line where we got the cursor from the connection with this:
<pre class="code">cur = conn.cursor(cursorclass=MySQLdb.cursors.DictCursor)</pre>
If you then start your script again, you should see a listing of Dicts :-)

<h3>References</h3>
* <a href="http://dev.mysql.com/doc/mysql/en/MySQL_APIs.html">API guides on MySQL.com</a>
* The pydocs of the MySQLdb package