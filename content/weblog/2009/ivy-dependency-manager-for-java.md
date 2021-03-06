---
date: '2009-05-31T12:00:00-00:00'
language: en
link: http://ant.apache.org/ivy/
tags:
- dependencies
- development
- ivy
- java
- tools
title: 'Ivy: A dependency manager for Java'
url_title: Apache Ivy, the agile dependency manager
---


In whatever language you are programming you eventually end up in a situation
where you are working with so many libraries that you need to manage them
somehow. Up until a couple of days ago, I used [Maven][] for this when dealing
with Java. But Maven might be a bit more than what you normally want. It not
only manages dependencies but also comes with its own project life-cycle
management system.  You can use Maven for compiling, testing, deploying,
generating the documentation, publishing it on a website, and so on. But at
least for me, all I normally want is a simple build system. If it has some
nice extras, fine, but I don't really like to be forced in some kind of
life-cycle management. So basically give me something like Make or Ant and
then give me something that manages dependencies. This is more or less how I
discovered [Ivy][]. 

-------------------------------

[Ivy][] is a dependency manager that integrates with [Ant][], and that's all:
nothing more and certainly nothing less.  In order to use it you have to first
of all have an ivy.xml file that holds the dependencies of your project as
well as some meta-information about it (for publishing). This is what such a
file could look like for a project that depends solely on the [Rome][] library:
    
    <?xml version="1.0" encoding="ISO-8859-1"?>
    <ivy-module version="1.0">
        <info 
            organisation="com.zerokspot.playground"
            module="rss"
            status="release">
        </info>
        <dependencies>
           <dependency org="rome" name="rome" rev="1.0" />
        </dependencies>
    </ivy-module>
        
Next, you will need to tell ant, that it should download those dependencies:
    
    <?xml version="1.0" encoding="UTF-8"?>
    <project name="rss-playground" default="build"
        xmlns:ivy="antlib:org.apache.ivy.ant">
        
        <property name="src.dir" value="src" />
        <property name="build.dir" value="bin" />
        <path id="classpath">
            <fileset dir="lib" includes="*.jar" />
        </path>
            
        <target name="build" depends="resolve">
            <javac srcdir="${src.dir}" destdir="${build.dir}" 
                  classpathref="classpath" />
        </target>
        
        <target name="resolve">
            <ivy:retrieve />
        </target>
    </project>    

The important part is the ``<ivy:retrieve />`` task in the target "resolve".
And that's it. But what if you want to use a different repository (for instance
an internal one)? That's also possible and rather simple: Ivy lets you define
a chain of [dependency resolvers][], which allows you to use libraries
accessible from SSH accounts, Maven2 compatible repositories, URLs in general
and much much more. 

Another really nice feature of Ivy is that you can also define certain
requirements for certain situations -- or [configurations][] as they are called
here. For instance: You want your package to depend on hsqldb while testing
but in production you also want to include postgresql. And since certain
configurations of your dependency's ivy.xml might not be available in your own
project, you can also define a [mapping][] to resolve those differences. Like a
library you're depending on has the configurations "hsqldb" and "postgresql",
while you only have "runtime" and "test". Then you can say, that "test" maps
to "hsqldb" and "runtime" to "postgresql".

That all is still completely focused on doing one thing: managing
dependencies. Just what I want :-) I can't believe how I could miss that
alternative to Maven for so long. Note that I just stumbled upon Ivy, so there
are definitely also some quirks here and there, but so far I really like what
I see.

[maven]: http://maven.apache.org/
[dependency resolvers]: http://ant.apache.org/ivy/history/2.1.0-rc1/settings/resolvers.html
[ivy]: http://ant.apache.org/ivy/
[rome]: https://rome.dev.java.net/
[ant]: http://ant.apache.org/
[mapping]: http://ant.apache.org/ivy/history/2.1.0-rc1/ivyfile/dependency.html#defaultconfmapping
[configurations]: http://ant.apache.org/ivy/history/2.1.0-rc1/ivyfile/configurations.html
