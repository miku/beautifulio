# Beautiful IO

> A tour through standard library pkg/io and various implementations of its interfaces.

Golab 2019, 2019–10–21, Florence
[Martin Czygan](mailto:martin.czygan@gmail.com)

<!-- Die Brautleute; short summaries at the beginning of the sections -->

----

# About me

SWE [@ubleipzig](https://ub.uni-leipzig.de) working mostly with Python and Go.

Taming data – open source – writing.

> [Explore IO](https://github.com/miku/exploreio) workshop at Golab 2017.

----

# Background

* Go Proverbs (2015)

> The bigger the interface, the weaker the abstraction.

Prominent examples are `io.Reader` and `io.Writer`.

----

# The IO package

* contains basic, widely used interfaces (within and outside standard library)
* utility functions

----

# Why beautiful?

> La bellezza è negli occhi di chi guarda

* small, versatile interfaces
* composable

----

# Love and praise

> This article aims to convince you to use io.Reader in your own code wherever
> you can. -- [@matryer](https://twitter.com/matryer)

> "Crossing Streams: a love letter to Go io.Reader" -- [@jmoiron](https://twitter.com/jmoiron)

> Which brings me to io.Reader, easily my favourite Go interface. --
> [@davecheney](https://twitter.com/davecheney)

----

# What's in pkg/io?

<!--
$ go doc io | grep ^type | wc -l
25
-->

* 25 types
* 21/25 are interfaces
* 12 functions, 3 constants, 6 errors

The concrete types are: `LimitedReader`, `PipeReader`, `PipeWriter`,
  `SectionReader`; functions: `Copy`, `CopyN`, `CopyBuffer`, `Pipe`,
  `ReadAtLeast`, `ReadFull`, `WriteString`, `LimitReader`, `MultiReader`,
  `TeeReader`, `NewSectionReader`, `MultiWriter`

----

# A few Interfaces


![](static/iointftab.png)

----

# Missing interfaces

You might find some missing pieces elsewhere.

![](static/go4extra.png)

----

# How many readers, writers are there?

```shell
$ guru -json implements /usr/local/go/src/io/io.go:#3309
```

I counted over 280 implementations of io.Reader and X of io.Writer.

----

# What is a Reader?

```go
type Reader interface {
        Read(p []byte) (n int, err error)
}
```

The reader implementation will populate a given byte slice.

* at most `len(p)` bytes are read
* to signal the end of a stream, return `io.EOF`
* `n` might or might not be zero, when `io.EOF` is returned

----



