---
date: '2012-05-20T12:00:00-00:00'
language: en
tags:
- tmux
- split-window
title: tmux, split-window and pwd
---


For the last couple of weeks I've been playing around with [tmux](http://tmux.sf.net) as kind of an alternative/addon/... to iTerm2 and my
terminal-focused work-style in general. One things that drove me nuts again and again is that -- for some reason -- I
could never get the split-window command to respect my current working directory. I tried the first two solutions
mentioned on the [ArchLinux-wiki](https://wiki.archlinux.org/index.php/Tmux#Split_window_and_retain_current_directory)
but the second simply did clash too much with how I'm using split-window and the first solution didn't work at all (for
some reason PWD didn't propagate correctly and so default-path always defaulted back to the cwd of tmux' parent shell).

With all other options gone I started to look into the [third solution](https://wiki.archlinux.org/index.php/Tmux#.2Fproc_method).
This one involved extracting the PWD from the pane's shell from the outside via /proc. But since I'm working on OSX, 
this doesn't work quite like this.

-----------------------------

The main differences here are that I can't use /proc to look for the PWD but have to rely on `lsof` and that `ps`
doesn't have a forest-rendering. So the resulting script looks like this:

<pre><code>
#!/usr/bin/env ruby
=begin
Helper script for tmux' split-window command that retains the original PWD
This script is based on https://wiki.archlinux.org/index.php/Tmux#.2Fproc_method
=end
SHELL = 'reattach-to-user-namespace -l zsh'


session_id, window_id, pane_id = `tmux display-message -p "#S:#I:#P"`.chomp.split(/:/)
tty = nil
in_sessions = false
in_right_session = false
in_right_window = false
in_right_pane = false

# First step is to find our tty
`tmux server-info`.each_line do |line|
    break if line.start_with? 'Terminals'
    if line.start_with? 'Sessions: '
        in_sessions = true
        next
    end
    if in_sessions and line =~ /^\s\d+: #{session_id}: \d+ window..*[flags=.*]/
        in_right_session = true
        next
    end   
    if in_right_session and line =~ /^\s{3}#{window_id}:/
        in_right_window = true
        next
    end
    if in_right_window and line =~ /^\s*#{pane_id}: \/dev\/(\S+) /
        tty = $1
        break
    end
end

raise "No matching tty found" if tty.nil?

# Now let's find the process with this tty
pid = `ps -t #{tty}`.split($/)[1].split(' ')[0]

# ... and now its PWD
pwd = `lsof -a -d cwd -p #{pid} -F n`.split($/)[1][1..-1]
args = ARGV * " "

# Let's put it all together by calling split-window with the shell and the path
`tmux split-window #{args} 'cd #{pwd}; #{SHELL}'`
</code></pre><a href="https://gist.github.com/2758863">Gist</a>

Thankfully, getting access to a process' PWD is pretty straight forward with `lsof -d cwd` and `ps` has a nice shortcut
for filtering for a specific tty. And yes, I know that the tty-scanning is really rocky up there and this is probably 
the one part of this script that will evolve over time. The solution for the missing "--forest" support is also not 
perfect yet, but so far it works just fine for me :-)

So, once you've put this script somewhere in your path (for me that ~/.local/bin/split-in-cwd.rb) and made it 
executable, all that is left to do, is to add it to your split-window bindings:

<pre>
unbind s
bind s run '~/.local/bin/split-in-cwd.rb -v'
unbind v
bind v run '~/.local/bin/split-in-cwd.rb -h'
</pre>