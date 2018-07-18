package util

import "io"

func rot13(b byte) byte {
	if (b >= 'A' && b < 'N') || (b >= 'a' && b < 'n') {
		return b + 13
	} else if (b >= 'M' && b < 'Z') || (b >= 'm' && b < 'z') {
		return b - 13
	} else {
		return b
	}
}

// Rot13Reader is a Reader that wraps another Reader, and applies the rot13
// function to the bytes it reads from the wrapped reader
type Rot13Reader struct {
	R io.Reader
}

func (r Rot13Reader) Read(bs []byte) (n int, err error) {
	n, err = r.R.Read(bs)
	for i := 0; i < n; i++ {
		bs[i] = rot13(bs[i])
	}
	return
}
