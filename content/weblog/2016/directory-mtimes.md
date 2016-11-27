---
date: 2016-11-27T08:13:48+01:00
title: "TIL: Directory modification times"
tags:
- til
- filesystems
---

Recently, I needed a way to check when the content of a folder had last been
updated. If it hadn't seen any changes within the last couple of days I wanted
to remove it.

One solution for this would be to recursively check the modification time of
every single file in that folder but I also got curious about what actually
defines the modification time of the folder itself.

As this is something that can be different from file-system to file-system,
let's look at some of the more popular systems for macOS, Linux, and
Windows. All of them are basically handling the mtime of a folder in pretty much
the same way:

The mtime of a folder is only modified if a file directly below it is added,
renamed, or removed. Changing anything about the content of any of the files or
their attributes won't change the mtime of the directory that contains them
([Source](http://stackoverflow.com/questions/3620684/directory-last-modified-date)).

To verify this I've created a [little Python script][] and tried it on my local
machine (macOS), a Linux box, and a Windows box.

[little python script]: https://gist.github.com/zerok/420056668b71a2d057233e101c62d445

There still seem to be some differences between the various file-systems,
though, esp. regarding the amount of precision they provide.

For HFS+ (macOS) it was pretty easy to find some details in the [specs][] that
specify the basic behavior:

[specs]: http://web.archive.org/web/20090214212148/http://developer.apple.com/technotes/tn/tn1150.html

> contentModDate: The date and time the folder's contents were last
> changed. This is the time when a file or folder was created or deleted inside
> this folder, or when a file or folder was moved in or out of this folder. See
> HFS Plus Dates for a description of the format.
>
> attributeModDate: The last date and time that any field in the folder's
> catalog record was changed. An implementation may treat this field as
> reserved. In Mac OS X, the BSD APIs use this field as the folder's change time
> (returned in the st_ctime field of struct stat). All versions of Mac OS 8 and
> 9 treat this field as reserved. See HFS Plus Dates for a description of the
> format.
>
> [...]
>
> HFS Plus Dates:
> HFS Plus stores dates in several data structures, including the volume header
> and catalog records. These dates are stored in unsigned 32-bit integers
> (UInt32) containing the number of seconds since midnight, January 1, 1904,
> GMT. This is slightly different from HFS, where the value represents local
> time.

Compared to more modern systems the mtime et al. are only exposed in seconds,
while others are exposing nanoseconds. That also shows when you look at the
`st_mtime_ns` value in Python (inspecting one of the source files of this blog):

```
>>> import pathlib
>>> pathlib.Path('32c3.md').stat().st_mtime_ns
1473583000000000000
```

[ext4](https://www.kernel.org/doc/Documentation/filesystems/ext4.txt) (Linux) on
the other hand offers nanosecond precision through some additional fields which
are then merged on the API layer for Python et al. to use:

```
>>> import pathlib
>>> pathlib.Path("README.md").stat().st_mtime_ns
1457287154109435711
```

On Windows NTFS seems to handle modification times in a similar way to ext4 and
HFS+ according to
the
[Description of NTFS date and time stamps for files and folders](https://support.microsoft.com/en-us/kb/299648) document,
but I couldn't yet find anywhere in what precision that value is actually
stored. It is exposed as nanoseconds according
to
[File Times](https://msdn.microsoft.com/en-us/library/windows/desktop/ms724290.aspx) and
also the little Python snippet returned a result similar to ext4.

While looking into all that I also learnt something about how the different
file-systems handle things like creation time ... a whole other can of worms,
though ([some details][]).

[some details]: http://superuser.com/questions/973547/how-can-i-display-all-8-ntfs-timestamps

Personally, I'm just glad that mtime doesn't seem to come close to the
configurable nightmare that atime is.
