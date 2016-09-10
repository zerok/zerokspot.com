---
date: '2005-07-18T12:00:00-00:00'
language: en
tags:
- development
title: Indexing script for docbook articles ;)
---


Today I wanted to have a index.html for a directory full of DocBook articles. Being quite lazy I simply wrote a small Ruby script for doing this since I couldn't get sed to do what I want:

-------------------------------



<pre class="code">#!/usr/bin/ruby -w

require 'rexml/document'

def getTitle(file)

        doc = REXML::Document.new(File.new(file))

        title = doc.elements.each("/article/title"){|t| return t.text}

        return File.filename(file)

end

puts &lt;&lt;EOS

&lt;!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd"&gt;

&lt;html xmlns="http://www.w3.org/1999/xhtml"&gt;

&lt;head&gt;

        &lt;title&gt;Guides&lt;/title&gt;

        &lt;meta http-equiv="Content-type" content="text/html;charset=UTF-8"/&gt;

        &lt;link rel="stylesheet" type="text/css" href="../_s/style.css"/&gt;

&lt;/head&gt;

&lt;body&gt;

&lt;h1&gt;Guides&lt;/h1&gt;

&lt;ul&gt;

EOS

dh = Dir.new(".")

dh.each do |f|

        next if ['.','..'].include?(f)

        path = dh.path+File::SEPARATOR+f

        indexxml = path+File::SEPARATOR+"index.xml"

        indexhtml=path+File::SEPARATOR+"index.html"

        if File.directory?(path) and File.exists?(indexxml)

                puts "&lt;li&gt;&lt;a href="#{indexhtml}"&gt;#{getTitle(indexxml)}&lt;/a&gt;&lt;/li&gt;"

        end

end

puts &lt;&lt;EOS

&lt;/ul&gt;

&lt;/body&gt;

&lt;/html&gt;

EOS

</pre>



Since I'm currently learning Ruby - or let's call it refreshing my Ruby knownledge - I'm trying to write most of the scripts I need everyday in Ruby to get some practice. Starting with today I will also post these small scripts here :)