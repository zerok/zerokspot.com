---
date: '2005-01-08T12:00:00-00:00'
language: en
tags:
- java
title: Connecting to a MySQL DB in Java
---


And here it is: Another small tutorial :-) This time we will build a small Java application that simply connects to a MySQL database and prints the conent of one of the database's tables.

-------------------------------



<h3>Requirements</h3>
* A MySQL server
* JDK
* <a href="http://dev.mysql.com/downloads/">JDBC drivers for MySQL databases</a>

<h3>How?</h3>
<h4>Loading the driver</h4>
First of all we'll have to tell Java, that there are some JDBC drivers for MySQL available which can be done by extending the classpath (These paths are for Gentoo Linux):
<pre class="code"> export CLASSPATH=$CLASSPATH:/usr/share/jdbc-mysql/lib/mysql-connector-java-3.0.15-ga-bin.jar</pre>

Now that Java should be able to find the driver we can try to load it:
<pre class="code">
import java.sql.*;
//import com.mysql.jdbc.Driver;
class Test{
	public static void main(String[] args){
		try {
			Class.forName("com.mysql.jdbc.Driver").newInstance();
		}
		catch(Exception e){
			System.out.println("Failed to load MySQL driver");
			System.exit(-1);
		}
	}
}
</pre>
As you can see, we don't import the driver directly but use some reflection to load the driver. According to the documentation this way is recommended because of some broken JavaVMs out there.

<h4>Establish a connection to the server</h4>
<pre class="code">
import java.sql.*;
//import com.mysql.jdbc.Driver;
class Test{
	public static void main(String[] args){
		try {
			Class.forName("com.mysql.jdbc.Driver").newInstance();
		}
		catch(Exception e){
			System.out.println("Failed to load MySQL driver");
			System.exit(-1);
		}

		String connstring = "jdbc:mysql://192.168.0.1:3306/phpbb20?user=myuser&password=mypassword";
		try {
			Connection conn = DriverManager.getConnection(connstring);
			System.out.println("Connection established");
			conn.close();
		}
		catch(SQLException e){
			System.out.println("Failed to open a DB connection");
			System.out.println("Message: "+e.getMessage());
		}

	}
}
</pre>
In this step we try to establish a connection to the database server. There are some Connection constructors from which I will use the one with only one String as parameter here. In this single string you put the options this way:
<pre>protocol:dbms://hostname:port/databasename</pre>
After the database name you can specifiy things like the username and the password to be used for this connection. If something goes wrong the Connection class will throw a SQLException.

After we have a connection, we can close it for now again.

<h4>To query a DB</h4>
JDBC offers quite a few ways to query a database, but I will just use the following ...
<pre class="code">
			Statement s = conn.createStatement();
			if (s.execute("SELECT * FROM phpbb_config")){
				System.out.println("Query successful");
			}
			else {
				System.out.println("Query failed");
			}
</pre>

If the query was successful, we can get the ResultSet and iterate through its content:
<pre class="code">
				ResultSet res = s.getResultSet();
				while(res.next()){
					System.out.println(res.getString("config_name")+" | "+res.getString("config_value"));
				}
</pre>
As you can see from the query, I use the phpbb_config table of phpBB 2.0.x as the test table. This table has two columns:
1. config_name
2. config_value

The code above will iterate through the whole result set as long as there's a next row and will print the table content in the format {config_name} | {config_value} . Both columns are of the VARCHAR type so they are mapped to the String class in Java.

<h4>The final code</h4>
Now we have ourselves a nice little Java application that queries the phpbb_config table and displays its content:
<pre class="code">
import java.sql.*;
//import com.mysql.jdbc.Driver;
class Test{
	public static void main(String[] args){
		try {
			Class.forName("com.mysql.jdbc.Driver").newInstance();
		}
		catch(Exception e){
			System.out.println("Failed to load MySQL driver");
			System.exit(-1);
		}
		
		String connstring = "jdbc:mysql://192.168.0.1:3306/phpbb20?user=myuser&password=mypassword";
		try {
			Connection conn = DriverManager.getConnection(connstring);
			System.out.println("Connection established");

			Statement s = conn.createStatement();
			if (s.execute("SELECT * FROM phpbb_config")){
				System.out.println("Query successful");
				ResultSet res = s.getResultSet();
				while(res.next()){
					System.out.println(res.getString("config_name")+" | "+res.getString("config_value"));
				}
			}
			else {
				System.out.println("Query failed");
			}

			conn.close();
		}
		catch(SQLException e){
			System.out.println("Failed to open a DB connection");
			System.out.println("Message: "+e.getMessage());
		}
	}
}
</pre>

<h3>Further reading</h3>
* <a href="http://java.sun.com/j2se/1.4.2/docs/api/java/sql/package-summary.html">java.sql documentation</a>
* <a href="http://dev.mysql.com/doc/connector/j/en/index.html">MySQL Connector/J developer documentation</a>