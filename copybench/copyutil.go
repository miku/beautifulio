package copybench

import (
	"crypto/rand"
	"io"
)

type randFile struct {
	Size int
	I    int
	r    io.Reader
}

func RandFile(size int) *randFile {
	return &randFile{
		Size: size,
		I:    0,
		r:    rand.Reader,
	}
}

// Read reads random bytes.
func (f *randFile) Read(p []byte) (int, error) {
	if f.I >= f.Size {
		return 0, io.EOF
	}
	gap := f.Size - f.I
	if gap < len(p) {
		f.I += gap
		return f.r.Read(p[:gap])
	}
	f.I += len(p)
	return f.r.Read(p)
}
