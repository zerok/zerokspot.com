# fish shell functions are magical!

For some of our projects at work we have decided that we want to stick with
Node’s LTS releases for the next couple of months. In order to not prevent
experimentation with newer versions must of us have [NVM][] installed.

In order to make switching between versions easier depending on what project I’m
currently working on, I’m playing a little bit around with a little shell helper
that should load the correct version for a project automatically from a `.nvmrc`
file.

As this is also the first time I'm trying to extend my shell since switching to
[fish][] I’ve learnt quite a bit. One thing in particular I’ve found quite
interesting: Functions support acting as “event handlers” out of the box. I use
the term “event” loosely here. To be more concrete, I’m talking about things
like signals, actual events, and changes to the value of a variable.

------------------------------------------------------------------------

The latter is what I was looking for for my little task at hand. If you declare
a function with the `--on-variable NAME` parameter, any change to the variable
will trigger that function.

``` fish
function hello_dir --on-variable PWD
    echo "Welcome to $PWD"
end
```

The function above (if you put it for instance in your
`~/.config/fish/fish/config.fish` file) will echo a nice welcome message
whenever you enter a new directory. That’s because, the global `$PWD`
environment variable has been set to a new value.

In other shells like Bash or ZSH you’d override the `cd` function for something
like that. Being able to simply add additional handlers on top of each other is
so much nicer!

But, as mentioned above, fish’s [function command][] can do even more. Want to do
something when a process signals for instance HUP? `--on-signal SIGHUP` has you
covered. Do you want to get notified when a certain child process exists? Use
`--on-process-exit 123` for the process with the PID 123.

As for what fish considers events:

``` fish
$ function event_handler --on-event CUSTOM_EVENT
$    echo "Custom event was triggered with these args: $argv"
$ end

$ emit CUSTOM_EVENT something
Custom event was triggered with these args: something
```

The shell already emits a couple of internal events, like `fish_preexec` and
`fish_postexec`. These sound like nice entry points for further exploration 😉

To get back to my original use-case, I'm not completely sure yet, how to
implement it. Obviously, I won't execute nvm with the content of some random
file by default but probably implement a little white-listing-logic here. We'll
see 😊

[NVM]: https://github.com/creationix/nvm
[function command]: http://fishshell.com/docs/current/commands.html#function
[fish]: http://fishshell.com/
