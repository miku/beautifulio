# Beautiful IO

Talk about beautiful IO with Go.

* 2019-10-21, 10:45-11:30, Firenze (Lungarno del Tempio, 44, 50121 Firenze FI, Italy)

![](f.jpg)

> Love letters have been written to parts of the IO subsystem.

* [Asynchronously Split an io.Reader in Go (golang)](https://rodaine.com/2015/04/async-split-io-reader-in-golang/)

> I have fallen in love with the flexibility of io.Reader and io.Writer when
> dealing with any stream of data in Go.

Target: 30-35 minutes, 15 sentences pm, 450 sentences, written down: 10 pages.

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
