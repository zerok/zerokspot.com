Git branches and PS1
####################

With the release of Mercurial 1.1 yesterday and me more and more getting into it again I complained on IRC that as much as I love git's named branches, they make it kind of easy to just do something to the *wrong branch*. If you have a branch per folder it's kind of more obvious in what branch you're operating all the time, in my opinion. 

But there is a simple solution for this: Martin today posted a `quick guide <http://www.mahner.org/weblog/git-branch-im-bash-prompt-anzeigen/>`_ on how to get the current branch name into your shell's $PS1.

-------------------------------

For this to work you need to source git's autocompletion which can be found in the contrib/completion-folder of the source distribution. After you've sourced it (through your .bashrc, .bash_profile, ...) you get some additional functions and esp. ``__git_ps1`` which gives you the current branch name in nice brackets. If you take a look into the ``git-completion.bash`` file, also look at this paragraph::
    
    #    4) Consider changing your PS1 to also show the current branch:
    #        PS1='[\u@\h \W$(__git_ps1 " (%s)")]\$ '
    #
    #       The argument to __git_ps1 will be displayed only if you
    #       are currently in a git repository.  The %s token will be
    #       the name of the current branch.
    
So add ``$(__git_ps1 " (%s)")`` to your PS1 and you should all be set. It's also sweet in this regard, that it also works when you're some folders within a git repository. Thanks Martin for this trick :D