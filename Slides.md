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

# Praise and love

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
$ guru -json implements /usr/local/go/src/io/io.go:#3309,#3800
```

I counted over 200 implementations of each, io.Reader and io.Writer in the Go
tree and subrepositories. 

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

There is some flexibility around the end of a stream.

> Callers should always process the n > 0 bytes returned before considering the
error err. Doing so correctly handles I/O errors that happen after reading
some bytes and also both of the allowed EOF behaviors.

----

# Notes 

```go
type Reader interface {
        Read(p []byte) (n int, err error)
}
```

* The byte slice is under the control of the caller.

> Implementations must not retain p.

This hints at the streaming nature of this interface.

----

# Implementations

* files
* network connections
* HTTP response bodies
* standard input and output
* compression
* hashing
* encoding
* formatting
* ...

Many uses in testing as well.

----

# Structural typing

* conversions are not required, a file implements `Read` and hence *is* a
  *io.Reader*

![](static/filetoreader.png)

----

# Streams

As layed out in the *love letter*, the use of `ioutil.ReadAll` is debatable.
It's in the standard library and useful, but not always necessary.

```go
b, err := ioutil.ReadAll(r)
...
```

----

# Streams

* you may lose the advantage to use the `Reader` in other places
* you may consume more memory

> Streams can trivially produce infinite output while using barely any memory at
> all - imagine an implementation behaving like /dev/zero or /dev/urandom.

* Memory control is an important advantage.

----

# Follow the stream

Instead of writing:

```go
b, _ := ioutil.ReadAll(resp.Body) // Pressure on memory.
fmt.Println(string(b))
```

You may want to connect streams:

```go
_, _ = io.Copy(os.Stdout, resp.Body)
```

----

# Stream advantages

* memory efficient
* can work with data, that does not fit in memory
* allows to work on different protocol parts differently (e.g. HTTP header vs HTTP body)

----

# Another example

We often need to unmarshal JSON.

```go
_ = json.Unmarshal(data, &v) // data might come from ioutil.ReadAll(resp.Body)
```

But we can decode it as well.

```go
_ = json.NewDecoder(resp.Body).Decode(&v)
```

In this case, the JSON data must be fully read, so this is a weak example.

----

# Glipse at composition

But what is we want need to preprocess the data, e.g. decompress it. Streams
compose well.

```go
zr, _ = gzip.NewReader(resp.Body)
_ json.NewDecoder(zr).Decode(&)
```

----

# How do you implement one yourself?

You only need a `Read` method with the correct signature.

* Example: `/dev/zero`

```go
type devZero struct{}

func (r *devZero) Read(p []byte) (int, error) {
        for i := 0; i < len(p); i++ {
                p[i] = '\x00'
        }
        return len(p), nil
}
```

This is already an infinite stream.

----

# Embed a reader

Often you want to transform a given data stream, so you embed it.

```go
type UpperReader struct {
	r io.Reader // Underlying stream
}

func (r *UpperReader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	copy(p, bytes.ToUpper(p))
	return n, err
}

func main() {
	if _, err := io.Copy(os.Stdout, &UpperReader{os.Stdin}); err != nil {
		log.Fatal(err)
	}
}
```

* Also try: https://tour.golang.org/methods/22 (Reader exercise, ROT13)

----

# The io.Writer interface

Analogous to the `io.Reader` interface.

```go
type Writer interface {
        Write(p []byte) (n int, err error)
}
```

> Write writes len(p) bytes from p to the underlying data stream. It returns the
> number of bytes written from p (0 <= n <= len(p)) and any error encountered
> that caused the write to stop early.

> Write must return a non-nil error if it returns n < len(p). Write must not
> modify the slice data, even temporarily.

As with readers:

> Implementations must not retain p. 

----

# An example

A writer that does not much, but is still useful - `/dev/null` in Go:

```go
type devNull struct{}

func (w *devNull) Write(p []byte) (int, error) {
	return len(p), nil
}

func main() {
	if n, err := io.Copy(&devNull{}, strings.NewReader("Hello World")); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("%d bytes copied", n)
	}
}
```

The standard library implementation is called `ioutil.Discard` (for an
interesting/frustrating bug related to ioutil.Discard, read
[#4589](https://github.com/golang/go/issues/4589)).

----

# Use case: File

Prototypical stream: A file.

* os.File

And alternatives and substitutions, e.g. dummy files for tests or file that
support atomic writes.

----

# Historical note

![](static/attunix.png)


> A file is simply a sequence of bytes. Its main attribute is its size. By
> contrast, on more conventional systems, a file has a dozen or so attributes.
> To specify and create a file it takes endless amount of chit-chat. If you are
> on a UNIX system you can simply ask for a file and use it interchangeble
> whereever you want a file. (XXX: Unix documentary)

If a file is just a sequence of bytes, more things will look like files.

----

# Use case: Networking

```go
type Conn interface {
        // Read reads data from the connection.
        // Read can be made to time out and return an Error with Timeout() == true
        // after a fixed time limit; see SetDeadline and SetReadDeadline.
        Read(b []byte) (n int, err error)
        ...
```

----

