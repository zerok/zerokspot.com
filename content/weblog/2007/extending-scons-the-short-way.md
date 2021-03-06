---
date: '2007-11-29T12:00:00-00:00'
language: en
tags:
- python
- rst
- scons
title: 'Extending SCons: The Short Way'
---


[SCons](http://www.scons.org/) is for me one of the most convenient build systems out there but at
least if you take a look at the manual and esp. in the [section about writing
your own builders](http://scons.org/doc/0.97.0d20070918/HTML/scons-user/x2369.html), it looks like a little bit of a pain to extend. Or actually
there doesn't seem to be a short way to extend build environments at all.

-------------------------------

Today, for example, I was looking for a way to convert RST documents into
HTML using SCons and therefor wanted to write and integrate a simple builder
that does this job. According to the [manual](http://scons.org/doc/0.97.0d20070918/HTML/scons-user/x2446.html) this would work somehow like
this (the *rstsuite* package being one of my toolkit packages):
    
    import os
    from rstsuite.scons import build_document
    
    env = Environment(ENV=os.environ)
    env['BUILDERS']['RSTDocument'] = Builder(action=build_document, 
        suffix='.html', src_suffix='.txt')
    env.RSTDocument(source='index.txt', target='index.html')
    

For this task, the *build_document* function would be quite short:
    
    def build_document(target, source, env):
        from docutils.core import publish_file
        assert(len(source)==1)
        assert(len(target)==1)
        publish_file(source_path=str(source[0]),
            destination_path=str(target[0]), 
            writer_name='html')
        return None
    
But I'm quite sure, that line 5 of the SConstruct file would eventually get on my nerves
if I had to write it ever single time I had a SConstruct file that should 
build some RST document. So I started looking through SCons' source files
in the hope to find some kind of dict or list that would let me set that
builder once and for all. That *Defaults.py* file really looked suspicious
and I really seem to have been lucky once again: It holds a dict for the 
default environment called "ConstructionEnvironment" that appears to be
what I was looking for. 

Now I could compress my SConstruct file to this one:
    
    import rstsuite.scons
    env=Environment()
    env.RSTDocument(source="index.txt", target="index.html")
    

with the *rstsuite.scons* module looking like this:
    
    from SCons.Defaults import ConstructionEnvironment
    from SCons.Builder import Builder

    def build_document(target, source, env):
        from docutils.core import publish_file
        assert(len(source)==1 and len(target)==1)
        publish_file(source_path=str(source[0]), 
            destination_path=str(target[0]), 
            writer_name='html')
        return None

    ConstructionEnvironment['BUILDERS']['RSTDocument'] = Builder(
        action=build_document,
        suffix='.html', 
        src_suffix='.txt')
    

If anyone knows of a better solution for extending the default environment, 
please let me know :-) For now at least, this solution seems to work, though.