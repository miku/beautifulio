# FIM

Find implemententations.

Usage:

```
$ fim io.Reader
```

```
/usr/local/go/src/io/io.go|77 col 6| interface type io.Reader
/usr/local/go/src/bufio/bufio.go|31 col 6| is implemented by pointer type *bufio.Reader
/usr/local/go/src/bytes/buffer.go|20 col 6| is implemented by pointer type *bytes.Buffer
/usr/local/go/src/bytes/reader.go|18 col 6| is implemented by pointer type *bytes.Reader
/usr/local/go/src/compress/flate/inflate.go|267 col 6| is implemented by pointer type *compress/flate.decompressor
/usr/local/go/src/compress/gzip/gunzip.go|74 col 6| is implemented by pointer type *compress/gzip.Reader
/usr/local/go/src/compress/zlib/reader.go|46 col 6| is implemented by pointer type *compress/zlib.reader
/usr/local/go/src/debug/elf/reader.go|37 col 6| is implemented by pointer type *debug/elf.readSeekerFromReader
/usr/local/go/src/encoding/base64/base64.go|386 col 6| is implemented by pointer type *encoding/base64.decoder
/usr/local/go/src/encoding/base64/base64.go|570 col 6| is implemented by pointer type *encoding/base64.newlineFilteringReader
/usr/local/go/src/encoding/gob/decode.go|39 col 6| is implemented by pointer type *encoding/gob.decBuffer
/usr/local/go/src/encoding/json/encode.go|278 col 6| is implemented by pointer type *encoding/json.encodeState
/usr/local/go/src/fmt/scan.go|157 col 6| is implemented by pointer type *fmt.ss
/usr/local/go/src/fmt/scan.go|84 col 6| is implemented by pointer type *fmt.stringReader
/home/tir/go/pkg/mod/golang.org/x/tools@v0.0.0-20191018212557-ed542cd5b28a/go/internal/gcimporter/iexport.go|659 col 6| is implemented by pointer type *golang.org/x/tools/go/internal/gcimporter.intWriter
/usr/local/go/src/internal/poll/fd_unix.go|18 col 6| is implemented by pointer type *internal/poll.FD
/usr/local/go/src/io/io.go|436 col 6| is implemented by pointer type *io.LimitedReader
/usr/local/go/src/io/pipe.go|117 col 6| is implemented by pointer type *io.PipeReader
/usr/local/go/src/io/io.go|461 col 6| is implemented by pointer type *io.SectionReader
/usr/local/go/src/io/multi.go|13 col 6| is implemented by pointer type *io.multiReader
/usr/local/go/src/io/pipe.go|32 col 6| is implemented by pointer type *io.pipe
/usr/local/go/src/io/io.go|529 col 6| is implemented by pointer type *io.teeReader
/usr/local/go/src/math/rand/rand.go|51 col 6| is implemented by pointer type *math/rand.Rand
/usr/local/go/src/os/types.go|16 col 6| is implemented by pointer type *os.File
/usr/local/go/src/strings/reader.go|17 col 6| is implemented by pointer type *strings.Reader
/usr/local/go/src/bufio/bufio.go|755 col 6| is implemented by struct type bufio.ReadWriter
/usr/local/go/src/debug/elf/reader.go|13 col 6| is implemented by struct type debug/elf.errorReader
/usr/local/go/src/go/internal/gcimporter/iimport.go|21 col 6| is implemented by struct type go/internal/gcimporter.intReader
/home/tir/go/pkg/mod/golang.org/x/tools@v0.0.0-20191018212557-ed542cd5b28a/go/internal/gcimporter/iimport.go|23 col 6| is implemented by struct type golang.org/x/tools/go/internal/gcimporter.intReader
/usr/local/go/src/io/multi.go|7 col 6| is implemented by struct type io.eofReader
/usr/local/go/src/io/ioutil/ioutil.go|110 col 6| is implemented by struct type io/ioutil.nopCloser
/usr/local/go/src/math/big/intconv.go|217 col 6| is implemented by struct type math/big.byteReader
/usr/local/go/src/os/exec/exec.go|589 col 6| is implemented by struct type os/exec.closeOnce
/usr/local/go/src/compress/flate/inflate.go|261 col 6| is implemented by interface type compress/flate.Reader
/usr/local/go/src/fmt/scan.go|21 col 6| is implemented by interface type fmt.ScanState
/usr/local/go/src/io/io.go|126 col 6| is implemented by interface type io.ReadCloser
/usr/local/go/src/io/io.go|145 col 6| is implemented by interface type io.ReadSeeker
/usr/local/go/src/io/io.go|138 col 6| is implemented by interface type io.ReadWriteCloser
/usr/local/go/src/io/io.go|157 col 6| is implemented by interface type io.ReadWriteSeeker
/usr/local/go/src/io/io.go|120 col 6| is implemented by interface type io.ReadWriter
```
