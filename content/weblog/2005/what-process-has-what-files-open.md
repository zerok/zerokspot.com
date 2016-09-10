---
date: '2005-03-09T12:00:00-00:00'
language: en
tags:
- development
title: What process has what files open?
---


I think everyone knows the problem: Sometimes some process changes files while you want to play with them. Be it unmounting a device, or simply some concurrent writing. Under Linux it's fairly easy to find out what process has what file descriptors open. That's all simply stored in /proc. Now I wanted to find | grep it and ... well, I simply don't really like Bash so I wrote a small script in Perl that will print all open files with their name and the process id that has the file open:

-------------------------------



<pre class="code">
#!/usr/bin/perl
my $procdir = '/proc';
opendir(PROCDIR,"/proc") or die("Can't open proc...");
my @procs;
while(defined($file = readdir(PROCDIR))){
	if ($file =~ /^[0-9]/) {
		if (-e '/proc/'.$file.'/fd/'){
			push @procs,$file;
		}
	}
}
closedir(PROCDIR);
foreach(@procs){
	my $pid = $_;
	my $fdpath = '/proc/'.$_."/fd/";
	opendir(DIR,$fdpath);
	while(defined($file = readdir(DIR))){
		if (-l $fdpath.$file){
			if(defined($act = readlink($fdpath.$file))){
				if ($act =~ /^\//){
					print "[PID:$pid]",$act,"\n";
				}
			}
		}
	}
	closedir(DIR);
}
</pre>

Perhaps someone else will find this useful :-) It's not fast, it's not nice, but it does the job :-)