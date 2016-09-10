---
date: '2007-06-11T12:00:00-00:00'
language: en
tags:
- latex
title: 'LaTeX: Section label vs. Figure label'
---


Just found out something about LaTeX that I ... actually hadn't anticipated: The whole cross-referencing seems to be quite fragile. 



-------------------------------



For example: I just got reviews to a seminar work I had to write. One reviewer noted, that all references to figures actually point to the section the figures appear instead. I looked at the code and so something like this:

<pre class="code tex">
\begin{figure}[h!t]\label{fig:asnstack}
	\begin{center}
		\includegraphics[width=5cm]{i/asnstack.png}
		\caption{ASN.1 und XML}
	\end{center}
\end{figure}
</pre>

Looks good to me.

Well, but in the output it clearly doesn't as has been described by my dear reviewer. So what's the problem? It may sound funny, but it seems like placing the \label before the \caption is a bad idea. Thanks to [Peter Newbury](http://www.iam.ubc.ca/~newbury/tex/figures.html#label) for describing this.

So basically if you have the same problem, just write it this way and make your peers happy:

<pre class="code tex">
\begin{figure}[h!t]
	\begin{center}
		\includegraphics[width=5cm]{i/asnstack.png}
		\caption{ASN.1 und XML}
		\label{fig:asnstack}
	\end{center}
\end{figure}
</pre>

Let's hope this is the only problem they noticed with your work ;-)