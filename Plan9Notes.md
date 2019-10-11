# On Plan 9

* http://doc.cat-v.org/plan_9/

With an eye on IO.

> Plan 9 began in the late 1980’s as an attempt to have it both ways: to build
a system that was centrally administered and cost-effective using cheap modern
microcomputers as its computing elements. The idea was to build a time-sharing
system out of workstations, but in a novel way.

> The early catch phrase was to build a UNIX out of a lot of little systems,
not a system out of a lot of little UNIXes.

> The problems with UNIX were too deep to fix, but some of its ideas could be
brought along. The best was its use of the file system to coordinate naming of
and access to resources, even those, such as devices, not traditionally treated
as files.

> The view of the system is built upon three principles. First, resources are
named and accessed like files in a hierarchical file system. Second, there is
a standard protocol, called 9P, for accessing these resources. Third, the
disjoint hierarchies provided by different services are joined together into
a single private hierarchical file name space. The unusual properties of Plan
9 stem from the consistent, aggressive application of these principles.

> Plan 9 provides the mechanism to assemble a personal view of the public space
with local names for globally accessible resources. Since the most important
resources of the network are files, the model of that view is file-oriented.

> The services available in the network all export file hierarchies.

> This is a different style of use from the idea of a ‘uniform global name
space’.

This reminds me of Erlang, Joe Armstrong single namespace for all functions.
Maybe PHP is similar, but surely not what JA envisioned.

> Similarly, in Plan 9 the name /dev/cons always refers to the user’s terminal
and /bin/date the correct version of the date command to run, but which files
those names represent depends on circumstances such as the architecture of the
machine executing date.

## File Server

> A central file server stores permanent files and presents them to the network
as a file hierarchy exported using 9P. The server is a stand-alone system,
accessible only over the network, designed to do its one job well. It runs no
user processes, only a fixed set of routines compiled into the boot image.
Rather than a set of disks or separate file systems, the main hierarchy
exported by the server is a single tree, representing files on many disks. That
hierarchy is shared by many users over a wide area on a variety of networks.
Other file trees exported by the server include special-purpose systems such as
temporary storage and, as explained below, a backup service.

Ok, they implemented some autobackup facility, like a gitfs.

> The most unusual feature of the file server comes from its use of a WORM
device for stable storage. Every morning at 5 o’clock, a dump of the file
system occurs automatically. The file system is frozen and all blocks modified
since the last dump are queued to be written to the WORM. Once the blocks are
queued, service is restored and the read-only root of the dumped file system
appears in a hierarchy of all dumps ever taken, named by its date. For example,
the directory /n/dump/1995/0315 is the root directory of an image of the file
system as it appeared in the early morning of March 15, 1995. It takes a few
minutes to queue the blocks, but the process to copy blocks to the WORM, which
runs in the background, may take hours.

It's also RO, so a bit functional, or like Hickey says, database as a value,
a snapshot as a value.

Here again a hint at the advantages of keeping track.

> People feel free to make large speculative changes to files in the knowledge
that they can be backed out with a single copy command.

Interesting take on designing around growth:

> Once a file is written to WORM, it cannot be removed, so our users never see
‘‘please clean up your files’’ messages and there is no df command. We regard
the WORM jukebox as an unlimited resource. The only issue is how long it will
take to fill. Our WORM has served a community of about 50 users for five years
and has absorbed daily dumps, consuming a total of 65% of the storage in the
jukebox. In that time, the manufacturer has improved the technology, doubling
the capacity of the individual disks. If we were to upgrade to the new media,
we would have more free space than in the original empty jukebox. Technology
has created storage faster than we can use it.

Notion of a single machine.

> The same operating system runs on all hardware. Except for performance, the
appearance of the system on, say, an SGI workstation is the same as on
a laptop. Since computing and file services are centralized, and terminals have
no permanent file storage, all terminals are functionally identical. In this
way, Plan 9 has one of the good properties of old timesharing systems, where
a user could sit in front of any machine and see the same system. In the modern
workstation community, machines tend to be owned by people who customize them
by storing private information on local disk. We reject this style of use,
although the system itself can be used this way. In our group, we have
a laboratory with many public-access machines—a terminal room—and a user may
sit down at any one of them and work.

