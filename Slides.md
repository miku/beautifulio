# Beautiful IO

> A tour through standard library pkg/io and various implementations of its interfaces.

Golab 2019, 2019–10–21, Firenze
[Martin Czygan](mailto:martin.czygan@gmail.com)

<!-- Die Brautleute; short summaries at the beginning of the sections -->

----

# About me

SWE [@ubleipzig](https://ub.uni-leipzig.de) working mostly with Python and Go.

Taming data – open source – writing.

> [Explore IO](https://github.com/miku/exploreio) workshop at Golab 2017.

----

# The IO package

* contains basic, widely used interfaces (within and outside standard library)
* simple, flexible helpers

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

* 25 types - of which 21 are interfaces, 12 functions, 3 constants, 6 errors
* Four concrete types are: `LimitedReader`, `PipeReader`, `PipeWriter`,
`SectionReader`.
* Functions: `Copy`, `CopyN`, `CopyBuffer`, `Pipe`, `ReadAtLeast`, `ReadFull`,
  `WriteString`, `LimitReader`, `MultiReader`, `TeeReader`, `NewSectionReader`,
  `MultiWriter`

----

# How many readers, writers are in the standard library?

* TODO: count

----

# How do you implement them?

```go
type Reader interface {
        Read(p []byte) (n int, err error)
}
```

```go
type Writer interface {
        Write(p []byte) (n int, err error)
}
```

----


