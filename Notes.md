# Notes

## General things about speaking on technical topics

* keep it applied
* the good parts
* context and curation

## Key points

* work with various kinds of data
* like the concept of small interfaces
* there are love letters out there
* how to write your own
* how are they actually used, in which contexts
* can the IO interfaces teach us a bit about composition

## Code snippets

* [ ] interfaces
* [ ] CopyBuffer performance, and CPU
* [ ] ReaderFrom performance, and allocation difference

## Resources

### Crossings Streams

* https://www.datadoghq.com/blog/crossing-streams-love-letter-gos-io-reader/

The use of `ioutil.ReadAll` is a mistake.

* [ ] How often it is used? -- /home/tir/code/miku/ebba409208989306191926e238614f85
* [ ] Clone 300 repositories, count go files, analyze go files
* [ ] Also use Github BigQuery dataset

### BigQuery GitHub

* https://codelabs.developers.google.com/codelabs/bigquery-github/index.html?index=..%2F..index#0
* https://console.cloud.google.com/bigquery?project=golab-255608&folder&organizationId&p=bigquery-public-data&d=github_repos&t=languages&page=table

```sql
SELECT repo_name, language
FROM `bigquery-public-data.github_repos.languages` as ls, ls.language as language
WHERE language.name = 'go'
LIMIT 100
```

----

The buffersize of 4K is relatively good. Why? Is it the pagesize? Cacheline?

----

> Eat writes! // Eat writes.

----

Hiding certain interfaces, temporarily:

```go
type writerOnly struct {
    io.Writer
}
```

Interesting: the ReadFrom itself uses writerOnly, in order to not get stuck in an infinite loop.

----

Inquiring about a reader.

Reader hides lots of things, but it can be interrogated with type assertions: `srcIsRegularFile(src io.Reader)`

----

* Optimizing read from file to a network connection.

```go
func (r *response) ReadFrom(r io.Reader) (n int64, err error) ... for sendfile
```

----

Wrapper examples: http.connReader.

Random bit: "Fortunately, almost nothing uses HTTP/1.x pipelining"

----

Neat bufio Pool helper:

```go
newBufioReader
```

----

> expectContinueReader, wraps a readCloser - which on first read, sends an HTTP/1.1 100 Continue header

----

Random notes:

* writeHeader with status code > 1000 panics

----

From the "Live of a Writer" - "the writers are wired together like this ..." (~1500)

----

Ensuring interface compliance, via:

```go
type interface closeWriter{ CloseWrite() error }

var _ closeWriter = (*net.TCPConn)(nil)
```

----

* multipart_test.go

```go
slowReader // just reads one byte at a time. Instead of called 256 times for a MB, it's called a million times.
```

A sentinel reader, that returns EOF on condition, combined with a MultiReader.

----

In iotest, various reader implementations:

```
src/testing/iotest/reader.go
21:func (r *oneByteReader) Read(p []byte) (int, error) {
36:func (r *halfReader) Read(p []byte) (int, error) {
53:func (r *dataErrReader) Read(p []byte) (n int, err error) {
82:func (r *timeoutReader) Read(p []byte) (int, error) {
```

----

Always error or endless zeros.

```
143:func (sr *slowReader) Read(p []byte) (n int, err error) {
371:func (alwaysError) Read(p []byte) (int, error) {
389:func (endlessZeros) Read(p []byte) (int, error) {
```

----

* `atLeastReader`

```go
// atLeastReader reads from R, stopping with EOF once at least N bytes have been
// read. It is different from an io.LimitedReader in that it doesn't cut short
// the last Read call, and in that it considers an early EOF an error.
type atLeastReader struct {
        R io.Reader
        N int64
}

func (r *atLeastReader) Read(p []byte) (int, error) {
        if r.N <= 0 {
                return 0, io.EOF
        }
        n, err := r.R.Read(p)
        r.N -= int64(n) // won't underflow unless len(p) >= n > 9223372036854775809
        if r.N > 0 && err == io.EOF {
                return n, io.ErrUnexpectedEOF
        }
        if r.N <= 0 && err == nil {
                return n, io.EOF
        }
        return n, err
}
```

----

The Downcaser.

```go
type downCaser struct {
        t *testing.T
        r io.ByteReader
}

func (d *downCaser) ReadByte() (c byte, err error) {
        c, err = d.r.ReadByte()
        if c >= 'A' && c <= 'Z' {
                c += 'a' - 'A'
        }
        return
}

func (d *downCaser) Read(p []byte) (int, error) {
        d.t.Fatalf("unexpected Read call on downCaser reader")
        panic("unreachable")
}
```

----

* nTimes

