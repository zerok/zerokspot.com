---
date: '2008-05-23T12:00:00-00:00'
language: en
tags:
- java
title: Small problems with Mevenide 3.1
---


A couple of days ago the [Mevenide](http://mevenide.codehaus.org/m2-site/) team released version 3.1 of their Maven plugin for the Java IDE Netbeans. Naturally an auto-update process kicked in here and it replaced the previously installed version 3.0.12 on my Netbeans installation.

When I restarted the IDE afterwards and wanted to run a project of mine, I was greeted with a nice little error message.

-------------------------------

<blockquote>
[ERROR]
<br/>[ERROR]
<br/>[ERROR]Maven cannot calculate your build plan, given the following information:
<br/>[ERROR]
<br/>[ERROR]Tasks:
<br/>[ERROR]- package
<br/>[ERROR]- org.codehaus.mevenide:netbeans-run-plugin:RELEASE:run-jar
<br/>[ERROR]
<br/>[ERROR]Current project:
<br/>[ERROR]Group-Id: MyProject
<br/>[ERROR]Artifact-Id: MyProject
<br/>[ERROR]Version: 0.0.1-SNAPSHOT
<br/>[ERROR]From file: /Users/zerok/Documents/workspace/MyProject/pom.xml
<br/>[ERROR]
<br/>[ERROR]
<br/>[ERROR]Error message: Failed to resolve plugin for mojo binding: org.apache.maven.plugins:maven-compiler-plugin:RELEASE:compile
<br/>[ERROR]Root error message: The PluginDescriptor for the plugin org.apache.maven.plugins:maven-compiler-plugin was not found. Should have been in realm: ClassRealm[/plugins/org.apache.maven.plugins:maven-compiler-plugin:RELEASE@48/thread:exec_Run MyProject_5, parent: null]
</blockquote>

For the last two weeks now I've ignored this error and simply did all my program-executing in the shell using a small bash script to copy over all the dependencies and setting the respective classpath.

But today I stopped being the only one affected in the office since my boss installed Netbeans again and naturally also received this message of love from the Maven plugin.

After some length googling around I eventually found a solution for this problem provided by [imyousuf](http://www.dzone.com/links/users/profile/166473.html) on [dzone](http://www.dzone.com/links/netbeans_maven_intragation_31_released_.html). The problem is, that for some reason the plugin can no longer resolve the "RELEASE" keyword correctly when loading the maven-compiler-plugin; and the solution is pretty simple:

Search your pom.xml for the configuration for this plugin:

<pre><code>
&lt;plugin&gt;
    &lt;artifactId&gt;maven-compiler-plugin&lt;/artifactId&gt;
    &lt;version&gt;RELEASE&lt;/version&gt;
    &lt;configuration&gt;
        &lt;source&gt;1.5&lt;/source&gt;
        &lt;target&gt;1.5&lt;/target&gt;
        &lt;encoding&gt;UTF-8&lt;/encoding&gt;
    &lt;/configuration&gt;
&lt;/plugin&gt;
</code></pre>

... and replace "RELEASE" in there with a concret version like for instance "2.0.2" -- which is the latest release of this plugin.

Another problem I ran into was that when you've created your project using an older version of the Maven plugin, it created a configuration for the maven-assembly-plugin that contains some now deprecated parts. If you see something like this in your pom.xml

<pre><code>
&lt;plugin&gt;
    &lt;artifactId&gt;maven-assembly-plugin&lt;/artifactId&gt;
    &lt;executions&gt;
        &lt;execution&gt;
            &lt;id&gt;nb&lt;/id&gt;
            &lt;phase&gt;package&lt;/phase&gt;
            &lt;goals&gt;
                &lt;goal&gt;directory&lt;/goal&gt;
            &lt;/goals&gt;
            &lt;configuration&gt;
                &lt;descriptor&gt;
                    ${basedir}/src/main/assemblies/netbeans-run.xml
                &lt;/descriptor&gt;
                &lt;finalName&gt;executable&lt;/finalName&gt;
            &lt;/configuration&gt;
        &lt;/execution&gt;
    &lt;/executions&gt;
&lt;/plugin&gt;
</code></pre>

There the whole &lt;goal&gt;-block seems now to be deprecated. After removing (or actually commenting it out) the project finally ran again from within Netbeans :-)