----

## Parallel Programming bits

> Parallel programs have three basic requirements: management of resources
shared between processes, an interface to the scheduler, and fine-grain process
synchronization using spin locks. On Plan 9, new processes are created using
the rfork system call. Rfork takes a single argument, a bit vector that
specifies which of the parent process’s resources should be shared, copied, or
created anew in the child. The resources controlled by rfork include the name
space, the environment, the file descriptor table, memory segments, and notes
(Plan 9’s analog of UNIX signals). One of the bits controls whether the rfork
call will create a new process; if the bit is off, the resulting modification
to the resources occurs in the process making the call. For example, a process
calls rfork(RFNAMEG) to disconnect its name space from its parent’s. Alef uses
a fine-grained fork in which all the resources, including memory, are shared
between parent and child, analogous to creating a kernel thread in many
systems.

A rendezvouz system call?

> The rendezvous system call provides a way for processes to synchronize. Alef
uses it to implement communication channels, queuing locks, multiple
reader/writer locks, and the sleep and wakeup mechanism. Rendezvous takes two
arguments, a tag and a value. When a process calls rendezvous with a tag it
sleeps until another process presents a matching tag. When a pair of tags
match, the values are exchanged between the two processes and both rendezvous
calls return. This primitive is sufficient to implement the full set of
synchronization routines.

Early use cases of threading.

> A Plan 9 process in a system call will block regardless of its ‘weight’. This
means that when a program wishes to read from a slow device without blocking
the entire calculation, it must fork a process to do the read for it. The
solution is to start a satellite process that does the I/O and delivers the
answer to the main program through shared memory or perhaps a pipe. This sounds
onerous but works easily and efficiently in practice; in fact, most interactive
Plan 9 applications, even relatively ordinary ones written in C, such as the
text editor Sam [Pike87], run as multiprocess programs.

Kernel structures around networks.

> The kernel plumbing used to build Plan 9 communications channels is called
streams [Rit84][Presotto]. A stream is a bidirectional channel connecting
a physical or pseudo-device to a user process. The user process inserts and
removes data at one end of the stream; a kernel process acting on behalf of
a device operates at the other end. A stream comprises a linear list of
processing modules. Each module has both an upstream (toward the process) and
downstream (toward the device) put routine. Calling the put routine of the
module on either end of the stream inserts data into the stream. Each module
calls the succeeding one to send data up or down the stream. Like UNIX streams
[Rit84], Plan 9 streams can be dynamically configured.

Oh, a replacement for TCP. In 9: 9P/IL.

> In Plan 9, the implementation of IL is smaller and faster than TCP. IL is our
main Internet transport protocol.

# The Use of Name Spaces in Plan 9

* http://doc.cat-v.org/plan_9/4th_edition/papers/names

> **All resources in Plan 9 look like file systems. That does not mean that they
are repositories for permanent files on disk, but that the interface to them is
file-oriented: finding files (resources) in a hierarchical name tree, attaching
to them by name, and accessing their contents by read and write calls. There
are dozens of file system types in Plan 9, but only a few represent traditional
files**. At this level of abstraction, files in Plan 9 are similar to objects,
except that files are already provided with naming, access, and protection
methods that must be created afresh for objects. Object-oriented readers may
approach the rest of this paper as a study in how to make objects look like
files.

What is a device file in UNIX?

> The file server is only one type of file system. A number of unusual
services are provided within the kernel as local file systems. These services
are not limited to I/O devices such as disks. They include network devices and
their associated protocols, the bitmap display and mouse, a representation of
processes similar to /proc [Killian], the name/value pairs that form the
‘environment’ passed to a new process, profiling services, and other resources.
Each of these is represented as a file system — directories containing sets of
files — but the constituent files do not represent permanent storage on disk.
Instead, they are closer in properties to UNIX device files.


