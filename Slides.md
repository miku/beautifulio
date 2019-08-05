# Beautiful IO

Golab 2019, 2019-10-21, Firenze
[Martin Czygan](mailto:martin.czygan@gmail.com)

----

# About me

SWE [@ubleipzig](https://ub.uni-leipzig.de) and web data engineer
[@internetarchive](https://archive.org) working mostly with Python and Go.

> [Explore IO](https://github.com/miku/exploreio) workshop at Golab 2017.

----

# The IO package

* contains basic, widely used interfaces (within and outside standard library)
* simple, flexible helpers

----

# Why beautiful?

> La bellezza è negli occhi di chi guarda

* small, versatile interfaces
* stackable
* reminiscent of UNIX pipes

----

# Love and praise

> This article aims to convince you to use io.Reader in your own code wherever
> you can. -- [@matryer](https://twitter.com/matryer)

> "Crossing Streams: a love letter to Go io.Reader" -- [@jmoiron](https://twitter.com/jmoiron)

> Which brings me to io.Reader, easily my favourite Go interface. --
> [@davecheney](https://twitter.com/davecheney)

----

# What's in io?

<!--
$ go doc io | grep ^type | wc -l
25
-->

* 25 types - of which 21 are interfaces, 7 functions, 3 constants, 6 errors

The four non-interfaces types are: `LimitedReader`, `PipeReader`, `PipeWriter`,
`SectionReader`. Functions: `Copy*`, `Pipe`, `ReadAtLeast`, `ReadFull`,
`WriteString`.

----

# How many readers are in the standard library?

* TODO: count

----

# How do you implement one?

```go
type Reader interface {
        Read(p []byte) (n int, err error)
}
```


