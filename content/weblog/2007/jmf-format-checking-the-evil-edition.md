---
date: '2007-02-05T12:00:00-00:00'
language: en
tags:
- java
- jmf
- reflection
title: 'JMF Format checking: The Evil Edition'
---


Desperate times sometimes require desperate solutions for stupid problems. For example today I have a deadline for a small JMF project. A part of the project is, that an RTSP client should be able to request only for example the audio streams of a certain object on the server so the client should have some kind of flag that specifies what streams should now really be used.


-------------------------------


But what happens, when you have an MP3 file and the user requests "video-only"-playback? Should really the player try to SETUP the whole thing or should it just be able to determine based on the server's response to a RTSP call wether there even _is_ a video stream? I would say so, yes, but there comes the next big problem right around the corner: From what I've seen of JMF so far, you can just simply pass something like "MPEG/RTP" into an javax.media.Format class and determine from this wether or not this is an audio or a video stream.

So now back to the desperate solutions thing: Format has two subclasses for handling Audio and Video formats ... AudioFormat and VideoFormat. Both of them have a bunch of constants that store strings like "MPEG/RTP" ;-) And Java has some nice reflection-abilities. I guess that code was inevitable:

<pre class="code java"><span class="k">private</span> <span class="k">static</span> <span class="kt">boolean</span> <span class="n">checkFormat</span><span class="o">(</span><span class="n">String</span> <span class="n">input</span><span class="o">,</span><span class="n">Class</span> <span class="n">formatClass</span><span class="o">){</span>
	<span class="k">for</span><span class="o">(</span><span class="n">Field</span> <span class="n">f</span> <span class="o">:</span> <span class="n">formatClass</span><span class="o">.</span><span class="n">getDeclaredFields</span><span class="o">()){</span>
		<span class="kt">int</span> <span class="n">m</span> <span class="o">=</span> <span class="n">f</span><span class="o">.</span><span class="n">getModifiers</span><span class="o">();</span>
		<span class="n">Object</span> <span class="n">v</span><span class="o">;</span>
		<span class="k">if</span> <span class="o">(</span><span class="n">Modifier</span><span class="o">.</span><span class="n">isFinal</span><span class="o">(</span><span class="n">m</span><span class="o">)</span> <span class="o">&amp;&amp;</span> <span class="n">Modifier</span><span class="o">.</span><span class="n">isStatic</span><span class="o">(</span><span class="n">m</span><span class="o">)){</span>
			<span class="k">try</span> <span class="o">{</span>
				<span class="n">v</span> <span class="o">=</span> <span class="n">f</span><span class="o">.</span><span class="n">get</span><span class="o">(</span>
					<span class="n">formatClass</span><span class="o">.</span><span class="n">getConstructor</span><span class="o">(</span><span class="n">String</span><span class="o">.</span><span class="n">class</span><span class="o">).</span><span class="n">newInstance</span><span class="o">(</span><span class="s">&quot;&quot;</span><span class="o">)</span>
					<span class="o">);</span>
				<span class="k">if</span> <span class="o">(</span><span class="n">v</span> <span class="k">instanceof</span> <span class="n">String</span><span class="o">){</span>
					<span class="k">if</span><span class="o">(((</span><span class="n">String</span><span class="o">)</span><span class="n">v</span><span class="o">).</span><span class="n">toLowerCase</span><span class="o">().</span><span class="n">equals</span><span class="o">(</span><span class="n">input</span><span class="o">.</span><span class="n">toLowerCase</span><span class="o">())){</span>
						<span class="k">return</span> <span class="kc">true</span><span class="o">;</span>
					<span class="o">}</span>

				<span class="o">}</span>
			<span class="o">}</span>
			<span class="k">catch</span><span class="o">(</span><span class="n">Exception</span> <span class="n">e</span><span class="o">){};</span>
		<span class="o">}</span>
	<span class="o">}</span>
	<span class="k">return</span> <span class="kc">false</span><span class="o">;</span>
<span class="o">}</span></pre>


This little method does one thing: It searches all the String constants of the given class for one with the value that is the same as the other input parameter.

So you can now do something like this:

<pre class="code java">
<span class="n">checkFormat</span><span class="o">(</span><span class="s">&quot;MPEG/RTP&quot;</span><span class="o">,</span><span class="n">javax</span><span class="o">.</span><span class="n">media</span><span class="o">.</span><span class="n">format</span><span class="o">.</span><span class="n">VideoFormat</span><span class="o">.</span><span class="n">class</span><span class="o">));</span>
</pre>

... and get true. Probably a stupid solution, but at least it works ;-)