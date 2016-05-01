# Quick switching to a project's directory

During her [talk about containers and sandboxes][talk] at [CraftConf][] last
Friday, Jess Frazelle used a little helper function in her shell to quickly
switch into a project's directory in her shell. This little gem can be found in
her [dotfiles repository on github][gogo]. This implementation focuses on being
able to easily get into a directory within her GOPATH.

-----

Right now I'm still switching from environment to environment quite frequently
so I need something a bit different, though. Not being all that familar with
scripting ZSH I need a ton of skimming through the docs but this is what I
eventually came up with: [cdp][]. It is more or less a simple function that
takes a parameter, scans a handful of directories for folders with the same name
as the parameter and cds into the first match.

Since I've always wanted to mess around with some autocompletion, that's also
support here. Hacky, but it works for now 😊

Being not really used to writing any Bash or ZSH code my code will probably be
offensive to quite a few folks, so I apologize 😊 I might improve this in the
future as I learn more about ZSH but for now this does the job for me 😉

A big THANK YOU to Jess Frazelle for the inspiration!

[cdp]: https://github.com/zerok/cdp
[talk]: http://beta.craft-conf.com/speaker/JessieFrazelle
[craftconf]: http://craft-conf.com
[jess frazelle]: https://blog.jessfraz.com/
[gogo]: https://github.com/jfrazelle/dotfiles/blob/3ef727195c9af666a309a27e9c10a8b5ac4b5b33/.functions#L335-L360
