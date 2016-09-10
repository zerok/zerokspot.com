---
date: '2006-11-11T12:00:00-00:00'
language: en
tags:
- python
- textmate
title: Explicitly splitting long lines in TextMate
---


Sometimes, when you're just in the flow of typing text (like for a paper) you might forget about some of the personal rules you set up on how you want your textfiles formated. I, for example, prefer it, when the length of each line isn't above 78 characters. I don't really want this an automatic method for formating source code files, but at least for some LaTeX/reST/Markdown texts this would be nice. But while TextMate offers a small command in the "Text" menu to unwrap paragraphs, I didn't find anything for actually splitting (and not just softwrapping) long lines.



-------------------------------



Not really a problem so let's write a simple command for this. And here is the result. I don't know if it works in all cases, but I tested it for situations where nothing is selected or whole lines are selected and it seems to work there. 

**Warning:** Don't try this command when you don't have full lines selected. This command wasn't built for this and I haven't yet found a way to check what kind of stuff got selected :) And don't try it on program source code ;-)

<pre class="code python">
#!/usr/bin/env python
import os,sys
from sys import stdout,stdin
              
def split_line(line,max_len,tab_size):
	if len(line) &lt;= max_len:
		return [line,]
              
	prefix = &apos;&apos;
	prefix_length = 0
	if line[0] in (&apos; &apos;,&quot;\t&quot;):
		for c in line:
			if c not in (&apos; &apos;,&quot;\t&quot;):
				break
			else:
				prefix += c
				if c == &apos;\t&apos;:
					prefix_length+=tab_size
	
	lines = []
	current_line = []
	line_no = 0
	words = line.split(&quot; &quot;)
	chars = prefix_length
	if len(words)==1:
		return [line,]
	while len(words) &gt; 0:
		if (chars+len(words[0])) &lt;= max_len:
			word = words.pop(0)
			current_line.append(word)
			chars += len(word)+1
		else:
			if line_no &gt; 0:
				lines.append(&quot;%s%s&quot;%(prefix,&quot; &quot;.join(current_line)))
			else:
				lines.append(&quot; &quot;.join(current_line))
			current_line = []
			line_no+=1
			chars = prefix_length
	if len(current_line) &gt; 0:
		if line_no &gt; 0:
			lines.append(&quot;%s%s&quot;%(prefix,&quot; &quot;.join(current_line)))
		else:
			lines.append(&quot; &quot;.join(current_line))
	return lines

if __name__ == &apos;__main__&apos;:
	max_len = int(os.getenv(&apos;TM_WRAP_LENGTH&apos;,78))
	tab_size = int(os.getenv(&apos;TM_TAB_SIZE&apos;,4))
	for line in stdin:
		stdout.write(&quot;\n&quot;.join(split_line(line,max_len,tab_size)))

</pre>

And here's the configuration for this command. I simply put it into a dummy bundle handling all my private extensions to the "Text" bundle ... well, actually this is the first extensions ;-)

<a class="thickbox figure" title="Make a command for replacing text coming from 'selected text' or 'current line'" href="http://zerokspot.com/uploads/tm_splitlines_config.png"><img src="http://zerokspot.com/uploads/tm_splitlines_config-mini.jpg" alt="Make a command for replacing text coming from 'selected text' or 'current line'"/></a>

**Warning:** The usual stuff. Use this on your own risk etc.

I hope this might be useful for some of you :-)