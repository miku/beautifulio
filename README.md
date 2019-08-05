# Beautiful IO

Talk about beautiful IO with Go.

* 2019-10-21, 10:45-11:30, Firenze (Lungarno del Tempio, 44, 50121 Firenze FI, Italy)

![](f.jpg)

> Love letters have been written to parts of the IO subsystem.

* [Asynchronously Split an io.Reader in Go (golang)](https://rodaine.com/2015/04/async-split-io-reader-in-golang/)

> I have fallen in love with the flexibility of io.Reader and io.Writer when
> dealing with any stream of data in Go.

Target: 30-35 minutes, 15 sentences pm, 450 sentences, written down: 10 pages,
maybe 50-80 slides, 40 slides and 15 minutes of coding.

# Learning goals

* io package overview (go doc)
* how custom readers are implemented (stdlib)
* how standard library readers and writers are implemented (stdlib)
* what are some strange implementations of readers and writer (scan github)
* examples of how you can combine interfaces
* examples where the flexibility pays off
* memory efficient data processing pipelines

Examples:

* image filters, real time (combinations)
* GAN generated images from TF
* function plotter

What rough classes of readers are there?

* persistence (files, databases, ...)
* transformation (compression, hashing, image filter)
* synthetic data
* distortion (flaky readers, slow internet connection, ...)

Performance considerations:

* ReaderFrom
* Allocations, buffer reuse
* ring buffers

# Run a few benchmarks

* Copy, CopyBuffer, CopyN - with various sizes

# Snippets

> Package io provides basic interfaces to I/O primitives.


----

This might not be relevant, as we only use zero file (Copy, CopyBuffer).

```
goos: linux
goarch: amd64
pkg: github.com/miku/beautifulio/copybench
BenchmarkSimple/size-1048576-4          20000000                50.0 ns/op
BenchmarkSimple/size-10485760-4         30000000                50.3 ns/op
BenchmarkSimple/size-104857600-4        30000000                51.2 ns/op
BenchmarkSimple/size-1073741824-4       20000000                54.5 ns/op
BenchmarkSimple/size-10737418240-4      20000000                60.1 ns/op
BenchmarkSimple/size-53687091200-4      10000000               117 ns/op
BenchmarkCopyBuffer/size-1048576-buf-1-4                30000000                54.8 ns/op
BenchmarkCopyBuffer/size-1048576-buf-8-4                30000000                51.8 ns/op
BenchmarkCopyBuffer/size-1048576-buf-1024-4             30000000                52.2 ns/op
BenchmarkCopyBuffer/size-1048576-buf-4096-4             30000000                54.7 ns/op
BenchmarkCopyBuffer/size-1048576-buf-8192-4             30000000                54.6 ns/op
BenchmarkCopyBuffer/size-1048576-buf-16384-4            30000000                52.2 ns/op
BenchmarkCopyBuffer/size-1048576-buf-32768-4            30000000                54.7 ns/op
BenchmarkCopyBuffer/size-10485760-buf-1-4               30000000                51.7 ns/op
BenchmarkCopyBuffer/size-10485760-buf-8-4               30000000                52.4 ns/op
BenchmarkCopyBuffer/size-10485760-buf-1024-4            30000000                54.7 ns/op
BenchmarkCopyBuffer/size-10485760-buf-4096-4            30000000                55.0 ns/op
BenchmarkCopyBuffer/size-10485760-buf-8192-4            30000000                52.2 ns/op
BenchmarkCopyBuffer/size-10485760-buf-16384-4           30000000                52.2 ns/op
BenchmarkCopyBuffer/size-10485760-buf-32768-4           30000000                55.2 ns/op
BenchmarkCopyBuffer/size-104857600-buf-1-4              20000000                52.2 ns/op
BenchmarkCopyBuffer/size-104857600-buf-8-4              30000000                55.3 ns/op
BenchmarkCopyBuffer/size-104857600-buf-1024-4           30000000                55.3 ns/op
BenchmarkCopyBuffer/size-104857600-buf-4096-4           20000000                55.4 ns/op
BenchmarkCopyBuffer/size-104857600-buf-8192-4           20000000                52.3 ns/op
BenchmarkCopyBuffer/size-104857600-buf-16384-4          30000000                52.2 ns/op
BenchmarkCopyBuffer/size-104857600-buf-32768-4          30000000                52.3 ns/op
BenchmarkCopyBuffer/size-1073741824-buf-1-4             20000000                55.8 ns/op
BenchmarkCopyBuffer/size-1073741824-buf-8-4             20000000                53.0 ns/op
BenchmarkCopyBuffer/size-1073741824-buf-1024-4          20000000                53.6 ns/op
BenchmarkCopyBuffer/size-1073741824-buf-4096-4          20000000                53.8 ns/op
BenchmarkCopyBuffer/size-1073741824-buf-8192-4          20000000                55.6 ns/op
BenchmarkCopyBuffer/size-1073741824-buf-16384-4         20000000                56.0 ns/op
BenchmarkCopyBuffer/size-1073741824-buf-32768-4         20000000                52.7 ns/op
BenchmarkCopyBuffer/size-10737418240-buf-1-4            20000000                61.9 ns/op
BenchmarkCopyBuffer/size-10737418240-buf-8-4            20000000                61.0 ns/op
BenchmarkCopyBuffer/size-10737418240-buf-1024-4         20000000                58.7 ns/op
BenchmarkCopyBuffer/size-10737418240-buf-4096-4         20000000                58.5 ns/op
BenchmarkCopyBuffer/size-10737418240-buf-8192-4         20000000                61.5 ns/op
BenchmarkCopyBuffer/size-10737418240-buf-16384-4        20000000                59.1 ns/op
BenchmarkCopyBuffer/size-10737418240-buf-32768-4        20000000                62.9 ns/op
BenchmarkCopyBuffer/size-53687091200-buf-1-4            10000000               112 ns/op
BenchmarkCopyBuffer/size-53687091200-buf-8-4            10000000               119 ns/op
BenchmarkCopyBuffer/size-53687091200-buf-1024-4         10000000               113 ns/op
BenchmarkCopyBuffer/size-53687091200-buf-4096-4         10000000               120 ns/op
BenchmarkCopyBuffer/size-53687091200-buf-8192-4         10000000               114 ns/op
BenchmarkCopyBuffer/size-53687091200-buf-16384-4        10000000               114 ns/op
BenchmarkCopyBuffer/size-53687091200-buf-32768-4        10000000               114 ns/op
PASS
ok      github.com/miku/beautifulio/copybench   186.691s
```
