---
date: '2007-11-14T12:00:00-00:00'
language: en
tags:
- postgresql
- trac
title: Multiple Tracs with the Same User Database
---


For my master thesis I'm working with a couple of other people on quite a 
big project which consists of multiple sub-projects. To get some order into 
the whole process and documentation we are about to use [Trac](http://trac.edgewall.org/), the free 
project management tool by the folks at edgewall. It simply let's you handle 
one project per instance with its own SVN repository and user base. But what 
if you want to have multiple instances with the same users but just different 
permissions per instance?

-------------------------------

At the first glance, Trac won't do that for you. But if you start looking at 
what backends it supports, there are quite a couple of nifty tricks you can 
apply to get it to work without relying on plugins or forks out there:

1.  The authentication can be handled through normal HTTP AUTH (for example
    through Apache's htpasswd). No one is stopping you from using one single
    htpasswd file for all your instances. This will at least solve the login
    problem, but not the existence of one session_attribute table per
    instance.
2.  To get around this first I thought thanks to some co-workers about the [ATTACH DATABASE](http://sqlite.org/lang_attach.html)
    statement in SQLite. But the problem here is, that ATTACH DATABASE
    is just a run-time command, meaning that you have to execute it manually
    after opening the connection to the database. 
    
So I asked in #postgresql and AlexB said, that what I wanted could be achieved
using a feature of PostgreSQL named "Schemas". First I didn't know what this
was all about, so I took a look at the [documentation](http://www.postgresql.org/docs/8.2/interactive/ddl-schemas.html): Schemas allow you to
have some kind of layering within a single database. For example you have
database DB in which should be used by multiple users. Now you want some 
tables to look different depending on which user accesses them while some
tables should look exactly the same for everyone. 

And this is exactly what I wanted for my multi-Trac installation. So I did
the following:

1.  Create a database called "trac" which should be used by all those Trac
    installations. By default every database uses the "public" schema.
2.  Then I created a role for every Trac installation and also a group role
    called "group_trac" to which I added all those roles.
3.  Now comes the interesting part: I had to create a schema for every
    trac-role in the "trac" database that belongs to the respective user. 
    This is done with following statement:
        
        CREATE SCHEMA trac_installation1 AUTHORIZATION trac_installation1;
        
    
    This means that the schema "trac_installation1" really belongs to the user
    with the same name.
4.  Now comes the boring part, where I had to create all those Trac 
    instances with the usual 'trac-admin /path/to/it/ initenv' command.
5.  The last step is the one, where the actual account-data-linking happens:
    Trac stores all the account data like email and realname in the 
    session_attribute table, which we now want to have shared over all
    the instances so that the users don't have to type their e-mail address
    again and again and again into the same dull form. 
    
    So I simply took one of the session\_attribute tables and moved it
    to the public schema. Since there seems to be no simple way to do that
    directly within PostgreSQL, I just used pg\_dump to dump only the one
    table (using the -t TABLENAME parameter), edited the statements to 
    use the `schemaname.tablename` format and imported it again. After importing
    that table into the public schema, I also granted the group_trac
    all the necessary permissions on this table (SELECT, UPDATE, DELETE).
    
    This works, 
    because PostgreSQL will first look for tables within a schema with the
    same name as the currently connected user (if it exists) and then fall back
    to the "public" schema. Again for details check out the great documentation
    available on [postgresql.org](http://www.postgresql.org/docs/8.2/interactive/).

And that's it :-)
    