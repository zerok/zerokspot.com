---
title: Advisory File locking on Linux
date: "2020-09-19T22:24:31+02:00"
tags:
- linux
- unix
---

Many folks have their first contact with file-locking when it goes wrong and a lock is flying around that should no longer relevant. In general, lock-files are used to prevent two processes or threads from doing the same operation. An extreme example is to prevent two processes from reading a file at the same time.

In my case, I ran into an issue where a shell script triggered by cronjob wanted to replace a file while a Java service was reading that same file. My initial thought was, that this should be a job for the `flock` command-line utility inside the cronjob. That would create a lock on that file while inside Java I’d check for that lock and acquire a read-only lock.

Inside a Java application, you’d use the FileLock object to create a lock:

	import java.io.RandomAccessFile;
	import java.nio.channels.FileLock;
	import java.io.File;
	
	RandomAccessFile f = new RandomAccessFile(new File("testfile"), "rw");
	FileLock lock = f.getChannel().tryLock();
	if (lock == null) {
		// Locking was not possible
	}
	lock.unlock();

At the same time, inside the cronjob I’d use something like this:

	flock testfile ./otherscript.sh

So the idea. Turns out, things aren’t that simple when it comes to file-locking. And so began my little dive into that topic…

On GNU/Linux are two distinct kinds of file-locking. *Mandatory* file-locking does not require any explicit support by applications who want to use them and is enforced by the operating system’s kernel/the file-system on every read, write, and open call to a given file-descriptor.

At the other end of the brutality spectrum is *advisory* file-locking. Here, processes that want to operate on a specific file follow a certain procedure before actually accessing the file. Using either the flock() or fcntl() syscall they first try to acquire a lock and once they’re done they release it again.

While mandatory file-locking is available on various Unix-systems and usually dependent on the used file-system, advisory locking is far more common.

## flock() vs fcntl()

The two syscalls are different in what kind of feature-set they provide. Sadly, since quite some time they are not compatible with each other:

> Since kernel 2.0, there is no interaction between the types of lock placed by flock(2) and fcntl().
> — [fcntl(2) manpage](https://linux.die.net/man/2/fcntl)

That being said, both can be listed by looking at the `/proc/locks` pseudo-file:

	$ cat /proc/locks
	1: FLOCK  ADVISORY  WRITE 1152 fc:01:2054788 0 EOF
	2: FLOCK  ADVISORY  WRITE 1181 fc:01:2613083 0 EOF
	3: FLOCK  ADVISORY  WRITE 1181 fc:01:2613082 0 EOF
	4: FLOCK  ADVISORY  WRITE 1181 fc:01:2613020 0 EOF
	5: FLOCK  ADVISORY  WRITE 1181 fc:01:1837577 0 EOF
	6: FLOCK  ADVISORY  WRITE 1181 fc:01:1837566 0 EOF
	7: OFDLCK ADVISORY  READ -1 00:06:6 0 EOF
	8: FLOCK  ADVISORY  WRITE 1178 fc:01:17282 0 EOF
	9: POSIX  ADVISORY  WRITE 1175 00:18:608 0 EOF
	10: POSIX  ADVISORY  WRITE 1145 00:18:591 0 EOF
	11: FLOCK  ADVISORY  WRITE 1124 00:18:559 0 EOF
	12: POSIX  ADVISORY  WRITE 671 00:18:474 0 EOF

The output shown above contains 12 locks in total with three different kinds. `POSIX` and `OFDLCK` are created using the `fcntl()` syscalls while `FLOCK` is created using `flock()`.

### flock()

[flock()](https://linux.die.net/man/2/flock) originated on BSD and is pretty simple where you just apply a certain operation to a specific file-descriptor:

	int flock(int fd, int operation);

The available operations are:

- `LOCK_SH` for creating a shared lock
- `LOCK_EX` for creating an exclusive lock
- `LOCK_UN` for removing a lock

All of these can also be combined (by ORing) with `LOCK_NB` in order to make the syscall non-blocking. Locks are associated with the process that is creating them and there can only be a single lock associating a specific process with a specific file.

The `flock` command creates, who’d have guessed, a FLOCK lock.

### fcntl()

Fcntl, on the other hand, is far more versatile as it is intended as a general-propose file manipulator:

	int fcntl(int fd, int cmd, ... /* arg */ );

Above I wrote that fcntl can be used for POSIX-compatible and Open File Descriptor locks (OFDLCK). I will focus on POSIX locks here since they are relevant for my use-case.

As the signature is already hinting at, the call can have a multitude of arguments depending on the use-case. For POSIX-compatible advisory locking, the signature would look like this:

	int fcntl(int fd, int cmd, flock_t* lock);

… with these three commands being relevant:

- `F_GETLK`
- `F_SETLK`
- `F_SETLKW`

The lock argument points at a structure that is either used for writing the output from or passing additional arguments. You can mark locks as shared or exclusive but the main “new” feature that fcntl has over flock is that you can only share certain parts of a file descriptor (like the first 10 bytes).

Cool, but not relevant for my use-case 😅 Anyway, the Java FileLock is a POSIX-one on macOS and Linux.

## The “solution”

OK, so I’d either have to somehow lock using fcntl in the cronjob or go with a completely different approach to handling lock files. In the end I could avoid the need for a lock file altogether but not before just implementing a little experimental flock-alternative that uses fcntl called [plock](https://github.com/zerok/plock).

	$ go run main.go --exclusive /tmp/lockfile ./test.sh &
	$ cat /proc/locks
	1: POSIX  ADVISORY  WRITE 214 00:70:2627719 0 EOF
	2: OFDLCK ADVISORY  READ  -1 00:06:1028 0 EOF

So while I didn’t need place in the end, this was at least quite an informative experiment for me and I hope this post included something new for you too 😊

Perhaps I’m motivated to write plock in Rust again just to see how. But that’s something for another time.
