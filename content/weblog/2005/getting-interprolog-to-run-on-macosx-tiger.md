---
date: '2005-12-20T12:00:00-00:00'
language: en
tags:
- development
title: Getting InterProlog to run on MacOSX Tiger
---


Take this as some kind of part one of probably a small series of posts about this whole topic since I just started to master the first few obstacles involved here. 

* SWI Prolog 5.4.7 for MacOSX
* Interprolog 2.21



-------------------------------



The installation manual is already quite detailed but there are still some things that are different with the current versions of MacOSX and Prolog. First of all you probably want to get at least one of the example scripts to work. So let's start with taking a look at the unixScripts/MacOSXVariables.sh file:

<pre class="code">
# Mac OS locations are typically different from Linux
# You might want to edit unixVariable.sh variables starting from these:
JAVA_BIN=/System/Library/Frameworks/JavaVM.framework/Versions/1.4.2/Home/bin
XSB_BIN_DIRECTORY=/MY_XSB_DIR/config/powerpc-apple-darwin7.9.0/bin
SWI_BIN_DIRECTORY=/usr/local/lib/swipl-5.4.7/bin/powerpc-darwin6.6
YAP_BIN_DIRECTORY=/usr/local/bin
</pre>

Since this tutorial is only about getting SWI Prolog to work with Interprolog we can ignore the XSB and YAP settings. The SWI_BIN_DIRECTORY is no longer correct on Tiger though, so let's correct it:

<pre class="code">
SWI_BIN_DIRECTORY=/usr/local/lib/swipl-5.4.7/bin/powerpc-darwin7.8.0
</pre>

Well, now let's take a look at the runSWIPLListener.sh script:
<pre class="code">
. unixVariables.sh
${JAVA_BIN}/java -classpath ${CLASSPATH}:../interprolog.jar com.declarativa.interprolog.gui.SWISubprocessEngineWindow $1 ${SWI_BIN_DIRECTORY}/pl
#For Mac version use /swipl rather than /pl
</pre>

Here we will have to include the MacOSXVariables.sh instead of the unixVariables.sh script. As the last comment suggests, "pl" also has to replaced with "swipl" to make this work under MacOSX.

Well, that's all that is needed to get at least this script working under MacOSX Tiger, but you should have gotten an idea, what is needed in general. I just think it's bad, that the demo scripts weren't written in a more generic way using for example `uname` and some bash magic to make the correct assumptions about the location of `swipl` and so on.

<div class="figure">
<img src="http://www.zerokspot.com/uploads/runswipllistener.png" alt="Listener Screenshot"/>
</div>