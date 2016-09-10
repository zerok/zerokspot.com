---
date: '2005-05-01T12:00:00-00:00'
language: en
tags:
- development
title: emerge -fp downloader
---


Before replacing Gentoo with Ubuntu on my server I want to give it another chance. The problem is, there seems to be a problem with at least one of the system packages (like autoconf or automake). I tried to compile and install ApacheHTTPD and PHP like I've done it on some other servers, but the Makefile doesn't seem to produce a working install label. For example the Apache Makefile never installs the modules and PHP doesn't install the module before chmodding it.

-------------------------------



So I've decided to first run `emerge -De system && emerge -De world` which should reinstall the whole system. Since my server is not online I have to get the sources from somewhere. So I generated with `emerge -pfDe system 2> list` a list of all the source packages that have to be downloaded an following small script that should fetch them all on a different machine :-)

<pre class="code">
#!/usr/bin/perl
if (@ARGV == 0){
	print "downloadlist.pl filelist\n";
}
my $file = $ARGV[0];
open(FILE,"< $file") or die("Couldn't open input file");
if (not -f "distfiles"){
	mkdir "distfiles";
}
download: while(&lt;FILE&gt;){
	if (/^$/){
		next;
	}
	else {
		my @candidates = split / /, $_;
		my @filename = split(/\//,@candidates[0]);
		my $filename = $filename[@filename-1];
		foreach $candidate (@candidates){
			`cd distfiles ; wget -c $candidate`;
			if ($? == 0){
				# Move the file to the storage 
				# directory
				#`mv $filename distfiles`;
				next download;
			}
		}
		print "Failed to download $filename\n";
		
	}
}
close(FILE);
</pre>

This will download all the files listed in the $list file (and should also try the mirrors if the download fails) and puts them into the $PWD/distfiles folder.