SSH Commands on Ubuntu Servers
##############################

Fighting with environment variables via SSH on an Ubuntu server? You're 
definitely not alone there. Normally when
running ``ssh <server> <command>`` the user's ~/.bashrc should be 
sourced (if the user has Bash as shell) but for some reason, none of my
variables got loaded.

After quite some googling, I found `this post <http://sayspy.blogspot.com/2006/10/anyone-know-how-to-get-sshbash-to-use.html>`_
that described the same problem I had but also contained a simple solution
in the comments section: Actually reading the .bashrc ;-) Ubuntu's default
.bashrc starts with following line::
    
    [ -z "$PS1" ] && return
    
The problem is now, that this condition is actually true via ``ssh <command>``.
To solve this you have more or less 2 options: Add you stuff before this line
or remove this line. 

Just stupid that it took me about an hour to find this problem ;)