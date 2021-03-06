---
date: '2007-09-20T12:00:00-00:00'
language: en
tags:
- http
- python
- twisted
title: Getting to know Twisted
---


Lately I've been in need for some simple client-server solutions that do specific jobs like transcoding a video using mencoder or simply adding an AVI container to a FLV videostream and MP3 audiostream. So I started taking a close look at the *Twisted* framework for Python ...

-------------------------------

[Twisted](http://twistedmatrix.com/trac/) is a long-standing project in the Python community that is specialised on networking. Therefor it offers solutions for remote objects, simple UDP or TCP communication and even some a little bit more high-level HTTP implementations. 

First of all: Why HTTP? Simple: You can send and receive data using common cli tools like cURL and it gives you an easy way to separate data from its metadata. 

Why did I choose Twisted for these jobs and not just any HTTP-server that I could abuse? Twisted has the advantage that it lets me down to the bytestream-level if I want to, yet also offers me enough abstraction if I don't want the byte-level. For example: For the tasks I mentioned above I need first of all huge amounts of data to be transferred to and from the server (think 300MB videos etc.) but I also needed some metadata like the size of the video in order to know how much content was to be expected to arrive through the pipe. 

For something like this HTTP POST is lightweight enough for my taste (although the first implementation simply reserved 8 Bytes for the size of the video and prepended it to the datastream :P ). The problem was, though, that most other HTTP implementations do too much. They parse the POST-body and therefor have to load the whole video into memory before they notice that something weird is going on here ;-) 

For such occasions where you first want to process lines (like the headers in HTTP) but the rest of the data on the byte-level (video in the body), Twisted offers a nice protocol called LineReceiver (`twisted.protcols.basic.LineReceiver`). LineReceiver offers two modes: the lineMode and the rawDataMode and you can switch between those two whenever you want.

Before I post an example for this, some design details about Twisted (at least as far as I understand it by now). Twisted is based on 3 basic concepts:

1. The *Protocol* handles the stream of data and basically is a dispatcher for actions to be taken based on what comes in.
2. The *Factory* is a factory for protocols and also holds the state of the system (since the protocol is re-instantiated whenever a new connection is opened).
3. When you do things the client-server-way, it's often because you want the server to take some load from the client and give the server this potentially longer job. If you have such a long-running job, Twisted has the *Deferred* for you. This implements basically a callback system that let's you trigger a function when the job was finished successfully or not.

Now on to a small code example for a simple HTTP-inspired server. The main part is the Protocol. What should this little example do? For starters, let's only accept POST requests and return the content-length. Not really useful, but it will show you a simple Twisted server.

To get started first of all we need a class that inherits the LineReceiver
protocol. Since we only care about the headers right now (and those are lines) we can easily adapt that protocol to our needs by also implementing the lineReceived method which gets called every time the server receives a new line.

    
    from twisted.protcols.basic import LineReceiver
    
    class MyLineReceiver(LineReceiver):
        def __init__(self):
            pass
        def lineReceived(self, line):
            pass            
            
2 methods. That's all you need. Now let's focus on the lineReceived method for now. It gets the received line passed so let's start messing around with it. First of all we want the Protocol if it gets something else than a POST request the usual 405 error message:

    
    def lineReceived(self, line):
        self._lineno += 1
        if self._lineno == 1:
            self._method == line.split(" ")[0].upper()
            if self._method != 'POST':
                self.sendLine('HTTP/1.0 405 Method Not Allowed')
                self.sendLine('')
                self.transport.loseConnection()
            else:
                self.sendLine('HTTP/1.0 200 OK')
                self.sendLine('')
                self.transport.loseConnection()
        
The LineReceiver protocol offers the sendLine method which allows you to write lines ... big suprise. Every protocol also has a transport member which lets you play with the underlying data connection. And also lets you close it with loseConnection ;-) Nothing more to discuss here. Let's move on.

    
    class MyLineReceiver(LineReceiver):
        def __init__(self):
            self._lineno = 0
            self._method = None
            self._headers = {}

        def lineReceived(self, line):
            self._lineno += 1
            if self._lineno == 1:
                self._method = line.split(" ")[0].upper()
                if self._method != 'POST':
                    self.sendLine('HTTP/1.0 405 Method Not Allowed')
                    self.sendLine('')
                    self.transport.loseConnection()
            else:
                if len(line)==0:
                    self.sendLine('HTTP/1.0 200 OK')
                    self.sendLine('Content-Type: text/plain')
                    self.sendLine('')
                    self.sendLine(self._headers.get('CONTENT-LENGTH','0'))
                    self.transport.loseConnection()
                else:
                    tmp = line.split(":")
                    key = tmp[0].upper()
                    value = ':'.join(tmp[1:])
                    self._headers[key] = value
    

That's the whole protocol. Just basic line parsing and finding the Content-Length header. If you now also want to handle huge amounts of post data, just switch the mode in the line where you notice that you got your first empty line using `self.setRawMode()`. From now on all the data will be send to the rawDataReceived method instead of the lineReceived method. That's it :-)

Now let's get this whole thing running. There are a couple of components that are required for this. First of all, the factory:

    
    from twisted.internet.protocol import ServerFactory
    
    class MyFactory(ServerFactory):
        protocol = MyLineReceiver
        
    
There isn't anything else needed for the factory, but to set its protocol member to the class of your protocol. Nearly done now. Only thing missing? The actual server :-)

There you have more or less two options: Going for a completely standalone server, or using the [twistd](http://twistedmatrix.com/projects/core/documentation/howto/basics.html#auto1) component that is supposed to make it easy to manage various server processes (haven't tried it yet, but I simply like its debug output ;-)):


    if __name__ == '__main__':
        from twisted.internet import reactor
        reactor.listenTCP(8765, MyFactory())
        reactor.run()
    else:
        from twisted.application import service, internet
        application = service.Application('dummyserver')
        internet.TCPServer(8765, MyFactory()).setServiceParent(application)
        

After merging all  this into [one file](/media/2007/dummytwistedserver.py.txt), just run `python dummytwistedserver.py` or `twistd -noy dummytwisttedserver.py` and you can easily mess with it using cURL :-)

**Note:** I only started to mess around with Twisted 2 days ago, so what I'm
describing right here might not be the best way to solve the problems
mentioned.

A good starting point to learn more is the [Twisted Document](http://twistedmatrix.com/projects/core/documentation/howto/index.html). While quite minimal, it gives good hints into the right direction. Same goes for the [API Reference](http://twistedmatrix.com/documents/current/api/).
