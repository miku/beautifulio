package copybench

import (
	"io"
	"io/ioutil"
	"testing"
)

const M = 1024 * 1024

func TestSimple(t *testing.T) {
	size := 100 * M
	f := RandFile(size)
	k, err := io.Copy(ioutil.Discard, f)
	if err != nil {
		t.Errorf("want %v, got %v", nil, err)
	}
	if k != int64(size) {
		t.Errorf("want %d, got %d", size, k)
	}
}
