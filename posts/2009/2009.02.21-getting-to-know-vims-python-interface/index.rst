Getting to VIM's Python interface
#################################

Scripting is usually one of the selling points for those types of text editors
that are specifically targeted at power-users. But most of the time you have
to learn a whole new scripting language to really be able to use it. Imagine
my surprise when I noticed that `VIM`_ besides having it's own scripting
language also let's you also write plugin in Perl, Python, Ruby and TCL. This
way you don't really have to learn a new language for scripting VIM, although
you might still want to look into it because there are a few problems.

Since I like Python quite a bit, the rest of this post will be mostly focused
on the Python interface.

.. _vim: http://www.vim.org/

-------------------------------

To be able to script in other languages, VIM has to be specifically compiled
with support for the language of your choice. Being a bit lazy here I simply
got myself a nice Python-enabled VIM with `MacVIM`_. Using MacPorts you can also get the something similar for the shell with this ::
   
   $ sudo port install vim +python

Once you have that, you get a small set of new commands:

* python

* pyfile

The first lets you execute Python code directly while the latter will load a
given Python file. 

If you want to write a plugin, you will most likely also like to take a look
at the ``vim`` module the interface offers, which let's you access things like
the current buffer or the current line (``vim.current.line``). You can also
get down to the current cursor position with ``vim.current.window.cursor``. 

Using the new commands mentioned above  you can easily define for instance a
function within a VIM plugin and then bind a command to it::
    
    python << EOF
    def _my_function(some_arg=None):
        import vim
        print len(vim.current.buffer)
    EOF
    command -nargs=* MyCommand :python _my_function(<f-args>)

The documentation of all that is actually quite good, so I will just defer to
`the official instructions
<http://vimdoc.sourceforge.net/htmldoc/if_pyth.html>`_. I just want to note a
few observations, though.

First of all, the Python interface seems not to offer you all the things you
can do in the native language for writing plugins. For instance: buffers don't
implement everything you'd normally expect from list objects in Python and
therefor there is for instance no such thing as ``buffer.insert``. This means
that you can't really create a new lines within a buffer as easily as with the
native ``append(line_num, text)`` function.

The Python interface has a ``buffer.append`` method but it only lets you
append content to the end of the buffer instead of allowing append/insert
operations at an arbitrary line number. In practice I simple resorted to
something like this for creating new lines after a given line-number::
    
    line_num = 123
    cmd = '%ds/$/\r/' % (line_num, )
    vim.command(cmd)

Another problem I ran into was something that I couldn't yet solve. For some
reason, I can't actually modify a buffer (either the current or any other line
inside a buffer)::
    
    :py import vim
    :py vim.current.line = '123'
    Traceback (most recent call last):
      File "<string>", line 1, in <module>
    TypeError: bad argument type for built-in operation

The interesting thing about this is that it only happens when I'm using the
terminal version of VIM. In MacVIM everything works just fine. And this isn't
really limited to the Python interface. I also tried the Ruby one and got
this::
    
    :ruby VIM::Buffer.current.line = '123'
    TypeError: can't modify frozen string

But there it only happens when I'm operating for instance on unsaved buffers.
If someone has a solution for this, please let me know :-) For now I had
to rewrite a small plugin I just wrote today in order to get it to work
*everywhere*. 

.. _macvim: http://code.google.com/p/macvim/