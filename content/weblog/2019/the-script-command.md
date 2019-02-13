---
title: "The script command"
date: 2019-02-13T19:08:25+01:00
tags:
- unix
- tools
---

While doing a testing session this afternoon,
[Chris](https://twitter.com/kautsig) mentioned a little command that I
hadn't used before: `script` allows you to record everything that
happens inside a terminal session and to play it back afterwards.

Turns out, it is shipped by default with macOS and most Linux
distributions (e.g. as part of the
[bsdutils](https://packages.debian.org/stretch/bsdutils) package). If
all you want to do, is keep a look of commands and their respective
output, just run `script OUTPUT_FILE`. Now, all the commands you enter
before closing the terminal session (end with `exit` or Ctrl-d) will
be recorded:

```
$ script session.typescript
Script started, output file is session.typescript

$ echo "hello"
hello

$ exit
Script done, output file is session.typescript
```

If you open `session.typescript` inside an editor you will see not
only the commands and their output but also your applied shell prompt!
Basically as much as possible of what was visible in the original
terminal session.

Sadly, if you want to go beyond the pure logging of commands then the
`script` command on macOS has a different parameter set compared to
the one being available for instance on a Debian installation.


## Recording a session (for replay)

In order to record a session on macOS that should also be replayable,
you would use the following command:

```
$ script -r session.typescript
```

On Linux you will actually have to instruct script to generate two
files: (1) that holds the commands and their respective output and (2)
a timing-file that stores when what should be displayed during replay:

```
$ script --timing=session.timing session.typescript
```

## Playing back a recording

Since the two commands produce different output files, the
replay-functionality also expects different inputs.

On macOS you can reuse the `script` command but with the `-p` flag:

```
$ script -p session.typescript
```

On Linux replaying is handled by a different command that expects
both, the timing file *and* the typescript:

```
$ scriptreplay --timing session.timing \
    --typescript session.typescript
```

## Platform-independent recording

If you need to share a replay between platforms, it's probably worth
taking a look at [asciinema](https://asciinema.org/):

```
# Record
$ asciinema rec session.cast

# Playback
$ asciinema play session.cast
```

asciinema stores the recording in a simple JSON format. You can find
its specification
[here](https://github.com/asciinema/asciinema/blob/master/doc/asciicast-v1.md).