```go
// nTimes is an io.Reader which yields the string s n times.
type nTimes struct {
        s   string
        n   int
        off int
}

func (r *nTimes) Read(p []byte) (n int, err error) {
        for {
                if r.n <= 0 || r.s == "" {
                        return n, io.EOF
                }
                n0 := copy(p, r.s[r.off:])
                p = p[n0:]
                n += n0
                r.off += n0
                if r.off == len(r.s) {
                        r.off = 0
                        r.n--
                }
                if len(p) == 0 {
                        return
                }
        }
}
```

----

Turning a io.Reader into a io.ReadSeeker - moderately efficient.

* readSeekerFromReader

----

From k8s, testing, a read delayer.

```go
type readDelayer struct {
        delay time.Duration
        io.ReadCloser
}

func (b *readDelayer) Read(p []byte) (n int, err error) {
        defer time.Sleep(b.delay)
        return b.ReadCloser.Read(p)
}
```

----

Rate limiting, subproject of k8s.

``` go
// Reader implements io.ReadCloser with a restriction on the rate of data
// transfer.
type Reader struct {
        io.Reader // Data source
        *Monitor  // Flow control monitor

        limit int64 // Rate limit in bytes per second (unlimited when <= 0)
        block bool  // What to do when no new bytes can be read due to the limit
}

// NewReader restricts all Read operations on r to limit bytes per second.
func NewReader(r io.Reader, limit int64) *Reader {
        return &Reader{r, New(0, 0), limit, true}
}

// Read reads up to len(p) bytes into p without exceeding the current transfer
// rate limit. It returns (0, nil) immediately if r is non-blocking and no new
// bytes can be read at this time.
func (r *Reader) Read(p []byte) (n int, err error) {
        p = p[:r.Limit(len(p), r.limit, r.block)]
        if len(p) > 0 {
                n, err = r.IO(r.Reader.Read(p))
        }
        return
}
...
```

----

* or just for Testing, FakeFile in k8s

----

A cacheing ReadCloser, for doing something with the content on EOF

```go
// cachingReadCloser is a wrapper around ReadCloser R that calls OnEOF
// handler with a full copy of the content read from R when EOF is
// reached.
type cachingReadCloser struct {
        // Underlying ReadCloser.
        R io.ReadCloser
        // OnEOF is called with a copy of the content of R when EOF is reached.
        OnEOF func(io.Reader)

        buf bytes.Buffer // buf stores a copy of the content of R.
}

// Read reads the next len(p) bytes from R or until R is drained. The
// return value n is the number of bytes read. If R has no data to
// return, err is io.EOF and OnEOF is called with a full copy of what
// has been read so far.
func (r *cachingReadCloser) Read(p []byte) (n int, err error) {
        n, err = r.R.Read(p)
        r.buf.Write(p[:n])
        if err == io.EOF {
                r.OnEOF(bytes.NewReader(r.buf.Bytes()))
        }
        return n, err
}

func (r *cachingReadCloser) Close() error {
        return r.R.Close()
}
```

----

Make sure you have a newline, sometimes important. From minio, s3select.

```go
// recordTransform will convert records to always have newline records.
type recordTransform struct {
        reader io.Reader
        // recordDelimiter can be up to 2 characters.
        recordDelimiter []byte
        oneByte         []byte
        useOneByte      bool
}
```

Just skip a number of bytes. SkipReader in minio, pkg/ioutil/ioutil.go.

* s3select, progressReader, counting data, scanner, read, ...

----

From moby project:

* pkg/chrootarchive/archive_test.go - slowEmptyTarReader

----

Benchmark with endless data; msgp

```go
package msgp

type timer interface {
        StartTimer()
        StopTimer()
}

// EndlessReader is an io.Reader
// that loops over the same data
// endlessly. It is used for benchmarking.
type EndlessReader struct {
        tb     timer
        data   []byte
        offset int
}

// NewEndlessReader returns a new endless reader
func NewEndlessReader(b []byte, tb timer) *EndlessReader {
        return &EndlessReader{tb: tb, data: b, offset: 0}
}

// Read implements io.Reader. In practice, it
// always returns (len(p), nil), although it
// fills the supplied slice while the benchmark
// timer is stopped.
func (c *EndlessReader) Read(p []byte) (int, error) {
        c.tb.StopTimer()
        var n int
        l := len(p)
        m := len(c.data)
        for n < l {
                nn := copy(p[n:], c.data[c.offset:])
                n += nn
                c.offset += nn
                c.offset %= m
        }
        c.tb.StartTimer()
        return n, nil
}
```

Writers have a few other areas, e.g. formatting.

* github/hub/vendor/github.com/kr/text/indent.go
