---
date: '2009-11-11T12:00:00-00:00'
language: en
tags:
- applicationcontext
- spring
title: Root contexts prefer not to share
---


Tonight I learned something funny with <a href="http://springframework.org/">Spring</a>: Don&#39;t be lazy when defining your root application context. I wanted to play around with Jersey and its integration with Spring, but was kind of lazy:

<pre class="code">    &lt;context-param&gt;
        &lt;param-name&gt;contextConfigLocation&lt;/param-name&gt;
        &lt;param-value&gt;WEB-INF/SpringMVC-servlet.xml&lt;/param-value&gt;
    &lt;/context-param&gt;
    &lt;servlet&gt;
        &lt;servlet-name&gt;SpringMVC&lt;/servlet-name&gt;
        &lt;servlet-class&gt;org.springframework.web.servlet.DispatcherServlet
        &lt;/servlet-class&gt;
    &lt;/servlet&gt;
    &lt;servlet&gt;
        &lt;servlet-name&gt;spring-jersey&lt;/servlet-name&gt;
        &lt;servlet-class&gt;com.sun.jersey.spi.spring.container.servlet.SpringServlet
        &lt;/servlet-class&gt;
    &lt;/servlet&gt;
</pre>

This is a really bad idea because you will end up with both the REST Resources as well as everything in Spring being able to access the same beans (those defined in the SpringMVC-servlet.xml) but with actually 2 contexts. If you, for instance, try to access a singleton bean called &quot;dataService&quot; from within a Resource and then also from within a SpringMVC controller, the hashCode of that so-called &quot;singleton&quot; will be different ...

So don&#39;t be lazy and give your root application context its own bean configuration if you intend to put something like a connection pool or something that really <strong>should</strong> be singleton in there ;-)
